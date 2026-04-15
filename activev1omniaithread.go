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

// AI assistant for conversational trading interactions.
//
// ActiveV1OmniAIThreadService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1OmniAIThreadService] method instead.
type ActiveV1OmniAIThreadService struct {
	options []option.RequestOption
	// AI assistant for conversational trading interactions.
	Messages ActiveV1OmniAIThreadMessageService
}

// NewActiveV1OmniAIThreadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1OmniAIThreadService(opts ...option.RequestOption) (r ActiveV1OmniAIThreadService) {
	r = ActiveV1OmniAIThreadService{}
	r.options = opts
	r.Messages = NewActiveV1OmniAIThreadMessageService(opts...)
	return
}

// Get a specific thread.
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

// Retrieves threads for the authenticated user.
func (r *ActiveV1OmniAIThreadService) ListThreads(ctx context.Context, query ActiveV1OmniAIThreadListThreadsParams, opts ...option.RequestOption) (res *ActiveV1OmniAIThreadListThreadsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/omni-ai/threads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ActiveV1OmniAIThreadGetThreadResponse struct {
	Data GetThreadResponse `json:"data" api:"required"`
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
	Data ListThreadsResponse `json:"data" api:"required"`
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

type ActiveV1OmniAIThreadGetThreadParams struct {
	// Account ID for the request
	AccountID string `query:"account_id" api:"required" json:"-"`
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
	AccountID string `query:"account_id" api:"required" json:"-"`
	// Maximum threads to return
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Page token for pagination
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"uuid" json:"-"`
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
