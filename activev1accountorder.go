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
// ActiveV1AccountOrderService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1AccountOrderService] method instead.
type ActiveV1AccountOrderService struct {
	options []option.RequestOption
}

// NewActiveV1AccountOrderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1AccountOrderService(opts ...option.RequestOption) (r ActiveV1AccountOrderService) {
	r = ActiveV1AccountOrderService{}
	r.options = opts
	return
}

// All filter parameters can be used independently or combined. The only constraint
// is that `security_id` and `security_id_source` must be provided together if
// either is specified.
func (r *ActiveV1AccountOrderService) CancelAllOrders(ctx context.Context, accountID int64, body ActiveV1AccountOrderCancelAllOrdersParams, opts ...option.RequestOption) (res *ActiveV1AccountOrderCancelAllOrdersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Cancel a specific order
func (r *ActiveV1AccountOrderService) CancelOrder(ctx context.Context, orderID string, body ActiveV1AccountOrderCancelOrderParams, opts ...option.RequestOption) (res *ActiveV1AccountOrderCancelOrderResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/accounts/%v/orders/%s", body.AccountID, url.PathEscape(orderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get order by ID
func (r *ActiveV1AccountOrderService) GetOrderByID(ctx context.Context, orderID string, query ActiveV1AccountOrderGetOrderByIDParams, opts ...option.RequestOption) (res *ActiveV1AccountOrderGetOrderByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/accounts/%v/orders/%s", query.AccountID, url.PathEscape(orderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List orders for an account with optional filtering
func (r *ActiveV1AccountOrderService) GetOrders(ctx context.Context, accountID int64, query ActiveV1AccountOrderGetOrdersParams, opts ...option.RequestOption) (res *ActiveV1AccountOrderGetOrdersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Replace an order with new parameters
func (r *ActiveV1AccountOrderService) ReplaceOrder(ctx context.Context, orderID string, params ActiveV1AccountOrderReplaceOrderParams, opts ...option.RequestOption) (res *ActiveV1AccountOrderReplaceOrderResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/accounts/%v/orders/%s", params.AccountID, url.PathEscape(orderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Submit new orders
func (r *ActiveV1AccountOrderService) SubmitOrders(ctx context.Context, accountID int64, body ActiveV1AccountOrderSubmitOrdersParams, opts ...option.RequestOption) (res *ActiveV1AccountOrderSubmitOrdersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v/orders", accountID)
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

type ActiveV1AccountOrderCancelAllOrdersResponse struct {
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
func (r ActiveV1AccountOrderGetOrderByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountOrderGetOrderByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountOrderGetOrdersResponse struct {
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
func (r ActiveV1AccountOrderReplaceOrderResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountOrderReplaceOrderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountOrderSubmitOrdersResponse struct {
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
func (r ActiveV1AccountOrderSubmitOrdersResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountOrderSubmitOrdersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountOrderCancelAllOrdersParams struct {
	// Filter by security ID(s). Accepts single value or indexed array.
	//
	// Examples:
	//
	// - Single: `security_id=037833100`
	// - Multiple: `security_id[0]=037833100&security_id[1]=594918104`
	SecurityID []string `query:"security_id,omitzero" json:"-"`
	// Source(s) for the security ID filter. Must match the count and order of
	// security_id.
	//
	// Examples:
	//
	// - Single: `security_id_source=CUSIP`
	// - Multiple: `security_id_source[0]=CUSIP&security_id_source[1]=FIGI`
	SecurityIDSource []string `query:"security_id_source,omitzero" json:"-"`
	// Filter by security type (e.g., COMMON_STOCK, OPTION)
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	SecurityType ActiveV1AccountOrderCancelAllOrdersParamsSecurityType `query:"security_type,omitzero" json:"-"`
	// Filter by order side (BUY or SELL)
	//
	// Any of "BUY", "SELL", "SELL_SHORT", "OTHER".
	Side ActiveV1AccountOrderCancelAllOrdersParamsSide `query:"side,omitzero" json:"-"`
	// Filter by order type (e.g., MARKET, LIMIT)
	//
	// Any of "MARKET", "LIMIT", "STOP", "STOP_LIMIT", "TRAILING_STOP",
	// "TRAILING_STOP_LIMIT", "OTHER".
	Type ActiveV1AccountOrderCancelAllOrdersParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1AccountOrderCancelAllOrdersParams]'s query
// parameters as `url.Values`.
func (r ActiveV1AccountOrderCancelAllOrdersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by security type (e.g., COMMON_STOCK, OPTION)
type ActiveV1AccountOrderCancelAllOrdersParamsSecurityType string

const (
	ActiveV1AccountOrderCancelAllOrdersParamsSecurityTypeCommonStock    ActiveV1AccountOrderCancelAllOrdersParamsSecurityType = "COMMON_STOCK"
	ActiveV1AccountOrderCancelAllOrdersParamsSecurityTypePreferredStock ActiveV1AccountOrderCancelAllOrdersParamsSecurityType = "PREFERRED_STOCK"
	ActiveV1AccountOrderCancelAllOrdersParamsSecurityTypeCorporateBond  ActiveV1AccountOrderCancelAllOrdersParamsSecurityType = "CORPORATE_BOND"
	ActiveV1AccountOrderCancelAllOrdersParamsSecurityTypeOption         ActiveV1AccountOrderCancelAllOrdersParamsSecurityType = "OPTION"
	ActiveV1AccountOrderCancelAllOrdersParamsSecurityTypeFuture         ActiveV1AccountOrderCancelAllOrdersParamsSecurityType = "FUTURE"
	ActiveV1AccountOrderCancelAllOrdersParamsSecurityTypeWarrant        ActiveV1AccountOrderCancelAllOrdersParamsSecurityType = "WARRANT"
	ActiveV1AccountOrderCancelAllOrdersParamsSecurityTypeCash           ActiveV1AccountOrderCancelAllOrdersParamsSecurityType = "CASH"
	ActiveV1AccountOrderCancelAllOrdersParamsSecurityTypeOther          ActiveV1AccountOrderCancelAllOrdersParamsSecurityType = "OTHER"
)

// Filter by order side (BUY or SELL)
type ActiveV1AccountOrderCancelAllOrdersParamsSide string

const (
	ActiveV1AccountOrderCancelAllOrdersParamsSideBuy       ActiveV1AccountOrderCancelAllOrdersParamsSide = "BUY"
	ActiveV1AccountOrderCancelAllOrdersParamsSideSell      ActiveV1AccountOrderCancelAllOrdersParamsSide = "SELL"
	ActiveV1AccountOrderCancelAllOrdersParamsSideSellShort ActiveV1AccountOrderCancelAllOrdersParamsSide = "SELL_SHORT"
	ActiveV1AccountOrderCancelAllOrdersParamsSideOther     ActiveV1AccountOrderCancelAllOrdersParamsSide = "OTHER"
)

// Filter by order type (e.g., MARKET, LIMIT)
type ActiveV1AccountOrderCancelAllOrdersParamsType string

const (
	ActiveV1AccountOrderCancelAllOrdersParamsTypeMarket            ActiveV1AccountOrderCancelAllOrdersParamsType = "MARKET"
	ActiveV1AccountOrderCancelAllOrdersParamsTypeLimit             ActiveV1AccountOrderCancelAllOrdersParamsType = "LIMIT"
	ActiveV1AccountOrderCancelAllOrdersParamsTypeStop              ActiveV1AccountOrderCancelAllOrdersParamsType = "STOP"
	ActiveV1AccountOrderCancelAllOrdersParamsTypeStopLimit         ActiveV1AccountOrderCancelAllOrdersParamsType = "STOP_LIMIT"
	ActiveV1AccountOrderCancelAllOrdersParamsTypeTrailingStop      ActiveV1AccountOrderCancelAllOrdersParamsType = "TRAILING_STOP"
	ActiveV1AccountOrderCancelAllOrdersParamsTypeTrailingStopLimit ActiveV1AccountOrderCancelAllOrdersParamsType = "TRAILING_STOP_LIMIT"
	ActiveV1AccountOrderCancelAllOrdersParamsTypeOther             ActiveV1AccountOrderCancelAllOrdersParamsType = "OTHER"
)

type ActiveV1AccountOrderCancelOrderParams struct {
	AccountID int64 `path:"account_id" api:"required" json:"-"`
	paramObj
}

type ActiveV1AccountOrderGetOrderByIDParams struct {
	AccountID int64 `path:"account_id" api:"required" json:"-"`
	paramObj
}

type ActiveV1AccountOrderGetOrdersParams struct {
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
	// Filter by security ID(s). Accepts single value or indexed array.
	//
	// Examples:
	//
	// - Single: `security_id=037833100`
	// - Multiple: `security_id[0]=037833100&security_id[1]=594918104`
	SecurityID []string `query:"security_id,omitzero" json:"-"`
	// Source(s) for the security ID filter. Must match the count and order of
	// security_id.
	//
	// Examples:
	//
	// - Single: `security_id_source=CUSIP`
	// - Multiple: `security_id_source[0]=CUSIP&security_id_source[1]=FIGI`
	SecurityIDSource []string `query:"security_id_source,omitzero" json:"-"`
	// Security type filter (e.g., COMMON_STOCK, PREFERRED_STOCK)
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	SecurityType ActiveV1AccountOrderGetOrdersParamsSecurityType `query:"security_type,omitzero" json:"-"`
	// Comma-separated order statuses to filter by
	//
	// Any of "PENDING_NEW", "NEW", "PARTIALLY_FILLED", "FILLED", "CANCELED",
	// "REJECTED", "EXPIRED", "PENDING_CANCEL", "PENDING_REPLACE", "REPLACED",
	// "DONE_FOR_DAY", "STOPPED", "SUSPENDED", "CALCULATED", "OTHER".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1AccountOrderGetOrdersParams]'s query parameters as
// `url.Values`.
func (r ActiveV1AccountOrderGetOrdersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Security type filter (e.g., COMMON_STOCK, PREFERRED_STOCK)
type ActiveV1AccountOrderGetOrdersParamsSecurityType string

const (
	ActiveV1AccountOrderGetOrdersParamsSecurityTypeCommonStock    ActiveV1AccountOrderGetOrdersParamsSecurityType = "COMMON_STOCK"
	ActiveV1AccountOrderGetOrdersParamsSecurityTypePreferredStock ActiveV1AccountOrderGetOrdersParamsSecurityType = "PREFERRED_STOCK"
	ActiveV1AccountOrderGetOrdersParamsSecurityTypeCorporateBond  ActiveV1AccountOrderGetOrdersParamsSecurityType = "CORPORATE_BOND"
	ActiveV1AccountOrderGetOrdersParamsSecurityTypeOption         ActiveV1AccountOrderGetOrdersParamsSecurityType = "OPTION"
	ActiveV1AccountOrderGetOrdersParamsSecurityTypeFuture         ActiveV1AccountOrderGetOrdersParamsSecurityType = "FUTURE"
	ActiveV1AccountOrderGetOrdersParamsSecurityTypeWarrant        ActiveV1AccountOrderGetOrdersParamsSecurityType = "WARRANT"
	ActiveV1AccountOrderGetOrdersParamsSecurityTypeCash           ActiveV1AccountOrderGetOrdersParamsSecurityType = "CASH"
	ActiveV1AccountOrderGetOrdersParamsSecurityTypeOther          ActiveV1AccountOrderGetOrdersParamsSecurityType = "OTHER"
)

type ActiveV1AccountOrderReplaceOrderParams struct {
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

func (r ActiveV1AccountOrderReplaceOrderParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1AccountOrderReplaceOrderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1AccountOrderReplaceOrderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountOrderSubmitOrdersParams struct {
	Body []ActiveV1AccountOrderSubmitOrdersParamsBodyUnion
	paramObj
}

func (r ActiveV1AccountOrderSubmitOrdersParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ActiveV1AccountOrderSubmitOrdersParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ActiveV1AccountOrderSubmitOrdersParamsBodyUnion struct {
	OfActiveV1AccountOrderSubmitOrderssBodyNewOrderMultilegRequest *ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequest `json:",omitzero,inline"`
	OfActiveV1AccountOrderSubmitOrderssBodyNewOrderRequest         *ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderRequest         `json:",omitzero,inline"`
	paramUnion
}

func (u ActiveV1AccountOrderSubmitOrdersParamsBodyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfActiveV1AccountOrderSubmitOrderssBodyNewOrderMultilegRequest, u.OfActiveV1AccountOrderSubmitOrderssBodyNewOrderRequest)
}
func (u *ActiveV1AccountOrderSubmitOrdersParamsBodyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Multileg strategy order request
//
// The properties Legs, OrderType, TimeInForce are required.
type ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequest struct {
	// Legs that compose the strategy.
	Legs []ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLeg `json:"legs,omitzero" api:"required"`
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

func (r ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequest) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single leg in a multileg strategy request.
//
// The properties Ratio, Security, SecurityType, Side are required.
type ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLeg struct {
	// Ratio for the leg.
	Ratio string `json:"ratio" api:"required"`
	// Security identifier for the leg.
	Security ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLegSecurityUnion `json:"security,omitzero" api:"required"`
	// Security type for the leg.
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	SecurityType SecurityType `json:"security_type,omitzero" api:"required"`
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

func (r ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLeg) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLeg
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLeg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLeg](
		"position_effect", "OPEN", "CLOSE",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLegSecurityUnion struct {
	OfString                                                                                param.Opt[string]                                                                           `json:",omitzero,inline"`
	OfActiveV1AccountOrderSubmitOrderssBodyNewOrderMultilegRequestLegSecuritySecurityIDPair *ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLegSecuritySecurityIDPair `json:",omitzero,inline"`
	paramUnion
}

func (u ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLegSecurityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfActiveV1AccountOrderSubmitOrderssBodyNewOrderMultilegRequestLegSecuritySecurityIDPair)
}
func (u *ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLegSecurityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ID, Source are required.
type ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLegSecuritySecurityIDPair struct {
	ID string `json:"id" api:"required"`
	// Security identifier source
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "CURRENCY", "FMP", "OEMS",
	// "SEDOL", "QUIK", "ISIN", "RIC", "COUNTRY", "EXCHANGE", "CTA", "BLOOMBERG",
	// "WERTPAPIER", "DUTCH", "VALOREN", "SICOVAM", "BELGIAN", "COMMON",
	// "CLEARING_HOUSE", "ISDA_FPML_SPECIFICATION", "ISDA_FPML_URL",
	// "LETTER_OF_CREDIT", "MARKETPLACE_ASSIGNED_IDENTIFIER", "MARKIT_RED_ENTITY_CLIP",
	// "MARKIT_RED_PAIR_CLIP", "CFTC", "ISDA_COMMODITY_REFERENCE_PRICE",
	// "LEGAL_ENTITY_IDENTIFIER", "SYNTHETIC", "FIDESSA_INSTRUMENT_MNEMONIC",
	// "INDEX_NAME", "UNIFORM_SYMBOL", "DIGITAL_TOKEN_IDENTIFIER", "MASSIVE", "OTHER".
	Source SecurityIDSource `json:"source,omitzero" api:"required"`
	paramObj
}

func (r ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLegSecuritySecurityIDPair) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLegSecuritySecurityIDPair
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLegSecuritySecurityIDPair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single-leg order request
//
// The properties OrderType, Quantity, SecurityType, Side, TimeInForce are
// required.
type ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderRequest struct {
	// Type of order
	//
	// Any of "MARKET", "LIMIT", "STOP", "STOP_LIMIT", "TRAILING_STOP",
	// "TRAILING_STOP_LIMIT", "OTHER".
	OrderType OrderType `json:"order_type,omitzero" api:"required"`
	// Quantity to trade. For COMMON_STOCK: shares (may be fractional if supported).
	// For OPTION (single-leg): contracts (must be an integer)
	Quantity string `json:"quantity" api:"required"`
	// Type of security
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	SecurityType SecurityType `json:"security_type,omitzero" api:"required"`
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
	ExpireAt param.Opt[time.Time] `json:"expire_at,omitzero" format:"date-time"`
	// Allow trading outside regular trading hours. Some brokers disallow options
	// outside RTH.
	ExtendedHours param.Opt[bool] `json:"extended_hours,omitzero"`
	// Limit offset for trailing stop-limit orders (signed)
	LimitOffset param.Opt[string] `json:"limit_offset,omitzero"`
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
	// Trailing offset amount (required for trailing orders)
	TrailingOffsetAmt param.Opt[string] `json:"trailing_offset_amt,omitzero"`
	// Required when security_type is OPTION. Specifies whether the order opens or
	// closes a position.
	//
	// Any of "OPEN", "CLOSE".
	PositionEffect string `json:"position_effect,omitzero"`
	// The source of the security identifier. Required if security_id is provided.
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "CURRENCY", "FMP", "OEMS",
	// "SEDOL", "QUIK", "ISIN", "RIC", "COUNTRY", "EXCHANGE", "CTA", "BLOOMBERG",
	// "WERTPAPIER", "DUTCH", "VALOREN", "SICOVAM", "BELGIAN", "COMMON",
	// "CLEARING_HOUSE", "ISDA_FPML_SPECIFICATION", "ISDA_FPML_URL",
	// "LETTER_OF_CREDIT", "MARKETPLACE_ASSIGNED_IDENTIFIER", "MARKIT_RED_ENTITY_CLIP",
	// "MARKIT_RED_PAIR_CLIP", "CFTC", "ISDA_COMMODITY_REFERENCE_PRICE",
	// "LEGAL_ENTITY_IDENTIFIER", "SYNTHETIC", "FIDESSA_INSTRUMENT_MNEMONIC",
	// "INDEX_NAME", "UNIFORM_SYMBOL", "DIGITAL_TOKEN_IDENTIFIER", "MASSIVE", "OTHER".
	SecurityIDSource SecurityIDSource `json:"security_id_source,omitzero"`
	// Execution strategy/router. Defaults to SOR where applicable.
	Strategy OrderStrategyUnionParam `json:"strategy,omitzero"`
	// Trailing offset type (PRICE or PERCENT_BPS)
	//
	// Any of "PRICE", "PERCENT_BPS".
	TrailingOffsetAmtType TrailingOffsetType `json:"trailing_offset_amt_type,omitzero"`
	paramObj
}

func (r ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderRequest) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderRequest](
		"position_effect", "OPEN", "CLOSE",
	)
}
