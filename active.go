// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"github.com/stainless-sdks/clear-street-go/option"
)

// ActiveService contains methods and other services that help with interacting
// with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveService] method instead.
type ActiveService struct {
	options []option.RequestOption
	V1      ActiveV1Service
}

// NewActiveService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewActiveService(opts ...option.RequestOption) (r ActiveService) {
	r = ActiveService{}
	r.options = opts
	r.V1 = NewActiveV1Service(opts...)
	return
}
