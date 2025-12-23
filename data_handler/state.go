package datahandler

import (
	"context"
	dr "indexer/data_retriver"
	"indexer/util"
	"sync"
	"time"

	"github.com/apache/arrow/go/v18/arrow/memory"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog/log"
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
		prometheus.GaugeOpts{Name: "block_fill_number"},
		func() float64 {
			return float64(FillBlockNumber)
		})

	promauto.NewGaugeFunc(
		prometheus.GaugeOpts{Name: "block_order_status_number"},
		func() float64 {
			return float64(OrderStatusBlockNumber)
		})

	promauto.NewGaugeFunc(
		prometheus.GaugeOpts{Name: "block_order_book_diff_number"},
		func() float64 {
			return float64(OrderBookDiffBlockNumber)
		})

	promauto.NewGaugeFunc(
		prometheus.GaugeOpts{Name: "block_idle_time"},
		func() float64 {
			return float64(time.Since(LatestBlockTime).Milliseconds())
		})

}

func Start(ctx context.Context, wg *sync.WaitGroup) {
	ueserStateReqTicker := time.NewTicker(1 * time.Minute)
	minuteSnapshotTicker := time.NewTicker(1 * time.Minute)

	go func() {
		defer wg.Done()
		log.Info().Msg("start data handler loop")
		for {
			select {
			case <-ctx.Done():
				log.Info().Msg("data handler loop exit")
				return
			case l4snapshot := <-dr.L4SnapShotChan:
				log.Info().Msg("l4 snapshot received")
				ApplyL4Snapshot(l4snapshot)
				log.Info().Msg("l4 snapshot applied")
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
				applyBlockOrderBookDiffPre(blockOrderBookDiff)
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
	clear(activeUsers)
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
		applyFillEvent(fillEvent)

	}
}

func applyFillEvent(fillEvent dr.FillEvent) {
	user, userId := GetUser(fillEvent.Address)
	activeUsers[userId] = fillEvent.Address
	user.Pnl += fillEvent.Fill.ClosedPnl
	if fillEvent.Fill.Liquidation != nil {
		_, liquidationUserId := GetUser(fillEvent.Fill.Liquidation.LiquidationUser)
		asset, assetId := GetAsset(fillEvent.Fill.Coin)
		if asset != nil {
			if userId != liquidationUserId {
				ProviderLiquidations = append(ProviderLiquidations, Liquidation{
					UserId:        userId,
					AssetId:       assetId,
					Szi:           fillEvent.Fill.Sz,
					MarkPx:        fillEvent.Fill.Liquidation.MarkPx,
					PositionValue: fillEvent.Fill.Px * fillEvent.Fill.Sz,
					ClosedPnl:     fillEvent.Fill.ClosedPnl,
				})
			} else {
				UnderLiquidations = append(UnderLiquidations, Liquidation{
					UserId:        userId,
					AssetId:       assetId,
					Szi:           fillEvent.Fill.Sz,
					MarkPx:        fillEvent.Fill.Liquidation.MarkPx,
					PositionValue: fillEvent.Fill.Px * fillEvent.Fill.Sz,
					ClosedPnl:     fillEvent.Fill.ClosedPnl,
				})
			}
		}

	}
}

func applyBlockOrderStatus(blockOrderStatus dr.BlockOrderStatus) {
	OrderStatusBlockNumber = blockOrderStatus.BlockNumber
	for _, orderStatus := range blockOrderStatus.Events {
		_, userId := GetUser(orderStatus.User)
		activeUsers[userId] = orderStatus.User
		order := orderStatus.Order
		asset, _ := GetAsset(order.Coin)
		book := asset.Book
		switch orderStatus.Order.IsTrigger {
		case true:
			switch orderStatus.Status {
			case dr.StatusOpen:
				switch order.Side {
				case dr.SideA:
					err := book.TriggerAsks.applyLevelAction(LevelAction{Px: order.TriggerPx, Sz: order.Sz, Oid: order.Oid, UserId: userId, ActionType: dr.ActionTypeNew})
					if err != nil {
						log.Error().Err(err).Msg("apply ask open trigger order failed")
					}
				case dr.SideB:
					err := book.TriggerBids.applyLevelAction(LevelAction{Px: order.TriggerPx, Sz: order.Sz, Oid: order.Oid, UserId: userId, ActionType: dr.ActionTypeNew})
					if err != nil {
						log.Error().Err(err).Msg("apply bid open trigger order failed")
					}
				}
			default:
				switch order.Side {
				case dr.SideA:
					err := book.TriggerAsks.applyLevelAction(LevelAction{Px: order.TriggerPx, Sz: order.Sz, Oid: order.Oid, UserId: userId, ActionType: dr.ActionTypeRemove})
					if err != nil {
						log.Error().Err(err).Msgf("apply ask close trigger order failed, orderstatus: %s", orderStatus.Status.String())
					}
				case dr.SideB:
					err := book.TriggerBids.applyLevelAction(LevelAction{Px: order.TriggerPx, Sz: order.Sz, Oid: order.Oid, UserId: userId, ActionType: dr.ActionTypeRemove})
					if err != nil {
						log.Error().Err(err).Msgf("apply bid close trigger order failed, orderstatus: %s", orderStatus.Status.String())
					}
				}
			}
		}
		for _, order := range orderStatus.Order.Children {
			switch order.IsTrigger {
			case true:
				switch orderStatus.Status {
				case dr.StatusOpen:
					switch order.Side {
					case dr.SideA:
						err := book.TriggerAsks.applyLevelAction(LevelAction{Px: order.TriggerPx, Sz: order.Sz, Oid: order.Oid, UserId: userId, ActionType: dr.ActionTypeNew})
						if err != nil {
							log.Error().Err(err).Msg("apply ask open trigger order failed")
						}
					case dr.SideB:
						err := book.TriggerBids.applyLevelAction(LevelAction{Px: order.TriggerPx, Sz: order.Sz, Oid: order.Oid, UserId: userId, ActionType: dr.ActionTypeNew})
						if err != nil {
							log.Error().Err(err).Msg("apply bid open trigger order failed")
						}
					}
				default:
					switch order.Side {
					case dr.SideA:
						err := book.TriggerAsks.applyLevelAction(LevelAction{Px: order.TriggerPx, Sz: order.Sz, Oid: order.Oid, UserId: userId, ActionType: dr.ActionTypeRemove})
						if err != nil {
							log.Error().Err(err).Msgf("apply ask close trigger order failed, orderstatus: %s", orderStatus.Status.String())
						}
					case dr.SideB:
						err := book.TriggerBids.applyLevelAction(LevelAction{Px: order.TriggerPx, Sz: order.Sz, Oid: order.Oid, UserId: userId, ActionType: dr.ActionTypeRemove})
						if err != nil {
							log.Error().Err(err).Msgf("apply bid close trigger order failed, orderstatus: %s", orderStatus.Status.String())
						}
					}
				}
			}
		}

	}
}

var (
	l4SnapshotBlockNumber  uint32 = 0
	preBlockOrderBookDiffs        = []dr.BlockOrderBookDiff{}
	isSync                        = false
)

func applyBlockOrderBookDiffPre(blockOrderBookDiff dr.BlockOrderBookDiff) {
	if isSync {
		applyBlockOrderBookDiff(blockOrderBookDiff)
		return
	}
	if l4SnapshotBlockNumber == 0 { // 尚未读取snapshot
		log.Info().Msgf("l4SnapshotBlockNumber=0, append preBlockOrderBookDiffs, blockNumber: %d, delay: %s", blockOrderBookDiff.BlockNumber, time.Since(time.Time(blockOrderBookDiff.BlockTime)))
		preBlockOrderBookDiffs = append(preBlockOrderBookDiffs, blockOrderBookDiff)
	} else {
		log.Info().Msgf("l4SnapshotBlockNumber>0, apply preBlockOrderBookDiffs, preBlockOrderBookDiffs range: %d ~ %d", preBlockOrderBookDiffs[0].BlockNumber, preBlockOrderBookDiffs[len(preBlockOrderBookDiffs)-1].BlockNumber)
		for _, preBlockOrderBookDiff := range preBlockOrderBookDiffs {
			if preBlockOrderBookDiff.BlockNumber > l4SnapshotBlockNumber {
				applyBlockOrderBookDiff(preBlockOrderBookDiff)
			}
		}
		if blockOrderBookDiff.BlockNumber > l4SnapshotBlockNumber {
			applyBlockOrderBookDiff(blockOrderBookDiff)
		}
		preBlockOrderBookDiffs = nil
		isSync = true
	}
}

func applyBlockOrderBookDiff(blockOrderBookDiff dr.BlockOrderBookDiff) {
	log.Info().Msgf("applying block order book diff, blockNumber: %d", blockOrderBookDiff.BlockNumber)
	OrderBookDiffBlockNumber = blockOrderBookDiff.BlockNumber
	for _, orderBookDiff := range blockOrderBookDiff.Events {
		_, userId := GetUser(orderBookDiff.User)
		setActive(userId, orderBookDiff.User)
		asset, _ := GetAsset(orderBookDiff.Coin)
		if asset != nil {
			book := asset.Book
			switch orderBookDiff.Side {
			case dr.SideA:
				err := book.LimitAsks.applyOrderBookDiff(userId, orderBookDiff)
				if err != nil {
					log.Error().Err(err).Msgf("apply ask bookdiff error, asset: %s userId: %d", orderBookDiff.Coin, userId)
				}
			case dr.SideB:
				err := book.LimitBids.applyOrderBookDiff(userId, orderBookDiff)
				if err != nil {
					log.Error().Err(err).Msgf("apply bid bookdiff error, asset: %s userId: %d", orderBookDiff.Coin, userId)
				}
			}
		}
	}
}

func ApplyL4Snapshot(l4Snapshot *dr.L4SnapShot) {
	log.Info().Msgf("applying l4 snapshot, blockNumber: %d", l4Snapshot.BlockNumber)
	l4SnapshotBlockNumber = l4Snapshot.BlockNumber
	for _, assetSnapShot := range l4Snapshot.AssetSnapShots {
		asset, _ := GetAsset(assetSnapShot.Name)
		book := asset.Book
		bookOrders := assetSnapShot.BookOrdersAndUntriggeredOrders.BookOrders
		untriggerredOrders := assetSnapShot.BookOrdersAndUntriggeredOrders.UntriggeredOrders
		book.LimitBids.applyaLimitAddrOrders(bookOrders.AskOrders)
		book.LimitAsks.applyaLimitAddrOrders(bookOrders.BidOrders)
		for _, untriggeredOrder := range untriggerredOrders {
			switch untriggeredOrder.Order.Side {
			case dr.SideA:
				book.TriggerAsks.applyTriggerAddrOrder(untriggeredOrder)
			case dr.SideB:
				book.TriggerBids.applyTriggerAddrOrder(untriggeredOrder)
			}
		}
	}
}

func generateMinuteSnapshot() {
	log.Info().Msg("Generating minute snapshot")
}
