// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/param"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// ActiveV1IrisFeedbackService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1IrisFeedbackService] method instead.
type ActiveV1IrisFeedbackService struct {
	Options []option.RequestOption
}

// NewActiveV1IrisFeedbackService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1IrisFeedbackService(opts ...option.RequestOption) (r ActiveV1IrisFeedbackService) {
	r = ActiveV1IrisFeedbackService{}
	r.Options = opts
	return
}

// Submit user feedback (thumbs up/down, rating, comment) for an assistant message.
func (r *ActiveV1IrisFeedbackService) NewFeedback(ctx context.Context, body ActiveV1IrisFeedbackNewFeedbackParams, opts ...option.RequestOption) (res *ActiveV1IrisFeedbackNewFeedbackResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/iris/feedback"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CreateFeedbackResponse struct {
	CreatedAt  string `json:"created_at,required"`
	FeedbackID string `json:"feedback_id,nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		FeedbackID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateFeedbackResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateFeedbackResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisFeedbackNewFeedbackResponse struct {
	Data CreateFeedbackResponse `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1IrisFeedbackNewFeedbackResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1IrisFeedbackNewFeedbackResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisFeedbackNewFeedbackParams struct {
	// Account ID for the request
	AccountID string `json:"account_id,required"`
	// Message to provide feedback on
	MessageID string `json:"message_id,required" format:"uuid"`
	// Feedback score (-1, 0, +1 or 1-5)
	Score int64 `json:"score,required"`
	// Thread containing the message
	ThreadID string `json:"thread_id,required" format:"uuid"`
	// Optional feedback comment
	Comment param.Opt[string] `json:"comment,omitzero"`
	// Optional metadata
	Metadata any `json:"metadata,omitzero"`
	paramObj
}

func (r ActiveV1IrisFeedbackNewFeedbackParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1IrisFeedbackNewFeedbackParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1IrisFeedbackNewFeedbackParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
