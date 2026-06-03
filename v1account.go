// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"encoding/json"
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
// V1AccountService contains methods and other services that help with interacting
// with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1AccountService] method instead.
type V1AccountService struct {
	options []option.RequestOption
}

// NewV1AccountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1AccountService(opts ...option.RequestOption) (r V1AccountService) {
	r = V1AccountService{}
	r.options = opts
	return
}

// Fetch account balance information
func (r *V1AccountService) GetAccountBalances(ctx context.Context, accountID int64, query V1AccountGetAccountBalancesParams, opts ...option.RequestOption) (res *V1AccountGetAccountBalancesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/balances", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Fetch account details by ID
func (r *V1AccountService) GetAccountByID(ctx context.Context, accountID int64, opts ...option.RequestOption) (res *V1AccountGetAccountByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List accounts the authenticated user has permission to access
func (r *V1AccountService) GetAccounts(ctx context.Context, query V1AccountGetAccountsParams, opts ...option.RequestOption) (res *V1AccountGetAccountsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves daily portfolio history for the specified account.
func (r *V1AccountService) GetPortfolioHistory(ctx context.Context, accountID int64, query V1AccountGetPortfolioHistoryParams, opts ...option.RequestOption) (res *V1AccountGetPortfolioHistoryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/portfolio-history", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Update account risk settings
func (r *V1AccountService) PatchAccountByID(ctx context.Context, accountID int64, body V1AccountPatchAccountByIDParams, opts ...option.RequestOption) (res *V1AccountPatchAccountByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Represents a trading account
type Account struct {
	// The unique identifier for the account
	ID int64 `json:"id" api:"required"`
	// The account holder entity identifier
	AccountHolderEntityID int64 `json:"account_holder_entity_id" api:"required"`
	// The full legal name of the account
	FullName string `json:"full_name" api:"required"`
	// The date the account was opened
	OpenDate time.Time `json:"open_date" api:"required" format:"date"`
	// The options level of the account
	OptionsLevel int64 `json:"options_level" api:"required"`
	// The short name of the account
	ShortName string `json:"short_name" api:"required"`
	// The current status of the account
	//
	// Any of "ACTIVE", "INACTIVE", "CLOSED".
	Status AccountStatus `json:"status" api:"required"`
	// The sub-type of account
	//
	// Any of "CASH", "MARGIN", "OTHER".
	Subtype AccountSubtype `json:"subtype" api:"required"`
	// The type of account
	//
	// Any of "CUSTOMER", "OTHER".
	Type AccountType `json:"type" api:"required"`
	// The date the account was closed, if applicable
	CloseDate time.Time `json:"close_date" api:"nullable" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccountHolderEntityID respjson.Field
		FullName              respjson.Field
		OpenDate              respjson.Field
		OptionsLevel          respjson.Field
		ShortName             respjson.Field
		Status                respjson.Field
		Subtype               respjson.Field
		Type                  respjson.Field
		CloseDate             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Account) RawJSON() string { return r.JSON.raw }
func (r *Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type AccountList []Account

type AccountSettings struct {
	// Risk settings for the account
	Risk RiskSettings `json:"risk" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Risk        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountSettings) RawJSON() string { return r.JSON.raw }
func (r *AccountSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account status
type AccountStatus string

const (
	AccountStatusActive   AccountStatus = "ACTIVE"
	AccountStatusInactive AccountStatus = "INACTIVE"
	AccountStatusClosed   AccountStatus = "CLOSED"
)

// Account subtype classification providing more granular categorization
type AccountSubtype string

const (
	AccountSubtypeCash   AccountSubtype = "CASH"
	AccountSubtypeMargin AccountSubtype = "MARGIN"
	AccountSubtypeOther  AccountSubtype = "OTHER"
)

// Account type classification
type AccountType string

const (
	AccountTypeCustomer AccountType = "CUSTOMER"
	AccountTypeOther    AccountType = "OTHER"
)

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

type PortfolioHistoryResponse struct {
	Segments []PortfolioHistorySegment `json:"segments" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Segments    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortfolioHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *PortfolioHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortfolioHistorySegment struct {
	// The date for this segment
	Date time.Time `json:"date" api:"required" format:"date"`
	// The equity at the end of the trading day.
	EodEquity string `json:"eod_equity" api:"required"`
	// Sum of the profit and loss realized from position closing trading activity.
	RealizedPnl string `json:"realized_pnl" api:"required"`
	// The equity at the start of the trading day.
	SodEquity string `json:"sod_equity" api:"required"`
	// Sum of the profit and loss from market changes.
	UnrealizedPnl string `json:"unrealized_pnl" api:"required"`
	// Amount bought MTM
	BoughtNotional string `json:"bought_notional" api:"nullable"`
	// Sum of the profit and loss from intraday trading activities for the trading day.
	DayPnl string `json:"day_pnl" api:"nullable"`
	// P&L after netting all realized and unrealized P&L, adjustments, dividends,
	// change in accruals, income and expenses
	NetPnl string `json:"net_pnl" api:"nullable"`
	// P&L attributable to start-of-day (carried) positions from market movement during
	// this trading day.
	PositionPnl string `json:"position_pnl" api:"nullable"`
	// Amount sold MTM
	SoldNotional string `json:"sold_notional" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date           respjson.Field
		EodEquity      respjson.Field
		RealizedPnl    respjson.Field
		SodEquity      respjson.Field
		UnrealizedPnl  respjson.Field
		BoughtNotional respjson.Field
		DayPnl         respjson.Field
		NetPnl         respjson.Field
		PositionPnl    respjson.Field
		SoldNotional   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortfolioHistorySegment) RawJSON() string { return r.JSON.raw }
func (r *PortfolioHistorySegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Risk settings for an account
type RiskSettings struct {
	// The maximum notional value available to the account
	MaxNotional string `json:"max_notional" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxNotional respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RiskSettings) RawJSON() string { return r.JSON.raw }
func (r *RiskSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RiskSettings to a RiskSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RiskSettingsParam.Overrides()
func (r RiskSettings) ToParam() RiskSettingsParam {
	return param.Override[RiskSettingsParam](json.RawMessage(r.RawJSON()))
}

// Risk settings for an account
type RiskSettingsParam struct {
	// The maximum notional value available to the account
	MaxNotional param.Opt[string] `json:"max_notional,omitzero"`
	paramObj
}

func (r RiskSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow RiskSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RiskSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountGetAccountBalancesResponse struct {
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
func (r V1AccountGetAccountBalancesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountGetAccountBalancesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountGetAccountByIDResponse struct {
	// Represents a trading account
	Data Account `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountGetAccountByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountGetAccountByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountGetAccountsResponse struct {
	Data AccountList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountGetAccountsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountGetAccountsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountGetPortfolioHistoryResponse struct {
	Data PortfolioHistoryResponse `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountGetPortfolioHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountGetPortfolioHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountPatchAccountByIDResponse struct {
	Data AccountSettings `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountPatchAccountByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountPatchAccountByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountGetAccountBalancesParams struct {
	// Limit the number of top margin contributors returned by the engine.
	TopMarginContributorsLimit param.Opt[int64] `query:"top_margin_contributors_limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1AccountGetAccountBalancesParams]'s query parameters as
// `url.Values`.
func (r V1AccountGetAccountBalancesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1AccountGetAccountsParams struct {
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	paramObj
}

// URLQuery serializes [V1AccountGetAccountsParams]'s query parameters as
// `url.Values`.
func (r V1AccountGetAccountsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1AccountGetPortfolioHistoryParams struct {
	// Start date for the portfolio history range, in YYYY-MM-DD format.
	StartDate time.Time `query:"start_date" api:"required" format:"date" json:"-"`
	// Defaults to today in America/New_York when omitted.
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [V1AccountGetPortfolioHistoryParams]'s query parameters as
// `url.Values`.
func (r V1AccountGetPortfolioHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1AccountPatchAccountByIDParams struct {
	// Risk settings for the account
	Risk RiskSettingsParam `json:"risk,omitzero"`
	paramObj
}

func (r V1AccountPatchAccountByIDParams) MarshalJSON() (data []byte, err error) {
	type shadow V1AccountPatchAccountByIDParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AccountPatchAccountByIDParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
