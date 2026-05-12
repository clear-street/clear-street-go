// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/internal/apiquery"
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/param"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Create and manage watchlists.
//
// V1WatchlistService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1WatchlistService] method instead.
type V1WatchlistService struct {
	options []option.RequestOption
}

// NewV1WatchlistService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1WatchlistService(opts ...option.RequestOption) (r V1WatchlistService) {
	r = V1WatchlistService{}
	r.options = opts
	return
}

// Add an instrument to a watchlist
func (r *V1WatchlistService) AddWatchlistItem(ctx context.Context, watchlistID string, body V1WatchlistAddWatchlistItemParams, opts ...option.RequestOption) (res *V1WatchlistAddWatchlistItemResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if watchlistID == "" {
		err = errors.New("missing required watchlist_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/watchlists/%s/items", watchlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create Watchlist
func (r *V1WatchlistService) NewWatchlist(ctx context.Context, body V1WatchlistNewWatchlistParams, opts ...option.RequestOption) (res *V1WatchlistNewWatchlistResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/watchlists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete a watchlist and all its items
func (r *V1WatchlistService) DeleteWatchlist(ctx context.Context, watchlistID string, opts ...option.RequestOption) (res *V1WatchlistDeleteWatchlistResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if watchlistID == "" {
		err = errors.New("missing required watchlist_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/watchlists/%s", watchlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Delete an instrument from a watchlist
func (r *V1WatchlistService) DeleteWatchlistItem(ctx context.Context, itemID string, body V1WatchlistDeleteWatchlistItemParams, opts ...option.RequestOption) (res *V1WatchlistDeleteWatchlistItemResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.WatchlistID == "" {
		err = errors.New("missing required watchlist_id parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/watchlists/%s/items/%s", body.WatchlistID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get a watchlist by ID with all its items
func (r *V1WatchlistService) GetWatchlistByID(ctx context.Context, watchlistID string, opts ...option.RequestOption) (res *V1WatchlistGetWatchlistByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if watchlistID == "" {
		err = errors.New("missing required watchlist_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/watchlists/%s", watchlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List watchlists for the authenticated user
func (r *V1WatchlistService) GetWatchlists(ctx context.Context, query V1WatchlistGetWatchlistsParams, opts ...option.RequestOption) (res *V1WatchlistGetWatchlistsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/watchlists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Response data for adding a watchlist item
type AddWatchlistItemData struct {
	// ID of the created item
	ItemID string `json:"item_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddWatchlistItemData) RawJSON() string { return r.JSON.raw }
func (r *AddWatchlistItemData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed watchlist with all items
type WatchlistDetail struct {
	// Watchlist ID
	ID string `json:"id" api:"required" format:"uuid"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Items in the watchlist
	Items []WatchlistItemEntry `json:"items" api:"required"`
	// Watchlist name
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Items       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WatchlistDetail) RawJSON() string { return r.JSON.raw }
func (r *WatchlistDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a user watchlist.
type WatchlistEntry struct {
	// The unique identifier for the watchlist.
	ID string `json:"id" api:"required" format:"uuid"`
	// The timestamp when the watchlist was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The user-provided watchlist name.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WatchlistEntry) RawJSON() string { return r.JSON.raw }
func (r *WatchlistEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WatchlistEntryList []WatchlistEntry

// A single item in a watchlist
type WatchlistItemEntry struct {
	// Item ID
	ID string `json:"id" api:"required" format:"uuid"`
	// When the item was added
	AddedAt time.Time `json:"added_at" api:"required" format:"date-time"`
	// Price when the item was added
	AddedPrice string `json:"added_price" api:"nullable"`
	// Instrument details
	Instrument Instrument `json:"instrument" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AddedAt     respjson.Field
		AddedPrice  respjson.Field
		Instrument  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WatchlistItemEntry) RawJSON() string { return r.JSON.raw }
func (r *WatchlistItemEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WatchlistAddWatchlistItemResponse struct {
	// Response data for adding a watchlist item
	Data AddWatchlistItemData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1WatchlistAddWatchlistItemResponse) RawJSON() string { return r.JSON.raw }
func (r *V1WatchlistAddWatchlistItemResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WatchlistNewWatchlistResponse struct {
	// Represents a user watchlist.
	Data WatchlistEntry `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1WatchlistNewWatchlistResponse) RawJSON() string { return r.JSON.raw }
func (r *V1WatchlistNewWatchlistResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WatchlistDeleteWatchlistResponse = any

type V1WatchlistDeleteWatchlistItemResponse = any

type V1WatchlistGetWatchlistByIDResponse struct {
	// Detailed watchlist with all items
	Data WatchlistDetail `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1WatchlistGetWatchlistByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1WatchlistGetWatchlistByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WatchlistGetWatchlistsResponse struct {
	Data WatchlistEntryList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1WatchlistGetWatchlistsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1WatchlistGetWatchlistsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WatchlistAddWatchlistItemParams struct {
	// OEMS instrument UUID
	InstrumentID InstrumentIDOrSymbol `json:"instrument_id" api:"required" format:"uuid"`
	paramObj
}

func (r V1WatchlistAddWatchlistItemParams) MarshalJSON() (data []byte, err error) {
	type shadow V1WatchlistAddWatchlistItemParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1WatchlistAddWatchlistItemParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WatchlistNewWatchlistParams struct {
	// The desired watchlist name.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r V1WatchlistNewWatchlistParams) MarshalJSON() (data []byte, err error) {
	type shadow V1WatchlistNewWatchlistParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1WatchlistNewWatchlistParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WatchlistDeleteWatchlistItemParams struct {
	WatchlistID string `path:"watchlist_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type V1WatchlistGetWatchlistsParams struct {
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	paramObj
}

// URLQuery serializes [V1WatchlistGetWatchlistsParams]'s query parameters as
// `url.Values`.
func (r V1WatchlistGetWatchlistsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
