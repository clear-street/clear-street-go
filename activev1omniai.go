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
	// AI assistant for conversational trading interactions.
	Feedback ActiveV1OmniAIFeedbackService
	// AI assistant for conversational trading interactions.
	Runs ActiveV1OmniAIRunService
	// AI assistant for conversational trading interactions.
	Threads ActiveV1OmniAIThreadService
}

// NewActiveV1OmniAIService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActiveV1OmniAIService(opts ...option.RequestOption) (r ActiveV1OmniAIService) {
	r = ActiveV1OmniAIService{}
	r.options = opts
	r.Feedback = NewActiveV1OmniAIFeedbackService(opts...)
	r.Runs = NewActiveV1OmniAIRunService(opts...)
	r.Threads = NewActiveV1OmniAIThreadService(opts...)
	return
}

type ActionButton struct {
	ButtonID string `json:"button_id" api:"required"`
	Label    string `json:"label" api:"required"`
	// Send a follow-up prompt as the next user message
	Action ButtonActionUnion `json:"action" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonID    respjson.Field
		Label       respjson.Field
		Action      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionButton) RawJSON() string { return r.JSON.raw }
func (r *ActionButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ButtonActionUnion contains all possible properties and values from
// [ButtonActionPrompt], [ButtonActionStructuredAction].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ButtonActionUnion struct {
	// This field is from variant [ButtonActionPrompt].
	Prompt     string `json:"prompt"`
	ActionType string `json:"action_type"`
	// This field is from variant [ButtonActionStructuredAction].
	ActionID string `json:"action_id"`
	JSON     struct {
		Prompt     respjson.Field
		ActionType respjson.Field
		ActionID   respjson.Field
		raw        string
	} `json:"-"`
}

func (u ButtonActionUnion) AsPrompt() (v ButtonActionPrompt) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ButtonActionUnion) AsStructuredAction() (v ButtonActionStructuredAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ButtonActionUnion) RawJSON() string { return u.JSON.raw }

func (r *ButtonActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Send a follow-up prompt as the next user message
type ButtonActionPrompt struct {
	// Any of "prompt".
	ActionType string `json:"action_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	PromptButtonAction
}

