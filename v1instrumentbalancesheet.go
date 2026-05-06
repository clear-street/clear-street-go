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
// V1InstrumentBalanceSheetService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1InstrumentBalanceSheetService] method instead.
type V1InstrumentBalanceSheetService struct {
	options []option.RequestOption
}

// NewV1InstrumentBalanceSheetService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1InstrumentBalanceSheetService(opts ...option.RequestOption) (r V1InstrumentBalanceSheetService) {
	r = V1InstrumentBalanceSheetService{}
	r.options = opts
	return
}

// Get balance sheet statements for an instrument.
//
// Retrieves quarterly balance sheet statements for a specific instrument, sorted
// by fiscal period (most recent first).
//
// Date range defaults:
//
// - `from_date`: None (no lower bound)
// - `to_date`: None (no upper bound)
func (r *V1InstrumentBalanceSheetService) GetInstrumentBalanceSheetStatements(ctx context.Context, instrumentID string, query V1InstrumentBalanceSheetGetInstrumentBalanceSheetStatementsParams, opts ...option.RequestOption) (res *V1InstrumentBalanceSheetGetInstrumentBalanceSheetStatementsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/instruments/%s/balance-sheets", instrumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A quarterly balance sheet statement for an instrument.
type InstrumentBalanceSheetStatement struct {
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
	// Account payables
	AccountPayables string `json:"account_payables" api:"nullable"`
	// Accounts receivables
	AccountsReceivables string `json:"accounts_receivables" api:"nullable"`
	// Accrued expenses
	AccruedExpenses string `json:"accrued_expenses" api:"nullable"`
	// Accumulated other comprehensive income/loss
	AccumulatedOtherComprehensiveIncomeLoss string `json:"accumulated_other_comprehensive_income_loss" api:"nullable"`
	// Additional paid-in capital
	AdditionalPaidInCapital string `json:"additional_paid_in_capital" api:"nullable"`
	// Capital lease obligations (total)
	CapitalLeaseObligations string `json:"capital_lease_obligations" api:"nullable"`
	// Capital lease obligations (current portion)
	CapitalLeaseObligationsCurrent string `json:"capital_lease_obligations_current" api:"nullable"`
	// Cash and cash equivalents
	CashAndCashEquivalents string `json:"cash_and_cash_equivalents" api:"nullable"`
	// Cash and short-term investments combined
	CashAndShortTermInvestments string `json:"cash_and_short_term_investments" api:"nullable"`
	// Common stock
	CommonStock string `json:"common_stock" api:"nullable"`
	// Deferred revenue
	DeferredRevenue string `json:"deferred_revenue" api:"nullable"`
	// Deferred revenue (non-current)
	DeferredRevenueNonCurrent string `json:"deferred_revenue_non_current" api:"nullable"`
	// Deferred tax liabilities (non-current)
	DeferredTaxLiabilitiesNonCurrent string `json:"deferred_tax_liabilities_non_current" api:"nullable"`
	// Goodwill
	Goodwill string `json:"goodwill" api:"nullable"`
	// Goodwill and intangible assets combined
	GoodwillAndIntangibleAssets string `json:"goodwill_and_intangible_assets" api:"nullable"`
	// Intangible assets
	IntangibleAssets string `json:"intangible_assets" api:"nullable"`
	// Inventory
	Inventory string `json:"inventory" api:"nullable"`
	// Long-term debt
	LongTermDebt string `json:"long_term_debt" api:"nullable"`
	// Long-term investments
	LongTermInvestments string `json:"long_term_investments" api:"nullable"`
	// Minority interest
	MinorityInterest string `json:"minority_interest" api:"nullable"`
	// Net debt (total debt minus cash)
	NetDebt string `json:"net_debt" api:"nullable"`
	// Net receivables
	NetReceivables string `json:"net_receivables" api:"nullable"`
	// Other assets
	OtherAssets string `json:"other_assets" api:"nullable"`
	// Other current assets
	OtherCurrentAssets string `json:"other_current_assets" api:"nullable"`
	// Other current liabilities
	OtherCurrentLiabilities string `json:"other_current_liabilities" api:"nullable"`
	// Other liabilities
	OtherLiabilities string `json:"other_liabilities" api:"nullable"`
	// Other non-current assets
	OtherNonCurrentAssets string `json:"other_non_current_assets" api:"nullable"`
	// Other non-current liabilities
	OtherNonCurrentLiabilities string `json:"other_non_current_liabilities" api:"nullable"`
	// Other payables
	OtherPayables string `json:"other_payables" api:"nullable"`
	// Other receivables
	OtherReceivables string `json:"other_receivables" api:"nullable"`
	// Other total stockholders equity
	OtherTotalStockholdersEquity string `json:"other_total_stockholders_equity" api:"nullable"`
	// Preferred stock
	PreferredStock string `json:"preferred_stock" api:"nullable"`
	// Prepaids
	Prepaids string `json:"prepaids" api:"nullable"`
	// Property, plant and equipment net of depreciation
	PropertyPlantAndEquipmentNet string `json:"property_plant_and_equipment_net" api:"nullable"`
	// Retained earnings
	RetainedEarnings string `json:"retained_earnings" api:"nullable"`
	// Short-term debt
	ShortTermDebt string `json:"short_term_debt" api:"nullable"`
	// Short-term investments
	ShortTermInvestments string `json:"short_term_investments" api:"nullable"`
	// Tax assets
	TaxAssets string `json:"tax_assets" api:"nullable"`
	// Tax payables
	TaxPayables string `json:"tax_payables" api:"nullable"`
	// Total assets
	TotalAssets string `json:"total_assets" api:"nullable"`
	// Total current assets
	TotalCurrentAssets string `json:"total_current_assets" api:"nullable"`
	// Total current liabilities
	TotalCurrentLiabilities string `json:"total_current_liabilities" api:"nullable"`
	// Total debt
	TotalDebt string `json:"total_debt" api:"nullable"`
	// Total equity
	TotalEquity string `json:"total_equity" api:"nullable"`
	// Total investments
	TotalInvestments string `json:"total_investments" api:"nullable"`
	// Total liabilities
	TotalLiabilities string `json:"total_liabilities" api:"nullable"`
	// Total liabilities and total equity
	TotalLiabilitiesAndTotalEquity string `json:"total_liabilities_and_total_equity" api:"nullable"`
	// Total non-current assets
	TotalNonCurrentAssets string `json:"total_non_current_assets" api:"nullable"`
	// Total non-current liabilities
	TotalNonCurrentLiabilities string `json:"total_non_current_liabilities" api:"nullable"`
	// Total payables
	TotalPayables string `json:"total_payables" api:"nullable"`
	// Total stockholders equity
	TotalStockholdersEquity string `json:"total_stockholders_equity" api:"nullable"`
	// Treasury stock
	TreasuryStock string `json:"treasury_stock" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptedDate                            respjson.Field
		FilingDate                              respjson.Field
		Period                                  respjson.Field
		PeriodType                              respjson.Field
		ReportedCurrency                        respjson.Field
		Year                                    respjson.Field
		AccountPayables                         respjson.Field
		AccountsReceivables                     respjson.Field
		AccruedExpenses                         respjson.Field
		AccumulatedOtherComprehensiveIncomeLoss respjson.Field
		AdditionalPaidInCapital                 respjson.Field
		CapitalLeaseObligations                 respjson.Field
		CapitalLeaseObligationsCurrent          respjson.Field
		CashAndCashEquivalents                  respjson.Field
		CashAndShortTermInvestments             respjson.Field
		CommonStock                             respjson.Field
		DeferredRevenue                         respjson.Field
		DeferredRevenueNonCurrent               respjson.Field
		DeferredTaxLiabilitiesNonCurrent        respjson.Field
		Goodwill                                respjson.Field
		GoodwillAndIntangibleAssets             respjson.Field
		IntangibleAssets                        respjson.Field
		Inventory                               respjson.Field
		LongTermDebt                            respjson.Field
		LongTermInvestments                     respjson.Field
		MinorityInterest                        respjson.Field
		NetDebt                                 respjson.Field
		NetReceivables                          respjson.Field
		OtherAssets                             respjson.Field
		OtherCurrentAssets                      respjson.Field
		OtherCurrentLiabilities                 respjson.Field
		OtherLiabilities                        respjson.Field
		OtherNonCurrentAssets                   respjson.Field
		OtherNonCurrentLiabilities              respjson.Field
		OtherPayables                           respjson.Field
		OtherReceivables                        respjson.Field
		OtherTotalStockholdersEquity            respjson.Field
		PreferredStock                          respjson.Field
		Prepaids                                respjson.Field
		PropertyPlantAndEquipmentNet            respjson.Field
		RetainedEarnings                        respjson.Field
		ShortTermDebt                           respjson.Field
		ShortTermInvestments                    respjson.Field
		TaxAssets                               respjson.Field
		TaxPayables                             respjson.Field
		TotalAssets                             respjson.Field
		TotalCurrentAssets                      respjson.Field
		TotalCurrentLiabilities                 respjson.Field
		TotalDebt                               respjson.Field
		TotalEquity                             respjson.Field
		TotalInvestments                        respjson.Field
		TotalLiabilities                        respjson.Field
		TotalLiabilitiesAndTotalEquity          respjson.Field
		TotalNonCurrentAssets                   respjson.Field
		TotalNonCurrentLiabilities              respjson.Field
		TotalPayables                           respjson.Field
		TotalStockholdersEquity                 respjson.Field
		TreasuryStock                           respjson.Field
		ExtraFields                             map[string]respjson.Field
		raw                                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentBalanceSheetStatement) RawJSON() string { return r.JSON.raw }
func (r *InstrumentBalanceSheetStatement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstrumentBalanceSheetStatementList []InstrumentBalanceSheetStatement

type V1InstrumentBalanceSheetGetInstrumentBalanceSheetStatementsResponse struct {
	Data InstrumentBalanceSheetStatementList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentBalanceSheetGetInstrumentBalanceSheetStatementsResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *V1InstrumentBalanceSheetGetInstrumentBalanceSheetStatementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentBalanceSheetGetInstrumentBalanceSheetStatementsParams struct {
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
// [V1InstrumentBalanceSheetGetInstrumentBalanceSheetStatementsParams]'s query
// parameters as `url.Values`.
func (r V1InstrumentBalanceSheetGetInstrumentBalanceSheetStatementsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
