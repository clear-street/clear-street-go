// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/param"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// ActiveV1WatchlistService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1WatchlistService] method instead.
type ActiveV1WatchlistService struct {
	Options []option.RequestOption
	Items   ActiveV1WatchlistItemService
}

// NewActiveV1WatchlistService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1WatchlistService(opts ...option.RequestOption) (r ActiveV1WatchlistService) {
	r = ActiveV1WatchlistService{}
	r.Options = opts
	r.Items = NewActiveV1WatchlistItemService(opts...)
	return
}

// Create a new watchlist
func (r *ActiveV1WatchlistService) NewWatchlist(ctx context.Context, body ActiveV1WatchlistNewWatchlistParams, opts ...option.RequestOption) (res *ActiveV1WatchlistNewWatchlistResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/watchlists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a watchlist and all its items
func (r *ActiveV1WatchlistService) DeleteWatchlist(ctx context.Context, watchlistID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if watchlistID == "" {
		err = errors.New("missing required watchlist_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/watchlists/%s", watchlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get a watchlist by ID with all its items
func (r *ActiveV1WatchlistService) GetWatchlistByID(ctx context.Context, watchlistID string, opts ...option.RequestOption) (res *ActiveV1WatchlistGetWatchlistByIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if watchlistID == "" {
		err = errors.New("missing required watchlist_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/watchlists/%s", watchlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List watchlists for the authenticated user
func (r *ActiveV1WatchlistService) GetWatchlists(ctx context.Context, opts ...option.RequestOption) (res *ActiveV1WatchlistGetWatchlistsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/watchlists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Detailed watchlist with all items
type WatchlistDetail struct {
	// Watchlist ID
	ID string `json:"id,required" format:"uuid"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Items in the watchlist
	Items []WatchlistItemEntry `json:"items,required"`
	// Watchlist name
	Name string `json:"name,required"`
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
	ID string `json:"id,required" format:"uuid"`
	// The timestamp when the watchlist was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The user-provided watchlist name.
	Name string `json:"name,required"`
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
	ID string `json:"id,required" format:"uuid"`
	// When the item was added
	AddedAt time.Time `json:"added_at,required" format:"date-time"`
	// Price when the item was added
	AddedPrice string `json:"added_price,nullable"`
	// Instrument details
	Instrument Instrument `json:"instrument,nullable"`
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

type ActiveV1WatchlistNewWatchlistResponse struct {
	// Represents a user watchlist.
	Data WatchlistEntry `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1WatchlistNewWatchlistResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1WatchlistNewWatchlistResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1WatchlistGetWatchlistByIDResponse struct {
	// Detailed watchlist with all items
	Data WatchlistDetail `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1WatchlistGetWatchlistByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1WatchlistGetWatchlistByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1WatchlistGetWatchlistsResponse struct {
	Data WatchlistEntryList `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1WatchlistGetWatchlistsResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1WatchlistGetWatchlistsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1WatchlistNewWatchlistParams struct {
	// The desired watchlist name.
	Name string `json:"name,required"`
	paramObj
}

func (r ActiveV1WatchlistNewWatchlistParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1WatchlistNewWatchlistParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1WatchlistNewWatchlistParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
