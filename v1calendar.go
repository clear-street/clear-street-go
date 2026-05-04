// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"github.com/clear-street/clear-street-go/option"
)

// V1CalendarService contains methods and other services that help with interacting
// with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CalendarService] method instead.
type V1CalendarService struct {
	options []option.RequestOption
	// Access financial calendars for events like earnings, dividends, and splits.
	MarketHours V1CalendarMarketHourService
}

// NewV1CalendarService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CalendarService(opts ...option.RequestOption) (r V1CalendarService) {
	r = V1CalendarService{}
	r.options = opts
	r.MarketHours = NewV1CalendarMarketHourService(opts...)
	return
}
