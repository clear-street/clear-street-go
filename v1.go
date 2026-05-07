// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"net/http"
	"slices"

	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
)

// Active Websocket.
//
// V1Service contains methods and other services that help with interacting with
// the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1Service] method instead.
type V1Service struct {
	options []option.RequestOption
	// Manage trading accounts, balances, and portfolio history.
	Accounts  V1AccountService
	Calendars V1CalendarService
	// Access financial calendars for events like earnings, dividends, and splits.
	Clock V1ClockService
	// Retrieve details and lists of tradable instruments.
	Instruments V1InstrumentService
	MarketData  V1MarketDataService
	// Retrieve market news and related instrument metadata.
	News   V1NewsService
	OmniAI V1OmniAIService
	// Endpoints for API service metadata.
	Version V1VersionService
	// Create and manage watchlists.
	Watchlists V1WatchlistService
}

// NewV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1Service(opts ...option.RequestOption) (r V1Service) {
	r = V1Service{}
	r.options = opts
	r.Accounts = NewV1AccountService(opts...)
	r.Calendars = NewV1CalendarService(opts...)
	r.Clock = NewV1ClockService(opts...)
	r.Instruments = NewV1InstrumentService(opts...)
	r.MarketData = NewV1MarketDataService(opts...)
	r.News = NewV1NewsService(opts...)
	r.OmniAI = NewV1OmniAIService(opts...)
	r.Version = NewV1VersionService(opts...)
	r.Watchlists = NewV1WatchlistService(opts...)
	return
}

// Upgrade the HTTP connection to a WebSocket and echo incoming messages.
func (r *V1Service) WebsocketHandler(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/ws"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Security type
type SecurityType string

const (
	SecurityTypeCommonStock    SecurityType = "COMMON_STOCK"
	SecurityTypePreferredStock SecurityType = "PREFERRED_STOCK"
	SecurityTypeOption         SecurityType = "OPTION"
	SecurityTypeCash           SecurityType = "CASH"
	SecurityTypeOther          SecurityType = "OTHER"
)
