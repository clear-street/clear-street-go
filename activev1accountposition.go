// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"github.com/clear-street/clear-street-go/option"
)

// ActiveV1AccountPositionService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1AccountPositionService] method instead.
type ActiveV1AccountPositionService struct {
	options      []option.RequestOption
	Instructions ActiveV1AccountPositionInstructionService
}

// NewActiveV1AccountPositionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1AccountPositionService(opts ...option.RequestOption) (r ActiveV1AccountPositionService) {
	r = ActiveV1AccountPositionService{}
	r.options = opts
	r.Instructions = NewActiveV1AccountPositionInstructionService(opts...)
	return
}
