package dataretriver

import (
	"encoding/json"
	"indexer/util"

	eth "github.com/ethereum/go-ethereum/common"
	"github.com/mailru/easyjson/opt"
)

//go:generate easyjson -all

type Grouping string

const (
	GroupingNA           Grouping = "na"
	GroupingNormalTpsl   Grouping = "normalTpsl"
	GroupingPositionTpls Grouping = "positionTpsl"
)

// Constants for default values
const (
	DefaultSlippage = 0.05 // 5%
)

type Tpsl string // Advanced order type

const (
	TakeProfit Tpsl = "tp"
	StopLoss   Tpsl = "sl"
)

type AssetInfo struct {
	Name          string `json:"name"`
	SzDecimals    int    `json:"szDecimals"`
	MaxLeverage   int    `json:"maxLeverage"`
	MarginTableId int    `json:"marginTableId"`
	OnlyIsolated  bool   `json:"onlyIsolated"`
	IsDelisted    bool   `json:"isDelisted"`
}

type MarginTier struct {
	LowerBound  string `json:"lowerBound"`
	MaxLeverage int    `json:"maxLeverage"`
}

type MarginTable struct {
	ID          int
	Description string       `json:"description"`
	MarginTiers []MarginTier `json:"marginTiers"`
}

type Meta struct {
	Universe     []AssetInfo   `json:"universe"`
	MarginTables []MarginTable `json:"marginTables"`
}

type AssetCtx struct {
	Funding      string   `json:"funding"`
	OpenInterest string   `json:"openInterest"`
	PrevDayPx    string   `json:"prevDayPx"`
	DayNtlVlm    string   `json:"dayNtlVlm"`
	Premium      string   `json:"premium"`
	OraclePx     string   `json:"oraclePx"`
	MarkPx       string   `json:"markPx"`
	MidPx        string   `json:"midPx,omitempty"`
	ImpactPxs    []string `json:"impactPxs"`
	DayBaseVlm   string   `json:"dayBaseVlm,omitempty"`
}

// This type has no JSON annotation because it cannot be directly unmarshalled from the response
type MetaAndAssetCtxs struct {
	Meta
	Ctxs []AssetCtx
}

type SpotAssetInfo struct {
	Name        string `json:"name"`
	Tokens      []int  `json:"tokens"`
	Index       int    `json:"index"`
	IsCanonical bool   `json:"isCanonical"`
}

type EvmContract struct {
	Address             string `json:"address"`
	EvmExtraWeiDecimals int    `json:"evm_extra_wei_decimals"`
}

type SpotTokenInfo struct {
	Name        string       `json:"name"`
	SzDecimals  int          `json:"szDecimals"`
	WeiDecimals int          `json:"weiDecimals"`
	Index       int          `json:"index"`
	TokenID     string       `json:"tokenId"`
	IsCanonical bool         `json:"isCanonical"`
	EvmContract *EvmContract `json:"evmContract"`
	FullName    *string      `json:"fullName"`
}

type SpotMeta struct {
	Universe []SpotAssetInfo `json:"universe"`
	Tokens   []SpotTokenInfo `json:"tokens"`
}

type SpotAssetCtx struct {
	DayNtlVlm         string  `json:"dayNtlVlm"`
	MarkPx            string  `json:"markPx"`
	MidPx             *string `json:"midPx"`
	PrevDayPx         string  `json:"prevDayPx"`
	CirculatingSupply string  `json:"circulatingSupply"`
	Coin              string  `json:"coin"`
}

// This type has no JSON annotation because it cannot be directly unmarshalled from the response
type SpotMetaAndAssetCtxs struct {
	Meta SpotMeta
	Ctxs []SpotAssetCtx
}

// WsMsg represents a WebSocket message with a channel and data payload.
type WsMsg struct {
	Channel string         `json:"channel"`
	Data    map[string]any `json:"data"`
}

