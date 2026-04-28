// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"encoding/json"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/respjson"
)

// ActiveV1OmniAIService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1OmniAIService] method instead.
type ActiveV1OmniAIService struct {
	options []option.RequestOption
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Thread/message/response endpoints require an
	// explicit account_id. Entitlement endpoints are caller-scoped and use
	// trading_account_ids.
	EntitlementAgreements ActiveV1OmniAIEntitlementAgreementService
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Thread/message/response endpoints require an
	// explicit account_id. Entitlement endpoints are caller-scoped and use
	// trading_account_ids.
	Entitlements ActiveV1OmniAIEntitlementService
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Thread/message/response endpoints require an
	// explicit account_id. Entitlement endpoints are caller-scoped and use
	// trading_account_ids.
	Messages ActiveV1OmniAIMessageService
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Thread/message/response endpoints require an
	// explicit account_id. Entitlement endpoints are caller-scoped and use
	// trading_account_ids.
	Responses ActiveV1OmniAIResponseService
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Thread/message/response endpoints require an
	// explicit account_id. Entitlement endpoints are caller-scoped and use
	// trading_account_ids.
	Threads ActiveV1OmniAIThreadService
}

// NewActiveV1OmniAIService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActiveV1OmniAIService(opts ...option.RequestOption) (r ActiveV1OmniAIService) {
	r = ActiveV1OmniAIService{}
	r.options = opts
	r.EntitlementAgreements = NewActiveV1OmniAIEntitlementAgreementService(opts...)
	r.Entitlements = NewActiveV1OmniAIEntitlementService(opts...)
	r.Messages = NewActiveV1OmniAIMessageService(opts...)
	r.Responses = NewActiveV1OmniAIResponseService(opts...)
	r.Threads = NewActiveV1OmniAIThreadService(opts...)
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

