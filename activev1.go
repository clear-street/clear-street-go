// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"github.com/stainless-sdks/clear-street-go/option"
)

// ActiveV1Service contains methods and other services that help with interacting
// with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1Service] method instead.
type ActiveV1Service struct {
	Options     []option.RequestOption
	Accounts    ActiveV1AccountService
	Assistant   ActiveV1AssistantService
	Calendars   ActiveV1CalendarService
	Instruments ActiveV1InstrumentService
	Screener    ActiveV1ScreenerService
	Version     ActiveV1VersionService
	Ws          ActiveV1WService
}

// NewActiveV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActiveV1Service(opts ...option.RequestOption) (r ActiveV1Service) {
	r = ActiveV1Service{}
	r.Options = opts
	r.Accounts = NewActiveV1AccountService(opts...)
	r.Assistant = NewActiveV1AssistantService(opts...)
	r.Calendars = NewActiveV1CalendarService(opts...)
	r.Instruments = NewActiveV1InstrumentService(opts...)
	r.Screener = NewActiveV1ScreenerService(opts...)
	r.Version = NewActiveV1VersionService(opts...)
	r.Ws = NewActiveV1WService(opts...)
	return
}

// Security identifier source
type SecurityIDSource string

const (
	SecurityIDSourceCms   SecurityIDSource = "CMS"
	SecurityIDSourceClst  SecurityIDSource = "CLST"
	SecurityIDSourceOpra  SecurityIDSource = "OPRA"
	SecurityIDSourceFigi  SecurityIDSource = "FIGI"
	SecurityIDSourceCusip SecurityIDSource = "CUSIP"
	SecurityIDSourceOther SecurityIDSource = "OTHER"
)

// Security type
type SecurityType string

const (
	SecurityTypeCommonStock    SecurityType = "COMMON_STOCK"
	SecurityTypePreferredStock SecurityType = "PREFERRED_STOCK"
	SecurityTypeCorporateBond  SecurityType = "CORPORATE_BOND"
	SecurityTypeOption         SecurityType = "OPTION"
	SecurityTypeFuture         SecurityType = "FUTURE"
	SecurityTypeWarrant        SecurityType = "WARRANT"
	SecurityTypeCash           SecurityType = "CASH"
	SecurityTypeOther          SecurityType = "OTHER"
)
