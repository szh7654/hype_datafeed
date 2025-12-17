package datahandler

import (
	"time"

	"github.com/apache/arrow/go/v18/arrow"
	"github.com/apache/arrow/go/v18/arrow/array"
	eth "github.com/ethereum/go-ethereum/common"
)

var (
	Positions map[uint64]Position = make(map[uint64]Position)

	positionBuilder *array.RecordBuilder = array.NewRecordBuilder(pool, arrow.NewSchema(
		[]arrow.Field{
			{Name: "time", Type: arrow.FixedWidthTypes.Timestamp_s},
			{Name: "asset_id", Type: arrow.PrimitiveTypes.Uint16},
			{Name: "user_id", Type: arrow.PrimitiveTypes.Uint32},
			{Name: "sz", Type: arrow.PrimitiveTypes.Float64},
			{Name: "entry_px", Type: arrow.PrimitiveTypes.Float64},
			{Name: "isolated", Type: arrow.FixedWidthTypes.Boolean},
			{Name: "one_way", Type: arrow.FixedWidthTypes.Boolean},
			{Name: "margin_used", Type: arrow.PrimitiveTypes.Float64},
			{Name: "leverage", Type: arrow.PrimitiveTypes.Uint8},
			{Name: "max_leverage", Type: arrow.PrimitiveTypes.Uint8},
			{Name: "unrealized_pnl", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	))
)

type Position struct {
	Szi           float64
	EntryPx       float64
	Isolated      bool
	OneWay        bool
	MarginUsed    float64
	Leverage      uint8
	MaxLeverage   uint8
	UnrealizedPnl float64
	LiquidationPx float64
}

func PositionId(assetId uint16, userId uint32) uint64 {
	return uint64(assetId)<<16 | uint64(userId)
}

func DecodePositionId(positionId uint64) (assetId uint16, userId uint32) {
	return uint16(positionId & 0xffff), uint32(positionId >> 16)
}

func getActivePositionUsers() map[uint32]eth.Address {
	res := map[uint32]eth.Address{}
	for positionId, _ := range Positions {
		_, userId := DecodePositionId(positionId)
		addr := UserAddrs[userId]
		res[userId] = addr
	}
	return res
}

func PositionRecord() arrow.Record {
	positionBuilder.Release()
	t := time.Now().Unix()
	for positionId, position := range Positions {
		assetId, userId := DecodePositionId(positionId)
		positionBuilder.Field(0).(*array.TimestampBuilder).Append(arrow.Timestamp(t))
		positionBuilder.Field(1).(*array.Uint16Builder).Append(assetId)
		positionBuilder.Field(2).(*array.Uint32Builder).Append(userId)
		positionBuilder.Field(3).(*array.Float64Builder).Append(position.Szi)
		positionBuilder.Field(4).(*array.Float64Builder).Append(position.EntryPx)
		positionBuilder.Field(5).(*array.BooleanBuilder).Append(position.Isolated)
		positionBuilder.Field(6).(*array.BooleanBuilder).Append(position.OneWay)
		positionBuilder.Field(7).(*array.Float64Builder).Append(position.MarginUsed)
		positionBuilder.Field(8).(*array.Uint8Builder).Append(position.Leverage)
		positionBuilder.Field(9).(*array.Uint8Builder).Append(position.MaxLeverage)
		positionBuilder.Field(10).(*array.Float64Builder).Append(position.UnrealizedPnl)
	}
	return positionBuilder.NewRecord()
}
