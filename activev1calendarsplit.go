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

// ActiveV1CalendarSplitService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1CalendarSplitService] method instead.
type ActiveV1CalendarSplitService struct {
	Options []option.RequestOption
}

// NewActiveV1CalendarSplitService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1CalendarSplitService(opts ...option.RequestOption) (r ActiveV1CalendarSplitService) {
	r = ActiveV1CalendarSplitService{}
	r.Options = opts
	return
}

// Retrieves upcoming stock splits.
func (r *ActiveV1CalendarSplitService) GetSplitsCalendar(ctx context.Context, query ActiveV1CalendarSplitGetSplitsCalendarParams, opts ...option.RequestOption) (res *ActiveV1CalendarSplitGetSplitsCalendarResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/calendars/splits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Represents a stock split event
type StockSplitEvent struct {
	// The date the split will occur
	Date time.Time `json:"date" api:"required" format:"date"`
	// The pre-split number of shares
	Denominator int64 `json:"denominator" api:"required"`
	// The post-split number of shares for every 'denominator' pre-split shares
	Numerator int64 `json:"numerator" api:"required"`
	// The symbol for the instrument
	Symbol string `json:"symbol" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Denominator respjson.Field
		Numerator   respjson.Field
		Symbol      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StockSplitEvent) RawJSON() string { return r.JSON.raw }
func (r *StockSplitEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StockSplitEventList []StockSplitEvent

type ActiveV1CalendarSplitGetSplitsCalendarResponse struct {
	Data StockSplitEventList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1CalendarSplitGetSplitsCalendarResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1CalendarSplitGetSplitsCalendarResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1CalendarSplitGetSplitsCalendarParams struct {
	// The start date for the query range, inclusive (YYYY-MM-DD)
	From param.Opt[time.Time] `query:"from,omitzero" format:"date" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD)
	To param.Opt[time.Time] `query:"to,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1CalendarSplitGetSplitsCalendarParams]'s query
// parameters as `url.Values`.
func (r ActiveV1CalendarSplitGetSplitsCalendarParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