type LimitOrderType struct {
	Tif Tif `json:"tif"` // TifAlo, TifIoc, TifGtc
}

type TriggerOrderType struct {
	TriggerPx float64 `json:"triggerPx" msgpack:"triggerPx"`
	IsMarket  bool    `json:"isMarket"  msgpack:"isMarket"`
	Tpsl      Tpsl    `json:"tpsl"      msgpack:"tpsl"` // "tp" or "sl"
}

type BuilderInfo struct {
	Builder string `json:"b" msgpack:"b"`
	Fee     int    `json:"f" msgpack:"f"`
}

type CancelRequest struct {
	Coin string `json:"coin"`
	Oid  int64  `json:"oid"`
}

type CancelByCloidRequest struct {
	Coin  string `json:"coin"`
	Cloid string `json:"cloid"`
}

type Cloid struct {
	Value string
}

func (c Cloid) ToRaw() string {
	return c.Value
}

type PerpDexSchemaInput struct {
	FullName        string  `json:"fullName"`
	CollateralToken int     `json:"collateralToken"`
	OracleUpdater   *string `json:"oracleUpdater"`
}

type SpotBalance struct {
	Coin     string `json:"coin"`
	Token    int    `json:"token"`
	Hold     string `json:"hold"`
	Total    string `json:"total"`
	EntryNtl string `json:"entryNtl"`
}

type SpotUserState struct {
	Balances []SpotBalance `json:"balances"`
}

type OpenOrder struct {
	Coin      string  `json:"coin"`
	LimitPx   float64 `json:"limitPx,string"`
	Oid       int64   `json:"oid"`
	Side      string  `json:"side"`
	Size      float64 `json:"sz,string"`
	Timestamp int64   `json:"timestamp"`
}

type FrontendOpenOrder struct {
	Coin             string    `json:"coin"`
	IsPositionTpSl   bool      `json:"isPositionTpsl"`
	IsTrigger        bool      `json:"isTrigger"`
	LimitPx          float64   `json:"limitPx,string"`
	Oid              int64     `json:"oid"`
	OrderType        string    `json:"orderType"`
	OrigSz           float64   `json:"origSz,string"`
	ReduceOnly       bool      `json:"reduceOnly"`
	Side             OrderSide `json:"side"`
	Sz               float64   `json:"sz,string"`
	Timestamp        int64     `json:"timestamp"`
	TriggerCondition string    `json:"triggerCondition"`
	TriggerPx        float64   `json:"triggerPx,string"`
}

type OrderSide string

const (
	OrderSideAsk OrderSide = "A"
	OrderSideBid OrderSide = "B"
)

type QueriedOrder struct {
	Coin             string         `json:"coin"`
	Side             OrderSide      `json:"side"`
	LimitPx          string         `json:"limitPx"`
	Sz               string         `json:"sz"`
	Oid              int64          `json:"oid"`
	Timestamp        int64          `json:"timestamp"`
	TriggerCondition string         `json:"triggerCondition"`
	IsTrigger        bool           `json:"isTrigger"`
	TriggerPx        string         `json:"triggerPx"`
	Children         []QueriedOrder `json:"children"`
	IsPositionTpsl   bool           `json:"isPositionTpsl"`
	ReduceOnly       bool           `json:"reduceOnly"`
	OrderType        string         `json:"orderType"`
	OrigSz           string         `json:"origSz"`
	Tif              Tif            `json:"tif"`
	Cloid            *string        `json:"cloid"`
}

type OrderQueryResponse struct {
	Order           QueriedOrder     `json:"order"`
	Status          OrderStatusValue `json:"status"`
	StatusTimestamp int64            `json:"statusTimestamp"`
}

type OrderStatusValue string

