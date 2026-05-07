// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"time"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/respjson"
)

// ActiveV1AccountPositionInstructionService contains methods and other services
// that help with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1AccountPositionInstructionService] method instead.
type ActiveV1AccountPositionInstructionService struct {
	options []option.RequestOption
}

// NewActiveV1AccountPositionInstructionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewActiveV1AccountPositionInstructionService(opts ...option.RequestOption) (r ActiveV1AccountPositionInstructionService) {
	r = ActiveV1AccountPositionInstructionService{}
	r.options = opts
	return
}

// The API representation of a single CSC instruction, combining the caller's
// request with the `oems-csc` lifecycle state.
type PositionInstruction struct {
	// Stable server-assigned id for the instruction (the engine instruction UUID).
	// Used as the `{instruction_id}` path parameter on DELETE.
	ID string `json:"id" api:"required" format:"uuid"`
	// Account the instruction belongs to.
	AccountID int64 `json:"account_id" api:"required"`
	// Caller-supplied instruction id (echoed from the submit request, or the
	// server-generated fallback when the caller omitted one).
	InstructionID string `json:"instruction_id" api:"required"`
	// The instruction type as understood by this API.
	//
	// Any of "EXERCISE", "DO_NOT_EXERCISE", "CONTRARY_EXERCISE".
	InstructionType PositionInstructionType `json:"instruction_type" api:"required"`
	// OEMS instrument identifier the instruction is for.
	InstrumentID string `json:"instrument_id" api:"required" format:"uuid"`
	// OSI option symbol (e.g. `AAPL 280121C00195000`). Display-only; resolved from the
	// instrument cache.
	Osi string `json:"osi" api:"required"`
	// Quantity of contracts.
	Quantity string `json:"quantity" api:"required"`
	// Current lifecycle status.
	//
	// Any of "SENT", "ACCEPTED", "REJECTED", "ENGINE_REJECTED", "CANCEL_REQUESTED",
	// "CANCELLED", "CANCEL_FAILED", "UNKNOWN".
	Status PositionInstructionStatus `json:"status" api:"required"`
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
		InstructionID    respjson.Field
		InstructionType  respjson.Field
		InstrumentID     respjson.Field
		Osi              respjson.Field
		Quantity         respjson.Field
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
func (r PositionInstruction) RawJSON() string { return r.JSON.raw }
func (r *PositionInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PositionInstructionList []PositionInstruction

// Public Active API lifecycle status for a position instruction.
//
// Maps 1:1 to the `oems-csc` wire enum while keeping the REST schema stable:
// api-gw owns serialization, OpenAPI generation, and the `Unknown` fallback for
// missing or unrecognized gRPC values.
type PositionInstructionStatus string

const (
	PositionInstructionStatusSent            PositionInstructionStatus = "SENT"
	PositionInstructionStatusAccepted        PositionInstructionStatus = "ACCEPTED"
	PositionInstructionStatusRejected        PositionInstructionStatus = "REJECTED"
	PositionInstructionStatusEngineRejected  PositionInstructionStatus = "ENGINE_REJECTED"
	PositionInstructionStatusCancelRequested PositionInstructionStatus = "CANCEL_REQUESTED"
	PositionInstructionStatusCancelled       PositionInstructionStatus = "CANCELLED"
	PositionInstructionStatusCancelFailed    PositionInstructionStatus = "CANCEL_FAILED"
	PositionInstructionStatusUnknown         PositionInstructionStatus = "UNKNOWN"
)

// The instruction type a caller wants `oems-csc` to take against an options
// position.
//
// Maps onto FIX `PosTransType` (tag 709) + `PosMaintAction` (tag 712) +
// `ContraryInstructionIndicator` (tag 719) per `oems-csc`'s `classify_action`.
type PositionInstructionType string

const (
	PositionInstructionTypeExercise         PositionInstructionType = "EXERCISE"
	PositionInstructionTypeDoNotExercise    PositionInstructionType = "DO_NOT_EXERCISE"
	PositionInstructionTypeContraryExercise PositionInstructionType = "CONTRARY_EXERCISE"
)
