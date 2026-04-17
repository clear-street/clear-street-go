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
// ActiveV1OmniAIResponseService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1OmniAIResponseService] method instead.
type ActiveV1OmniAIResponseService struct {
	options []option.RequestOption
}

// NewActiveV1OmniAIResponseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1OmniAIResponseService(opts ...option.RequestOption) (r ActiveV1OmniAIResponseService) {
	r = ActiveV1OmniAIResponseService{}
	r.options = opts
	return
}

// Requests cancellation of a queued or running response. If the response has
// already reached a terminal status, this is an idempotent success. A canceled
// turn still produces a final assistant message with outcome `canceled` in the
// thread history.
func (r *ActiveV1OmniAIResponseService) CancelResponse(ctx context.Context, responseID string, body ActiveV1OmniAIResponseCancelResponseParams, opts ...option.RequestOption) (res *ActiveV1OmniAIResponseCancelResponseResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/omni-ai/responses/%s", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Returns the current snapshot of an in-progress or completed response. While the
// status is `queued` or `running`, the content may be partial and may include
// `thinking` parts. Poll this endpoint periodically until the status reaches a
// terminal value (`succeeded`, `failed`, or `canceled`).
//
// Once terminal, the finalized assistant message is available in thread history
// via `GET /omni-ai/threads/{thread_id}/messages`.
func (r *ActiveV1OmniAIResponseService) GetResponse(ctx context.Context, responseID string, query ActiveV1OmniAIResponseGetResponseParams, opts ...option.RequestOption) (res *ActiveV1OmniAIResponseGetResponseResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/omni-ai/responses/%s", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ActiveV1OmniAIResponseCancelResponseResponse struct {
	Data CancelResponsePayload `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1OmniAIResponseCancelResponseResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1OmniAIResponseCancelResponseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1OmniAIResponseGetResponseResponse struct {
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
func (r ActiveV1OmniAIResponseGetResponseResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1OmniAIResponseGetResponseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1OmniAIResponseCancelResponseParams struct {
	// Account ID for the request
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1OmniAIResponseCancelResponseParams]'s query
// parameters as `url.Values`.
func (r ActiveV1OmniAIResponseCancelResponseParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActiveV1OmniAIResponseGetResponseParams struct {
	// Account ID for the request
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1OmniAIResponseGetResponseParams]'s query parameters
// as `url.Values`.
func (r ActiveV1OmniAIResponseGetResponseParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
