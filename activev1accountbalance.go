// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// Manage trading accounts and view balances.
//
// ActiveV1AccountBalanceService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1AccountBalanceService] method instead.
type ActiveV1AccountBalanceService struct {
	Options []option.RequestOption
}

// NewActiveV1AccountBalanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1AccountBalanceService(opts ...option.RequestOption) (r ActiveV1AccountBalanceService) {
	r = ActiveV1AccountBalanceService{}
	r.Options = opts
	return
}

// Fetch account balance information
func (r *ActiveV1AccountBalanceService) GetAccountBalances(ctx context.Context, accountID int64, opts ...option.RequestOption) (res *ActiveV1AccountBalanceGetAccountBalancesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v/balances", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Represents the balance details for a trading account
type AccountBalances struct {
	// The unique identifier for the account
	AccountID int64 `json:"account_id" api:"required"`
	// The Reg T balance for the account
	Balance RegTBalance `json:"balance" api:"required"`
	// Realized profit or loss since start of day.
	DailyRealizedPnl string `json:"daily_realized_pnl" api:"required"`
	// Total profit or loss since start of day.
	DailyTotalPnl string `json:"daily_total_pnl" api:"required"`
	// Total unrealized profit or loss across all positions relative to prior close.
	DailyUnrealizedPnl string `json:"daily_unrealized_pnl" api:"required"`
	// The applicable margin model for the account
	//
	// Any of "OTHER", "NONE", "PORTFOLIO_MARGIN", "RISK_BASED_HAIRCUT_BROKER_DEALER",
	// "REG_T", "RISK_BASED_HAIRCUT_MARKET_MAKER", "CIRO", "FUTURES_NLV",
	// "FUTURES_TOT_EQ".
	MarginType MarginType `json:"margin_type" api:"required"`
	// Timestamp for the start-of-day values
	SodAsof APITimestamp `json:"sod_asof" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID          respjson.Field
		Balance            respjson.Field
		DailyRealizedPnl   respjson.Field
		DailyTotalPnl      respjson.Field
		DailyUnrealizedPnl respjson.Field
		MarginType         respjson.Field
		SodAsof            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBalances) RawJSON() string { return r.JSON.raw }
func (r *AccountBalances) UnmarshalJSON(data []byte) error {
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

// The Reg T balance for the account
type RegTBalance struct {
	// The total buying power available in the account
	BuyingPower string `json:"buying_power" api:"required"`
	// Currency identifier for all monetary values
	Currency string `json:"currency" api:"required"`
	// Day-trading buying power.
	DaytradingBuyingPower string `json:"daytrading_buying_power" api:"required"`
	// The total equity in the account (market value of all assets minus liabilities)
	Equity string `json:"equity" api:"required"`
	// The total market value of all long positions
	LongMarketValue string `json:"long_market_value" api:"required"`
	// Margin requirement for trade-date balances.
	MaintenanceMargin string `json:"maintenance_margin" api:"required"`
	// Margin excess for trade-date balances.
	MarginExcess string `json:"margin_excess" api:"required"`
	// Applied multiplier for margin calculations.
	Multiplier string `json:"multiplier" api:"required"`
	// Notional exposure from open risk-increasing orders.
	OpenOrderNotionalValue string `json:"open_order_notional_value" api:"required"`
	// Regulation T buying power.
	RegtBuyingPower string `json:"regt_buying_power" api:"required"`
	// The amount of cash that is settled and available for withdrawal or trading
	SettledCash string `json:"settled_cash" api:"required"`
	// The total market value of all short positions (represented as a positive value)
	ShortMarketValue string `json:"short_market_value" api:"required"`
	// Start-of-day cash balance.
	SodCash string `json:"sod_cash" api:"required"`
	// Start-of-day day-trading buying power.
	SodDaytradingBuyingPower string `json:"sod_daytrading_buying_power" api:"required"`
	// Start-of-day equity based on cash and positions.
	SodEquity string `json:"sod_equity" api:"required"`
	// Start-of-day long position market value (ex-cash).
	SodLongMarketValue string `json:"sod_long_market_value" api:"required"`
	// Start-of-day margin excess.
	SodMarginExcess string `json:"sod_margin_excess" api:"required"`
	// Start-of-day margin requirement.
	SodMarginRequirement string `json:"sod_margin_requirement" api:"required"`
	// Start-of-day Regulation T buying power.
	SodRegTBuyingPower string `json:"sod_reg_t_buying_power" api:"required"`
	// Start-of-day short position market value (ex-cash).
	SodShortMarketValue string `json:"sod_short_market_value" api:"required"`
	// Aggregated cash value.
	TradeCash string `json:"trade_cash" api:"required"`
	// Trade-date unsettled cash credits.
	UnsettledCashCredits string `json:"unsettled_cash_credits" api:"required"`
	// Trade-date unsettled cash debits.
	UnsettledCashDebits string `json:"unsettled_cash_debits" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BuyingPower              respjson.Field
		Currency                 respjson.Field
		DaytradingBuyingPower    respjson.Field
		Equity                   respjson.Field
		LongMarketValue          respjson.Field
		MaintenanceMargin        respjson.Field
		MarginExcess             respjson.Field
		Multiplier               respjson.Field
		OpenOrderNotionalValue   respjson.Field
		RegtBuyingPower          respjson.Field
		SettledCash              respjson.Field
		ShortMarketValue         respjson.Field
		SodCash                  respjson.Field
		SodDaytradingBuyingPower respjson.Field
		SodEquity                respjson.Field
		SodLongMarketValue       respjson.Field
		SodMarginExcess          respjson.Field
		SodMarginRequirement     respjson.Field
		SodRegTBuyingPower       respjson.Field
		SodShortMarketValue      respjson.Field
		TradeCash                respjson.Field
		UnsettledCashCredits     respjson.Field
		UnsettledCashDebits      respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegTBalance) RawJSON() string { return r.JSON.raw }
func (r *RegTBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountBalanceGetAccountBalancesResponse struct {
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
func (r ActiveV1AccountBalanceGetAccountBalancesResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountBalanceGetAccountBalancesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
