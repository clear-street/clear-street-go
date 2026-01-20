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

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/apiquery"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// ActiveV1InstrumentEventService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1InstrumentEventService] method instead.
type ActiveV1InstrumentEventService struct {
	Options []option.RequestOption
}

// NewActiveV1InstrumentEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1InstrumentEventService(opts ...option.RequestOption) (r ActiveV1InstrumentEventService) {
	r = ActiveV1InstrumentEventService{}
	r.Options = opts
	return
}

// Retrieves corporate events (dividends, splits, etc.) for an instrument.
func (r *ActiveV1InstrumentEventService) GetInstrumentEvents(ctx context.Context, securityID string, params ActiveV1InstrumentEventGetInstrumentEventsParams, opts ...option.RequestOption) (res *ActiveV1InstrumentEventGetInstrumentEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if securityID == "" {
		err = errors.New("missing required security_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/instruments/%v/%s/events", params.SecurityIDSource, securityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Represents an instrument event (dividends, splits, etc.)
type InstrumentEvent struct {
	// The date of the event
	Date time.Time `json:"date,required" format:"date"`
	// A brief description of the event
	Description string `json:"description,required"`
	// The type of event
	//
	// Any of "EARNINGS", "DIVIDEND", "SPLIT", "MERGER_ACQUISITION".
	EventType InstrumentEventEventType `json:"event_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Description respjson.Field
		EventType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentEvent) RawJSON() string { return r.JSON.raw }
func (r *InstrumentEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event
type InstrumentEventEventType string

const (
	InstrumentEventEventTypeEarnings          InstrumentEventEventType = "EARNINGS"
	InstrumentEventEventTypeDividend          InstrumentEventEventType = "DIVIDEND"
	InstrumentEventEventTypeSplit             InstrumentEventEventType = "SPLIT"
	InstrumentEventEventTypeMergerAcquisition InstrumentEventEventType = "MERGER_ACQUISITION"
)

type InstrumentEventList []InstrumentEvent

type ActiveV1InstrumentEventGetInstrumentEventsResponse struct {
	Data InstrumentEventList `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1InstrumentEventGetInstrumentEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1InstrumentEventGetInstrumentEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1InstrumentEventGetInstrumentEventsParams struct {
	// Security identifier source
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "OTHER".
	SecurityIDSource SecurityIDSource `path:"security_id_source,omitzero,required" json:"-"`
	// The start date for the query range, inclusive (YYYY-MM-DD)
	FromDate string `query:"from_date,required" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD)
	ToDate string `query:"to_date,required" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1InstrumentEventGetInstrumentEventsParams]'s query
// parameters as `url.Values`.
func (r ActiveV1InstrumentEventGetInstrumentEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
