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
// ActiveV1OmniAIThreadService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1OmniAIThreadService] method instead.
type ActiveV1OmniAIThreadService struct {
	options []option.RequestOption
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Every endpoint requires an explicit account_id.
	Messages ActiveV1OmniAIThreadMessageService
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Every endpoint requires an explicit account_id.
	Response ActiveV1OmniAIThreadResponseService
}

// NewActiveV1OmniAIThreadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1OmniAIThreadService(opts ...option.RequestOption) (r ActiveV1OmniAIThreadService) {
	r = ActiveV1OmniAIThreadService{}
	r.options = opts
	r.Messages = NewActiveV1OmniAIThreadMessageService(opts...)
	r.Response = NewActiveV1OmniAIThreadResponseService(opts...)
	return
}

// Atomically creates a new thread and submits the first user turn. The response
// contains a `response_id` that should be polled via
// `GET /omni-ai/responses/{response_id}` for assistant output.
//
// Two creation modes are supported:
//
//   - **instant** — provide `text` with a natural-language prompt.
//   - **deep_insights** — provide a `target` ticker and optional `thesis` for
//     long-form research.
func (r *ActiveV1OmniAIThreadService) NewThread(ctx context.Context, body ActiveV1OmniAIThreadNewThreadParams, opts ...option.RequestOption) (res *ActiveV1OmniAIThreadNewThreadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/omni-ai/threads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns metadata (title, timestamps) for a single thread. Does not include
// messages — use `GET /omni-ai/threads/{thread_id}/messages` for conversation
// history.
func (r *ActiveV1OmniAIThreadService) GetThread(ctx context.Context, threadID string, query ActiveV1OmniAIThreadGetThreadParams, opts ...option.RequestOption) (res *ActiveV1OmniAIThreadGetThreadResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/omni-ai/threads/%s", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns thread metadata ordered by most recently created first. Use `page_size`
// and `page_token` for pagination. Thread objects contain only metadata (title,
// timestamps) — use the messages endpoint for conversation history.
func (r *ActiveV1OmniAIThreadService) ListThreads(ctx context.Context, query ActiveV1OmniAIThreadListThreadsParams, opts ...option.RequestOption) (res *ActiveV1OmniAIThreadListThreadsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/omni-ai/threads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ActiveV1OmniAIThreadNewThreadResponse struct {
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
func (r ActiveV1OmniAIThreadNewThreadResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1OmniAIThreadNewThreadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1OmniAIThreadGetThreadResponse struct {
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
func (r ActiveV1OmniAIThreadGetThreadResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1OmniAIThreadGetThreadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1OmniAIThreadListThreadsResponse struct {
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
func (r ActiveV1OmniAIThreadListThreadsResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1OmniAIThreadListThreadsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1OmniAIThreadNewThreadParams struct {
	AccountID int64 `json:"account_id" api:"required"`
	// Thread creation mode.
	//
	// Any of "instant", "deep_insights".
	Type   ActiveV1OmniAIThreadNewThreadParamsType `json:"type,omitzero" api:"required"`
	Text   param.Opt[string]                       `json:"text,omitzero"`
	Thesis param.Opt[string]                       `json:"thesis,omitzero"`
	// Deep-insights target payload.
	Target ActiveV1OmniAIThreadNewThreadParamsTarget `json:"target,omitzero"`
	// Any of "PREFILL_ORDER", "OPEN_CHART", "OPEN_SCREENER".
	Capabilities []string `json:"capabilities,omitzero"`
	paramObj
}

func (r ActiveV1OmniAIThreadNewThreadParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1OmniAIThreadNewThreadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1OmniAIThreadNewThreadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Thread creation mode.
type ActiveV1OmniAIThreadNewThreadParamsType string

const (
	ActiveV1OmniAIThreadNewThreadParamsTypeInstant      ActiveV1OmniAIThreadNewThreadParamsType = "instant"
	ActiveV1OmniAIThreadNewThreadParamsTypeDeepInsights ActiveV1OmniAIThreadNewThreadParamsType = "deep_insights"
)

// Deep-insights target payload.
//
// The properties Ticker, Type are required.
type ActiveV1OmniAIThreadNewThreadParamsTarget struct {
	Ticker string `json:"ticker" api:"required"`
	// Deep-insights target type. Launch supports ticker-only.
	//
	// Any of "ticker".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ActiveV1OmniAIThreadNewThreadParamsTarget) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1OmniAIThreadNewThreadParamsTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1OmniAIThreadNewThreadParamsTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActiveV1OmniAIThreadNewThreadParamsTarget](
		"type", "ticker",
	)
}

type ActiveV1OmniAIThreadGetThreadParams struct {
	// Account ID for the request
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1OmniAIThreadGetThreadParams]'s query parameters as
// `url.Values`.
func (r ActiveV1OmniAIThreadGetThreadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActiveV1OmniAIThreadListThreadsParams struct {
	// Account ID for the request
	AccountID int64            `query:"account_id" api:"required" json:"-"`
	PageSize  param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1OmniAIThreadListThreadsParams]'s query parameters
// as `url.Values`.
func (r ActiveV1OmniAIThreadListThreadsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
