// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/clear-street/clear-street-go/internal/apijson"
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
// ActiveV1OmniAIMessageFeedbackService contains methods and other services that
// help with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1OmniAIMessageFeedbackService] method instead.
type ActiveV1OmniAIMessageFeedbackService struct {
	options []option.RequestOption
}

// NewActiveV1OmniAIMessageFeedbackService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1OmniAIMessageFeedbackService(opts ...option.RequestOption) (r ActiveV1OmniAIMessageFeedbackService) {
	r = ActiveV1OmniAIMessageFeedbackService{}
	r.options = opts
	return
}

// Attaches a score and optional comment to a finalized assistant message. Feedback
// is only valid for messages with role `ASSISTANT` that have reached a terminal
// outcome.
func (r *ActiveV1OmniAIMessageFeedbackService) NewFeedback(ctx context.Context, messageID string, body ActiveV1OmniAIMessageFeedbackNewFeedbackParams, opts ...option.RequestOption) (res *ActiveV1OmniAIMessageFeedbackNewFeedbackResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/omni-ai/messages/%s/feedback", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ActiveV1OmniAIMessageFeedbackNewFeedbackResponse struct {
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
func (r ActiveV1OmniAIMessageFeedbackNewFeedbackResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1OmniAIMessageFeedbackNewFeedbackResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1OmniAIMessageFeedbackNewFeedbackParams struct {
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

func (r ActiveV1OmniAIMessageFeedbackNewFeedbackParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1OmniAIMessageFeedbackNewFeedbackParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1OmniAIMessageFeedbackNewFeedbackParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
