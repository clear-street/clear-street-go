// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"time"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/respjson"
)

// V1ExecutionService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ExecutionService] method instead.
type V1ExecutionService struct {
	options []option.RequestOption
}

// NewV1ExecutionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1ExecutionService(opts ...option.RequestOption) (r V1ExecutionService) {
	r = V1ExecutionService{}
	r.options = opts
	return
}

// Represents a single fill of an order for an account.
type Execution struct {
	// Unique identifier for this execution report.
	ID string `json:"id" api:"required" format:"uuid"`
	// OEMS instrument identifier.
	InstrumentID string `json:"instrument_id" api:"required" format:"uuid"`
	// Identifier of the order this execution belongs to.
	OrderID string `json:"order_id" api:"required" format:"uuid"`
	// Fill price.
	Price string `json:"price" api:"required"`
	// Filled quantity.
	Quantity string `json:"quantity" api:"required"`
	// Side of the fill.
	//
	// Any of "BUY", "SELL", "SELL_SHORT", "OTHER".
	Side Side `json:"side" api:"required"`
	// Trading symbol.
	Symbol string `json:"symbol" api:"required"`
	// Transaction timestamp in nanosecond precision (UTC).
	TransactionTime time.Time `json:"transaction_time" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		InstrumentID    respjson.Field
		OrderID         respjson.Field
		Price           respjson.Field
		Quantity        respjson.Field
		Side            respjson.Field
		Symbol          respjson.Field
		TransactionTime respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Execution) RawJSON() string { return r.JSON.raw }
func (r *Execution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecutionList []Execution
