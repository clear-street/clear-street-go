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

// View account positions.
//
// V1AccountPositionService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1AccountPositionService] method instead.
type V1AccountPositionService struct {
	options []option.RequestOption
}

// NewV1AccountPositionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1AccountPositionService(opts ...option.RequestOption) (r V1AccountPositionService) {
	r = V1AccountPositionService{}
	r.options = opts
	return
}

// Delete a position within an account for an instrument.
//
// Retrieves orders generated to close the position.
func (r *V1AccountPositionService) ClosePosition(ctx context.Context, instrumentID string, params V1AccountPositionClosePositionParams, opts ...option.RequestOption) (res *V1AccountPositionClosePositionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%v/positions/%s", params.AccountID, instrumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Delete all positions within an account.
//
// Closes all positions for the specified trading account.
func (r *V1AccountPositionService) ClosePositions(ctx context.Context, accountID int64, body V1AccountPositionClosePositionsParams, opts ...option.RequestOption) (res *V1AccountPositionClosePositionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/positions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Retrieves all positions for the specified trading account.
func (r *V1AccountPositionService) GetPositions(ctx context.Context, accountID int64, query V1AccountPositionGetPositionsParams, opts ...option.RequestOption) (res *V1AccountPositionGetPositionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/positions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Represents a holding of a particular instrument in an account
type Position struct {
	// The account this position belongs to
	AccountID int64 `json:"account_id" api:"required"`
	// The quantity of a position that is free to be operated on.
	AvailableQuantity string `json:"available_quantity" api:"required"`
	// OEMS instrument UUID
	InstrumentID string `json:"instrument_id" api:"required" format:"uuid"`
	// Type of security
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	InstrumentType SecurityType `json:"instrument_type" api:"required"`
	// The current market value of the position
	MarketValue string `json:"market_value" api:"required"`
	// The type of position
	//
	// Any of "LONG", "SHORT", "LONG_CALL", "SHORT_CALL", "LONG_PUT", "SHORT_PUT".
	PositionType PositionType `json:"position_type" api:"required"`
	// The number of shares or contracts. Can be positive (long) or negative (short)
	Quantity string `json:"quantity" api:"required"`
	// The trading symbol for the instrument
	Symbol string `json:"symbol" api:"required"`
	// The average price paid per share or contract for this position
	AvgPrice string `json:"avg_price" api:"nullable"`
	// The closing price used to value the position for the last trading day
	ClosingPrice string `json:"closing_price" api:"nullable"`
	// The market date associated with `closing_price`
	ClosingPriceDate time.Time `json:"closing_price_date" api:"nullable" format:"date"`
	// The total cost basis for this position
	CostBasis string `json:"cost_basis" api:"nullable"`
	// The unrealized profit or loss for this position relative to the previous close
	DailyUnrealizedPnl string `json:"daily_unrealized_pnl" api:"nullable"`
	// The unrealized profit/loss for the position for the current day, expressed as a
	// percentage of the baseline value (range: 0-100).
	DailyUnrealizedPnlPct string `json:"daily_unrealized_pnl_pct" api:"nullable"`
	// The current market price of the instrument
	InstrumentPrice string `json:"instrument_price" api:"nullable"`
	// OEMS instrument identifier of the underlying instrument, if resolvable
	UnderlierInstrumentID string `json:"underlier_instrument_id" api:"nullable" format:"uuid"`
	// The total unrealized profit or loss for this position based on current market
	// value
	UnrealizedPnl string `json:"unrealized_pnl" api:"nullable"`
	// The unrealized profit/loss for the position, expressed as a percentage of the
	// position's cost basis (range: 0-100).
	UnrealizedPnlPct string `json:"unrealized_pnl_pct" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID             respjson.Field
		AvailableQuantity     respjson.Field
		InstrumentID          respjson.Field
		InstrumentType        respjson.Field
		MarketValue           respjson.Field
		PositionType          respjson.Field
		Quantity              respjson.Field
		Symbol                respjson.Field
		AvgPrice              respjson.Field
		ClosingPrice          respjson.Field
		ClosingPriceDate      respjson.Field
		CostBasis             respjson.Field
		DailyUnrealizedPnl    respjson.Field
		DailyUnrealizedPnlPct respjson.Field
		InstrumentPrice       respjson.Field
		UnderlierInstrumentID respjson.Field
		UnrealizedPnl         respjson.Field
		UnrealizedPnlPct      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Position) RawJSON() string { return r.JSON.raw }
func (r *Position) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PositionList []Position

// Position type classification
type PositionType string

const (
	PositionTypeLong      PositionType = "LONG"
	PositionTypeShort     PositionType = "SHORT"
	PositionTypeLongCall  PositionType = "LONG_CALL"
	PositionTypeShortCall PositionType = "SHORT_CALL"
	PositionTypeLongPut   PositionType = "LONG_PUT"
	PositionTypeShortPut  PositionType = "SHORT_PUT"
)

type V1AccountPositionClosePositionResponse struct {
	Data OrderList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountPositionClosePositionResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountPositionClosePositionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountPositionClosePositionsResponse struct {
	Data OrderList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountPositionClosePositionsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountPositionClosePositionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountPositionGetPositionsResponse struct {
	Data PositionList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountPositionGetPositionsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountPositionGetPositionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountPositionClosePositionParams struct {
	AccountID    int64           `path:"account_id" api:"required" json:"-"`
	CancelOrders param.Opt[bool] `json:"cancel_orders,omitzero"`
	paramObj
}

func (r V1AccountPositionClosePositionParams) MarshalJSON() (data []byte, err error) {
	type shadow V1AccountPositionClosePositionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AccountPositionClosePositionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountPositionClosePositionsParams struct {
	CancelOrders param.Opt[bool] `json:"cancel_orders,omitzero"`
	paramObj
}

func (r V1AccountPositionClosePositionsParams) MarshalJSON() (data []byte, err error) {
	type shadow V1AccountPositionClosePositionsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AccountPositionClosePositionsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountPositionGetPositionsParams struct {
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Comma-separated OEMS instrument UUIDs
	InstrumentIDs []string `query:"instrument_ids,omitzero" format:"uuid" json:"-"`
	// Field to sort by
	//
	// Any of "SYMBOL", "INSTRUMENT_TYPE", "QUANTITY", "MARKET_VALUE", "POSITION_TYPE",
	// "UNREALIZED_PNL", "DAILY_UNREALIZED_PNL".
	SortBy V1AccountPositionGetPositionsParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Sort direction
	//
	// Any of "ASC", "DESC".
	SortDirection V1AccountPositionGetPositionsParamsSortDirection `query:"sort_direction,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1AccountPositionGetPositionsParams]'s query parameters as
// `url.Values`.
func (r V1AccountPositionGetPositionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Field to sort by
type V1AccountPositionGetPositionsParamsSortBy string

const (
	V1AccountPositionGetPositionsParamsSortBySymbol             V1AccountPositionGetPositionsParamsSortBy = "SYMBOL"
	V1AccountPositionGetPositionsParamsSortByInstrumentType     V1AccountPositionGetPositionsParamsSortBy = "INSTRUMENT_TYPE"
	V1AccountPositionGetPositionsParamsSortByQuantity           V1AccountPositionGetPositionsParamsSortBy = "QUANTITY"
	V1AccountPositionGetPositionsParamsSortByMarketValue        V1AccountPositionGetPositionsParamsSortBy = "MARKET_VALUE"
	V1AccountPositionGetPositionsParamsSortByPositionType       V1AccountPositionGetPositionsParamsSortBy = "POSITION_TYPE"
	V1AccountPositionGetPositionsParamsSortByUnrealizedPnl      V1AccountPositionGetPositionsParamsSortBy = "UNREALIZED_PNL"
	V1AccountPositionGetPositionsParamsSortByDailyUnrealizedPnl V1AccountPositionGetPositionsParamsSortBy = "DAILY_UNREALIZED_PNL"
)

// Sort direction
type V1AccountPositionGetPositionsParamsSortDirection string

const (
	V1AccountPositionGetPositionsParamsSortDirectionAsc  V1AccountPositionGetPositionsParamsSortDirection = "ASC"
	V1AccountPositionGetPositionsParamsSortDirectionDesc V1AccountPositionGetPositionsParamsSortDirection = "DESC"
)
