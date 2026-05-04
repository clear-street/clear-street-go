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

// Access financial calendars for events like earnings, dividends, and splits.
//
// V1CalendarMarketHourService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CalendarMarketHourService] method instead.
type V1CalendarMarketHourService struct {
	options []option.RequestOption
}

// NewV1CalendarMarketHourService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CalendarMarketHourService(opts ...option.RequestOption) (r V1CalendarMarketHourService) {
	r = V1CalendarMarketHourService{}
	r.options = opts
	return
}

// Retrieves comprehensive trading hours including pre-market, regular, and
// after-hours sessions. Returns market status, session times, and next session
// schedules.
func (r *V1CalendarMarketHourService) GetMarketHoursCalendar(ctx context.Context, query V1CalendarMarketHourGetMarketHoursCalendarParams, opts ...option.RequestOption) (res *V1CalendarMarketHourGetMarketHoursCalendarResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/calendars/market-hours"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Day type for market hours - indicates the type of trading day
type DayType string

const (
	DayTypeTradingDay DayType = "TRADING_DAY"
	DayTypeEarlyClose DayType = "EARLY_CLOSE"
	DayTypeHoliday    DayType = "HOLIDAY"
	DayTypeWeekend    DayType = "WEEKEND"
)

// Comprehensive market hours information for a specific market and date
type MarketHoursDetail struct {
	// Current time in market timezone with offset
	CurrentTime time.Time `json:"current_time" api:"required" format:"date-time"`
	// The date for which market hours are provided
	Date time.Time `json:"date" api:"required" format:"date"`
	// Market type identifier
	//
	// Any of "us_equities", "us_options".
	Market MarketType `json:"market" api:"required"`
	// Human-readable market name
	MarketName string `json:"market_name" api:"required"`
	// Next trading day's session schedules (without time_until fields)
	NextSessions TradingSessions `json:"next_sessions" api:"required"`
	// Market status information
	Status MarketStatus `json:"status" api:"required"`
	// IANA timezone identifier for the market
	Timezone string `json:"timezone" api:"required"`
	// Trading session schedules for the requested date with time_until fields
	TodaySessions TradingSessions `json:"today_sessions" api:"required"`
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

type MarketHoursDetailList []MarketHoursDetail

// Session type for market hours
type MarketSessionType string

const (
	MarketSessionTypePreMarket  MarketSessionType = "pre_market"
	MarketSessionTypeRegular    MarketSessionType = "regular"
	MarketSessionTypeAfterHours MarketSessionType = "after_hours"
)

// Market status information
type MarketStatus struct {
	// The type of trading day
	//
	// Any of "TRADING_DAY", "EARLY_CLOSE", "HOLIDAY", "WEEKEND".
	DayType DayType `json:"day_type" api:"required"`
	// Whether the market is currently open (real-time)
	IsOpen bool `json:"is_open" api:"required"`
	// Current session type if market is open, null if closed
	//
	// Any of "pre_market", "regular", "after_hours".
	CurrentSession MarketSessionType `json:"current_session" api:"nullable"`
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
func (r MarketStatus) RawJSON() string { return r.JSON.raw }
func (r *MarketStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Market type for market hours calendar endpoint
type MarketType string

const (
	MarketTypeUsEquities MarketType = "us_equities"
	MarketTypeUsOptions  MarketType = "us_options"
)

// Session schedule with open and close timestamps
type SessionSchedule struct {
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
func (r SessionSchedule) RawJSON() string { return r.JSON.raw }
func (r *SessionSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trading sessions for a market day with full timestamps
type TradingSessions struct {
	// After-hours session schedule, null if not available
	AfterHours SessionSchedule `json:"after_hours" api:"nullable"`
	// Pre-market session schedule, null if not available
	PreMarket SessionSchedule `json:"pre_market" api:"nullable"`
	// Regular trading session schedule, null if holiday/weekend
	Regular SessionSchedule `json:"regular" api:"nullable"`
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
func (r TradingSessions) RawJSON() string { return r.JSON.raw }
func (r *TradingSessions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CalendarMarketHourGetMarketHoursCalendarResponse struct {
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
func (r V1CalendarMarketHourGetMarketHoursCalendarResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CalendarMarketHourGetMarketHoursCalendarResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CalendarMarketHourGetMarketHoursCalendarParams struct {
	// The date to query market hours for (YYYY-MM-DD). Defaults to today.
	Date string `query:"date" api:"required" json:"-"`
	// Market type to query (us_equities, us_options). If omitted, returns all markets.
	//
	// Any of "us_equities", "us_options".
	Market MarketType `query:"market,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CalendarMarketHourGetMarketHoursCalendarParams]'s query
// parameters as `url.Values`.
func (r V1CalendarMarketHourGetMarketHoursCalendarParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
