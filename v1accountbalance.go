// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/internal/apiquery"
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/param"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Manage trading accounts, balances, and portfolio history.
//
// V1AccountBalanceService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1AccountBalanceService] method instead.
type V1AccountBalanceService struct {
	options []option.RequestOption
}

// NewV1AccountBalanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1AccountBalanceService(opts ...option.RequestOption) (r V1AccountBalanceService) {
	r = V1AccountBalanceService{}
	r.options = opts
	return
}

// Fetch account balance information
func (r *V1AccountBalanceService) GetAccountBalances(ctx context.Context, accountID int64, query V1AccountBalanceGetAccountBalancesParams, opts ...option.RequestOption) (res *V1AccountBalanceGetAccountBalancesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/balances", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Represents the balance details for a trading account
type AccountBalances struct {
	// The unique identifier for the account
	AccountID int64 `json:"account_id" api:"required"`
	// The total buying power available in the account.
	BuyingPower string `json:"buying_power" api:"required"`
	// Currency identifier for all monetary values.
	Currency string `json:"currency" api:"required"`
	// Realized profit or loss since start of day.
	DailyRealizedPnl string `json:"daily_realized_pnl" api:"required"`
	// Total profit or loss since start of day.
	DailyTotalPnl string `json:"daily_total_pnl" api:"required"`
	// Total unrealized profit or loss across all positions relative to prior close.
	DailyUnrealizedPnl string `json:"daily_unrealized_pnl" api:"required"`
	// The total equity in the account.
	Equity string `json:"equity" api:"required"`
	// The total market value of all long positions.
	LongMarketValue string `json:"long_market_value" api:"required"`
	// The applicable margin model for the account
	//
	// Any of "OTHER", "NONE", "PORTFOLIO_MARGIN", "RISK_BASED_HAIRCUT_BROKER_DEALER",
	// "REG_T", "RISK_BASED_HAIRCUT_MARKET_MAKER", "CIRO", "FUTURES_NLV",
	// "FUTURES_TOT_EQ".
	MarginType MarginType `json:"margin_type" api:"required"`
	// Signed buying-power correction from open orders.
	OpenOrderAdjustment string `json:"open_order_adjustment" api:"required"`
	// The amount of cash that is settled and available for withdrawal or trading.
	SettledCash string `json:"settled_cash" api:"required"`
	// Start-of-day snapshot balances.
	Sod AccountBalancesSod `json:"sod" api:"required"`
	// Trade-date effective cash.
	TradeCash string `json:"trade_cash" api:"required"`
	// Trade-date unsettled cash credits.
	UnsettledCredits string `json:"unsettled_credits" api:"required"`
	// Trade-date unsettled cash debits.
	UnsettledDebits string `json:"unsettled_debits" api:"required"`
	// The amount of cash currently available to withdraw.
	WithdrawableCash string `json:"withdrawable_cash" api:"required"`
	// Margin-account-only details.
	MarginDetails MarginDetails `json:"margin_details" api:"nullable"`
	// Applied multiplier for margin calculations.
	Multiplier string `json:"multiplier" api:"nullable"`
	// The total market value of all short positions.
	ShortMarketValue string `json:"short_market_value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID           respjson.Field
		BuyingPower         respjson.Field
		Currency            respjson.Field
		DailyRealizedPnl    respjson.Field
		DailyTotalPnl       respjson.Field
		DailyUnrealizedPnl  respjson.Field
		Equity              respjson.Field
		LongMarketValue     respjson.Field
		MarginType          respjson.Field
		OpenOrderAdjustment respjson.Field
		SettledCash         respjson.Field
		Sod                 respjson.Field
		TradeCash           respjson.Field
		UnsettledCredits    respjson.Field
		UnsettledDebits     respjson.Field
		WithdrawableCash    respjson.Field
		MarginDetails       respjson.Field
		Multiplier          respjson.Field
		ShortMarketValue    respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBalances) RawJSON() string { return r.JSON.raw }
func (r *AccountBalances) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBalancesSod struct {
	// Start-of-day buying power.
	BuyingPower string `json:"buying_power" api:"required"`
	// Start-of-day equity.
	Equity string `json:"equity" api:"required"`
	// Start-of-day long market value.
	LongMarketValue string `json:"long_market_value" api:"required"`
	// Start-of-day short market value.
	ShortMarketValue string `json:"short_market_value" api:"required"`
	// Timestamp for the start-of-day values.
	Asof time.Time `json:"asof" api:"nullable" format:"date"`
	// Start-of-day day-trade buying power.
	DayTradeBuyingPower string `json:"day_trade_buying_power" api:"nullable"`
	// Start-of-day maintenance margin excess.
	MaintenanceMarginExcess string `json:"maintenance_margin_excess" api:"nullable"`
	// Start-of-day maintenance margin requirement.
	MaintenanceMarginRequirement string `json:"maintenance_margin_requirement" api:"nullable"`
	// Start-of-day trade cash.
	TradeCash string `json:"trade_cash" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BuyingPower                  respjson.Field
		Equity                       respjson.Field
		LongMarketValue              respjson.Field
		ShortMarketValue             respjson.Field
		Asof                         respjson.Field
		DayTradeBuyingPower          respjson.Field
		MaintenanceMarginExcess      respjson.Field
		MaintenanceMarginRequirement respjson.Field
		TradeCash                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBalancesSod) RawJSON() string { return r.JSON.raw }
func (r *AccountBalancesSod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarginDetails struct {
	// The number of day trades executed over the 5 most recent trading days.
	DayTradeCount int64 `json:"day_trade_count" api:"required"`
	// Initial margin excess for trade-date balances.
	InitialMarginExcess string `json:"initial_margin_excess" api:"required"`
	// Initial margin requirement for trade-date balances.
	InitialMarginRequirement string `json:"initial_margin_requirement" api:"required"`
	// Maintenance margin excess for trade-date balances.
	MaintenanceMarginExcess string `json:"maintenance_margin_excess" api:"required"`
	// Maintenance margin requirement for trade-date balances.
	MaintenanceMarginRequirement string `json:"maintenance_margin_requirement" api:"required"`
	// `true` if the account is currently flagged as a PDT, otherwise `false`.
	PatternDayTrader bool `json:"pattern_day_trader" api:"required"`
	// The amount of day-trade buying power used during the current trading day.
	DayTradeBuyingPowerUsage string `json:"day_trade_buying_power_usage" api:"nullable"`
	// Optional top margin contributors, returned only when explicitly requested.
	TopContributors []MarginTopContributor `json:"top_contributors"`
	// Current usage totals.
	Usage MarginDetailsUsage `json:"usage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DayTradeCount                respjson.Field
		InitialMarginExcess          respjson.Field
		InitialMarginRequirement     respjson.Field
		MaintenanceMarginExcess      respjson.Field
		MaintenanceMarginRequirement respjson.Field
		PatternDayTrader             respjson.Field
		DayTradeBuyingPowerUsage     respjson.Field
		TopContributors              respjson.Field
		Usage                        respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarginDetails) RawJSON() string { return r.JSON.raw }
func (r *MarginDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarginDetailsUsage struct {
	// The total margin available in the current model.
	Total string `json:"total" api:"required"`
	// The amount of margin that is currently being utilized.
	Used string `json:"used" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Total       respjson.Field
		Used        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarginDetailsUsage) RawJSON() string { return r.JSON.raw }
func (r *MarginDetailsUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarginTopContributor struct {
	// Day-trade buying power consumed by fills against this underlying on the current
	// trade date. Populated only for pattern day trader accounts.
	DayTradeBuyingPowerUsage string `json:"day_trade_buying_power_usage" api:"required"`
	// Initial margin requirement attributable to this underlying.
	InitialMarginRequirement string `json:"initial_margin_requirement" api:"required"`
	// Maintenance margin requirement attributable to this underlying.
	MaintenanceMarginRequirement string `json:"maintenance_margin_requirement" api:"required"`
	// Net market value attributable to this underlying.
	MarketValue string `json:"market_value" api:"required"`
	// UUID of the underlying security contributing to margin requirement.
	UnderlyingInstrumentID string `json:"underlying_instrument_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DayTradeBuyingPowerUsage     respjson.Field
		InitialMarginRequirement     respjson.Field
		MaintenanceMarginRequirement respjson.Field
		MarketValue                  respjson.Field
		UnderlyingInstrumentID       respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarginTopContributor) RawJSON() string { return r.JSON.raw }
func (r *MarginTopContributor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An account's margin type
type MarginType string

const (
	MarginTypeOther                        MarginType = "OTHER"
	MarginTypeNone                         MarginType = "NONE"
	MarginTypePortfolioMargin              MarginType = "PORTFOLIO_MARGIN"
	MarginTypeRiskBasedHaircutBrokerDealer MarginType = "RISK_BASED_HAIRCUT_BROKER_DEALER"
	MarginTypeRegT                         MarginType = "REG_T"
	MarginTypeRiskBasedHaircutMarketMaker  MarginType = "RISK_BASED_HAIRCUT_MARKET_MAKER"
	MarginTypeCiro                         MarginType = "CIRO"
	MarginTypeFuturesNlv                   MarginType = "FUTURES_NLV"
	MarginTypeFuturesTotEq                 MarginType = "FUTURES_TOT_EQ"
)

type V1AccountBalanceGetAccountBalancesResponse struct {
	// Represents the balance details for a trading account
	Data AccountBalances `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountBalanceGetAccountBalancesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountBalanceGetAccountBalancesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountBalanceGetAccountBalancesParams struct {
	// Limit the number of top margin contributors returned by the engine.
	TopMarginContributorsLimit param.Opt[int64] `query:"top_margin_contributors_limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1AccountBalanceGetAccountBalancesParams]'s query
// parameters as `url.Values`.
func (r V1AccountBalanceGetAccountBalancesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
