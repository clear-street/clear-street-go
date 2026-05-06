// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"github.com/clear-street/clear-street-go/option"
)

// ActiveV1Service contains methods and other services that help with interacting
// with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1Service] method instead.
type ActiveV1Service struct {
	options  []option.RequestOption
	Accounts ActiveV1AccountService
}

// NewActiveV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActiveV1Service(opts ...option.RequestOption) (r ActiveV1Service) {
	r = ActiveV1Service{}
	r.options = opts
	r.Accounts = NewActiveV1AccountService(opts...)
	return
}
