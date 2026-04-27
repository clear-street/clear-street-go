// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/param"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Manage API keys for authentication.
//
// ActiveV1APIKeyService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1APIKeyService] method instead.
type ActiveV1APIKeyService struct {
	options []option.RequestOption
}

// NewActiveV1APIKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActiveV1APIKeyService(opts ...option.RequestOption) (r ActiveV1APIKeyService) {
	r = ActiveV1APIKeyService{}
	r.options = opts
	return
}

// Create a new API key
func (r *ActiveV1APIKeyService) New(ctx context.Context, body ActiveV1APIKeyNewParams, opts ...option.RequestOption) (res *ActiveV1APIKeyNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/api_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List API keys for the authenticated user
func (r *ActiveV1APIKeyService) List(ctx context.Context, opts ...option.RequestOption) (res *ActiveV1APIKeyListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/api_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Revoke a specific API key
func (r *ActiveV1APIKeyService) Revoke(ctx context.Context, id string, opts ...option.RequestOption) (res *ActiveV1APIKeyRevokeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/api_keys/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Revoke all API keys for the authenticated user
func (r *ActiveV1APIKeyService) RevokeAll(ctx context.Context, opts ...option.RequestOption) (res *ActiveV1APIKeyRevokeAllResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/api_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type APIKey struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	APIKey    string    `json:"api_key" api:"required" format:"password"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		APIKey      respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKey) RawJSON() string { return r.JSON.raw }
func (r *APIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyListEntry struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"nullable"`
	RevokedAt time.Time `json:"revoked_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Name        respjson.Field
		RevokedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyListEntry) RawJSON() string { return r.JSON.raw }
func (r *APIKeyListEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyListEntryList []APIKeyListEntry

type Revocation struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	RevokedAt time.Time `json:"revoked_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		RevokedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Revocation) RawJSON() string { return r.JSON.raw }
func (r *Revocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RevocationList []Revocation

type ActiveV1APIKeyNewResponse struct {
	Data APIKey `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1APIKeyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1APIKeyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1APIKeyListResponse struct {
	Data APIKeyListEntryList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1APIKeyListResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1APIKeyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1APIKeyRevokeResponse struct {
	Data Revocation `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1APIKeyRevokeResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1APIKeyRevokeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1APIKeyRevokeAllResponse struct {
	Data RevocationList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1APIKeyRevokeAllResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1APIKeyRevokeAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1APIKeyNewParams struct {
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ActiveV1APIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1APIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1APIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
