// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"encoding/json"
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
// messages from thread history. Thread/message/response endpoints require an
// explicit account_id. Entitlement endpoints are caller-scoped and use
// account_ids.
//
// V1OmniAIResponseService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1OmniAIResponseService] method instead.
type V1OmniAIResponseService struct {
	options []option.RequestOption
}

// NewV1OmniAIResponseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1OmniAIResponseService(opts ...option.RequestOption) (r V1OmniAIResponseService) {
	r = V1OmniAIResponseService{}
	r.options = opts
	return
}

// Cancel a response.
//
// Requests cancellation of a queued or running response. If the response has
// already reached a terminal status, this is an idempotent success. A canceled
// turn still produces a final assistant message with outcome `canceled` in the
// thread history.
func (r *V1OmniAIResponseService) CancelResponse(ctx context.Context, responseID string, body V1OmniAIResponseCancelResponseParams, opts ...option.RequestOption) (res *V1OmniAIResponseCancelResponseResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/omni-ai/responses/%s", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Poll a response for assistant output.
//
// Returns the current snapshot of an in-progress or completed response. While the
// status is `queued` or `running`, the content may be partial and may include
// `thinking` parts. Poll this endpoint periodically until the status reaches a
// terminal value (`succeeded`, `failed`, or `canceled`).
//
// Once terminal, the finalized assistant message is available in thread history
// via `GET /omni-ai/threads/{thread_id}/messages`.
func (r *V1OmniAIResponseService) GetResponseByID(ctx context.Context, responseID string, query V1OmniAIResponseGetResponseByIDParams, opts ...option.RequestOption) (res *V1OmniAIResponseGetResponseByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/omni-ai/responses/%s", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type CancelResponsePayload struct {
	Canceled bool `json:"canceled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Canceled    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CancelResponsePayload) RawJSON() string { return r.JSON.raw }
func (r *CancelResponsePayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shared sanitized error payload.
type ErrorStatus struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	Details any    `json:"details" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Details     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ErrorStatus) RawJSON() string { return r.JSON.raw }
func (r *ErrorStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dynamic pollable response.
type Response struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Dynamic lifecycle status for a pollable response.
	//
	// Any of "queued", "running", "succeeded", "failed", "canceled".
	Status        ResponseStatus `json:"status" api:"required"`
	ThreadID      string         `json:"thread_id" api:"required" format:"uuid"`
	UserMessageID string         `json:"user_message_id" api:"required" format:"uuid"`
	// Dynamic response content container. May include thinking parts.
	Content ResponseContent `json:"content" api:"nullable"`
	// Shared sanitized error payload.
	Error           ErrorStatus `json:"error" api:"nullable"`
	OutputMessageID string      `json:"output_message_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Status          respjson.Field
		ThreadID        respjson.Field
		UserMessageID   respjson.Field
		Content         respjson.Field
		Error           respjson.Field
		OutputMessageID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Response) RawJSON() string { return r.JSON.raw }
func (r *Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dynamic response content container. May include thinking parts.
type ResponseContent struct {
	Parts []ResponseContentPartUnion `json:"parts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseContent) RawJSON() string { return r.JSON.raw }
func (r *ResponseContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseContentPartUnion contains all possible properties and values from
// [ResponseContentPartObject], [ResponseContentPartObject2],
// [ResponseContentPartObject3], [ResponseContentPartObject4],
// [ResponseContentPartObject5], [ResponseContentPartObject6].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseContentPartUnion struct {
	// This field is from variant [ResponseContentPartObject].
	Text string `json:"text"`
	Type string `json:"type"`
	// This field is from variant [ResponseContentPartObject2].
	Thoughts []string `json:"thoughts"`
	// This field is from variant [ResponseContentPartObject3].
	Action StructuredActionUnion `json:"action"`
	// This field is from variant [ResponseContentPartObject3].
	ActionID string `json:"action_id"`
	// This field is a union of [ChartPayload], [SuggestedActionsPayload], [any]
	Payload ResponseContentPartUnionPayload `json:"payload"`
	JSON    struct {
		Text     respjson.Field
		Type     respjson.Field
		Thoughts respjson.Field
		Action   respjson.Field
		ActionID respjson.Field
		Payload  respjson.Field
		raw      string
	} `json:"-"`
}

func (u ResponseContentPartUnion) AsResponseContentPartObject() (v ResponseContentPartObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseContentPartUnion) AsResponseContentPartObject2() (v ResponseContentPartObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseContentPartUnion) AsResponseContentPartObject3() (v ResponseContentPartObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseContentPartUnion) AsResponseContentPartObject4() (v ResponseContentPartObject4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseContentPartUnion) AsResponseContentPartObject5() (v ResponseContentPartObject5) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseContentPartUnion) AsResponseContentPartObject6() (v ResponseContentPartObject6) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseContentPartUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseContentPartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseContentPartUnionPayload is an implicit subunion of
// [ResponseContentPartUnion]. ResponseContentPartUnionPayload provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResponseContentPartUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfContentPartCustomPayloadPayload]
type ResponseContentPartUnionPayload struct {
	// This field will be present if the value is a [any] instead of an object.
	OfContentPartCustomPayloadPayload any `json:",inline"`
	// This field is from variant [ChartPayload].
	ChartID       string         `json:"chartId"`
	ActionButtons []ActionButton `json:"actionButtons"`
	// This field is from variant [ChartPayload].
	DataChart DataChart `json:"dataChart"`
	JSON      struct {
		OfContentPartCustomPayloadPayload respjson.Field
		ChartID                           respjson.Field
		ActionButtons                     respjson.Field
		DataChart                         respjson.Field
		raw                               string
	} `json:"-"`
}

func (r *ResponseContentPartUnionPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part.
type ResponseContentPartObject struct {
	// Any of "text".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartTextPayload
}

// Returns the unmodified JSON received from the API
func (r ResponseContentPartObject) RawJSON() string { return r.JSON.raw }
func (r *ResponseContentPartObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Thinking content part shown on dynamic response polling.
type ResponseContentPartObject2 struct {
	// Any of "thinking".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartThinkingPayload
}

// Returns the unmodified JSON received from the API
func (r ResponseContentPartObject2) RawJSON() string { return r.JSON.raw }
func (r *ResponseContentPartObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured action content part.
type ResponseContentPartObject3 struct {
	// Any of "structured_action".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartStructuredActionPayload
}

// Returns the unmodified JSON received from the API
func (r ResponseContentPartObject3) RawJSON() string { return r.JSON.raw }
func (r *ResponseContentPartObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chart payload content part.
type ResponseContentPartObject4 struct {
	// Any of "chart".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartChartPayload
}

// Returns the unmodified JSON received from the API
func (r ResponseContentPartObject4) RawJSON() string { return r.JSON.raw }
func (r *ResponseContentPartObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Suggested actions payload content part.
type ResponseContentPartObject5 struct {
	// Any of "suggested_actions".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartSuggestedActionsPayload
}

// Returns the unmodified JSON received from the API
func (r ResponseContentPartObject5) RawJSON() string { return r.JSON.raw }
func (r *ResponseContentPartObject5) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Escape-hatch custom payload content part.
type ResponseContentPartObject6 struct {
	// Any of "custom".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	ContentPartCustomPayload
}

// Returns the unmodified JSON received from the API
func (r ResponseContentPartObject6) RawJSON() string { return r.JSON.raw }
func (r *ResponseContentPartObject6) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dynamic lifecycle status for a pollable response.
type ResponseStatus string

const (
	ResponseStatusQueued    ResponseStatus = "queued"
	ResponseStatusRunning   ResponseStatus = "running"
	ResponseStatusSucceeded ResponseStatus = "succeeded"
	ResponseStatusFailed    ResponseStatus = "failed"
	ResponseStatusCanceled  ResponseStatus = "canceled"
)

type V1OmniAIResponseCancelResponseResponse struct {
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
func (r V1OmniAIResponseCancelResponseResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIResponseCancelResponseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIResponseGetResponseByIDResponse struct {
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
func (r V1OmniAIResponseGetResponseByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIResponseGetResponseByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIResponseCancelResponseParams struct {
	// Account ID for the request
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1OmniAIResponseCancelResponseParams]'s query parameters as
// `url.Values`.
func (r V1OmniAIResponseCancelResponseParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1OmniAIResponseGetResponseByIDParams struct {
	// Account ID for the request
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1OmniAIResponseGetResponseByIDParams]'s query parameters
// as `url.Values`.
func (r V1OmniAIResponseGetResponseByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
