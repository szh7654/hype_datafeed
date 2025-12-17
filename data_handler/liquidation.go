package datahandler

import "time"

var (
	Liquidations []Liquidation = make([]Liquidation, 0)
)

type Liquidation struct {
	UserId        uint32
	AssetId       uint16
	Szi           float64
	PositionValue float64
	ClosedPnl     float64
	time          time.Time
}
