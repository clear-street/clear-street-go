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
// V1InstrumentService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1InstrumentService] method instead.
type V1InstrumentService struct {
	options []option.RequestOption
	// Retrieve details and lists of tradable instruments.
	AnalystReporting V1InstrumentAnalystReportingService
	// Retrieve details and lists of tradable instruments.
	BalanceSheets V1InstrumentBalanceSheetService
	// Retrieve details and lists of tradable instruments.
	CashFlowStatements V1InstrumentCashFlowStatementService
	// Retrieve details and lists of tradable instruments.
	Events V1InstrumentEventService
	// Retrieve details and lists of tradable instruments.
	Fundamentals V1InstrumentFundamentalService
	// Retrieve details and lists of tradable instruments.
	IncomeStatements V1InstrumentIncomeStatementService
	// Retrieve details and lists of tradable instruments.
	Options V1InstrumentOptionService
}

// NewV1InstrumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1InstrumentService(opts ...option.RequestOption) (r V1InstrumentService) {
	r = V1InstrumentService{}
	r.options = opts
	r.AnalystReporting = NewV1InstrumentAnalystReportingService(opts...)
	r.BalanceSheets = NewV1InstrumentBalanceSheetService(opts...)
	r.CashFlowStatements = NewV1InstrumentCashFlowStatementService(opts...)
	r.Events = NewV1InstrumentEventService(opts...)
	r.Fundamentals = NewV1InstrumentFundamentalService(opts...)
	r.IncomeStatements = NewV1InstrumentIncomeStatementService(opts...)
	r.Options = NewV1InstrumentOptionService(opts...)
	return
}

