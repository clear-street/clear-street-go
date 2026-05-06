// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/param"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Create and manage watchlists.
//
// V1WatchlistItemService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1WatchlistItemService] method instead.
type V1WatchlistItemService struct {
	options []option.RequestOption
}

// NewV1WatchlistItemService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1WatchlistItemService(opts ...option.RequestOption) (r V1WatchlistItemService) {
	r = V1WatchlistItemService{}
	r.options = opts
	return
}

// Add an instrument to a watchlist
func (r *V1WatchlistItemService) AddWatchlistItem(ctx context.Context, watchlistID string, body V1WatchlistItemAddWatchlistItemParams, opts ...option.RequestOption) (res *V1WatchlistItemAddWatchlistItemResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if watchlistID == "" {
		err = errors.New("missing required watchlist_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/watchlists/%s/items", watchlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete an instrument from a watchlist
func (r *V1WatchlistItemService) DeleteWatchlistItem(ctx context.Context, itemID string, body V1WatchlistItemDeleteWatchlistItemParams, opts ...option.RequestOption) (res *V1WatchlistItemDeleteWatchlistItemResponse, err error) {
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

type V1WatchlistItemAddWatchlistItemResponse struct {
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
func (r V1WatchlistItemAddWatchlistItemResponse) RawJSON() string { return r.JSON.raw }
func (r *V1WatchlistItemAddWatchlistItemResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WatchlistItemDeleteWatchlistItemResponse = any

type V1WatchlistItemAddWatchlistItemParams struct {
	// OEMS instrument UUID
	InstrumentID InstrumentIDOrSymbol `json:"instrument_id" api:"required" format:"uuid"`
	paramObj
}

func (r V1WatchlistItemAddWatchlistItemParams) MarshalJSON() (data []byte, err error) {
	type shadow V1WatchlistItemAddWatchlistItemParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1WatchlistItemAddWatchlistItemParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1WatchlistItemDeleteWatchlistItemParams struct {
	WatchlistID string `path:"watchlist_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
