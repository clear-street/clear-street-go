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
// V1AccountExerciseService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1AccountExerciseService] method instead.
type V1AccountExerciseService struct {
	options []option.RequestOption
}

// NewV1AccountExerciseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1AccountExerciseService(opts ...option.RequestOption) (r V1AccountExerciseService) {
	r = V1AccountExerciseService{}
	r.options = opts
	return
}

// Cancel an outstanding exercise / DNE / CEA instruction by its server- assigned
// `id`. Returns the updated instruction with status `CANCEL_REQUESTED`; the
// terminal `CANCELLED` / `CANCEL_FAILED` state arrives asynchronously via
// subsequent GETs.
func (r *V1AccountExerciseService) CancelExercise(ctx context.Context, exerciseID string, body V1AccountExerciseCancelExerciseParams, opts ...option.RequestOption) (res *V1AccountExerciseCancelExerciseResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if exerciseID == "" {
		err = errors.New("missing required exercise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%v/exercises/%s", body.AccountID, exerciseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns the current lifecycle state of exercise / DNE / CEA instructions for the
// account. Optionally filter by a specific instrument.
func (r *V1AccountExerciseService) GetExercises(ctx context.Context, accountID int64, query V1AccountExerciseGetExercisesParams, opts ...option.RequestOption) (res *V1AccountExerciseGetExercisesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/exercises", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Submit one or more option lifecycle instructions against the account. Each row
// is routed to `oems-csc` independently; per-row rejections are surfaced on the
// corresponding response entry without failing the batch.
func (r *V1AccountExerciseService) SubmitExercises(ctx context.Context, accountID int64, body V1AccountExerciseSubmitExercisesParams, opts ...option.RequestOption) (res *V1AccountExerciseSubmitExercisesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/exercises", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type V1AccountExerciseCancelExerciseResponse struct {
	// The API representation of a single CSC instruction, combining the caller's
	// request with the `oems-csc` lifecycle state.
	Data ExerciseInstruction `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountExerciseCancelExerciseResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountExerciseCancelExerciseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountExerciseGetExercisesResponse struct {
	Data ExerciseInstructionList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountExerciseGetExercisesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountExerciseGetExercisesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountExerciseSubmitExercisesResponse struct {
	Data ExerciseInstructionList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountExerciseSubmitExercisesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountExerciseSubmitExercisesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountExerciseCancelExerciseParams struct {
	AccountID int64 `path:"account_id" api:"required" json:"-"`
	paramObj
}

type V1AccountExerciseGetExercisesParams struct {
	// Filter by OEMS instrument id.
	InstrumentID param.Opt[string] `query:"instrument_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [V1AccountExerciseGetExercisesParams]'s query parameters as
// `url.Values`.
func (r V1AccountExerciseGetExercisesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1AccountExerciseSubmitExercisesParams struct {
	Exercises []V1AccountExerciseSubmitExercisesParamsExercise
	paramObj
}

func (r V1AccountExerciseSubmitExercisesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Exercises)
}
func (r *V1AccountExerciseSubmitExercisesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One exercise / DNE / CEA instruction requested by a client.
//
// Cancel is not an action — use
// `DELETE /accounts/{account_id}/exercises/{exercise_id}`.
//
// The properties Action, InstrumentID, Quantity are required.
type V1AccountExerciseSubmitExercisesParamsExercise struct {
	// Instruction type.
	//
	// Any of "EXERCISE", "DO_NOT_EXERCISE", "CONTRARY_EXERCISE".
	Action ExerciseAction `json:"action,omitzero" api:"required"`
	// OEMS instrument identifier. api-gw resolves this to `security_id` +
	// `security_id_source` via the instrument cache before dispatching to `oems-csc`.
	// Unknown ids return 404.
	InstrumentID string `json:"instrument_id" api:"required" format:"uuid"`
	// Quantity of contracts to exercise / DNE / CEA.
	Quantity string `json:"quantity" api:"required"`
	// Caller-supplied correlation id. Echoed back on the response and used as the FIX
	// `pos_req_id` (tag 710) for idempotency. If omitted the server generates a UUID.
	ClientExerciseID param.Opt[string] `json:"client_exercise_id,omitzero"`
	paramObj
}

func (r V1AccountExerciseSubmitExercisesParamsExercise) MarshalJSON() (data []byte, err error) {
	type shadow V1AccountExerciseSubmitExercisesParamsExercise
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AccountExerciseSubmitExercisesParamsExercise) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
