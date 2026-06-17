// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"github.com/clear-street/clear-street-go/option"
)

// V1Service contains methods and other services that help with interacting with
// the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1Service] method instead.
type V1Service struct {
	options []option.RequestOption
	// Manage trading accounts, balances, and portfolio history.
	Accounts V1AccountService
	// Endpoints for API service metadata.
	APIVersion V1APIVersionService
	// Access clocks and financial calendars for market sessions and events.
	Calendar V1CalendarService
	// Retrieve instrument analytics, market data, news, and related reference data.
	InstrumentData V1InstrumentDataService
	// Retrieve core details and discovery endpoints for tradable instruments.
	Instruments V1InstrumentService
	OmniAI      V1OmniAIService
	// Place, monitor, and manage trading orders.
	Orders V1OrderService
	// View positions and manage position instructions.
	Positions V1PositionService
	// Search instruments and manage saved screeners.
	Screener V1ScreenerService
	// Create and manage watchlists.
	Watchlist V1WatchlistService
}

// NewV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1Service(opts ...option.RequestOption) (r V1Service) {
	r = V1Service{}
	r.options = opts
	r.Accounts = NewV1AccountService(opts...)
	r.APIVersion = NewV1APIVersionService(opts...)
	r.Calendar = NewV1CalendarService(opts...)
	r.InstrumentData = NewV1InstrumentDataService(opts...)
	r.Instruments = NewV1InstrumentService(opts...)
	r.OmniAI = NewV1OmniAIService(opts...)
	r.Orders = NewV1OrderService(opts...)
	r.Positions = NewV1PositionService(opts...)
	r.Screener = NewV1ScreenerService(opts...)
	r.Watchlist = NewV1WatchlistService(opts...)
	return
}

// Security type
type SecurityType string

const (
	SecurityTypeCommonStock SecurityType = "COMMON_STOCK"
	SecurityTypeOption      SecurityType = "OPTION"
	SecurityTypeCash        SecurityType = "CASH"
)

// Sort direction sorted results
type SortDirection string

const (
	SortDirectionAsc  SortDirection = "ASC"
	SortDirectionDesc SortDirection = "DESC"
)
