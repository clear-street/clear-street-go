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
// V1InstrumentAnalystReportingService contains methods and other services that
// help with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1InstrumentAnalystReportingService] method instead.
type V1InstrumentAnalystReportingService struct {
	options []option.RequestOption
}

// NewV1InstrumentAnalystReportingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1InstrumentAnalystReportingService(opts ...option.RequestOption) (r V1InstrumentAnalystReportingService) {
	r = V1InstrumentAnalystReportingService{}
	r.options = opts
	return
}

// Retrieves analyst ratings and price targets for an instrument.
func (r *V1InstrumentAnalystReportingService) GetInstrumentAnalystConsensus(ctx context.Context, instrumentID string, query V1InstrumentAnalystReportingGetInstrumentAnalystConsensusParams, opts ...option.RequestOption) (res *V1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/instruments/%s/analyst-reporting", instrumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

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

type V1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse struct {
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
func (r V1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *V1InstrumentAnalystReportingGetInstrumentAnalystConsensusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentAnalystReportingGetInstrumentAnalystConsensusParams struct {
	// The start date for the query range, inclusive (YYYY-MM-DD)
	From param.Opt[time.Time] `query:"from,omitzero" format:"date" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD)
	To param.Opt[time.Time] `query:"to,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes
// [V1InstrumentAnalystReportingGetInstrumentAnalystConsensusParams]'s query
// parameters as `url.Values`.
func (r V1InstrumentAnalystReportingGetInstrumentAnalystConsensusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
