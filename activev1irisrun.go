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

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/apiquery"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/param"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// ActiveV1IrisRunService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1IrisRunService] method instead.
type ActiveV1IrisRunService struct {
	Options []option.RequestOption
}

// NewActiveV1IrisRunService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActiveV1IrisRunService(opts ...option.RequestOption) (r ActiveV1IrisRunService) {
	r = ActiveV1IrisRunService{}
	r.Options = opts
	return
}

// Cancel a running assistant run.
func (r *ActiveV1IrisRunService) CancelRun(ctx context.Context, runID string, body ActiveV1IrisRunCancelRunParams, opts ...option.RequestOption) (res *ActiveV1IrisRunCancelRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/iris/runs/%s", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Poll for the current status of a run and any new events since the last poll.
func (r *ActiveV1IrisRunService) GetRun(ctx context.Context, runID string, query ActiveV1IrisRunGetRunParams, opts ...option.RequestOption) (res *ActiveV1IrisRunGetRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/iris/runs/%s", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Begins an agentic conversation run. If thread_id is provided, continues an
// existing conversation; otherwise creates a new thread.
func (r *ActiveV1IrisRunService) StartRun(ctx context.Context, body ActiveV1IrisRunStartRunParams, opts ...option.RequestOption) (res *ActiveV1IrisRunStartRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/iris/runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CancelRunResponse struct {
	Canceled bool `json:"canceled,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Canceled    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CancelRunResponse) RawJSON() string { return r.JSON.raw }
func (r *CancelRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Capability allowlist for structured actions.
//
// Clients declare which capabilities they support when starting a run. Iris will
// only emit structured actions that match the declared capabilities.
type Capability string

const (
	CapabilityNavigate       Capability = "NAVIGATE"
	CapabilityOpenChatWindow Capability = "OPEN_CHAT_WINDOW"
	CapabilityPrefillOrder   Capability = "PREFILL_ORDER"
	CapabilityOpenChart      Capability = "OPEN_CHART"
	CapabilityOpenScreener   Capability = "OPEN_SCREENER"
)

// ContentPartUnion contains all possible properties and values from
// [ContentPartObject], [ContentPartObject], [ContentPartType].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ContentPartUnion struct {
	// This field is from variant [ContentPartObject].
	Text string `json:"text"`
	Type string `json:"type"`
	// This field is from variant [ContentPartObject].
	Action ContentPartObjectActionUnion `json:"action"`
	// This field is from variant [ContentPartObject].
	ActionID string `json:"action_id"`
	JSON     struct {
		Text     respjson.Field
		Type     respjson.Field
		Action   respjson.Field
		ActionID respjson.Field
		raw      string
	} `json:"-"`
}

func (u ContentPartUnion) AsContentPartObject() (v ContentPartObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentPartUnion) AsVariant2() (v ContentPartObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentPartUnion) AsContentPartType() (v ContentPartType) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContentPartUnion) RawJSON() string { return u.JSON.raw }

func (r *ContentPartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Plain text content
type ContentPartObject struct {
	Text string `json:"text,required"`
	// Any of "text".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentPartObject) RawJSON() string { return r.JSON.raw }
func (r *ContentPartObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom/extensible content
type ContentPartType struct {
	// Any of "custom".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentPartType) RawJSON() string { return r.JSON.raw }
func (r *ContentPartType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetRunResponse struct {
	Events        []any  `json:"events,required"`
	Run           Run    `json:"run,required"`
	NextPageToken string `json:"next_page_token,nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events        respjson.Field
		Run           respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetRunResponse) RawJSON() string { return r.JSON.raw }
func (r *GetRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message content containing text and structured action parts.
type MessageContent struct {
	Parts []ContentPartUnion `json:"parts,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageContent) RawJSON() string { return r.JSON.raw }
func (r *MessageContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageRole string

const (
	MessageRoleUnspecified MessageRole = "UNSPECIFIED"
	MessageRoleSystem      MessageRole = "SYSTEM"
	MessageRoleUser        MessageRole = "USER"
	MessageRoleAssistant   MessageRole = "ASSISTANT"
	MessageRoleTool        MessageRole = "TOOL"
)

type Run struct {
	CreatedAt string `json:"created_at,required"`
	Model     string `json:"model,required"`
	Provider  string `json:"provider,required"`
	// Any of "UNSPECIFIED", "QUEUED", "RUNNING", "SUCCEEDED", "FAILED", "CANCELED".
	Status       RunStatus    `json:"status,required"`
	ID           string       `json:"id,nullable" format:"uuid"`
	Capabilities []Capability `json:"capabilities"`
	EndedAt      string       `json:"ended_at,nullable"`
	Error        any          `json:"error,nullable"`
	Metadata     any          `json:"metadata,nullable"`
	Parameters   any          `json:"parameters,nullable"`
	StartedAt    string       `json:"started_at,nullable"`
	ThreadID     string       `json:"thread_id,nullable" format:"uuid"`
	Usage        any          `json:"usage,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		Model        respjson.Field
		Provider     respjson.Field
		Status       respjson.Field
		ID           respjson.Field
		Capabilities respjson.Field
		EndedAt      respjson.Field
		Error        respjson.Field
		Metadata     respjson.Field
		Parameters   respjson.Field
		StartedAt    respjson.Field
		ThreadID     respjson.Field
		Usage        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Run) RawJSON() string { return r.JSON.raw }
func (r *Run) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RunStatus string

const (
	RunStatusUnspecified RunStatus = "UNSPECIFIED"
	RunStatusQueued      RunStatus = "QUEUED"
	RunStatusRunning     RunStatus = "RUNNING"
	RunStatusSucceeded   RunStatus = "SUCCEEDED"
	RunStatusFailed      RunStatus = "FAILED"
	RunStatusCanceled    RunStatus = "CANCELED"
)

type StartRunResponse struct {
	Run         Run     `json:"run,required"`
	Thread      Thread  `json:"thread,required"`
	UserMessage Message `json:"user_message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Run         respjson.Field
		Thread      respjson.Field
		UserMessage respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StartRunResponse) RawJSON() string { return r.JSON.raw }
func (r *StartRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisRunCancelRunResponse struct {
	Data CancelRunResponse `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1IrisRunCancelRunResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1IrisRunCancelRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisRunGetRunResponse struct {
	Data GetRunResponse `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1IrisRunGetRunResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1IrisRunGetRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisRunStartRunResponse struct {
	Data StartRunResponse `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1IrisRunStartRunResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1IrisRunStartRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisRunCancelRunParams struct {
	// Account ID for the request
	AccountID string `json:"account_id,required"`
	// Reason for cancellation
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r ActiveV1IrisRunCancelRunParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1IrisRunCancelRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1IrisRunCancelRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1IrisRunGetRunParams struct {
	// Account ID for the request
	AccountID string `query:"account_id,required" json:"-"`
	// Maximum events to return
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Page token for incremental polling
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1IrisRunGetRunParams]'s query parameters as
// `url.Values`.
func (r ActiveV1IrisRunGetRunParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActiveV1IrisRunStartRunParams struct {
	// Account ID for the request
	AccountID string `json:"account_id,required"`
	// The user's natural language command
	CommandText string `json:"command_text,required"`
	// Optional model override
	Model param.Opt[string] `json:"model,omitzero"`
	// Optional LLM provider override
	Provider param.Opt[string] `json:"provider,omitzero"`
	// Optional thread ID to continue an existing conversation
	ThreadID param.Opt[string] `json:"thread_id,omitzero" format:"uuid"`
	// Optional title for new threads
	ThreadTitle param.Opt[string] `json:"thread_title,omitzero"`
	// Optional context for the agent
	Context any `json:"context,omitzero"`
	// Optional metadata
	Metadata any `json:"metadata,omitzero"`
	// Optional LLM parameters
	Parameters any `json:"parameters,omitzero"`
	// Capabilities for structured actions
	Capabilities []Capability `json:"capabilities,omitzero"`
	paramObj
}

func (r ActiveV1IrisRunStartRunParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1IrisRunStartRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1IrisRunStartRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