// Returns the unmodified JSON received from the API
func (r ButtonActionPrompt) RawJSON() string { return r.JSON.raw }
func (r *ButtonActionPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trigger a structured action already present in the same message
type ButtonActionStructuredAction struct {
	// Any of "structured_action".
	ActionType string `json:"action_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	StructuredActionButtonAction
}

// Returns the unmodified JSON received from the API
func (r ButtonActionStructuredAction) RawJSON() string { return r.JSON.raw }
func (r *ButtonActionStructuredAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CancelRunResponse struct {
	Canceled bool `json:"canceled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Canceled    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CancelRunResponse) RawJSON() string { return r.JSON.raw }
func (r *CancelRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Capability allowlist for structured actions.
//
// Clients declare which capabilities they support when starting a run. Omni AI
// will only emit structured actions that match the declared capabilities.
type Capability string

const (
	CapabilityNavigate       Capability = "NAVIGATE"
	CapabilityOpenChatWindow Capability = "OPEN_CHAT_WINDOW"
	CapabilityPrefillOrder   Capability = "PREFILL_ORDER"
	CapabilityOpenChart      Capability = "OPEN_CHART"
	CapabilityOpenScreener   Capability = "OPEN_SCREENER"
)

// ChartKindUnion contains all possible properties and values from
// [ChartKindSymbolChart], [ChartKindDataChart].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChartKindUnion struct {
	// This field is from variant [ChartKindSymbolChart].
	Symbol string `json:"symbol"`
	// This field is from variant [ChartKindSymbolChart].
	Timeframe string `json:"timeframe"`
	ChartType string `json:"chart_type"`
	// This field is from variant [ChartKindDataChart].
	Series []ChartSeries `json:"series"`
	JSON   struct {
		Symbol    respjson.Field
		Timeframe respjson.Field
		ChartType respjson.Field
		Series    respjson.Field
		raw       string
	} `json:"-"`
}

func (u ChartKindUnion) AsSymbolChart() (v ChartKindSymbolChart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChartKindUnion) AsDataChart() (v ChartKindDataChart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChartKindUnion) RawJSON() string { return u.JSON.raw }

func (r *ChartKindUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chart for a specific ticker symbol
type ChartKindSymbolChart struct {
	// Any of "symbol_chart".
	ChartType string `json:"chart_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChartType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	SymbolChart
}

// Returns the unmodified JSON received from the API
func (r ChartKindSymbolChart) RawJSON() string { return r.JSON.raw }
func (r *ChartKindSymbolChart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chart built from explicit data series
type ChartKindDataChart struct {
	// Any of "data_chart".
	ChartType string `json:"chart_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChartType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	DataChart
}

// Returns the unmodified JSON received from the API
func (r ChartKindDataChart) RawJSON() string { return r.JSON.raw }
func (r *ChartKindDataChart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type ChartSeries struct {
	Name   string       `json:"name" api:"required"`
	Points []ChartPoint `json:"points" api:"required"`
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

// ContentPartUnion contains all possible properties and values from
// [ContentPartText], [ContentPartStructuredActionUnion], [ContentPartThinking],
// [ContentPartChart], [ContentPartSuggestedActions], [ContentPartType].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ContentPartUnion struct {
	// This field is from variant [ContentPartText].
	Text string `json:"text"`
	Type string `json:"type"`
	// This field is from variant [ContentPartStructuredActionUnion].
	Orders []OrderPayload `json:"orders"`
	// This field is from variant [ContentPartStructuredActionUnion].
	AccountID  int64  `json:"account_id"`
	ActionType string `json:"action_type"`
	// This field is from variant [ContentPartStructuredActionUnion].
	Symbol string `json:"symbol"`
	Extras any    `json:"extras"`
	// This field is from variant [ContentPartStructuredActionUnion].
	Timeframe string `json:"timeframe"`
	// This field is from variant [ContentPartStructuredActionUnion].
	Filters []ScreenerFilter `json:"filters"`
	// This field is from variant [ContentPartStructuredActionUnion].
	FieldFilter []string `json:"field_filter"`
	// This field is from variant [ContentPartStructuredActionUnion].
	PageSize int64 `json:"page_size"`
	// This field is from variant [ContentPartStructuredActionUnion].
	SortBy string `json:"sort_by"`
	// This field is from variant [ContentPartStructuredActionUnion].
	SortDirection string `json:"sort_direction"`
	// This field is from variant [ContentPartStructuredActionUnion].
	ThreadID string `json:"thread_id"`
	// This field is from variant [ContentPartStructuredActionUnion].
	Title string `json:"title"`
	// This field is from variant [ContentPartStructuredActionUnion].
	Route string `json:"route"`
	// This field is from variant [ContentPartStructuredActionUnion].
	Params any `json:"params"`
	// This field is from variant [ContentPartThinking].
	Thoughts []string `json:"thoughts"`
	// This field is from variant [ContentPartChart].
	ChartID       string         `json:"chart_id"`
	ActionButtons []ActionButton `json:"action_buttons"`
	// This field is from variant [ContentPartChart].
	ChartKind ChartKindUnion `json:"chart_kind"`
	JSON      struct {
		Text          respjson.Field
		Type          respjson.Field
		Orders        respjson.Field
		AccountID     respjson.Field
		ActionType    respjson.Field
		Symbol        respjson.Field
		Extras        respjson.Field
		Timeframe     respjson.Field
		Filters       respjson.Field
		FieldFilter   respjson.Field
		PageSize      respjson.Field
		SortBy        respjson.Field
		SortDirection respjson.Field
		ThreadID      respjson.Field
		Title         respjson.Field
		Route         respjson.Field
		Params        respjson.Field
		Thoughts      respjson.Field
		ChartID       respjson.Field
		ActionButtons respjson.Field
		ChartKind     respjson.Field
		raw           string
	} `json:"-"`
}

func (u ContentPartUnion) AsText() (v ContentPartText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentPartUnion) AsStructuredAction() (v ContentPartStructuredActionUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentPartUnion) AsThinking() (v ContentPartThinking) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentPartUnion) AsChart() (v ContentPartChart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentPartUnion) AsSuggestedActions() (v ContentPartSuggestedActions) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentPartUnion) AsContentPartType() (v ContentPartType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContentPartUnion) RawJSON() string { return u.JSON.raw }

func (r *ContentPartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Plain text content
type ContentPartText struct {
	// Any of "text".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartText
}

// Returns the unmodified JSON received from the API
func (r ContentPartText) RawJSON() string { return r.JSON.raw }
func (r *ContentPartText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContentPartStructuredActionUnion contains all possible properties and values
// from [ContentPartStructuredActionPrefillOrder],
// [ContentPartStructuredActionOpenChart],
// [ContentPartStructuredActionOpenScreener],
// [ContentPartStructuredActionOpenChatWindow],
// [ContentPartStructuredActionNavigate].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ContentPartStructuredActionUnion struct {
	// This field is from variant [ContentPartStructuredActionPrefillOrder].
	Orders []OrderPayload `json:"orders"`
	// This field is from variant [ContentPartStructuredActionPrefillOrder].
	AccountID  int64  `json:"account_id"`
	ActionType string `json:"action_type"`
	Type       string `json:"type"`
	// This field is from variant [ContentPartStructuredActionOpenChart].
	Symbol string `json:"symbol"`
	Extras any    `json:"extras"`
	// This field is from variant [ContentPartStructuredActionOpenChart].
	Timeframe string `json:"timeframe"`
	// This field is from variant [ContentPartStructuredActionOpenScreener].
	Filters []ScreenerFilter `json:"filters"`
	// This field is from variant [ContentPartStructuredActionOpenScreener].
	FieldFilter []string `json:"field_filter"`
	// This field is from variant [ContentPartStructuredActionOpenScreener].
	PageSize int64 `json:"page_size"`
	// This field is from variant [ContentPartStructuredActionOpenScreener].
	SortBy string `json:"sort_by"`
	// This field is from variant [ContentPartStructuredActionOpenScreener].
	SortDirection string `json:"sort_direction"`
	// This field is from variant [ContentPartStructuredActionOpenChatWindow].
	ThreadID string `json:"thread_id"`
	// This field is from variant [ContentPartStructuredActionOpenChatWindow].
	Title string `json:"title"`
	// This field is from variant [ContentPartStructuredActionNavigate].
	Route string `json:"route"`
	// This field is from variant [ContentPartStructuredActionNavigate].
	Params any `json:"params"`
	JSON   struct {
		Orders        respjson.Field
		AccountID     respjson.Field
		ActionType    respjson.Field
		Type          respjson.Field
		Symbol        respjson.Field
		Extras        respjson.Field
		Timeframe     respjson.Field
		Filters       respjson.Field
		FieldFilter   respjson.Field
		PageSize      respjson.Field
		SortBy        respjson.Field
		SortDirection respjson.Field
		ThreadID      respjson.Field
		Title         respjson.Field
		Route         respjson.Field
		Params        respjson.Field
		raw           string
	} `json:"-"`
}

func (u ContentPartStructuredActionUnion) AsPrefillOrder() (v ContentPartStructuredActionPrefillOrder) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentPartStructuredActionUnion) AsOpenChart() (v ContentPartStructuredActionOpenChart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentPartStructuredActionUnion) AsOpenScreener() (v ContentPartStructuredActionOpenScreener) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentPartStructuredActionUnion) AsOpenChatWindow() (v ContentPartStructuredActionOpenChatWindow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentPartStructuredActionUnion) AsNavigate() (v ContentPartStructuredActionNavigate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContentPartStructuredActionUnion) RawJSON() string { return u.JSON.raw }

func (r *ContentPartStructuredActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prefill an order ticket for user confirmation
type ContentPartStructuredActionPrefillOrder struct {
	// Any of "prefill_order".
	ActionType string `json:"action_type" api:"required"`
	// Any of "structured_action".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	PrefillOrderAction
}

// Returns the unmodified JSON received from the API
func (r ContentPartStructuredActionPrefillOrder) RawJSON() string { return r.JSON.raw }
func (r *ContentPartStructuredActionPrefillOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Open a chart for a symbol
type ContentPartStructuredActionOpenChart struct {
	// Any of "open_chart".
	ActionType string `json:"action_type" api:"required"`
	// Any of "structured_action".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	OpenChartAction
}

// Returns the unmodified JSON received from the API
func (r ContentPartStructuredActionOpenChart) RawJSON() string { return r.JSON.raw }
func (r *ContentPartStructuredActionOpenChart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Open a stock screener with filters
type ContentPartStructuredActionOpenScreener struct {
	// Any of "open_screener".
	ActionType string `json:"action_type" api:"required"`
	// Any of "structured_action".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	OpenScreenerAction
}

// Returns the unmodified JSON received from the API
func (r ContentPartStructuredActionOpenScreener) RawJSON() string { return r.JSON.raw }
func (r *ContentPartStructuredActionOpenScreener) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Open a chat window
type ContentPartStructuredActionOpenChatWindow struct {
	// Any of "open_chat_window".
	ActionType string `json:"action_type" api:"required"`
	// Any of "structured_action".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	OpenChatWindowAction
}

// Returns the unmodified JSON received from the API
func (r ContentPartStructuredActionOpenChatWindow) RawJSON() string { return r.JSON.raw }
func (r *ContentPartStructuredActionOpenChatWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Navigate to a client route
type ContentPartStructuredActionNavigate struct {
	// Any of "navigate".
	ActionType string `json:"action_type" api:"required"`
	// Any of "structured_action".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	NavigateAction
}

// Returns the unmodified JSON received from the API
func (r ContentPartStructuredActionNavigate) RawJSON() string { return r.JSON.raw }
func (r *ContentPartStructuredActionNavigate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model reasoning/thinking content and tool call status indicators
type ContentPartThinking struct {
	// Any of "thinking".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartThinking
}

// Returns the unmodified JSON received from the API
func (r ContentPartThinking) RawJSON() string { return r.JSON.raw }
func (r *ContentPartThinking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed inline chart (symbol or data-driven)
type ContentPartChart struct {
	// Any of "chart".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartChart
}

// Returns the unmodified JSON received from the API
func (r ContentPartChart) RawJSON() string { return r.JSON.raw }
func (r *ContentPartChart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message-level follow-up action buttons
type ContentPartSuggestedActions struct {
	// Any of "suggested_actions".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartSuggestedActions
}

// Returns the unmodified JSON received from the API
func (r ContentPartSuggestedActions) RawJSON() string { return r.JSON.raw }
func (r *ContentPartSuggestedActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom/extensible content
type ContentPartType struct {
	// Any of "custom".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentPartType) RawJSON() string { return r.JSON.raw }
func (r *ContentPartType) UnmarshalJSON(data []byte) error {
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

type DataChart struct {
	Series []ChartSeries `json:"series" api:"required"`
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

type GetRunResponse struct {
	Events        []any  `json:"events" api:"required"`
	Run           Run    `json:"run" api:"required"`
	NextPageToken string `json:"next_page_token" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events        respjson.Field
		Run           respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetRunResponse) RawJSON() string { return r.JSON.raw }
func (r *GetRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetThreadResponse struct {
	Thread Thread `json:"thread" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Thread      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetThreadResponse) RawJSON() string { return r.JSON.raw }
func (r *GetThreadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListMessagesResponse struct {
	Messages      []Message `json:"messages" api:"required"`
	NextPageToken string    `json:"next_page_token" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Messages      respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListMessagesResponse) RawJSON() string { return r.JSON.raw }
func (r *ListMessagesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListThreadsResponse struct {
	Threads       []Thread `json:"threads" api:"required"`
	NextPageToken string   `json:"next_page_token" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Threads       respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListThreadsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListThreadsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Message struct {
	// Denormalized text content for search/display
	ContentText string `json:"content_text" api:"required"`
	CreatedAt   string `json:"created_at" api:"required"`
	// Any of "UNSPECIFIED", "SYSTEM", "USER", "ASSISTANT", "TOOL".
	Role         MessageRole `json:"role" api:"required"`
	Seq          int64       `json:"seq" api:"required"`
	ID           string      `json:"id" api:"nullable" format:"uuid"`
	AuthorUserID string      `json:"author_user_id" api:"nullable"`
	// Parsed content parts (text, thinking, and structured actions)
	Content  MessageContent `json:"content" api:"nullable"`
	Metadata any            `json:"metadata" api:"nullable"`
	RunID    string         `json:"run_id" api:"nullable" format:"uuid"`
	ThreadID string         `json:"thread_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentText  respjson.Field
		CreatedAt    respjson.Field
		Role         respjson.Field
		Seq          respjson.Field
		ID           respjson.Field
		AuthorUserID respjson.Field
		Content      respjson.Field
		Metadata     respjson.Field
		RunID        respjson.Field
		ThreadID     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Message) RawJSON() string { return r.JSON.raw }
func (r *Message) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message content containing text and structured action parts.
type MessageContent struct {
	Parts []ContentPartUnion `json:"parts" api:"required"`
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

type MessageRole string

const (
	MessageRoleUnspecified MessageRole = "UNSPECIFIED"
	MessageRoleSystem      MessageRole = "SYSTEM"
	MessageRoleUser        MessageRole = "USER"
	MessageRoleAssistant   MessageRole = "ASSISTANT"
	MessageRoleTool        MessageRole = "TOOL"
)

// Action to navigate to a client route.
type NavigateAction struct {
	// Route path or key
	Route string `json:"route" api:"required"`
	// Route parameters
	Params any `json:"params"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Route       respjson.Field
		Params      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NavigateAction) RawJSON() string { return r.JSON.raw }
func (r *NavigateAction) UnmarshalJSON(data []byte) error {
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

// Action to open a chat window.
type OpenChatWindowAction struct {
	// Additional configuration
	Extras any `json:"extras"`
	// Thread ID to open (None to open a new chat window)
	ThreadID string `json:"thread_id" api:"nullable"`
	// Window title
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Extras      respjson.Field
		ThreadID    respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenChatWindowAction) RawJSON() string { return r.JSON.raw }
func (r *OpenChatWindowAction) UnmarshalJSON(data []byte) error {
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
	// Order type
	//
	// Any of "MARKET", "LIMIT", "STOP", "STOP_LIMIT", "TRAILING_STOP",
	// "TRAILING_STOP_LIMIT", "OTHER".
	OrderType OrderType `json:"order_type" api:"required"`
	// Quantity (shares for stocks, contracts for options)
	Quantity string `json:"quantity" api:"required"`
	// Type of security
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	SecurityType SecurityType `json:"security_type" api:"required"`
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
	// Stop price (required for STOP and STOP_LIMIT orders)
	StopPrice string `json:"stop_price" api:"nullable"`
	// Execution strategy (simplified enum, not the full strategy params for now)
	//
	// Any of "SOR", "VWAP", "TWAP", "DARK", "DMA", "AP", "POV".
	Strategy OrderStrategyType `json:"strategy" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrderType    respjson.Field
		Quantity     respjson.Field
		SecurityType respjson.Field
		Side         respjson.Field
		Symbol       respjson.Field
		TimeInForce  respjson.Field
		LimitPrice   respjson.Field
		StopPrice    respjson.Field
		Strategy     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
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
	// The orders to prefill
	Orders []OrderPayload `json:"orders" api:"required"`
	// Account to prefill for (if known from context)
	AccountID int64 `json:"account_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Orders      respjson.Field
		AccountID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrefillOrderAction) RawJSON() string { return r.JSON.raw }
func (r *PrefillOrderAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptButtonAction struct {
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

type Run struct {
	CreatedAt string `json:"created_at" api:"required"`
	// Any of "UNSPECIFIED", "QUEUED", "RUNNING", "SUCCEEDED", "FAILED", "CANCELED".
	Status       RunStatus    `json:"status" api:"required"`
	ID           string       `json:"id" api:"nullable" format:"uuid"`
	Capabilities []Capability `json:"capabilities"`
	EndedAt      string       `json:"ended_at" api:"nullable"`
	Error        any          `json:"error" api:"nullable"`
	StartedAt    string       `json:"started_at" api:"nullable"`
	ThreadID     string       `json:"thread_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		Status       respjson.Field
		ID           respjson.Field
		Capabilities respjson.Field
		EndedAt      respjson.Field
		Error        respjson.Field
		StartedAt    respjson.Field
		ThreadID     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Run) RawJSON() string { return r.JSON.raw }
func (r *Run) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RunStatus string

const (
	RunStatusUnspecified RunStatus = "UNSPECIFIED"
	RunStatusQueued      RunStatus = "QUEUED"
	RunStatusRunning     RunStatus = "RUNNING"
	RunStatusSucceeded   RunStatus = "SUCCEEDED"
	RunStatusFailed      RunStatus = "FAILED"
	RunStatusCanceled    RunStatus = "CANCELED"
)

type StartRunResponse struct {
	Run         Run     `json:"run" api:"required"`
	Thread      Thread  `json:"thread" api:"required"`
	UserMessage Message `json:"user_message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Run         respjson.Field
		Thread      respjson.Field
		UserMessage respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StartRunResponse) RawJSON() string { return r.JSON.raw }
func (r *StartRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StructuredActionButtonAction struct {
	ActionID string `json:"action_id" api:"nullable" format:"uuid"`
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

type Thread struct {
	AccountID   string `json:"account_id" api:"required"`
	CreatedAt   string `json:"created_at" api:"required"`
	Description string `json:"description" api:"required"`
	OwnerUserID string `json:"owner_user_id" api:"required"`
	Title       string `json:"title" api:"required"`
	UpdatedAt   string `json:"updated_at" api:"required"`
	ID          string `json:"id" api:"nullable" format:"uuid"`
	Metadata    any    `json:"metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID   respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		OwnerUserID respjson.Field
		Title       respjson.Field
		UpdatedAt   respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Thread) RawJSON() string { return r.JSON.raw }
func (r *Thread) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
