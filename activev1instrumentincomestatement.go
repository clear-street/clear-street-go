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
// ActiveV1InstrumentIncomeStatementService contains methods and other services
// that help with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1InstrumentIncomeStatementService] method instead.
type ActiveV1InstrumentIncomeStatementService struct {
	options []option.RequestOption
}

// NewActiveV1InstrumentIncomeStatementService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1InstrumentIncomeStatementService(opts ...option.RequestOption) (r ActiveV1InstrumentIncomeStatementService) {
	r = ActiveV1InstrumentIncomeStatementService{}
	r.options = opts
	return
}

// Retrieves quarterly income statements for a specific instrument, sorted by
// fiscal period (most recent first).
//
// Date range defaults:
//
// - `from_date`: None (no lower bound)
// - `to_date`: None (no upper bound)
func (r *ActiveV1InstrumentIncomeStatementService) GetInstrumentIncomeStatements(ctx context.Context, securityID string, params ActiveV1InstrumentIncomeStatementGetInstrumentIncomeStatementsParams, opts ...option.RequestOption) (res *ActiveV1InstrumentIncomeStatementGetInstrumentIncomeStatementsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if securityID == "" {
		err = errors.New("missing required security_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/instruments/%v/%s/income-statements", params.SecurityIDSource, url.PathEscape(securityID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// A quarterly income statement for an instrument.
type InstrumentIncomeStatement struct {
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
	// Bottom line net income after all adjustments
	BottomLineNetIncome string `json:"bottom_line_net_income" api:"nullable"`
	// Total costs and expenses
	CostAndExpenses string `json:"cost_and_expenses" api:"nullable"`
	// Direct costs attributable to producing goods sold
	CostOfRevenue string `json:"cost_of_revenue" api:"nullable"`
	// Depreciation and amortization expenses
	DepreciationAndAmortization string `json:"depreciation_and_amortization" api:"nullable"`
	// Earnings before interest and taxes
	Ebit string `json:"ebit" api:"nullable"`
	// Earnings before interest, taxes, depreciation, and amortization
	Ebitda string `json:"ebitda" api:"nullable"`
	// Basic earnings per share
	Eps string `json:"eps" api:"nullable"`
	// Diluted earnings per share
	EpsDiluted string `json:"eps_diluted" api:"nullable"`
	// General administrative overhead expenses
	GeneralAndAdministrativeExpenses string `json:"general_and_administrative_expenses" api:"nullable"`
	// Revenue minus cost of revenue
	GrossProfit string `json:"gross_profit" api:"nullable"`
	// Income before income tax expense
	IncomeBeforeTax string `json:"income_before_tax" api:"nullable"`
	// Income tax expense for the period
	IncomeTaxExpense string `json:"income_tax_expense" api:"nullable"`
	// Interest paid on debt
	InterestExpense string `json:"interest_expense" api:"nullable"`
	// Interest earned on investments and cash
	InterestIncome string `json:"interest_income" api:"nullable"`
	// Total net income for the period
	NetIncome string `json:"net_income" api:"nullable"`
	// Deductions from net income
	NetIncomeDeductions string `json:"net_income_deductions" api:"nullable"`
	// Net income from continuing operations
	NetIncomeFromContinuingOperations string `json:"net_income_from_continuing_operations" api:"nullable"`
	// Net income from discontinued operations
	NetIncomeFromDiscontinuedOperations string `json:"net_income_from_discontinued_operations" api:"nullable"`
	// Net interest income (interest income minus interest expense)
	NetInterestIncome string `json:"net_interest_income" api:"nullable"`
	// Non-operating income excluding interest
	NonOperatingIncomeExcludingInterest string `json:"non_operating_income_excluding_interest" api:"nullable"`
	// Total operating expenses
	OperatingExpenses string `json:"operating_expenses" api:"nullable"`
	// Income from core business operations
	OperatingIncome string `json:"operating_income" api:"nullable"`
	// Other adjustments to net income
	OtherAdjustmentsToNetIncome string `json:"other_adjustments_to_net_income" api:"nullable"`
	// Other miscellaneous expenses
	OtherExpenses string `json:"other_expenses" api:"nullable"`
	// Expenditure on research and development activities
	ResearchAndDevelopmentExpenses string `json:"research_and_development_expenses" api:"nullable"`
	// Total revenue from sales of goods and services
	Revenue string `json:"revenue" api:"nullable"`
	// Expenditure on marketing and sales activities
	SellingAndMarketingExpenses string `json:"selling_and_marketing_expenses" api:"nullable"`
	// Combined selling, general, and administrative expenses
	SellingGeneralAndAdministrativeExpenses string `json:"selling_general_and_administrative_expenses" api:"nullable"`
	// Net of other income and expenses
	TotalOtherIncomeExpensesNet string `json:"total_other_income_expenses_net" api:"nullable"`
	// Weighted average shares outstanding (basic)
	WeightedAverageShsOut string `json:"weighted_average_shs_out" api:"nullable"`
	// Weighted average shares outstanding (diluted)
	WeightedAverageShsOutDil string `json:"weighted_average_shs_out_dil" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptedDate                            respjson.Field
		FilingDate                              respjson.Field
		Period                                  respjson.Field
		PeriodType                              respjson.Field
		ReportedCurrency                        respjson.Field
		Year                                    respjson.Field
		BottomLineNetIncome                     respjson.Field
		CostAndExpenses                         respjson.Field
		CostOfRevenue                           respjson.Field
		DepreciationAndAmortization             respjson.Field
		Ebit                                    respjson.Field
		Ebitda                                  respjson.Field
		Eps                                     respjson.Field
		EpsDiluted                              respjson.Field
		GeneralAndAdministrativeExpenses        respjson.Field
		GrossProfit                             respjson.Field
		IncomeBeforeTax                         respjson.Field
		IncomeTaxExpense                        respjson.Field
		InterestExpense                         respjson.Field
		InterestIncome                          respjson.Field
		NetIncome                               respjson.Field
		NetIncomeDeductions                     respjson.Field
		NetIncomeFromContinuingOperations       respjson.Field
		NetIncomeFromDiscontinuedOperations     respjson.Field
		NetInterestIncome                       respjson.Field
		NonOperatingIncomeExcludingInterest     respjson.Field
		OperatingExpenses                       respjson.Field
		OperatingIncome                         respjson.Field
		OtherAdjustmentsToNetIncome             respjson.Field
		OtherExpenses                           respjson.Field
		ResearchAndDevelopmentExpenses          respjson.Field
		Revenue                                 respjson.Field
		SellingAndMarketingExpenses             respjson.Field
		SellingGeneralAndAdministrativeExpenses respjson.Field
		TotalOtherIncomeExpensesNet             respjson.Field
		WeightedAverageShsOut                   respjson.Field
		WeightedAverageShsOutDil                respjson.Field
		ExtraFields                             map[string]respjson.Field
		raw                                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentIncomeStatement) RawJSON() string { return r.JSON.raw }
func (r *InstrumentIncomeStatement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstrumentIncomeStatementList []InstrumentIncomeStatement

type ActiveV1InstrumentIncomeStatementGetInstrumentIncomeStatementsResponse struct {
	Data InstrumentIncomeStatementList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1InstrumentIncomeStatementGetInstrumentIncomeStatementsResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ActiveV1InstrumentIncomeStatementGetInstrumentIncomeStatementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1InstrumentIncomeStatementGetInstrumentIncomeStatementsParams struct {
	// Security identifier source
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "CURRENCY", "FMP", "OEMS",
	// "SEDOL", "QUIK", "ISIN", "RIC", "COUNTRY", "EXCHANGE", "CTA", "BLOOMBERG",
	// "WERTPAPIER", "DUTCH", "VALOREN", "SICOVAM", "BELGIAN", "COMMON",
	// "CLEARING_HOUSE", "ISDA_FPML_SPECIFICATION", "ISDA_FPML_URL",
	// "LETTER_OF_CREDIT", "MARKETPLACE_ASSIGNED_IDENTIFIER", "MARKIT_RED_ENTITY_CLIP",
	// "MARKIT_RED_PAIR_CLIP", "CFTC", "ISDA_COMMODITY_REFERENCE_PRICE",
	// "LEGAL_ENTITY_IDENTIFIER", "SYNTHETIC", "FIDESSA_INSTRUMENT_MNEMONIC",
	// "INDEX_NAME", "UNIFORM_SYMBOL", "DIGITAL_TOKEN_IDENTIFIER", "MASSIVE", "OTHER".
	SecurityIDSource SecurityIDSource `path:"security_id_source,omitzero" api:"required" json:"-"`
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
// [ActiveV1InstrumentIncomeStatementGetInstrumentIncomeStatementsParams]'s query
// parameters as `url.Values`.
func (r ActiveV1InstrumentIncomeStatementGetInstrumentIncomeStatementsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
