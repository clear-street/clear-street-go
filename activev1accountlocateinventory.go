// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/apiquery"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// ActiveV1AccountLocateInventoryService contains methods and other services that
// help with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1AccountLocateInventoryService] method instead.
type ActiveV1AccountLocateInventoryService struct {
	Options []option.RequestOption
}

// NewActiveV1AccountLocateInventoryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1AccountLocateInventoryService(opts ...option.RequestOption) (r ActiveV1AccountLocateInventoryService) {
	r = ActiveV1AccountLocateInventoryService{}
	r.Options = opts
	return
}

// Retrieves available inventory for short stock locates.
func (r *ActiveV1AccountLocateInventoryService) GetLocateInventory(ctx context.Context, accountID int64, query ActiveV1AccountLocateInventoryGetLocateInventoryParams, opts ...option.RequestOption) (res *ActiveV1AccountLocateInventoryGetLocateInventoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v/locates/inventory", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Represents the available locate inventory for a symbol
type LocateInventoryItem struct {
	// The account the locate inventory belongs to
	AccountID int64 `json:"account_id,required"`
	// The available quantity of shares that can be located to borrow
	Available int64 `json:"available,required"`
	// The quantity of shares reserved for locate orders that have been `OFFERED` but
	// not yet `FILLED`
	Reserved int64 `json:"reserved,required"`
	// The symbol of the security
	Symbol string `json:"symbol,required"`
	// The quantity of shares that have been `FILLED` and are currently borrowed
	Used int64 `json:"used,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID   respjson.Field
		Available   respjson.Field
		Reserved    respjson.Field
		Symbol      respjson.Field
		Used        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocateInventoryItem) RawJSON() string { return r.JSON.raw }
func (r *LocateInventoryItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocateInventoryItemList []LocateInventoryItem

type ActiveV1AccountLocateInventoryGetLocateInventoryResponse struct {
	Data LocateInventoryItemList `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1AccountLocateInventoryGetLocateInventoryResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountLocateInventoryGetLocateInventoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountLocateInventoryGetLocateInventoryParams struct {
	// The instrument symbol
	Symbol string `query:"symbol,required" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1AccountLocateInventoryGetLocateInventoryParams]'s
// query parameters as `url.Values`.
func (r ActiveV1AccountLocateInventoryGetLocateInventoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
