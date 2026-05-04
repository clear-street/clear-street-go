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

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/internal/apiquery"
	shimjson "github.com/clear-street/clear-street-go/internal/encoding/json"
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/param"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Place, monitor, and manage trading orders.
//
// V1AccountOrderService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1AccountOrderService] method instead.
type V1AccountOrderService struct {
	options []option.RequestOption
}

// NewV1AccountOrderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1AccountOrderService(opts ...option.RequestOption) (r V1AccountOrderService) {
	r = V1AccountOrderService{}
	r.options = opts
	return
}

// Cancel all orders for an account
func (r *V1AccountOrderService) CancelAllOpenOrders(ctx context.Context, accountID int64, body V1AccountOrderCancelAllOpenOrdersParams, opts ...option.RequestOption) (res *V1AccountOrderCancelAllOpenOrdersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Cancel a specific order
func (r *V1AccountOrderService) CancelOpenOrder(ctx context.Context, orderID string, body V1AccountOrderCancelOpenOrderParams, opts ...option.RequestOption) (res *V1AccountOrderCancelOpenOrderResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%v/orders/%s", body.AccountID, url.PathEscape(orderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get Order By ID
func (r *V1AccountOrderService) GetOrderByID(ctx context.Context, orderID string, query V1AccountOrderGetOrderByIDParams, opts ...option.RequestOption) (res *V1AccountOrderGetOrderByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%v/orders/%s", query.AccountID, url.PathEscape(orderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List orders for an account with optional filtering
func (r *V1AccountOrderService) GetOrders(ctx context.Context, accountID int64, query V1AccountOrderGetOrdersParams, opts ...option.RequestOption) (res *V1AccountOrderGetOrdersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Replace an order with new parameters
func (r *V1AccountOrderService) ReplaceOrder(ctx context.Context, orderID string, params V1AccountOrderReplaceOrderParams, opts ...option.RequestOption) (res *V1AccountOrderReplaceOrderResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%v/orders/%s", params.AccountID, url.PathEscape(orderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Submit new orders
func (r *V1AccountOrderService) SubmitOrders(ctx context.Context, accountID int64, body V1AccountOrderSubmitOrdersParams, opts ...option.RequestOption) (res *V1AccountOrderSubmitOrdersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Arrival Price strategy
type ApStrategy struct {
	// Maximum percentage of market volume to participate in (0-100)
	MaxPercent APIDecimal64 `json:"max_percent" api:"nullable"`
	// Minimum percentage of market volume to participate in (0-100)
	MinPercent APIDecimal64 `json:"min_percent" api:"nullable"`
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
	MaxPercent param.Opt[APIDecimal64] `json:"max_percent,omitzero"`
	// Minimum percentage of market volume to participate in (0-100)
	MinPercent param.Opt[APIDecimal64] `json:"min_percent,omitzero"`
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
	EndAt time.Time `json:"end_at" api:"nullable" format:"date-time"`
	// UTC timestamp to start execution (defaults to order placement time)
	StartAt time.Time `json:"start_at" api:"nullable" format:"date-time"`
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
	MaxPercent APIDecimal64 `json:"max_percent" api:"nullable"`
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
	MaxPercent param.Opt[APIDecimal64] `json:"max_percent,omitzero"`
	BaseStrategyParams
}

func (r DarkStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*DarkStrategyParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Direct Market Access strategy
type DmaStrategy struct {
	// Destination exchange (MIC code)
	Destination string `json:"destination" api:"required"`
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
	Destination string `json:"destination" api:"required"`
	paramObj
}

func (r DmaStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow DmaStrategyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DmaStrategyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A trading order with its current state and execution details.
//
// This is the unified API representation of an order across its lifecycle,
// combining data from execution reports, order status queries, and parent/child
// tracking.
type Order struct {
	// Engine-assigned unique identifier for this order (UUID).
	ID string `json:"id" api:"required"`
	// Account placing the order
	AccountID int64 `json:"account_id" api:"required"`
	// Client-provided identifier echoed back (FIX tag 11).
	ClientOrderID string `json:"client_order_id" api:"required"`
	// Timestamp when order was created (UTC)
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Cumulative filled quantity
	FilledQuantity string `json:"filled_quantity" api:"required"`
	// OEMS instrument UUID for the traded instrument.
	InstrumentID string `json:"instrument_id" api:"required" format:"uuid"`
	// Type of security
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	InstrumentType SecurityType `json:"instrument_type" api:"required"`
	// Remaining unfilled quantity
	LeavesQuantity string `json:"leaves_quantity" api:"required"`
	// Type of order (MARKET, LIMIT, etc.)
	//
	// Any of "MARKET", "LIMIT", "STOP", "STOP_LIMIT", "TRAILING_STOP",
	// "TRAILING_STOP_LIMIT", "OTHER".
	OrderType OrderType `json:"order_type" api:"required"`
	// Total order quantity
	Quantity string `json:"quantity" api:"required"`
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
	// Whether the order is eligible for extended-hours trading.
	ExtendedHours bool `json:"extended_hours" api:"nullable"`
	// Limit offset for trailing stop-limit orders (signed)
	LimitOffset string `json:"limit_offset" api:"nullable"`
	// Limit price (for LIMIT and STOP_LIMIT orders)
	LimitPrice string `json:"limit_price" api:"nullable"`
	// Parent order queue state, present when the order is awaiting release or
	// released.
	//
	// Any of "AWAITING_RELEASE", "RELEASED".
	QueueState QueueState `json:"queue_state" api:"nullable"`
	// Scheduled release time for orders awaiting release.
	ReleasesAt time.Time `json:"releases_at" api:"nullable" format:"date-time"`
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
	// OEMS instrument ID of the option's underlying instrument. Populated only for
	// OPTIONS orders; `null` for non-options and for options whose underlier cannot be
	// resolved from the instrument cache.
	UnderlyingInstrumentID string `json:"underlying_instrument_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccountID              respjson.Field
		ClientOrderID          respjson.Field
		CreatedAt              respjson.Field
		FilledQuantity         respjson.Field
		InstrumentID           respjson.Field
		InstrumentType         respjson.Field
		LeavesQuantity         respjson.Field
		OrderType              respjson.Field
		Quantity               respjson.Field
		Side                   respjson.Field
		Status                 respjson.Field
		Symbol                 respjson.Field
		TimeInForce            respjson.Field
		UpdatedAt              respjson.Field
		Venue                  respjson.Field
		AverageFillPrice       respjson.Field
		Details                respjson.Field
		ExpiresAt              respjson.Field
		ExtendedHours          respjson.Field
		LimitOffset            respjson.Field
		LimitPrice             respjson.Field
		QueueState             respjson.Field
		ReleasesAt             respjson.Field
		StopPrice              respjson.Field
		Strategy               respjson.Field
		TrailingOffsetAmt      respjson.Field
		TrailingOffsetAmtType  respjson.Field
		TrailingWatermarkPx    respjson.Field
		TrailingWatermarkTs    respjson.Field
		UnderlyingInstrumentID respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Order) RawJSON() string { return r.JSON.raw }
func (r *Order) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderList []Order

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
	// This field is from variant [OrderStrategySor], [OrderStrategyVwap],
	// [OrderStrategyTwap], [OrderStrategyAp], [OrderStrategyPov], [OrderStrategyDark].
	EndAt time.Time `json:"end_at"`
	// This field is from variant [OrderStrategySor], [OrderStrategyVwap],
	// [OrderStrategyTwap], [OrderStrategyAp], [OrderStrategyPov], [OrderStrategyDark].
	StartAt time.Time `json:"start_at"`
	// This field is from variant [OrderStrategySor], [OrderStrategyVwap],
	// [OrderStrategyTwap], [OrderStrategyAp], [OrderStrategyPov], [OrderStrategyDark].
	Urgency Urgency `json:"urgency"`
	Type    string  `json:"type"`
	// This field is from variant [OrderStrategyVwap].
	MaxPercent APIDecimal64 `json:"max_percent"`
	// This field is from variant [OrderStrategyVwap].
	MinPercent APIDecimal64 `json:"min_percent"`
	// This field is from variant [OrderStrategyPov].
	TargetPercent APIDecimal64 `json:"target_percent"`
	// This field is from variant [OrderStrategyDma].
	Destination string `json:"destination"`
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
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseStrategyParamsResp
}

// Returns the unmodified JSON received from the API
func (r OrderStrategySor) RawJSON() string { return r.JSON.raw }
func (r *OrderStrategySor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volume Weighted Average Price - matches VWAP over a period
type OrderStrategyVwap struct {
	// Any of "VWAP".
	Type string `json:"type" api:"required"`
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
	Type string `json:"type" api:"required"`
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
	Type string `json:"type" api:"required"`
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
	Type string `json:"type" api:"required"`
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
	Type string `json:"type" api:"required"`
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
	Type string `json:"type" api:"required"`
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

func OrderStrategyParamOfPov(targetPercent APIDecimal64, type_ string) OrderStrategyUnionParam {
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

func OrderStrategyParamOfDma(destination string, type_ string) OrderStrategyUnionParam {
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

// Smart Order Router (default) - routes to best available venue
type OrderStrategySorParam struct {
	Type string `json:"type,omitzero" api:"required"`
	BaseStrategyParams
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
	Type string `json:"type,omitzero" api:"required"`
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
	Type string `json:"type,omitzero" api:"required"`
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
	Type string `json:"type,omitzero" api:"required"`
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
	Type string `json:"type,omitzero" api:"required"`
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
	Type string `json:"type,omitzero" api:"required"`
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
	Type string `json:"type,omitzero" api:"required"`
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
	OrderTypeMarket            OrderType = "MARKET"
	OrderTypeLimit             OrderType = "LIMIT"
	OrderTypeStop              OrderType = "STOP"
	OrderTypeStopLimit         OrderType = "STOP_LIMIT"
	OrderTypeTrailingStop      OrderType = "TRAILING_STOP"
	OrderTypeTrailingStopLimit OrderType = "TRAILING_STOP_LIMIT"
	OrderTypeOther             OrderType = "OTHER"
)

// Percentage of Volume strategy
type PovStrategy struct {
	// Target percentage of market volume to participate in (0-100)
	TargetPercent APIDecimal64 `json:"target_percent" api:"required"`
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
	TargetPercent APIDecimal64 `json:"target_percent" api:"required"`
	BaseStrategyParams
}

func (r PovStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*PovStrategyParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Parent order queue or hold state.
type QueueState string

const (
	QueueStateAwaitingRelease QueueState = "AWAITING_RELEASE"
	QueueStateReleased        QueueState = "RELEASED"
)

// Side of an order
type Side string

const (
	SideBuy       Side = "BUY"
	SideSell      Side = "SELL"
	SideSellShort Side = "SELL_SHORT"
	SideOther     Side = "OTHER"
)

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

// Trailing offset type for trailing stop orders.
type TrailingOffsetType string

const (
	TrailingOffsetTypePrice      TrailingOffsetType = "PRICE"
	TrailingOffsetTypePercentBps TrailingOffsetType = "PERCENT_BPS"
)

// Time Weighted Average Price strategy
type TwapStrategy struct {
	// Maximum percentage of market volume to participate in (0-50)
	MaxPercent APIDecimal64 `json:"max_percent" api:"nullable"`
	// Minimum percentage of market volume to participate in (0-100)
	MinPercent APIDecimal64 `json:"min_percent" api:"nullable"`
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
	MaxPercent param.Opt[APIDecimal64] `json:"max_percent,omitzero"`
	// Minimum percentage of market volume to participate in (0-100)
	MinPercent param.Opt[APIDecimal64] `json:"min_percent,omitzero"`
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
	MaxPercent APIDecimal64 `json:"max_percent" api:"nullable"`
	// Minimum percentage of market volume to participate in (0-100)
	MinPercent APIDecimal64 `json:"min_percent" api:"nullable"`
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
	MaxPercent param.Opt[APIDecimal64] `json:"max_percent,omitzero"`
	// Minimum percentage of market volume to participate in (0-100)
	MinPercent param.Opt[APIDecimal64] `json:"min_percent,omitzero"`
	BaseStrategyParams
}

func (r VwapStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*VwapStrategyParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type V1AccountOrderCancelAllOpenOrdersResponse struct {
	Data OrderList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountOrderCancelAllOpenOrdersResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountOrderCancelAllOpenOrdersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountOrderCancelOpenOrderResponse struct {
	// A trading order with its current state and execution details.
	//
	// This is the unified API representation of an order across its lifecycle,
	// combining data from execution reports, order status queries, and parent/child
	// tracking.
	Data Order `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountOrderCancelOpenOrderResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountOrderCancelOpenOrderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountOrderGetOrderByIDResponse struct {
	// A trading order with its current state and execution details.
	//
	// This is the unified API representation of an order across its lifecycle,
	// combining data from execution reports, order status queries, and parent/child
	// tracking.
	Data Order `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountOrderGetOrderByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountOrderGetOrderByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountOrderGetOrdersResponse struct {
	Data OrderList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountOrderGetOrdersResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountOrderGetOrdersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountOrderReplaceOrderResponse struct {
	// A trading order with its current state and execution details.
	//
	// This is the unified API representation of an order across its lifecycle,
	// combining data from execution reports, order status queries, and parent/child
	// tracking.
	Data Order `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountOrderReplaceOrderResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountOrderReplaceOrderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountOrderSubmitOrdersResponse struct {
	Data OrderList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountOrderSubmitOrdersResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountOrderSubmitOrdersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountOrderCancelAllOpenOrdersParams struct {
	// Comma-separated OEMS instrument UUIDs
	InstrumentIDs []string `query:"instrument_ids,omitzero" format:"uuid" json:"-"`
	// Filter by instrument type (e.g., COMMON_STOCK, OPTION)
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	InstrumentType V1AccountOrderCancelAllOpenOrdersParamsInstrumentType `query:"instrument_type,omitzero" json:"-"`
	// Filter by order side (BUY or SELL)
	//
	// Any of "BUY", "SELL", "SELL_SHORT", "OTHER".
	Side V1AccountOrderCancelAllOpenOrdersParamsSide `query:"side,omitzero" json:"-"`
	// Filter by order type (e.g., MARKET, LIMIT)
	//
	// Any of "MARKET", "LIMIT", "STOP", "STOP_LIMIT", "TRAILING_STOP",
	// "TRAILING_STOP_LIMIT", "OTHER".
	Type V1AccountOrderCancelAllOpenOrdersParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1AccountOrderCancelAllOpenOrdersParams]'s query parameters
// as `url.Values`.
func (r V1AccountOrderCancelAllOpenOrdersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by instrument type (e.g., COMMON_STOCK, OPTION)
type V1AccountOrderCancelAllOpenOrdersParamsInstrumentType string

const (
	V1AccountOrderCancelAllOpenOrdersParamsInstrumentTypeCommonStock    V1AccountOrderCancelAllOpenOrdersParamsInstrumentType = "COMMON_STOCK"
	V1AccountOrderCancelAllOpenOrdersParamsInstrumentTypePreferredStock V1AccountOrderCancelAllOpenOrdersParamsInstrumentType = "PREFERRED_STOCK"
	V1AccountOrderCancelAllOpenOrdersParamsInstrumentTypeCorporateBond  V1AccountOrderCancelAllOpenOrdersParamsInstrumentType = "CORPORATE_BOND"
	V1AccountOrderCancelAllOpenOrdersParamsInstrumentTypeOption         V1AccountOrderCancelAllOpenOrdersParamsInstrumentType = "OPTION"
	V1AccountOrderCancelAllOpenOrdersParamsInstrumentTypeFuture         V1AccountOrderCancelAllOpenOrdersParamsInstrumentType = "FUTURE"
	V1AccountOrderCancelAllOpenOrdersParamsInstrumentTypeWarrant        V1AccountOrderCancelAllOpenOrdersParamsInstrumentType = "WARRANT"
	V1AccountOrderCancelAllOpenOrdersParamsInstrumentTypeCash           V1AccountOrderCancelAllOpenOrdersParamsInstrumentType = "CASH"
	V1AccountOrderCancelAllOpenOrdersParamsInstrumentTypeOther          V1AccountOrderCancelAllOpenOrdersParamsInstrumentType = "OTHER"
)

// Filter by order side (BUY or SELL)
type V1AccountOrderCancelAllOpenOrdersParamsSide string

const (
	V1AccountOrderCancelAllOpenOrdersParamsSideBuy       V1AccountOrderCancelAllOpenOrdersParamsSide = "BUY"
	V1AccountOrderCancelAllOpenOrdersParamsSideSell      V1AccountOrderCancelAllOpenOrdersParamsSide = "SELL"
	V1AccountOrderCancelAllOpenOrdersParamsSideSellShort V1AccountOrderCancelAllOpenOrdersParamsSide = "SELL_SHORT"
	V1AccountOrderCancelAllOpenOrdersParamsSideOther     V1AccountOrderCancelAllOpenOrdersParamsSide = "OTHER"
)

// Filter by order type (e.g., MARKET, LIMIT)
type V1AccountOrderCancelAllOpenOrdersParamsType string

const (
	V1AccountOrderCancelAllOpenOrdersParamsTypeMarket            V1AccountOrderCancelAllOpenOrdersParamsType = "MARKET"
	V1AccountOrderCancelAllOpenOrdersParamsTypeLimit             V1AccountOrderCancelAllOpenOrdersParamsType = "LIMIT"
	V1AccountOrderCancelAllOpenOrdersParamsTypeStop              V1AccountOrderCancelAllOpenOrdersParamsType = "STOP"
	V1AccountOrderCancelAllOpenOrdersParamsTypeStopLimit         V1AccountOrderCancelAllOpenOrdersParamsType = "STOP_LIMIT"
	V1AccountOrderCancelAllOpenOrdersParamsTypeTrailingStop      V1AccountOrderCancelAllOpenOrdersParamsType = "TRAILING_STOP"
	V1AccountOrderCancelAllOpenOrdersParamsTypeTrailingStopLimit V1AccountOrderCancelAllOpenOrdersParamsType = "TRAILING_STOP_LIMIT"
	V1AccountOrderCancelAllOpenOrdersParamsTypeOther             V1AccountOrderCancelAllOpenOrdersParamsType = "OTHER"
)

type V1AccountOrderCancelOpenOrderParams struct {
	AccountID int64 `path:"account_id" api:"required" json:"-"`
	paramObj
}

type V1AccountOrderGetOrderByIDParams struct {
	AccountID int64 `path:"account_id" api:"required" json:"-"`
	paramObj
}

type V1AccountOrderGetOrdersParams struct {
	// The start date and time for the query range, inclusive (ISO 8601 format)
	From     param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	PageSize param.Opt[int64]     `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Filter by symbol
	Symbol param.Opt[string] `query:"symbol,omitzero" json:"-"`
	// The end date and time for the query range, inclusive (ISO 8601 format)
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Comma-separated OEMS instrument UUIDs. Matches options orders whose resolved
	// underlier is any of the given IDs.
	UnderlyingInstrumentIDs param.Opt[string] `query:"underlying_instrument_ids,omitzero" json:"-"`
	// Comma-separated OEMS instrument UUIDs
	InstrumentIDs []string `query:"instrument_ids,omitzero" format:"uuid" json:"-"`
	// Instrument type filter (e.g., COMMON_STOCK, OPTION)
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	InstrumentType V1AccountOrderGetOrdersParamsInstrumentType `query:"instrument_type,omitzero" json:"-"`
	// Comma-separated order statuses to filter by
	//
	// Any of "PENDING_NEW", "NEW", "PARTIALLY_FILLED", "FILLED", "CANCELED",
	// "REJECTED", "EXPIRED", "PENDING_CANCEL", "PENDING_REPLACE", "REPLACED",
	// "DONE_FOR_DAY", "STOPPED", "SUSPENDED", "CALCULATED", "OTHER".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1AccountOrderGetOrdersParams]'s query parameters as
// `url.Values`.
func (r V1AccountOrderGetOrdersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Instrument type filter (e.g., COMMON_STOCK, OPTION)
type V1AccountOrderGetOrdersParamsInstrumentType string

const (
	V1AccountOrderGetOrdersParamsInstrumentTypeCommonStock    V1AccountOrderGetOrdersParamsInstrumentType = "COMMON_STOCK"
	V1AccountOrderGetOrdersParamsInstrumentTypePreferredStock V1AccountOrderGetOrdersParamsInstrumentType = "PREFERRED_STOCK"
	V1AccountOrderGetOrdersParamsInstrumentTypeCorporateBond  V1AccountOrderGetOrdersParamsInstrumentType = "CORPORATE_BOND"
	V1AccountOrderGetOrdersParamsInstrumentTypeOption         V1AccountOrderGetOrdersParamsInstrumentType = "OPTION"
	V1AccountOrderGetOrdersParamsInstrumentTypeFuture         V1AccountOrderGetOrdersParamsInstrumentType = "FUTURE"
	V1AccountOrderGetOrdersParamsInstrumentTypeWarrant        V1AccountOrderGetOrdersParamsInstrumentType = "WARRANT"
	V1AccountOrderGetOrdersParamsInstrumentTypeCash           V1AccountOrderGetOrdersParamsInstrumentType = "CASH"
	V1AccountOrderGetOrdersParamsInstrumentTypeOther          V1AccountOrderGetOrdersParamsInstrumentType = "OTHER"
)

type V1AccountOrderReplaceOrderParams struct {
	AccountID int64 `path:"account_id" api:"required" json:"-"`
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

func (r V1AccountOrderReplaceOrderParams) MarshalJSON() (data []byte, err error) {
	type shadow V1AccountOrderReplaceOrderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AccountOrderReplaceOrderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountOrderSubmitOrdersParams struct {
	Orders []V1AccountOrderSubmitOrdersParamsOrderUnion
	paramObj
}

func (r V1AccountOrderSubmitOrdersParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Orders)
}
func (r *V1AccountOrderSubmitOrdersParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V1AccountOrderSubmitOrdersParamsOrderUnion struct {
	OfV1AccountOrderSubmitOrderssOrderNewOrderMultilegRequest *V1AccountOrderSubmitOrdersParamsOrderNewOrderMultilegRequest `json:",omitzero,inline"`
	OfV1AccountOrderSubmitOrderssOrderNewOrderRequest         *V1AccountOrderSubmitOrdersParamsOrderNewOrderRequest         `json:",omitzero,inline"`
	paramUnion
}

func (u V1AccountOrderSubmitOrdersParamsOrderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfV1AccountOrderSubmitOrderssOrderNewOrderMultilegRequest, u.OfV1AccountOrderSubmitOrderssOrderNewOrderRequest)
}
func (u *V1AccountOrderSubmitOrdersParamsOrderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Multileg strategy order request
//
// The properties Legs, OrderType, TimeInForce are required.
type V1AccountOrderSubmitOrdersParamsOrderNewOrderMultilegRequest struct {
	// Legs that compose the strategy.
	Legs []V1AccountOrderSubmitOrdersParamsOrderNewOrderMultilegRequestLeg `json:"legs,omitzero" api:"required"`
	// Type of order (currently MARKET or LIMIT for multileg strategy submission)
	//
	// Any of "MARKET", "LIMIT", "STOP", "STOP_LIMIT", "TRAILING_STOP",
	// "TRAILING_STOP_LIMIT", "OTHER".
	OrderType OrderType `json:"order_type,omitzero" api:"required"`
	// Time in force
	//
	// Any of "DAY", "GOOD_TILL_CANCEL", "IMMEDIATE_OR_CANCEL", "FILL_OR_KILL",
	// "GOOD_TILL_DATE", "AT_THE_OPENING", "AT_THE_CLOSE", "GOOD_TILL_CROSSING",
	// "GOOD_THROUGH_CROSSING", "AT_CROSSING", "OTHER".
	TimeInForce TimeInForce `json:"time_in_force,omitzero" api:"required"`
	// Optional client-provided unique ID (idempotency). Required to be unique per
	// account.
	ID param.Opt[string] `json:"id,omitzero"`
	// Strategy price, required for LIMIT orders.
	LimitPrice param.Opt[string] `json:"limit_price,omitzero"`
	// Optional strategy-level quantity. Multiplies leg quantities. Defaults to 1.
	Quantity param.Opt[string] `json:"quantity,omitzero"`
	paramObj
}

func (r V1AccountOrderSubmitOrdersParamsOrderNewOrderMultilegRequest) MarshalJSON() (data []byte, err error) {
	type shadow V1AccountOrderSubmitOrdersParamsOrderNewOrderMultilegRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AccountOrderSubmitOrdersParamsOrderNewOrderMultilegRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single leg in a multileg strategy request.
//
// The properties InstrumentType, Ratio, Security, Side are required.
type V1AccountOrderSubmitOrdersParamsOrderNewOrderMultilegRequestLeg struct {
	// Security type for the leg.
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	InstrumentType SecurityType `json:"instrument_type,omitzero" api:"required"`
	// Ratio for the leg.
	Ratio string `json:"ratio" api:"required"`
	// Trading symbol (e.g. "AAPL" or OSI symbol for options)
	Security string `json:"security" api:"required"`
	// Leg side.
	//
	// Any of "BUY", "SELL", "SELL_SHORT", "OTHER".
	Side Side `json:"side,omitzero" api:"required"`
	// Optional leg reference identifier.
	ID param.Opt[string] `json:"id,omitzero"`
	// Optional leg position effect.
	//
	// Any of "OPEN", "CLOSE".
	PositionEffect string `json:"position_effect,omitzero"`
	paramObj
}

func (r V1AccountOrderSubmitOrdersParamsOrderNewOrderMultilegRequestLeg) MarshalJSON() (data []byte, err error) {
	type shadow V1AccountOrderSubmitOrdersParamsOrderNewOrderMultilegRequestLeg
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AccountOrderSubmitOrdersParamsOrderNewOrderMultilegRequestLeg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AccountOrderSubmitOrdersParamsOrderNewOrderMultilegRequestLeg](
		"position_effect", "OPEN", "CLOSE",
	)
}

// Single-leg order request
//
// The properties InstrumentType, OrderType, Quantity, Side, TimeInForce are
// required.
type V1AccountOrderSubmitOrdersParamsOrderNewOrderRequest struct {
	// Type of security
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	InstrumentType SecurityType `json:"instrument_type,omitzero" api:"required"`
	// Type of order
	//
	// Any of "MARKET", "LIMIT", "STOP", "STOP_LIMIT", "TRAILING_STOP",
	// "TRAILING_STOP_LIMIT", "OTHER".
	OrderType OrderType `json:"order_type,omitzero" api:"required"`
	// Quantity to trade. For COMMON_STOCK: shares (may be fractional if supported).
	// For OPTION (single-leg): contracts (must be an integer)
	Quantity string `json:"quantity" api:"required"`
	// Side of the order
	//
	// Any of "BUY", "SELL", "SELL_SHORT", "OTHER".
	Side Side `json:"side,omitzero" api:"required"`
	// Time in force
	//
	// Any of "DAY", "GOOD_TILL_CANCEL", "IMMEDIATE_OR_CANCEL", "FILL_OR_KILL",
	// "GOOD_TILL_DATE", "AT_THE_OPENING", "AT_THE_CLOSE", "GOOD_TILL_CROSSING",
	// "GOOD_THROUGH_CROSSING", "AT_CROSSING", "OTHER".
	TimeInForce TimeInForce `json:"time_in_force,omitzero" api:"required"`
	// Optional client-provided unique ID (idempotency). Required to be unique per
	// account.
	ID param.Opt[string] `json:"id,omitzero"`
	// The timestamp when the order should expire (UTC). Required when time_in_force is
	// GOOD_TILL_DATE.
	ExpiresAt param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	// Allow trading outside regular trading hours. Some brokers disallow options
	// outside RTH.
	ExtendedHours param.Opt[bool] `json:"extended_hours,omitzero"`
	// OEMS instrument UUID. Either `symbol` or `instrument_id` must be provided.
	InstrumentID param.Opt[string] `json:"instrument_id,omitzero" format:"uuid"`
	// Limit offset for trailing stop-limit orders (signed)
	LimitOffset param.Opt[string] `json:"limit_offset,omitzero"`
	// Limit price (required for LIMIT and STOP_LIMIT orders)
	LimitPrice param.Opt[string] `json:"limit_price,omitzero"`
	// Stop price (required for STOP and STOP_LIMIT orders)
	StopPrice param.Opt[string] `json:"stop_price,omitzero"`
	// Trading symbol. For equities, use the ticker symbol (e.g., "AAPL"). For options,
	// use the OSI symbol (e.g., "AAPL 250117C00190000"). Either `symbol` or
	// `instrument_id` must be provided.
	Symbol param.Opt[string] `json:"symbol,omitzero"`
	// Trailing offset amount (required for trailing orders)
	TrailingOffsetAmt param.Opt[string] `json:"trailing_offset_amt,omitzero"`
	// Required when instrument_type is OPTION. Specifies whether the order opens or
	// closes a position.
	//
	// Any of "OPEN", "CLOSE".
	PositionEffect string `json:"position_effect,omitzero"`
	// Execution strategy/router. Defaults to SOR where applicable.
	Strategy OrderStrategyUnionParam `json:"strategy,omitzero"`
	// Trailing offset type (PRICE or PERCENT_BPS)
	//
	// Any of "PRICE", "PERCENT_BPS".
	TrailingOffsetAmtType TrailingOffsetType `json:"trailing_offset_amt_type,omitzero"`
	paramObj
}

func (r V1AccountOrderSubmitOrdersParamsOrderNewOrderRequest) MarshalJSON() (data []byte, err error) {
	type shadow V1AccountOrderSubmitOrdersParamsOrderNewOrderRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AccountOrderSubmitOrdersParamsOrderNewOrderRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1AccountOrderSubmitOrdersParamsOrderNewOrderRequest](
		"position_effect", "OPEN", "CLOSE",
	)
}
