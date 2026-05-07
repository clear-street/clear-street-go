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
	shimjson "github.com/clear-street/clear-street-go/internal/encoding/json"
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/param"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Submit and monitor option exercise, DNE, CEA, and cancel instructions.
//
// V1AccountPositionInstructionService contains methods and other services that
// help with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1AccountPositionInstructionService] method instead.
type V1AccountPositionInstructionService struct {
	options []option.RequestOption
}

// NewV1AccountPositionInstructionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1AccountPositionInstructionService(opts ...option.RequestOption) (r V1AccountPositionInstructionService) {
	r = V1AccountPositionInstructionService{}
	r.options = opts
	return
}

// Cancel an outstanding exercise / DNE / CEA instruction by its server- assigned
// `id`. Returns the updated instruction with status `CANCEL_REQUESTED`; the
// terminal `CANCELLED` / `CANCEL_FAILED` state arrives asynchronously via
// subsequent GETs.
func (r *V1AccountPositionInstructionService) CancelPositionInstruction(ctx context.Context, instructionID string, body V1AccountPositionInstructionCancelPositionInstructionParams, opts ...option.RequestOption) (res *V1AccountPositionInstructionCancelPositionInstructionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instructionID == "" {
		err = errors.New("missing required instruction_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%v/positions/instructions/%s", body.AccountID, instructionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns the current lifecycle state of exercise / DNE / CEA instructions for the
// account. Optionally filter by a specific instrument.
func (r *V1AccountPositionInstructionService) GetPositionInstructions(ctx context.Context, accountID int64, query V1AccountPositionInstructionGetPositionInstructionsParams, opts ...option.RequestOption) (res *V1AccountPositionInstructionGetPositionInstructionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/positions/instructions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Submit one or more option lifecycle instructions against the account. Each row
// is routed to `oems-csc` independently; per-row rejections are surfaced on the
// corresponding response entry without failing the batch.
func (r *V1AccountPositionInstructionService) SubmitPositionInstructions(ctx context.Context, accountID int64, body V1AccountPositionInstructionSubmitPositionInstructionsParams, opts ...option.RequestOption) (res *V1AccountPositionInstructionSubmitPositionInstructionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/positions/instructions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type V1AccountPositionInstructionCancelPositionInstructionResponse struct {
	// The API representation of a single CSC instruction, combining the caller's
	// request with the `oems-csc` lifecycle state.
	Data PositionInstruction `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountPositionInstructionCancelPositionInstructionResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *V1AccountPositionInstructionCancelPositionInstructionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountPositionInstructionGetPositionInstructionsResponse struct {
	Data PositionInstructionList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountPositionInstructionGetPositionInstructionsResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *V1AccountPositionInstructionGetPositionInstructionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountPositionInstructionSubmitPositionInstructionsResponse struct {
	Data PositionInstructionList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountPositionInstructionSubmitPositionInstructionsResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *V1AccountPositionInstructionSubmitPositionInstructionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountPositionInstructionCancelPositionInstructionParams struct {
	AccountID int64 `path:"account_id" api:"required" json:"-"`
	paramObj
}

type V1AccountPositionInstructionGetPositionInstructionsParams struct {
	// Filter by OEMS instrument id or symbol (CMS / OSI).
	InstrumentID param.Opt[InstrumentIDOrSymbol] `query:"instrument_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes
// [V1AccountPositionInstructionGetPositionInstructionsParams]'s query parameters
// as `url.Values`.
func (r V1AccountPositionInstructionGetPositionInstructionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1AccountPositionInstructionSubmitPositionInstructionsParams struct {
	Instructions []V1AccountPositionInstructionSubmitPositionInstructionsParamsInstruction
	paramObj
}

func (r V1AccountPositionInstructionSubmitPositionInstructionsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Instructions)
}
func (r *V1AccountPositionInstructionSubmitPositionInstructionsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One exercise / DNE / CEA instruction requested by a client.
//
// Cancel is not an instruction type — use
// `DELETE /accounts/{account_id}/positions/instructions/{instruction_id}`.
//
// The properties InstructionType, InstrumentID, Quantity are required.
type V1AccountPositionInstructionSubmitPositionInstructionsParamsInstruction struct {
	// Instruction type.
	//
	// Any of "EXERCISE", "DO_NOT_EXERCISE", "CONTRARY_EXERCISE".
	InstructionType PositionInstructionType `json:"instruction_type,omitzero" api:"required"`
	// OEMS instrument identifier. api-gw resolves this to `security_id` +
	// `security_id_source` via the instrument cache before dispatching to `oems-csc`.
	// Unknown ids return 404.
	InstrumentID string `json:"instrument_id" api:"required" format:"uuid"`
	// Quantity of contracts to exercise / DNE / CEA.
	Quantity string `json:"quantity" api:"required"`
	// Caller-supplied instruction id. Echoed back on the response and used as the FIX
	// `pos_req_id` (tag 710) for idempotency. If omitted the server generates a UUID.
	InstructionID param.Opt[string] `json:"instruction_id,omitzero"`
	paramObj
}

func (r V1AccountPositionInstructionSubmitPositionInstructionsParamsInstruction) MarshalJSON() (data []byte, err error) {
	type shadow V1AccountPositionInstructionSubmitPositionInstructionsParamsInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AccountPositionInstructionSubmitPositionInstructionsParamsInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
