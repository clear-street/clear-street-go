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
// ActiveV1AccountService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1AccountService] method instead.
type ActiveV1AccountService struct {
	options []option.RequestOption
	// Manage trading accounts, balances, and portfolio history.
	Balances ActiveV1AccountBalanceService
	// Place, monitor, and manage trading orders.
	Orders ActiveV1AccountOrderService
	// Manage trading accounts, balances, and portfolio history.
	PortfolioHistory ActiveV1AccountPortfolioHistoryService
	// View account positions.
	Positions ActiveV1AccountPositionService
}

// NewActiveV1AccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActiveV1AccountService(opts ...option.RequestOption) (r ActiveV1AccountService) {
	r = ActiveV1AccountService{}
	r.options = opts
	r.Balances = NewActiveV1AccountBalanceService(opts...)
	r.Orders = NewActiveV1AccountOrderService(opts...)
	r.PortfolioHistory = NewActiveV1AccountPortfolioHistoryService(opts...)
	r.Positions = NewActiveV1AccountPositionService(opts...)
	return
}

// Fetch account details by ID
func (r *ActiveV1AccountService) GetAccountByID(ctx context.Context, accountID int64, opts ...option.RequestOption) (res *ActiveV1AccountGetAccountByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List accounts the authenticated user has permission to access
func (r *ActiveV1AccountService) GetAccounts(ctx context.Context, query ActiveV1AccountGetAccountsParams, opts ...option.RequestOption) (res *ActiveV1AccountGetAccountsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Update account risk settings
func (r *ActiveV1AccountService) PatchAccountByID(ctx context.Context, accountID int64, body ActiveV1AccountPatchAccountByIDParams, opts ...option.RequestOption) (res *ActiveV1AccountPatchAccountByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v", accountID)
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
		ID                    respjson.Field
		AccountHolderEntityID respjson.Field
		FullName              respjson.Field
		Kind                  respjson.Field
		OpenDate              respjson.Field
		ShortName             respjson.Field
		Status                respjson.Field
		Subkind               respjson.Field
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
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
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
