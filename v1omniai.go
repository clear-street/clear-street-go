// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"encoding/json"

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
	// trading_account_ids.
	Entitlements V1OmniAIEntitlementService
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Thread/message/response endpoints require an
	// explicit account_id. Entitlement endpoints are caller-scoped and use
	// trading_account_ids.
	Messages V1OmniAIMessageService
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Thread/message/response endpoints require an
	// explicit account_id. Entitlement endpoints are caller-scoped and use
	// trading_account_ids.
	Responses V1OmniAIResponseService
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Thread/message/response endpoints require an
	// explicit account_id. Entitlement endpoints are caller-scoped and use
	// trading_account_ids.
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
	// Follow-up prompt to submit as the next user message.
	Prompt PromptButtonAction `json:"prompt" api:"nullable"`
	// Structured action in the same message to execute on click.
	StructuredAction StructuredActionButtonAction `json:"structuredAction" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonID         respjson.Field
		Label            respjson.Field
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
	// Buttons associated with this chart.
	ActionButtons []ActionButton `json:"actionButtons"`
	// Explicit series-driven chart definition.
	DataChart DataChart `json:"dataChart" api:"nullable"`
	// Symbol-driven chart definition.
	SymbolChart SymbolChart `json:"symbolChart" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChartID       respjson.Field
		ActionButtons respjson.Field
		DataChart     respjson.Field
		SymbolChart   respjson.Field
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		ActionID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// Action to open a chart for a symbol.
type OpenChartAction struct {
	// Trading symbol to chart
	Symbol string `json:"symbol" api:"required"`
	// Additional chart configuration (indicators, overlays, etc.)
	Extras any `json:"extras"`
	// Chart timeframe (e.g., "1D", "1W", "1M", "3M", "1Y", "5Y")
	Timeframe string `json:"timeframe" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Symbol      respjson.Field
		Extras      respjson.Field
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
	// Stable entitlement agreement family key.
	//
	// Any of "omni_account_data_access".
	AgreementKey              EntitlementAgreementKey `json:"agreement_key" api:"required"`
	Reason                    string                  `json:"reason" api:"required"`
	RequestedEntitlementCodes []EntitlementCode       `json:"requested_entitlement_codes" api:"required"`
	TradingAccountIDs         []int64                 `json:"trading_account_ids" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgreementKey              respjson.Field
		Reason                    respjson.Field
		RequestedEntitlementCodes respjson.Field
		TradingAccountIDs         respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
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
	// Optional field/column selection for screener results.
	FieldFilter []string `json:"field_filter" api:"nullable"`
	// Optional page size.
	PageSize int64 `json:"page_size" api:"nullable"`
	// Optional sort field for screener rows.
	SortBy string `json:"sort_by" api:"nullable"`
	// Optional sort direction (`ASC` or `DESC`).
	SortDirection string `json:"sort_direction" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters       respjson.Field
		FieldFilter   respjson.Field
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
	Orders []CancelOrderRequest `json:"orders" api:"required"`
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

// New-order prefill action.
type PrefillNewOrderAction struct {
	// Orders to prefill using the same shape accepted by the orders API.
	Orders []NewOrderRequest `json:"orders" api:"required"`
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

// PrefillOrderActionUnion contains all possible properties and values from
// [PrefillOrderActionPrefillNewOrderAction],
// [PrefillOrderActionPrefillCancelOrderAction].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PrefillOrderActionUnion struct {
	// This field is a union of [[]NewOrderRequest], [[]CancelOrderRequest]
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
// will be valid: OfNewOrderRequestArray OfCancelOrderRequestArray]
type PrefillOrderActionUnionOrders struct {
	// This field will be present if the value is a [[]NewOrderRequest] instead of an
	// object.
	OfNewOrderRequestArray []NewOrderRequest `json:",inline"`
	// This field will be present if the value is a [[]CancelOrderRequest] instead of
	// an object.
	OfCancelOrderRequestArray []CancelOrderRequest `json:",inline"`
	JSON                      struct {
		OfNewOrderRequestArray    respjson.Field
		OfCancelOrderRequestArray respjson.Field
		raw                       string
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

// A single filter criterion for the screener.
type ScreenerFilter struct {
	// Field to filter on (e.g., "market_cap", "sector", "price")
	Field string `json:"field" api:"required"`
	// Comparison operator (e.g., "eq", "gte", "lte", "in")
	Operator string `json:"operator" api:"required"`
	// Filter value
	Value any `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScreenerFilter) RawJSON() string { return r.JSON.raw }
func (r *ScreenerFilter) UnmarshalJSON(data []byte) error {
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
	// UUID of a `structured_action` content part in the same message.
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionButtons respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SuggestedActionsPayload) RawJSON() string { return r.JSON.raw }
func (r *SuggestedActionsPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chart for a single symbol and timeframe.
type SymbolChart struct {
	Symbol    string `json:"symbol" api:"required"`
	Timeframe string `json:"timeframe" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Symbol      respjson.Field
		Timeframe   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SymbolChart) RawJSON() string { return r.JSON.raw }
func (r *SymbolChart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
