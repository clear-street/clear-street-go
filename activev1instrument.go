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

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/apiquery"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/param"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// ActiveV1InstrumentService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1InstrumentService] method instead.
type ActiveV1InstrumentService struct {
	Options          []option.RequestOption
	AnalystReporting ActiveV1InstrumentAnalystReportingService
	Events           ActiveV1InstrumentEventService
	Reporting        ActiveV1InstrumentReportingService
	Venues           ActiveV1InstrumentVenueService
}

// NewActiveV1InstrumentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1InstrumentService(opts ...option.RequestOption) (r ActiveV1InstrumentService) {
	r = ActiveV1InstrumentService{}
	r.Options = opts
	r.AnalystReporting = NewActiveV1InstrumentAnalystReportingService(opts...)
	r.Events = NewActiveV1InstrumentEventService(opts...)
	r.Reporting = NewActiveV1InstrumentReportingService(opts...)
	r.Venues = NewActiveV1InstrumentVenueService(opts...)
	return
}

// Retrieves detailed information for a specific instrument.
func (r *ActiveV1InstrumentService) GetInstrumentByID(ctx context.Context, securityID string, query ActiveV1InstrumentGetInstrumentByIDParams, opts ...option.RequestOption) (res *ActiveV1InstrumentGetInstrumentByIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if securityID == "" {
		err = errors.New("missing required security_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/instruments/%v/%s", query.SecurityIDSource, url.PathEscape(securityID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves a list of tradeable instruments.
func (r *ActiveV1InstrumentService) GetInstruments(ctx context.Context, query ActiveV1InstrumentGetInstrumentsParams, opts ...option.RequestOption) (res *ActiveV1InstrumentGetInstrumentsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/instruments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Represents a tradable financial instrument, including supplemental information
type Instrument struct {
	// The number of shares currently available to borrow
	AvailableToBorrow int64 `json:"available_to_borrow,nullable"`
	// The average daily trading volume over the past 30 days
	AverageVolume int64 `json:"average_volume,nullable"`
	// The beta value, measuring the instrument's volatility relative to the overall
	// market
	Beta string `json:"beta,nullable"`
	// The fee associated with borrowing the instrument, expressed as a decimal
	BorrowFee string `json:"borrow_fee,nullable"`
	// A detailed description of the instrument or company
	Description string `json:"description,nullable"`
	// The trailing twelve months (TTM) dividend yield
	DividendYield string `json:"dividend_yield,nullable"`
	// The trailing twelve months (TTM) earnings per share
	EarningsPerShare string `json:"earnings_per_share,nullable"`
	// The highest price over the last 52 weeks
	FiftyTwoWeekHigh string `json:"fifty_two_week_high,nullable"`
	// The lowest price over the last 52 weeks
	FiftyTwoWeekLow string `json:"fifty_two_week_low,nullable"`
	// The specific industry of the instrument's issuer
	Industry string `json:"industry,nullable"`
	// The date the instrument was first listed
	ListDate time.Time `json:"list_date,nullable" format:"date"`
	// URL to a representative logo image for the instrument or issuer
	LogoURL string `json:"logo_url,nullable"`
	// A cap on how much of your equity you can put into a single symbol on the long
	// side
	LongConcentrationLimit string `json:"long_concentration_limit,nullable"`
	// The total market capitalization
	MarketCap string `json:"market_cap,nullable"`
	// The closing price from the previous trading day
	PreviousClose string `json:"previous_close,nullable"`
	// The price-to-earnings (P/E) ratio for the trailing twelve months (TTM)
	PriceToEarnings string `json:"price_to_earnings,nullable"`
	// Real-time market quote data for the instrument
	Quote InstrumentQuote `json:"quote,nullable"`
	// The business sector of the instrument's issuer
	Sector string `json:"sector,nullable"`
	// A cap on how much of your equity you can allocate to a single symbol on the
	// short side
	ShortConcentrationLimit string `json:"short_concentration_limit,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableToBorrow       respjson.Field
		AverageVolume           respjson.Field
		Beta                    respjson.Field
		BorrowFee               respjson.Field
		Description             respjson.Field
		DividendYield           respjson.Field
		EarningsPerShare        respjson.Field
		FiftyTwoWeekHigh        respjson.Field
		FiftyTwoWeekLow         respjson.Field
		Industry                respjson.Field
		ListDate                respjson.Field
		LogoURL                 respjson.Field
		LongConcentrationLimit  respjson.Field
		MarketCap               respjson.Field
		PreviousClose           respjson.Field
		PriceToEarnings         respjson.Field
		Quote                   respjson.Field
		Sector                  respjson.Field
		ShortConcentrationLimit respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
	InstrumentCore
}

// Returns the unmodified JSON received from the API
func (r Instrument) RawJSON() string { return r.JSON.raw }
func (r *Instrument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstrumentCore struct {
	// Unique instrument identifier
	ID string `json:"id,required" format:"uuid"`
	// The ISO country code of the instrument's issue
	CountryOfIssue string `json:"country_of_issue,required"`
	// The ISO currency code in which the instrument is traded
	Currency string `json:"currency,required"`
	// Indicates if the instrument is classified as Easy-To-Borrow
	EasyToBorrow bool `json:"easy_to_borrow,required"`
	// Indicates if the instrument is liquidation only and cannot be bought
	IsLiquidationOnly bool `json:"is_liquidation_only,required"`
	// Indicates if the instrument is marginable
	IsMarginable bool `json:"is_marginable,required"`
	// Indicates if the instrument is restricted from trading
	IsRestricted bool `json:"is_restricted,required"`
	// Indicates if short selling is prohibited for the instrument
	IsShortProhibited bool `json:"is_short_prohibited,required"`
	// Indicates if the instrument is on the Regulation SHO Threshold Security List
	IsThresholdSecurity bool `json:"is_threshold_security,required"`
	// Deprecated. Use `security_ids`.
	//
	// A primary security identifier for this instrument.
	//
	// Deprecated: deprecated
	SecurityID string `json:"security_id,required"`
	// Deprecated. Use `security_ids`.
	//
	// The source for `security_id`.
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "CURRENCY", "FMP", "OEMS",
	// "SEDOL", "QUIK", "ISIN", "RIC", "COUNTRY", "EXCHANGE", "CTA", "BLOOMBERG",
	// "WERTPAPIER", "DUTCH", "VALOREN", "SICOVAM", "BELGIAN", "COMMON",
	// "CLEARING_HOUSE", "ISDA_FPML_SPECIFICATION", "ISDA_FPML_URL",
	// "LETTER_OF_CREDIT", "MARKETPLACE_ASSIGNED_IDENTIFIER", "MARKIT_RED_ENTITY_CLIP",
	// "MARKIT_RED_PAIR_CLIP", "CFTC", "ISDA_COMMODITY_REFERENCE_PRICE",
	// "LEGAL_ENTITY_IDENTIFIER", "SYNTHETIC", "FIDESSA_INSTRUMENT_MNEMONIC",
	// "INDEX_NAME", "UNIFORM_SYMBOL", "DIGITAL_TOKEN_IDENTIFIER", "MASSIVE", "OTHER".
	SecurityIDSource SecurityIDSource `json:"security_id_source,required"`
	// All known security identifiers for this instrument
	SecurityIDs []InstrumentSecurityID `json:"security_ids,required"`
	// The trading symbol for the instrument
	Symbol string `json:"symbol,required"`
	// The MIC code of the primary listing venue
	Venue string `json:"venue,required"`
	// The expiration date for options instruments
	Expiry time.Time `json:"expiry,nullable" format:"date"`
	// The percent of a long position's value you must post as margin
	LongMarginRate string `json:"long_margin_rate,nullable"`
	// The full name of the instrument or its issuer
	Name string `json:"name,nullable"`
	// The type of security (e.g., Common Stock, ETF)
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	SecurityType SecurityType `json:"security_type,nullable"`
	// The percent of a short position's value you must post as margin
	ShortMarginRate string `json:"short_margin_rate,nullable"`
	// The strike price for options instruments
	StrikePrice string `json:"strike_price,nullable"`
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
		SecurityID          respjson.Field
		SecurityIDSource    respjson.Field
		SecurityIDs         respjson.Field
		Symbol              respjson.Field
		Venue               respjson.Field
		Expiry              respjson.Field
		LongMarginRate      respjson.Field
		Name                respjson.Field
		SecurityType        respjson.Field
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
	Date time.Time `json:"date,required" format:"date"`
	// The actual earnings per share (EPS) for the period
	EpsActual string `json:"eps_actual,nullable"`
	// The estimated earnings per share (EPS) for the period
	EpsEstimate string `json:"eps_estimate,nullable"`
	// The percentage difference between actual and estimated EPS
	EpsSurprisePercent string `json:"eps_surprise_percent,nullable"`
	// The actual total revenue for the period
	RevenueActual string `json:"revenue_actual,nullable"`
	// The estimated total revenue for the period
	RevenueEstimate string `json:"revenue_estimate,nullable"`
	// The percentage difference between actual and estimated revenue
	RevenueSurprisePercent string `json:"revenue_surprise_percent,nullable"`
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

// Real-time market quote data for a specific instrument
type InstrumentQuote struct {
	// The highest trade price during the current trading day
	High string `json:"high,required"`
	// The most recent trade price
	LastPrice string `json:"last_price,required"`
	// The lowest trade price during the current trading day
	Low string `json:"low,required"`
	// The opening price for the current trading day
	Open string `json:"open,required"`
	// The total number of shares traded during the current trading day
	Volume int64 `json:"volume,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		High        respjson.Field
		LastPrice   respjson.Field
		Low         respjson.Field
		Open        respjson.Field
		Volume      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentQuote) RawJSON() string { return r.JSON.raw }
func (r *InstrumentQuote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a tradable financial instrument, as a more concise item listing only
// key fields.
type InstrumentSecurityID struct {
	// The identifier for the instrument
	SecurityID string `json:"security_id,required"`
	// The source system for the security identifier
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "CURRENCY", "FMP", "OEMS",
	// "SEDOL", "QUIK", "ISIN", "RIC", "COUNTRY", "EXCHANGE", "CTA", "BLOOMBERG",
	// "WERTPAPIER", "DUTCH", "VALOREN", "SICOVAM", "BELGIAN", "COMMON",
	// "CLEARING_HOUSE", "ISDA_FPML_SPECIFICATION", "ISDA_FPML_URL",
	// "LETTER_OF_CREDIT", "MARKETPLACE_ASSIGNED_IDENTIFIER", "MARKIT_RED_ENTITY_CLIP",
	// "MARKIT_RED_PAIR_CLIP", "CFTC", "ISDA_COMMODITY_REFERENCE_PRICE",
	// "LEGAL_ENTITY_IDENTIFIER", "SYNTHETIC", "FIDESSA_INSTRUMENT_MNEMONIC",
	// "INDEX_NAME", "UNIFORM_SYMBOL", "DIGITAL_TOKEN_IDENTIFIER", "MASSIVE", "OTHER".
	SecurityIDSource SecurityIDSource `json:"security_id_source,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SecurityID       respjson.Field
		SecurityIDSource respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentSecurityID) RawJSON() string { return r.JSON.raw }
func (r *InstrumentSecurityID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1InstrumentGetInstrumentByIDResponse struct {
	// Represents a tradable financial instrument, including supplemental information
	Data Instrument `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1InstrumentGetInstrumentByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1InstrumentGetInstrumentByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1InstrumentGetInstrumentsResponse struct {
	Data InstrumentCoreList `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1InstrumentGetInstrumentsResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1InstrumentGetInstrumentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1InstrumentGetInstrumentByIDParams struct {
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
	SecurityIDSource SecurityIDSource `path:"security_id_source,omitzero,required" json:"-"`
	paramObj
}

type ActiveV1InstrumentGetInstrumentsParams struct {
	// Filter by easy to borrow status
	EasyToBorrow param.Opt[bool] `query:"easy_to_borrow,omitzero" json:"-"`
	// Filter IDs to those containing this substring. For options, this is required.
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
	IsThresholdSecurity param.Opt[bool] `query:"is_threshold_security,omitzero" json:"-"`
	// The number of items to return per page (only used when page_token is not
	// provided)
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Filter by security ID(s). Accepts single value or indexed array.
	//
	// Examples:
	//
	// - Single: `security_id=037833100`
	// - Multiple: `security_id[0]=037833100&security_id[1]=594918104`
	SecurityID []string `query:"security_id,omitzero" json:"-"`
	// Source(s) for the security ID filter. Must match the count and order of
	// security_id.
	//
	// Examples:
	//
	// - Single: `security_id_source=CUSIP`
	// - Multiple: `security_id_source[0]=CUSIP&security_id_source[1]=FIGI`
	SecurityIDSource []string `query:"security_id_source,omitzero" json:"-"`
	// Filter by security type, required and defaults to COMMON_STOCK
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	SecurityType ActiveV1InstrumentGetInstrumentsParamsSecurityType `query:"security_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1InstrumentGetInstrumentsParams]'s query parameters
// as `url.Values`.
func (r ActiveV1InstrumentGetInstrumentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by security type, required and defaults to COMMON_STOCK
type ActiveV1InstrumentGetInstrumentsParamsSecurityType string

const (
	ActiveV1InstrumentGetInstrumentsParamsSecurityTypeCommonStock    ActiveV1InstrumentGetInstrumentsParamsSecurityType = "COMMON_STOCK"
	ActiveV1InstrumentGetInstrumentsParamsSecurityTypePreferredStock ActiveV1InstrumentGetInstrumentsParamsSecurityType = "PREFERRED_STOCK"
	ActiveV1InstrumentGetInstrumentsParamsSecurityTypeCorporateBond  ActiveV1InstrumentGetInstrumentsParamsSecurityType = "CORPORATE_BOND"
	ActiveV1InstrumentGetInstrumentsParamsSecurityTypeOption         ActiveV1InstrumentGetInstrumentsParamsSecurityType = "OPTION"
	ActiveV1InstrumentGetInstrumentsParamsSecurityTypeFuture         ActiveV1InstrumentGetInstrumentsParamsSecurityType = "FUTURE"
	ActiveV1InstrumentGetInstrumentsParamsSecurityTypeWarrant        ActiveV1InstrumentGetInstrumentsParamsSecurityType = "WARRANT"
	ActiveV1InstrumentGetInstrumentsParamsSecurityTypeCash           ActiveV1InstrumentGetInstrumentsParamsSecurityType = "CASH"
	ActiveV1InstrumentGetInstrumentsParamsSecurityTypeOther          ActiveV1InstrumentGetInstrumentsParamsSecurityType = "OTHER"
)