const (
	// Placed successfully
	OrderStatusValueOpen OrderStatusValue = "open"
	// Filled
	OrderStatusValueFilled OrderStatusValue = "filled"
	// Canceled by user
	OrderStatusValueCanceled OrderStatusValue = "canceled"
	// Trigger order triggered
	OrderStatusValueTriggered OrderStatusValue = "triggered"
	// Rejected at time of placement
	OrderStatusValueRejected OrderStatusValue = "rejected"
	// Canceled because insufficient margin to fill
	OrderStatusValueMarginCanceled OrderStatusValue = "marginCanceled"
	// Vaults only. Canceled due to a user's withdrawal from vault
	OrderStatusValueVaultWithdrawalCanceled OrderStatusValue = "vaultWithdrawalCanceled"
	// Canceled due to order being too aggressive when open interest was at cap
	OrderStatusValueOpenInterestCapCanceled OrderStatusValue = "openInterestCapCanceled"
	// Canceled due to self-trade prevention
	OrderStatusValueSelfTradeCanceled OrderStatusValue = "selfTradeCanceled"
	// Canceled reduced-only order that does not reduce position
	OrderStatusValueReduceOnlyCanceled OrderStatusValue = "reduceOnlyCanceled"
	// TP/SL only. Canceled due to sibling ordering being filled
	OrderStatusValueSiblingFilledCanceled OrderStatusValue = "siblingFilledCanceled"
	// Canceled due to asset delisting
	OrderStatusValueDelistedCanceled OrderStatusValue = "delistedCanceled"
	// Canceled due to liquidation
	OrderStatusValueLiquidatedCanceled OrderStatusValue = "liquidatedCanceled"
	// API only. Canceled due to exceeding scheduled cancel deadline (dead man's switch)
	OrderStatusValueScheduledCancel OrderStatusValue = "scheduledCancel"
	// Rejected due to invalid tick price
	OrderStatusValueTickRejected OrderStatusValue = "tickRejected"
	// Rejected due to order notional below minimum
	OrderStatusValueMinTradeNtlRejected OrderStatusValue = "minTradeNtlRejected"
	// Rejected due to insufficient margin
	OrderStatusValuePerpMarginRejected OrderStatusValue = "perpMarginRejected"
	// Rejected due to reduce only
	OrderStatusValueReduceOnlyRejected OrderStatusValue = "reduceOnlyRejected"
	// Rejected due to post-only immediate match
	OrderStatusValueBadAloPxRejected OrderStatusValue = "badAloPxRejected"
	// Rejected due to IOC not able to match
	OrderStatusValueIocCancelRejected OrderStatusValue = "iocCancelRejected"
	// Rejected due to invalid TP/SL price
	OrderStatusValueBadTriggerPxRejected OrderStatusValue = "badTriggerPxRejected"
	// Rejected due to lack of liquidity for market order
	OrderStatusValueMarketOrderNoLiquidityRejected OrderStatusValue = "marketOrderNoLiquidityRejected"
	// Rejected due to open interest cap
	OrderStatusValuePositionIncreaseAtOpenInterestCapRejected OrderStatusValue = "positionIncreaseAtOpenInterestCapRejected"
	// Rejected due to open interest cap
	OrderStatusValuePositionFlipAtOpenInterestCapRejected OrderStatusValue = "positionFlipAtOpenInterestCapRejected"
	// Rejected due to price too aggressive at open interest cap
	OrderStatusValueTooAggressiveAtOpenInterestCapRejected OrderStatusValue = "tooAggressiveAtOpenInterestCapRejected"
	// Rejected due to open interest cap
	OrderStatusValueOpenInterestIncreaseRejected OrderStatusValue = "openInterestIncreaseRejected"
	// Rejected due to insufficient spot balance
	OrderStatusValueInsufficientSpotBalanceRejected OrderStatusValue = "insufficientSpotBalanceRejected"
	// Rejected due to price too far from oracle
	OrderStatusValueOracleRejected OrderStatusValue = "oracleRejected"
	// Rejected due to exceeding margin tier limit at current leverage
	OrderStatusValuePerpMaxPositionRejected OrderStatusValue = "perpMaxPositionRejected"
)

