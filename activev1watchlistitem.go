// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/param"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// ActiveV1WatchlistItemService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1WatchlistItemService] method instead.
type ActiveV1WatchlistItemService struct {
	Options []option.RequestOption
}

// NewActiveV1WatchlistItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1WatchlistItemService(opts ...option.RequestOption) (r ActiveV1WatchlistItemService) {
	r = ActiveV1WatchlistItemService{}
	r.Options = opts
	return
}

// Add an instrument to a watchlist
func (r *ActiveV1WatchlistItemService) AddWatchlistItem(ctx context.Context, watchlistID string, body ActiveV1WatchlistItemAddWatchlistItemParams, opts ...option.RequestOption) (res *ActiveV1WatchlistItemAddWatchlistItemResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if watchlistID == "" {
		err = errors.New("missing required watchlist_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/watchlists/%s/items", watchlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete an instrument from a watchlist
func (r *ActiveV1WatchlistItemService) DeleteWatchlistItem(ctx context.Context, itemID string, body ActiveV1WatchlistItemDeleteWatchlistItemParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.WatchlistID == "" {
		err = errors.New("missing required watchlist_id parameter")
		return
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/watchlists/%s/items/%s", body.WatchlistID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Response data for adding a watchlist item
type AddWatchlistItemData struct {
	// ID of the created item
	ItemID string `json:"item_id,required" format:"uuid"`
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

type ActiveV1WatchlistItemAddWatchlistItemResponse struct {
	// Response data for adding a watchlist item
	Data AddWatchlistItemData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1WatchlistItemAddWatchlistItemResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1WatchlistItemAddWatchlistItemResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1WatchlistItemAddWatchlistItemParams struct {
	// OEMS instrument ID (mutually exclusive with security_id/security_id_source)
	InstrumentID param.Opt[string] `json:"instrument_id,omitzero" format:"uuid"`
	// Security identifier
	SecurityID param.Opt[string] `json:"security_id,omitzero"`
	// Security identifier source
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "CURRENCY", "FMP", "OEMS",
	// "SEDOL", "QUIK", "ISIN", "RIC", "COUNTRY", "EXCHANGE", "CTA", "BLOOMBERG",
	// "WERTPAPIER", "DUTCH", "VALOREN", "SICOVAM", "BELGIAN", "COMMON",
	// "CLEARING_HOUSE", "ISDA_FPML_SPECIFICATION", "ISDA_FPML_URL",
	// "LETTER_OF_CREDIT", "MARKETPLACE_ASSIGNED_IDENTIFIER", "MARKIT_RED_ENTITY_CLIP",
	// "MARKIT_RED_PAIR_CLIP", "CFTC", "ISDA_COMMODITY_REFERENCE_PRICE",
	// "LEGAL_ENTITY_IDENTIFIER", "SYNTHETIC", "FIDESSA_INSTRUMENT_MNEMONIC",
	// "INDEX_NAME", "UNIFORM_SYMBOL", "DIGITAL_TOKEN_IDENTIFIER", "MASSIVE", "OTHER".
	SecurityIDSource SecurityIDSource `json:"security_id_source,omitzero"`
	paramObj
}

func (r ActiveV1WatchlistItemAddWatchlistItemParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1WatchlistItemAddWatchlistItemParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1WatchlistItemAddWatchlistItemParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1WatchlistItemDeleteWatchlistItemParams struct {
	WatchlistID string `path:"watchlist_id,required" format:"uuid" json:"-"`
	paramObj
}
