package dataretriver

import (
	"context"
	"indexer/util"
	"sync"
	"time"

	"github.com/bytedance/sonic"
	eth "github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog/log"
)

var (
	BlockFillChan               = make(chan BlockFill, 1000)
	BlockOrderStatusChan        = make(chan BlockOrderStatus, 1000)
	BlockOrderBookDiffChan      = make(chan BlockOrderBookDiff, 1000)
	UserAndPositionSnapshotChan = make(chan map[uint64]Position, 1000)

	UserStateWorkerChan = make(chan AddrWithUserId, 100000)

	UserStateResponseChan = make(chan UserStateWithId, 1000)

	L4SnapShotChan = make(chan *L4SnapShot, 2)

	// 使用 NewGaugeFunc 创建通道背压监控指标
	_ = promauto.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name:        "channel_queue_length",
			ConstLabels: map[string]string{"channel_name": "BlockFillChan"},
		},
		func() float64 {
			return float64(len(BlockFillChan))
		},
	)
	_ = promauto.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name:        "channel_queue_length",
			ConstLabels: map[string]string{"channel_name": "BlockOrderStatusChan"},
		},
		func() float64 {
			return float64(len(BlockOrderStatusChan))
		},
	)
	_ = promauto.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name:        "channel_queue_length",
			ConstLabels: map[string]string{"channel_name": "BlockOrderBookDiffChan"},
		},
		func() float64 {
			return float64(len(BlockOrderBookDiffChan))
		},
	)
	_ = promauto.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name:        "channel_queue_length",
			ConstLabels: map[string]string{"channel_name": "UserAndPositionSnapshotChan"},
		},
		func() float64 {
			return float64(len(UserAndPositionSnapshotChan))
		},
	)
	_ = promauto.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name:        "channel_queue_length",
			ConstLabels: map[string]string{"channel_name": "UserStateWorkerChan"},
		},
		func() float64 {
			return float64(len(UserStateWorkerChan))
		},
	)
	_ = promauto.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name:        "channel_queue_length",
			ConstLabels: map[string]string{"channel_name": "UserStateResponseChan"},
		},
		func() float64 {
			return float64(len(UserStateResponseChan))
		},
	)
)

type UserStateWithId struct {
	UserId    uint32
	UserState *UserState
}
type AddrWithUserId struct {
	Address eth.Address
	UserId  uint32
}

func Start(ctx context.Context, wg *sync.WaitGroup) {
	fetchUserStateRoutine(ctx, wg)
	fetchAssetRoutine(ctx, wg)
	readFillBlockRoutine(ctx, wg)
	readOrderStatusBlockRoutine(ctx, wg)
	readOrderBookDiffBlockRoutine(ctx, wg)

	l4Snapshot := FetchL4Snapshot()
	L4SnapShotChan <- l4Snapshot
	log.Info().Msg("L4SnapShot sent")
}

func fetchUserStateRoutine(ctx context.Context, wg *sync.WaitGroup) {
	for range 5 {
		go fetchUserStateWorker(ctx, wg)
	}
}

func fetchAssetRoutine(ctx context.Context, wg *sync.WaitGroup) {
	go func() {
		//TODO: Implement
	}()
}

func fetchUserStateWorker(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("fetchUserStateWorker stopped")
			return
		case addressWithUserId := <-UserStateWorkerChan:
			userState, err := FetchUserState(addressWithUserId.Address.String())
			if err != nil {
				log.Error().Err(err).Send()
				time.Sleep(100 * time.Millisecond)
				UserStateWorkerChan <- addressWithUserId // retry
				continue
			}
			UserStateResponseChan <- UserStateWithId{
				UserId:    addressWithUserId.UserId,
				UserState: userState,
			}
		}
	}
}

func readFillBlockRoutine(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	tailer := util.NewTailer("/root/hl/data/node_fills_by_block/hourly/", ctx, wg)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				log.Info().Msg("readFillBlockRoutine stopped")
				return
			case line := <-tailer.Lines:
				var blockFill BlockFill
				err := sonic.Unmarshal([]byte(line), &blockFill)
				if err != nil {
					log.Error().Err(err).Send()
					continue
				}
				BlockFillChan <- blockFill
			}
		}
	}()
}

func readOrderStatusBlockRoutine(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	tailer := util.NewTailer("/root/hl/data/node_order_statuses_by_block/hourly/", ctx, wg)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				log.Info().Msg("readOrderStatusBlockRoutine stopped")
				return
			case line := <-tailer.Lines:
				var blockOrderStatus BlockOrderStatus
				err := sonic.UnmarshalString(line, &blockOrderStatus)
				if err != nil {
					log.Error().Err(err).Send()
					continue
				}
				BlockOrderStatusChan <- blockOrderStatus
			}
		}
	}()
}

func readOrderBookDiffBlockRoutine(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	tailer := util.NewTailer("/root/hl/data/node_raw_book_diffs_by_block/hourly/", ctx, wg)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				log.Info().Msg("readOrderBookDiffBlockRoutine stopped")
				return
			case line := <-tailer.Lines:
				var blockOrderBookDiff BlockOrderBookDiff
				err := sonic.UnmarshalString(line, &blockOrderBookDiff)
				if err != nil {
					log.Error().Err(err).Send()
					continue
				}
				BlockOrderBookDiffChan <- blockOrderBookDiff
			}
		}
	}()
}
