// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// Retrieve details and lists of tradable instruments.
//
// ActiveV1InstrumentVenueService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1InstrumentVenueService] method instead.
type ActiveV1InstrumentVenueService struct {
	Options []option.RequestOption
}

// NewActiveV1InstrumentVenueService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1InstrumentVenueService(opts ...option.RequestOption) (r ActiveV1InstrumentVenueService) {
	r = ActiveV1InstrumentVenueService{}
	r.Options = opts
	return
}

// Retrieves a list of available trading venues and exchanges.
func (r *ActiveV1InstrumentVenueService) GetVenues(ctx context.Context, opts ...option.RequestOption) (res *ActiveV1InstrumentVenueGetVenuesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/instruments/venues"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A trading venue with its characteristics and capabilities
type Venue struct {
	// The ISO country code where the venue operates
	Country string `json:"country" api:"required"`
	// The display characteristics of the venue
	//
	// Any of "LIT", "DARK", "PERIODIC_AUCTION", "RFQ".
	DisplayType VenueDisplayType `json:"display_type" api:"required"`
	// Indicates whether GOOD_TILL_DATE orders accept date-only or timestamp
	// specifications
	GtdAccepts VenueGtdAccepts `json:"gtd_accepts" api:"required"`
	// The minimum quantity increment for orders at this venue
	LotSize int64 `json:"lot_size" api:"required"`
	// The Market Identifier Code (MIC) for the venue
	Mic string `json:"mic" api:"required"`
	// The display name of the venue
	Name string `json:"name" api:"required"`
	// Trading sessions available at this venue
	Sessions []VenueSession `json:"sessions" api:"required"`
	// Order types supported by this venue
	SupportedOrderTypes []string `json:"supported_order_types" api:"required"`
	// Time-in-force options supported by this venue
	SupportedTifs []string `json:"supported_tifs" api:"required"`
	// The minimum price increment for orders at this venue
	TickSize string `json:"tick_size" api:"required"`
	// IANA timezone identifier for the venue's local time
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country             respjson.Field
		DisplayType         respjson.Field
		GtdAccepts          respjson.Field
		LotSize             respjson.Field
		Mic                 respjson.Field
		Name                respjson.Field
		Sessions            respjson.Field
		SupportedOrderTypes respjson.Field
		SupportedTifs       respjson.Field
		TickSize            respjson.Field
		Timezone            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Venue) RawJSON() string { return r.JSON.raw }
func (r *Venue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The display characteristics of the venue
type VenueDisplayType string

const (
	VenueDisplayTypeLit             VenueDisplayType = "LIT"
	VenueDisplayTypeDark            VenueDisplayType = "DARK"
	VenueDisplayTypePeriodicAuction VenueDisplayType = "PERIODIC_AUCTION"
	VenueDisplayTypeRfq             VenueDisplayType = "RFQ"
)

// Indicates whether GOOD_TILL_DATE orders accept date-only or timestamp
// specifications
type VenueGtdAccepts struct {
	// Whether the venue accepts date-only expiration (YYYY-MM-DD)
	Date bool `json:"date" api:"required"`
	// Whether the venue accepts precise timestamp expiration
	Timestamp bool `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VenueGtdAccepts) RawJSON() string { return r.JSON.raw }
func (r *VenueGtdAccepts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A trading session within a venue's trading day
type VenueSession struct {
	// Session end time in venue's local timezone (HH:MM format, 24-hour)
	EndLocal string `json:"end_local" api:"required"`
	// The name of the trading session
	Name string `json:"name" api:"required"`
	// Session start time in venue's local timezone (HH:MM format, 24-hour)
	StartLocal string `json:"start_local" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndLocal    respjson.Field
		Name        respjson.Field
		StartLocal  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VenueSession) RawJSON() string { return r.JSON.raw }
func (r *VenueSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VenueList []Venue

type ActiveV1InstrumentVenueGetVenuesResponse struct {
	Data VenueList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1InstrumentVenueGetVenuesResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1InstrumentVenueGetVenuesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
