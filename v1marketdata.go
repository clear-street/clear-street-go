// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"github.com/clear-street/clear-street-go/option"
)

// V1MarketDataService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1MarketDataService] method instead.
type V1MarketDataService struct {
	options []option.RequestOption
	// Real-time market data snapshots.
	DailySummary V1MarketDataDailySummaryService
	// Real-time market data snapshots.
	Snapshot V1MarketDataSnapshotService
}

// NewV1MarketDataService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1MarketDataService(opts ...option.RequestOption) (r V1MarketDataService) {
	r = V1MarketDataService{}
	r.options = opts
	r.DailySummary = NewV1MarketDataDailySummaryService(opts...)
	r.Snapshot = NewV1MarketDataSnapshotService(opts...)
	return
}
