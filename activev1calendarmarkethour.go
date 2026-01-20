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

// Retrieves trading hours and market holidays.
func (r *ActiveV1CalendarMarketHourService) GetMarketHoursCalendar(ctx context.Context, query ActiveV1CalendarMarketHourGetMarketHoursCalendarParams, opts ...option.RequestOption) (res *ActiveV1CalendarMarketHourGetMarketHoursCalendarResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/calendars/market-hours"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Trading hours and market status for a specific venue and date
type MarketHours struct {
	// The date for which market hours are provided
	Date time.Time `json:"date,required" format:"date"`
	// Whether the market is open for trading on this date
	IsOpen bool `json:"is_open,required"`
	// IANA timezone identifier for the venue
	Timezone string `json:"timezone,required"`
	// The MIC code of the venue
	Venue string `json:"venue,required"`
	// Market close time in local venue timezone (HH:MM:SS). Null if market is closed
	CloseTime string `json:"close_time,nullable"`
	// Name of the holiday if market is closed for a holiday. Null otherwise
	HolidayName string `json:"holiday_name,nullable"`
	// Next market close timestamp in UTC
	NextClose time.Time `json:"next_close,nullable" format:"date-time"`
	// Next market open timestamp in UTC
	NextOpen time.Time `json:"next_open,nullable" format:"date-time"`
	// Market open time in local venue timezone (HH:MM:SS). Null if market is closed
	OpenTime string `json:"open_time,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		IsOpen      respjson.Field
		Timezone    respjson.Field
		Venue       respjson.Field
		CloseTime   respjson.Field
		HolidayName respjson.Field
		NextClose   respjson.Field
		NextOpen    respjson.Field
		OpenTime    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketHours) RawJSON() string { return r.JSON.raw }
func (r *MarketHours) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MarketHoursList []MarketHours

type ActiveV1CalendarMarketHourGetMarketHoursCalendarResponse struct {
	Data MarketHoursList `json:"data,required"`
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
	// The date to query market hours for (YYYY-MM-DD)
	Date string `query:"date,required" json:"-"`
	// The MIC code of the venue
	Venue param.Opt[string] `query:"venue,omitzero" json:"-"`
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
