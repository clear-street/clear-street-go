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
// V1InstrumentEventService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1InstrumentEventService] method instead.
type V1InstrumentEventService struct {
	options []option.RequestOption
}

// NewV1InstrumentEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1InstrumentEventService(opts ...option.RequestOption) (r V1InstrumentEventService) {
	r = V1InstrumentEventService{}
	r.options = opts
	return
}

// List instrument events across all securities.
//
// Retrieves all instrument events grouped by date.
func (r *V1InstrumentEventService) GetAllInstrumentEvents(ctx context.Context, query V1InstrumentEventGetAllInstrumentEventsParams, opts ...option.RequestOption) (res *V1InstrumentEventGetAllInstrumentEventsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/instruments/events"
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
func (r *V1InstrumentEventService) GetInstrumentEvents(ctx context.Context, instrumentID InstrumentIDOrSymbol, query V1InstrumentEventGetInstrumentEventsParams, opts ...option.RequestOption) (res *V1InstrumentEventGetInstrumentEventsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/instruments/%s/events", instrumentID)
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

type V1InstrumentEventGetAllInstrumentEventsResponse struct {
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
func (r V1InstrumentEventGetAllInstrumentEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentEventGetAllInstrumentEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentEventGetInstrumentEventsResponse struct {
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
func (r V1InstrumentEventGetInstrumentEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentEventGetInstrumentEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentEventGetAllInstrumentEventsParams struct {
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

// URLQuery serializes [V1InstrumentEventGetAllInstrumentEventsParams]'s query
// parameters as `url.Values`.
func (r V1InstrumentEventGetAllInstrumentEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1InstrumentEventGetInstrumentEventsParams struct {
	// The start date for the query range, inclusive (YYYY-MM-DD).
	FromDate param.Opt[string] `query:"from_date,omitzero" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD).
	ToDate param.Opt[string] `query:"to_date,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentEventGetInstrumentEventsParams]'s query
// parameters as `url.Values`.
func (r V1InstrumentEventGetInstrumentEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
