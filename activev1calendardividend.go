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
// ActiveV1CalendarDividendService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1CalendarDividendService] method instead.
type ActiveV1CalendarDividendService struct {
	Options []option.RequestOption
}

// NewActiveV1CalendarDividendService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1CalendarDividendService(opts ...option.RequestOption) (r ActiveV1CalendarDividendService) {
	r = ActiveV1CalendarDividendService{}
	r.Options = opts
	return
}

// Retrieves upcoming dividend payments.
func (r *ActiveV1CalendarDividendService) GetDividendsCalendar(ctx context.Context, query ActiveV1CalendarDividendGetDividendsCalendarParams, opts ...option.RequestOption) (res *ActiveV1CalendarDividendGetDividendsCalendarResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/calendars/dividends"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Represents a single dividend event
type DividendCalendarEvent struct {
	// The dividend amount adjusted for any stock splits
	AdjustedDividend string `json:"adjusted_dividend" api:"required"`
	// The ex-dividend date
	Date time.Time `json:"date" api:"required" format:"date"`
	// The dividend amount per share
	Dividend string `json:"dividend" api:"required"`
	// The symbol for the instrument
	Symbol string `json:"symbol" api:"required"`
	// The date the dividend was declared
	DeclarationDate time.Time `json:"declaration_date" api:"nullable" format:"date"`
	// The frequency of the dividend payment
	//
	// Any of "ANNUALLY", "SEMI_ANNUALLY", "QUARTERLY", "MONTHLY", "OTHER".
	Frequency DividendCalendarEventFrequency `json:"frequency" api:"nullable"`
	// The payment date for the dividend
	PaymentDate time.Time `json:"payment_date" api:"nullable" format:"date"`
	// The record date for the dividend
	RecordDate time.Time `json:"record_date" api:"nullable" format:"date"`
	// The dividend yield as a percentage decimal
	Yield string `json:"yield" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdjustedDividend respjson.Field
		Date             respjson.Field
		Dividend         respjson.Field
		Symbol           respjson.Field
		DeclarationDate  respjson.Field
		Frequency        respjson.Field
		PaymentDate      respjson.Field
		RecordDate       respjson.Field
		Yield            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DividendCalendarEvent) RawJSON() string { return r.JSON.raw }
func (r *DividendCalendarEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency of the dividend payment
type DividendCalendarEventFrequency string

const (
	DividendCalendarEventFrequencyAnnually     DividendCalendarEventFrequency = "ANNUALLY"
	DividendCalendarEventFrequencySemiAnnually DividendCalendarEventFrequency = "SEMI_ANNUALLY"
	DividendCalendarEventFrequencyQuarterly    DividendCalendarEventFrequency = "QUARTERLY"
	DividendCalendarEventFrequencyMonthly      DividendCalendarEventFrequency = "MONTHLY"
	DividendCalendarEventFrequencyOther        DividendCalendarEventFrequency = "OTHER"
)

type DividendCalendarEventList []DividendCalendarEvent

type ActiveV1CalendarDividendGetDividendsCalendarResponse struct {
	Data DividendCalendarEventList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1CalendarDividendGetDividendsCalendarResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1CalendarDividendGetDividendsCalendarResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1CalendarDividendGetDividendsCalendarParams struct {
	// The start date for the query range, inclusive (YYYY-MM-DD)
	From param.Opt[time.Time] `query:"from,omitzero" format:"date" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD)
	To param.Opt[time.Time] `query:"to,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1CalendarDividendGetDividendsCalendarParams]'s query
// parameters as `url.Values`.
func (r ActiveV1CalendarDividendGetDividendsCalendarParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
