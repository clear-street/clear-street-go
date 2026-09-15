// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
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
// V1OrderService contains methods and other services that help with interacting
// with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1OrderService] method instead.
type V1OrderService struct {
	options []option.RequestOption
}

// NewV1OrderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1OrderService(opts ...option.RequestOption) (r V1OrderService) {
	r = V1OrderService{}
	r.options = opts
	return
}

// Cancel all orders for an account
func (r *V1OrderService) CancelAllOpenOrders(ctx context.Context, accountID int64, body V1OrderCancelAllOpenOrdersParams, opts ...option.RequestOption) (res *V1OrderCancelAllOpenOrdersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Cancel a specific order
func (r *V1OrderService) CancelOpenOrder(ctx context.Context, orderID string, body V1OrderCancelOpenOrderParams, opts ...option.RequestOption) (res *V1OrderCancelOpenOrderResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%v/orders/%s", body.AccountID, url.PathEscape(orderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Retrieves filled and partially-filled execution reports for the specified
// trading account, ordered by transaction time (nanosecond precision) descending.
func (r *V1OrderService) GetExecutions(ctx context.Context, accountID int64, query V1OrderGetExecutionsParams, opts ...option.RequestOption) (res *V1OrderGetExecutionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/executions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Fetch a single order. The `{order_id}` path parameter accepts either the order's
// `id` or its `client_order_id`. A `client_order_id` can only be used while the
// order is open; after that, use the `id` returned in every order response, or
// find the order with the list-orders endpoint's `order_ids` filter, which accepts
// both identifiers at any time.
func (r *V1OrderService) GetOrderByID(ctx context.Context, orderID string, query V1OrderGetOrderByIDParams, opts ...option.RequestOption) (res *V1OrderGetOrderByIDResponse, err error) {
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
func (r *V1OrderService) GetOrders(ctx context.Context, accountID int64, query V1OrderGetOrdersParams, opts ...option.RequestOption) (res *V1OrderGetOrdersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Replace an order with new parameters
func (r *V1OrderService) ReplaceOrder(ctx context.Context, orderID string, params V1OrderReplaceOrderParams, opts ...option.RequestOption) (res *V1OrderReplaceOrderResponse, err error) {
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
func (r *V1OrderService) SubmitOrders(ctx context.Context, accountID int64, body V1OrderSubmitOrdersParams, opts ...option.RequestOption) (res *V1OrderSubmitOrdersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Represents a single fill of an order for an account.
type Execution struct {
	// Unique identifier for this execution report.
	ID string `json:"id" api:"required" format:"uuid"`
	// Identifier of the order this execution belongs to.
	OrderID string `json:"order_id" api:"required" format:"uuid"`
	// Filled quantity.
	Quantity string `json:"quantity" api:"required"`
	// Side of the fill.
	//
	// Any of "BUY", "SELL".
	Side Side `json:"side" api:"required"`
	// Transaction timestamp in nanosecond precision (UTC).
	TransactionTime time.Time `json:"transaction_time" api:"required" format:"date-time"`
	// Unique instrument identifier. `null` when this fill has no single resolvable
	// instrument. When a null/undefined value is observed, it indicates it does not
	// apply.
	InstrumentID string `json:"instrument_id" api:"nullable" format:"uuid"`
	// Fill price. `null` for multileg fills, whose price lives only at the leg level.
	// When a null/undefined value is observed, it indicates it does not apply.
	Price string `json:"price" api:"nullable"`
	// Trading symbol. `null` when this fill has no single resolvable instrument. When
	// a null/undefined value is observed, it indicates it does not apply.
	Symbol string `json:"symbol" api:"nullable"`
	// Underlying instrument identifier for an option fill. Omitted for a
	// non-derivative fill, when the underlier could not be resolved, or for a multileg
	// fill (per-leg underliers live in `legs[]`). When a null/undefined value is
	// observed, it indicates it does not apply.
	UnderlyingInstrumentID string `json:"underlying_instrument_id" api:"nullable" format:"uuid"`
	// Type of the underlying instrument, alongside `underlying_instrument_id`. When a
	// null/undefined value is observed, it indicates it does not apply.
	//
	// Any of "COMMON_STOCK", "INDEX", "OPTION", "CASH".
	UnderlyingInstrumentType SecurityType `json:"underlying_instrument_type" api:"nullable"`
	// Venue where this fill occurred, as reported by that venue. Distinct from an
	// order's `venue`, which is the routing destination. Codes are not normalized, so
	// the format varies by venue. When a null/undefined value is observed, it
	// indicates that there is no available data.
	Venue string `json:"venue" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		OrderID                  respjson.Field
		Quantity                 respjson.Field
		Side                     respjson.Field
		TransactionTime          respjson.Field
		InstrumentID             respjson.Field
		Price                    respjson.Field
		Symbol                   respjson.Field
		UnderlyingInstrumentID   respjson.Field
		UnderlyingInstrumentType respjson.Field
		Venue                    respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Execution) RawJSON() string { return r.JSON.raw }
func (r *Execution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecutionList []Execution

type InstrumentIDOrSymbol = string

// Request to submit a new order (PlaceOrderRequest from spec)
//
// The properties OrderType, Quantity, Side, TimeInForce are required.
type NewOrderRequestParam struct {
	// Type of order
	//
	// Any of "MARKET", "LIMIT", "STOP", "STOP_LIMIT", "TRAILING_STOP",
	// "TRAILING_STOP_LIMIT".
	OrderType RequestOrderType `json:"order_type,omitzero" api:"required"`
	// Quantity to trade. For COMMON_STOCK: shares (may be fractional if supported).
	// For OPTION (single-leg): contracts (must be an integer)
	Quantity string `json:"quantity" api:"required"`
	// Side of the order
	//
	// Any of "BUY", "SELL".
	Side Side `json:"side,omitzero" api:"required"`
	// Time in force
	//
	// Any of "DAY", "GOOD_TILL_CANCEL", "IMMEDIATE_OR_CANCEL", "FILL_OR_KILL",
	// "GOOD_TILL_DATE", "AT_OPEN", "AT_CLOSE".
	TimeInForce RequestTimeInForce `json:"time_in_force,omitzero" api:"required"`
	// Optional client-provided unique ID (idempotency). Required to be unique per
	// account.
	ID param.Opt[string] `json:"id,omitzero"`
	// The timestamp when the order should expire (UTC). Required when time_in_force is
	// GOOD_TILL_DATE.
	ExpiresAt param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	// Allow trading outside regular trading hours. Some brokers disallow options
	// outside RTH.
	ExtendedHours param.Opt[bool] `json:"extended_hours,omitzero"`
	// Limit offset for trailing stop-limit orders (signed)
	LimitOffset param.Opt[string] `json:"limit_offset,omitzero"`
	// Limit price (required for LIMIT and STOP_LIMIT orders)
	LimitPrice param.Opt[string] `json:"limit_price,omitzero"`
	// Stop price (required for STOP and STOP_LIMIT orders)
	StopPrice param.Opt[string] `json:"stop_price,omitzero"`
	// Trading symbol. For equities, use the ticker symbol (e.g., "TSLA"). For options,
	// use the OSI symbol (e.g., "TSLA 250117C00190000"). Either `symbol` or
	// `instrument_id` must be provided.
	Symbol param.Opt[string] `json:"symbol,omitzero"`
	// Trailing offset amount (required for trailing orders)
	TrailingOffset param.Opt[string] `json:"trailing_offset,omitzero"`
	// Instrument ID (UUID) or symbol (equity ticker or OSI option symbol). Either
	// `symbol` or `instrument_id` must be provided.
	InstrumentID param.Opt[InstrumentIDOrSymbol] `json:"instrument_id,omitzero"`
	// Optional open/close intent for this order. When omitted, the platform determines
	// the position effect.
	//
	// Any of "OPEN", "CLOSE".
	PositionIntent RequestPositionEffect `json:"position_intent,omitzero"`
	// Trailing offset type (PRICE or PERCENT_BPS)
	//
	// Any of "PRICE", "BPS".
	TrailingOffsetType TrailingOffsetType `json:"trailing_offset_type,omitzero"`
	paramObj
}

func (r NewOrderRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow NewOrderRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewOrderRequestParam) UnmarshalJSON(data []byte) error {
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
	// Client-provided identifier echoed back.
	ClientOrderID string `json:"client_order_id" api:"required"`
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
	// Side of the order (BUY or SELL)
	//
	// Any of "BUY", "SELL".
	Side Side `json:"side" api:"required"`
	// Current status of the order
	//
	// Any of "PENDING_NEW", "QUEUED", "PENDING_TRIGGER", "NEW", "PARTIALLY_FILLED",
	// "FILLED", "CANCELED", "REJECTED", "EXPIRED", "PENDING_CANCEL",
	// "PENDING_REPLACE", "REPLACED", "DONE_FOR_DAY", "STOPPED", "SUSPENDED",
	// "CALCULATED", "OTHER".
	Status OrderStatus `json:"status" api:"required"`
	// Time in force instruction
	//
	// Any of "DAY", "GOOD_TILL_CANCEL", "IMMEDIATE_OR_CANCEL", "FILL_OR_KILL",
	// "GOOD_TILL_DATE", "AT_OPEN", "AT_CLOSE", "OTHER".
	TimeInForce TimeInForce `json:"time_in_force" api:"required"`
	// Timestamp of the most recent update (UTC)
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// MIC code of the venue where the order is routed
	Venue string `json:"venue" api:"required"`
	// Average fill price across all executions. For multileg orders this is the
	// venue's strategy-level average when reported, else the signed net package price
	// derived from the leg averages: net debit positive, net credit negative, zero
	// possible. When a null/undefined value is observed, it indicates that there is no
	// available data.
	AverageFillPrice string `json:"average_fill_price" api:"nullable"`
	// Contains execution, rejection or cancellation details, if any
	Details []string `json:"details"`
	// Timestamp when the order will expire (UTC). Present when time_in_force is
	// GOOD_TILL_DATE. When a null/undefined value is observed, it indicates it does
	// not apply.
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// Whether the order is eligible for extended-hours trading.
	ExtendedHours bool `json:"extended_hours" api:"nullable"`
	// Instrument identifier for the traded instrument. `null` when the order has no
	// single resolvable instrument. When a null/undefined value is observed, it
	// indicates it does not apply.
	InstrumentID string `json:"instrument_id" api:"nullable" format:"uuid"`
	// Type of security. `null` when the order has no single resolvable instrument.
	// When a null/undefined value is observed, it indicates it does not apply.
	//
	// Any of "COMMON_STOCK", "INDEX", "OPTION", "CASH".
	InstrumentType SecurityType `json:"instrument_type" api:"nullable"`
	// Limit offset for trailing stop-limit orders (signed) When a null/undefined value
	// is observed, it indicates it does not apply.
	LimitOffset string `json:"limit_offset" api:"nullable"`
	// Limit price (for LIMIT and STOP_LIMIT orders) When a null/undefined value is
	// observed, it indicates it does not apply.
	LimitPrice string `json:"limit_price" api:"nullable"`
	// Parent order queue state, present when the order is awaiting release or
	// released. When a null/undefined value is observed, it indicates it does not
	// apply.
	//
	// Any of "AWAITING_RELEASE", "RELEASED".
	QueueState QueueState `json:"queue_state" api:"nullable"`
	// Scheduled release time for orders awaiting release. When a null/undefined value
	// is observed, it indicates it does not apply.
	ReleasesAt time.Time `json:"releases_at" api:"nullable" format:"date-time"`
	// Stop price (for STOP and STOP_LIMIT orders) When a null/undefined value is
	// observed, it indicates it does not apply.
	StopPrice string `json:"stop_price" api:"nullable"`
	// Trading symbol. `null` when the order has no single resolvable instrument. When
	// a null/undefined value is observed, it indicates it does not apply.
	Symbol string `json:"symbol" api:"nullable"`
	// Current trailing limit price computed by the trailing strategy When a
	// null/undefined value is observed, it indicates it does not apply.
	TrailingLimitPx string `json:"trailing_limit_px" api:"nullable"`
	// Trailing offset amount for trailing orders When a null/undefined value is
	// observed, it indicates it does not apply.
	TrailingOffset string `json:"trailing_offset" api:"nullable"`
	// Trailing offset type for trailing orders When a null/undefined value is
	// observed, it indicates it does not apply.
	//
	// Any of "PRICE", "BPS".
	TrailingOffsetType TrailingOffsetType `json:"trailing_offset_type" api:"nullable"`
	// Current trailing stop price computed by the trailing strategy When a
	// null/undefined value is observed, it indicates it does not apply.
	TrailingStopPx string `json:"trailing_stop_px" api:"nullable"`
	// Trailing watermark price for trailing orders When a null/undefined value is
	// observed, it indicates it does not apply.
	TrailingWatermarkPx string `json:"trailing_watermark_px" api:"nullable"`
	// Trailing watermark timestamp for trailing orders When a null/undefined value is
	// observed, it indicates it does not apply.
	TrailingWatermarkTs time.Time `json:"trailing_watermark_ts" api:"nullable" format:"date-time"`
	// Instrument ID of the option's underlying instrument. Populated only for options
	// orders. A `null` means one of two things: the order is not an option, so the
	// field does not apply; or the order is an option whose underlier has not yet been
	// resolved. When a null/undefined value is observed, it indicates it does not
	// apply.
	UnderlyingInstrumentID string `json:"underlying_instrument_id" api:"nullable" format:"uuid"`
	// Type of the underlying instrument, alongside `underlying_instrument_id`. When a
	// null/undefined value is observed, it indicates it does not apply.
	//
	// Any of "COMMON_STOCK", "INDEX", "OPTION", "CASH".
	UnderlyingInstrumentType SecurityType `json:"underlying_instrument_type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		AccountID                respjson.Field
		ClientOrderID            respjson.Field
		CreatedAt                respjson.Field
		FilledQuantity           respjson.Field
		LeavesQuantity           respjson.Field
		OrderType                respjson.Field
		Quantity                 respjson.Field
		Side                     respjson.Field
		Status                   respjson.Field
		TimeInForce              respjson.Field
		UpdatedAt                respjson.Field
		Venue                    respjson.Field
		AverageFillPrice         respjson.Field
		Details                  respjson.Field
		ExpiresAt                respjson.Field
		ExtendedHours            respjson.Field
		InstrumentID             respjson.Field
		InstrumentType           respjson.Field
		LimitOffset              respjson.Field
		LimitPrice               respjson.Field
		QueueState               respjson.Field
		ReleasesAt               respjson.Field
		StopPrice                respjson.Field
		Symbol                   respjson.Field
		TrailingLimitPx          respjson.Field
		TrailingOffset           respjson.Field
		TrailingOffsetType       respjson.Field
		TrailingStopPx           respjson.Field
		TrailingWatermarkPx      respjson.Field
		TrailingWatermarkTs      respjson.Field
		UnderlyingInstrumentID   respjson.Field
		UnderlyingInstrumentType respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
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
	OrderStatusQueued          OrderStatus = "QUEUED"
	OrderStatusPendingTrigger  OrderStatus = "PENDING_TRIGGER"
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

// Parent order queue or hold state.
type QueueState string

const (
	QueueStateAwaitingRelease QueueState = "AWAITING_RELEASE"
	QueueStateReleased        QueueState = "RELEASED"
)

// Request to replace (modify) an existing order
//
// At least one field must be provided.
type ReplaceOrderRequestParam struct {
	// New limit offset for trailing stop-limit orders (signed)
	LimitOffset param.Opt[string] `json:"limit_offset,omitzero"`
	// New limit price for the order
	LimitPrice param.Opt[string] `json:"limit_price,omitzero"`
	// New quantity for the order
	Quantity param.Opt[string] `json:"quantity,omitzero"`
	// New stop price for the order
	StopPrice param.Opt[string] `json:"stop_price,omitzero"`
	// New trailing offset for trailing orders
	TrailingOffset param.Opt[string] `json:"trailing_offset,omitzero"`
	// New trailing offset type (PRICE or BPS)
	//
	// Any of "PRICE", "BPS".
	TrailingOffsetType TrailingOffsetType `json:"trailing_offset_type,omitzero"`
	paramObj
}

func (r ReplaceOrderRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ReplaceOrderRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReplaceOrderRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Strict order-type enum for order submission/replacement requests.
type RequestOrderType string

const (
	RequestOrderTypeMarket            RequestOrderType = "MARKET"
	RequestOrderTypeLimit             RequestOrderType = "LIMIT"
	RequestOrderTypeStop              RequestOrderType = "STOP"
	RequestOrderTypeStopLimit         RequestOrderType = "STOP_LIMIT"
	RequestOrderTypeTrailingStop      RequestOrderType = "TRAILING_STOP"
	RequestOrderTypeTrailingStopLimit RequestOrderType = "TRAILING_STOP_LIMIT"
)

// Position effect for a multileg strategy leg: client-attested open/close intent.
// Required on every leg of a multileg order submission.
type RequestPositionEffect string

const (
	RequestPositionEffectOpen  RequestPositionEffect = "OPEN"
	RequestPositionEffectClose RequestPositionEffect = "CLOSE"
)

// Strict time-in-force enum for order submission requests.
type RequestTimeInForce string

const (
	RequestTimeInForceDay               RequestTimeInForce = "DAY"
	RequestTimeInForceGoodTillCancel    RequestTimeInForce = "GOOD_TILL_CANCEL"
	RequestTimeInForceImmediateOrCancel RequestTimeInForce = "IMMEDIATE_OR_CANCEL"
	RequestTimeInForceFillOrKill        RequestTimeInForce = "FILL_OR_KILL"
	RequestTimeInForceGoodTillDate      RequestTimeInForce = "GOOD_TILL_DATE"
	RequestTimeInForceAtOpen            RequestTimeInForce = "AT_OPEN"
	RequestTimeInForceAtClose           RequestTimeInForce = "AT_CLOSE"
)

// Side of the order (BUY or SELL).
type Side string

const (
	SideBuy  Side = "BUY"
	SideSell Side = "SELL"
)

// Time in force
type TimeInForce string

const (
	TimeInForceDay               TimeInForce = "DAY"
	TimeInForceGoodTillCancel    TimeInForce = "GOOD_TILL_CANCEL"
	TimeInForceImmediateOrCancel TimeInForce = "IMMEDIATE_OR_CANCEL"
	TimeInForceFillOrKill        TimeInForce = "FILL_OR_KILL"
	TimeInForceGoodTillDate      TimeInForce = "GOOD_TILL_DATE"
	TimeInForceAtOpen            TimeInForce = "AT_OPEN"
	TimeInForceAtClose           TimeInForce = "AT_CLOSE"
	TimeInForceOther             TimeInForce = "OTHER"
)

// Trailing offset type for trailing stop orders.
type TrailingOffsetType string

const (
	TrailingOffsetTypePrice TrailingOffsetType = "PRICE"
	TrailingOffsetTypeBps   TrailingOffsetType = "BPS"
)

type V1OrderCancelAllOpenOrdersResponse struct {
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
func (r V1OrderCancelAllOpenOrdersResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OrderCancelAllOpenOrdersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OrderCancelOpenOrderResponse struct {
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
func (r V1OrderCancelOpenOrderResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OrderCancelOpenOrderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OrderGetExecutionsResponse struct {
	Data ExecutionList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1OrderGetExecutionsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OrderGetExecutionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OrderGetOrderByIDResponse struct {
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
func (r V1OrderGetOrderByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OrderGetOrderByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OrderGetOrdersResponse struct {
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
func (r V1OrderGetOrdersResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OrderGetOrdersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OrderReplaceOrderResponse struct {
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
func (r V1OrderReplaceOrderResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OrderReplaceOrderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OrderSubmitOrdersResponse struct {
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
func (r V1OrderSubmitOrdersResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OrderSubmitOrdersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OrderCancelAllOpenOrdersParams struct {
	// Comma-separated instrument IDs (UUID) or symbols (equity tickers or OSI option
	// symbols).
	InstrumentIDs []InstrumentIDOrSymbol `query:"instrument_ids,omitzero" json:"-"`
	// Filter by instrument type (e.g., COMMON_STOCK, OPTION)
	//
	// Any of "COMMON_STOCK", "INDEX", "OPTION", "CASH".
	InstrumentType V1OrderCancelAllOpenOrdersParamsInstrumentType `query:"instrument_type,omitzero" json:"-"`
	// Filter by order side (BUY or SELL)
	//
	// Any of "BUY", "SELL".
	Side V1OrderCancelAllOpenOrdersParamsSide `query:"side,omitzero" json:"-"`
	// Filter by order type (e.g., MARKET, LIMIT)
	//
	// Any of "MARKET", "LIMIT", "STOP", "STOP_LIMIT", "TRAILING_STOP",
	// "TRAILING_STOP_LIMIT", "OTHER".
	Type V1OrderCancelAllOpenOrdersParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1OrderCancelAllOpenOrdersParams]'s query parameters as
// `url.Values`.
func (r V1OrderCancelAllOpenOrdersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by instrument type (e.g., COMMON_STOCK, OPTION)
type V1OrderCancelAllOpenOrdersParamsInstrumentType string

const (
	V1OrderCancelAllOpenOrdersParamsInstrumentTypeCommonStock V1OrderCancelAllOpenOrdersParamsInstrumentType = "COMMON_STOCK"
	V1OrderCancelAllOpenOrdersParamsInstrumentTypeIndex       V1OrderCancelAllOpenOrdersParamsInstrumentType = "INDEX"
	V1OrderCancelAllOpenOrdersParamsInstrumentTypeOption      V1OrderCancelAllOpenOrdersParamsInstrumentType = "OPTION"
	V1OrderCancelAllOpenOrdersParamsInstrumentTypeCash        V1OrderCancelAllOpenOrdersParamsInstrumentType = "CASH"
)

// Filter by order side (BUY or SELL)
type V1OrderCancelAllOpenOrdersParamsSide string

const (
	V1OrderCancelAllOpenOrdersParamsSideBuy  V1OrderCancelAllOpenOrdersParamsSide = "BUY"
	V1OrderCancelAllOpenOrdersParamsSideSell V1OrderCancelAllOpenOrdersParamsSide = "SELL"
)

// Filter by order type (e.g., MARKET, LIMIT)
type V1OrderCancelAllOpenOrdersParamsType string

const (
	V1OrderCancelAllOpenOrdersParamsTypeMarket            V1OrderCancelAllOpenOrdersParamsType = "MARKET"
	V1OrderCancelAllOpenOrdersParamsTypeLimit             V1OrderCancelAllOpenOrdersParamsType = "LIMIT"
	V1OrderCancelAllOpenOrdersParamsTypeStop              V1OrderCancelAllOpenOrdersParamsType = "STOP"
	V1OrderCancelAllOpenOrdersParamsTypeStopLimit         V1OrderCancelAllOpenOrdersParamsType = "STOP_LIMIT"
	V1OrderCancelAllOpenOrdersParamsTypeTrailingStop      V1OrderCancelAllOpenOrdersParamsType = "TRAILING_STOP"
	V1OrderCancelAllOpenOrdersParamsTypeTrailingStopLimit V1OrderCancelAllOpenOrdersParamsType = "TRAILING_STOP_LIMIT"
	V1OrderCancelAllOpenOrdersParamsTypeOther             V1OrderCancelAllOpenOrdersParamsType = "OTHER"
)

type V1OrderCancelOpenOrderParams struct {
	AccountID int64 `path:"account_id" api:"required" json:"-"`
	paramObj
}

type V1OrderGetExecutionsParams struct {
	// The start date and time for the query range, inclusive (ISO 8601 format)
	From param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// The end date and time for the query range, inclusive (ISO 8601 format)
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Comma-separated instrument identifiers (UUIDs) or symbols (e.g. `AAPL`) to
	// filter by. When provided, only executions for any of the listed instruments are
	// returned.
	InstrumentIDs []string `query:"instrument_ids,omitzero" json:"-"`
	// Comma-separated order IDs to filter by. When provided, only executions belonging
	// to an order in this set are returned.
	OrderIDs []string `query:"order_ids,omitzero" json:"-"`
	// Comma-separated instrument IDs (UUID) or symbols (equity tickers or OSI option
	// symbols). Matches option fills whose resolved underlier is any of the given
	// instruments.
	UnderlyingInstrumentIDs []string `query:"underlying_instrument_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1OrderGetExecutionsParams]'s query parameters as
// `url.Values`.
func (r V1OrderGetExecutionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1OrderGetOrderByIDParams struct {
	AccountID int64 `path:"account_id" api:"required" json:"-"`
	paramObj
}

type V1OrderGetOrdersParams struct {
	// The start date and time for the query range, inclusive (ISO 8601 format)
	From param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Filter by symbol
	Symbol param.Opt[string] `query:"symbol,omitzero" json:"-"`
	// The end date and time for the query range, inclusive (ISO 8601 format)
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Comma-separated instrument IDs (UUID) or symbols (equity tickers or OSI option
	// symbols).
	InstrumentIDs []InstrumentIDOrSymbol `query:"instrument_ids,omitzero" json:"-"`
	// Instrument type filter (e.g., COMMON_STOCK, OPTION)
	//
	// Any of "COMMON_STOCK", "INDEX", "OPTION", "CASH".
	InstrumentType V1OrderGetOrdersParamsInstrumentType `query:"instrument_type,omitzero" json:"-"`
	// Comma-separated list of order identifiers. Each value may be an order's `id` or
	// its `client_order_id`; only orders matching one of the given identifiers are
	// returned.
	OrderIDs []string `query:"order_ids,omitzero" json:"-"`
	// Comma-separated order statuses to filter by
	//
	// Any of "PENDING_NEW", "QUEUED", "PENDING_TRIGGER", "NEW", "PARTIALLY_FILLED",
	// "FILLED", "CANCELED", "REJECTED", "EXPIRED", "PENDING_CANCEL",
	// "PENDING_REPLACE", "REPLACED", "DONE_FOR_DAY", "STOPPED", "SUSPENDED",
	// "CALCULATED", "OTHER".
	Status []string `query:"status,omitzero" json:"-"`
	// Comma-separated instrument IDs (UUID) or symbols (equity tickers or OSI option
	// symbols). Matches options orders whose resolved underlier is any of the given
	// instruments.
	UnderlyingInstrumentIDs []InstrumentIDOrSymbol          `query:"underlying_instrument_ids,omitzero" json:"-"`
	UpdatedAt               V1OrderGetOrdersParamsUpdatedAt `query:"updated_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1OrderGetOrdersParams]'s query parameters as `url.Values`.
func (r V1OrderGetOrdersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Instrument type filter (e.g., COMMON_STOCK, OPTION)
type V1OrderGetOrdersParamsInstrumentType string

const (
	V1OrderGetOrdersParamsInstrumentTypeCommonStock V1OrderGetOrdersParamsInstrumentType = "COMMON_STOCK"
	V1OrderGetOrdersParamsInstrumentTypeIndex       V1OrderGetOrdersParamsInstrumentType = "INDEX"
	V1OrderGetOrdersParamsInstrumentTypeOption      V1OrderGetOrdersParamsInstrumentType = "OPTION"
	V1OrderGetOrdersParamsInstrumentTypeCash        V1OrderGetOrdersParamsInstrumentType = "CASH"
)

type V1OrderGetOrdersParamsUpdatedAt struct {
	// > **Alpha** — this parameter is experimental and may change or be removed at any
	// > time.
	Gt param.Opt[time.Time] `query:"gt,omitzero" format:"date-time" json:"-"`
	// > **Alpha** — this parameter is experimental and may change or be removed at any
	// > time.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// > **Alpha** — this parameter is experimental and may change or be removed at any
	// > time.
	Lt param.Opt[time.Time] `query:"lt,omitzero" format:"date-time" json:"-"`
	// > **Alpha** — this parameter is experimental and may change or be removed at any
	// > time.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [V1OrderGetOrdersParamsUpdatedAt]'s query parameters as
// `url.Values`.
func (r V1OrderGetOrdersParamsUpdatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1OrderReplaceOrderParams struct {
	AccountID int64 `path:"account_id" api:"required" json:"-"`
	// Request to replace (modify) an existing order
	//
	// At least one field must be provided.
	ReplaceOrderRequest ReplaceOrderRequestParam
	paramObj
}

func (r V1OrderReplaceOrderParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ReplaceOrderRequest)
}
func (r *V1OrderReplaceOrderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OrderSubmitOrdersParams struct {
	Orders []NewOrderRequestParam
	paramObj
}

func (r V1OrderSubmitOrdersParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Orders)
}
func (r *V1OrderSubmitOrdersParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
