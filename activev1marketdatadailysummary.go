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

// Real-time market data snapshots.
//
// ActiveV1MarketDataDailySummaryService contains methods and other services that
// help with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1MarketDataDailySummaryService] method instead.
type ActiveV1MarketDataDailySummaryService struct {
	options []option.RequestOption
}

// NewActiveV1MarketDataDailySummaryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1MarketDataDailySummaryService(opts ...option.RequestOption) (r ActiveV1MarketDataDailySummaryService) {
	r = ActiveV1MarketDataDailySummaryService{}
	r.options = opts
	return
}

// Returns the most recent OHLV and current price for the requested OEMS
// instruments. Backed by the in-memory Polygon snapshot cache.
//
// Response contract: every request returns one row per **unique** `instrument_id`,
// in first-seen request order. Unresolvable IDs come back with `symbol = null` and
// every market-data field `null`; resolvable IDs with no cache entry come back
// with `symbol` populated but market-data fields `null`.
//
// **Note (temporary):** ID resolution currently goes through the supplemental
// screener (OEMS instrument_id → FMP fmp_symbol → metadata_id → realtime cache).
// Removed when the market-data service serves daily aggregates directly, or when
// Polygon symbology is loaded into the instrument cache.
func (r *ActiveV1MarketDataDailySummaryService) GetDailySummaries(ctx context.Context, query ActiveV1MarketDataDailySummaryGetDailySummariesParams, opts ...option.RequestOption) (res *ActiveV1MarketDataDailySummaryGetDailySummariesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/market-data/daily-summary"
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
	// OEMS instrument identifier. Always populated; echoes the request ID.
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

type ActiveV1MarketDataDailySummaryGetDailySummariesResponse struct {
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
func (r ActiveV1MarketDataDailySummaryGetDailySummariesResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1MarketDataDailySummaryGetDailySummariesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1MarketDataDailySummaryGetDailySummariesParams struct {
	// Comma-separated OEMS instrument UUIDs (required, 1..=100)
	InstrumentIDs string `query:"instrument_ids" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1MarketDataDailySummaryGetDailySummariesParams]'s
// query parameters as `url.Values`.
func (r ActiveV1MarketDataDailySummaryGetDailySummariesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
