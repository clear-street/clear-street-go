// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"time"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/respjson"
)

// V1PrivateMarketSpvService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1PrivateMarketSpvService] method instead.
type V1PrivateMarketSpvService struct {
	options []option.RequestOption
}

// NewV1PrivateMarketSpvService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1PrivateMarketSpvService(opts ...option.RequestOption) (r V1PrivateMarketSpvService) {
	r = V1PrivateMarketSpvService{}
	r.options = opts
	return
}

// Party charging a fee.
type ChargedBy string

const (
	ChargedByFundManager ChargedBy = "FUND_MANAGER"
	ChargedByClearStreet ChargedBy = "CLEAR_STREET"
	ChargedByThirdParty  ChargedBy = "THIRD_PARTY"
)

// Fee timing/cadence.
type FeeFrequency string

const (
	FeeFrequencyOneTime     FeeFrequency = "ONE_TIME"
	FeeFrequencyAnnual      FeeFrequency = "ANNUAL"
	FeeFrequencyAtExit      FeeFrequency = "AT_EXIT"
	FeeFrequencyPassThrough FeeFrequency = "PASS_THROUGH"
)

// Kind of SPV fee.
type FeeType string

const (
	FeeTypeManagement     FeeType = "MANAGEMENT"
	FeeTypeCarry          FeeType = "CARRY"
	FeeTypePlacement      FeeType = "PLACEMENT"
	FeeTypeAdministrative FeeType = "ADMINISTRATIVE"
	FeeTypeOther          FeeType = "OTHER"
)

// An OPEN SPV's identity, exact economics, and typed fee schedule.
type SpvDetail struct {
	// Stable SPV identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// Company whose shares the vehicle holds.
	CompanyID string `json:"company_id" api:"required" format:"uuid"`
	// Terms currency.
	//
	// Any of "USD".
	Currency Currency `json:"currency" api:"required"`
	// Legal/display name.
	Name string `json:"name" api:"required"`
	// Lifecycle state.
	//
	// Any of "DRAFT", "OPEN", "CLOSED", "LIQUIDATING", "DISSOLVED".
	Status SpvStatus `json:"status" api:"required"`
	// Price per share including fees.
	AllInPricePerShare string `json:"all_in_price_per_share" api:"nullable"`
	// Custodian.
	CustodianName string `json:"custodian_name" api:"nullable"`
	// Per-share fee.
	FeePerShare string `json:"fee_per_share" api:"nullable"`
	// Typed fee schedule.
	FeeTerms []SpvFeeTermResource `json:"fee_terms"`
	// Percentage of dollar allocation funded, derived from the allocation pair.
	FundedPercent string `json:"funded_percent" api:"nullable"`
	// Funding deadline.
	FundingDeadline time.Time `json:"funding_deadline" api:"nullable" format:"date-time"`
	// SPV manager.
	ManagerName string `json:"manager_name" api:"nullable"`
	// Minimum investment amount.
	MinimumInvestmentAmount string `json:"minimum_investment_amount" api:"nullable"`
	// Time the vehicle opened.
	OpenedAt time.Time `json:"opened_at" api:"nullable" format:"date-time"`
	// Price per share excluding fees.
	PricePerShare string `json:"price_per_share" api:"nullable"`
	// Remaining dollar allocation.
	RemainingAllocationAmount string `json:"remaining_allocation_amount" api:"nullable"`
	// Remaining share allocation.
	RemainingShareAllocation string `json:"remaining_share_allocation" api:"nullable"`
	// Underlying share class, when specified.
	ShareClass string `json:"share_class" api:"nullable"`
	// Plain-text vehicle structure.
	StructureDescription string `json:"structure_description" api:"nullable"`
	// Total dollar allocation.
	TotalAllocationAmount string `json:"total_allocation_amount" api:"nullable"`
	// Total share allocation.
	TotalShareAllocation string `json:"total_share_allocation" api:"nullable"`
	// Exact company valuation.
	Valuation string `json:"valuation" api:"nullable"`
	// Meaning of `valuation`.
	//
	// Any of "PRE_MONEY", "POST_MONEY", "REFERENCE", "IMPLIED".
	ValuationBasis ValuationBasis `json:"valuation_basis" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		CompanyID                 respjson.Field
		Currency                  respjson.Field
		Name                      respjson.Field
		Status                    respjson.Field
		AllInPricePerShare        respjson.Field
		CustodianName             respjson.Field
		FeePerShare               respjson.Field
		FeeTerms                  respjson.Field
		FundedPercent             respjson.Field
		FundingDeadline           respjson.Field
		ManagerName               respjson.Field
		MinimumInvestmentAmount   respjson.Field
		OpenedAt                  respjson.Field
		PricePerShare             respjson.Field
		RemainingAllocationAmount respjson.Field
		RemainingShareAllocation  respjson.Field
		ShareClass                respjson.Field
		StructureDescription      respjson.Field
		TotalAllocationAmount     respjson.Field
		TotalShareAllocation      respjson.Field
		Valuation                 respjson.Field
		ValuationBasis            respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpvDetail) RawJSON() string { return r.JSON.raw }
func (r *SpvDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One typed SPV fee term.
type SpvFeeTermResource struct {
	// Charging party.
	//
	// Any of "FUND_MANAGER", "CLEAR_STREET", "THIRD_PARTY".
	ChargedBy ChargedBy `json:"charged_by" api:"required"`
	// Terms currency.
	//
	// Any of "USD".
	Currency Currency `json:"currency" api:"required"`
	// Plain-text fee disclosure.
	Description string `json:"description" api:"required"`
	// Fee kind.
	//
	// Any of "MANAGEMENT", "CARRY", "PLACEMENT", "ADMINISTRATIVE", "OTHER".
	FeeType FeeType `json:"fee_type" api:"required"`
	// Timing/cadence.
	//
	// Any of "ONE_TIME", "ANNUAL", "AT_EXIT", "PASS_THROUGH".
	Frequency FeeFrequency `json:"frequency" api:"required"`
	// Exact fixed amount, when amount-based.
	Amount string `json:"amount" api:"nullable"`
	// Charge duration in years, when specified.
	DurationYears string `json:"duration_years" api:"nullable"`
	// Carry hurdle as a decimal fraction, when specified.
	HurdleRate string `json:"hurdle_rate" api:"nullable"`
	// Decimal fraction between zero and one, when percentage-based.
	Rate string `json:"rate" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChargedBy     respjson.Field
		Currency      respjson.Field
		Description   respjson.Field
		FeeType       respjson.Field
		Frequency     respjson.Field
		Amount        respjson.Field
		DurationYears respjson.Field
		HurdleRate    respjson.Field
		Rate          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpvFeeTermResource) RawJSON() string { return r.JSON.raw }
func (r *SpvFeeTermResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
