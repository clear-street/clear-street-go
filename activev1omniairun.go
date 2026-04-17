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

// AI assistant for conversational trading interactions.
//
// ActiveV1OmniAIRunService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1OmniAIRunService] method instead.
type ActiveV1OmniAIRunService struct {
	options []option.RequestOption
}

// NewActiveV1OmniAIRunService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1OmniAIRunService(opts ...option.RequestOption) (r ActiveV1OmniAIRunService) {
	r = ActiveV1OmniAIRunService{}
	r.options = opts
	return
}

// Cancel a running assistant run.
func (r *ActiveV1OmniAIRunService) CancelRun(ctx context.Context, runID string, body ActiveV1OmniAIRunCancelRunParams, opts ...option.RequestOption) (res *ActiveV1OmniAIRunCancelRunResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/omni-ai/runs/%s", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Get run status and events.
//
// Poll for the current status of a run and any new events since the last poll.
func (r *ActiveV1OmniAIRunService) GetRun(ctx context.Context, runID string, query ActiveV1OmniAIRunGetRunParams, opts ...option.RequestOption) (res *ActiveV1OmniAIRunGetRunResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/omni-ai/runs/%s", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Start a new assistant run.
//
// Begins an agentic conversation run. If thread_id is provided, continues an
// existing conversation; otherwise creates a new thread.
func (r *ActiveV1OmniAIRunService) StartRun(ctx context.Context, body ActiveV1OmniAIRunStartRunParams, opts ...option.RequestOption) (res *ActiveV1OmniAIRunStartRunResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/omni-ai/runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ActiveV1OmniAIRunCancelRunResponse struct {
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
func (r ActiveV1OmniAIRunCancelRunResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1OmniAIRunCancelRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1OmniAIRunGetRunResponse struct {
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
func (r ActiveV1OmniAIRunGetRunResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1OmniAIRunGetRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1OmniAIRunStartRunResponse struct {
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
func (r ActiveV1OmniAIRunStartRunResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1OmniAIRunStartRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1OmniAIRunCancelRunParams struct {
	// Account ID for the request
	AccountID string `json:"account_id" api:"required"`
	// Reason for cancellation
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r ActiveV1OmniAIRunCancelRunParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1OmniAIRunCancelRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1OmniAIRunCancelRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1OmniAIRunGetRunParams struct {
	// Account ID for the request
	AccountID string `query:"account_id" api:"required" json:"-"`
	// Maximum events to return
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Page token for incremental polling
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1OmniAIRunGetRunParams]'s query parameters as
// `url.Values`.
func (r ActiveV1OmniAIRunGetRunParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActiveV1OmniAIRunStartRunParams struct {
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

func (r ActiveV1OmniAIRunStartRunParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1OmniAIRunStartRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1OmniAIRunStartRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
