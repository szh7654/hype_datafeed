package dataretriver

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	eth "github.com/ethereum/go-ethereum/common"
	jlexer "github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
	"github.com/rs/zerolog/log"
)

type StrTime time.Time

const nsTimeStrLen = len("2006-01-02T15:04:05.000")

func (ct *StrTime) UnmarshalJSON(data []byte) error {
	var str string
	err := sonic.Unmarshal(data, &str)
	if err != nil {
		log.Warn().Msgf("Error unmarshaling Dir: cannot decode string from JSON data '%s': %v", string(data), err)
		return nil
	}
	if len(str) > nsTimeStrLen {
		str = str[:nsTimeStrLen]
	}

	t, err := time.Parse("2006-01-02T15:04:05.000", str)
	if err != nil {
		log.Warn().Msgf("Error unmarshaling Dir: cannot parse time '%s': %v", str, err)
		return nil
	}
	*ct = StrTime(t)
	return nil
}

func (ct StrTime) MarshalJSON() ([]byte, error) {
	str := time.Time(ct).Format("2006-01-02T15:04:05.000")
	return sonic.Marshal(str)
}

type MsTime time.Time

func (ct *MsTime) UnmarshalJSON(data []byte) error {
	var str int64
	err := sonic.Unmarshal(data, &str)
	if err != nil {
		log.Warn().Msgf("Error unmarshaling Dir: cannot decode string from JSON data '%s': %v", string(data), err)
		return nil
	}
	t := time.UnixMilli(str)
	*ct = MsTime(t)
	return nil
}

func (ct MsTime) MarshalJSON() ([]byte, error) {
	ms := time.Time(ct).Format("2006-01-02T15:04:05.000")
	return sonic.Marshal(ms)
}

func (ct StrTime) Time() time.Time {
	return time.Time(ct)
}

func (v *MetaAndAssetCtxs) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	if err := sonic.Unmarshal(data, &arr); err != nil {
		return err
	}
	if len(arr) < 2 {
		return fmt.Errorf("expected at least 2 elements in response, got %d", len(arr))
	}
	err := sonic.Unmarshal(arr[0], &v.Meta)
	if err != nil {
		return fmt.Errorf("failed to unmarshal meta: %w, raw: %s", err, string(data))
	}
	err = sonic.Unmarshal(arr[1], &v.Ctxs)
	if err != nil {
		return fmt.Errorf("failed to unmarshal asset contexts: %w", err)
	}
	return nil
}

type FillEvent struct {
	Address eth.Address `json:"address"`
	Fill    Fill        `json:"event"`
}

func (v *FillEvent) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	if err := sonic.Unmarshal(data, &arr); err != nil {
		return err
	}

	if len(arr) != 2 {
		return fmt.Errorf("expected array of length 2, got %d", len(arr))
	}

	var addressStr string
	if err := sonic.Unmarshal(arr[0], &addressStr); err != nil {
		return err
	}
	v.Address = eth.HexToAddress(addressStr)

	if err := sonic.Unmarshal(arr[1], &v.Fill); err != nil {
		return err
	}

	return nil
}

type OptFloat64 float64

func (of *OptFloat64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	str := strings.Trim(string(data), "\" ")

	if str != "" {
		if f, err := strconv.ParseFloat(str, 64); err == nil {
			*of = OptFloat64(f)
		}
	}
	return nil
}

func (of OptFloat64) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(float64(of))
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
func (d *Dir) UnmarshalJSON(data []byte) error {
	var str string
	if err := sonic.Unmarshal(data, &str); err != nil {
		log.Warn().Msgf("Error unmarshaling Dir: cannot decode string from JSON data '%s': %v", string(data), err)
		*d = DirUnknown
		return nil
	}
	d.FromString(str)
	return nil
}

func (d Dir) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(d.String())
}

func (d Dir) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(d.String())
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

func (s *Status) UnmarshalJSON(data []byte) error {
	var str string
	if err := sonic.Unmarshal(data, &str); err != nil {
		log.Warn().Msgf("Error unmarshaling Status: cannot decode string from JSON data '%s': %v", string(data), err)
		*s = StatusUnknown
		return nil
	}
	s.FromString(str)
	return nil
}

