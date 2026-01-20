// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"github.com/stainless-sdks/clear-street-go/option"
)

// ActiveV1CalendarService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1CalendarService] method instead.
type ActiveV1CalendarService struct {
	Options             []option.RequestOption
	Dividends           ActiveV1CalendarDividendService
	Earnings            ActiveV1CalendarEarningService
	Economic            ActiveV1CalendarEconomicService
	MarketHours         ActiveV1CalendarMarketHourService
	MergersAcquisitions ActiveV1CalendarMergersAcquisitionService
	Splits              ActiveV1CalendarSplitService
	Summary             ActiveV1CalendarSummaryService
}

// NewActiveV1CalendarService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1CalendarService(opts ...option.RequestOption) (r ActiveV1CalendarService) {
	r = ActiveV1CalendarService{}
	r.Options = opts
	r.Dividends = NewActiveV1CalendarDividendService(opts...)
	r.Earnings = NewActiveV1CalendarEarningService(opts...)
	r.Economic = NewActiveV1CalendarEconomicService(opts...)
	r.MarketHours = NewActiveV1CalendarMarketHourService(opts...)
	r.MergersAcquisitions = NewActiveV1CalendarMergersAcquisitionService(opts...)
	r.Splits = NewActiveV1CalendarSplitService(opts...)
	r.Summary = NewActiveV1CalendarSummaryService(opts...)
	return
}
