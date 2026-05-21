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

// Continue an existing conversation thread.
//
// Appends a new user message to the thread and starts an assistant response. Only
// one response may be active per thread at a time — if the previous turn is still
// in progress, this endpoint returns **409 Conflict**. Wait for the active
// response to reach a terminal status before submitting the next turn.
//
// Poll the returned `response_id` via `GET /omni-ai/responses/{response_id}` for
// assistant output.
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

// Create a new conversation thread.
//
// Atomically creates a new thread and submits the first user turn. The response
// contains a `response_id` that should be polled via
// `GET /omni-ai/responses/{response_id}` for assistant output.
//
// Two creation modes are supported:
//
//   - **instant** — provide `text` with a natural-language prompt.
//   - **deep_insights** — provide a `target` ticker and optional `thesis` for
//     long-form research.
func (r *V1OmniAIThreadService) NewThread(ctx context.Context, body V1OmniAIThreadNewThreadParams, opts ...option.RequestOption) (res *V1OmniAIThreadNewThreadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/omni-ai/threads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List finalized messages in a thread.
//
// Returns the latest page of **finalized** messages by default, with messages
// within each page ordered chronologically. Messages from in-progress assistant
// turns are excluded — use `GET /omni-ai/threads/{thread_id}/response` or
// `GET /omni-ai/responses/{response_id}` for live output.
//
// If the last finalized message has role `USER`, an active response likely exists
// and should be polled separately.
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

// Get a specific thread.
//
// Returns metadata (title, timestamps) for a single thread. Does not include
// messages — use `GET /omni-ai/threads/{thread_id}/messages` for conversation
// history.
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

// Get the active response for a thread.
//
// Convenience endpoint to look up the currently active response for a thread
// without knowing the `response_id`. Useful when reloading a thread whose last
// finalized message is a `USER` message — this indicates an assistant turn is
// likely in progress.
//
// Returns **404** if no active response exists (the thread is idle).
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

// List conversation threads.
//
// Returns thread metadata ordered by most recently created first. Use `page_size`
// and `page_token` for pagination. Thread objects contain only metadata (title,
// timestamps) — use the messages endpoint for conversation history.
func (r *V1OmniAIThreadService) GetThreads(ctx context.Context, query V1OmniAIThreadGetThreadsParams, opts ...option.RequestOption) (res *V1OmniAIThreadGetThreadsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/omni-ai/threads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
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
	JSON      struct {
		OfContentPartCustomPayloadPayload respjson.Field
		ChartID                           respjson.Field
		ActionButtons                     respjson.Field
		DataChart                         respjson.Field
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

// Thread metadata returned by list/get thread endpoints.
type Thread struct {
	ID        string `json:"id" api:"required" format:"uuid"`
	CreatedAt string `json:"created_at" api:"required"`
	Title     string `json:"title" api:"required"`
	UpdatedAt string `json:"updated_at" api:"required"`
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
	// Thread metadata returned by list/get thread endpoints.
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
	AccountID int64  `json:"account_id" api:"required"`
	Text      string `json:"text" api:"required"`
	// Any of "PREFILL_ORDER", "OPEN_CHART", "OPEN_SCREENER",
	// "OPEN_ENTITLEMENT_CONSENT".
	Capabilities []string `json:"capabilities,omitzero"`
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
	AccountID int64 `json:"account_id" api:"required"`
	// Thread creation mode.
	//
	// Any of "instant", "deep_insights".
	Type   V1OmniAIThreadNewThreadParamsType `json:"type,omitzero" api:"required"`
	Text   param.Opt[string]                 `json:"text,omitzero"`
	Thesis param.Opt[string]                 `json:"thesis,omitzero"`
	// Deep-insights target payload.
	Target V1OmniAIThreadNewThreadParamsTarget `json:"target,omitzero"`
	// Any of "PREFILL_ORDER", "OPEN_CHART", "OPEN_SCREENER",
	// "OPEN_ENTITLEMENT_CONSENT".
	Capabilities []string `json:"capabilities,omitzero"`
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
	// Account ID for the request
	AccountID int64 `query:"account_id" api:"required" json:"-"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1OmniAIThreadGetThreadByIDParams struct {
	// Account ID for the request
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1OmniAIThreadGetThreadByIDParams]'s query parameters as
// `url.Values`.
func (r V1OmniAIThreadGetThreadByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1OmniAIThreadGetThreadResponseParams struct {
	// Account ID for the request
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1OmniAIThreadGetThreadResponseParams]'s query parameters
// as `url.Values`.
func (r V1OmniAIThreadGetThreadResponseParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1OmniAIThreadGetThreadsParams struct {
	// Account ID for the request
	AccountID int64 `query:"account_id" api:"required" json:"-"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
