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

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/apiquery"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/param"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// ActiveV1AccountService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1AccountService] method instead.
type ActiveV1AccountService struct {
	Options          []option.RequestOption
	Balances         ActiveV1AccountBalanceService
	Locates          ActiveV1AccountLocateService
	Orders           ActiveV1AccountOrderService
	PortfolioHistory ActiveV1AccountPortfolioHistoryService
	Positions        ActiveV1AccountPositionService
}

// NewActiveV1AccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActiveV1AccountService(opts ...option.RequestOption) (r ActiveV1AccountService) {
	r = ActiveV1AccountService{}
	r.Options = opts
	r.Balances = NewActiveV1AccountBalanceService(opts...)
	r.Locates = NewActiveV1AccountLocateService(opts...)
	r.Orders = NewActiveV1AccountOrderService(opts...)
	r.PortfolioHistory = NewActiveV1AccountPortfolioHistoryService(opts...)
	r.Positions = NewActiveV1AccountPositionService(opts...)
	return
}

// Fetch account details by ID
func (r *ActiveV1AccountService) GetAccountByID(ctx context.Context, accountID int64, opts ...option.RequestOption) (res *ActiveV1AccountGetAccountByIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List accounts the authenticated user has permission to access
func (r *ActiveV1AccountService) GetAccounts(ctx context.Context, query ActiveV1AccountGetAccountsParams, opts ...option.RequestOption) (res *ActiveV1AccountGetAccountsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update account risk settings
func (r *ActiveV1AccountService) PatchAccountByID(ctx context.Context, accountID int64, body ActiveV1AccountPatchAccountByIDParams, opts ...option.RequestOption) (res *ActiveV1AccountPatchAccountByIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Represents a trading account
type Account struct {
	// The unique identifier for the account
	ID int64 `json:"id" api:"required"`
	// The full legal name of the account
	FullName string `json:"full_name" api:"required"`
	// The type of account
	//
	// Any of "HOUSE", "PAB", "CUSTOMER", "COUNTERPARTY", "OTHER".
	Kind AccountKind `json:"kind" api:"required"`
	// The date the account was opened
	OpenDate time.Time `json:"open_date" api:"required" format:"date"`
	// The short name of the account
	ShortName string `json:"short_name" api:"required"`
	// The current status of the account
	//
	// Any of "ACTIVE", "INACTIVE", "CLOSED".
	Status AccountStatus `json:"status" api:"required"`
	// The sub-type of account
	//
	// Any of "AFFILIATE", "ALLOCATION", "ARRANGING", "BANK", "BLOCK_TRADING",
	// "CARRY_BROKER", "CASH", "CLIENT", "COLLATERAL", "COURTESY_MASTER", "CROSS",
	// "DEPOSIT", "DVP", "ERROR", "EXECUTION", "FACILITATION", "FUNDING_SOURCE",
	// "HEDGE", "MARGIN", "MUTUAL_FUND", "OPERATING", "OTHER", "RELATED_MASTER",
	// "REPO", "SECURITIES_LENDING", "SHADOW_AWAY", "TRADING",
	// "TRIPARTY_COLLATERAL_AWAY", "UNKNOWN".
	Subkind AccountSubkind `json:"subkind" api:"required"`
	// The date the account was closed, if applicable
	CloseDate time.Time `json:"close_date" api:"nullable" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		FullName    respjson.Field
		Kind        respjson.Field
		OpenDate    respjson.Field
		ShortName   respjson.Field
		Status      respjson.Field
		Subkind     respjson.Field
		CloseDate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Account) RawJSON() string { return r.JSON.raw }
func (r *Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account kind classification
type AccountKind string

const (
	AccountKindHouse        AccountKind = "HOUSE"
	AccountKindPab          AccountKind = "PAB"
	AccountKindCustomer     AccountKind = "CUSTOMER"
	AccountKindCounterparty AccountKind = "COUNTERPARTY"
	AccountKindOther        AccountKind = "OTHER"
)

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

// Account sub-kind classification providing more granular categorization
type AccountSubkind string

const (
	AccountSubkindAffiliate              AccountSubkind = "AFFILIATE"
	AccountSubkindAllocation             AccountSubkind = "ALLOCATION"
	AccountSubkindArranging              AccountSubkind = "ARRANGING"
	AccountSubkindBank                   AccountSubkind = "BANK"
	AccountSubkindBlockTrading           AccountSubkind = "BLOCK_TRADING"
	AccountSubkindCarryBroker            AccountSubkind = "CARRY_BROKER"
	AccountSubkindCash                   AccountSubkind = "CASH"
	AccountSubkindClient                 AccountSubkind = "CLIENT"
	AccountSubkindCollateral             AccountSubkind = "COLLATERAL"
	AccountSubkindCourtesyMaster         AccountSubkind = "COURTESY_MASTER"
	AccountSubkindCross                  AccountSubkind = "CROSS"
	AccountSubkindDeposit                AccountSubkind = "DEPOSIT"
	AccountSubkindDvp                    AccountSubkind = "DVP"
	AccountSubkindError                  AccountSubkind = "ERROR"
	AccountSubkindExecution              AccountSubkind = "EXECUTION"
	AccountSubkindFacilitation           AccountSubkind = "FACILITATION"
	AccountSubkindFundingSource          AccountSubkind = "FUNDING_SOURCE"
	AccountSubkindHedge                  AccountSubkind = "HEDGE"
	AccountSubkindMargin                 AccountSubkind = "MARGIN"
	AccountSubkindMutualFund             AccountSubkind = "MUTUAL_FUND"
	AccountSubkindOperating              AccountSubkind = "OPERATING"
	AccountSubkindOther                  AccountSubkind = "OTHER"
	AccountSubkindRelatedMaster          AccountSubkind = "RELATED_MASTER"
	AccountSubkindRepo                   AccountSubkind = "REPO"
	AccountSubkindSecuritiesLending      AccountSubkind = "SECURITIES_LENDING"
	AccountSubkindShadowAway             AccountSubkind = "SHADOW_AWAY"
	AccountSubkindTrading                AccountSubkind = "TRADING"
	AccountSubkindTripartyCollateralAway AccountSubkind = "TRIPARTY_COLLATERAL_AWAY"
	AccountSubkindUnknown                AccountSubkind = "UNKNOWN"
)

// A trading order with its current state and execution details.
//
// This is the unified API representation of an order across its lifecycle,
// combining data from execution reports, order status queries, and parent/child
// tracking.
type Order struct {
	// Client-provided unique identifier for this order
	ID string `json:"id" api:"required"`
	// Account placing the order
	AccountID int64 `json:"account_id" api:"required"`
	// Timestamp when order was created (UTC)
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Cumulative filled quantity
	FilledQuantity string `json:"filled_quantity" api:"required"`
	// Remaining unfilled quantity
	LeavesQuantity string `json:"leaves_quantity" api:"required"`
	// Type of order (MARKET, LIMIT, etc.)
	//
	// Any of "MARKET", "LIMIT", "STOP", "STOP_LIMIT", "TRAILING_STOP",
	// "TRAILING_STOP_LIMIT", "OTHER".
	OrderType OrderType `json:"order_type" api:"required"`
	// Total order quantity
	Quantity string `json:"quantity" api:"required"`
	// The identifier for the traded instrument (CMS/CUSIP/ISIN/FIGI for equities or
	// option OPRA OSI)
	SecurityID string `json:"security_id" api:"required"`
	// The source of the security identifier
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "CURRENCY", "FMP", "OEMS",
	// "SEDOL", "QUIK", "ISIN", "RIC", "COUNTRY", "EXCHANGE", "CTA", "BLOOMBERG",
	// "WERTPAPIER", "DUTCH", "VALOREN", "SICOVAM", "BELGIAN", "COMMON",
	// "CLEARING_HOUSE", "ISDA_FPML_SPECIFICATION", "ISDA_FPML_URL",
	// "LETTER_OF_CREDIT", "MARKETPLACE_ASSIGNED_IDENTIFIER", "MARKIT_RED_ENTITY_CLIP",
	// "MARKIT_RED_PAIR_CLIP", "CFTC", "ISDA_COMMODITY_REFERENCE_PRICE",
	// "LEGAL_ENTITY_IDENTIFIER", "SYNTHETIC", "FIDESSA_INSTRUMENT_MNEMONIC",
	// "INDEX_NAME", "UNIFORM_SYMBOL", "DIGITAL_TOKEN_IDENTIFIER", "MASSIVE", "OTHER".
	SecurityIDSource SecurityIDSource `json:"security_id_source" api:"required"`
	// Type of security
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	SecurityType SecurityType `json:"security_type" api:"required"`
	// Side of the order (BUY, SELL, SELL_SHORT)
	//
	// Any of "BUY", "SELL", "SELL_SHORT", "OTHER".
	Side Side `json:"side" api:"required"`
	// Current status of the order
	//
	// Any of "PENDING_NEW", "NEW", "PARTIALLY_FILLED", "FILLED", "CANCELED",
	// "REJECTED", "EXPIRED", "PENDING_CANCEL", "PENDING_REPLACE", "REPLACED",
	// "DONE_FOR_DAY", "STOPPED", "SUSPENDED", "CALCULATED", "OTHER".
	Status OrderStatus `json:"status" api:"required"`
	// Trading symbol
	Symbol string `json:"symbol" api:"required"`
	// Time in force instruction
	//
	// Any of "DAY", "GOOD_TILL_CANCEL", "IMMEDIATE_OR_CANCEL", "FILL_OR_KILL",
	// "GOOD_TILL_DATE", "AT_THE_OPENING", "AT_THE_CLOSE", "GOOD_TILL_CROSSING",
	// "GOOD_THROUGH_CROSSING", "AT_CROSSING", "OTHER".
	TimeInForce TimeInForce `json:"time_in_force" api:"required"`
	// Timestamp of the most recent update (UTC)
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// MIC code of the venue where the order is routed
	Venue string `json:"venue" api:"required"`
	// Average fill price across all executions
	AverageFillPrice string `json:"average_fill_price" api:"nullable"`
	// Contains execution, rejection or cancellation details, if any
	Details []string `json:"details"`
	// Timestamp when the order will expire (UTC). Present when time_in_force is
	// GOOD_TILL_DATE.
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// Limit offset for trailing stop-limit orders (signed)
	LimitOffset string `json:"limit_offset" api:"nullable"`
	// Limit price (for LIMIT and STOP_LIMIT orders)
	LimitPrice string `json:"limit_price" api:"nullable"`
	// Stop price (for STOP and STOP_LIMIT orders)
	StopPrice string `json:"stop_price" api:"nullable"`
	// Execution strategy for this order
	Strategy OrderStrategyUnion `json:"strategy" api:"nullable"`
	// Trailing offset amount for trailing orders
	TrailingOffsetAmt string `json:"trailing_offset_amt" api:"nullable"`
	// Trailing offset type for trailing orders
	//
	// Any of "PRICE", "PERCENT_BPS".
	TrailingOffsetAmtType TrailingOffsetType `json:"trailing_offset_amt_type" api:"nullable"`
	// Trailing watermark price for trailing orders
	TrailingWatermarkPx string `json:"trailing_watermark_px" api:"nullable"`
	// Trailing watermark timestamp for trailing orders
	TrailingWatermarkTs time.Time `json:"trailing_watermark_ts" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccountID             respjson.Field
		CreatedAt             respjson.Field
		FilledQuantity        respjson.Field
		LeavesQuantity        respjson.Field
		OrderType             respjson.Field
		Quantity              respjson.Field
		SecurityID            respjson.Field
		SecurityIDSource      respjson.Field
		SecurityType          respjson.Field
		Side                  respjson.Field
		Status                respjson.Field
		Symbol                respjson.Field
		TimeInForce           respjson.Field
		UpdatedAt             respjson.Field
		Venue                 respjson.Field
		AverageFillPrice      respjson.Field
		Details               respjson.Field
		ExpiresAt             respjson.Field
		LimitOffset           respjson.Field
		LimitPrice            respjson.Field
		StopPrice             respjson.Field
		Strategy              respjson.Field
		TrailingOffsetAmt     respjson.Field
		TrailingOffsetAmtType respjson.Field
		TrailingWatermarkPx   respjson.Field
		TrailingWatermarkTs   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Order) RawJSON() string { return r.JSON.raw }
func (r *Order) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderList []Order

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

type ActiveV1AccountGetAccountByIDResponse struct {
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
func (r ActiveV1AccountGetAccountByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountGetAccountByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountGetAccountsResponse struct {
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
func (r ActiveV1AccountGetAccountsResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountGetAccountsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountPatchAccountByIDResponse struct {
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
func (r ActiveV1AccountPatchAccountByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountPatchAccountByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountGetAccountsParams struct {
	// The number of items to return per page (only used when page_token is not
	// provided)
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1AccountGetAccountsParams]'s query parameters as
// `url.Values`.
func (r ActiveV1AccountGetAccountsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActiveV1AccountPatchAccountByIDParams struct {
	// Risk settings for the account
	Risk RiskSettingsParam `json:"risk,omitzero"`
	paramObj
}

func (r ActiveV1AccountPatchAccountByIDParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1AccountPatchAccountByIDParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1AccountPatchAccountByIDParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
