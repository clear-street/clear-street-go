// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/apiquery"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/param"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// AI assistant for conversational trading interactions.
//
// ActiveV1IrisThreadService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1IrisThreadService] method instead.
type ActiveV1IrisThreadService struct {
	Options []option.RequestOption
	// AI assistant for conversational trading interactions.
	Messages ActiveV1IrisThreadMessageService
}

// NewActiveV1IrisThreadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1IrisThreadService(opts ...option.RequestOption) (r ActiveV1IrisThreadService) {
	r = ActiveV1IrisThreadService{}
	r.Options = opts
	r.Messages = NewActiveV1IrisThreadMessageService(opts...)
	return
}

// Get a specific thread.
func (r *ActiveV1IrisThreadService) GetThread(ctx context.Context, threadID string, query ActiveV1IrisThreadGetThreadParams, opts ...option.RequestOption) (res *ActiveV1IrisThreadGetThreadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/iris/threads/%s", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves threads for the authenticated user.
func (r *ActiveV1IrisThreadService) ListThreads(ctx context.Context, query ActiveV1IrisThreadListThreadsParams, opts ...option.RequestOption) (res *ActiveV1IrisThreadListThreadsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/iris/threads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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

type ActiveV1IrisThreadGetThreadResponse struct {
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
func (r ActiveV1IrisThreadGetThreadResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1IrisThreadGetThreadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisThreadListThreadsResponse struct {
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
func (r ActiveV1IrisThreadListThreadsResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1IrisThreadListThreadsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisThreadGetThreadParams struct {
	// Account ID for the request
	AccountID string `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1IrisThreadGetThreadParams]'s query parameters as
// `url.Values`.
func (r ActiveV1IrisThreadGetThreadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActiveV1IrisThreadListThreadsParams struct {
	// Account ID for the request
	AccountID string `query:"account_id" api:"required" json:"-"`
	// Maximum threads to return
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Page token for pagination
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1IrisThreadListThreadsParams]'s query parameters as
// `url.Values`.
func (r ActiveV1IrisThreadListThreadsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
