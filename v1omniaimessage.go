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
// V1OmniAIMessageService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1OmniAIMessageService] method instead.
type V1OmniAIMessageService struct {
	options []option.RequestOption
}

// NewV1OmniAIMessageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1OmniAIMessageService(opts ...option.RequestOption) (r V1OmniAIMessageService) {
	r = V1OmniAIMessageService{}
	r.options = opts
	return
}

// Create feedback on a finalized assistant message.
//
// Attaches a score and optional comment to a finalized assistant message. Feedback
// is only valid for messages with role `ASSISTANT` that have reached a terminal
// outcome.
func (r *V1OmniAIMessageService) Feedback(ctx context.Context, messageID string, body V1OmniAIMessageFeedbackParams, opts ...option.RequestOption) (res *V1OmniAIMessageFeedbackResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/omni-ai/messages/%s/feedback", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a finalized message by ID.
//
// Returns a single finalized message. Returns **404** if the message belongs to an
// in-progress assistant turn (use the response endpoint for live output). Once the
// turn completes, the message becomes available here.
func (r *V1OmniAIMessageService) GetMessage(ctx context.Context, messageID string, query V1OmniAIMessageGetMessageParams, opts ...option.RequestOption) (res *V1OmniAIMessageGetMessageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/omni-ai/messages/%s", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type V1OmniAIMessageFeedbackResponse struct {
	Data CreateFeedbackResponse `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1OmniAIMessageFeedbackResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIMessageFeedbackResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIMessageGetMessageResponse struct {
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
func (r V1OmniAIMessageGetMessageResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIMessageGetMessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIMessageFeedbackParams struct {
	// Account ID for the request
	AccountID int64 `json:"account_id" api:"required"`
	// Feedback score (-1, 0, +1 or 1-5)
	Score int64 `json:"score" api:"required"`
	// Optional feedback comment
	Comment param.Opt[string] `json:"comment,omitzero"`
	// Optional metadata
	Metadata any `json:"metadata,omitzero"`
	paramObj
}

func (r V1OmniAIMessageFeedbackParams) MarshalJSON() (data []byte, err error) {
	type shadow V1OmniAIMessageFeedbackParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1OmniAIMessageFeedbackParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIMessageGetMessageParams struct {
	// Account ID for the request
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1OmniAIMessageGetMessageParams]'s query parameters as
// `url.Values`.
func (r V1OmniAIMessageGetMessageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
