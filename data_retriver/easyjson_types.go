package dataretriver

import (
	"strconv"
	"strings"

	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	"github.com/rs/zerolog/log"
)

type OptFloat64 float64

func (of *OptFloat64) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
		return
	}
	str := in.UnsafeString()

	if str != "" {
		// Deleted:*of = OptFloat64(float64(str))
		if f, err := strconv.ParseFloat(str, 64); err == nil {
			*of = OptFloat64(f)
		}
	}
}

type Dir int

// 定义枚举常量，使用 iota
const (
	DirOpenLong Dir = iota
	DirOpenShort
	DirCloseLong
	DirCloseShort
	DirBuy
	DirSell
	DirLongToShort
	DirShortToLong
	DirSpotDustConversion
	DirNetChildVaults
	DirLiquidatedIsolatedShort
	DirLiquidatedIsolatedLong
	DirAutoDeleveraging
	DirUnknown
)

func (d Dir) String() string {
	switch d {
	case DirOpenLong:
		return "Open Long"
	case DirOpenShort:
		return "Open Short"
	case DirCloseLong:
		return "Close Long"
	case DirCloseShort:
		return "Close Short"
	case DirBuy:
		return "Buy"
	case DirSell:
		return "Sell"
	case DirLongToShort:
		return "Long > Short"
	case DirShortToLong:
		return "Short > Long"
	case DirSpotDustConversion:
		return "Spot Dust Conversion"
	case DirNetChildVaults:
		return "Net Child Vaults"
	case DirLiquidatedIsolatedShort:
		return "Liquidated Isolated Short"
	case DirLiquidatedIsolatedLong:
		return "Liquidated Isolated Long"
	case DirAutoDeleveraging:
		return "Auto-Deleveraging"
	default:
		return "Unknown"
	}
}

func (d *Dir) FromString(dirStr string) {
	switch dirStr {
	case "Open Long":
		*d = DirOpenLong
	case "Open Short":
		*d = DirOpenShort
	case "Close Long":
		*d = DirCloseLong
	case "Close Short":
		*d = DirCloseShort
	case "Buy":
		*d = DirBuy
	case "Sell":
		*d = DirSell
	case "Long > Short":
		*d = DirLongToShort
	case "Short > Long":
		*d = DirShortToLong
	case "Spot Dust Conversion":
		*d = DirSpotDustConversion
	case "Net Child Vaults":
		*d = DirNetChildVaults
	case "Liquidated Isolated Short":
		*d = DirLiquidatedIsolatedShort
	case "Liquidated Isolated Long":
		*d = DirLiquidatedIsolatedLong
	case "Auto-Deleveraging":
		*d = DirAutoDeleveraging
	default:
		log.Warn().Msgf("Unsupported dir: %s", dirStr)
		*d = DirUnknown
	}
}

// UnmarshalEasyJSON 实现 easyjson 的反序列化接口
// 这是关键步骤，让 easyjson 知道如何将 JSON 字符串解析为此枚举类型
func (d *Dir) UnmarshalEasyJSON(in *jlexer.Lexer) {
	str := in.UnsafeString()
	d.FromString(str)
}

func (d Dir) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(d.String())
}

// ----------------------------
// Status 枚举
// ----------------------------

type Status int

const (
	StatusUnknown Status = iota
	StatusOpen
	StatusFilled
	StatusRejected
	StatusCanceled
	StatusReduceOnlyCanceled
	StatusMarginCanceled
	StatusSelfTradeCanceled
	StatusTriggered
	StatusSiblingFilledCanceled
	StatusScheduledCancel
)

func (s *Status) FromString(statusStr string) {
	// 特殊情况优先检查
	if strings.HasSuffix(statusStr, "Rejected") { // 注意：Rust 代码是 ends_with，Go 是 HasSuffix
		*s = StatusRejected
		return
	}

	// 使用 switch 进行精确匹配
	switch statusStr {
	case "open":
		*s = StatusOpen
	case "filled":
		*s = StatusFilled
	case "canceled":
		*s = StatusCanceled
	case "reduceOnlyCanceled":
		*s = StatusReduceOnlyCanceled
	case "marginCanceled":
		*s = StatusMarginCanceled
	case "selfTradeCanceled":
		*s = StatusSelfTradeCanceled
	case "triggered":
		*s = StatusTriggered
	case "siblingFilledCanceled":
		*s = StatusSiblingFilledCanceled
	case "scheduledCancel":
		*s = StatusScheduledCancel
	default:
		log.Warn().Msgf("Unknown order status: %s", statusStr)
		*s = StatusUnknown
	}
}

