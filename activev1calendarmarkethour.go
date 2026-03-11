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
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// Access financial calendars for events like earnings, dividends, and splits.
//
// ActiveV1CalendarMarketHourService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1CalendarMarketHourService] method instead.
type ActiveV1CalendarMarketHourService struct {
	Options []option.RequestOption
}

// NewActiveV1CalendarMarketHourService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1CalendarMarketHourService(opts ...option.RequestOption) (r ActiveV1CalendarMarketHourService) {
	r = ActiveV1CalendarMarketHourService{}
	r.Options = opts
	return
}

// Retrieves comprehensive trading hours including pre-market, regular, and
// after-hours sessions. Returns market status, session times, and next session
// schedules.
func (r *ActiveV1CalendarMarketHourService) GetMarketHoursCalendar(ctx context.Context, query ActiveV1CalendarMarketHourGetMarketHoursCalendarParams, opts ...option.RequestOption) (res *ActiveV1CalendarMarketHourGetMarketHoursCalendarResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/calendars/market-hours"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Comprehensive market hours information for a specific market and date
type MarketHoursDetail struct {
	// Current time in market timezone with offset
	CurrentTime time.Time `json:"current_time" api:"required" format:"date-time"`
	// The date for which market hours are provided
	Date time.Time `json:"date" api:"required" format:"date"`
	// Market type identifier
	//
	// Any of "us_equities", "us_options".
	Market MarketHoursDetailMarket `json:"market" api:"required"`
	// Human-readable market name
	MarketName string `json:"market_name" api:"required"`
	// Next trading day's session schedules (without time_until fields)
	NextSessions MarketHoursDetailNextSessions `json:"next_sessions" api:"required"`
	// Market status information
	Status MarketHoursDetailStatus `json:"status" api:"required"`
	// IANA timezone identifier for the market
	Timezone string `json:"timezone" api:"required"`
	// Trading session schedules for the requested date with time_until fields
	TodaySessions MarketHoursDetailTodaySessions `json:"today_sessions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentTime   respjson.Field
		Date          respjson.Field
		Market        respjson.Field
		MarketName    respjson.Field
		NextSessions  respjson.Field
		Status        respjson.Field
		Timezone      respjson.Field
		TodaySessions respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketHoursDetail) RawJSON() string { return r.JSON.raw }
func (r *MarketHoursDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Market type identifier
type MarketHoursDetailMarket string

const (
	MarketHoursDetailMarketUsEquities MarketHoursDetailMarket = "us_equities"
	MarketHoursDetailMarketUsOptions  MarketHoursDetailMarket = "us_options"
)

// Next trading day's session schedules (without time_until fields)
type MarketHoursDetailNextSessions struct {
	// After-hours session schedule, null if not available
	AfterHours MarketHoursDetailNextSessionsAfterHours `json:"after_hours" api:"nullable"`
	// Pre-market session schedule, null if not available
	PreMarket MarketHoursDetailNextSessionsPreMarket `json:"pre_market" api:"nullable"`
	// Regular trading session schedule, null if holiday/weekend
	Regular MarketHoursDetailNextSessionsRegular `json:"regular" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AfterHours  respjson.Field
		PreMarket   respjson.Field
		Regular     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketHoursDetailNextSessions) RawJSON() string { return r.JSON.raw }
func (r *MarketHoursDetailNextSessions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// After-hours session schedule, null if not available
type MarketHoursDetailNextSessionsAfterHours struct {
	// Session close timestamp with timezone offset
	Close time.Time `json:"close" api:"required" format:"date-time"`
	// Session open timestamp with timezone offset
	Open time.Time `json:"open" api:"required" format:"date-time"`
	// ISO 8601 duration until session closes. Null if session is not currently open.
	TimeUntilClose string `json:"time_until_close" api:"nullable" format:"duration"`
	// ISO 8601 duration until session opens. Null if session has already started or
	// closed.
	TimeUntilOpen string `json:"time_until_open" api:"nullable" format:"duration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Close          respjson.Field
		Open           respjson.Field
		TimeUntilClose respjson.Field
		TimeUntilOpen  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketHoursDetailNextSessionsAfterHours) RawJSON() string { return r.JSON.raw }
func (r *MarketHoursDetailNextSessionsAfterHours) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pre-market session schedule, null if not available
type MarketHoursDetailNextSessionsPreMarket struct {
	// Session close timestamp with timezone offset
	Close time.Time `json:"close" api:"required" format:"date-time"`
	// Session open timestamp with timezone offset
	Open time.Time `json:"open" api:"required" format:"date-time"`
	// ISO 8601 duration until session closes. Null if session is not currently open.
	TimeUntilClose string `json:"time_until_close" api:"nullable" format:"duration"`
	// ISO 8601 duration until session opens. Null if session has already started or
	// closed.
	TimeUntilOpen string `json:"time_until_open" api:"nullable" format:"duration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Close          respjson.Field
		Open           respjson.Field
		TimeUntilClose respjson.Field
		TimeUntilOpen  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketHoursDetailNextSessionsPreMarket) RawJSON() string { return r.JSON.raw }
func (r *MarketHoursDetailNextSessionsPreMarket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Regular trading session schedule, null if holiday/weekend
type MarketHoursDetailNextSessionsRegular struct {
	// Session close timestamp with timezone offset
	Close time.Time `json:"close" api:"required" format:"date-time"`
	// Session open timestamp with timezone offset
	Open time.Time `json:"open" api:"required" format:"date-time"`
	// ISO 8601 duration until session closes. Null if session is not currently open.
	TimeUntilClose string `json:"time_until_close" api:"nullable" format:"duration"`
	// ISO 8601 duration until session opens. Null if session has already started or
	// closed.
	TimeUntilOpen string `json:"time_until_open" api:"nullable" format:"duration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Close          respjson.Field
		Open           respjson.Field
		TimeUntilClose respjson.Field
		TimeUntilOpen  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketHoursDetailNextSessionsRegular) RawJSON() string { return r.JSON.raw }
func (r *MarketHoursDetailNextSessionsRegular) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Market status information
type MarketHoursDetailStatus struct {
	// The type of trading day
	//
	// Any of "TRADING_DAY", "EARLY_CLOSE", "HOLIDAY", "WEEKEND".
	DayType string `json:"day_type" api:"required"`
	// Whether the market is currently open (real-time)
	IsOpen bool `json:"is_open" api:"required"`
	// Current session type if market is open, null if closed
	//
	// Any of "pre_market", "regular", "after_hours".
	CurrentSession string `json:"current_session" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DayType        respjson.Field
		IsOpen         respjson.Field
		CurrentSession respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketHoursDetailStatus) RawJSON() string { return r.JSON.raw }
func (r *MarketHoursDetailStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trading session schedules for the requested date with time_until fields
type MarketHoursDetailTodaySessions struct {
	// After-hours session schedule, null if not available
	AfterHours MarketHoursDetailTodaySessionsAfterHours `json:"after_hours" api:"nullable"`
	// Pre-market session schedule, null if not available
	PreMarket MarketHoursDetailTodaySessionsPreMarket `json:"pre_market" api:"nullable"`
	// Regular trading session schedule, null if holiday/weekend
	Regular MarketHoursDetailTodaySessionsRegular `json:"regular" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AfterHours  respjson.Field
		PreMarket   respjson.Field
		Regular     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketHoursDetailTodaySessions) RawJSON() string { return r.JSON.raw }
func (r *MarketHoursDetailTodaySessions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// After-hours session schedule, null if not available
type MarketHoursDetailTodaySessionsAfterHours struct {
	// Session close timestamp with timezone offset
	Close time.Time `json:"close" api:"required" format:"date-time"`
	// Session open timestamp with timezone offset
	Open time.Time `json:"open" api:"required" format:"date-time"`
	// ISO 8601 duration until session closes. Null if session is not currently open.
	TimeUntilClose string `json:"time_until_close" api:"nullable" format:"duration"`
	// ISO 8601 duration until session opens. Null if session has already started or
	// closed.
	TimeUntilOpen string `json:"time_until_open" api:"nullable" format:"duration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Close          respjson.Field
		Open           respjson.Field
		TimeUntilClose respjson.Field
		TimeUntilOpen  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketHoursDetailTodaySessionsAfterHours) RawJSON() string { return r.JSON.raw }
func (r *MarketHoursDetailTodaySessionsAfterHours) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pre-market session schedule, null if not available
type MarketHoursDetailTodaySessionsPreMarket struct {
	// Session close timestamp with timezone offset
	Close time.Time `json:"close" api:"required" format:"date-time"`
	// Session open timestamp with timezone offset
	Open time.Time `json:"open" api:"required" format:"date-time"`
	// ISO 8601 duration until session closes. Null if session is not currently open.
	TimeUntilClose string `json:"time_until_close" api:"nullable" format:"duration"`
	// ISO 8601 duration until session opens. Null if session has already started or
	// closed.
	TimeUntilOpen string `json:"time_until_open" api:"nullable" format:"duration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Close          respjson.Field
		Open           respjson.Field
		TimeUntilClose respjson.Field
		TimeUntilOpen  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketHoursDetailTodaySessionsPreMarket) RawJSON() string { return r.JSON.raw }
func (r *MarketHoursDetailTodaySessionsPreMarket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Regular trading session schedule, null if holiday/weekend
type MarketHoursDetailTodaySessionsRegular struct {
	// Session close timestamp with timezone offset
	Close time.Time `json:"close" api:"required" format:"date-time"`
	// Session open timestamp with timezone offset
	Open time.Time `json:"open" api:"required" format:"date-time"`
	// ISO 8601 duration until session closes. Null if session is not currently open.
	TimeUntilClose string `json:"time_until_close" api:"nullable" format:"duration"`
	// ISO 8601 duration until session opens. Null if session has already started or
	// closed.
	TimeUntilOpen string `json:"time_until_open" api:"nullable" format:"duration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Close          respjson.Field
		Open           respjson.Field
		TimeUntilClose respjson.Field
		TimeUntilOpen  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketHoursDetailTodaySessionsRegular) RawJSON() string { return r.JSON.raw }
func (r *MarketHoursDetailTodaySessionsRegular) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketHoursDetailList []MarketHoursDetail

type ActiveV1CalendarMarketHourGetMarketHoursCalendarResponse struct {
	Data MarketHoursDetailList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1CalendarMarketHourGetMarketHoursCalendarResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1CalendarMarketHourGetMarketHoursCalendarResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1CalendarMarketHourGetMarketHoursCalendarParams struct {
	// The date to query market hours for (YYYY-MM-DD). Defaults to today.
	Date string `query:"date" api:"required" json:"-"`
	// Market type to query (us_equities, us_options). If omitted, returns all markets.
	//
	// Any of "us_equities", "us_options".
	Market ActiveV1CalendarMarketHourGetMarketHoursCalendarParamsMarket `query:"market,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1CalendarMarketHourGetMarketHoursCalendarParams]'s
// query parameters as `url.Values`.
func (r ActiveV1CalendarMarketHourGetMarketHoursCalendarParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Market type to query (us_equities, us_options). If omitted, returns all markets.
type ActiveV1CalendarMarketHourGetMarketHoursCalendarParamsMarket string

const (
	ActiveV1CalendarMarketHourGetMarketHoursCalendarParamsMarketUsEquities ActiveV1CalendarMarketHourGetMarketHoursCalendarParamsMarket = "us_equities"
	ActiveV1CalendarMarketHourGetMarketHoursCalendarParamsMarketUsOptions  ActiveV1CalendarMarketHourGetMarketHoursCalendarParamsMarket = "us_options"
)
