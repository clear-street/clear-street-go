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

// ActiveV1AssistantPromptService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1AssistantPromptService] method instead.
type ActiveV1AssistantPromptService struct {
	Options []option.RequestOption
}

// NewActiveV1AssistantPromptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1AssistantPromptService(opts ...option.RequestOption) (r ActiveV1AssistantPromptService) {
	r = ActiveV1AssistantPromptService{}
	r.Options = opts
	return
}

// Retrieve the status and outputs of a prompt workflow by ID.
func (r *ActiveV1AssistantPromptService) GetPromptResult(ctx context.Context, id string, query ActiveV1AssistantPromptGetPromptResultParams, opts ...option.RequestOption) (res *ActiveV1AssistantPromptGetPromptResultResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/assistant/prompts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Forwards an arbitrary JSON payload to an Iris prompt identified by `slug` and
// returns a handle that can be used to poll for results.
func (r *ActiveV1AssistantPromptService) RunPrompt(ctx context.Context, body ActiveV1AssistantPromptRunPromptParams, opts ...option.RequestOption) (res *ActiveV1AssistantPromptRunPromptResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/assistant/prompts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PromptResult struct {
	PromptID string `json:"prompt_id" api:"required"`
	Response string `json:"response" api:"required"`
	// Any of "RUNNING", "SUCCESS", "FAILED", "UNSPECIFIED", "UNKNOWN".
	Status  PromptStatus                  `json:"status" api:"required"`
	Error   string                        `json:"error" api:"nullable"`
	Outputs map[string]PromptResultOutput `json:"outputs"`
	Raw     any                           `json:"raw"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PromptID    respjson.Field
		Response    respjson.Field
		Status      respjson.Field
		Error       respjson.Field
		Outputs     respjson.Field
		Raw         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptResult) RawJSON() string { return r.JSON.raw }
func (r *PromptResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptResultOutput struct {
	IsOutputNode    bool   `json:"is_output_node" api:"required"`
	Status          string `json:"status" api:"required"`
	DisplayValue    any    `json:"display_value"`
	ErrorMessage    string `json:"error_message" api:"nullable"`
	RawErrorMessage any    `json:"raw_error_message"`
	Value           any    `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsOutputNode    respjson.Field
		Status          respjson.Field
		DisplayValue    respjson.Field
		ErrorMessage    respjson.Field
		RawErrorMessage respjson.Field
		Value           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptResultOutput) RawJSON() string { return r.JSON.raw }
func (r *PromptResultOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptStatus string

const (
	PromptStatusRunning     PromptStatus = "RUNNING"
	PromptStatusSuccess     PromptStatus = "SUCCESS"
	PromptStatusFailed      PromptStatus = "FAILED"
	PromptStatusUnspecified PromptStatus = "UNSPECIFIED"
	PromptStatusUnknown     PromptStatus = "UNKNOWN"
)

type RunPromptResponse struct {
	RequestID string `json:"request_id" api:"required"`
	Response  string `json:"response" api:"required"`
	// Any of "RUNNING", "SUCCESS", "FAILED", "UNSPECIFIED", "UNKNOWN".
	Status PromptStatus `json:"status" api:"required"`
	Error  string       `json:"error" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		Response    respjson.Field
		Status      respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RunPromptResponse) RawJSON() string { return r.JSON.raw }
func (r *RunPromptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AssistantPromptGetPromptResultResponse struct {
	Data PromptResult `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1AssistantPromptGetPromptResultResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AssistantPromptGetPromptResultResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AssistantPromptRunPromptResponse struct {
	Data RunPromptResponse `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1AssistantPromptRunPromptResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AssistantPromptRunPromptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AssistantPromptGetPromptResultParams struct {
	// When true, return intermediate outputs for all nodes in the workflow.
	ReturnAllOutputs param.Opt[bool] `query:"return_all_outputs,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1AssistantPromptGetPromptResultParams]'s query
// parameters as `url.Values`.
func (r ActiveV1AssistantPromptGetPromptResultParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActiveV1AssistantPromptRunPromptParams struct {
	// JSON payload forwarded to the prompt workflow.
	Body any `json:"body,omitzero" api:"required"`
	// Unique slug identifying the prompt workflow to execute.
	Slug     string            `json:"slug" api:"required"`
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r ActiveV1AssistantPromptRunPromptParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1AssistantPromptRunPromptParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1AssistantPromptRunPromptParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