func (s Status) String() string {
	switch s {
	case StatusOpen:
		return "open"
	case StatusFilled:
		return "filled"
	case StatusCanceled:
		return "canceled"
	case StatusReduceOnlyCanceled:
		return "reduceOnlyCanceled"
	case StatusMarginCanceled:
		return "marginCanceled"
	case StatusSelfTradeCanceled:
		return "selfTradeCanceled"
	case StatusTriggered:
		return "triggered"
	case StatusSiblingFilledCanceled:
		return "siblingFilledCanceled"
	case StatusScheduledCancel:
		return "scheduledCancel"
	case StatusRejected: // Don't forget the Rejected case in String()
		return "rejected"
	default:
		return "unknown"
	}
}

func (s *Status) UnmarshalEasyJSON(in *jlexer.Lexer) {
	str := in.UnsafeString()
	s.FromString(str)
}

func (s Status) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(s.String())
}

// ----------------------------
// Side 枚举
// ----------------------------

type Side int

const (
	SideUnknown Side = iota
	SideA
	SideB
)

func (s *Side) FromString(sideStr string) {
	switch sideStr {
	case "A":
		*s = SideA
	case "B":
		*s = SideB
	default:
		log.Warn().Msgf("Unknown side: %s", sideStr)
		*s = SideUnknown
	}
}

func (s Side) String() string {
	switch s {
	case SideA:
		return "A"
	case SideB:
		return "B"
	default:
		return "Unknown"
	}
}

func (s *Side) UnmarshalEasyJSON(in *jlexer.Lexer) {
	str := in.UnsafeString()
	s.FromString(str)
}

func (s Side) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(s.String())
}

// ----------------------------
// TriggerConditionType 枚举
// ----------------------------

type TriggerConditionType int

const (
	TriggerConditionUnknown TriggerConditionType = iota
	TriggerConditionNone
	TriggerConditionPriceBelow
	TriggerConditionPriceAbove
	TriggerConditionTriggered
)

func (tct *TriggerConditionType) FromString(conditionStr string) {
	if conditionStr == "N/A" {
		*tct = TriggerConditionNone
		return
	}
	// starts_with 检查
	if strings.HasPrefix(conditionStr, "Price below ") {
		*tct = TriggerConditionPriceBelow
		return
	}
	if strings.HasPrefix(conditionStr, "Price above ") {
		*tct = TriggerConditionPriceAbove
		return
	}

	switch conditionStr {
	case "Triggered":
		*tct = TriggerConditionTriggered
	default:
		log.Warn().Msgf("Unknown trigger condition: %s", conditionStr)
		*tct = TriggerConditionUnknown
	}
}

func (tct TriggerConditionType) String() string {
	switch tct {
	case TriggerConditionNone:
		return "N/A"
	case TriggerConditionPriceBelow:
		return "Price below"
	case TriggerConditionPriceAbove:
		return "Price above"
	case TriggerConditionTriggered:
		return "Triggered"
	default:
		return "Unknown"
	}
}

func (tct *TriggerConditionType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	str := in.UnsafeString()
	tct.FromString(str)
}

func (tct TriggerConditionType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(tct.String())
}

// ----------------------------
// OrderType 枚举
// ----------------------------

type OrderType int

const (
	OrderTypeUnknown OrderType = iota
	OrderTypeLimit
	OrderTypeStopMarket
	OrderTypeMarket
	OrderTypeTakeProfitMarket
	OrderTypeStopLimit
	OrderTypeTakeProfitLimit
	OrderTypeSpotDustConversion
)

func (ot *OrderType) FromString(orderTypeStr string) {
	switch orderTypeStr {
	case "Limit":
		*ot = OrderTypeLimit
	case "Stop Market":
		*ot = OrderTypeStopMarket
	case "Market":
		*ot = OrderTypeMarket
	case "Take Profit Market":
		*ot = OrderTypeTakeProfitMarket
	case "Stop Limit":
		*ot = OrderTypeStopLimit
	case "Take Profit Limit":
		*ot = OrderTypeTakeProfitLimit
	case "Spot Dust Conversion":
		*ot = OrderTypeSpotDustConversion
	default:
		log.Warn().Msgf("Unknown order_type: %s", orderTypeStr)
		*ot = OrderTypeUnknown
	}
}

