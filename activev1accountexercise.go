// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"time"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/respjson"
)

// ActiveV1AccountExerciseService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1AccountExerciseService] method instead.
type ActiveV1AccountExerciseService struct {
	options []option.RequestOption
}

// NewActiveV1AccountExerciseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1AccountExerciseService(opts ...option.RequestOption) (r ActiveV1AccountExerciseService) {
	r = ActiveV1AccountExerciseService{}
	r.options = opts
	return
}

// The action a caller wants `oems-csc` to take against an options position.
//
// Maps onto FIX `PosTransType` (tag 709) + `PosMaintAction` (tag 712) +
// `ContraryInstructionIndicator` (tag 719) per `oems-csc`'s `classify_action`.
type ExerciseAction string

const (
	ExerciseActionExercise         ExerciseAction = "EXERCISE"
	ExerciseActionDoNotExercise    ExerciseAction = "DO_NOT_EXERCISE"
	ExerciseActionContraryExercise ExerciseAction = "CONTRARY_EXERCISE"
)

// The API representation of a single CSC instruction, combining the caller's
// request with the `oems-csc` lifecycle state.
type ExerciseInstruction struct {
	// Stable server-assigned id for the instruction (the engine instruction UUID).
	// Used as the `{exercise_id}` path parameter on DELETE.
	ID string `json:"id" api:"required" format:"uuid"`
	// Account the instruction belongs to.
	AccountID int64 `json:"account_id" api:"required"`
	// The instruction type as understood by this API.
	//
	// Any of "EXERCISE", "DO_NOT_EXERCISE", "CONTRARY_EXERCISE".
	Action ExerciseAction `json:"action" api:"required"`
	// Caller-supplied correlation id (echoed from the submit request, or the
	// server-generated fallback when the caller omitted one).
	ClientExerciseID string `json:"client_exercise_id" api:"required"`
	// OEMS instrument identifier the instruction is for.
	InstrumentID string `json:"instrument_id" api:"required" format:"uuid"`
	// Quantity of contracts.
	Quantity string `json:"quantity" api:"required"`
	// Security identifier (display-only; resolved from the instrument cache).
	SecurityID string `json:"security_id" api:"required"`
	// Security identifier source (display-only).
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "CURRENCY", "FMP", "OEMS",
	// "SEDOL", "QUIK", "ISIN", "RIC", "COUNTRY", "EXCHANGE", "CTA", "BLOOMBERG",
	// "WERTPAPIER", "DUTCH", "VALOREN", "SICOVAM", "BELGIAN", "COMMON",
	// "CLEARING_HOUSE", "ISDA_FPML_SPECIFICATION", "ISDA_FPML_URL",
	// "LETTER_OF_CREDIT", "MARKETPLACE_ASSIGNED_IDENTIFIER", "MARKIT_RED_ENTITY_CLIP",
	// "MARKIT_RED_PAIR_CLIP", "CFTC", "ISDA_COMMODITY_REFERENCE_PRICE",
	// "LEGAL_ENTITY_IDENTIFIER", "SYNTHETIC", "FIDESSA_INSTRUMENT_MNEMONIC",
	// "INDEX_NAME", "UNIFORM_SYMBOL", "DIGITAL_TOKEN_IDENTIFIER", "MASSIVE", "OTHER".
	SecurityIDSource SecurityIDSource `json:"security_id_source" api:"required"`
	// Current lifecycle status.
	//
	// Any of "SENT", "ACCEPTED", "REJECTED", "ENGINE_REJECTED", "CANCEL_REQUESTED",
	// "CANCELLED", "CANCEL_FAILED", "UNKNOWN".
	Status ExerciseStatus `json:"status" api:"required"`
	// Trading symbol resolved from the instrument cache. Empty if the instrument
	// cannot be resolved (e.g. expired option).
	Symbol string `json:"symbol" api:"required"`
	// Quantity accepted by OCC. Populated after `ACCEPTED`.
	AcceptedQuantity string `json:"accepted_quantity" api:"nullable"`
	// Row creation timestamp surfaced from `oems-csc`.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Inline error detail when a batch entry was rejected (omitted on success).
	Error string `json:"error" api:"nullable"`
	// Reason text populated on terminal reject / cancel-failed statuses.
	RejectionReason string `json:"rejection_reason" api:"nullable"`
	// Last update timestamp surfaced from `oems-csc`.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AccountID        respjson.Field
		Action           respjson.Field
		ClientExerciseID respjson.Field
		InstrumentID     respjson.Field
		Quantity         respjson.Field
		SecurityID       respjson.Field
		SecurityIDSource respjson.Field
		Status           respjson.Field
		Symbol           respjson.Field
		AcceptedQuantity respjson.Field
		CreatedAt        respjson.Field
		Error            respjson.Field
		RejectionReason  respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExerciseInstruction) RawJSON() string { return r.JSON.raw }
func (r *ExerciseInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExerciseInstructionList []ExerciseInstruction

// Public Active API lifecycle status for an exercise instruction.
//
// Maps 1:1 to the `oems-csc` wire enum while keeping the REST schema stable:
// api-gw owns serialization, OpenAPI generation, and the `Unknown` fallback for
// missing or unrecognized gRPC values.
type ExerciseStatus string

const (
	ExerciseStatusSent            ExerciseStatus = "SENT"
	ExerciseStatusAccepted        ExerciseStatus = "ACCEPTED"
	ExerciseStatusRejected        ExerciseStatus = "REJECTED"
	ExerciseStatusEngineRejected  ExerciseStatus = "ENGINE_REJECTED"
	ExerciseStatusCancelRequested ExerciseStatus = "CANCEL_REQUESTED"
	ExerciseStatusCancelled       ExerciseStatus = "CANCELLED"
	ExerciseStatusCancelFailed    ExerciseStatus = "CANCEL_FAILED"
	ExerciseStatusUnknown         ExerciseStatus = "UNKNOWN"
)
