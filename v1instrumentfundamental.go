// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Retrieve details and lists of tradable instruments.
//
// V1InstrumentFundamentalService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1InstrumentFundamentalService] method instead.
type V1InstrumentFundamentalService struct {
	options []option.RequestOption
}

// NewV1InstrumentFundamentalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1InstrumentFundamentalService(opts ...option.RequestOption) (r V1InstrumentFundamentalService) {
	r = V1InstrumentFundamentalService{}
	r.options = opts
	return
}

// Retrieves supplemental fundamentals and company profile data for an instrument.
func (r *V1InstrumentFundamentalService) GetInstrumentFundamentals(ctx context.Context, instrumentID InstrumentIDOrSymbol, opts ...option.RequestOption) (res *V1InstrumentFundamentalGetInstrumentFundamentalsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/instruments/%s/fundamentals", instrumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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

type V1InstrumentFundamentalGetInstrumentFundamentalsResponse struct {
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
func (r V1InstrumentFundamentalGetInstrumentFundamentalsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentFundamentalGetInstrumentFundamentalsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
