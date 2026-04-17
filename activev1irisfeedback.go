// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"net/http"
	"slices"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/param"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Deprecated /iris/_ routes. Use /omni-ai/_ instead.
//
// ActiveV1IrisFeedbackService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1IrisFeedbackService] method instead.
type ActiveV1IrisFeedbackService struct {
	options []option.RequestOption
}

// NewActiveV1IrisFeedbackService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1IrisFeedbackService(opts ...option.RequestOption) (r ActiveV1IrisFeedbackService) {
	r = ActiveV1IrisFeedbackService{}
	r.options = opts
	return
}

// **Deprecated**: Use `POST /omni-ai/feedback` instead.
//
// Deprecated: deprecated
func (r *ActiveV1IrisFeedbackService) NewFeedbackDeprecated(ctx context.Context, body ActiveV1IrisFeedbackNewFeedbackDeprecatedParams, opts ...option.RequestOption) (res *ActiveV1IrisFeedbackNewFeedbackDeprecatedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/iris/feedback"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ActiveV1IrisFeedbackNewFeedbackDeprecatedResponse struct {
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
func (r ActiveV1IrisFeedbackNewFeedbackDeprecatedResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1IrisFeedbackNewFeedbackDeprecatedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisFeedbackNewFeedbackDeprecatedParams struct {
	// Account ID for the request
	AccountID string `json:"account_id" api:"required"`
	// Message to provide feedback on
	MessageID string `json:"message_id" api:"required" format:"uuid"`
	// Feedback score (-1, 0, +1 or 1-5)
	Score int64 `json:"score" api:"required"`
	// Thread containing the message
	ThreadID string `json:"thread_id" api:"required" format:"uuid"`
	// Optional feedback comment
	Comment param.Opt[string] `json:"comment,omitzero"`
	// Optional metadata
	Metadata any `json:"metadata,omitzero"`
	paramObj
}

func (r ActiveV1IrisFeedbackNewFeedbackDeprecatedParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1IrisFeedbackNewFeedbackDeprecatedParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1IrisFeedbackNewFeedbackDeprecatedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