func (s Status) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(s.String())
}

func (s Status) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(s.String())
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

func (s *Side) UnmarshalJSON(data []byte) error {
	var str string
	if err := sonic.Unmarshal(data, &str); err != nil {
		log.Warn().Msgf("Error unmarshaling Side: cannot decode string from JSON data '%s': %v", string(data), err)
		*s = SideUnknown
		return nil
	}
	s.FromString(str)
	return nil
}

func (s Side) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(s.String())
}

func (s Side) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(s.String())
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

func (tct *TriggerConditionType) UnmarshalJSON(data []byte) error {
	var str string
	if err := sonic.Unmarshal(data, &str); err != nil {
		log.Warn().Msgf("Error unmarshaling TriggerConditionType: cannot decode string from JSON data '%s': %v", string(data), err)
		*tct = TriggerConditionUnknown
		return nil
	}
	tct.FromString(str)
	return nil
}

func (tct TriggerConditionType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(tct.String())
}

func (tct TriggerConditionType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(tct.String())
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

func (ot OrderType) IsLimit() bool {
	if ot == OrderTypeLimit || ot == OrderTypeStopLimit || ot == OrderTypeTakeProfitLimit {
		return true
	}
	return false
}

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

func (ot *OrderType) UnmarshalJSON(data []byte) error {
	var str string
	if err := sonic.Unmarshal(data, &str); err != nil {
		log.Warn().Msgf("Error unmarshaling OrderType: cannot decode string from JSON data '%s': %v", string(data), err)
		*ot = OrderTypeUnknown
		return nil
	}
	ot.FromString(str)
	return nil
}

func (ot OrderType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(ot.String())
}

func (ot OrderType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(ot.String())
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

func (lt *LeverageType) UnmarshalJSON(data []byte) error {
	var str string
	if err := sonic.Unmarshal(data, &str); err != nil {
		log.Warn().Msgf("Error unmarshaling LeverageType: cannot decode string from JSON data '%s': %v", string(data), err)
		*lt = LeverageTypeUnknown
		return nil
	}
	lt.FromString(str)
	return nil
}

func (lt LeverageType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(lt.String())
}

func (lt LeverageType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(lt.String())
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

func (pt *PositionType) UnmarshalJSON(data []byte) error {
	var str string
	if err := sonic.Unmarshal(data, &str); err != nil {
		log.Warn().Msgf("Error unmarshaling PositionType: cannot decode string from JSON data '%s': %v", string(data), err)
		*pt = PositionTypeUnknown
		return nil
	}
	pt.FromString(str)
	return nil
}

func (pt PositionType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(pt.String())
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

func (tif *Tif) UnmarshalJSON(data []byte) error {
	var str string
	if err := sonic.Unmarshal(data, &str); err != nil {
		// Check if it's a null value
		if string(data) == "null" {
			*tif = TifNull
			return nil
		}
		log.Warn().Msgf("Error unmarshaling Tif: cannot decode string from JSON data '%s': %v", string(data), err)
		*tif = TifUnknown
		return nil
	}
	tif.FromString(str)
	return nil
}

func (tif Tif) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(tif.String())
}

func (tif Tif) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(tif.String())
}

type ActionType int

const (
	ActionTypeNew ActionType = iota
	ActionTypeUpdate
	ActionTypeRemove
	ActionTypeUnknown
)

func (v ActionType) MarshalJSON() ([]byte, error) {
	switch v {
	case ActionTypeNew:
		return sonic.Marshal("new")
	case ActionTypeUpdate:
		return sonic.Marshal("update")
	case ActionTypeRemove:
		return sonic.Marshal("remove")
	default:
		return sonic.Marshal("unknown")
	}
}

type OrderBookDiffAction struct {
	ActionType ActionType
	OrigSz     float64
	NewSz      float64
}

func (v *OrderBookDiffAction) UnmarshalJSON(data []byte) error {
	// 1. 使用 sonic AST 解析 JSON 数据
	node, err := sonic.Get(data)
	if err != nil {
		return err // 返回 sonic 解析错误
	}

	// 2. 根据 AST 节点类型进行分发处理
	switch node.Type() {
	case ast.V_STRING:
		str, err := node.String()
		if err != nil {
			log.Warn().Msgf("cannot decode string from JSON data '%s': %v", string(data), err)
			v.ActionType = ActionTypeUnknown
		}
		if str == "remove" {
			v.ActionType = ActionTypeRemove
		} else {
			log.Warn().Msgf("Unknown action type: %s", str)
			v.ActionType = ActionTypeUnknown
		}
	case ast.V_OBJECT:
		// 处理对象类型: "raw_book_diff": { ... }
		// 遍历对象的键（应该只有一个键：'update' 或 'new'）
		sanner := func(path ast.Sequence, node *ast.Node) bool {

			switch *path.Key {
			case "update":
				v.ActionType = ActionTypeUpdate
				v.OrigSz = parseFloatStr("origSz", node)
				v.NewSz = parseFloatStr("newSz", node)
			case "new":
				v.ActionType = ActionTypeNew
				v.NewSz = parseFloatStr("sz", node)
			}
			return true
		}
		err = node.ForEach(sanner)
		if err != nil {
			v.ActionType = ActionTypeUnknown
			return nil
		}

	default:
		v.ActionType = ActionTypeUnknown
		log.Warn().Msgf("Error unmarshaling OrderBookDiffAction: unsupported JSON type: %v", node.Type())
	}

	return nil
}

func parseFloatStr(key string, node *ast.Node) float64 {
	floatStr, err := node.Get(key).String()
	if err != nil {
		log.Warn().Msgf("Error unmarshaling OrderBookDiffAction: cannot decode %s from JSON data '%v': %v", key, node, err)
		return 0
	}
	floatVal, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		log.Warn().Msgf("Error unmarshaling OrderBookDiffAction: cannot decode %s from JSON data '%v': %v", key, node, err)
		return 0
	}
	return floatVal
}

type L4SnapShot struct {
	BlockNumber    uint32
	AssetSnapShots []AssetSnapShot
}

type AssetSnapShot struct {
	Name                           string
	BookOrdersAndUntriggeredOrders BookOrdersAndUntriggeredOrders
}

type BookOrdersAndUntriggeredOrders struct {
	BookOrders        BookOrders  `json:"book_orders"`
	UntriggeredOrders []AddrOrder `json:"untriggered_orders"`
}

type BookOrders struct {
	BidOrders []AddrOrder
	AskOrders []AddrOrder
}

type AddrOrder struct {
	User  eth.Address
	Order Order
}

func (v *L4SnapShot) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	if err := sonic.Unmarshal(data, &arr); err != nil {
		panic(err)
	}
	if len(arr) < 2 {
		panic("expected at least 2 elements in response, got %d")
	}
	if err := sonic.Unmarshal(arr[0], &v.BlockNumber); err != nil {
		panic(err)
	}

	if err := sonic.Unmarshal(arr[1], &v.AssetSnapShots); err != nil {
		panic(err)
	}

	return nil
}

func (v *AssetSnapShot) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	if err := sonic.Unmarshal(data, &arr); err != nil {
		panic(err)
	}
	if len(arr) < 2 {
		panic("")
	}
	if err := sonic.Unmarshal(arr[0], &v.Name); err != nil {
		panic(err)
	}
	if err := sonic.Unmarshal(arr[1], &v.BookOrdersAndUntriggeredOrders); err != nil {
		panic(err)
	}
	return nil
}

func (v *BookOrders) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	if err := sonic.Unmarshal(data, &arr); err != nil {
		panic(err)
	}
	if len(arr) < 2 {
		panic("")
	}
	if err := sonic.Unmarshal(arr[0], &v.BidOrders); err != nil {
		panic(err)
	}
	if err := sonic.Unmarshal(arr[1], &v.AskOrders); err != nil {
		panic(err)
	}
	return nil
}

func (v *AddrOrder) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	if err := sonic.Unmarshal(data, &arr); err != nil {
		panic(err)
	}
	if len(arr) < 2 {
		panic("")
	}
	if err := sonic.Unmarshal(arr[0], &v.User); err != nil {
		panic(err)
	}
	if err := sonic.Unmarshal(arr[1], &v.Order); err != nil {
		panic(err)
	}
	return nil
}
