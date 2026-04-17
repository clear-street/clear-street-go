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
// ActiveV1IrisRunService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1IrisRunService] method instead.
type ActiveV1IrisRunService struct {
	options []option.RequestOption
}

// NewActiveV1IrisRunService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActiveV1IrisRunService(opts ...option.RequestOption) (r ActiveV1IrisRunService) {
	r = ActiveV1IrisRunService{}
	r.options = opts
	return
}

// **Deprecated**: Use `DELETE /omni-ai/runs/{run_id}` instead.
//
// Deprecated: deprecated
func (r *ActiveV1IrisRunService) CancelRunDeprecated(ctx context.Context, runID string, body ActiveV1IrisRunCancelRunDeprecatedParams, opts ...option.RequestOption) (res *ActiveV1IrisRunCancelRunDeprecatedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/iris/runs/%s", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// **Deprecated**: Use `GET /omni-ai/runs/{run_id}` instead.
//
// Deprecated: deprecated
func (r *ActiveV1IrisRunService) GetRunDeprecated(ctx context.Context, runID string, query ActiveV1IrisRunGetRunDeprecatedParams, opts ...option.RequestOption) (res *ActiveV1IrisRunGetRunDeprecatedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/iris/runs/%s", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// **Deprecated**: Use `POST /omni-ai/runs` instead.
//
// Deprecated: deprecated
func (r *ActiveV1IrisRunService) StartRunDeprecated(ctx context.Context, body ActiveV1IrisRunStartRunDeprecatedParams, opts ...option.RequestOption) (res *ActiveV1IrisRunStartRunDeprecatedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/iris/runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ActiveV1IrisRunCancelRunDeprecatedResponse struct {
	Data CancelRunResponse `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1IrisRunCancelRunDeprecatedResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1IrisRunCancelRunDeprecatedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisRunGetRunDeprecatedResponse struct {
	Data GetRunResponse `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1IrisRunGetRunDeprecatedResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1IrisRunGetRunDeprecatedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisRunStartRunDeprecatedResponse struct {
	Data StartRunResponse `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1IrisRunStartRunDeprecatedResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1IrisRunStartRunDeprecatedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisRunCancelRunDeprecatedParams struct {
	// Account ID for the request
	AccountID string `json:"account_id" api:"required"`
	// Reason for cancellation
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r ActiveV1IrisRunCancelRunDeprecatedParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1IrisRunCancelRunDeprecatedParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1IrisRunCancelRunDeprecatedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisRunGetRunDeprecatedParams struct {
	// Account ID for the request
	AccountID string `query:"account_id" api:"required" json:"-"`
	// Maximum events to return
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Page token for incremental polling
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1IrisRunGetRunDeprecatedParams]'s query parameters
// as `url.Values`.
func (r ActiveV1IrisRunGetRunDeprecatedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActiveV1IrisRunStartRunDeprecatedParams struct {
	// Account ID for the request
	AccountID string `json:"account_id" api:"required"`
	// The user's natural language command
	CommandText string `json:"command_text" api:"required"`
	// Optional thread ID to continue an existing conversation
	ThreadID param.Opt[string] `json:"thread_id,omitzero" format:"uuid"`
	// Optional title for new threads
	ThreadTitle param.Opt[string] `json:"thread_title,omitzero"`
	// Capabilities for structured actions
	Capabilities []Capability `json:"capabilities,omitzero"`
	paramObj
}

func (r ActiveV1IrisRunStartRunDeprecatedParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1IrisRunStartRunDeprecatedParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1IrisRunStartRunDeprecatedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
