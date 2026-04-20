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
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Thread-centric AI assistant for conversational trading. Create threads to start
// conversations, poll response objects for in-progress output, and read finalized
// messages from thread history. Every endpoint requires an explicit account_id.
//
// ActiveV1OmniAIThreadResponseService contains methods and other services that
// help with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1OmniAIThreadResponseService] method instead.
type ActiveV1OmniAIThreadResponseService struct {
	options []option.RequestOption
}

// NewActiveV1OmniAIThreadResponseService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1OmniAIThreadResponseService(opts ...option.RequestOption) (r ActiveV1OmniAIThreadResponseService) {
	r = ActiveV1OmniAIThreadResponseService{}
	r.options = opts
	return
}

// Get the active response for a thread.
//
// Convenience endpoint to look up the currently active response for a thread
// without knowing the `response_id`. Useful when reloading a thread whose last
// finalized message is a `USER` message — this indicates an assistant turn is
// likely in progress.
//
// Returns **404** if no active response exists (the thread is idle).
func (r *ActiveV1OmniAIThreadResponseService) GetThreadResponse(ctx context.Context, threadID string, query ActiveV1OmniAIThreadResponseGetThreadResponseParams, opts ...option.RequestOption) (res *ActiveV1OmniAIThreadResponseGetThreadResponseResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/omni-ai/threads/%s/response", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ActiveV1OmniAIThreadResponseGetThreadResponseResponse struct {
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
func (r ActiveV1OmniAIThreadResponseGetThreadResponseResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1OmniAIThreadResponseGetThreadResponseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1OmniAIThreadResponseGetThreadResponseParams struct {
	// Account ID for the request
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1OmniAIThreadResponseGetThreadResponseParams]'s
// query parameters as `url.Values`.
func (r ActiveV1OmniAIThreadResponseGetThreadResponseParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
