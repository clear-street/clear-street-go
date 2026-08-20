// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"encoding/json"
	"time"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/respjson"
)

// V1OmniAIService contains methods and other services that help with interacting
// with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1OmniAIService] method instead.
type V1OmniAIService struct {
	options []option.RequestOption
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Thread/message/response endpoints require an
	// explicit account_id. Entitlement endpoints are caller-scoped and use
	// account_ids.
	Entitlements V1OmniAIEntitlementService
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Thread/message/response endpoints require an
	// explicit account_id. Entitlement endpoints are caller-scoped and use
	// account_ids.
	Messages V1OmniAIMessageService
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Thread/message/response endpoints require an
	// explicit account_id. Entitlement endpoints are caller-scoped and use
	// account_ids.
	Responses V1OmniAIResponseService
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Thread/message/response endpoints require an
	// explicit account_id. Entitlement endpoints are caller-scoped and use
	// account_ids.
	Threads V1OmniAIThreadService
}

// NewV1OmniAIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1OmniAIService(opts ...option.RequestOption) (r V1OmniAIService) {
	r = V1OmniAIService{}
	r.options = opts
	r.Entitlements = NewV1OmniAIEntitlementService(opts...)
	r.Messages = NewV1OmniAIMessageService(opts...)
	r.Responses = NewV1OmniAIResponseService(opts...)
	r.Threads = NewV1OmniAIThreadService(opts...)
	return
}

// Button metadata shared by chart and suggested-actions payloads.
type ActionButton struct {
	// Stable button identifier within the content part.
	ButtonID string `json:"buttonId" api:"required"`
	// User-visible label.
	Label string `json:"label" api:"required"`
	// Interaction-tracking identity. Absent on messages created before tracking. When
	// a null/undefined value is observed, it indicates that there is no available
	// data.
	ItemID string `json:"itemId" api:"nullable" format:"uuid"`
	// Follow-up prompt to submit as the next user message. When a null/undefined value
	// is observed, it indicates it does not apply.
	Prompt PromptButtonAction `json:"prompt" api:"nullable"`
	// Structured action in the same message to execute on click. When a null/undefined
	// value is observed, it indicates it does not apply.
	StructuredAction StructuredActionButtonAction `json:"structuredAction" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonID         respjson.Field
		Label            respjson.Field
		ItemID           respjson.Field
		Prompt           respjson.Field
		StructuredAction respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionButton) RawJSON() string { return r.JSON.raw }
func (r *ActionButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed chart payload rendered inline in assistant content.
type ChartPayload struct {
	// Stable chart identifier scoped to the content part.
	ChartID string `json:"chartId" api:"required"`
	// Whether the current user clicked this chart.
	Clicked bool `json:"clicked" api:"required"`
	// Buttons associated with this chart.
	ActionButtons []ActionButton `json:"actionButtons"`
	// Explicit series-driven chart definition. When a null/undefined value is
	// observed, it indicates it does not apply.
	DataChart DataChart `json:"dataChart" api:"nullable"`
	// Interaction-tracking identity. Absent on messages created before tracking. When
	// a null/undefined value is observed, it indicates that there is no available
	// data.
	ItemID string `json:"itemId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChartID       respjson.Field
		Clicked       respjson.Field
		ActionButtons respjson.Field
		DataChart     respjson.Field
		ItemID        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChartPayload) RawJSON() string { return r.JSON.raw }
