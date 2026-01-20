// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"encoding/json"
	"errors"
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

// ActiveV1AccountOrderService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1AccountOrderService] method instead.
type ActiveV1AccountOrderService struct {
	Options []option.RequestOption
}

// NewActiveV1AccountOrderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1AccountOrderService(opts ...option.RequestOption) (r ActiveV1AccountOrderService) {
	r = ActiveV1AccountOrderService{}
	r.Options = opts
	return
}

// All filter parameters can be used independently or combined. The only constraint
// is that `security_id` and `security_id_source` must be provided together if
// either is specified.
func (r *ActiveV1AccountOrderService) CancelAllOrders(ctx context.Context, accountID int64, body ActiveV1AccountOrderCancelAllOrdersParams, opts ...option.RequestOption) (res *ActiveV1AccountOrderCancelAllOrdersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Cancel a specific order
func (r *ActiveV1AccountOrderService) CancelOrder(ctx context.Context, orderID string, body ActiveV1AccountOrderCancelOrderParams, opts ...option.RequestOption) (res *ActiveV1AccountOrderCancelOrderResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/accounts/%v/orders/%s", body.AccountID, orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get order by ID
func (r *ActiveV1AccountOrderService) GetOrderByID(ctx context.Context, orderID string, query ActiveV1AccountOrderGetOrderByIDParams, opts ...option.RequestOption) (res *ActiveV1AccountOrderGetOrderByIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/accounts/%v/orders/%s", query.AccountID, orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List orders for an account with optional filtering
func (r *ActiveV1AccountOrderService) GetOrders(ctx context.Context, accountID int64, query ActiveV1AccountOrderGetOrdersParams, opts ...option.RequestOption) (res *ActiveV1AccountOrderGetOrdersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Replace an order with new parameters
func (r *ActiveV1AccountOrderService) ReplaceOrder(ctx context.Context, orderID string, params ActiveV1AccountOrderReplaceOrderParams, opts ...option.RequestOption) (res *ActiveV1AccountOrderReplaceOrderResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/accounts/%v/orders/%s", params.AccountID, orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Submit new orders
func (r *ActiveV1AccountOrderService) SubmitOrders(ctx context.Context, accountID int64, body ActiveV1AccountOrderSubmitOrdersParams, opts ...option.RequestOption) (res *ActiveV1AccountOrderSubmitOrdersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Arrival Price strategy
type ApStrategy struct {
	// Maximum percentage of market volume to participate in (0-100)
	MaxPercent int64 `json:"max_percent,nullable"`
	// Minimum percentage of market volume to participate in (0-100)
	MinPercent int64 `json:"min_percent,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxPercent  respjson.Field
		MinPercent  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseStrategyParamsResp
}

// Returns the unmodified JSON received from the API
func (r ApStrategy) RawJSON() string { return r.JSON.raw }
func (r *ApStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ApStrategy to a ApStrategyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ApStrategyParam.Overrides()
func (r ApStrategy) ToParam() ApStrategyParam {
	return param.Override[ApStrategyParam](json.RawMessage(r.RawJSON()))
}

// Arrival Price strategy
type ApStrategyParam struct {
	// Maximum percentage of market volume to participate in (0-100)
	MaxPercent param.Opt[int64] `json:"max_percent,omitzero"`
	// Minimum percentage of market volume to participate in (0-100)
	MinPercent param.Opt[int64] `json:"min_percent,omitzero"`
	BaseStrategyParams
}

func (r ApStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ApStrategyParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Base parameters common to most algorithmic strategies
type BaseStrategyParamsResp struct {
	// UTC timestamp to end execution (defaults to market close)
	EndAt time.Time `json:"end_at,nullable" format:"date-time"`
	// UTC timestamp to start execution (defaults to order placement time)
	StartAt time.Time `json:"start_at,nullable" format:"date-time"`
	// Urgency level for execution aggressiveness
	//
	// Any of "SUPER_PASSIVE", "PASSIVE", "MODERATE", "AGGRESSIVE", "SUPER_AGGRESSIVE".
	Urgency Urgency `json:"urgency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndAt       respjson.Field
		StartAt     respjson.Field
		Urgency     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseStrategyParamsResp) RawJSON() string { return r.JSON.raw }
func (r *BaseStrategyParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BaseStrategyParamsResp to a BaseStrategyParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BaseStrategyParams.Overrides()
func (r BaseStrategyParamsResp) ToParam() BaseStrategyParams {
	return param.Override[BaseStrategyParams](json.RawMessage(r.RawJSON()))
}

// Base parameters common to most algorithmic strategies
type BaseStrategyParams struct {
	// UTC timestamp to end execution (defaults to market close)
	EndAt param.Opt[time.Time] `json:"end_at,omitzero" format:"date-time"`
	// UTC timestamp to start execution (defaults to order placement time)
	StartAt param.Opt[time.Time] `json:"start_at,omitzero" format:"date-time"`
	// Urgency level for execution aggressiveness
	//
	// Any of "SUPER_PASSIVE", "PASSIVE", "MODERATE", "AGGRESSIVE", "SUPER_AGGRESSIVE".
	Urgency Urgency `json:"urgency,omitzero"`
	paramObj
}

func (r BaseStrategyParams) MarshalJSON() (data []byte, err error) {
	type shadow BaseStrategyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseStrategyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dark Pool strategy
type DarkStrategy struct {
	// Maximum percentage of market volume to participate in (0-100)
	MaxPercent int64 `json:"max_percent,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxPercent  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseStrategyParamsResp
}

// Returns the unmodified JSON received from the API
func (r DarkStrategy) RawJSON() string { return r.JSON.raw }
func (r *DarkStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DarkStrategy to a DarkStrategyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DarkStrategyParam.Overrides()
func (r DarkStrategy) ToParam() DarkStrategyParam {
	return param.Override[DarkStrategyParam](json.RawMessage(r.RawJSON()))
}

// Dark Pool strategy
type DarkStrategyParam struct {
	// Maximum percentage of market volume to participate in (0-100)
	MaxPercent param.Opt[int64] `json:"max_percent,omitzero"`
	BaseStrategyParams
}

func (r DarkStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*DarkStrategyParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Destination exchange for DMA orders (Market Identifier Code)
type Destination string

const (
	DestinationArcx Destination = "ARCX"
	DestinationBats Destination = "BATS"
	DestinationBaty Destination = "BATY"
	DestinationEdga Destination = "EDGA"
	DestinationEdgx Destination = "EDGX"
	DestinationEprl Destination = "EPRL"
	DestinationIexg Destination = "IEXG"
	DestinationMemx Destination = "MEMX"
	DestinationXase Destination = "XASE"
	DestinationXbos Destination = "XBOS"
	DestinationXcis Destination = "XCIS"
	DestinationXnms Destination = "XNMS"
	DestinationXnys Destination = "XNYS"
)

// Direct Market Access strategy
type DmaStrategy struct {
	// Destination exchange (MIC code)
	//
	// Any of "ARCX", "BATS", "BATY", "EDGA", "EDGX", "EPRL", "IEXG", "MEMX", "XASE",
	// "XBOS", "XCIS", "XNMS", "XNYS".
	Destination Destination `json:"destination,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Destination respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DmaStrategy) RawJSON() string { return r.JSON.raw }
func (r *DmaStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DmaStrategy to a DmaStrategyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DmaStrategyParam.Overrides()
func (r DmaStrategy) ToParam() DmaStrategyParam {
	return param.Override[DmaStrategyParam](json.RawMessage(r.RawJSON()))
}

// Direct Market Access strategy
//
// The property Destination is required.
type DmaStrategyParam struct {
	// Destination exchange (MIC code)
	//
	// Any of "ARCX", "BATS", "BATY", "EDGA", "EDGX", "EPRL", "IEXG", "MEMX", "XASE",
	// "XBOS", "XCIS", "XNMS", "XNYS".
	Destination Destination `json:"destination,omitzero,required"`
	paramObj
}

func (r DmaStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow DmaStrategyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DmaStrategyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Order status
type OrderStatus string

const (
	OrderStatusPendingNew      OrderStatus = "PENDING_NEW"
	OrderStatusNew             OrderStatus = "NEW"
	OrderStatusPartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	OrderStatusFilled          OrderStatus = "FILLED"
	OrderStatusCanceled        OrderStatus = "CANCELED"
	OrderStatusRejected        OrderStatus = "REJECTED"
	OrderStatusExpired         OrderStatus = "EXPIRED"
	OrderStatusPendingCancel   OrderStatus = "PENDING_CANCEL"
	OrderStatusPendingReplace  OrderStatus = "PENDING_REPLACE"
	OrderStatusReplaced        OrderStatus = "REPLACED"
	OrderStatusDoneForDay      OrderStatus = "DONE_FOR_DAY"
	OrderStatusStopped         OrderStatus = "STOPPED"
	OrderStatusSuspended       OrderStatus = "SUSPENDED"
	OrderStatusCalculated      OrderStatus = "CALCULATED"
	OrderStatusOther           OrderStatus = "OTHER"
)

// OrderStrategyUnion contains all possible properties and values from
// [OrderStrategySor], [OrderStrategyVwap], [OrderStrategyTwap], [OrderStrategyAp],
// [OrderStrategyPov], [OrderStrategyDark], [OrderStrategyDma].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OrderStrategyUnion struct {
	EndAt   time.Time `json:"end_at"`
	StartAt time.Time `json:"start_at"`
	// This field is from variant [OrderStrategySor].
	Urgency    Urgency `json:"urgency"`
	Type       string  `json:"type"`
	MaxPercent int64   `json:"max_percent"`
	MinPercent int64   `json:"min_percent"`
	// This field is from variant [OrderStrategyPov].
	TargetPercent int64 `json:"target_percent"`
	// This field is from variant [OrderStrategyDma].
	Destination Destination `json:"destination"`
	JSON        struct {
		EndAt         respjson.Field
		StartAt       respjson.Field
		Urgency       respjson.Field
		Type          respjson.Field
		MaxPercent    respjson.Field
		MinPercent    respjson.Field
		TargetPercent respjson.Field
		Destination   respjson.Field
		raw           string
	} `json:"-"`
}

func (u OrderStrategyUnion) AsSor() (v OrderStrategySor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OrderStrategyUnion) AsVwap() (v OrderStrategyVwap) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OrderStrategyUnion) AsTwap() (v OrderStrategyTwap) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OrderStrategyUnion) AsAp() (v OrderStrategyAp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OrderStrategyUnion) AsPov() (v OrderStrategyPov) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OrderStrategyUnion) AsDark() (v OrderStrategyDark) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OrderStrategyUnion) AsDma() (v OrderStrategyDma) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OrderStrategyUnion) RawJSON() string { return u.JSON.raw }

func (r *OrderStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OrderStrategyUnion to a OrderStrategyUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OrderStrategyUnionParam.Overrides()
func (r OrderStrategyUnion) ToParam() OrderStrategyUnionParam {
	return param.Override[OrderStrategyUnionParam](json.RawMessage(r.RawJSON()))
}

// Smart Order Router (default) - routes to best available venue
type OrderStrategySor struct {
	// Any of "SOR".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	SorStrategy
}

// Returns the unmodified JSON received from the API
func (r OrderStrategySor) RawJSON() string { return r.JSON.raw }
func (r *OrderStrategySor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volume Weighted Average Price - matches VWAP over a period
type OrderStrategyVwap struct {
	// Any of "VWAP".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	VwapStrategy
}

// Returns the unmodified JSON received from the API
func (r OrderStrategyVwap) RawJSON() string { return r.JSON.raw }
func (r *OrderStrategyVwap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time Weighted Average Price - spreads execution evenly over time
type OrderStrategyTwap struct {
	// Any of "TWAP".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	TwapStrategy
}

// Returns the unmodified JSON received from the API
func (r OrderStrategyTwap) RawJSON() string { return r.JSON.raw }
func (r *OrderStrategyTwap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Arrival Price - aims to match price at order placement time
type OrderStrategyAp struct {
	// Any of "AP".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ApStrategy
}

// Returns the unmodified JSON received from the API
func (r OrderStrategyAp) RawJSON() string { return r.JSON.raw }
func (r *OrderStrategyAp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Percentage of Volume - participates as a percentage of market volume
type OrderStrategyPov struct {
	// Any of "POV".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	PovStrategy
}

// Returns the unmodified JSON received from the API
func (r OrderStrategyPov) RawJSON() string { return r.JSON.raw }
func (r *OrderStrategyPov) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dark Pool - routes to dark pool venues
type OrderStrategyDark struct {
	// Any of "DARK".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	DarkStrategy
}

// Returns the unmodified JSON received from the API
func (r OrderStrategyDark) RawJSON() string { return r.JSON.raw }
func (r *OrderStrategyDark) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Direct Market Access - sends directly to a specified exchange
type OrderStrategyDma struct {
	// Any of "DMA".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	DmaStrategy
}

// Returns the unmodified JSON received from the API
func (r OrderStrategyDma) RawJSON() string { return r.JSON.raw }
func (r *OrderStrategyDma) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func OrderStrategyParamOfSor(type_ string) OrderStrategyUnionParam {
	var variant OrderStrategySorParam
	variant.Type = type_
	return OrderStrategyUnionParam{OfSor: &variant}
}

func OrderStrategyParamOfVwap(type_ string) OrderStrategyUnionParam {
	var variant OrderStrategyVwapParam
	variant.Type = type_
	return OrderStrategyUnionParam{OfVwap: &variant}
}

func OrderStrategyParamOfTwap(type_ string) OrderStrategyUnionParam {
	var variant OrderStrategyTwapParam
	variant.Type = type_
	return OrderStrategyUnionParam{OfTwap: &variant}
}

func OrderStrategyParamOfAp(type_ string) OrderStrategyUnionParam {
	var variant OrderStrategyApParam
	variant.Type = type_
	return OrderStrategyUnionParam{OfAp: &variant}
}

func OrderStrategyParamOfPov(targetPercent int64, type_ string) OrderStrategyUnionParam {
	var variant OrderStrategyPovParam
	variant.TargetPercent = targetPercent
	variant.Type = type_
	return OrderStrategyUnionParam{OfPov: &variant}
}

func OrderStrategyParamOfDark(type_ string) OrderStrategyUnionParam {
	var variant OrderStrategyDarkParam
	variant.Type = type_
	return OrderStrategyUnionParam{OfDark: &variant}
}

func OrderStrategyParamOfDma(destination Destination, type_ string) OrderStrategyUnionParam {
	var variant OrderStrategyDmaParam
	variant.Destination = destination
	variant.Type = type_
	return OrderStrategyUnionParam{OfDma: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OrderStrategyUnionParam struct {
	OfSor  *OrderStrategySorParam  `json:",omitzero,inline"`
	OfVwap *OrderStrategyVwapParam `json:",omitzero,inline"`
	OfTwap *OrderStrategyTwapParam `json:",omitzero,inline"`
	OfAp   *OrderStrategyApParam   `json:",omitzero,inline"`
	OfPov  *OrderStrategyPovParam  `json:",omitzero,inline"`
	OfDark *OrderStrategyDarkParam `json:",omitzero,inline"`
	OfDma  *OrderStrategyDmaParam  `json:",omitzero,inline"`
	paramUnion
}

func (u OrderStrategyUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSor,
		u.OfVwap,
		u.OfTwap,
		u.OfAp,
		u.OfPov,
		u.OfDark,
		u.OfDma)
}
func (u *OrderStrategyUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OrderStrategyUnionParam) asAny() any {
	if !param.IsOmitted(u.OfSor) {
		return u.OfSor
	} else if !param.IsOmitted(u.OfVwap) {
		return u.OfVwap
	} else if !param.IsOmitted(u.OfTwap) {
		return u.OfTwap
	} else if !param.IsOmitted(u.OfAp) {
		return u.OfAp
	} else if !param.IsOmitted(u.OfPov) {
		return u.OfPov
	} else if !param.IsOmitted(u.OfDark) {
		return u.OfDark
	} else if !param.IsOmitted(u.OfDma) {
		return u.OfDma
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OrderStrategyUnionParam) GetTargetPercent() *int64 {
	if vt := u.OfPov; vt != nil {
		return &vt.TargetPercent
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OrderStrategyUnionParam) GetDestination() *string {
	if vt := u.OfDma; vt != nil {
		return (*string)(&vt.Destination)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OrderStrategyUnionParam) GetUrgency() *string {
	if vt := u.OfSor; vt != nil {
		return (*string)(&vt.Urgency)
	} else if vt := u.OfVwap; vt != nil {
		return (*string)(&vt.Urgency)
	} else if vt := u.OfTwap; vt != nil {
		return (*string)(&vt.Urgency)
	} else if vt := u.OfAp; vt != nil {
		return (*string)(&vt.Urgency)
	} else if vt := u.OfPov; vt != nil {
		return (*string)(&vt.Urgency)
	} else if vt := u.OfDark; vt != nil {
		return (*string)(&vt.Urgency)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OrderStrategyUnionParam) GetType() *string {
	if vt := u.OfSor; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVwap; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTwap; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAp; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPov; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDark; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDma; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OrderStrategyUnionParam) GetMaxPercent() *int64 {
	if vt := u.OfVwap; vt != nil && vt.MaxPercent.Valid() {
		return &vt.MaxPercent.Value
	} else if vt := u.OfTwap; vt != nil && vt.MaxPercent.Valid() {
		return &vt.MaxPercent.Value
	} else if vt := u.OfAp; vt != nil && vt.MaxPercent.Valid() {
		return &vt.MaxPercent.Value
	} else if vt := u.OfDark; vt != nil && vt.MaxPercent.Valid() {
		return &vt.MaxPercent.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OrderStrategyUnionParam) GetMinPercent() *int64 {
	if vt := u.OfVwap; vt != nil && vt.MinPercent.Valid() {
		return &vt.MinPercent.Value
	} else if vt := u.OfTwap; vt != nil && vt.MinPercent.Valid() {
		return &vt.MinPercent.Value
	} else if vt := u.OfAp; vt != nil && vt.MinPercent.Valid() {
		return &vt.MinPercent.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's EndAt property, if present.
func (u OrderStrategyUnionParam) GetEndAt() *time.Time {
	if vt := u.OfSor; vt != nil && vt.EndAt.Valid() {
		return &vt.EndAt.Value
	} else if vt := u.OfVwap; vt != nil && vt.EndAt.Valid() {
		return &vt.EndAt.Value
	} else if vt := u.OfTwap; vt != nil && vt.EndAt.Valid() {
		return &vt.EndAt.Value
	} else if vt := u.OfAp; vt != nil && vt.EndAt.Valid() {
		return &vt.EndAt.Value
	} else if vt := u.OfPov; vt != nil && vt.EndAt.Valid() {
		return &vt.EndAt.Value
	} else if vt := u.OfDark; vt != nil && vt.EndAt.Valid() {
		return &vt.EndAt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's StartAt property, if present.
func (u OrderStrategyUnionParam) GetStartAt() *time.Time {
	if vt := u.OfSor; vt != nil && vt.StartAt.Valid() {
		return &vt.StartAt.Value
	} else if vt := u.OfVwap; vt != nil && vt.StartAt.Valid() {
		return &vt.StartAt.Value
	} else if vt := u.OfTwap; vt != nil && vt.StartAt.Valid() {
		return &vt.StartAt.Value
	} else if vt := u.OfAp; vt != nil && vt.StartAt.Valid() {
		return &vt.StartAt.Value
	} else if vt := u.OfPov; vt != nil && vt.StartAt.Valid() {
		return &vt.StartAt.Value
	} else if vt := u.OfDark; vt != nil && vt.StartAt.Valid() {
		return &vt.StartAt.Value
	}
	return nil
}

// Smart Order Router (default) - routes to best available venue
type OrderStrategySorParam struct {
	Type string `json:"type,omitzero,required"`
	SorStrategyParam
}

func (r OrderStrategySorParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*OrderStrategySorParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Volume Weighted Average Price - matches VWAP over a period
type OrderStrategyVwapParam struct {
	Type string `json:"type,omitzero,required"`
	VwapStrategyParam
}

func (r OrderStrategyVwapParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*OrderStrategyVwapParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Time Weighted Average Price - spreads execution evenly over time
type OrderStrategyTwapParam struct {
	Type string `json:"type,omitzero,required"`
	TwapStrategyParam
}

func (r OrderStrategyTwapParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*OrderStrategyTwapParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Arrival Price - aims to match price at order placement time
type OrderStrategyApParam struct {
	Type string `json:"type,omitzero,required"`
	ApStrategyParam
}

func (r OrderStrategyApParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*OrderStrategyApParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Percentage of Volume - participates as a percentage of market volume
type OrderStrategyPovParam struct {
	Type string `json:"type,omitzero,required"`
	PovStrategyParam
}

func (r OrderStrategyPovParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*OrderStrategyPovParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Dark Pool - routes to dark pool venues
type OrderStrategyDarkParam struct {
	Type string `json:"type,omitzero,required"`
	DarkStrategyParam
}

func (r OrderStrategyDarkParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*OrderStrategyDarkParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Direct Market Access - sends directly to a specified exchange
type OrderStrategyDmaParam struct {
	Type string `json:"type,omitzero,required"`
	DmaStrategyParam
}

func (r OrderStrategyDmaParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*OrderStrategyDmaParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Order type
type OrderType string

const (
	OrderTypeMarket    OrderType = "MARKET"
	OrderTypeLimit     OrderType = "LIMIT"
	OrderTypeStop      OrderType = "STOP"
	OrderTypeStopLimit OrderType = "STOP_LIMIT"
	OrderTypeOther     OrderType = "OTHER"
)

// Percentage of Volume strategy
type PovStrategy struct {
	// Target percentage of market volume to participate in (0-100)
	TargetPercent int64 `json:"target_percent,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TargetPercent respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
	BaseStrategyParamsResp
}

// Returns the unmodified JSON received from the API
func (r PovStrategy) RawJSON() string { return r.JSON.raw }
func (r *PovStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PovStrategy to a PovStrategyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PovStrategyParam.Overrides()
func (r PovStrategy) ToParam() PovStrategyParam {
	return param.Override[PovStrategyParam](json.RawMessage(r.RawJSON()))
}

// Percentage of Volume strategy
type PovStrategyParam struct {
	// Target percentage of market volume to participate in (0-100)
	TargetPercent int64 `json:"target_percent,required"`
	BaseStrategyParams
}

func (r PovStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*PovStrategyParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Side of an order
type Side string

const (
	SideBuy       Side = "BUY"
	SideSell      Side = "SELL"
	SideSellShort Side = "SELL_SHORT"
	SideOther     Side = "OTHER"
)

// Base parameters common to most algorithmic strategies
type SorStrategy struct {
	// UTC timestamp to end execution (defaults to market close)
	EndAt time.Time `json:"end_at,nullable" format:"date-time"`
	// UTC timestamp to start execution (defaults to order placement time)
	StartAt time.Time `json:"start_at,nullable" format:"date-time"`
	// Urgency level for execution aggressiveness
	//
	// Any of "SUPER_PASSIVE", "PASSIVE", "MODERATE", "AGGRESSIVE", "SUPER_AGGRESSIVE".
	Urgency Urgency `json:"urgency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndAt       respjson.Field
		StartAt     respjson.Field
		Urgency     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SorStrategy) RawJSON() string { return r.JSON.raw }
func (r *SorStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SorStrategy to a SorStrategyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SorStrategyParam.Overrides()
func (r SorStrategy) ToParam() SorStrategyParam {
	return param.Override[SorStrategyParam](json.RawMessage(r.RawJSON()))
}

// Base parameters common to most algorithmic strategies
type SorStrategyParam struct {
	// UTC timestamp to end execution (defaults to market close)
	EndAt param.Opt[time.Time] `json:"end_at,omitzero" format:"date-time"`
	// UTC timestamp to start execution (defaults to order placement time)
	StartAt param.Opt[time.Time] `json:"start_at,omitzero" format:"date-time"`
	// Urgency level for execution aggressiveness
	//
	// Any of "SUPER_PASSIVE", "PASSIVE", "MODERATE", "AGGRESSIVE", "SUPER_AGGRESSIVE".
	Urgency Urgency `json:"urgency,omitzero"`
	paramObj
}

func (r SorStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow SorStrategyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SorStrategyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time in force
type TimeInForce string

const (
	TimeInForceDay                 TimeInForce = "DAY"
	TimeInForceGoodTillCancel      TimeInForce = "GOOD_TILL_CANCEL"
	TimeInForceImmediateOrCancel   TimeInForce = "IMMEDIATE_OR_CANCEL"
	TimeInForceFillOrKill          TimeInForce = "FILL_OR_KILL"
	TimeInForceGoodTillDate        TimeInForce = "GOOD_TILL_DATE"
	TimeInForceAtTheOpening        TimeInForce = "AT_THE_OPENING"
	TimeInForceAtTheClose          TimeInForce = "AT_THE_CLOSE"
	TimeInForceGoodTillCrossing    TimeInForce = "GOOD_TILL_CROSSING"
	TimeInForceGoodThroughCrossing TimeInForce = "GOOD_THROUGH_CROSSING"
	TimeInForceAtCrossing          TimeInForce = "AT_CROSSING"
	TimeInForceOther               TimeInForce = "OTHER"
)

// Time Weighted Average Price strategy
type TwapStrategy struct {
	// Maximum percentage of market volume to participate in (0-50)
	MaxPercent int64 `json:"max_percent,nullable"`
	// Minimum percentage of market volume to participate in (0-100)
	MinPercent int64 `json:"min_percent,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxPercent  respjson.Field
		MinPercent  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseStrategyParamsResp
}

// Returns the unmodified JSON received from the API
func (r TwapStrategy) RawJSON() string { return r.JSON.raw }
func (r *TwapStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TwapStrategy to a TwapStrategyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TwapStrategyParam.Overrides()
func (r TwapStrategy) ToParam() TwapStrategyParam {
	return param.Override[TwapStrategyParam](json.RawMessage(r.RawJSON()))
}

// Time Weighted Average Price strategy
type TwapStrategyParam struct {
	// Maximum percentage of market volume to participate in (0-50)
	MaxPercent param.Opt[int64] `json:"max_percent,omitzero"`
	// Minimum percentage of market volume to participate in (0-100)
	MinPercent param.Opt[int64] `json:"min_percent,omitzero"`
	BaseStrategyParams
}

func (r TwapStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*TwapStrategyParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Urgency level for algorithmic execution
type Urgency string

const (
	UrgencySuperPassive    Urgency = "SUPER_PASSIVE"
	UrgencyPassive         Urgency = "PASSIVE"
	UrgencyModerate        Urgency = "MODERATE"
	UrgencyAggressive      Urgency = "AGGRESSIVE"
	UrgencySuperAggressive Urgency = "SUPER_AGGRESSIVE"
)

// Volume Weighted Average Price strategy
type VwapStrategy struct {
	// Maximum percentage of market volume to participate in (0-50)
	MaxPercent int64 `json:"max_percent,nullable"`
	// Minimum percentage of market volume to participate in (0-100)
	MinPercent int64 `json:"min_percent,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxPercent  respjson.Field
		MinPercent  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseStrategyParamsResp
}

// Returns the unmodified JSON received from the API
func (r VwapStrategy) RawJSON() string { return r.JSON.raw }
func (r *VwapStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VwapStrategy to a VwapStrategyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VwapStrategyParam.Overrides()
func (r VwapStrategy) ToParam() VwapStrategyParam {
	return param.Override[VwapStrategyParam](json.RawMessage(r.RawJSON()))
}

// Volume Weighted Average Price strategy
type VwapStrategyParam struct {
	// Maximum percentage of market volume to participate in (0-50)
	MaxPercent param.Opt[int64] `json:"max_percent,omitzero"`
	// Minimum percentage of market volume to participate in (0-100)
	MinPercent param.Opt[int64] `json:"min_percent,omitzero"`
	BaseStrategyParams
}

func (r VwapStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*VwapStrategyParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type ActiveV1AccountOrderCancelAllOrdersResponse struct {
	Data OrderList `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1AccountOrderCancelAllOrdersResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountOrderCancelAllOrdersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountOrderCancelOrderResponse struct {
	// A trading order with its current state and execution details.
	//
	// This is the unified API representation of an order across its lifecycle,
	// combining data from execution reports, order status queries, and parent/child
	// tracking.
	Data Order `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1AccountOrderCancelOrderResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountOrderCancelOrderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountOrderGetOrderByIDResponse struct {
	// A trading order with its current state and execution details.
	//
	// This is the unified API representation of an order across its lifecycle,
	// combining data from execution reports, order status queries, and parent/child
	// tracking.
	Data Order `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1AccountOrderGetOrderByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountOrderGetOrderByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountOrderGetOrdersResponse struct {
	Data OrderList `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1AccountOrderGetOrdersResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountOrderGetOrdersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountOrderReplaceOrderResponse struct {
	// A trading order with its current state and execution details.
	//
	// This is the unified API representation of an order across its lifecycle,
	// combining data from execution reports, order status queries, and parent/child
	// tracking.
	Data Order `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1AccountOrderReplaceOrderResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountOrderReplaceOrderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountOrderSubmitOrdersResponse struct {
	Data OrderList `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1AccountOrderSubmitOrdersResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountOrderSubmitOrdersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountOrderCancelAllOrdersParams struct {
	// Filter by security identifier (e.g., CUSIP, ISIN). Must be provided with
	// security_id_source.
	SecurityID param.Opt[string] `query:"security_id,omitzero" json:"-"`
	// Type of security identifier. Must be provided with security_id.
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "OTHER".
	SecurityIDSource SecurityIDSource `query:"security_id_source,omitzero" json:"-"`
	// Filter by security type (e.g., COMMON_STOCK, OPTION)
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	SecurityType SecurityType `query:"security_type,omitzero" json:"-"`
	// Filter by order side (BUY or SELL)
	//
	// Any of "BUY", "SELL", "SELL_SHORT", "OTHER".
	Side Side `query:"side,omitzero" json:"-"`
	// Filter by order type (e.g., MARKET, LIMIT)
	//
	// Any of "MARKET", "LIMIT", "STOP", "STOP_LIMIT", "OTHER".
	Type OrderType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1AccountOrderCancelAllOrdersParams]'s query
// parameters as `url.Values`.
func (r ActiveV1AccountOrderCancelAllOrdersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActiveV1AccountOrderCancelOrderParams struct {
	AccountID int64 `path:"account_id,required" json:"-"`
	paramObj
}

type ActiveV1AccountOrderGetOrderByIDParams struct {
	AccountID int64 `path:"account_id,required" json:"-"`
	paramObj
}

type ActiveV1AccountOrderGetOrdersParams struct {
	// The start date and time for the query range, inclusive (ISO 8601 format)
	From string `query:"from,required" json:"-"`
	// The end date and time for the query range, inclusive (ISO 8601 format)
	To string `query:"to,required" json:"-"`
	// The number of items to return per page (only used when page_token is not
	// provided)
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Filter by security ID
	SecurityID param.Opt[string] `query:"security_id,omitzero" json:"-"`
	// Filter by symbol
	Symbol param.Opt[string] `query:"symbol,omitzero" json:"-"`
	// Source for the security ID filter
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "OTHER".
	SecurityIDSource SecurityIDSource `query:"security_id_source,omitzero" json:"-"`
	// Security type filter (e.g., COMMON_STOCK, PREFERRED_STOCK)
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	SecurityType SecurityType `query:"security_type,omitzero" json:"-"`
	// Filter by order status
	//
	// Any of "PENDING_NEW", "NEW", "PARTIALLY_FILLED", "FILLED", "CANCELED",
	// "REJECTED", "EXPIRED", "PENDING_CANCEL", "PENDING_REPLACE", "REPLACED",
	// "DONE_FOR_DAY", "STOPPED", "SUSPENDED", "CALCULATED", "OTHER".
	Status OrderStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1AccountOrderGetOrdersParams]'s query parameters as
// `url.Values`.
func (r ActiveV1AccountOrderGetOrdersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActiveV1AccountOrderReplaceOrderParams struct {
	AccountID int64 `path:"account_id,required" json:"-"`
	// New limit price for the order
	LimitPrice param.Opt[string] `json:"limit_price,omitzero"`
	// New quantity for the order
	Quantity param.Opt[string] `json:"quantity,omitzero"`
	// New stop price for the order
	StopPrice param.Opt[string] `json:"stop_price,omitzero"`
	// New time in force for the order
	//
	// Any of "DAY", "GOOD_TILL_CANCEL", "IMMEDIATE_OR_CANCEL", "FILL_OR_KILL",
	// "GOOD_TILL_DATE", "AT_THE_OPENING", "AT_THE_CLOSE", "GOOD_TILL_CROSSING",
	// "GOOD_THROUGH_CROSSING", "AT_CROSSING", "OTHER".
	TimeInForce TimeInForce `json:"time_in_force,omitzero"`
	paramObj
}

func (r ActiveV1AccountOrderReplaceOrderParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1AccountOrderReplaceOrderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1AccountOrderReplaceOrderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountOrderSubmitOrdersParams struct {
	Body []ActiveV1AccountOrderSubmitOrdersParamsBody
	paramObj
}

func (r ActiveV1AccountOrderSubmitOrdersParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ActiveV1AccountOrderSubmitOrdersParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Request to submit a new order (PlaceOrderRequest from spec)
//
// The properties OrderID, OrderType, Quantity, SecurityType, Side, TimeInForce are
// required.
type ActiveV1AccountOrderSubmitOrdersParamsBody struct {
	// Client-provided unique ID (idempotency). Required to be unique per account.
	OrderID string `json:"order_id,required"`
	// Type of order
	//
	// Any of "MARKET", "LIMIT", "STOP", "STOP_LIMIT", "OTHER".
	OrderType OrderType `json:"order_type,omitzero,required"`
	// Quantity to trade. For COMMON_STOCK: shares (may be fractional if supported).
	// For OPTION (single-leg): contracts (must be an integer)
	Quantity string `json:"quantity,required"`
	// Type of security
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	SecurityType SecurityType `json:"security_type,omitzero,required"`
	// Side of the order
	//
	// Any of "BUY", "SELL", "SELL_SHORT", "OTHER".
	Side Side `json:"side,omitzero,required"`
	// Time in force
	//
	// Any of "DAY", "GOOD_TILL_CANCEL", "IMMEDIATE_OR_CANCEL", "FILL_OR_KILL",
	// "GOOD_TILL_DATE", "AT_THE_OPENING", "AT_THE_CLOSE", "GOOD_TILL_CROSSING",
	// "GOOD_THROUGH_CROSSING", "AT_CROSSING", "OTHER".
	TimeInForce TimeInForce `json:"time_in_force,omitzero,required"`
	// The timestamp when the order should expire (UTC). Required when time_in_force is
	// GOOD_TILL_DATE.
	ExpireAt param.Opt[time.Time] `json:"expire_at,omitzero" format:"date-time"`
	// Allow trading outside regular trading hours. Some brokers disallow options
	// outside RTH.
	ExtendedHours param.Opt[bool] `json:"extended_hours,omitzero"`
	// Limit price (required for LIMIT and STOP_LIMIT orders)
	LimitPrice param.Opt[string] `json:"limit_price,omitzero"`
	// Unique identifier for the instrument (CMS/CUSIP/ISIN/FIGI for equities or option
	// OPRA OSI). Required if symbol is not provided.
	SecurityID param.Opt[string] `json:"security_id,omitzero"`
	// Stop price (required for STOP and STOP_LIMIT orders)
	StopPrice param.Opt[string] `json:"stop_price,omitzero"`
	// Trading symbol. For equities, use the ticker symbol (e.g., "AAPL"). For options,
	// use the OSI symbol (e.g., "AAPL 250117C00190000"). If provided without
	// security_id, the system will derive security_id and source based on
	// security_type (CMS for equities, OPRA for options).
	Symbol param.Opt[string] `json:"symbol,omitzero"`
	// Execution venue to route the order to. If not specified, the system will choose
	// the best venue.
	Venue param.Opt[string] `json:"venue,omitzero"`
	// Required when security_type is OPTION. Specifies whether the order opens or
	// closes a position.
	//
	// Any of "OPEN", "CLOSE".
	PositionEffect string `json:"position_effect,omitzero"`
	// The source of the security identifier. Required if security_id is provided.
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "OTHER".
	SecurityIDSource SecurityIDSource `json:"security_id_source,omitzero"`
	// Execution strategy/router. Defaults to SOR where applicable.
	Strategy OrderStrategyUnionParam `json:"strategy,omitzero"`
	paramObj
}

func (r ActiveV1AccountOrderSubmitOrdersParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1AccountOrderSubmitOrdersParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1AccountOrderSubmitOrdersParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActiveV1AccountOrderSubmitOrdersParamsBody](
		"position_effect", "OPEN", "CLOSE",
	)
}
