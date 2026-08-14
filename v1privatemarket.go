// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"github.com/clear-street/clear-street-go/option"
)

// V1PrivateMarketService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1PrivateMarketService] method instead.
type V1PrivateMarketService struct {
	options []option.RequestOption
	// Browse private-market offerings and their indicative terms. Access requires the
	// account holder to hold an accreditation attestation.
	Offerings V1PrivateMarketOfferingService
}

// NewV1PrivateMarketService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1PrivateMarketService(opts ...option.RequestOption) (r V1PrivateMarketService) {
	r = V1PrivateMarketService{}
	r.options = opts
	r.Offerings = NewV1PrivateMarketOfferingService(opts...)
	return
}
