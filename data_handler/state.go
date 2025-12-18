package datahandler

import (
	"context"
	dr "indexer/data_retriver"
	"indexer/util"
	"sync"
	"time"

	"github.com/apache/arrow/go/v18/arrow/memory"
)

var pool = memory.NewGoAllocator()

func Start(ctx context.Context, wg *sync.WaitGroup) {

	users := util.ExtractUser()
	InitUsers(users)
	client := dr.NewClient(dr.LocalAPIURL, false)
	metaAndAssetCtxs, err := client.MetaAndAssetCtxs()
	if err != nil {
		panic(err)
	}
	InitAssets(*metaAndAssetCtxs)

	ueserStateReqTicker := time.NewTicker(1 * time.Minute)
	minuteSnapshotTicker := time.NewTicker(1 * time.Minute)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ueserStateReqTicker.C:
				for userId, addr := range getActivePositionUsers() {
					dr.UserStateWorkerChan <- dr.AddressWithUserId{
						UserId:  userId,
						Address: addr,
					}
				}

			case <-minuteSnapshotTicker.C:
				generateMinuteSnapshot()

			case UserStateWithId := <-dr.UserStateResponseChan:
				applyUserState(UserStateWithId.UserState, UserStateWithId.UserId)

			case blockfill := <-dr.BlockFillChan:
				applyBlockFill(blockfill)
			}
		}
	}()

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
		}
	}
}

func applyBlockFill(blockfill dr.BlockFill) {
	panic("unimplemented")
}

func generateMinuteSnapshot() {
	panic("unimplemented")
}
