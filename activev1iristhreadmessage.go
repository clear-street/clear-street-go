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
// ActiveV1IrisThreadMessageService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1IrisThreadMessageService] method instead.
type ActiveV1IrisThreadMessageService struct {
	Options []option.RequestOption
}

// NewActiveV1IrisThreadMessageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1IrisThreadMessageService(opts ...option.RequestOption) (r ActiveV1IrisThreadMessageService) {
	r = ActiveV1IrisThreadMessageService{}
	r.Options = opts
	return
}

// List messages in a thread.
func (r *ActiveV1IrisThreadMessageService) ListMessages(ctx context.Context, threadID string, query ActiveV1IrisThreadMessageListMessagesParams, opts ...option.RequestOption) (res *ActiveV1IrisThreadMessageListMessagesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/iris/threads/%s/messages", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ListMessagesResponse struct {
	Messages      []Message `json:"messages" api:"required"`
	NextPageToken string    `json:"next_page_token" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Messages      respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListMessagesResponse) RawJSON() string { return r.JSON.raw }
func (r *ListMessagesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisThreadMessageListMessagesResponse struct {
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
func (r ActiveV1IrisThreadMessageListMessagesResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1IrisThreadMessageListMessagesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisThreadMessageListMessagesParams struct {
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

// URLQuery serializes [ActiveV1IrisThreadMessageListMessagesParams]'s query
// parameters as `url.Values`.
func (r ActiveV1IrisThreadMessageListMessagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
