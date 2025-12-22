package dataretriver

import (
	"context"
	"indexer/util"
	"sync"
	"time"

	"github.com/bytedance/sonic"
	eth "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

var (
	BlockFillChan               = make(chan BlockFill, 1000)
	BlockOrderStatusChan        = make(chan BlockOrderStatus, 1000)
	BlockOrderBookDiffChan      = make(chan BlockOrderBookDiff, 1000)
	UserAndPositionSnapshotChan = make(chan map[uint64]Position, 1000)

	UserStateWorkerChan = make(chan AddrWithUserId, 100000)

	UserStateResponseChan = make(chan UserStateWithId, 1000)
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
}

func fetchUserStateRoutine(ctx context.Context, wg *sync.WaitGroup) {
	for range 5 {
		go fetchUserStateWorker(ctx, wg)
	}
}

func fetchAssetRoutine(ctx context.Context, wg *sync.WaitGroup) {
	go func() {

	}()
}

func fetchUserStateWorker(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
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
	tailer := util.NewTailer("/root/hl/data/node_orders_status_by_block/hourly/", ctx, wg)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
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
	tailer := util.NewTailer("/root/hl/data/node_orders_book_diff_by_block/hourly/", ctx, wg)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
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