type OrderQueryStatus string

const (
	OrderQueryStatusSuccess OrderQueryStatus = "order"
	OrderQueryStatusError   OrderQueryStatus = "unknownOid"
)

type OrderQueryResult struct {
	Status OrderQueryStatus   `json:"status"`
	Order  OrderQueryResponse `json:"order,omitempty"`
}

type FundingHistory struct {
	Coin        string `json:"coin"`
	FundingRate string `json:"fundingRate"`
	Premium     string `json:"premium"`
	Time        int64  `json:"time"`
}

type UserFundingHistory struct {
	Delta Delta  `json:"delta"`
	Hash  string `json:"hash"`
	Time  int64  `json:"time"`
}

type Delta struct {
	Coin        string `json:"coin"`
	FundingRate string `json:"fundingRate"`
	Size        string `json:"size"`
	Type        string `json:"type"`
	USDC        string `json:"usdc"`
}

type UserFees struct {
	ActiveReferralDiscount string       `json:"activeReferralDiscount"`
	DailyUserVolume        []UserVolume `json:"dailyUserVlm"`
	FeeSchedule            FeeSchedule  `json:"feeSchedule"`
	UserAddRate            string       `json:"userAddRate"`
	UserCrossRate          string       `json:"userCrossRate"`
	UserSpotCrossRate      string       `json:"userSpotCrossRate"`
	UserSpotAddRate        string       `json:"userSpotAddRate"`
}

type UserActiveAssetData struct {
	User             string   `json:"user"`
	Coin             string   `json:"coin"`
	Leverage         Leverage `json:"leverage"`
	MaxTradeSzs      []string `json:"maxTradeSzs"`
	AvailableToTrade []string `json:"availableToTrade"`
	MarkPx           string   `json:"markPx"`
}

type UserVolume struct {
	Date      string `json:"date"`
	Exchange  string `json:"exchange"`
	UserAdd   string `json:"userAdd"`
	UserCross string `json:"userCross"`
}

type FeeSchedule struct {
	Add              string `json:"add"`
	Cross            string `json:"cross"`
	ReferralDiscount string `json:"referralDiscount"`
	Tiers            Tiers  `json:"tiers"`
}

type Tiers struct {
	MM  []MMTier  `json:"mm"`
	VIP []VIPTier `json:"vip"`
}

type MMTier struct {
	Add                 string `json:"add"`
	MakerFractionCutoff string `json:"makerFractionCutoff"`
}

type VIPTier struct {
	Add       string `json:"add"`
	Cross     string `json:"cross"`
	NtlCutoff string `json:"ntlCutoff"`
}

type StakingSummary struct {
	Delegated              string `json:"delegated"`
	Undelegated            string `json:"undelegated"`
	TotalPendingWithdrawal string `json:"totalPendingWithdrawal"`
	NPendingWithdrawals    int    `json:"nPendingWithdrawals"`
}

type StakingDelegation struct {
	Validator            string `json:"validator"`
	Amount               string `json:"amount"`
	LockedUntilTimestamp int64  `json:"lockedUntilTimestamp"`
}

type StakingReward struct {
	Time        int64  `json:"time"`
	Source      string `json:"source"`
	TotalAmount string `json:"totalAmount"`
}

type ReferralState struct {
	ReferralCode string   `json:"referralCode"`
	Referrer     string   `json:"referrer"`
	Referred     []string `json:"referred"`
}

type SubAccount struct {
	Name        string   `json:"name"`
	User        string   `json:"user"`
	Permissions []string `json:"permissions"`
}

type MultiSigSigner struct {
	User      string `json:"user"`
	Threshold int    `json:"threshold"`
}

type OrderStatus struct {
	Resting *OrderStatusResting `json:"resting,omitempty"`
	Filled  *OrderStatusFilled  `json:"filled,omitempty"`
	Error   *string             `json:"error,omitempty"`
}

