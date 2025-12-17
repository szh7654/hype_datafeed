package datahandler

import (
	dr "indexer/data_retriver"
	"indexer/util"
	"time"

	"github.com/apache/arrow/go/v18/arrow/memory"
)

var pool = memory.NewGoAllocator()

func StartLoop() {
	ueserStateReqTicker := time.NewTicker(1 * time.Minute)
	sendUserAndPositionSnapshotTicker := time.NewTicker(1 * time.Minute)
	for {
		select {
		case <-ueserStateReqTicker.C:
			dr.UserStateReqChan <- getActivePositionUsers()

		case <-sendUserAndPositionSnapshotTicker.C:
			sendUserAndPositionSnapshot()
		case UserStateWithId := <-dr.UserStateChan:
			applyUserState(UserStateWithId.UserState, UserStateWithId.UserId)
		}
	}
}

func applyUserState(userState *dr.UserState, userId uint32) {
	user := Users[userId]
	user.AccountValue = util.ToF64(userState.CrossMarginSummary.AccountValue)
	user.TotalNtlPos = util.ToF64(userState.CrossMarginSummary.TotalNtlPos)

	user.CrossAccountValue = util.ToF64(userState.CrossMarginSummary.AccountValue)

	// TODO: 补全liquidationPx

	for _, position := range userState.AssetPositions {
		if assetId, ok := NameToAssetId[position.Position.Coin]; ok {
			positionId := PositionId(assetId, userId)
			var liquidationPx float64 = 0.0
			if position.Position.LiquidationPx != nil {
				liquidationPx = util.ToF64(*position.Position.LiquidationPx)
			}
			Positions[positionId] = Position{
				Szi:           util.ToF64(position.Position.Szi),
				EntryPx:       util.ToF64P(position.Position.EntryPx),
				Isolated:      position.Position.Leverage.Type == "isolated",
				OneWay:        position.Type == "oneWay",
				MarginUsed:    util.ToF64(position.Position.MarginUsed),
				Leverage:      uint8(position.Position.Leverage.Value),
				MaxLeverage:   uint8(position.Position.MaxLeverage),
				UnrealizedPnl: util.ToF64(position.Position.UnrealizedPnl),
				LiquidationPx: liquidationPx,
			}
		}
	}
}

func sendUserAndPositionSnapshot() {

}
