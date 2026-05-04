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
// messages from thread history. Thread/message/response endpoints require an
// explicit account_id. Entitlement endpoints are caller-scoped and use
// trading_account_ids.
//
// V1OmniAIThreadService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1OmniAIThreadService] method instead.
type V1OmniAIThreadService struct {
	options []option.RequestOption
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Thread/message/response endpoints require an
	// explicit account_id. Entitlement endpoints are caller-scoped and use
	// trading_account_ids.
	Messages V1OmniAIThreadMessageService
}

// NewV1OmniAIThreadService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1OmniAIThreadService(opts ...option.RequestOption) (r V1OmniAIThreadService) {
	r = V1OmniAIThreadService{}
	r.options = opts
	r.Messages = NewV1OmniAIThreadMessageService(opts...)
	return
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

// Get a specific thread.
//
// Returns metadata (title, timestamps) for a single thread. Does not include
// messages — use `GET /omni-ai/threads/{thread_id}/messages` for conversation
// history.
func (r *V1OmniAIThreadService) GetThread(ctx context.Context, threadID string, query V1OmniAIThreadGetThreadParams, opts ...option.RequestOption) (res *V1OmniAIThreadGetThreadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/omni-ai/threads/%s", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List conversation threads.
//
// Returns thread metadata ordered by most recently created first. Use `page_size`
// and `page_token` for pagination. Thread objects contain only metadata (title,
// timestamps) — use the messages endpoint for conversation history.
func (r *V1OmniAIThreadService) ListThreads(ctx context.Context, query V1OmniAIThreadListThreadsParams, opts ...option.RequestOption) (res *V1OmniAIThreadListThreadsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/omni-ai/threads"
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
func (r *V1OmniAIThreadService) Response(ctx context.Context, threadID string, query V1OmniAIThreadResponseParams, opts ...option.RequestOption) (res *V1OmniAIThreadResponseResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/omni-ai/threads/%s/response", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
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

type V1OmniAIThreadGetThreadResponse struct {
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
func (r V1OmniAIThreadGetThreadResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIThreadGetThreadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIThreadListThreadsResponse struct {
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
func (r V1OmniAIThreadListThreadsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIThreadListThreadsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIThreadResponseResponse struct {
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
func (r V1OmniAIThreadResponseResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIThreadResponseResponse) UnmarshalJSON(data []byte) error {
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

type V1OmniAIThreadGetThreadParams struct {
	// Account ID for the request
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1OmniAIThreadGetThreadParams]'s query parameters as
// `url.Values`.
func (r V1OmniAIThreadGetThreadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1OmniAIThreadListThreadsParams struct {
	// Account ID for the request
	AccountID int64            `query:"account_id" api:"required" json:"-"`
	PageSize  param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	paramObj
}

// URLQuery serializes [V1OmniAIThreadListThreadsParams]'s query parameters as
// `url.Values`.
func (r V1OmniAIThreadListThreadsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1OmniAIThreadResponseParams struct {
	// Account ID for the request
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1OmniAIThreadResponseParams]'s query parameters as
// `url.Values`.
func (r V1OmniAIThreadResponseParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
