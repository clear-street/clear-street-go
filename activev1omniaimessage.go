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
// ActiveV1OmniAIMessageService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1OmniAIMessageService] method instead.
type ActiveV1OmniAIMessageService struct {
	options []option.RequestOption
	// Thread-centric AI assistant for conversational trading. Create threads to start
	// conversations, poll response objects for in-progress output, and read finalized
	// messages from thread history. Every endpoint requires an explicit account_id.
	Feedback ActiveV1OmniAIMessageFeedbackService
}

// NewActiveV1OmniAIMessageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1OmniAIMessageService(opts ...option.RequestOption) (r ActiveV1OmniAIMessageService) {
	r = ActiveV1OmniAIMessageService{}
	r.options = opts
	r.Feedback = NewActiveV1OmniAIMessageFeedbackService(opts...)
	return
}

// Get a finalized message by ID.
//
// Returns a single finalized message. Returns **404** if the message belongs to an
// in-progress assistant turn (use the response endpoint for live output). Once the
// turn completes, the message becomes available here.
func (r *ActiveV1OmniAIMessageService) GetMessage(ctx context.Context, messageID string, query ActiveV1OmniAIMessageGetMessageParams, opts ...option.RequestOption) (res *ActiveV1OmniAIMessageGetMessageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/omni-ai/messages/%s", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ActiveV1OmniAIMessageGetMessageResponse struct {
	// Final immutable message.
	Data Message `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1OmniAIMessageGetMessageResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1OmniAIMessageGetMessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1OmniAIMessageGetMessageParams struct {
	// Account ID for the request
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1OmniAIMessageGetMessageParams]'s query parameters
// as `url.Values`.
func (r ActiveV1OmniAIMessageGetMessageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
