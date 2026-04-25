// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"github.com/clear-street/clear-street-go/option"
)

// ActiveV1MarketDataService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1MarketDataService] method instead.
type ActiveV1MarketDataService struct {
	options []option.RequestOption
	// Real-time market data snapshots.
	DailySummary ActiveV1MarketDataDailySummaryService
	// Real-time market data snapshots.
	Snapshot ActiveV1MarketDataSnapshotService
}

// NewActiveV1MarketDataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1MarketDataService(opts ...option.RequestOption) (r ActiveV1MarketDataService) {
	r = ActiveV1MarketDataService{}
	r.options = opts
	r.DailySummary = NewActiveV1MarketDataDailySummaryService(opts...)
	r.Snapshot = NewActiveV1MarketDataSnapshotService(opts...)
	return
}
