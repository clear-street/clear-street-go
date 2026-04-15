// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// Endpoints for API service metadata.
//
// ActiveV1VersionService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1VersionService] method instead.
type ActiveV1VersionService struct {
	options []option.RequestOption
}

// NewActiveV1VersionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActiveV1VersionService(opts ...option.RequestOption) (r ActiveV1VersionService) {
	r = ActiveV1VersionService{}
	r.options = opts
	return
}

// Returns the current version string for this API endpoint.
func (r *ActiveV1VersionService) GetVersion(ctx context.Context, opts ...option.RequestOption) (res *ActiveV1VersionGetVersionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Allows clients to set their preferred API version.
func (r *ActiveV1VersionService) UpdateVersion(ctx context.Context, opts ...option.RequestOption) (res *ActiveV1VersionUpdateVersionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return res, err
}

// API version information
type Version struct {
	// API version string
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Version) RawJSON() string { return r.JSON.raw }
func (r *Version) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1VersionGetVersionResponse struct {
	// API version information
	Data Version `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1VersionGetVersionResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1VersionGetVersionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1VersionUpdateVersionResponse struct {
	// API version information
	Data Version `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1VersionUpdateVersionResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1VersionUpdateVersionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
