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
	shimjson "github.com/stainless-sdks/clear-street-go/internal/encoding/json"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/param"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// Manage locate requests for short selling.
//
// ActiveV1AccountLocateService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1AccountLocateService] method instead.
type ActiveV1AccountLocateService struct {
	Options []option.RequestOption
	// Manage locate requests for short selling.
	Inventory ActiveV1AccountLocateInventoryService
}

// NewActiveV1AccountLocateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1AccountLocateService(opts ...option.RequestOption) (r ActiveV1AccountLocateService) {
	r = ActiveV1AccountLocateService{}
	r.Options = opts
	r.Inventory = NewActiveV1AccountLocateInventoryService(opts...)
	return
}

// Submits a new short stock locate request.
func (r *ActiveV1AccountLocateService) NewLocateRequest(ctx context.Context, accountID int64, body ActiveV1AccountLocateNewLocateRequestParams, opts ...option.RequestOption) (res *ActiveV1AccountLocateNewLocateRequestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v/locates", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves all locate requests for the specified account.
func (r *ActiveV1AccountLocateService) GetLocateRequests(ctx context.Context, accountID int64, query ActiveV1AccountLocateGetLocateRequestsParams, opts ...option.RequestOption) (res *ActiveV1AccountLocateGetLocateRequestsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v/locates", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Modifies an existing locate request.
func (r *ActiveV1AccountLocateService) UpdateLocateRequest(ctx context.Context, accountID int64, body ActiveV1AccountLocateUpdateLocateRequestParams, opts ...option.RequestOption) (res *ActiveV1AccountLocateUpdateLocateRequestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v/locates", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Represents a single locate order and its status
type LocateOrder struct {
	// The unique system-generated ID for the locate order
	LocateOrderID string `json:"locate_order_id" api:"required"`
	// The quantity of shares that have been located
	LocatedQuantity int64 `json:"located_quantity" api:"required"`
	// The client Market Participant Identifier, assigned by Clear Street
	Mpid string `json:"mpid" api:"required"`
	// The timestamp when the locate order was received from the client in UTC
	RequestedAt time.Time `json:"requested_at" api:"required" format:"date-time"`
	// The quantity of shares requested by the client
	RequestedQuantity int64 `json:"requested_quantity" api:"required"`
	// The status of the locate order
	//
	// Any of "PENDING", "OFFERED", "FILLED", "REJECTED", "DECLINED", "EXPIRED",
	// "CANCELED".
	Status LocateOrderStatus `json:"status" api:"required"`
	// The symbol of the security to locate
	Symbol string `json:"symbol" api:"required"`
	// The borrow rate for the security if held overnight, expressed as a decimal
	BorrowRate string `json:"borrow_rate" api:"nullable"`
	// Comments provided by the trading desk
	DeskComment string `json:"desk_comment" api:"nullable"`
	// The timestamp when the locate order will expire, set once the order has been
	// processed, in UTC
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// A unique ID for the locate order, available after the order has been `OFFERED`
	LocateID string `json:"locate_id" api:"nullable"`
	// The timestamp when the security was located in UTC
	LocatedAt time.Time `json:"located_at" api:"nullable" format:"date-time"`
	// The reference ID provided when submitting the locate order
	ReferenceID string `json:"reference_id" api:"nullable"`
	// The total cost of the locate
	TotalCost string `json:"total_cost" api:"nullable"`
	// Comments provided by the trader when submitting the locate order
	TraderComment string `json:"trader_comment" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LocateOrderID     respjson.Field
		LocatedQuantity   respjson.Field
		Mpid              respjson.Field
		RequestedAt       respjson.Field
		RequestedQuantity respjson.Field
		Status            respjson.Field
		Symbol            respjson.Field
		BorrowRate        respjson.Field
		DeskComment       respjson.Field
		ExpiresAt         respjson.Field
		LocateID          respjson.Field
		LocatedAt         respjson.Field
		ReferenceID       respjson.Field
		TotalCost         respjson.Field
		TraderComment     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocateOrder) RawJSON() string { return r.JSON.raw }
func (r *LocateOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocateOrderList []LocateOrder

// The status of a locate order
type LocateOrderStatus string

const (
	LocateOrderStatusPending  LocateOrderStatus = "PENDING"
	LocateOrderStatusOffered  LocateOrderStatus = "OFFERED"
	LocateOrderStatusFilled   LocateOrderStatus = "FILLED"
	LocateOrderStatusRejected LocateOrderStatus = "REJECTED"
	LocateOrderStatusDeclined LocateOrderStatus = "DECLINED"
	LocateOrderStatusExpired  LocateOrderStatus = "EXPIRED"
	LocateOrderStatusCanceled LocateOrderStatus = "CANCELED"
)

type ActiveV1AccountLocateNewLocateRequestResponse struct {
	Data LocateOrderList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1AccountLocateNewLocateRequestResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountLocateNewLocateRequestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountLocateGetLocateRequestsResponse struct {
	Data LocateOrderList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1AccountLocateGetLocateRequestsResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountLocateGetLocateRequestsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountLocateUpdateLocateRequestResponse struct {
	// Represents a single locate order and its status
	Data LocateOrder `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1AccountLocateUpdateLocateRequestResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountLocateUpdateLocateRequestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountLocateNewLocateRequestParams struct {
	Body []ActiveV1AccountLocateNewLocateRequestParamsBody
	paramObj
}

func (r ActiveV1AccountLocateNewLocateRequestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ActiveV1AccountLocateNewLocateRequestParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Request to create a new locate order
//
// The properties Quantity, Symbol are required.
type ActiveV1AccountLocateNewLocateRequestParamsBody struct {
	// The quantity of shares to locate
	Quantity int64 `json:"quantity" api:"required"`
	// The symbol of the security to locate
	Symbol string `json:"symbol" api:"required"`
	// Optional comments to associate with the locate request
	Comments param.Opt[string] `json:"comments,omitzero"`
	// A client-provided reference ID to identify the locate order
	ReferenceID param.Opt[string] `json:"reference_id,omitzero"`
	paramObj
}

func (r ActiveV1AccountLocateNewLocateRequestParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1AccountLocateNewLocateRequestParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1AccountLocateNewLocateRequestParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountLocateGetLocateRequestsParams struct {
	// The number of items to return per page (only used when page_token is not
	// provided)
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Filter by client reference ID
	ReferenceID param.Opt[string] `query:"reference_id,omitzero" json:"-"`
	// Filter by locate order status
	//
	// Any of "PENDING", "OFFERED", "FILLED", "REJECTED", "DECLINED", "EXPIRED",
	// "CANCELED".
	Status LocateOrderStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1AccountLocateGetLocateRequestsParams]'s query
// parameters as `url.Values`.
func (r ActiveV1AccountLocateGetLocateRequestsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActiveV1AccountLocateUpdateLocateRequestParams struct {
	// Whether to accept (`true`) or decline (`false`) the locate offer
	Accept bool `json:"accept" api:"required"`
	paramObj
}

func (r ActiveV1AccountLocateUpdateLocateRequestParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1AccountLocateUpdateLocateRequestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1AccountLocateUpdateLocateRequestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
