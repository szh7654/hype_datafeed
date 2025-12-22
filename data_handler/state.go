package datahandler

import (
	"context"
	dr "indexer/data_retriver"
	"indexer/util"
	"sync"
	"time"

	"github.com/apache/arrow/go/v18/arrow/memory"
	eth "github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var pool = memory.NewGoAllocator()

func init() {
	promauto.NewGaugeFunc(
		prometheus.GaugeOpts{Name: "user_count"},
		func() float64 {
			return float64(len(Users))
		})

	promauto.NewGaugeFunc(
		prometheus.GaugeOpts{Name: "asset_count"},
		func() float64 {
			return float64(len(Assets))
		})

	promauto.NewGaugeFunc(
		prometheus.GaugeOpts{Name: "position_count"},
		func() float64 {
			return float64(len(Positions))
		})

	promauto.NewGaugeFunc(
		prometheus.GaugeOpts{Name: "active_user_count"},
		func() float64 {
			return float64(len(activeUsers))
		})

	promauto.NewGaugeFunc(
		prometheus.GaugeOpts{Name: "fill_block_number"},
		func() float64 {
			return float64(FillBlockNumber)
		})

	promauto.NewGaugeFunc(
		prometheus.GaugeOpts{Name: "order_status_block_number"},
		func() float64 {
			return float64(OrderStatusBlockNumber)
		})

	promauto.NewGaugeFunc(
		prometheus.GaugeOpts{Name: "order_book_diff_block_number"},
		func() float64 {
			return float64(OrderBookDiffBlockNumber)
		})

	promauto.NewGaugeFunc(
		prometheus.GaugeOpts{Name: "block_idle_time"},
		func() float64 {
			return float64(time.Now().Sub(LatestBlockTime).Milliseconds())
		})

}

func preStart() {
	users := util.ExtractUser()
	InitUsers(users)
	metaAndAssetCtxs, err := dr.FetchMetaAndAssetCtxs()
	if err != nil {
		panic(err)
	}
	InitAssets(*metaAndAssetCtxs)

	l4Snapshot := dr.FetchL4Snapshot()
	applyL4Snapshot(l4Snapshot)
}

func Start(ctx context.Context, wg *sync.WaitGroup) {
	preStart()
	ueserStateReqTicker := time.NewTicker(1 * time.Minute)
	minuteSnapshotTicker := time.NewTicker(1 * time.Minute)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return

			case <-ueserStateReqTicker.C:
				sendAndResetActiveUser()

			case <-minuteSnapshotTicker.C:
				generateMinuteSnapshot()

			case UserStateWithId := <-dr.UserStateResponseChan:
				applyUserState(UserStateWithId.UserState, UserStateWithId.UserId)

			case blockfill := <-dr.BlockFillChan:
				applyBlockFill(blockfill)
			case blockOrderStatus := <-dr.BlockOrderStatusChan:
				applyBlockOrderStatus(blockOrderStatus)

			case blockOrderBookDiff := <-dr.BlockOrderBookDiffChan:
				applyBlockOrderBookDiff(blockOrderBookDiff)
			}
		}
	}()

}

func sendAndResetActiveUser() {
	for userId, addr := range activeUsers {
		dr.UserStateWorkerChan <- dr.AddrWithUserId{
			UserId:  userId,
			Address: addr,
		}
	}
	activeUsers = make(map[uint32]eth.Address, len(activeUsers))
}

func applyUserState(userState *dr.UserState, userId uint32) {
	user := Users[userId]
	user.AccountValue = userState.MarginSummary.AccountValue
	user.TotalNtlPos = userState.MarginSummary.TotalNtlPos

	user.CrossAccountValue = userState.CrossMarginSummary.AccountValue

	// TODO: 补全liquidationPx

	for _, position := range userState.AssetPositions {
		if assetId, ok := NameToAssetId[position.Position.Coin]; ok {
			positionId := PositionId(assetId, userId)
			Positions[positionId] = Position{
				Szi:           position.Position.Szi,
				EntryPx:       float64(position.Position.EntryPx),
				Isolated:      position.Position.Leverage.Type == dr.LeverageTypeIsolated,
				OneWay:        position.PositionType == dr.PositionTypeOneWay,
				MarginUsed:    position.Position.MarginUsed,
				Leverage:      uint8(position.Position.Leverage.Value),
				MaxLeverage:   uint8(position.Position.MaxLeverage),
				UnrealizedPnl: position.Position.UnrealizedPnl,
				LiquidationPx: float64(position.Position.LiquidationPx),
			}
			activeUsers[userId] = UserAddrs[userId]
		}
	}
}