func (ot OrderType) String() string {
	switch ot {
	case OrderTypeLimit:
		return "Limit"
	case OrderTypeStopMarket:
		return "Stop Market"
	case OrderTypeMarket:
		return "Market"
	case OrderTypeTakeProfitMarket:
		return "Take Profit Market"
	case OrderTypeStopLimit:
		return "Stop Limit"
	case OrderTypeTakeProfitLimit:
		return "Take Profit Limit"
	case OrderTypeSpotDustConversion:
		return "Spot Dust Conversion"
	default:
		return "Unknown"
	}
}

func (ot *OrderType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	str := in.UnsafeString()
	ot.FromString(str)
}

func (ot OrderType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(ot.String())
}

type LeverageType int

const (
	LeverageTypeUnknown LeverageType = iota
	LeverageTypeCross
	LeverageTypeIsolated
)

func (lt *LeverageType) FromString(leverageTypeStr string) {
	switch leverageTypeStr {
	case "cross":
		*lt = LeverageTypeCross
	case "isolated":
		*lt = LeverageTypeIsolated
	default:
		log.Warn().Msgf("Unknown leverage_type: %s", leverageTypeStr)
		*lt = LeverageTypeUnknown
	}
}

func (lt LeverageType) String() string {
	switch lt {
	case LeverageTypeCross:
		return "cross"
	case LeverageTypeIsolated:
		return "isolated"
	default:
		return "Unknown"
	}
}

func (lt *LeverageType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	str := in.UnsafeString()
	lt.FromString(str)
}

func (lt LeverageType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(lt.String())
}

type PositionType int

const (
	PositionTypeOneWay PositionType = iota
	PositionTypeHedge
	PositionTypeUnknown
)

func (pt *PositionType) FromString(positionTypeStr string) {
	switch positionTypeStr {
	case "oneWay":
		*pt = PositionTypeOneWay
	case "hedge":
		*pt = PositionTypeHedge
	default:
		log.Warn().Msgf("Unknown position_type: %s", positionTypeStr)
		*pt = PositionTypeUnknown
	}
}

func (pt PositionType) String() string {
	switch pt {
	case PositionTypeOneWay:
		return "oneWay"
	case PositionTypeHedge:
		return "hedge"
	default:
		return "unknown"
	}
}

func (pt *PositionType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	str := in.UnsafeString()
	pt.FromString(str)
}

func (pt PositionType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(pt.String())
}

// ----------------------------
// Tif 枚举
// ----------------------------

type Tif int

const (
	TifUnknown Tif = iota
	TifNull
	TifAlo
	TifGtc
	TifIoc
	TifFrontendMarket
	TifLiquidationMarket
)

func (tif *Tif) FromString(tifStr string) {
	switch tifStr {
	case "":
		*tif = TifNull
	case "Alo":
		*tif = TifAlo
	case "Gtc":
		*tif = TifGtc
	case "Ioc":
		*tif = TifIoc
	case "FrontendMarket":
		*tif = TifFrontendMarket
	case "LiquidationMarket":
		*tif = TifLiquidationMarket
	default:
		log.Warn().Msgf("Unknown tif: %s", tifStr)
		*tif = TifUnknown
	}
}

// Handle non-string JSON values like Rust's custom Deserialize
// This function mimics the behavior of the Rust Deserialize impl and from_value
func (tif *Tif) FromJSONValue(in *jlexer.Lexer) {
	if in.IsStart() || in.IsNull() {
		*tif = TifNull
		in.Skip()
		return
	}
	str := in.UnsafeString()
	if !in.Ok() && in.IsNull() {
		*tif = TifNull
		return
	}
	tif.FromString(str)
}

func (tif Tif) String() string {
	switch tif {
	case TifNull:
		return "" // Empty string for Null
	case TifAlo:
		return "Alo"
	case TifGtc:
		return "Gtc"
	case TifIoc:
		return "Ioc"
	case TifFrontendMarket:
		return "FrontendMarket"
	case TifLiquidationMarket:
		return "LiquidationMarket"
	default:
		return "Unknown"
	}
}

func (tif *Tif) UnmarshalEasyJSON(in *jlexer.Lexer) {
	tif.FromJSONValue(in)
}

func (tif Tif) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(tif.String())
}