type OrderStatusResting struct {
	Oid      int64   `json:"oid"`
	ClientID *string `json:"cloid"`
	Status   string  `json:"status"`
}

type OrderStatusFilled struct {
	TotalSz string `json:"totalSz"`
	AvgPx   string `json:"avgPx"`
	Oid     int    `json:"oid"`
}

func (s *OrderStatus) String() string {
	data, _ := json.Marshal(s)
	return string(data)
}

type OrderResponse struct {
	Statuses []OrderStatus
}

type BulkOrderResponse struct {
	Status string        `json:"status"`
	Data   []OrderStatus `json:"data,omitempty"`
	Error  string        `json:"error,omitempty"`
}

type CancelResponse struct {
	Status string     `json:"status"`
	Data   *OpenOrder `json:"data,omitempty"`
	Error  string     `json:"error,omitempty"`
}

type BulkCancelResponse struct {
	Status string      `json:"status"`
	Data   []OpenOrder `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

type ModifyResponse struct {
	Status string        `json:"status"`
	Data   []OrderStatus `json:"data,omitempty"`
	Error  string        `json:"error,omitempty"`
}

type TransferResponse struct {
	Status string `json:"status"`
	TxHash string `json:"txHash,omitempty"`
	Error  string `json:"error,omitempty"`
}

type ApprovalResponse struct {
	Status string `json:"status"`
	TxHash string `json:"txHash,omitempty"`
	Error  string `json:"error,omitempty"`
}

type CreateVaultResponse struct {
	Status string `json:"status"`
	Data   string `json:"data,omitempty"`
	Error  string `json:"error,omitempty"`
}

type CreateSubAccountResponse struct {
	Status string      `json:"status"`
	Data   *SubAccount `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

type SetReferrerResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type ScheduleCancelResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type AgentApprovalResponse struct {
	Status string `json:"status"`
	TxHash string `json:"txHash,omitempty"`
	Error  string `json:"error,omitempty"`
}

type MultiSigConversionResponse struct {
	Status string `json:"status"`
	TxHash string `json:"txHash,omitempty"`
	Error  string `json:"error,omitempty"`
}

type SpotDeployResponse struct {
	Status string `json:"status"`
	TxHash string `json:"txHash,omitempty"`
	Error  string `json:"error,omitempty"`
}

type ValidatorResponse struct {
	Status string `json:"status"`
	TxHash string `json:"txHash,omitempty"`
	Error  string `json:"error,omitempty"`
}

type MultiSigResponse struct {
	Status string `json:"status"`
	TxHash string `json:"txHash,omitempty"`
	Error  string `json:"error,omitempty"`
}

// PerpDex represents a perpetual DEX
type PerpDex struct {
	Name                     string     `json:"name"`
	FullName                 string     `json:"fullName"`
	Deployer                 string     `json:"deployer"`
	OracleUpdater            *string    `json:"oracleUpdater"`
	FeeRecipient             *string    `json:"feeRecipient"`
	AssetToStreamingOiCap    [][]string `json:"assetToStreamingOiCap"`    // Array of [coin, cap] tuples
	AssetToFundingMultiplier [][]string `json:"assetToFundingMultiplier"` // Array of [coin, multiplier] tuples
}

// PerpDexLimits represents limits for a builder-deployed perp DEX
type PerpDexLimits struct {
	TotalOiCap     string     `json:"totalOiCap"`
	OiSzCapPerPerp string     `json:"oiSzCapPerPerp"`
	MaxTransferNtl string     `json:"maxTransferNtl"`
	CoinToOiCap    [][]string `json:"coinToOiCap"` // Array of [coin, cap] tuples
}

// PerpDexStatus represents status for a builder-deployed perp DEX
type PerpDexStatus struct {
	TotalNetDeposit string `json:"totalNetDeposit"`
}

// PerpDeployAuctionStatus represents the status of a perp deploy auction
type PerpDeployAuctionStatus struct {
	StartTimeSeconds int64   `json:"startTimeSeconds"`
	DurationSeconds  int64   `json:"durationSeconds"`
	StartGas         string  `json:"startGas"`
	CurrentGas       string  `json:"currentGas"`
	EndGas           *string `json:"endGas"`
}

type BlockFill struct {
	LocalTime   util.NSTime `json:"local_time"`
	BlockTime   util.NSTime `json:"block_time"`
	BlockNumber uint32      `json:"block_number"`
	Events      []FillEvent `json:"events"`
}

type BlockOrderStatus struct {
	LocalTime   util.NSTime `json:"local_time"`
	BlockTime   util.NSTime `json:"block_time"`
	BlockNumber uint32      `json:"block_number"`
	Events      []any       `json:"events"`
}

type BlockOrderBookDiff struct {
	LocalTime   util.NSTime `json:"local_time"`
	BlockTime   util.NSTime `json:"block_time"`
	BlockNumber uint32      `json:"block_number"`
	Events      []any       `json:"events"`
}

type FillEvent struct {
	Address eth.Address `json:"address"`
	Fill    Fill        `json:"event"`
}

type Fill struct {
	Coin          string    `json:"coin"`
	Px            float64   `json:"px,string"`
	Sz            float64   `json:"sz,string"`
	Side          Side      `json:"side"`
	Time          int64     `json:"time"`
	StartPosition float64   `json:"startPosition"`
	Dir           string    `json:"dir"`
	ClosedPnl     float64   `json:"closedPnl"`
	Oid           uint64    `json:"oid"`
	Crossed       bool      `json:"crossed"`
	Fee           float64   `json:"fee"`
	TwapId        opt.Int64 `json:"twapId"`
}

type AssetPosition struct {
	Position     Position     `json:"position"`
	PositionType PositionType `json:"type"`
}

type Position struct {
	Coin           string     `json:"coin"`
	EntryPx        OptFloat64 `json:"entryPx"`
	Leverage       Leverage   `json:"leverage"`
	LiquidationPx  OptFloat64 `json:"liquidationPx"`
	MarginUsed     float64    `json:"marginUsed,string"`
	PositionValue  float64    `json:"positionValue,string"`
	ReturnOnEquity float64    `json:"returnOnEquity,string"`
	Szi            float64    `json:"szi,string"`
	UnrealizedPnl  float64    `json:"unrealizedPnl,string"`
	//CumFunding     *CumFunding `json:"cumFunding,omitempty"`
	MaxLeverage int `json:"maxLeverage"`
}

type Leverage struct {
	Type  LeverageType `json:"type"`
	Value int          `json:"value"`
}

type CumFunding struct {
	AllTime     string `json:"allTime"`
	SinceChange string `json:"sinceChange"`
	SinceOpen   string `json:"sinceOpen"`
}

type UserState struct {
	AssetPositions     []AssetPosition    `json:"assetPositions"`
	CrossMarginSummary CrossMarginSummary `json:"crossMarginSummary"`
	MarginSummary      MarginSummary      `json:"marginSummary"`
	// Withdrawable       string          `json:"withdrawable"`
}

type MarginSummary struct {
	AccountValue float64 `json:"accountValue,string"`
	//TotalMarginUsed float64 `json:"totalMarginUsed,string"`
	TotalNtlPos float64 `json:"totalNtlPos,string"`
	// TotalRawUsd     string  `json:"totalRawUsd"`
}

type CrossMarginSummary struct {
	AccountValue float64 `json:"accountValue,string"`
	// TotalMarginUsed string  `json:"totalMarginUsed"`
	// TotalNtlPos     string  `json:"totalNtlPos"`
	// TotalRawUsd     string  `json:"totalRawUsd"`
}
