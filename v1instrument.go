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

// Retrieve core details and discovery endpoints for tradable instruments.
//
// V1InstrumentService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1InstrumentService] method instead.
type V1InstrumentService struct {
	options []option.RequestOption
}

// NewV1InstrumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1InstrumentService(opts ...option.RequestOption) (r V1InstrumentService) {
	r = V1InstrumentService{}
	r.options = opts
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

// List options contracts.
//
// Returns options contracts for a given underlier with options-specific metadata.
// Exactly one underlier identifier must be provided.
func (r *V1InstrumentService) GetOptionContracts(ctx context.Context, query V1InstrumentGetOptionContractsParams, opts ...option.RequestOption) (res *V1InstrumentGetOptionContractsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/instruments/options/contracts"
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

// Represents a tradable financial instrument.
type Instrument struct {
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
	// Available options expiration dates for this instrument. Present only when
	// `include_options_expiry_dates=true` in the request.
	OptionsExpiryDates []time.Time `json:"options_expiry_dates" api:"nullable" format:"date"`
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
		OptionsExpiryDates  respjson.Field
		PreviousClose       respjson.Field
		ShortMarginRate     respjson.Field
		StrikePrice         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
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

type V1InstrumentGetOptionContractsResponse struct {
	Data OptionsContractList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentGetOptionContractsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentGetOptionContractsResponse) UnmarshalJSON(data []byte) error {
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
	// Filter by instrument type. OPTION is not supported on this endpoint; use GET
	// /instruments/options/contracts to list option contracts. If omitted, returns all
	// supported instrument types except options.
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

// Filter by instrument type. OPTION is not supported on this endpoint; use GET
// /instruments/options/contracts to list option contracts. If omitted, returns all
// supported instrument types except options.
type V1InstrumentGetInstrumentsParamsInstrumentType string

const (
	V1InstrumentGetInstrumentsParamsInstrumentTypeCommonStock    V1InstrumentGetInstrumentsParamsInstrumentType = "COMMON_STOCK"
	V1InstrumentGetInstrumentsParamsInstrumentTypePreferredStock V1InstrumentGetInstrumentsParamsInstrumentType = "PREFERRED_STOCK"
	V1InstrumentGetInstrumentsParamsInstrumentTypeOption         V1InstrumentGetInstrumentsParamsInstrumentType = "OPTION"
	V1InstrumentGetInstrumentsParamsInstrumentTypeCash           V1InstrumentGetInstrumentsParamsInstrumentType = "CASH"
	V1InstrumentGetInstrumentsParamsInstrumentTypeOther          V1InstrumentGetInstrumentsParamsInstrumentType = "OTHER"
)

type V1InstrumentGetOptionContractsParams struct {
	// Filter to contracts expiring on this date (YYYY-MM-DD)
	Expiry   param.Opt[time.Time] `query:"expiry,omitzero" format:"date" json:"-"`
	PageSize param.Opt[int64]     `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Underlier symbol (e.g., AAPL, SPX)
	Underlier param.Opt[string] `query:"underlier,omitzero" json:"-"`
	// OEMS instrument UUID or symbol of the underlying equity/index
	UnderlyingInstrumentID param.Opt[InstrumentIDOrSymbol] `query:"underlying_instrument_id,omitzero" format:"uuid" json:"-"`
	// Filter by contract type: CALL or PUT
	//
	// Any of "CALL", "PUT".
	ContractType ContractType `query:"contract_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentGetOptionContractsParams]'s query parameters as
// `url.Values`.
func (r V1InstrumentGetOptionContractsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

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
