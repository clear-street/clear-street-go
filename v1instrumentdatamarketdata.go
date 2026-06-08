// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/internal/apiquery"
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Retrieve instrument analytics, market data, news, and related reference data.
//
// V1InstrumentDataMarketDataService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1InstrumentDataMarketDataService] method instead.
type V1InstrumentDataMarketDataService struct {
	options []option.RequestOption
}

// NewV1InstrumentDataMarketDataService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1InstrumentDataMarketDataService(opts ...option.RequestOption) (r V1InstrumentDataMarketDataService) {
	r = V1InstrumentDataMarketDataService{}
	r.options = opts
	return
}

// Returns the most recent open, high, low, volume (OHLV) and current price for the
// requested instruments.
//
// Response contract: every request returns one row per **unique** `instrument_id`,
// in first-seen request order. Unresolvable IDs come back with `symbol = null` and
// every market-data field `null`; resolvable IDs with no available data come back
// with `symbol` populated but market-data fields `null`.
func (r *V1InstrumentDataMarketDataService) GetDailySummaries(ctx context.Context, query V1InstrumentDataMarketDataGetDailySummariesParams, opts ...option.RequestOption) (res *V1InstrumentDataMarketDataGetDailySummariesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/market-data/daily-summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get market data snapshots for one or more securities.
func (r *V1InstrumentDataMarketDataService) GetSnapshots(ctx context.Context, query V1InstrumentDataMarketDataGetSnapshotsParams, opts ...option.RequestOption) (res *V1InstrumentDataMarketDataGetSnapshotsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/market-data/snapshot"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Daily aggregate (OHLV) summary for a single instrument.
//
// Returned by `GET /market-data/daily-summary`. Every field except `instrument_id`
// is `Option`:
//
//   - Unresolvable `instrument_id` → all other fields `None` (including `symbol`).
//   - Resolvable `instrument_id` with no realtime cache entry → `symbol` populated,
//     OHLV/`trade_date` `None`.
//   - `trade_date` reflects the session the OHLV represents (today during trading
//     hours, the last trading date during weekends/holidays).
type DailySummary struct {
	// Unique instrument identifier. Always populated; echoes the request ID.
	InstrumentID string `json:"instrument_id" api:"required" format:"uuid"`
	// Session high.
	High string `json:"high" api:"nullable"`
	// Session low.
	Low string `json:"low" api:"nullable"`
	// Opening price for the session.
	Open string `json:"open" api:"nullable"`
	// Display symbol for the security. `None` for unresolvable IDs.
	Symbol string `json:"symbol" api:"nullable"`
	// Session date the OHLV represents, US/Eastern.
	TradeDate time.Time `json:"trade_date" api:"nullable" format:"date"`
	// Session cumulative trading volume.
	Volume int64 `json:"volume" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InstrumentID respjson.Field
		High         respjson.Field
		Low          respjson.Field
		Open         respjson.Field
		Symbol       respjson.Field
		TradeDate    respjson.Field
		Volume       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DailySummary) RawJSON() string { return r.JSON.raw }
func (r *DailySummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DailySummaryList []DailySummary

// Market data snapshot for a single security.
type MarketDataSnapshot struct {
	// Unique instrument identifier.
	InstrumentID string `json:"instrument_id" api:"required"`
	// Display symbol for the security.
	Symbol string `json:"symbol" api:"required"`
	// Cumulative traded volume reported on the most recent trade, in shares for
	// equities or contracts for options. Absent when no trade is available.
	CumulativeVolume int64 `json:"cumulative_volume" api:"nullable"`
	// Theoretical price and Greeks for option instruments. `None` for equities, and
	// for options whose Greeks have not yet been observed
	Greeks SnapshotGreeks `json:"greeks" api:"nullable"`
	// Most recent quote if available.
	LastQuote SnapshotQuote `json:"last_quote" api:"nullable"`
	// Most recent last-sale trade if available.
	LastTrade SnapshotLastTrade `json:"last_trade" api:"nullable"`
	// Security name if available.
	Name string `json:"name" api:"nullable"`
	// Session metrics computed from previous close and last trade, if available.
	Session SnapshotSession `json:"session" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InstrumentID     respjson.Field
		Symbol           respjson.Field
		CumulativeVolume respjson.Field
		Greeks           respjson.Field
		LastQuote        respjson.Field
		LastTrade        respjson.Field
		Name             respjson.Field
		Session          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketDataSnapshot) RawJSON() string { return r.JSON.raw }
func (r *MarketDataSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketDataSnapshotList []MarketDataSnapshot

// Theoretical price and Greeks for an options snapshot. All values are **per
// share**; no contract multiplier is applied.
type SnapshotGreeks struct {
	// Delta: ∂V/∂S, range \[-1, 1\].
	Delta string `json:"delta" api:"required"`
	// Gamma: ∂²V/∂S².
	Gamma string `json:"gamma" api:"required"`
	// Implied volatility, annualized (`0.20` == 20%).
	Iv string `json:"iv" api:"required"`
	// Rho per 1.0 rate point.
	Rho string `json:"rho" api:"required"`
	// Theoretical option price in USD per share.
	TheoPrice string `json:"theo_price" api:"required"`
	// Theta per trading day.
	Theta string `json:"theta" api:"required"`
	// Timestamp when the Greeks were calculated.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Vega per 1.0 vol point.
	Vega string `json:"vega" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delta       respjson.Field
		Gamma       respjson.Field
		Iv          respjson.Field
		Rho         respjson.Field
		TheoPrice   respjson.Field
		Theta       respjson.Field
		Timestamp   respjson.Field
		Vega        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapshotGreeks) RawJSON() string { return r.JSON.raw }
func (r *SnapshotGreeks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Last-trade fields for a market data snapshot.
type SnapshotLastTrade struct {
	// Most recent last-sale eligible trade price.
	Price string `json:"price" api:"required"`
	// Share quantity of the most recent last-sale eligible trade.
	Size int64 `json:"size" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Price       respjson.Field
		Size        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapshotLastTrade) RawJSON() string { return r.JSON.raw }
func (r *SnapshotLastTrade) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// L1 quote fields for a market data snapshot.
type SnapshotQuote struct {
	// Current best ask.
	Ask string `json:"ask" api:"required"`
	// Current best bid.
	Bid string `json:"bid" api:"required"`
	// Midpoint of bid and ask.
	Midpoint string `json:"midpoint" api:"required"`
	// Size at the best ask, in shares.
	AskSize int64 `json:"ask_size" api:"nullable"`
	// Size at the best bid, in shares.
	BidSize int64 `json:"bid_size" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ask         respjson.Field
		Bid         respjson.Field
		Midpoint    respjson.Field
		AskSize     respjson.Field
		BidSize     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapshotQuote) RawJSON() string { return r.JSON.raw }
func (r *SnapshotQuote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Session-level pricing metrics for a market data snapshot.
type SnapshotSession struct {
	// Absolute change from previous close to last trade.
	Change string `json:"change" api:"required"`
	// Percent change from previous close to last trade.
	ChangePercent string `json:"change_percent" api:"required"`
	// Previous session close price.
	PreviousClose string `json:"previous_close" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Change        respjson.Field
		ChangePercent respjson.Field
		PreviousClose respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SnapshotSession) RawJSON() string { return r.JSON.raw }
func (r *SnapshotSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentDataMarketDataGetDailySummariesResponse struct {
	Data DailySummaryList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentDataMarketDataGetDailySummariesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentDataMarketDataGetDailySummariesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentDataMarketDataGetSnapshotsResponse struct {
	Data MarketDataSnapshotList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentDataMarketDataGetSnapshotsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentDataMarketDataGetSnapshotsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentDataMarketDataGetDailySummariesParams struct {
	// Comma-separated instrument identifiers (required, 1..=100)
	InstrumentIDs string `query:"instrument_ids" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentDataMarketDataGetDailySummariesParams]'s query
// parameters as `url.Values`.
func (r V1InstrumentDataMarketDataGetDailySummariesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1InstrumentDataMarketDataGetSnapshotsParams struct {
	// Comma-separated instrument identifiers.
	InstrumentIDs []string `query:"instrument_ids,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentDataMarketDataGetSnapshotsParams]'s query
// parameters as `url.Values`.
func (r V1InstrumentDataMarketDataGetSnapshotsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
