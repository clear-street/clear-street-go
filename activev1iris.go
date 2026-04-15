// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"github.com/stainless-sdks/clear-street-go/option"
)

// ActiveV1IrisService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1IrisService] method instead.
type ActiveV1IrisService struct {
	options []option.RequestOption
	// Deprecated /iris/_ routes. Use /omni-ai/_ instead.
	Feedback ActiveV1IrisFeedbackService
	// Deprecated /iris/_ routes. Use /omni-ai/_ instead.
	Runs ActiveV1IrisRunService
	// Deprecated /iris/_ routes. Use /omni-ai/_ instead.
	Threads ActiveV1IrisThreadService
}

// NewActiveV1IrisService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActiveV1IrisService(opts ...option.RequestOption) (r ActiveV1IrisService) {
	r = ActiveV1IrisService{}
	r.options = opts
	r.Feedback = NewActiveV1IrisFeedbackService(opts...)
	r.Runs = NewActiveV1IrisRunService(opts...)
	r.Threads = NewActiveV1IrisThreadService(opts...)
	return
}
