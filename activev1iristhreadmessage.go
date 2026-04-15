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

// Deprecated /iris/_ routes. Use /omni-ai/_ instead.
//
// ActiveV1IrisThreadMessageService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1IrisThreadMessageService] method instead.
type ActiveV1IrisThreadMessageService struct {
	options []option.RequestOption
}

// NewActiveV1IrisThreadMessageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1IrisThreadMessageService(opts ...option.RequestOption) (r ActiveV1IrisThreadMessageService) {
	r = ActiveV1IrisThreadMessageService{}
	r.options = opts
	return
}

// **Deprecated**: Use `GET /omni-ai/threads/{thread_id}/messages` instead.
//
// Deprecated: deprecated
func (r *ActiveV1IrisThreadMessageService) ListMessagesDeprecated(ctx context.Context, threadID string, query ActiveV1IrisThreadMessageListMessagesDeprecatedParams, opts ...option.RequestOption) (res *ActiveV1IrisThreadMessageListMessagesDeprecatedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/iris/threads/%s/messages", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ActiveV1IrisThreadMessageListMessagesDeprecatedResponse struct {
	Data ListMessagesResponse `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1IrisThreadMessageListMessagesDeprecatedResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1IrisThreadMessageListMessagesDeprecatedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisThreadMessageListMessagesDeprecatedParams struct {
	// Account ID for the request
	AccountID string `query:"account_id" api:"required" json:"-"`
	// Return messages after this sequence number
	AfterSeq param.Opt[int64] `query:"after_seq,omitzero" json:"-"`
	// Maximum messages to return
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Page token for pagination
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1IrisThreadMessageListMessagesDeprecatedParams]'s
// query parameters as `url.Values`.
func (r ActiveV1IrisThreadMessageListMessagesDeprecatedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
