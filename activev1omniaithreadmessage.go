// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
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
// messages from thread history. Every endpoint requires an explicit account_id.
//
// ActiveV1OmniAIThreadMessageService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1OmniAIThreadMessageService] method instead.
type ActiveV1OmniAIThreadMessageService struct {
	options []option.RequestOption
}

// NewActiveV1OmniAIThreadMessageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1OmniAIThreadMessageService(opts ...option.RequestOption) (r ActiveV1OmniAIThreadMessageService) {
	r = ActiveV1OmniAIThreadMessageService{}
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
func (r *ActiveV1OmniAIThreadMessageService) NewMessage(ctx context.Context, threadID string, body ActiveV1OmniAIThreadMessageNewMessageParams, opts ...option.RequestOption) (res *ActiveV1OmniAIThreadMessageNewMessageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/omni-ai/threads/%s/messages", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List finalized messages in a thread.
//
// Returns **finalized** messages in chronological order. Messages from in-progress
// assistant turns are excluded — use `GET /omni-ai/threads/{thread_id}/response`
// or `GET /omni-ai/responses/{response_id}` for live output.
//
// If the last finalized message has role `USER`, an active response likely exists
// and should be polled separately.
func (r *ActiveV1OmniAIThreadMessageService) ListMessages(ctx context.Context, threadID string, query ActiveV1OmniAIThreadMessageListMessagesParams, opts ...option.RequestOption) (res *ActiveV1OmniAIThreadMessageListMessagesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/omni-ai/threads/%s/messages", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ActiveV1OmniAIThreadMessageNewMessageResponse struct {
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
func (r ActiveV1OmniAIThreadMessageNewMessageResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1OmniAIThreadMessageNewMessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1OmniAIThreadMessageListMessagesResponse struct {
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
func (r ActiveV1OmniAIThreadMessageListMessagesResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1OmniAIThreadMessageListMessagesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1OmniAIThreadMessageNewMessageParams struct {
	AccountID int64  `json:"account_id" api:"required"`
	Text      string `json:"text" api:"required"`
	// Any of "PREFILL_ORDER", "OPEN_CHART", "OPEN_SCREENER".
	Capabilities []string `json:"capabilities,omitzero"`
	paramObj
}

func (r ActiveV1OmniAIThreadMessageNewMessageParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1OmniAIThreadMessageNewMessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1OmniAIThreadMessageNewMessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1OmniAIThreadMessageListMessagesParams struct {
	// Account ID for the request
	AccountID int64            `query:"account_id" api:"required" json:"-"`
	PageSize  param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1OmniAIThreadMessageListMessagesParams]'s query
// parameters as `url.Values`.
func (r ActiveV1OmniAIThreadMessageListMessagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
