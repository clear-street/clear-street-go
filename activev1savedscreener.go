// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"encoding/json"
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

// Search and manage saved screeners.
//
// ActiveV1SavedScreenerService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1SavedScreenerService] method instead.
type ActiveV1SavedScreenerService struct {
	options []option.RequestOption
}

// NewActiveV1SavedScreenerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1SavedScreenerService(opts ...option.RequestOption) (r ActiveV1SavedScreenerService) {
	r = ActiveV1SavedScreenerService{}
	r.options = opts
	return
}

// Create a saved screener configuration.
//
// Persists a screener configuration for the authenticated user.
func (r *ActiveV1SavedScreenerService) NewScreener(ctx context.Context, body ActiveV1SavedScreenerNewScreenerParams, opts ...option.RequestOption) (res *ActiveV1SavedScreenerNewScreenerResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/saved-screeners"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete a saved screener configuration.
//
// Deletes the screener configuration for the authenticated user.
func (r *ActiveV1SavedScreenerService) DeleteScreener(ctx context.Context, screenerID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if screenerID == "" {
		err = errors.New("missing required screener_id parameter")
		return err
	}
	path := fmt.Sprintf("active/v1/saved-screeners/%s", screenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get a saved screener configuration by ID.
//
// Returns a single screener configuration for the authenticated user.
func (r *ActiveV1SavedScreenerService) GetScreenerByID(ctx context.Context, screenerID string, opts ...option.RequestOption) (res *ActiveV1SavedScreenerGetScreenerByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if screenerID == "" {
		err = errors.New("missing required screener_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/saved-screeners/%s", screenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List saved screener configurations.
//
// Returns all screener configurations for the authenticated user.
func (r *ActiveV1SavedScreenerService) GetScreeners(ctx context.Context, opts ...option.RequestOption) (res *ActiveV1SavedScreenerGetScreenersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/saved-screeners"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a saved screener configuration.
//
// Replaces the screener configuration for the authenticated user. If `name` is
// null, the existing name is preserved.
func (r *ActiveV1SavedScreenerService) ReplaceScreener(ctx context.Context, screenerID string, body ActiveV1SavedScreenerReplaceScreenerParams, opts ...option.RequestOption) (res *ActiveV1SavedScreenerReplaceScreenerResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if screenerID == "" {
		err = errors.New("missing required screener_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/saved-screeners/%s", screenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// A single filter criterion for a screener
type SavedScreenerFilter struct {
	// The field name to filter on
	FieldName string `json:"field_name" api:"required"`
	// The filter operation (lt, lte, gt, gte, eq, rgx, bw, ew)
	Operation string `json:"operation" api:"required"`
	// The filter value
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldName   respjson.Field
		Operation   respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SavedScreenerFilter) RawJSON() string { return r.JSON.raw }
func (r *SavedScreenerFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SavedScreenerFilter to a SavedScreenerFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SavedScreenerFilterParam.Overrides()
func (r SavedScreenerFilter) ToParam() SavedScreenerFilterParam {
	return param.Override[SavedScreenerFilterParam](json.RawMessage(r.RawJSON()))
}

// A single filter criterion for a screener
//
// The properties FieldName, Operation, Value are required.
type SavedScreenerFilterParam struct {
	// The field name to filter on
	FieldName string `json:"field_name" api:"required"`
	// The filter operation (lt, lte, gt, gte, eq, rgx, bw, ew)
	Operation string `json:"operation" api:"required"`
	// The filter value
	Value string `json:"value" api:"required"`
	paramObj
}

func (r SavedScreenerFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow SavedScreenerFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SavedScreenerFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A saved screener configuration entry
type ScreenerEntry struct {
	// Unique identifier for this screener
	ID string `json:"id" api:"required" format:"uuid"`
	// When this screener was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Filter criteria for this screener
	Filters []SavedScreenerFilter `json:"filters" api:"required"`
	// The name of this screener configuration
	Name string `json:"name" api:"required"`
	// When this screener was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// List of field names to include when running this screener
	FieldFilter []string `json:"field_filter" api:"nullable"`
	// Field name to sort results by
	SortBy string `json:"sort_by" api:"nullable"`
	// Sort direction for results
	SortDirection string `json:"sort_direction" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Filters       respjson.Field
		Name          respjson.Field
		UpdatedAt     respjson.Field
		FieldFilter   respjson.Field
		SortBy        respjson.Field
		SortDirection respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScreenerEntry) RawJSON() string { return r.JSON.raw }
func (r *ScreenerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScreenerEntryList []ScreenerEntry

type ActiveV1SavedScreenerNewScreenerResponse struct {
	// A saved screener configuration entry
	Data ScreenerEntry `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1SavedScreenerNewScreenerResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1SavedScreenerNewScreenerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1SavedScreenerGetScreenerByIDResponse struct {
	// A saved screener configuration entry
	Data ScreenerEntry `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1SavedScreenerGetScreenerByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1SavedScreenerGetScreenerByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1SavedScreenerGetScreenersResponse struct {
	Data ScreenerEntryList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1SavedScreenerGetScreenersResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1SavedScreenerGetScreenersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1SavedScreenerReplaceScreenerResponse struct {
	// A saved screener configuration entry
	Data ScreenerEntry `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1SavedScreenerReplaceScreenerResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1SavedScreenerReplaceScreenerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1SavedScreenerNewScreenerParams struct {
	// The name for this screener configuration
	Name param.Opt[string] `json:"name,omitzero"`
	// Field name to sort results by
	SortBy param.Opt[string] `json:"sort_by,omitzero"`
	// List of field names to include when running this screener
	FieldFilter []string `json:"field_filter,omitzero"`
	// Filter criteria for this screener
	Filters []SavedScreenerFilterParam `json:"filters,omitzero"`
	// Sort direction for results
	//
	// Any of "ASC", "DESC".
	SortDirection ActiveV1SavedScreenerNewScreenerParamsSortDirection `json:"sort_direction,omitzero"`
	paramObj
}

func (r ActiveV1SavedScreenerNewScreenerParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1SavedScreenerNewScreenerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1SavedScreenerNewScreenerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sort direction for results
type ActiveV1SavedScreenerNewScreenerParamsSortDirection string

const (
	ActiveV1SavedScreenerNewScreenerParamsSortDirectionAsc  ActiveV1SavedScreenerNewScreenerParamsSortDirection = "ASC"
	ActiveV1SavedScreenerNewScreenerParamsSortDirectionDesc ActiveV1SavedScreenerNewScreenerParamsSortDirection = "DESC"
)

type ActiveV1SavedScreenerReplaceScreenerParams struct {
	// The name for this screener configuration
	Name param.Opt[string] `json:"name,omitzero"`
	// Field name to sort results by
	SortBy param.Opt[string] `json:"sort_by,omitzero"`
	// List of field names to include when running this screener
	FieldFilter []string `json:"field_filter,omitzero"`
	// Filter criteria for this screener
	Filters []SavedScreenerFilterParam `json:"filters,omitzero"`
	// Sort direction for results
	//
	// Any of "ASC", "DESC".
	SortDirection ActiveV1SavedScreenerReplaceScreenerParamsSortDirection `json:"sort_direction,omitzero"`
	paramObj
}

func (r ActiveV1SavedScreenerReplaceScreenerParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1SavedScreenerReplaceScreenerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1SavedScreenerReplaceScreenerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sort direction for results
type ActiveV1SavedScreenerReplaceScreenerParamsSortDirection string

const (
	ActiveV1SavedScreenerReplaceScreenerParamsSortDirectionAsc  ActiveV1SavedScreenerReplaceScreenerParamsSortDirection = "ASC"
	ActiveV1SavedScreenerReplaceScreenerParamsSortDirectionDesc ActiveV1SavedScreenerReplaceScreenerParamsSortDirection = "DESC"
)
