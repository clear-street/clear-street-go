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

// ActiveV1CalendarMergersAcquisitionService contains methods and other services
// that help with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1CalendarMergersAcquisitionService] method instead.
type ActiveV1CalendarMergersAcquisitionService struct {
	Options []option.RequestOption
}

// NewActiveV1CalendarMergersAcquisitionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewActiveV1CalendarMergersAcquisitionService(opts ...option.RequestOption) (r ActiveV1CalendarMergersAcquisitionService) {
	r = ActiveV1CalendarMergersAcquisitionService{}
	r.Options = opts
	return
}

// Retrieves upcoming M&A events.
func (r *ActiveV1CalendarMergersAcquisitionService) GetMergersAndAcquisitionsCalendar(ctx context.Context, query ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarParams, opts ...option.RequestOption) (res *ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/calendars/mergers-acquisitions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Represents a merger or acquisition event
type MergersAcquisitionsEvent struct {
	// The symbol of the acquiring company
	AcquirerSymbol string `json:"acquirer_symbol,required"`
	// The symbol of the target company being acquired
	TargetSymbol string `json:"target_symbol,required"`
	// The date of the transaction
	TransactionDate time.Time `json:"transaction_date,required" format:"date"`
	// The timestamp when the merger or acquisition was accepted in UTC
	AcceptedAt time.Time `json:"accepted_at,nullable" format:"date-time"`
	// The CIK of the acquiring company
	AcquirerCik string `json:"acquirer_cik,nullable"`
	// The name of the acquiring company
	AcquirerName string `json:"acquirer_name,nullable"`
	// A URL link to more details about the merger or acquisition
	Link string `json:"link,nullable"`
	// The CIK of the target company
	TargetCik string `json:"target_cik,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcquirerSymbol  respjson.Field
		TargetSymbol    respjson.Field
		TransactionDate respjson.Field
		AcceptedAt      respjson.Field
		AcquirerCik     respjson.Field
		AcquirerName    respjson.Field
		Link            respjson.Field
		TargetCik       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergersAcquisitionsEvent) RawJSON() string { return r.JSON.raw }
func (r *MergersAcquisitionsEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MergersAcquisitionsEventList []MergersAcquisitionsEvent

type ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarResponse struct {
	Data MergersAcquisitionsEventList `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarParams struct {
	// The start date for the query range, inclusive (YYYY-MM-DD)
	FromDate string `query:"from_date,required" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD)
	ToDate string `query:"to_date,required" json:"-"`
	paramObj
}

// URLQuery serializes
// [ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarParams]'s
// query parameters as `url.Values`.
func (r ActiveV1CalendarMergersAcquisitionGetMergersAndAcquisitionsCalendarParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
