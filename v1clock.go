// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Access financial calendars for events like earnings, dividends, and splits.
//
// V1ClockService contains methods and other services that help with interacting
// with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ClockService] method instead.
type V1ClockService struct {
	options []option.RequestOption
}

// NewV1ClockService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1ClockService(opts ...option.RequestOption) (r V1ClockService) {
	r = V1ClockService{}
	r.options = opts
	return
}

// Returns the current server time in UTC.
func (r *V1ClockService) GetClock(ctx context.Context, opts ...option.RequestOption) (res *V1ClockGetClockResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/clock"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Current server time and market clock information
type ClockDetail struct {
	// Current server time in UTC
	Clock time.Time `json:"clock" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clock       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClockDetail) RawJSON() string { return r.JSON.raw }
func (r *ClockDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ClockGetClockResponse struct {
	// Current server time and market clock information
	Data ClockDetail `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1ClockGetClockResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ClockGetClockResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