type CancelResponsePayload struct {
	Canceled bool `json:"canceled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Canceled    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CancelResponsePayload) RawJSON() string { return r.JSON.raw }
func (r *CancelResponsePayload) UnmarshalJSON(data []byte) error {
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

type CreateFeedbackResponse struct {
	CreatedAt  string `json:"created_at" api:"required"`
	FeedbackID string `json:"feedback_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		FeedbackID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateFeedbackResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateFeedbackResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response payload for continuing a thread with a new message.
type CreateMessageResponse struct {
	ResponseID    string `json:"response_id" api:"required" format:"uuid"`
	ThreadID      string `json:"thread_id" api:"required" format:"uuid"`
	UserMessageID string `json:"user_message_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResponseID    respjson.Field
		ThreadID      respjson.Field
		UserMessageID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateMessageResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateMessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response payload for thread creation.
type CreateThreadResponse struct {
	ResponseID    string `json:"response_id" api:"required" format:"uuid"`
	ThreadID      string `json:"thread_id" api:"required" format:"uuid"`
	UserMessageID string `json:"user_message_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResponseID    respjson.Field
		ThreadID      respjson.Field
		UserMessageID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateThreadResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateThreadResponse) UnmarshalJSON(data []byte) error {
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

// Shared sanitized error payload.
type ErrorStatus struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	Details any    `json:"details" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Details     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ErrorStatus) RawJSON() string { return r.JSON.raw }
func (r *ErrorStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Final immutable message.
type Message struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Finalized immutable message content container. Never includes thinking parts.
	Content   MessageContent `json:"content" api:"required"`
	CreatedAt string         `json:"created_at" api:"required"`
	// Immutable terminal outcome for a finalized assistant message.
	//
	// Any of "completed", "errored", "canceled".
	Outcome MessageOutcome `json:"outcome" api:"required"`
	// Finalized message role in the public contract.
	//
	// Any of "USER", "ASSISTANT".
	Role     MessageRole `json:"role" api:"required"`
	Seq      int64       `json:"seq" api:"required"`
	ThreadID string      `json:"thread_id" api:"required" format:"uuid"`
	// Shared sanitized error payload.
	Error ErrorStatus `json:"error" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Content     respjson.Field
		CreatedAt   respjson.Field
		Outcome     respjson.Field
		Role        respjson.Field
		Seq         respjson.Field
		ThreadID    respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Message) RawJSON() string { return r.JSON.raw }
func (r *Message) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Finalized immutable message content container. Never includes thinking parts.
type MessageContent struct {
	Parts []MessageContentPartUnion `json:"parts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageContent) RawJSON() string { return r.JSON.raw }
func (r *MessageContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageContentPartUnion contains all possible properties and values from
// [MessageContentPartObject], [MessageContentPartObject2],
// [MessageContentPartObject3], [MessageContentPartObject4],
// [MessageContentPartObject5].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MessageContentPartUnion struct {
	// This field is from variant [MessageContentPartObject].
	Text string `json:"text"`
	Type string `json:"type"`
	// This field is from variant [MessageContentPartObject2].
	Action StructuredActionUnion `json:"action"`
	// This field is from variant [MessageContentPartObject2].
	ActionID string `json:"action_id"`
	// This field is a union of [ChartPayload], [SuggestedActionsPayload], [any]
	Payload MessageContentPartUnionPayload `json:"payload"`
	JSON    struct {
		Text     respjson.Field
		Type     respjson.Field
		Action   respjson.Field
		ActionID respjson.Field
		Payload  respjson.Field
		raw      string
	} `json:"-"`
}

func (u MessageContentPartUnion) AsMessageContentPartObject() (v MessageContentPartObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageContentPartUnion) AsMessageContentPartObject2() (v MessageContentPartObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageContentPartUnion) AsMessageContentPartObject3() (v MessageContentPartObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageContentPartUnion) AsMessageContentPartObject4() (v MessageContentPartObject4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageContentPartUnion) AsMessageContentPartObject5() (v MessageContentPartObject5) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageContentPartUnion) RawJSON() string { return u.JSON.raw }

func (r *MessageContentPartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageContentPartUnionPayload is an implicit subunion of
// [MessageContentPartUnion]. MessageContentPartUnionPayload provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MessageContentPartUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfContentPartCustomPayloadPayload]
type MessageContentPartUnionPayload struct {
	// This field will be present if the value is a [any] instead of an object.
	OfContentPartCustomPayloadPayload any `json:",inline"`
	// This field is from variant [ChartPayload].
	ChartID       string         `json:"chartId"`
	ActionButtons []ActionButton `json:"actionButtons"`
	// This field is from variant [ChartPayload].
	DataChart DataChart `json:"dataChart"`
	// This field is from variant [ChartPayload].
	SymbolChart SymbolChart `json:"symbolChart"`
	JSON        struct {
		OfContentPartCustomPayloadPayload respjson.Field
		ChartID                           respjson.Field
		ActionButtons                     respjson.Field
		DataChart                         respjson.Field
		SymbolChart                       respjson.Field
		raw                               string
	} `json:"-"`
}

func (r *MessageContentPartUnionPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part.
type MessageContentPartObject struct {
	// Any of "text".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartTextPayload
}

// Returns the unmodified JSON received from the API
func (r MessageContentPartObject) RawJSON() string { return r.JSON.raw }
func (r *MessageContentPartObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured action content part.
type MessageContentPartObject2 struct {
	// Any of "structured_action".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartStructuredActionPayload
}

// Returns the unmodified JSON received from the API
func (r MessageContentPartObject2) RawJSON() string { return r.JSON.raw }
func (r *MessageContentPartObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chart payload content part.
type MessageContentPartObject3 struct {
	// Any of "chart".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartChartPayload
}

// Returns the unmodified JSON received from the API
func (r MessageContentPartObject3) RawJSON() string { return r.JSON.raw }
func (r *MessageContentPartObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Suggested actions payload content part.
type MessageContentPartObject4 struct {
	// Any of "suggested_actions".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartSuggestedActionsPayload
}

// Returns the unmodified JSON received from the API
func (r MessageContentPartObject4) RawJSON() string { return r.JSON.raw }
func (r *MessageContentPartObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Escape-hatch custom payload content part.
type MessageContentPartObject5 struct {
	// Any of "custom".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartCustomPayload
}

// Returns the unmodified JSON received from the API
func (r MessageContentPartObject5) RawJSON() string { return r.JSON.raw }
func (r *MessageContentPartObject5) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageList []Message

// Immutable terminal outcome for a finalized assistant message.
type MessageOutcome string

const (
	MessageOutcomeCompleted MessageOutcome = "completed"
	MessageOutcomeErrored   MessageOutcome = "errored"
	MessageOutcomeCanceled  MessageOutcome = "canceled"
)

// Finalized message role in the public contract.
type MessageRole string

const (
	MessageRoleUser      MessageRole = "USER"
	MessageRoleAssistant MessageRole = "ASSISTANT"
)

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
	AgreementKey              string   `json:"agreement_key" api:"required"`
	Reason                    string   `json:"reason" api:"required"`
	RequestedEntitlementCodes []string `json:"requested_entitlement_codes" api:"required"`
	TradingAccountIDs         []int64  `json:"trading_account_ids" api:"required"`
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

// Order payload for prefilling an order ticket.
//
// This schema aligns with the NewOrderRequest schema used for order submission,
// containing the fields needed to prefill an order ticket or submit via API.
type OrderPayload struct {
	// Type of security
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	InstrumentType SecurityType `json:"instrument_type" api:"required"`
	// Order type
	//
	// Any of "MARKET", "LIMIT", "STOP", "STOP_LIMIT", "TRAILING_STOP",
	// "TRAILING_STOP_LIMIT", "OTHER".
	OrderType OrderType `json:"order_type" api:"required"`
	// Quantity (shares for stocks, contracts for options)
	Quantity string `json:"quantity" api:"required"`
	// Order side
	//
	// Any of "BUY", "SELL", "SELL_SHORT", "OTHER".
	Side Side `json:"side" api:"required"`
	// Trading symbol (e.g., "AAPL" for equities, OSI for options)
	Symbol string `json:"symbol" api:"required"`
	// Time in force
	//
	// Any of "DAY", "GOOD_TILL_CANCEL", "IMMEDIATE_OR_CANCEL", "FILL_OR_KILL",
	// "GOOD_TILL_DATE", "AT_THE_OPENING", "AT_THE_CLOSE", "GOOD_TILL_CROSSING",
	// "GOOD_THROUGH_CROSSING", "AT_CROSSING", "OTHER".
	TimeInForce TimeInForce `json:"time_in_force" api:"required"`
	// Limit price (required for LIMIT and STOP_LIMIT orders)
	LimitPrice string `json:"limit_price" api:"nullable"`
	// Existing order identifier. Required for cancel actions.
	OrderID string `json:"order_id" api:"nullable"`
	// Stop price (required for STOP and STOP_LIMIT orders)
	StopPrice string `json:"stop_price" api:"nullable"`
	// Execution strategy (simplified enum, not the full strategy params for now)
	//
	// Any of "SOR", "VWAP", "TWAP", "DARK", "DMA", "AP", "POV".
	Strategy OrderStrategyType `json:"strategy" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InstrumentType respjson.Field
		OrderType      respjson.Field
		Quantity       respjson.Field
		Side           respjson.Field
		Symbol         respjson.Field
		TimeInForce    respjson.Field
		LimitPrice     respjson.Field
		OrderID        respjson.Field
		StopPrice      respjson.Field
		Strategy       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderPayload) RawJSON() string { return r.JSON.raw }
func (r *OrderPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Simplified order strategy type for prefill actions.
//
// This is a simplified enum compared to the full OrderStrategy with params,
// suitable for indicating the desired strategy without full configuration.
type OrderStrategyType string

const (
	OrderStrategyTypeSor  OrderStrategyType = "SOR"
	OrderStrategyTypeVwap OrderStrategyType = "VWAP"
	OrderStrategyTypeTwap OrderStrategyType = "TWAP"
	OrderStrategyTypeDark OrderStrategyType = "DARK"
	OrderStrategyTypeDma  OrderStrategyType = "DMA"
	OrderStrategyTypeAp   OrderStrategyType = "AP"
	OrderStrategyTypePov  OrderStrategyType = "POV"
)

// Action to prefill order details for user confirmation.
//
// The user must review and authorize the order before submission to the trading
// API. This action provides parsed order data that can be used to prefill an order
// ticket UI or submitted directly via the orders API after user confirmation.
type PrefillOrderAction struct {
	// Order operation represented by this prefill action.
	//
	// Any of "NEW", "CANCEL".
	ActionType PrefillOrderActionType `json:"action_type" api:"required"`
	// The orders to prefill
	Orders []OrderPayload `json:"orders" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType  respjson.Field
		Orders      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrefillOrderAction) RawJSON() string { return r.JSON.raw }
func (r *PrefillOrderAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Operation represented by a prefill order action.
type PrefillOrderActionType string

const (
	PrefillOrderActionTypeNew    PrefillOrderActionType = "NEW"
	PrefillOrderActionTypeCancel PrefillOrderActionType = "CANCEL"
)

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

// Dynamic pollable response.
type Response struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Dynamic lifecycle status for a pollable response.
	//
	// Any of "queued", "running", "succeeded", "failed", "canceled".
	Status        ResponseStatus `json:"status" api:"required"`
	ThreadID      string         `json:"thread_id" api:"required" format:"uuid"`
	UserMessageID string         `json:"user_message_id" api:"required" format:"uuid"`
	// Dynamic response content container. May include thinking parts.
	Content ResponseContent `json:"content" api:"nullable"`
	// Shared sanitized error payload.
	Error           ErrorStatus `json:"error" api:"nullable"`
	OutputMessageID string      `json:"output_message_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Status          respjson.Field
		ThreadID        respjson.Field
		UserMessageID   respjson.Field
		Content         respjson.Field
		Error           respjson.Field
		OutputMessageID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Response) RawJSON() string { return r.JSON.raw }
func (r *Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dynamic response content container. May include thinking parts.
type ResponseContent struct {
	Parts []ResponseContentPartUnion `json:"parts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseContent) RawJSON() string { return r.JSON.raw }
func (r *ResponseContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseContentPartUnion contains all possible properties and values from
// [ResponseContentPartObject], [ResponseContentPartObject2],
// [ResponseContentPartObject3], [ResponseContentPartObject4],
// [ResponseContentPartObject5], [ResponseContentPartObject6].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseContentPartUnion struct {
	// This field is from variant [ResponseContentPartObject].
	Text string `json:"text"`
	Type string `json:"type"`
	// This field is from variant [ResponseContentPartObject2].
	Thoughts []string `json:"thoughts"`
	// This field is from variant [ResponseContentPartObject3].
	Action StructuredActionUnion `json:"action"`
	// This field is from variant [ResponseContentPartObject3].
	ActionID string `json:"action_id"`
	// This field is a union of [ChartPayload], [SuggestedActionsPayload], [any]
	Payload ResponseContentPartUnionPayload `json:"payload"`
	JSON    struct {
		Text     respjson.Field
		Type     respjson.Field
		Thoughts respjson.Field
		Action   respjson.Field
		ActionID respjson.Field
		Payload  respjson.Field
		raw      string
	} `json:"-"`
}

func (u ResponseContentPartUnion) AsResponseContentPartObject() (v ResponseContentPartObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseContentPartUnion) AsResponseContentPartObject2() (v ResponseContentPartObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseContentPartUnion) AsResponseContentPartObject3() (v ResponseContentPartObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseContentPartUnion) AsResponseContentPartObject4() (v ResponseContentPartObject4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseContentPartUnion) AsResponseContentPartObject5() (v ResponseContentPartObject5) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseContentPartUnion) AsResponseContentPartObject6() (v ResponseContentPartObject6) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseContentPartUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseContentPartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseContentPartUnionPayload is an implicit subunion of
// [ResponseContentPartUnion]. ResponseContentPartUnionPayload provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseContentPartUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfContentPartCustomPayloadPayload]
type ResponseContentPartUnionPayload struct {
	// This field will be present if the value is a [any] instead of an object.
	OfContentPartCustomPayloadPayload any `json:",inline"`
	// This field is from variant [ChartPayload].
	ChartID       string         `json:"chartId"`
	ActionButtons []ActionButton `json:"actionButtons"`
	// This field is from variant [ChartPayload].
	DataChart DataChart `json:"dataChart"`
	// This field is from variant [ChartPayload].
	SymbolChart SymbolChart `json:"symbolChart"`
	JSON        struct {
		OfContentPartCustomPayloadPayload respjson.Field
		ChartID                           respjson.Field
		ActionButtons                     respjson.Field
		DataChart                         respjson.Field
		SymbolChart                       respjson.Field
		raw                               string
	} `json:"-"`
}

func (r *ResponseContentPartUnionPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part.
type ResponseContentPartObject struct {
	// Any of "text".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartTextPayload
}

// Returns the unmodified JSON received from the API
func (r ResponseContentPartObject) RawJSON() string { return r.JSON.raw }
func (r *ResponseContentPartObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Thinking content part shown on dynamic response polling.
type ResponseContentPartObject2 struct {
	// Any of "thinking".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartThinkingPayload
}

// Returns the unmodified JSON received from the API
func (r ResponseContentPartObject2) RawJSON() string { return r.JSON.raw }
func (r *ResponseContentPartObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured action content part.
type ResponseContentPartObject3 struct {
	// Any of "structured_action".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartStructuredActionPayload
}

// Returns the unmodified JSON received from the API
func (r ResponseContentPartObject3) RawJSON() string { return r.JSON.raw }
func (r *ResponseContentPartObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chart payload content part.
type ResponseContentPartObject4 struct {
	// Any of "chart".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartChartPayload
}

// Returns the unmodified JSON received from the API
func (r ResponseContentPartObject4) RawJSON() string { return r.JSON.raw }
func (r *ResponseContentPartObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Suggested actions payload content part.
type ResponseContentPartObject5 struct {
	// Any of "suggested_actions".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartSuggestedActionsPayload
}

// Returns the unmodified JSON received from the API
func (r ResponseContentPartObject5) RawJSON() string { return r.JSON.raw }
func (r *ResponseContentPartObject5) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Escape-hatch custom payload content part.
type ResponseContentPartObject6 struct {
	// Any of "custom".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartCustomPayload
}

// Returns the unmodified JSON received from the API
func (r ResponseContentPartObject6) RawJSON() string { return r.JSON.raw }
func (r *ResponseContentPartObject6) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dynamic lifecycle status for a pollable response.
type ResponseStatus string

const (
	ResponseStatusQueued    ResponseStatus = "queued"
	ResponseStatusRunning   ResponseStatus = "running"
	ResponseStatusSucceeded ResponseStatus = "succeeded"
	ResponseStatusFailed    ResponseStatus = "failed"
	ResponseStatusCanceled  ResponseStatus = "canceled"
)

// StructuredActionUnion contains all possible properties and values from
// [StructuredActionPrefillOrder], [StructuredActionOpenChart],
// [StructuredActionOpenScreener], [StructuredActionOpenEntitlementConsent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type StructuredActionUnion struct {
	// This field is from variant [StructuredActionPrefillOrder].
	PrefillOrder PrefillOrderAction `json:"prefill_order"`
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
	PrefillOrder PrefillOrderAction `json:"prefill_order" api:"required"`
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

// Thread metadata returned by list/get thread endpoints.
type Thread struct {
	ID          string `json:"id" api:"required" format:"uuid"`
	CreatedAt   string `json:"created_at" api:"required"`
	Description string `json:"description" api:"required"`
	Title       string `json:"title" api:"required"`
	UpdatedAt   string `json:"updated_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Title       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Thread) RawJSON() string { return r.JSON.raw }
func (r *Thread) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ThreadList []Thread
