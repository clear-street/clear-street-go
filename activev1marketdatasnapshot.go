// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/internal/apiquery"
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/param"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Real-time market data snapshots.
//
// ActiveV1MarketDataSnapshotService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1MarketDataSnapshotService] method instead.
type ActiveV1MarketDataSnapshotService struct {
	options []option.RequestOption
}

// NewActiveV1MarketDataSnapshotService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1MarketDataSnapshotService(opts ...option.RequestOption) (r ActiveV1MarketDataSnapshotService) {
	r = ActiveV1MarketDataSnapshotService{}
	r.options = opts
	return
}

// Get market data snapshots for one or more securities.
func (r *ActiveV1MarketDataSnapshotService) GetSnapshots(ctx context.Context, query ActiveV1MarketDataSnapshotGetSnapshotsParams, opts ...option.RequestOption) (res *ActiveV1MarketDataSnapshotGetSnapshotsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/market-data/snapshot"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Market data snapshot for a single security.
type MarketDataSnapshot struct {
	// OEMS instrument identifier.
	InstrumentID string `json:"instrument_id" api:"required"`
	// Display symbol for the security.
	Symbol string `json:"symbol" api:"required"`
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
		InstrumentID respjson.Field
		Symbol       respjson.Field
		LastQuote    respjson.Field
		LastTrade    respjson.Field
		Name         respjson.Field
		Session      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketDataSnapshot) RawJSON() string { return r.JSON.raw }
func (r *MarketDataSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketDataSnapshotList []MarketDataSnapshot

// Last-trade fields for a market data snapshot.
type SnapshotLastTrade struct {
	// Most recent last-sale eligible trade price.
	Price string `json:"price" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Price       respjson.Field
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

type ActiveV1MarketDataSnapshotGetSnapshotsResponse struct {
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
func (r ActiveV1MarketDataSnapshotGetSnapshotsResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1MarketDataSnapshotGetSnapshotsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1MarketDataSnapshotGetSnapshotsParams struct {
	// Comma-separated OEMS instrument UUIDs.
	IDs param.Opt[string] `query:"ids,omitzero" json:"-"`
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
	paramObj
}

// URLQuery serializes [ActiveV1MarketDataSnapshotGetSnapshotsParams]'s query
// parameters as `url.Values`.
func (r ActiveV1MarketDataSnapshotGetSnapshotsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