// Retrieves detailed information for a specific instrument.
func (r *V1InstrumentService) GetInstrumentByID(ctx context.Context, instrumentID InstrumentIDOrSymbol, query V1InstrumentGetInstrumentByIDParams, opts ...option.RequestOption) (res *V1InstrumentGetInstrumentByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/instruments/%s", instrumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves a list of tradeable instruments.
func (r *V1InstrumentService) GetInstruments(ctx context.Context, query V1InstrumentGetInstrumentsParams, opts ...option.RequestOption) (res *V1InstrumentGetInstrumentsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/instruments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search instruments by symbol, alternate identifier, or company name.
//
// The `q` parameter is case-insensitive and supports ticker symbols, alternate
// identifiers such as CUSIP, ISIN, OPRA root, and CMS identifiers, and company
// names for non-option instruments. Results are ranked by match quality plus
// instrument quality signals including log-scaled ADV, listing status,
// marginability, easy-to-borrow status, and OTC, restricted, and liquidation-only
// penalties. Defaults to the `EQUITY` asset class (common stocks, preferred
// shares, ADRs, ETFs, and exchange-traded mutual funds). Pass `asset_class=OPTION`
// to search option contracts by symbol or alternate identifier.
func (r *V1InstrumentService) SearchInstruments(ctx context.Context, query V1InstrumentSearchInstrumentsParams, opts ...option.RequestOption) (res *V1InstrumentSearchInstrumentsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/instruments/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
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

// The type of options contract
type ContractType string

const (
	ContractTypeCall ContractType = "CALL"
	ContractTypePut  ContractType = "PUT"
)

// The exercise style of an options contract
type ExerciseStyle string

const (
	ExerciseStyleAmerican ExerciseStyle = "AMERICAN"
	ExerciseStyleEuropean ExerciseStyle = "EUROPEAN"
)

// Fiscal period type for earnings reports
type FiscalPeriodType string

const (
	FiscalPeriodTypeQuarterly FiscalPeriodType = "QUARTERLY"
	FiscalPeriodTypeAnnual    FiscalPeriodType = "ANNUAL"
	FiscalPeriodTypeTtm       FiscalPeriodType = "TTM"
	FiscalPeriodTypeBiannual  FiscalPeriodType = "BIANNUAL"
)

// Represents a tradable financial instrument.
type Instrument struct {
	// Available options expiration dates for this instrument. Present only when
	// `include_options_expiry_dates=true` in the request.
	OptionsExpiryDates []time.Time `json:"options_expiry_dates" api:"nullable" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OptionsExpiryDates respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
	InstrumentCore
}

// Returns the unmodified JSON received from the API
func (r Instrument) RawJSON() string { return r.JSON.raw }
func (r *Instrument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstrumentCore struct {
	// Unique OEMS instrument identifier (UUID)
	ID string `json:"id" api:"required" format:"uuid"`
	// The ISO country code of the instrument's issue
	CountryOfIssue string `json:"country_of_issue" api:"required"`
	// The ISO currency code in which the instrument is traded
	Currency string `json:"currency" api:"required"`
	// Indicates if the instrument is classified as Easy-To-Borrow
	EasyToBorrow bool `json:"easy_to_borrow" api:"required"`
	// Indicates if the instrument is liquidation only and cannot be bought
	IsLiquidationOnly bool `json:"is_liquidation_only" api:"required"`
	// Indicates if the instrument is marginable
	IsMarginable bool `json:"is_marginable" api:"required"`
	// Indicates if the instrument is restricted from trading
	IsRestricted bool `json:"is_restricted" api:"required"`
	// Indicates if short selling is prohibited for the instrument
	IsShortProhibited bool `json:"is_short_prohibited" api:"required"`
	// Indicates if the instrument is on the Regulation SHO Threshold Security List
	IsThresholdSecurity bool `json:"is_threshold_security" api:"required"`
	// Indicates if the instrument is tradable
	IsTradable bool `json:"is_tradable" api:"required"`
	// The trading symbol for the instrument
	Symbol string `json:"symbol" api:"required"`
	// The MIC code of the primary listing venue
	Venue string `json:"venue" api:"required"`
	// Average daily share volume from the security definition.
	Adv string `json:"adv" api:"nullable"`
	// The expiration date for options instruments
	Expiry time.Time `json:"expiry" api:"nullable" format:"date"`
	// The type of security (e.g., Common Stock, ETF)
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "OPTION", "CASH", "OTHER".
	InstrumentType SecurityType `json:"instrument_type" api:"nullable"`
	// The percent of a long position's value you must post as margin
	LongMarginRate string `json:"long_margin_rate" api:"nullable"`
	// The full name of the instrument or its issuer
	Name string `json:"name" api:"nullable"`
	// Notional ADV (`adv × previous_close`). The primary liquidity signal used by
	// `/instruments/search` ranking. Computed at response time so it stays consistent
	// with whatever `adv` and `previous_close` show.
	NotionalAdv string `json:"notional_adv" api:"nullable"`
	// Last close price from the security definition.
	PreviousClose string `json:"previous_close" api:"nullable"`
	// The percent of a short position's value you must post as margin
	ShortMarginRate string `json:"short_margin_rate" api:"nullable"`
	// The strike price for options instruments
	StrikePrice string `json:"strike_price" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CountryOfIssue      respjson.Field
		Currency            respjson.Field
		EasyToBorrow        respjson.Field
		IsLiquidationOnly   respjson.Field
		IsMarginable        respjson.Field
		IsRestricted        respjson.Field
		IsShortProhibited   respjson.Field
		IsThresholdSecurity respjson.Field
		IsTradable          respjson.Field
		Symbol              respjson.Field
		Venue               respjson.Field
		Adv                 respjson.Field
		Expiry              respjson.Field
		InstrumentType      respjson.Field
		LongMarginRate      respjson.Field
		Name                respjson.Field
		NotionalAdv         respjson.Field
		PreviousClose       respjson.Field
		ShortMarginRate     respjson.Field
		StrikePrice         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentCore) RawJSON() string { return r.JSON.raw }
func (r *InstrumentCore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstrumentCoreList []InstrumentCore

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

// The listing type of an options contract
type ListingType string

const (
	ListingTypeStandard ListingType = "STANDARD"
	ListingTypeFlex     ListingType = "FLEX"
	ListingTypeOtc      ListingType = "OTC"
)

// An options contract with options-specific metadata
type OptionsContract struct {
	// OEMS instrument identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether this is a CALL or PUT
	//
	// Any of "CALL", "PUT".
	ContractType ContractType `json:"contract_type" api:"required"`
	// ISO currency code
	Currency string `json:"currency" api:"required"`
	// MIC code of the primary listing venue
	Exchange string `json:"exchange" api:"required"`
	// Exercise style
	//
	// Any of "AMERICAN", "EUROPEAN".
	ExerciseStyle ExerciseStyle `json:"exercise_style" api:"required"`
	// Expiration date
	Expiry time.Time `json:"expiry" api:"required" format:"date"`
	// Whether the contract is liquidation-only
	IsLiquidationOnly bool `json:"is_liquidation_only" api:"required"`
	// Whether the contract is marginable
	IsMarginable bool `json:"is_marginable" api:"required"`
	// Whether the contract is restricted from trading
	IsRestricted bool `json:"is_restricted" api:"required"`
	// Listing type
	//
	// Any of "STANDARD", "FLEX", "OTC".
	ListingType ListingType `json:"listing_type" api:"required"`
	// Contract multiplier (100 for standard options)
	Multiplier string `json:"multiplier" api:"required"`
	// Strike price
	StrikePrice string `json:"strike_price" api:"required"`
	// OSI symbol (e.g. "AAPL 251219C00150000")
	Symbol string `json:"symbol" api:"required"`
	// Open interest (number of outstanding contracts), if available
	OpenInterest int64 `json:"open_interest" api:"nullable"`
	// OEMS instrument ID of the underlying instrument, if resolvable
	UnderlyingInstrumentID string `json:"underlying_instrument_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		ContractType           respjson.Field
		Currency               respjson.Field
		Exchange               respjson.Field
		ExerciseStyle          respjson.Field
		Expiry                 respjson.Field
		IsLiquidationOnly      respjson.Field
		IsMarginable           respjson.Field
		IsRestricted           respjson.Field
		ListingType            respjson.Field
		Multiplier             respjson.Field
		StrikePrice            respjson.Field
		Symbol                 respjson.Field
		OpenInterest           respjson.Field
		UnderlyingInstrumentID respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptionsContract) RawJSON() string { return r.JSON.raw }
func (r *OptionsContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptionsContractList []OptionsContract

type V1InstrumentGetInstrumentByIDResponse struct {
	// Represents a tradable financial instrument.
	Data Instrument `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentGetInstrumentByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentGetInstrumentByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentGetInstrumentsResponse struct {
	Data InstrumentCoreList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentGetInstrumentsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentGetInstrumentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentSearchInstrumentsResponse struct {
	Data InstrumentCoreList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentSearchInstrumentsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentSearchInstrumentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentGetInstrumentByIDParams struct {
	// When true, include unique options expiry dates for this instrument
	IncludeOptionsExpiryDates param.Opt[bool] `query:"include_options_expiry_dates,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentGetInstrumentByIDParams]'s query parameters as
// `url.Values`.
func (r V1InstrumentGetInstrumentByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1InstrumentGetInstrumentsParams struct {
	// Filter by easy to borrow status
	EasyToBorrow param.Opt[bool] `query:"easy_to_borrow,omitzero" json:"-"`
	// Filter IDs to those containing this substring. For options, and when
	// instrument_type is omitted and no instrument_ids filters are provided, this is
	// required.
	IDFilter param.Opt[string] `query:"id_filter,omitzero" json:"-"`
	// Filter by liquidation only status
	IsLiquidationOnly param.Opt[bool] `query:"is_liquidation_only,omitzero" json:"-"`
	// Filter by marginable status
	IsMarginable param.Opt[bool] `query:"is_marginable,omitzero" json:"-"`
	// Filter by restricted status
	IsRestricted param.Opt[bool] `query:"is_restricted,omitzero" json:"-"`
	// Filter by short prohibited status
	IsShortProhibited param.Opt[bool] `query:"is_short_prohibited,omitzero" json:"-"`
	// Filter by threshold security status
	IsThresholdSecurity param.Opt[bool]  `query:"is_threshold_security,omitzero" json:"-"`
	PageSize            param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Comma-separated OEMS instrument UUIDs
	InstrumentIDs []string `query:"instrument_ids,omitzero" format:"uuid" json:"-"`
	// Filter by instrument type. If omitted, returns all supported instrument types.
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "OPTION", "CASH", "OTHER".
	InstrumentType V1InstrumentGetInstrumentsParamsInstrumentType `query:"instrument_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentGetInstrumentsParams]'s query parameters as
// `url.Values`.
func (r V1InstrumentGetInstrumentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by instrument type. If omitted, returns all supported instrument types.
type V1InstrumentGetInstrumentsParamsInstrumentType string

const (
	V1InstrumentGetInstrumentsParamsInstrumentTypeCommonStock    V1InstrumentGetInstrumentsParamsInstrumentType = "COMMON_STOCK"
	V1InstrumentGetInstrumentsParamsInstrumentTypePreferredStock V1InstrumentGetInstrumentsParamsInstrumentType = "PREFERRED_STOCK"
	V1InstrumentGetInstrumentsParamsInstrumentTypeOption         V1InstrumentGetInstrumentsParamsInstrumentType = "OPTION"
	V1InstrumentGetInstrumentsParamsInstrumentTypeCash           V1InstrumentGetInstrumentsParamsInstrumentType = "CASH"
	V1InstrumentGetInstrumentsParamsInstrumentTypeOther          V1InstrumentGetInstrumentsParamsInstrumentType = "OTHER"
)

type V1InstrumentSearchInstrumentsParams struct {
	// Search term applied case-insensitively to ticker symbols, alternate identifiers
	// (CUSIP, ISIN, OPRA root, CMS), and company names for non-option instruments.
	// Option searches match symbols and alternate identifiers.
	Q string `query:"q" api:"required" json:"-"`
	// Comma-separated asset classes (EQUITY|OPTION|WARRANT|BOND|FX|OTHER). Defaults to
	// EQUITY.
	AssetClass param.Opt[string] `query:"asset_class,omitzero" json:"-"`
	// Optional listing-country filter (e.g., US).
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// Optional ISO currency filter (e.g., USD).
	Currency param.Opt[string] `query:"currency,omitzero" json:"-"`
	// Include inactive instruments. Default false.
	IncludeInactive param.Opt[bool] `query:"include_inactive,omitzero" json:"-"`
	// Include restricted instruments. Default true (penalized in ranking).
	IncludeRestricted param.Opt[bool]  `query:"include_restricted,omitzero" json:"-"`
	PageSize          param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentSearchInstrumentsParams]'s query parameters as
// `url.Values`.
func (r V1InstrumentSearchInstrumentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
