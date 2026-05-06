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
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/param"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Retrieve details and lists of tradable instruments.
//
// V1InstrumentCashFlowStatementService contains methods and other services that
// help with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1InstrumentCashFlowStatementService] method instead.
type V1InstrumentCashFlowStatementService struct {
	options []option.RequestOption
}

// NewV1InstrumentCashFlowStatementService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1InstrumentCashFlowStatementService(opts ...option.RequestOption) (r V1InstrumentCashFlowStatementService) {
	r = V1InstrumentCashFlowStatementService{}
	r.options = opts
	return
}

// Get cash flow statements for an instrument.
//
// Retrieves historical cash flow statements for the specified instrument. Cash
// flow statements show cash inflows and outflows from operating, investing, and
// financing activities.
func (r *V1InstrumentCashFlowStatementService) GetInstrumentCashFlowStatements(ctx context.Context, instrumentID string, query V1InstrumentCashFlowStatementGetInstrumentCashFlowStatementsParams, opts ...option.RequestOption) (res *V1InstrumentCashFlowStatementGetInstrumentCashFlowStatementsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/instruments/%s/cash-flow-statements", instrumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A quarterly cash flow statement for an instrument.
type InstrumentCashFlowStatement struct {
	// The date and time when the filing was accepted by the SEC
	AcceptedDate time.Time `json:"accepted_date" api:"required" format:"date-time"`
	// The date the financial statement was filed
	FilingDate time.Time `json:"filing_date" api:"required" format:"date"`
	// The fiscal period identifier (e.g., "Q1", "Q2", "Q3", "Q4")
	Period string `json:"period" api:"required"`
	// The type of fiscal period
	//
	// Any of "QUARTERLY", "ANNUAL", "TTM", "BIANNUAL".
	PeriodType FiscalPeriodType `json:"period_type" api:"required"`
	// The currency in which the statement is reported (ISO 4217)
	ReportedCurrency string `json:"reported_currency" api:"required"`
	// The fiscal year of the statement
	Year int64 `json:"year" api:"required"`
	// Change in accounts payables
	AccountsPayables string `json:"accounts_payables" api:"nullable"`
	// Change in accounts receivables
	AccountsReceivables string `json:"accounts_receivables" api:"nullable"`
	// Net acquisitions
	AcquisitionsNet string `json:"acquisitions_net" api:"nullable"`
	// Capital expenditure
	CapitalExpenditure string `json:"capital_expenditure" api:"nullable"`
	// Cash and cash equivalents at beginning of period
	CashAtBeginningOfPeriod string `json:"cash_at_beginning_of_period" api:"nullable"`
	// Cash and cash equivalents at end of period
	CashAtEndOfPeriod string `json:"cash_at_end_of_period" api:"nullable"`
	// Change in working capital
	ChangeInWorkingCapital string `json:"change_in_working_capital" api:"nullable"`
	// Common dividends paid
	CommonDividendsPaid string `json:"common_dividends_paid" api:"nullable"`
	// Common stock issuance
	CommonStockIssuance string `json:"common_stock_issuance" api:"nullable"`
	// Common stock repurchased (buybacks)
	CommonStockRepurchased string `json:"common_stock_repurchased" api:"nullable"`
	// Deferred income tax expense
	DeferredIncomeTax string `json:"deferred_income_tax" api:"nullable"`
	// Depreciation and amortization expense
	DepreciationAndAmortization string `json:"depreciation_and_amortization" api:"nullable"`
	// Effect of foreign exchange changes on cash
	EffectOfForexChangesOnCash string `json:"effect_of_forex_changes_on_cash" api:"nullable"`
	// Free cash flow (operating cash flow minus capital expenditure)
	FreeCashFlow string `json:"free_cash_flow" api:"nullable"`
	// Income taxes paid
	IncomeTaxesPaid string `json:"income_taxes_paid" api:"nullable"`
	// Interest paid
	InterestPaid string `json:"interest_paid" api:"nullable"`
	// Change in inventory
	Inventory string `json:"inventory" api:"nullable"`
	// Investments in property, plant, and equipment
	InvestmentsInPropertyPlantAndEquipment string `json:"investments_in_property_plant_and_equipment" api:"nullable"`
	// Long-term net debt issuance
	LongTermNetDebtIssuance string `json:"long_term_net_debt_issuance" api:"nullable"`
	// Net cash provided by financing activities
	NetCashProvidedByFinancingActivities string `json:"net_cash_provided_by_financing_activities" api:"nullable"`
	// Net cash provided by investing activities
	NetCashProvidedByInvestingActivities string `json:"net_cash_provided_by_investing_activities" api:"nullable"`
	// Net cash provided by operating activities
	NetCashProvidedByOperatingActivities string `json:"net_cash_provided_by_operating_activities" api:"nullable"`
	// Net change in cash during the period
	NetChangeInCash string `json:"net_change_in_cash" api:"nullable"`
	// Net common stock issuance
	NetCommonStockIssuance string `json:"net_common_stock_issuance" api:"nullable"`
	// Net debt issuance (long-term + short-term)
	NetDebtIssuance string `json:"net_debt_issuance" api:"nullable"`
	// Net dividends paid (common + preferred)
	NetDividendsPaid string `json:"net_dividends_paid" api:"nullable"`
	// Net income for the period
	NetIncome string `json:"net_income" api:"nullable"`
	// Net preferred stock issuance
	NetPreferredStockIssuance string `json:"net_preferred_stock_issuance" api:"nullable"`
	// Net stock issuance (common + preferred)
	NetStockIssuance string `json:"net_stock_issuance" api:"nullable"`
	// Operating cash flow (alternative calculation)
	OperatingCashFlow string `json:"operating_cash_flow" api:"nullable"`
	// Other financing activities
	OtherFinancingActivities string `json:"other_financing_activities" api:"nullable"`
	// Other investing activities
	OtherInvestingActivities string `json:"other_investing_activities" api:"nullable"`
	// Other non-cash items
	OtherNonCashItems string `json:"other_non_cash_items" api:"nullable"`
	// Change in other working capital
	OtherWorkingCapital string `json:"other_working_capital" api:"nullable"`
	// Preferred dividends paid
	PreferredDividendsPaid string `json:"preferred_dividends_paid" api:"nullable"`
	// Purchases of investments
	PurchasesOfInvestments string `json:"purchases_of_investments" api:"nullable"`
	// Sales and maturities of investments
	SalesMaturitiesOfInvestments string `json:"sales_maturities_of_investments" api:"nullable"`
	// Short-term net debt issuance
	ShortTermNetDebtIssuance string `json:"short_term_net_debt_issuance" api:"nullable"`
	// Stock-based compensation expense
	StockBasedCompensation string `json:"stock_based_compensation" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptedDate                           respjson.Field
		FilingDate                             respjson.Field
		Period                                 respjson.Field
		PeriodType                             respjson.Field
		ReportedCurrency                       respjson.Field
		Year                                   respjson.Field
		AccountsPayables                       respjson.Field
		AccountsReceivables                    respjson.Field
		AcquisitionsNet                        respjson.Field
		CapitalExpenditure                     respjson.Field
		CashAtBeginningOfPeriod                respjson.Field
		CashAtEndOfPeriod                      respjson.Field
		ChangeInWorkingCapital                 respjson.Field
		CommonDividendsPaid                    respjson.Field
		CommonStockIssuance                    respjson.Field
		CommonStockRepurchased                 respjson.Field
		DeferredIncomeTax                      respjson.Field
		DepreciationAndAmortization            respjson.Field
		EffectOfForexChangesOnCash             respjson.Field
		FreeCashFlow                           respjson.Field
		IncomeTaxesPaid                        respjson.Field
		InterestPaid                           respjson.Field
		Inventory                              respjson.Field
		InvestmentsInPropertyPlantAndEquipment respjson.Field
		LongTermNetDebtIssuance                respjson.Field
		NetCashProvidedByFinancingActivities   respjson.Field
		NetCashProvidedByInvestingActivities   respjson.Field
		NetCashProvidedByOperatingActivities   respjson.Field
		NetChangeInCash                        respjson.Field
		NetCommonStockIssuance                 respjson.Field
		NetDebtIssuance                        respjson.Field
		NetDividendsPaid                       respjson.Field
		NetIncome                              respjson.Field
		NetPreferredStockIssuance              respjson.Field
		NetStockIssuance                       respjson.Field
		OperatingCashFlow                      respjson.Field
		OtherFinancingActivities               respjson.Field
		OtherInvestingActivities               respjson.Field
		OtherNonCashItems                      respjson.Field
		OtherWorkingCapital                    respjson.Field
		PreferredDividendsPaid                 respjson.Field
		PurchasesOfInvestments                 respjson.Field
		SalesMaturitiesOfInvestments           respjson.Field
		ShortTermNetDebtIssuance               respjson.Field
		StockBasedCompensation                 respjson.Field
		ExtraFields                            map[string]respjson.Field
		raw                                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentCashFlowStatement) RawJSON() string { return r.JSON.raw }
func (r *InstrumentCashFlowStatement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstrumentCashFlowStatementList []InstrumentCashFlowStatement

type V1InstrumentCashFlowStatementGetInstrumentCashFlowStatementsResponse struct {
	Data InstrumentCashFlowStatementList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentCashFlowStatementGetInstrumentCashFlowStatementsResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *V1InstrumentCashFlowStatementGetInstrumentCashFlowStatementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentCashFlowStatementGetInstrumentCashFlowStatementsParams struct {
	// The start date for the query range, inclusive (YYYY-MM-DD).
	FromDate param.Opt[string] `query:"from_date,omitzero" json:"-"`
	PageSize param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD).
	ToDate param.Opt[string] `query:"to_date,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [V1InstrumentCashFlowStatementGetInstrumentCashFlowStatementsParams]'s query
// parameters as `url.Values`.
func (r V1InstrumentCashFlowStatementGetInstrumentCashFlowStatementsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
