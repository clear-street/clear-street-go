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

// Retrieve instrument analytics, market data, news, and related reference data.
//
// V1InstrumentDataService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1InstrumentDataService] method instead.
type V1InstrumentDataService struct {
	options []option.RequestOption
	// Retrieve instrument analytics, market data, news, and related reference data.
	MarketData V1InstrumentDataMarketDataService
	// Retrieve instrument analytics, market data, news, and related reference data.
	News V1InstrumentDataNewsService
}

// NewV1InstrumentDataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1InstrumentDataService(opts ...option.RequestOption) (r V1InstrumentDataService) {
	r = V1InstrumentDataService{}
	r.options = opts
	r.MarketData = NewV1InstrumentDataMarketDataService(opts...)
	r.News = NewV1InstrumentDataNewsService(opts...)
	return
}

// List instrument events across all securities.
//
// Retrieves all instrument events grouped by date.
func (r *V1InstrumentDataService) GetAllInstrumentEvents(ctx context.Context, query V1InstrumentDataGetAllInstrumentEventsParams, opts ...option.RequestOption) (res *V1InstrumentDataGetAllInstrumentEventsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/instruments/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves analyst ratings and price targets for an instrument.
func (r *V1InstrumentDataService) GetInstrumentAnalystConsensus(ctx context.Context, instrumentID InstrumentIDOrSymbol, query V1InstrumentDataGetInstrumentAnalystConsensusParams, opts ...option.RequestOption) (res *V1InstrumentDataGetInstrumentAnalystConsensusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/instruments/%s/analyst-reporting", instrumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
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
func (r *V1InstrumentDataService) GetInstrumentBalanceSheetStatements(ctx context.Context, instrumentID InstrumentIDOrSymbol, query V1InstrumentDataGetInstrumentBalanceSheetStatementsParams, opts ...option.RequestOption) (res *V1InstrumentDataGetInstrumentBalanceSheetStatementsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/instruments/%s/balance-sheets", instrumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get cash flow statements for an instrument.
//
// Retrieves historical cash flow statements for the specified instrument. Cash
// flow statements show cash inflows and outflows from operating, investing, and
// financing activities.
func (r *V1InstrumentDataService) GetInstrumentCashFlowStatements(ctx context.Context, instrumentID InstrumentIDOrSymbol, query V1InstrumentDataGetInstrumentCashFlowStatementsParams, opts ...option.RequestOption) (res *V1InstrumentDataGetInstrumentCashFlowStatementsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/instruments/%s/cash-flow-statements", instrumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves corporate events (dividends, splits, etc.) for an instrument, grouped
// by event type.
//
// Date range defaults:
//
// - `from_date`: today - 365 days
// - `to_date`: today + 60 days
func (r *V1InstrumentDataService) GetInstrumentEvents(ctx context.Context, instrumentID InstrumentIDOrSymbol, query V1InstrumentDataGetInstrumentEventsParams, opts ...option.RequestOption) (res *V1InstrumentDataGetInstrumentEventsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/instruments/%s/events", instrumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves supplemental fundamentals and company profile data for an instrument.
func (r *V1InstrumentDataService) GetInstrumentFundamentals(ctx context.Context, instrumentID InstrumentIDOrSymbol, opts ...option.RequestOption) (res *V1InstrumentDataGetInstrumentFundamentalsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/instruments/%s/fundamentals", instrumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves quarterly income statements for a specific instrument, sorted by
// fiscal period (most recent first).
//
// Date range defaults:
//
// - `from_date`: None (no lower bound)
// - `to_date`: None (no upper bound)
func (r *V1InstrumentDataService) GetInstrumentIncomeStatements(ctx context.Context, instrumentID InstrumentIDOrSymbol, query V1InstrumentDataGetInstrumentIncomeStatementsParams, opts ...option.RequestOption) (res *V1InstrumentDataGetInstrumentIncomeStatementsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/instruments/%s/income-statements", instrumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Event types supported by the all-events endpoint.
type AllEventsEventType string

const (
	AllEventsEventTypeEarnings   AllEventsEventType = "EARNINGS"
	AllEventsEventTypeDividend   AllEventsEventType = "DIVIDEND"
	AllEventsEventTypeStockSplit AllEventsEventType = "STOCK_SPLIT"
	AllEventsEventTypeIpo        AllEventsEventType = "IPO"
)

// Analyst recommendation distribution
type AnalystDistribution struct {
	// Number of buy recommendations
	Buy int64 `json:"buy" api:"required"`
	// Number of hold recommendations
	Hold int64 `json:"hold" api:"required"`
	// Number of sell recommendations
	Sell int64 `json:"sell" api:"required"`
	// Number of strong buy recommendations
	StrongBuy int64 `json:"strong_buy" api:"required"`
	// Number of strong sell recommendations
	StrongSell int64 `json:"strong_sell" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buy         respjson.Field
		Hold        respjson.Field
		Sell        respjson.Field
		StrongBuy   respjson.Field
		StrongSell  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalystDistribution) RawJSON() string { return r.JSON.raw }
func (r *AnalystDistribution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Analyst rating category
type AnalystRating string

const (
	AnalystRatingStrongBuy  AnalystRating = "STRONG_BUY"
	AnalystRatingBuy        AnalystRating = "BUY"
	AnalystRatingHold       AnalystRating = "HOLD"
	AnalystRatingSell       AnalystRating = "SELL"
	AnalystRatingStrongSell AnalystRating = "STRONG_SELL"
)

// Fiscal period type for earnings reports
type FiscalPeriodType string

const (
	FiscalPeriodTypeQuarterly FiscalPeriodType = "QUARTERLY"
	FiscalPeriodTypeAnnual    FiscalPeriodType = "ANNUAL"
	FiscalPeriodTypeTtm       FiscalPeriodType = "TTM"
	FiscalPeriodTypeBiannual  FiscalPeriodType = "BIANNUAL"
)

// All-events payload grouped by date.
type InstrumentAllEventsData struct {
	// Events grouped by date in descending order.
	EventDates []InstrumentEventsByDate `json:"event_dates" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventDates  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentAllEventsData) RawJSON() string { return r.JSON.raw }
func (r *InstrumentAllEventsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Aggregated analyst consensus metrics
type InstrumentAnalystConsensus struct {
	// The date the consensus snapshot was generated
	Date time.Time `json:"date" api:"required" format:"date"`
	// Count of individual analyst recommendations by category
	Distribution AnalystDistribution `json:"distribution" api:"nullable"`
	// Aggregated analyst price target statistics
	PriceTarget PriceTarget `json:"price_target" api:"nullable"`
	// Consensus analyst rating
	//
	// Any of "STRONG_BUY", "BUY", "HOLD", "SELL", "STRONG_SELL".
	Rating AnalystRating `json:"rating" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date         respjson.Field
		Distribution respjson.Field
		PriceTarget  respjson.Field
		Rating       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentAnalystConsensus) RawJSON() string { return r.JSON.raw }
func (r *InstrumentAnalystConsensus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

// Represents a dividend event for an instrument
type InstrumentDividendEvent struct {
	// The adjusted dividend amount accounting for any splits.
	AdjustedDividendAmount string `json:"adjusted_dividend_amount" api:"required"`
	// The day the stock starts trading without the right to receive that dividend.
	ExDate time.Time `json:"ex_date" api:"required" format:"date"`
	// The declaration date of the dividend
	DeclarationDate time.Time `json:"declaration_date" api:"nullable" format:"date"`
	// The dividend amount per share.
	DividendAmount string `json:"dividend_amount" api:"nullable"`
	// The dividend yield as a percentage of the stock price.
	DividendYield string `json:"dividend_yield" api:"nullable"`
	// The frequency of the dividend payments (e.g., "Quarterly", "Annual").
	Frequency string `json:"frequency" api:"nullable"`
	// The payment date is the date on which a declared stock dividend is scheduled to
	// be paid.
	PaymentDate time.Time `json:"payment_date" api:"nullable" format:"date"`
	// The record date, set by a company's board of directors, is when a company
	// compiles a list of shareholders of the stock for which it has declared a
	// dividend.
	RecordDate time.Time `json:"record_date" api:"nullable" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdjustedDividendAmount respjson.Field
		ExDate                 respjson.Field
		DeclarationDate        respjson.Field
		DividendAmount         respjson.Field
		DividendYield          respjson.Field
		Frequency              respjson.Field
		PaymentDate            respjson.Field
		RecordDate             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentDividendEvent) RawJSON() string { return r.JSON.raw }
func (r *InstrumentDividendEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents instrument earnings data
type InstrumentEarnings struct {
	// The date when the earnings report was published
	Date time.Time `json:"date" api:"required" format:"date"`
	// The actual earnings per share (EPS) for the period
	EpsActual string `json:"eps_actual" api:"nullable"`
	// The estimated earnings per share (EPS) for the period
	EpsEstimate string `json:"eps_estimate" api:"nullable"`
	// The percentage difference between actual and estimated EPS
	EpsSurprisePercent string `json:"eps_surprise_percent" api:"nullable"`
	// The actual total revenue for the period
	RevenueActual string `json:"revenue_actual" api:"nullable"`
	// The estimated total revenue for the period
	RevenueEstimate string `json:"revenue_estimate" api:"nullable"`
	// The percentage difference between actual and estimated revenue
	RevenueSurprisePercent string `json:"revenue_surprise_percent" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date                   respjson.Field
		EpsActual              respjson.Field
		EpsEstimate            respjson.Field
		EpsSurprisePercent     respjson.Field
		RevenueActual          respjson.Field
		RevenueEstimate        respjson.Field
		RevenueSurprisePercent respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentEarnings) RawJSON() string { return r.JSON.raw }
func (r *InstrumentEarnings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unified envelope for the all-events response.
type InstrumentEventEnvelope struct {
	// Symbol associated with the event.
	Symbol string `json:"symbol" api:"required"`
	// Event type discriminator.
	//
	// Any of "EARNINGS", "DIVIDEND", "STOCK_SPLIT", "IPO".
	Type AllEventsEventType `json:"type" api:"required"`
	// Dividend payload when type is DIVIDEND.
	DividendEventData InstrumentDividendEvent `json:"dividend_event_data" api:"nullable"`
	// Earnings payload when type is EARNINGS.
	EarningsEventData InstrumentEarnings `json:"earnings_event_data" api:"nullable"`
	// OEMS instrument identifier, when the instrument is found in the instrument
	// cache.
	InstrumentID string `json:"instrument_id" api:"nullable" format:"uuid"`
	// IPO payload when type is IPO.
	IpoEventData InstrumentEventIpoItem `json:"ipo_event_data" api:"nullable"`
	// Instrument name associated with the event, when available.
	Name string `json:"name" api:"nullable"`
	// Stock split payload when type is STOCK_SPLIT.
	StockSplitEventData InstrumentSplitEvent `json:"stock_split_event_data" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Symbol              respjson.Field
		Type                respjson.Field
		DividendEventData   respjson.Field
		EarningsEventData   respjson.Field
		InstrumentID        respjson.Field
		IpoEventData        respjson.Field
		Name                respjson.Field
		StockSplitEventData respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentEventEnvelope) RawJSON() string { return r.JSON.raw }
func (r *InstrumentEventEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IPO event in the all-events date grouping response.
type InstrumentEventIpoItem struct {
	// IPO action.
	Actions string `json:"actions" api:"nullable"`
	// IPO announced timestamp.
	AnnouncedAt time.Time `json:"announced_at" api:"nullable" format:"date-time"`
	// IPO company name.
	Company string `json:"company" api:"nullable"`
	// IPO exchange.
	Exchange string `json:"exchange" api:"nullable"`
	// IPO market cap.
	MarketCap string `json:"market_cap" api:"nullable"`
	// IPO price range.
	PriceRange string `json:"price_range" api:"nullable"`
	// IPO shares offered.
	Shares string `json:"shares" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actions     respjson.Field
		AnnouncedAt respjson.Field
		Company     respjson.Field
		Exchange    respjson.Field
		MarketCap   respjson.Field
		PriceRange  respjson.Field
		Shares      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentEventIpoItem) RawJSON() string { return r.JSON.raw }
func (r *InstrumentEventIpoItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instrument events for a single date.
type InstrumentEventsByDate struct {
	// Event date.
	Date time.Time `json:"date" api:"required" format:"date"`
	// Flat event envelopes for this date.
	Events []InstrumentEventEnvelope `json:"events" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Events      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentEventsByDate) RawJSON() string { return r.JSON.raw }
func (r *InstrumentEventsByDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Grouped instrument events by type
type InstrumentEventsData struct {
	// Dividend distribution events
	Dividends []InstrumentDividendEvent `json:"dividends" api:"required"`
	// Earnings announcement events
	Earnings []InstrumentEarnings `json:"earnings" api:"required"`
	// OEMS instrument UUID from the request
	InstrumentID string `json:"instrument_id" api:"required" format:"uuid"`
	// Stock split events
	Splits []InstrumentSplitEvent `json:"splits" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dividends    respjson.Field
		Earnings     respjson.Field
		InstrumentID respjson.Field
		Splits       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentEventsData) RawJSON() string { return r.JSON.raw }
func (r *InstrumentEventsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Supplemental fundamentals and company profile data for an instrument.
type InstrumentFundamentals struct {
	// The average daily trading volume over the past 30 days
	AverageVolume int64 `json:"average_volume" api:"nullable"`
	// The beta value, measuring the instrument's volatility relative to the overall
	// market
	Beta string `json:"beta" api:"nullable"`
	// A detailed description of the instrument or company
	Description string `json:"description" api:"nullable"`
	// The trailing twelve months (TTM) dividend yield
	DividendYield string `json:"dividend_yield" api:"nullable"`
	// The trailing twelve months (TTM) earnings per share
	EarningsPerShare string `json:"earnings_per_share" api:"nullable"`
	// The highest price over the last 52 weeks
	FiftyTwoWeekHigh string `json:"fifty_two_week_high" api:"nullable"`
	// The lowest price over the last 52 weeks
	FiftyTwoWeekLow string `json:"fifty_two_week_low" api:"nullable"`
	// The specific industry of the instrument's issuer
	Industry string `json:"industry" api:"nullable"`
	// The date the instrument was first listed
	ListDate time.Time `json:"list_date" api:"nullable" format:"date"`
	// URL to a representative logo image for the instrument or issuer
	LogoURL string `json:"logo_url" api:"nullable"`
	// The total market capitalization
	MarketCap string `json:"market_cap" api:"nullable"`
	// The closing price from the previous trading day
	PreviousClose string `json:"previous_close" api:"nullable"`
	// The price-to-earnings (P/E) ratio for the trailing twelve months (TTM)
	PriceToEarnings string `json:"price_to_earnings" api:"nullable"`
	// The business sector of the instrument's issuer
	Sector string `json:"sector" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AverageVolume    respjson.Field
		Beta             respjson.Field
		Description      respjson.Field
		DividendYield    respjson.Field
		EarningsPerShare respjson.Field
		FiftyTwoWeekHigh respjson.Field
		FiftyTwoWeekLow  respjson.Field
		Industry         respjson.Field
		ListDate         respjson.Field
		LogoURL          respjson.Field
		MarketCap        respjson.Field
		PreviousClose    respjson.Field
		PriceToEarnings  respjson.Field
		Sector           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentFundamentals) RawJSON() string { return r.JSON.raw }
func (r *InstrumentFundamentals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

// Represents a stock split event for an instrument
type InstrumentSplitEvent struct {
	// The date of the stock split
	Date time.Time `json:"date" api:"required" format:"date"`
	// The denominator of the split ratio
	Denominator string `json:"denominator" api:"required"`
	// The numerator of the split ratio
	Numerator string `json:"numerator" api:"required"`
	// The type of stock split (e.g., "stock-split", "stock-dividend", "bonus-issue")
	SplitType string `json:"split_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Denominator respjson.Field
		Numerator   respjson.Field
		SplitType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentSplitEvent) RawJSON() string { return r.JSON.raw }
func (r *InstrumentSplitEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Analyst price target statistics
type PriceTarget struct {
	// Average analyst price target
	Average string `json:"average" api:"required"`
	// ISO 4217 currency code of the price targets
	Currency string `json:"currency" api:"required"`
	// Highest analyst price target
	High string `json:"high" api:"required"`
	// Lowest analyst price target
	Low string `json:"low" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Average     respjson.Field
		Currency    respjson.Field
		High        respjson.Field
		Low         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PriceTarget) RawJSON() string { return r.JSON.raw }
func (r *PriceTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentDataGetAllInstrumentEventsResponse struct {
	// All-events payload grouped by date.
	Data InstrumentAllEventsData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentDataGetAllInstrumentEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentDataGetAllInstrumentEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentDataGetInstrumentAnalystConsensusResponse struct {
	// Aggregated analyst consensus metrics
	Data InstrumentAnalystConsensus `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentDataGetInstrumentAnalystConsensusResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentDataGetInstrumentAnalystConsensusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentDataGetInstrumentBalanceSheetStatementsResponse struct {
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
func (r V1InstrumentDataGetInstrumentBalanceSheetStatementsResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *V1InstrumentDataGetInstrumentBalanceSheetStatementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentDataGetInstrumentCashFlowStatementsResponse struct {
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
func (r V1InstrumentDataGetInstrumentCashFlowStatementsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentDataGetInstrumentCashFlowStatementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentDataGetInstrumentEventsResponse struct {
	// Grouped instrument events by type
	Data InstrumentEventsData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentDataGetInstrumentEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentDataGetInstrumentEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentDataGetInstrumentFundamentalsResponse struct {
	// Supplemental fundamentals and company profile data for an instrument.
	Data InstrumentFundamentals `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentDataGetInstrumentFundamentalsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentDataGetInstrumentFundamentalsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentDataGetInstrumentIncomeStatementsResponse struct {
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
func (r V1InstrumentDataGetInstrumentIncomeStatementsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentDataGetInstrumentIncomeStatementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentDataGetAllInstrumentEventsParams struct {
	// The start date for the query range, inclusive (YYYY-MM-DD).
	FromDate param.Opt[string] `query:"from_date,omitzero" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD).
	ToDate param.Opt[string] `query:"to_date,omitzero" json:"-"`
	// Filter by event type(s). Comma-delimited list. Example:
	// `event_types=EARNINGS,IPO`.
	EventTypes []AllEventsEventType `query:"event_types,omitzero" json:"-"`
	// Filter by OEMS instrument ID(s). Comma-delimited list of UUIDs. Example:
	// `instrument_ids=550e8400-e29b-41d4-a716-446655440000`.
	InstrumentIDs []string `query:"instrument_ids,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentDataGetAllInstrumentEventsParams]'s query
// parameters as `url.Values`.
func (r V1InstrumentDataGetAllInstrumentEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1InstrumentDataGetInstrumentAnalystConsensusParams struct {
	// The start date for the query range, inclusive (YYYY-MM-DD)
	From param.Opt[time.Time] `query:"from,omitzero" format:"date" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD)
	To param.Opt[time.Time] `query:"to,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentDataGetInstrumentAnalystConsensusParams]'s
// query parameters as `url.Values`.
func (r V1InstrumentDataGetInstrumentAnalystConsensusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1InstrumentDataGetInstrumentBalanceSheetStatementsParams struct {
	// The start date for the query range, inclusive (YYYY-MM-DD).
	FromDate param.Opt[string] `query:"from_date,omitzero" json:"-"`
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD).
	ToDate param.Opt[string] `query:"to_date,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [V1InstrumentDataGetInstrumentBalanceSheetStatementsParams]'s query parameters
// as `url.Values`.
func (r V1InstrumentDataGetInstrumentBalanceSheetStatementsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1InstrumentDataGetInstrumentCashFlowStatementsParams struct {
	// The start date for the query range, inclusive (YYYY-MM-DD).
	FromDate param.Opt[string] `query:"from_date,omitzero" json:"-"`
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD).
	ToDate param.Opt[string] `query:"to_date,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentDataGetInstrumentCashFlowStatementsParams]'s
// query parameters as `url.Values`.
func (r V1InstrumentDataGetInstrumentCashFlowStatementsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1InstrumentDataGetInstrumentEventsParams struct {
	// The start date for the query range, inclusive (YYYY-MM-DD).
	FromDate param.Opt[string] `query:"from_date,omitzero" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD).
	ToDate param.Opt[string] `query:"to_date,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentDataGetInstrumentEventsParams]'s query
// parameters as `url.Values`.
func (r V1InstrumentDataGetInstrumentEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1InstrumentDataGetInstrumentIncomeStatementsParams struct {
	// The start date for the query range, inclusive (YYYY-MM-DD).
	FromDate param.Opt[string] `query:"from_date,omitzero" json:"-"`
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD).
	ToDate param.Opt[string] `query:"to_date,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentDataGetInstrumentIncomeStatementsParams]'s
// query parameters as `url.Values`.
func (r V1InstrumentDataGetInstrumentIncomeStatementsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
