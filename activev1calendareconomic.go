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

// ActiveV1CalendarEconomicService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1CalendarEconomicService] method instead.
type ActiveV1CalendarEconomicService struct {
	Options []option.RequestOption
}

// NewActiveV1CalendarEconomicService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1CalendarEconomicService(opts ...option.RequestOption) (r ActiveV1CalendarEconomicService) {
	r = ActiveV1CalendarEconomicService{}
	r.Options = opts
	return
}

// Retrieves upcoming economic events and indicators.
func (r *ActiveV1CalendarEconomicService) GetEconomicCalendar(ctx context.Context, query ActiveV1CalendarEconomicGetEconomicCalendarParams, opts ...option.RequestOption) (res *ActiveV1CalendarEconomicGetEconomicCalendarResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/calendars/economic"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Represents a single economic calendar event
type EconomicCalendarEvent struct {
	// The ISO 3166-1 alpha-2 country code
	Country string `json:"country,required"`
	// The ISO 4217 currency code
	Currency string `json:"currency,required"`
	// The name of the economic event
	EventName string `json:"event_name,required"`
	// The date and time of the event in UTC
	EventTimestamp time.Time `json:"event_timestamp,required" format:"date-time"`
	// The expected market impact of the event
	//
	// Any of "LOW", "MEDIUM", "HIGH".
	Impact EconomicCalendarEventImpact `json:"impact,required"`
	// The actual value reported for the event
	ActualValue string `json:"actual_value,nullable"`
	// The percentage change between the actual and previous values
	ChangePercent string `json:"change_percent,nullable"`
	// The market consensus estimate for the event's value
	EstimatedValue string `json:"estimated_value,nullable"`
	// The previous value for this event
	PreviousValue string `json:"previous_value,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country        respjson.Field
		Currency       respjson.Field
		EventName      respjson.Field
		EventTimestamp respjson.Field
		Impact         respjson.Field
		ActualValue    respjson.Field
		ChangePercent  respjson.Field
		EstimatedValue respjson.Field
		PreviousValue  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EconomicCalendarEvent) RawJSON() string { return r.JSON.raw }
func (r *EconomicCalendarEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The expected market impact of the event
type EconomicCalendarEventImpact string

const (
	EconomicCalendarEventImpactLow    EconomicCalendarEventImpact = "LOW"
	EconomicCalendarEventImpactMedium EconomicCalendarEventImpact = "MEDIUM"
	EconomicCalendarEventImpactHigh   EconomicCalendarEventImpact = "HIGH"
)

type EconomicCalendarEventList []EconomicCalendarEvent

type ActiveV1CalendarEconomicGetEconomicCalendarResponse struct {
	Data EconomicCalendarEventList `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1CalendarEconomicGetEconomicCalendarResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1CalendarEconomicGetEconomicCalendarResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1CalendarEconomicGetEconomicCalendarParams struct {
	// The start date for the query range, inclusive (YYYY-MM-DD)
	FromDate string `query:"from_date,required" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD)
	ToDate string `query:"to_date,required" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1CalendarEconomicGetEconomicCalendarParams]'s query
// parameters as `url.Values`.
func (r ActiveV1CalendarEconomicGetEconomicCalendarParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
