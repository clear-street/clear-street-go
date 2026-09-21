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
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/param"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Thread-centric AI assistant for conversational trading. Create threads to start
// conversations, poll response objects for in-progress output, and read finalized
// messages from thread history. Thread/message/response endpoints require an
// explicit account_id. Entitlement endpoints are caller-scoped and use
// account_ids.
//
// V1OmniAIThreadService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1OmniAIThreadService] method instead.
type V1OmniAIThreadService struct {
	options []option.RequestOption
}

// NewV1OmniAIThreadService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1OmniAIThreadService(opts ...option.RequestOption) (r V1OmniAIThreadService) {
	r = V1OmniAIThreadService{}
	r.options = opts
	return
}

// Append a user message to an existing thread and start an assistant response.
// Poll the returned `response_id` via `GET /omni-ai/responses/{response_id}` for
// assistant output.
//
// Only one response may be active per thread. Wait for it to reach a terminal
// status before submitting another turn; otherwise this endpoint returns 409.
//
// The first accepted selected-account message links an unlinked thread. A linked
// thread keeps its account regardless of omission or another selection. A changed
// scope also returns 409 without accepting a turn.
func (r *V1OmniAIThreadService) NewMessage(ctx context.Context, threadID string, body V1OmniAIThreadNewMessageParams, opts ...option.RequestOption) (res *V1OmniAIThreadNewMessageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/omni-ai/threads/%s/messages", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Atomically create a conversation and submit its first user turn. Use `instant`
// with `text` for a prompt, or `deep_insights` with a ticker `target` and optional
// `thesis` for long-form research.
//
// Poll the returned `response_id` via `GET /omni-ai/responses/{response_id}` for
// assistant output.
//
// Omit `account_id` to start without an account. The first accepted turn with a
// selected account links that account permanently. Reuse `Idempotency-Key` only
// for an identical request.
func (r *V1OmniAIThreadService) NewThread(ctx context.Context, body V1OmniAIThreadNewThreadParams, opts ...option.RequestOption) (res *V1OmniAIThreadNewThreadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/omni-ai/threads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List finalized messages, including messages created before the account link.
// Return the latest page by default, in chronological order within each page. Use
// the returned page token to navigate history.
//
// In-progress assistant output is not included. Poll
// `GET /omni-ai/responses/{response_id}` until the response reaches a terminal
// status, then read its finalized message here.
func (r *V1OmniAIThreadService) GetMessages(ctx context.Context, threadID string, query V1OmniAIThreadGetMessagesParams, opts ...option.RequestOption) (res *V1OmniAIThreadGetMessagesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/omni-ai/threads/%s/messages", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Read an owned thread's metadata. Use `GET /omni-ai/threads/{thread_id}/messages`
// for conversation history.
//
// Omission or another account selection does not change authorization.
func (r *V1OmniAIThreadService) GetThreadByID(ctx context.Context, threadID string, query V1OmniAIThreadGetThreadByIDParams, opts ...option.RequestOption) (res *V1OmniAIThreadGetThreadByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/omni-ai/threads/%s", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Look up the currently active response without knowing its `response_id`. Use
// this endpoint when reopening a thread whose assistant turn may still be in
// progress.
//
// An idle owned thread returns HTTP 200 with `data: null`.
func (r *V1OmniAIThreadService) GetThreadResponse(ctx context.Context, threadID string, query V1OmniAIThreadGetThreadResponseParams, opts ...option.RequestOption) (res *V1OmniAIThreadGetThreadResponseResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/omni-ai/threads/%s/response", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List authorized conversation metadata, newest first. Use `page_size` and
// `page_token` for pagination, and the messages endpoint for conversation history.
//
// With `account_id`, list only conversations linked to that account and require
// current account access. Without it, list only conversations with no linked
// account.
func (r *V1OmniAIThreadService) GetThreads(ctx context.Context, query V1OmniAIThreadGetThreadsParams, opts ...option.RequestOption) (res *V1OmniAIThreadGetThreadsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/omni-ai/threads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A snapshot of the widget the user asks about.
type ContextItem struct {
	// Relevant widget data, selections, and units. Use strings for exact decimals and
	// large IDs.
	Data map[string]any `json:"data" api:"required"`
	// Nonblank descriptive kind. New kinds do not require a backend release.
	Kind string `json:"kind" api:"required"`
	// Nonblank attachment label for conversation rendering.
	Label string `json:"label" api:"required"`
	// Client-reported snapshot time. Omit when unknown.
	CapturedAt time.Time `json:"captured_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Kind        respjson.Field
		Label       respjson.Field
		CapturedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextItem) RawJSON() string { return r.JSON.raw }
func (r *ContextItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ContextItem to a ContextItemParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ContextItemParam.Overrides()
func (r ContextItem) ToParam() ContextItemParam {
	return param.Override[ContextItemParam](json.RawMessage(r.RawJSON()))
}

// A snapshot of the widget the user asks about.
//
// The properties Data, Kind, Label are required.
type ContextItemParam struct {
	// Relevant widget data, selections, and units. Use strings for exact decimals and
	// large IDs.
	Data map[string]any `json:"data,omitzero" api:"required"`
	// Nonblank descriptive kind. New kinds do not require a backend release.
	Kind string `json:"kind" api:"required"`
	// Nonblank attachment label for conversation rendering.
	Label string `json:"label" api:"required"`
	// Client-reported snapshot time. Omit when unknown.
	CapturedAt param.Opt[time.Time] `json:"captured_at,omitzero" format:"date-time"`
	paramObj
}

func (r ContextItemParam) MarshalJSON() (data []byte, err error) {
	type shadow ContextItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContextItemParam) UnmarshalJSON(data []byte) error {
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

// Final immutable message.
type Message struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Finalized immutable message content container. Never includes thinking parts.
	Content   MessageContent `json:"content" api:"required"`
	CreatedAt time.Time      `json:"created_at" api:"required" format:"date-time"`
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
	// Immutable snapshots attached to this user message. Omitted when none were
	// supplied. When a null/undefined value is observed, it indicates that there is no
	// available data.
	Context TurnContext `json:"context" api:"nullable"`
	// When a null/undefined value is observed, it indicates it does not apply.
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
		Context     respjson.Field
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
// [MessageContentPartContentPartText],
// [MessageContentPartContentPartStructuredAction],
// [MessageContentPartContentPartChart],
// [MessageContentPartContentPartSuggestedActions],
// [MessageContentPartContentPartCustom].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MessageContentPartUnion struct {
	// This field is from variant [MessageContentPartContentPartText].
	Text string `json:"text"`
	Type string `json:"type"`
	// This field is from variant [MessageContentPartContentPartStructuredAction].
	Action StructuredActionUnion `json:"action"`
	// This field is from variant [MessageContentPartContentPartStructuredAction].
	ActionID string `json:"action_id"`
	// This field is from variant [MessageContentPartContentPartStructuredAction].
	Clicked bool `json:"clicked"`
	// This field is from variant [MessageContentPartContentPartStructuredAction].
	ClickedItemIDs []string `json:"clicked_item_ids"`
	// This field is from variant [MessageContentPartContentPartStructuredAction].
	ItemID string `json:"item_id"`
	// This field is a union of [ChartPayload], [SuggestedActionsPayload], [any]
	Payload MessageContentPartUnionPayload `json:"payload"`
	JSON    struct {
		Text           respjson.Field
		Type           respjson.Field
		Action         respjson.Field
		ActionID       respjson.Field
		Clicked        respjson.Field
		ClickedItemIDs respjson.Field
		ItemID         respjson.Field
		Payload        respjson.Field
		raw            string
	} `json:"-"`
}

func (u MessageContentPartUnion) AsContentPartText() (v MessageContentPartContentPartText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageContentPartUnion) AsContentPartStructuredAction() (v MessageContentPartContentPartStructuredAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageContentPartUnion) AsContentPartChart() (v MessageContentPartContentPartChart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageContentPartUnion) AsContentPartSuggestedActions() (v MessageContentPartContentPartSuggestedActions) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageContentPartUnion) AsContentPartCustom() (v MessageContentPartContentPartCustom) {
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
	ChartID string `json:"chartId"`
	// This field is from variant [ChartPayload].
	Clicked       bool           `json:"clicked"`
	ActionButtons []ActionButton `json:"actionButtons"`
	// This field is from variant [ChartPayload].
	DataChart DataChart `json:"dataChart"`
	// This field is from variant [ChartPayload].
	ItemID string `json:"itemId"`
	// This field is from variant [SuggestedActionsPayload].
	ClickedItemIDs []string `json:"clickedItemIds"`
	JSON           struct {
		OfContentPartCustomPayloadPayload respjson.Field
		ChartID                           respjson.Field
		Clicked                           respjson.Field
		ActionButtons                     respjson.Field
		DataChart                         respjson.Field
		ItemID                            respjson.Field
		ClickedItemIDs                    respjson.Field
		raw                               string
	} `json:"-"`
}

func (r *MessageContentPartUnionPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part.
type MessageContentPartContentPartText struct {
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
func (r MessageContentPartContentPartText) RawJSON() string { return r.JSON.raw }
func (r *MessageContentPartContentPartText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured action content part.
type MessageContentPartContentPartStructuredAction struct {
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
func (r MessageContentPartContentPartStructuredAction) RawJSON() string { return r.JSON.raw }
func (r *MessageContentPartContentPartStructuredAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chart payload content part.
type MessageContentPartContentPartChart struct {
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
func (r MessageContentPartContentPartChart) RawJSON() string { return r.JSON.raw }
func (r *MessageContentPartContentPartChart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Suggested actions payload content part.
type MessageContentPartContentPartSuggestedActions struct {
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
func (r MessageContentPartContentPartSuggestedActions) RawJSON() string { return r.JSON.raw }
func (r *MessageContentPartContentPartSuggestedActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Escape-hatch custom payload content part.
type MessageContentPartContentPartCustom struct {
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
func (r MessageContentPartContentPartCustom) RawJSON() string { return r.JSON.raw }
func (r *MessageContentPartContentPartCustom) UnmarshalJSON(data []byte) error {
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

// Thread metadata.
type Thread struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Title     string    `json:"title" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
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

// Client snapshots attached to one instant-chat user message.
//
// Context is separate from visible message text and does not grant account access.
// The compact JSON representation must not exceed 64 KiB.
type TurnContext struct {
	// One to four snapshots. Each snapshot's data may contain at most 32 levels of
	// nesting.
	Items []ContextItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnContext) RawJSON() string { return r.JSON.raw }
func (r *TurnContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TurnContext to a TurnContextParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TurnContextParam.Overrides()
func (r TurnContext) ToParam() TurnContextParam {
	return param.Override[TurnContextParam](json.RawMessage(r.RawJSON()))
}

// Client snapshots attached to one instant-chat user message.
//
// Context is separate from visible message text and does not grant account access.
// The compact JSON representation must not exceed 64 KiB.
//
// The property Items is required.
type TurnContextParam struct {
	// One to four snapshots. Each snapshot's data may contain at most 32 levels of
	// nesting.
	Items []ContextItemParam `json:"items,omitzero" api:"required"`
	paramObj
}

func (r TurnContextParam) MarshalJSON() (data []byte, err error) {
	type shadow TurnContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TurnContextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIThreadNewMessageResponse struct {
	// Response payload for continuing a thread with a new message.
	Data CreateMessageResponse `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1OmniAIThreadNewMessageResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIThreadNewMessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIThreadNewThreadResponse struct {
	// Response payload for thread creation.
	Data CreateThreadResponse `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1OmniAIThreadNewThreadResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIThreadNewThreadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIThreadGetMessagesResponse struct {
	Data MessageList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1OmniAIThreadGetMessagesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIThreadGetMessagesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIThreadGetThreadByIDResponse struct {
	// Thread metadata.
	Data Thread `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1OmniAIThreadGetThreadByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIThreadGetThreadByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIThreadGetThreadResponseResponse struct {
	// Dynamic pollable response.
	Data Response `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1OmniAIThreadGetThreadResponseResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIThreadGetThreadResponseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIThreadGetThreadsResponse struct {
	Data ThreadList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1OmniAIThreadGetThreadsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIThreadGetThreadsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIThreadNewMessageParams struct {
	Text string `json:"text" api:"required"`
	// Selected account for creation or the first account-linked turn. Omit for an
	// unlinked conversation. An existing account link remains authoritative even when
	// another account is selected.
	AccountID param.Opt[int64] `json:"account_id,omitzero"`
	// Any of "PREFILL_ORDER", "OPEN_CHART", "OPEN_SCREENER",
	// "OPEN_ENTITLEMENT_CONSENT".
	Capabilities []string `json:"capabilities,omitzero"`
	// Snapshots for this instant-chat message. Omission does not remove earlier
	// attachments.
	Context TurnContextParam `json:"context,omitzero"`
	paramObj
}

func (r V1OmniAIThreadNewMessageParams) MarshalJSON() (data []byte, err error) {
	type shadow V1OmniAIThreadNewMessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1OmniAIThreadNewMessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIThreadNewThreadParams struct {
	// Thread creation mode.
	//
	// Any of "instant", "deep_insights".
	Type V1OmniAIThreadNewThreadParamsType `json:"type,omitzero" api:"required"`
	// Selected account for creation or the first account-linked turn. Omit for an
	// unlinked conversation. An existing account link remains authoritative even when
	// another account is selected.
	AccountID param.Opt[int64]  `json:"account_id,omitzero"`
	Text      param.Opt[string] `json:"text,omitzero"`
	Thesis    param.Opt[string] `json:"thesis,omitzero"`
	// Deep-insights target payload.
	Target V1OmniAIThreadNewThreadParamsTarget `json:"target,omitzero"`
	// Any of "PREFILL_ORDER", "OPEN_CHART", "OPEN_SCREENER",
	// "OPEN_ENTITLEMENT_CONSENT".
	Capabilities []string `json:"capabilities,omitzero"`
	// Snapshots for the first instant-chat message. Omit to attach no new context.
	Context TurnContextParam `json:"context,omitzero"`
	paramObj
}

func (r V1OmniAIThreadNewThreadParams) MarshalJSON() (data []byte, err error) {
	type shadow V1OmniAIThreadNewThreadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1OmniAIThreadNewThreadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Thread creation mode.
type V1OmniAIThreadNewThreadParamsType string

const (
	V1OmniAIThreadNewThreadParamsTypeInstant      V1OmniAIThreadNewThreadParamsType = "instant"
	V1OmniAIThreadNewThreadParamsTypeDeepInsights V1OmniAIThreadNewThreadParamsType = "deep_insights"
)

// Deep-insights target payload.
//
// The properties Ticker, Type are required.
type V1OmniAIThreadNewThreadParamsTarget struct {
	Ticker string `json:"ticker" api:"required"`
	// Deep-insights target type. Launch supports ticker-only.
	//
	// Any of "ticker".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r V1OmniAIThreadNewThreadParamsTarget) MarshalJSON() (data []byte, err error) {
	type shadow V1OmniAIThreadNewThreadParamsTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1OmniAIThreadNewThreadParamsTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1OmniAIThreadNewThreadParamsTarget](
		"type", "ticker",
	)
}

type V1OmniAIThreadGetMessagesParams struct {
	// Lists only conversations for this account, or unlinked conversations when
	// omitted. Other reads authorize the resource's linked account. Omit when no
	// account is selected; empty values and the string null are invalid.
	AccountID param.Opt[int64] `query:"account_id,omitzero" json:"-"`
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	paramObj
}

// URLQuery serializes [V1OmniAIThreadGetMessagesParams]'s query parameters as
// `url.Values`.
func (r V1OmniAIThreadGetMessagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1OmniAIThreadGetThreadByIDParams struct {
	// Lists only conversations for this account, or unlinked conversations when
	// omitted. Other reads authorize the resource's linked account. Omit when no
	// account is selected; empty values and the string null are invalid.
	AccountID param.Opt[int64] `query:"account_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1OmniAIThreadGetThreadByIDParams]'s query parameters as
// `url.Values`.
func (r V1OmniAIThreadGetThreadByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1OmniAIThreadGetThreadResponseParams struct {
	// Lists only conversations for this account, or unlinked conversations when
	// omitted. Other reads authorize the resource's linked account. Omit when no
	// account is selected; empty values and the string null are invalid.
	AccountID param.Opt[int64] `query:"account_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1OmniAIThreadGetThreadResponseParams]'s query parameters
// as `url.Values`.
func (r V1OmniAIThreadGetThreadResponseParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1OmniAIThreadGetThreadsParams struct {
	// Lists only conversations for this account, or unlinked conversations when
	// omitted. Other reads authorize the resource's linked account. Omit when no
	// account is selected; empty values and the string null are invalid.
	AccountID param.Opt[int64] `query:"account_id,omitzero" json:"-"`
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	paramObj
}

// URLQuery serializes [V1OmniAIThreadGetThreadsParams]'s query parameters as
// `url.Values`.
func (r V1OmniAIThreadGetThreadsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
