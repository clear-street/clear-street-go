// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
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

// Access financial calendars for events like earnings, dividends, and splits.
//
// ActiveV1CalendarSummaryService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1CalendarSummaryService] method instead.
type ActiveV1CalendarSummaryService struct {
	options []option.RequestOption
}

// NewActiveV1CalendarSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1CalendarSummaryService(opts ...option.RequestOption) (r ActiveV1CalendarSummaryService) {
	r = ActiveV1CalendarSummaryService{}
	r.options = opts
	return
}

// Retrieves a consolidated view of all calendar events.
func (r *ActiveV1CalendarSummaryService) GetCalendarSummary(ctx context.Context, query ActiveV1CalendarSummaryGetCalendarSummaryParams, opts ...option.RequestOption) (res *ActiveV1CalendarSummaryGetCalendarSummaryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/calendars/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Summary of events for a specific date
type CalendarDateSummary struct {
	// The date of the events
	Date time.Time `json:"date" api:"required" format:"date"`
	// The number of dividend events on this date
	DividendsCount int64 `json:"dividends_count" api:"required"`
	// The number of earnings announcements on this date
	EarningsCount int64 `json:"earnings_count" api:"required"`
	// The number of economic events on this date
	EconomicEventsCount int64 `json:"economic_events_count" api:"required"`
	// The number of mergers and acquisitions on this date
	MergersAcquisitionsCount int64 `json:"mergers_acquisitions_count" api:"required"`
	// The number of stock split events on this date
	StockSplitsCount int64 `json:"stock_splits_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date                     respjson.Field
		DividendsCount           respjson.Field
		EarningsCount            respjson.Field
		EconomicEventsCount      respjson.Field
		MergersAcquisitionsCount respjson.Field
		StockSplitsCount         respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalendarDateSummary) RawJSON() string { return r.JSON.raw }
func (r *CalendarDateSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CalendarDateSummaryList []CalendarDateSummary

type ActiveV1CalendarSummaryGetCalendarSummaryResponse struct {
	Data CalendarDateSummaryList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1CalendarSummaryGetCalendarSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1CalendarSummaryGetCalendarSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1CalendarSummaryGetCalendarSummaryParams struct {
	// The start date for the query range, inclusive (YYYY-MM-DD)
	From param.Opt[time.Time] `query:"from,omitzero" format:"date" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD)
	To param.Opt[time.Time] `query:"to,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1CalendarSummaryGetCalendarSummaryParams]'s query
// parameters as `url.Values`.
func (r ActiveV1CalendarSummaryGetCalendarSummaryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
