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
	// Manage trading accounts, balances, and portfolio history.
	Balances V1AccountBalanceService
	// Submit and monitor option exercise, DNE, CEA, and cancel instructions.
	Exercises V1AccountExerciseService
	// Place, monitor, and manage trading orders.
	Orders V1AccountOrderService
	// Manage trading accounts, balances, and portfolio history.
	PortfolioHistory V1AccountPortfolioHistoryService
	// View account positions.
	Positions V1AccountPositionService
}

// NewV1AccountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1AccountService(opts ...option.RequestOption) (r V1AccountService) {
	r = V1AccountService{}
	r.options = opts
	r.Balances = NewV1AccountBalanceService(opts...)
	r.Exercises = NewV1AccountExerciseService(opts...)
	r.Orders = NewV1AccountOrderService(opts...)
	r.PortfolioHistory = NewV1AccountPortfolioHistoryService(opts...)
	r.Positions = NewV1AccountPositionService(opts...)
	return
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

type V1AccountGetAccountsParams struct {
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	paramObj
}

// URLQuery serializes [V1AccountGetAccountsParams]'s query parameters as
// `url.Values`.
func (r V1AccountGetAccountsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
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
