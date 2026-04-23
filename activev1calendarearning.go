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
	"github.com/clear-street/clear-street-go/packages/param"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Access financial calendars for events like earnings, dividends, and splits.
//
// ActiveV1CalendarEarningService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1CalendarEarningService] method instead.
type ActiveV1CalendarEarningService struct {
	options []option.RequestOption
}

// NewActiveV1CalendarEarningService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1CalendarEarningService(opts ...option.RequestOption) (r ActiveV1CalendarEarningService) {
	r = ActiveV1CalendarEarningService{}
	r.options = opts
	return
}

// Retrieves upcoming earnings announcements.
func (r *ActiveV1CalendarEarningService) GetEarningsCalendar(ctx context.Context, query ActiveV1CalendarEarningGetEarningsCalendarParams, opts ...option.RequestOption) (res *ActiveV1CalendarEarningGetEarningsCalendarResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/calendars/earnings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Represents a single earnings announcement event
type EarningsCalendarEvent struct {
	// The date of the earnings announcement
	Date time.Time `json:"date" api:"required" format:"date"`
	// The date of the last update to this event
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date"`
	// The symbol for the instrument
	Symbol string `json:"symbol" api:"required"`
	// The actual reported earnings per share
	EpsActual string `json:"eps_actual" api:"nullable"`
	// The consensus estimated earnings per share
	EpsEstimated string `json:"eps_estimated" api:"nullable"`
	// The actual reported revenue
	RevenueActual string `json:"revenue_actual" api:"nullable"`
	// The consensus estimated revenue
	RevenueEstimated string `json:"revenue_estimated" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date             respjson.Field
		LastUpdated      respjson.Field
		Symbol           respjson.Field
		EpsActual        respjson.Field
		EpsEstimated     respjson.Field
		RevenueActual    respjson.Field
		RevenueEstimated respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EarningsCalendarEvent) RawJSON() string { return r.JSON.raw }
func (r *EarningsCalendarEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EarningsCalendarEventList []EarningsCalendarEvent

type ActiveV1CalendarEarningGetEarningsCalendarResponse struct {
	Data EarningsCalendarEventList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1CalendarEarningGetEarningsCalendarResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1CalendarEarningGetEarningsCalendarResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1CalendarEarningGetEarningsCalendarParams struct {
	// The start date for the query range, inclusive (YYYY-MM-DD)
	From param.Opt[time.Time] `query:"from,omitzero" format:"date" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD)
	To param.Opt[time.Time] `query:"to,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1CalendarEarningGetEarningsCalendarParams]'s query
// parameters as `url.Values`.
func (r ActiveV1CalendarEarningGetEarningsCalendarParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
