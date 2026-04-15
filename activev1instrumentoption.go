// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"github.com/clear-street/clear-street-go/option"
)

// ActiveV1InstrumentOptionService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1InstrumentOptionService] method instead.
type ActiveV1InstrumentOptionService struct {
	options []option.RequestOption
	// Retrieve details and lists of tradable instruments.
	Contracts ActiveV1InstrumentOptionContractService
}

// NewActiveV1InstrumentOptionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1InstrumentOptionService(opts ...option.RequestOption) (r ActiveV1InstrumentOptionService) {
	r = ActiveV1InstrumentOptionService{}
	r.options = opts
	r.Contracts = NewActiveV1InstrumentOptionContractService(opts...)
	return
}
