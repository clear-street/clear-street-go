// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/internal/apiquery"
	shimjson "github.com/clear-street/clear-street-go/internal/encoding/json"
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/param"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// View positions and manage position instructions.
//
// V1PositionService contains methods and other services that help with interacting
// with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1PositionService] method instead.
type V1PositionService struct {
	options []option.RequestOption
}

// NewV1PositionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1PositionService(opts ...option.RequestOption) (r V1PositionService) {
	r = V1PositionService{}
	r.options = opts
	return
}

// Cancel an outstanding exercise / DNE / CEA instruction by its server- assigned
// `id`. Returns the updated instruction with status `CANCEL_REQUESTED`; the
// terminal `CANCELLED` / `CANCEL_FAILED` state arrives asynchronously via
// subsequent GETs.
func (r *V1PositionService) CancelPositionInstruction(ctx context.Context, instructionID string, body V1PositionCancelPositionInstructionParams, opts ...option.RequestOption) (res *V1PositionCancelPositionInstructionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instructionID == "" {
		err = errors.New("missing required instruction_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%v/positions/instructions/%s", body.AccountID, instructionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Delete a position within an account for an instrument.
//
// Retrieves orders generated to close the position.
func (r *V1PositionService) ClosePosition(ctx context.Context, instrumentID InstrumentIDOrSymbol, params V1PositionClosePositionParams, opts ...option.RequestOption) (res *V1PositionClosePositionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%v/positions/%s", params.AccountID, instrumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Delete all positions within an account.
//
// Closes all positions for the specified trading account.
func (r *V1PositionService) ClosePositions(ctx context.Context, accountID int64, body V1PositionClosePositionsParams, opts ...option.RequestOption) (res *V1PositionClosePositionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/positions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Returns the current lifecycle state of exercise / DNE / CEA instructions for the
// account. Optionally filter by a specific instrument.
func (r *V1PositionService) GetPositionInstructions(ctx context.Context, accountID int64, query V1PositionGetPositionInstructionsParams, opts ...option.RequestOption) (res *V1PositionGetPositionInstructionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/positions/instructions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves all positions for the specified trading account.
func (r *V1PositionService) GetPositions(ctx context.Context, accountID int64, query V1PositionGetPositionsParams, opts ...option.RequestOption) (res *V1PositionGetPositionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/positions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Submit one or more option lifecycle instructions against the account. Each row
// is routed to `oems-csc` independently; per-row rejections are surfaced on the
// corresponding response entry without failing the batch.
func (r *V1PositionService) SubmitPositionInstructions(ctx context.Context, accountID int64, body V1PositionSubmitPositionInstructionsParams, opts ...option.RequestOption) (res *V1PositionSubmitPositionInstructionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/positions/instructions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Represents a holding of a particular instrument in an account
type Position struct {
	// The account this position belongs to
	AccountID int64 `json:"account_id" api:"required"`
	// The quantity of a position that is free to be operated on.
	AvailableQuantity string `json:"available_quantity" api:"required"`
	// OEMS instrument UUID
	InstrumentID string `json:"instrument_id" api:"required" format:"uuid"`
	// Type of security
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "OPTION", "CASH", "OTHER".
	InstrumentType SecurityType `json:"instrument_type" api:"required"`
	// The current market value of the position
	MarketValue string `json:"market_value" api:"required"`
	// The type of position
	//
	// Any of "LONG", "SHORT", "LONG_CALL", "SHORT_CALL", "LONG_PUT", "SHORT_PUT".
	PositionType PositionType `json:"position_type" api:"required"`
	// The number of shares or contracts. Can be positive (long) or negative (short)
	Quantity string `json:"quantity" api:"required"`
	// The trading symbol for the instrument
	Symbol string `json:"symbol" api:"required"`
	// The average price paid per share or contract for this position
	AvgPrice string `json:"avg_price" api:"nullable"`
	// The closing price used to value the position for the last trading day
	ClosingPrice string `json:"closing_price" api:"nullable"`
	// The market date associated with `closing_price`
	ClosingPriceDate time.Time `json:"closing_price_date" api:"nullable" format:"date"`
	// The total cost basis for this position
	CostBasis string `json:"cost_basis" api:"nullable"`
	// The unrealized profit or loss for this position relative to the previous close
	DailyUnrealizedPnl string `json:"daily_unrealized_pnl" api:"nullable"`
	// The unrealized profit/loss for the position for the current day, expressed as a
	// percentage of the baseline value (range: 0-100).
	DailyUnrealizedPnlPct string `json:"daily_unrealized_pnl_pct" api:"nullable"`
	// The current market price of the instrument
	InstrumentPrice string `json:"instrument_price" api:"nullable"`
	// OEMS instrument identifier of the underlying instrument, if resolvable
	UnderlyingInstrumentID string `json:"underlying_instrument_id" api:"nullable" format:"uuid"`
	// The total unrealized profit or loss for this position based on current market
	// value
	UnrealizedPnl string `json:"unrealized_pnl" api:"nullable"`
	// The unrealized profit/loss for the position, expressed as a percentage of the
	// position's cost basis (range: 0-100).
	UnrealizedPnlPct string `json:"unrealized_pnl_pct" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID              respjson.Field
		AvailableQuantity      respjson.Field
		InstrumentID           respjson.Field
		InstrumentType         respjson.Field
		MarketValue            respjson.Field
		PositionType           respjson.Field
		Quantity               respjson.Field
		Symbol                 respjson.Field
		AvgPrice               respjson.Field
		ClosingPrice           respjson.Field
		ClosingPriceDate       respjson.Field
		CostBasis              respjson.Field
		DailyUnrealizedPnl     respjson.Field
		DailyUnrealizedPnlPct  respjson.Field
		InstrumentPrice        respjson.Field
		UnderlyingInstrumentID respjson.Field
		UnrealizedPnl          respjson.Field
		UnrealizedPnlPct       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Position) RawJSON() string { return r.JSON.raw }
func (r *Position) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	// Quantity of contracts.
	Quantity string `json:"quantity" api:"required"`
	// Current lifecycle status.
	//
	// Any of "SENT", "ACCEPTED", "REJECTED", "ENGINE_REJECTED", "CANCEL_REQUESTED",
	// "CANCELLED", "CANCEL_FAILED", "UNKNOWN".
	Status PositionInstructionStatus `json:"status" api:"required"`
	// Trading symbol resolved from the instrument cache (OSI for options, since
	// exercises are options-only). Empty if the instrument cannot be resolved (e.g.
	// expired option). Display-only.
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

type PositionList []Position

// Position type classification
type PositionType string

const (
	PositionTypeLong      PositionType = "LONG"
	PositionTypeShort     PositionType = "SHORT"
	PositionTypeLongCall  PositionType = "LONG_CALL"
	PositionTypeShortCall PositionType = "SHORT_CALL"
	PositionTypeLongPut   PositionType = "LONG_PUT"
	PositionTypeShortPut  PositionType = "SHORT_PUT"
)

type V1PositionCancelPositionInstructionResponse struct {
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
func (r V1PositionCancelPositionInstructionResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PositionCancelPositionInstructionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PositionClosePositionResponse struct {
	Data OrderList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1PositionClosePositionResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PositionClosePositionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PositionClosePositionsResponse struct {
	Data OrderList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1PositionClosePositionsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PositionClosePositionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PositionGetPositionInstructionsResponse struct {
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
func (r V1PositionGetPositionInstructionsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PositionGetPositionInstructionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PositionGetPositionsResponse struct {
	Data PositionList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1PositionGetPositionsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PositionGetPositionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PositionSubmitPositionInstructionsResponse struct {
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
func (r V1PositionSubmitPositionInstructionsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PositionSubmitPositionInstructionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PositionCancelPositionInstructionParams struct {
	AccountID int64 `path:"account_id" api:"required" json:"-"`
	paramObj
}

type V1PositionClosePositionParams struct {
	AccountID int64 `path:"account_id" api:"required" json:"-"`
	// Whether to cancel existing open orders for the position before submitting
	// closing orders.
	CancelOrders param.Opt[bool] `json:"cancel_orders,omitzero"`
	paramObj
}

func (r V1PositionClosePositionParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PositionClosePositionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PositionClosePositionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PositionClosePositionsParams struct {
	// Whether to cancel existing open orders for the position before submitting
	// closing orders.
	CancelOrders param.Opt[bool] `json:"cancel_orders,omitzero"`
	paramObj
}

func (r V1PositionClosePositionsParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PositionClosePositionsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PositionClosePositionsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PositionGetPositionInstructionsParams struct {
	// Filter by OEMS instrument id or symbol (CMS / OSI).
	InstrumentID param.Opt[InstrumentIDOrSymbol] `query:"instrument_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [V1PositionGetPositionInstructionsParams]'s query parameters
// as `url.Values`.
func (r V1PositionGetPositionInstructionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1PositionGetPositionsParams struct {
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Comma-separated OEMS instrument UUIDs
	InstrumentIDs []string `query:"instrument_ids,omitzero" format:"uuid" json:"-"`
	// Field to sort by
	//
	// Any of "SYMBOL", "INSTRUMENT_TYPE", "QUANTITY", "MARKET_VALUE", "POSITION_TYPE",
	// "UNREALIZED_PNL", "DAILY_UNREALIZED_PNL".
	SortBy V1PositionGetPositionsParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Sort direction
	//
	// Any of "ASC", "DESC".
	SortDirection V1PositionGetPositionsParamsSortDirection `query:"sort_direction,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1PositionGetPositionsParams]'s query parameters as
// `url.Values`.
func (r V1PositionGetPositionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Field to sort by
type V1PositionGetPositionsParamsSortBy string

const (
	V1PositionGetPositionsParamsSortBySymbol             V1PositionGetPositionsParamsSortBy = "SYMBOL"
	V1PositionGetPositionsParamsSortByInstrumentType     V1PositionGetPositionsParamsSortBy = "INSTRUMENT_TYPE"
	V1PositionGetPositionsParamsSortByQuantity           V1PositionGetPositionsParamsSortBy = "QUANTITY"
	V1PositionGetPositionsParamsSortByMarketValue        V1PositionGetPositionsParamsSortBy = "MARKET_VALUE"
	V1PositionGetPositionsParamsSortByPositionType       V1PositionGetPositionsParamsSortBy = "POSITION_TYPE"
	V1PositionGetPositionsParamsSortByUnrealizedPnl      V1PositionGetPositionsParamsSortBy = "UNREALIZED_PNL"
	V1PositionGetPositionsParamsSortByDailyUnrealizedPnl V1PositionGetPositionsParamsSortBy = "DAILY_UNREALIZED_PNL"
)

// Sort direction
type V1PositionGetPositionsParamsSortDirection string

const (
	V1PositionGetPositionsParamsSortDirectionAsc  V1PositionGetPositionsParamsSortDirection = "ASC"
	V1PositionGetPositionsParamsSortDirectionDesc V1PositionGetPositionsParamsSortDirection = "DESC"
)

type V1PositionSubmitPositionInstructionsParams struct {
	Instructions []V1PositionSubmitPositionInstructionsParamsInstruction
	paramObj
}

func (r V1PositionSubmitPositionInstructionsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Instructions)
}
func (r *V1PositionSubmitPositionInstructionsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One exercise / DNE / CEA instruction requested by a client.
//
// Cancel is not an instruction type — use
// `DELETE /accounts/{account_id}/positions/instructions/{instruction_id}`.
//
// The properties InstructionType, InstrumentID, Quantity are required.
type V1PositionSubmitPositionInstructionsParamsInstruction struct {
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

func (r V1PositionSubmitPositionInstructionsParamsInstruction) MarshalJSON() (data []byte, err error) {
	type shadow V1PositionSubmitPositionInstructionsParamsInstruction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PositionSubmitPositionInstructionsParamsInstruction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