func (r *ChartPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single chart coordinate.
type ChartPoint struct {
	X string  `json:"x" api:"required"`
	Y float64 `json:"y" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChartPoint) RawJSON() string { return r.JSON.raw }
func (r *ChartPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Named data series within a chart.
type ChartSeries struct {
	Name   string       `json:"name" api:"required"`
	Points []ChartPoint `json:"points"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Points      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChartSeries) RawJSON() string { return r.JSON.raw }
func (r *ChartSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chart payload content part.
type ContentPartChartPayload struct {
	// Typed chart payload rendered inline in assistant content.
	Payload ChartPayload `json:"payload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Payload     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentPartChartPayload) RawJSON() string { return r.JSON.raw }
func (r *ContentPartChartPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Escape-hatch custom payload content part.
type ContentPartCustomPayload struct {
	Payload any `json:"payload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Payload     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentPartCustomPayload) RawJSON() string { return r.JSON.raw }
func (r *ContentPartCustomPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured action content part.
type ContentPartStructuredActionPayload struct {
	// Structured actions that Omni AI can return to clients.
	//
	// These actions provide machine-readable instructions for the client to execute,
	// such as prefilling an order ticket, opening a chart, or navigating to a route.
	Action   StructuredActionUnion `json:"action" api:"required"`
	ActionID string                `json:"action_id" api:"required" format:"uuid"`
	// Whether the current user clicked this action.
	Clicked bool `json:"clicked" api:"required"`
	// IDs of nested items clicked by the current user.
	ClickedItemIDs []string `json:"clicked_item_ids" format:"uuid"`
	// Interaction-tracking identity. Absent on messages created before tracking. When
	// a null/undefined value is observed, it indicates that there is no available
	// data.
	ItemID string `json:"item_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action         respjson.Field
		ActionID       respjson.Field
		Clicked        respjson.Field
		ClickedItemIDs respjson.Field
		ItemID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentPartStructuredActionPayload) RawJSON() string { return r.JSON.raw }
func (r *ContentPartStructuredActionPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Suggested actions payload content part.
type ContentPartSuggestedActionsPayload struct {
	// Suggested follow-up buttons rendered at the end of an assistant message.
	Payload SuggestedActionsPayload `json:"payload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Payload     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentPartSuggestedActionsPayload) RawJSON() string { return r.JSON.raw }
func (r *ContentPartSuggestedActionsPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part.
type ContentPartTextPayload struct {
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentPartTextPayload) RawJSON() string { return r.JSON.raw }
func (r *ContentPartTextPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Thinking content part shown on dynamic response polling.
type ContentPartThinkingPayload struct {
	Thoughts []string `json:"thoughts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Thoughts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentPartThinkingPayload) RawJSON() string { return r.JSON.raw }
func (r *ContentPartThinkingPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chart represented by explicit data series.
type DataChart struct {
	Series []ChartSeries `json:"series"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Series      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataChart) RawJSON() string { return r.JSON.raw }
func (r *DataChart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable entitlement agreement family key.
type EntitlementAgreementKey string

const (
	EntitlementAgreementKeyOmniAccountDataAccess EntitlementAgreementKey = "omni_account_data_access"
)

// Stable entitlement code granted by an agreement.
type EntitlementCode string

const (
	EntitlementCodeOmniAccountData EntitlementCode = "omni.account_data"
)

// Action to open a chart for a symbol.
type OpenChartAction struct {
	// Trading symbol to chart
	Symbol string `json:"symbol" api:"required"`
	// Additional chart configuration (indicators, overlays, etc.) When a
	// null/undefined value is observed, it indicates it does not apply.
	Extras any `json:"extras"`
	// Interaction-tracking identity. Absent on messages created before tracking. When
	// a null/undefined value is observed, it indicates that there is no available
	// data.
	ItemID string `json:"item_id" api:"nullable" format:"uuid"`
	// Chart timeframe (e.g., "1D", "1W", "1M", "3M", "1Y", "5Y") When a null/undefined
	// value is observed, it indicates it does not apply.
	Timeframe string `json:"timeframe" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Symbol      respjson.Field
		Extras      respjson.Field
		ItemID      respjson.Field
		Timeframe   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenChartAction) RawJSON() string { return r.JSON.raw }
func (r *OpenChartAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action to open entitlement consent flow for one or more accounts.
type OpenEntitlementConsentAction struct {
	AccountIDs []int64 `json:"account_ids" api:"required"`
	// Stable entitlement agreement family key.
	//
	// Any of "omni_account_data_access".
	AgreementKey     EntitlementAgreementKey `json:"agreement_key" api:"required"`
	EntitlementCodes []EntitlementCode       `json:"entitlement_codes" api:"required"`
	Reason           string                  `json:"reason" api:"required"`
	// Interaction-tracking identity. Absent on messages created before tracking. When
	// a null/undefined value is observed, it indicates that there is no available
	// data.
	ItemID string `json:"item_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountIDs       respjson.Field
		AgreementKey     respjson.Field
		EntitlementCodes respjson.Field
		Reason           respjson.Field
		ItemID           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenEntitlementConsentAction) RawJSON() string { return r.JSON.raw }
func (r *OpenEntitlementConsentAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action to open a stock screener with filters.
type OpenScreenerAction struct {
	// Filter criteria for the screener
	Filters []ScreenerFilter `json:"filters" api:"required"`
	// Optional field/column selection for screener results. When a null/undefined
	// value is observed, it indicates it does not apply.
	Columns []string `json:"columns" api:"nullable"`
	// Interaction-tracking identity. Absent on messages created before tracking. When
	// a null/undefined value is observed, it indicates that there is no available
	// data.
	ItemID string `json:"item_id" api:"nullable" format:"uuid"`
	// Optional page size. When a null/undefined value is observed, it indicates it
	// does not apply.
	PageSize int64 `json:"page_size" api:"nullable"`
	// Optional sort field for screener rows. When a null/undefined value is observed,
	// it indicates it does not apply.
	SortBy string `json:"sort_by" api:"nullable"`
	// Optional sort direction (`ASC` or `DESC`). When a null/undefined value is
	// observed, it indicates it does not apply.
	SortDirection string `json:"sort_direction" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters       respjson.Field
		Columns       respjson.Field
		ItemID        respjson.Field
		PageSize      respjson.Field
		SortBy        respjson.Field
		SortDirection respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenScreenerAction) RawJSON() string { return r.JSON.raw }
func (r *OpenScreenerAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cancel-order prefill action.
type PrefillCancelOrderAction struct {
	// Orders to cancel using the same identifiers required by the cancel-order API.
	Orders []PrefillCancelOrderRequest `json:"orders" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Orders      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrefillCancelOrderAction) RawJSON() string { return r.JSON.raw }
func (r *PrefillCancelOrderAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to cancel an existing order
//
// Note: In the API, order cancellation is done via DELETE request without a body.
// The order_id and account_id come from the URL path parameters.
type PrefillCancelOrderRequest struct {
	// Account ID (from path parameter)
	AccountID int64 `json:"account_id" api:"required"`
	// Order ID to cancel (from path parameter)
	OrderID string `json:"order_id" api:"required"`
	// Interaction-tracking identity. Absent on messages created before tracking. When
	// a null/undefined value is observed, it indicates that there is no available
	// data.
	ItemID string `json:"item_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID   respjson.Field
		OrderID     respjson.Field
		ItemID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrefillCancelOrderRequest) RawJSON() string { return r.JSON.raw }
func (r *PrefillCancelOrderRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Modify-order prefill action.
type PrefillModifyOrderAction struct {
	// Modification targets and deltas needed to construct replace-order API requests.
	Orders []PrefillModifyOrderRequest `json:"orders" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Orders      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrefillModifyOrderAction) RawJSON() string { return r.JSON.raw }
func (r *PrefillModifyOrderAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to replace (modify) an existing order
//
// At least one field must be provided.
type PrefillModifyOrderRequest struct {
	// Account ID that owns the order.
	AccountID int64 `json:"account_id"`
	// Interaction-tracking identity. Absent on messages created before tracking. When
	// a null/undefined value is observed, it indicates that there is no available
	// data.
	ItemID string `json:"item_id" api:"nullable" format:"uuid"`
	// New limit offset for trailing stop-limit orders (signed)
	LimitOffset string `json:"limit_offset" api:"nullable"`
	// New limit price for the order
	LimitPrice string `json:"limit_price" api:"nullable"`
	// Order ID to modify.
	OrderID string `json:"order_id"`
	// New quantity for the order
	Quantity string `json:"quantity" api:"nullable"`
	// New stop price for the order
	StopPrice string `json:"stop_price" api:"nullable"`
	// New trailing offset for trailing orders
	TrailingOffset string `json:"trailing_offset" api:"nullable"`
	// New trailing offset type (PRICE or BPS)
	//
	// Any of "PRICE", "BPS".
	TrailingOffsetType TrailingOffsetType `json:"trailing_offset_type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID          respjson.Field
		ItemID             respjson.Field
		LimitOffset        respjson.Field
		LimitPrice         respjson.Field
		OrderID            respjson.Field
		Quantity           respjson.Field
		StopPrice          respjson.Field
		TrailingOffset     respjson.Field
		TrailingOffsetType respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrefillModifyOrderRequest) RawJSON() string { return r.JSON.raw }
func (r *PrefillModifyOrderRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// New-order prefill action.
type PrefillNewOrderAction struct {
	// Orders to prefill using the same shape accepted by the orders API.
	Orders []PrefillNewOrderRequest `json:"orders" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Orders      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrefillNewOrderAction) RawJSON() string { return r.JSON.raw }
func (r *PrefillNewOrderAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to submit a new order (PlaceOrderRequest from spec)
type PrefillNewOrderRequest struct {
	// Type of order
	//
	// Any of "MARKET", "LIMIT", "STOP", "STOP_LIMIT", "TRAILING_STOP",
	// "TRAILING_STOP_LIMIT".
	OrderType RequestOrderType `json:"order_type" api:"required"`
	// Quantity to trade. For COMMON_STOCK: shares (may be fractional if supported).
	// For OPTION (single-leg): contracts (must be an integer)
	Quantity string `json:"quantity" api:"required"`
	// Side of the order
	//
	// Any of "BUY", "SELL".
	Side Side `json:"side" api:"required"`
	// Time in force
	//
	// Any of "DAY", "GOOD_TILL_CANCEL", "IMMEDIATE_OR_CANCEL", "FILL_OR_KILL",
	// "GOOD_TILL_DATE", "AT_THE_OPENING", "AT_THE_CLOSE".
	TimeInForce RequestTimeInForce `json:"time_in_force" api:"required"`
	// Optional client-provided unique ID (idempotency). Required to be unique per
	// account.
	ID string `json:"id" api:"nullable"`
	// The timestamp when the order should expire (UTC). Required when time_in_force is
	// GOOD_TILL_DATE.
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// Allow trading outside regular trading hours. Some brokers disallow options
	// outside RTH.
	ExtendedHours bool `json:"extended_hours" api:"nullable"`
	// Instrument ID (UUID) or symbol (equity ticker or OSI option symbol). Either
	// `symbol` or `instrument_id` must be provided.
	InstrumentID InstrumentIDOrSymbol `json:"instrument_id" api:"nullable"`
	// Interaction-tracking identity. Absent on messages created before tracking. When
	// a null/undefined value is observed, it indicates that there is no available
	// data.
	ItemID string `json:"item_id" api:"nullable" format:"uuid"`
	// Limit offset for trailing stop-limit orders (signed)
	LimitOffset string `json:"limit_offset" api:"nullable"`
	// Limit price (required for LIMIT and STOP_LIMIT orders)
	LimitPrice string `json:"limit_price" api:"nullable"`
	// Optional open/close intent for this order. When omitted, the platform determines
	// the position effect.
	//
	// Any of "OPEN", "CLOSE".
	PositionIntent RequestPositionEffect `json:"position_intent" api:"nullable"`
	// Stop price (required for STOP and STOP_LIMIT orders)
	StopPrice string `json:"stop_price" api:"nullable"`
	// Trading symbol. For equities, use the ticker symbol (e.g., "TSLA"). For options,
	// use the OSI symbol (e.g., "TSLA 250117C00190000"). Either `symbol` or
	// `instrument_id` must be provided.
	Symbol string `json:"symbol" api:"nullable"`
	// Trailing offset amount (required for trailing orders)
	TrailingOffset string `json:"trailing_offset" api:"nullable"`
	// Trailing offset type (PRICE or PERCENT_BPS)
	//
	// Any of "PRICE", "BPS".
	TrailingOffsetType TrailingOffsetType `json:"trailing_offset_type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrderType          respjson.Field
		Quantity           respjson.Field
		Side               respjson.Field
		TimeInForce        respjson.Field
		ID                 respjson.Field
		ExpiresAt          respjson.Field
		ExtendedHours      respjson.Field
		InstrumentID       respjson.Field
		ItemID             respjson.Field
		LimitOffset        respjson.Field
		LimitPrice         respjson.Field
		PositionIntent     respjson.Field
		StopPrice          respjson.Field
		Symbol             respjson.Field
		TrailingOffset     respjson.Field
		TrailingOffsetType respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrefillNewOrderRequest) RawJSON() string { return r.JSON.raw }
func (r *PrefillNewOrderRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PrefillOrderActionUnion contains all possible properties and values from
// [PrefillOrderActionPrefillNewOrderAction],
// [PrefillOrderActionPrefillCancelOrderAction],
// [PrefillOrderActionPrefillModifyOrderAction].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PrefillOrderActionUnion struct {
	// This field is a union of [[]PrefillNewOrderRequest],
	// [[]PrefillCancelOrderRequest], [[]PrefillModifyOrderRequest]
	Orders     PrefillOrderActionUnionOrders `json:"orders"`
	ActionType string                        `json:"action_type"`
	JSON       struct {
		Orders     respjson.Field
		ActionType respjson.Field
		raw        string
	} `json:"-"`
}

func (u PrefillOrderActionUnion) AsPrefillNewOrderAction() (v PrefillOrderActionPrefillNewOrderAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PrefillOrderActionUnion) AsPrefillCancelOrderAction() (v PrefillOrderActionPrefillCancelOrderAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PrefillOrderActionUnion) AsPrefillModifyOrderAction() (v PrefillOrderActionPrefillModifyOrderAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PrefillOrderActionUnion) RawJSON() string { return u.JSON.raw }

func (r *PrefillOrderActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PrefillOrderActionUnionOrders is an implicit subunion of
// [PrefillOrderActionUnion]. PrefillOrderActionUnionOrders provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PrefillOrderActionUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPrefillNewOrderRequestArray OfPrefillCancelOrderRequestArray
// OfPrefillModifyOrderRequestArray]
type PrefillOrderActionUnionOrders struct {
	// This field will be present if the value is a [[]PrefillNewOrderRequest] instead
	// of an object.
	OfPrefillNewOrderRequestArray []PrefillNewOrderRequest `json:",inline"`
	// This field will be present if the value is a [[]PrefillCancelOrderRequest]
	// instead of an object.
	OfPrefillCancelOrderRequestArray []PrefillCancelOrderRequest `json:",inline"`
	// This field will be present if the value is a [[]PrefillModifyOrderRequest]
	// instead of an object.
	OfPrefillModifyOrderRequestArray []PrefillModifyOrderRequest `json:",inline"`
	JSON                             struct {
		OfPrefillNewOrderRequestArray    respjson.Field
		OfPrefillCancelOrderRequestArray respjson.Field
		OfPrefillModifyOrderRequestArray respjson.Field
		raw                              string
	} `json:"-"`
}

func (r *PrefillOrderActionUnionOrders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Create one or more new orders.
type PrefillOrderActionPrefillNewOrderAction struct {
	// Any of "NEW".
	ActionType string `json:"action_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	PrefillNewOrderAction
}

// Returns the unmodified JSON received from the API
func (r PrefillOrderActionPrefillNewOrderAction) RawJSON() string { return r.JSON.raw }
func (r *PrefillOrderActionPrefillNewOrderAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cancel one or more existing orders.
type PrefillOrderActionPrefillCancelOrderAction struct {
	// Any of "CANCEL".
	ActionType string `json:"action_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	PrefillCancelOrderAction
}

// Returns the unmodified JSON received from the API
func (r PrefillOrderActionPrefillCancelOrderAction) RawJSON() string { return r.JSON.raw }
func (r *PrefillOrderActionPrefillCancelOrderAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Modify one or more existing orders.
type PrefillOrderActionPrefillModifyOrderAction struct {
	// Any of "MODIFY".
	ActionType string `json:"action_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	PrefillModifyOrderAction
}

// Returns the unmodified JSON received from the API
func (r PrefillOrderActionPrefillModifyOrderAction) RawJSON() string { return r.JSON.raw }
func (r *PrefillOrderActionPrefillModifyOrderAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prompt-style button behavior.
type PromptButtonAction struct {
	// Prompt text to submit as the next user turn.
	Prompt string `json:"prompt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptButtonAction) RawJSON() string { return r.JSON.raw }
func (r *PromptButtonAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// StructuredActionUnion contains all possible properties and values from
// [StructuredActionPrefillOrder], [StructuredActionOpenChart],
// [StructuredActionOpenScreener], [StructuredActionOpenEntitlementConsent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type StructuredActionUnion struct {
	// This field is from variant [StructuredActionPrefillOrder].
	PrefillOrder PrefillOrderActionUnion `json:"prefill_order"`
	// This field is from variant [StructuredActionOpenChart].
	OpenChart OpenChartAction `json:"open_chart"`
	// This field is from variant [StructuredActionOpenScreener].
	OpenScreener OpenScreenerAction `json:"open_screener"`
	// This field is from variant [StructuredActionOpenEntitlementConsent].
	OpenEntitlementConsent OpenEntitlementConsentAction `json:"open_entitlement_consent"`
	JSON                   struct {
		PrefillOrder           respjson.Field
		OpenChart              respjson.Field
		OpenScreener           respjson.Field
		OpenEntitlementConsent respjson.Field
		raw                    string
	} `json:"-"`
}

func (u StructuredActionUnion) AsPrefillOrder() (v StructuredActionPrefillOrder) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StructuredActionUnion) AsOpenChart() (v StructuredActionOpenChart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StructuredActionUnion) AsOpenScreener() (v StructuredActionOpenScreener) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StructuredActionUnion) AsOpenEntitlementConsent() (v StructuredActionOpenEntitlementConsent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u StructuredActionUnion) RawJSON() string { return u.JSON.raw }

func (r *StructuredActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prefill an order ticket for user confirmation
type StructuredActionPrefillOrder struct {
	// Prefill an order ticket for user confirmation
	PrefillOrder PrefillOrderActionUnion `json:"prefill_order" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PrefillOrder respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StructuredActionPrefillOrder) RawJSON() string { return r.JSON.raw }
func (r *StructuredActionPrefillOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Open a chart for a symbol
type StructuredActionOpenChart struct {
	// Open a chart for a symbol
	OpenChart OpenChartAction `json:"open_chart" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OpenChart   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StructuredActionOpenChart) RawJSON() string { return r.JSON.raw }
func (r *StructuredActionOpenChart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Open a stock screener with filters
type StructuredActionOpenScreener struct {
	// Open a stock screener with filters
	OpenScreener OpenScreenerAction `json:"open_screener" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OpenScreener respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StructuredActionOpenScreener) RawJSON() string { return r.JSON.raw }
func (r *StructuredActionOpenScreener) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Open entitlement consent flow
type StructuredActionOpenEntitlementConsent struct {
	// Open entitlement consent flow
	OpenEntitlementConsent OpenEntitlementConsentAction `json:"open_entitlement_consent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OpenEntitlementConsent respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StructuredActionOpenEntitlementConsent) RawJSON() string { return r.JSON.raw }
func (r *StructuredActionOpenEntitlementConsent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured-action button behavior.
type StructuredActionButtonAction struct {
	// UUID of a `structured_action` content part in the same message. When a
	// null/undefined value is observed, it indicates it does not apply.
	ActionID string `json:"actionId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StructuredActionButtonAction) RawJSON() string { return r.JSON.raw }
func (r *StructuredActionButtonAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Suggested follow-up buttons rendered at the end of an assistant message.
type SuggestedActionsPayload struct {
	// Ordered message-level buttons.
	ActionButtons []ActionButton `json:"actionButtons"`
	// IDs of buttons clicked by the current user.
	ClickedItemIDs []string `json:"clickedItemIds" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionButtons  respjson.Field
		ClickedItemIDs respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SuggestedActionsPayload) RawJSON() string { return r.JSON.raw }
func (r *SuggestedActionsPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