var (
	FillBlockNumber          uint32    = 0
	OrderStatusBlockNumber   uint32    = 0
	OrderBookDiffBlockNumber uint32    = 0
	LatestBlockTime          time.Time = time.Time{}
	blockLatency                       = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "block_latency",
		Buckets: prometheus.DefBuckets, // 或者自定义 []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}
	})
)

func applyBlockFill(blockfill dr.BlockFill) {
	FillBlockNumber = blockfill.BlockNumber
	latency := time.Time(blockfill.BlockTime).Sub(time.Time(blockfill.LocalTime))
	blockLatency.Observe(latency.Seconds())
	LatestBlockTime = util.MaxTime(LatestBlockTime, time.Time(blockfill.BlockTime))
	for _, fillEvent := range blockfill.Events {
		user, userId := GetUser(fillEvent.Address)
		activeUsers[userId] = fillEvent.Address
		user.Pnl += fillEvent.Fill.ClosedPnl
	}
}

func applyBlockOrderStatus(blockOrderStatus dr.BlockOrderStatus) {
	OrderStatusBlockNumber = blockOrderStatus.BlockNumber
	for _, orderStatus := range blockOrderStatus.Events {
		_, userId := GetUser(orderStatus.User)
		activeUsers[userId] = orderStatus.User
		order := orderStatus.Order
		asset, assetId := GetAsset(order.Coin)
		if asset != nil {
			book := books[assetId]
			switch orderStatus.Order.IsTrigger {
			case true:
				switch orderStatus.Status {
				case dr.StatusOpen:
					switch order.Side {
					case dr.SideA:
						book.TriggerAsks.applyLevelAction(LevelAction{Px: order.TriggerPx, Sz: order.Sz, Oid: order.Oid, UserId: userId, ActionType: dr.ActionTypeNew})
					case dr.SideB:
						book.TriggerBids.applyLevelAction(LevelAction{Px: order.TriggerPx, Sz: order.Sz, Oid: order.Oid, UserId: userId, ActionType: dr.ActionTypeNew})
					}					
				default:
					switch order.Side {
					case dr.SideA:
						book.TriggerAsks.applyLevelAction(LevelAction{Px: order.TriggerPx, Sz: order.Sz, Oid: order.Oid, UserId: userId, ActionType: dr.ActionTypeRemove})
					case dr.SideB:
						book.TriggerBids.applyLevelAction(LevelAction{Px: order.TriggerPx, Sz: order.Sz, Oid: order.Oid, UserId: userId, ActionType: dr.ActionTypeRemove})
					}
				}
			}
		}
	}
}

func applyBlockOrderBookDiff(blockOrderBookDiff dr.BlockOrderBookDiff) {
	OrderBookDiffBlockNumber = blockOrderBookDiff.BlockNumber
	for _, orderBookDiff := range blockOrderBookDiff.Events {
		_, userId := GetUser(orderBookDiff.User)
		activeUsers[userId] = orderBookDiff.User
		asset, assetId := GetAsset(orderBookDiff.Coin)
		if asset != nil {
			book := books[assetId]
			switch orderBookDiff.Side {
			case dr.SideA:
				book.LimitAsks.applyOrderBookDiff(userId, orderBookDiff)
			case dr.SideB:
				book.LimitBids.applyOrderBookDiff(userId, orderBookDiff)
			}
		}
	}
}

func applyL4Snapshot(l4Snapshot *dr.L4SnapShot) {
	OrderBookDiffBlockNumber = l4Snapshot.BlockNumber
	for _, assetOrders := range l4Snapshot.AssetOrders {
		if id, ok := NameToAssetId[assetOrders.Name]; ok {
			book := books[id]
			book.applyAsks(assetOrders.AskOrders)
			book.applyBids(assetOrders.BidOrders)
		}
	}
}

func generateMinuteSnapshot() {
	panic("unimplemented")
}
