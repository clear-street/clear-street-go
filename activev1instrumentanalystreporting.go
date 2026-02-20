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

// ActiveV1InstrumentAnalystReportingService contains methods and other services
// that help with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1InstrumentAnalystReportingService] method instead.
type ActiveV1InstrumentAnalystReportingService struct {
	Options []option.RequestOption
}

// NewActiveV1InstrumentAnalystReportingService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewActiveV1InstrumentAnalystReportingService(opts ...option.RequestOption) (r ActiveV1InstrumentAnalystReportingService) {
	r = ActiveV1InstrumentAnalystReportingService{}
	r.Options = opts
	return
}

// Retrieves analyst ratings and price targets for an instrument.
func (r *ActiveV1InstrumentAnalystReportingService) GetInstrumentAnalystConsensus(ctx context.Context, securityID string, params ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusParams, opts ...option.RequestOption) (res *ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if securityID == "" {
		err = errors.New("missing required security_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/instruments/%v/%s/analyst-reporting", params.SecurityIDSource, url.PathEscape(securityID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Analyst recommendation distribution
type AnalystDistribution struct {
	// Number of buy recommendations
	Buy int64 `json:"buy,required"`
	// Number of hold recommendations
	Hold int64 `json:"hold,required"`
	// Number of sell recommendations
	Sell int64 `json:"sell,required"`
	// Number of strong buy recommendations
	StrongBuy int64 `json:"strong_buy,required"`
	// Number of strong sell recommendations
	StrongSell int64 `json:"strong_sell,required"`
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

// Aggregated analyst consensus metrics
type InstrumentAnalystConsensus struct {
	// The date the consensus snapshot was generated
	Date time.Time `json:"date,required" format:"date"`
	// Count of individual analyst recommendations by category
	Distribution AnalystDistribution `json:"distribution,nullable"`
	// Aggregated analyst price target statistics
	PriceTarget PriceTarget `json:"price_target,nullable"`
	// Consensus analyst rating
	//
	// Any of "STRONG_BUY", "BUY", "HOLD", "SELL", "STRONG_SELL".
	Rating AnalystRating `json:"rating,nullable"`
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

// Analyst price target statistics
type PriceTarget struct {
	// Average analyst price target
	Average string `json:"average,required"`
	// ISO 4217 currency code of the price targets
	Currency string `json:"currency,required"`
	// Highest analyst price target
	High string `json:"high,required"`
	// Lowest analyst price target
	Low string `json:"low,required"`
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

type ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse struct {
	// Aggregated analyst consensus metrics
	Data InstrumentAnalystConsensus `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusParams struct {
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
	// The start date for the query range, inclusive (YYYY-MM-DD)
	From param.Opt[time.Time] `query:"from,omitzero" format:"date" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD)
	To param.Opt[time.Time] `query:"to,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes
// [ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusParams]'s query
// parameters as `url.Values`.
func (r ActiveV1InstrumentAnalystReportingGetInstrumentAnalystConsensusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
