// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

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
// ActiveV1AccountPositionService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1AccountPositionService] method instead.
type ActiveV1AccountPositionService struct {
	options []option.RequestOption
}

// NewActiveV1AccountPositionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1AccountPositionService(opts ...option.RequestOption) (r ActiveV1AccountPositionService) {
	r = ActiveV1AccountPositionService{}
	r.options = opts
	return
}

// Delete a position within an account for an instrument.
//
// Retrieves orders generated to close the position.
func (r *ActiveV1AccountPositionService) ClosePosition(ctx context.Context, securityID string, params ActiveV1AccountPositionClosePositionParams, opts ...option.RequestOption) (res *ActiveV1AccountPositionClosePositionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if securityID == "" {
		err = errors.New("missing required security_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("active/v1/accounts/%v/positions/%v/%s", params.AccountID, params.SecurityIDSource, url.PathEscape(securityID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Delete all positions within an account.
//
// Closes all positions for the specified trading account.
func (r *ActiveV1AccountPositionService) ClosePositions(ctx context.Context, accountID int64, body ActiveV1AccountPositionClosePositionsParams, opts ...option.RequestOption) (res *ActiveV1AccountPositionClosePositionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v/positions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Retrieves all positions for the specified trading account.
func (r *ActiveV1AccountPositionService) GetPositions(ctx context.Context, accountID int64, query ActiveV1AccountPositionGetPositionsParams, opts ...option.RequestOption) (res *ActiveV1AccountPositionGetPositionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v/positions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Represents a holding of a particular instrument in an account
type Position struct {
	// The account this position belongs to
	AccountID int64 `json:"account_id" api:"required"`
	// The quantity of a position that is free to be operated on.
	AvailableQuantity string `json:"available_quantity" api:"required"`
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
	// An identifier for the instrument which, when paired with `security_id_source`,
	// identifies one or more financial instruments.
	SecurityID string `json:"security_id" api:"required"`
	// The source of the security identifier
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "CURRENCY", "FMP", "OEMS",
	// "SEDOL", "QUIK", "ISIN", "RIC", "COUNTRY", "EXCHANGE", "CTA", "BLOOMBERG",
	// "WERTPAPIER", "DUTCH", "VALOREN", "SICOVAM", "BELGIAN", "COMMON",
	// "CLEARING_HOUSE", "ISDA_FPML_SPECIFICATION", "ISDA_FPML_URL",
	// "LETTER_OF_CREDIT", "MARKETPLACE_ASSIGNED_IDENTIFIER", "MARKIT_RED_ENTITY_CLIP",
	// "MARKIT_RED_PAIR_CLIP", "CFTC", "ISDA_COMMODITY_REFERENCE_PRICE",
	// "LEGAL_ENTITY_IDENTIFIER", "SYNTHETIC", "FIDESSA_INSTRUMENT_MNEMONIC",
	// "INDEX_NAME", "UNIFORM_SYMBOL", "DIGITAL_TOKEN_IDENTIFIER", "MASSIVE", "OTHER".
	SecurityIDSource SecurityIDSource `json:"security_id_source" api:"required"`
	// The trading symbol for the instrument
	Symbol string `json:"symbol" api:"required"`
	// The average price paid per share or contract for this position
	AvgPrice string `json:"avg_price" api:"nullable"`
	// The closing price used to value the position for the last trading day
	ClosingPrice string `json:"closing_price" api:"nullable"`
	// The total cost basis for this position
	CostBasis string `json:"cost_basis" api:"nullable"`
	// The unrealized profit or loss for this position relative to the previous close
	DailyUnrealizedPnl string `json:"daily_unrealized_pnl" api:"nullable"`
	// The unrealized profit/loss for the position for the current day, expressed as a
	// percentage of the baseline value (range: 0-100).
	DailyUnrealizedPnlPct string `json:"daily_unrealized_pnl_pct" api:"nullable"`
	// The current market price of the instrument
	MarketPrice string `json:"market_price" api:"nullable"`
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
		InstrumentType        respjson.Field
		MarketValue           respjson.Field
		PositionType          respjson.Field
		Quantity              respjson.Field
		SecurityID            respjson.Field
		SecurityIDSource      respjson.Field
		Symbol                respjson.Field
		AvgPrice              respjson.Field
		ClosingPrice          respjson.Field
		CostBasis             respjson.Field
		DailyUnrealizedPnl    respjson.Field
		DailyUnrealizedPnlPct respjson.Field
		MarketPrice           respjson.Field
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

type ActiveV1AccountPositionClosePositionResponse struct {
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
func (r ActiveV1AccountPositionClosePositionResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountPositionClosePositionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountPositionClosePositionsResponse struct {
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
func (r ActiveV1AccountPositionClosePositionsResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountPositionClosePositionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountPositionGetPositionsResponse struct {
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
func (r ActiveV1AccountPositionGetPositionsResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1AccountPositionGetPositionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountPositionClosePositionParams struct {
	AccountID int64 `path:"account_id" api:"required" json:"-"`
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
	SecurityIDSource SecurityIDSource `path:"security_id_source,omitzero" api:"required" json:"-"`
	CancelOrders     param.Opt[bool]  `json:"cancel_orders,omitzero"`
	paramObj
}

func (r ActiveV1AccountPositionClosePositionParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1AccountPositionClosePositionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1AccountPositionClosePositionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountPositionClosePositionsParams struct {
	CancelOrders param.Opt[bool] `json:"cancel_orders,omitzero"`
	paramObj
}

func (r ActiveV1AccountPositionClosePositionsParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1AccountPositionClosePositionsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1AccountPositionClosePositionsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountPositionGetPositionsParams struct {
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Filter by security ID(s). Accepts single value or indexed array.
	//
	// Examples:
	//
	// - Single: `security_id=037833100`
	// - Multiple: `security_id[0]=037833100&security_id[1]=594918104`
	SecurityID []string `query:"security_id,omitzero" json:"-"`
	// Source(s) for the security ID filter. Must match the count and order of
	// security_id.
	//
	// Examples:
	//
	// - Single: `security_id_source=CUSIP`
	// - Multiple: `security_id_source[0]=CUSIP&security_id_source[1]=FIGI`
	SecurityIDSource []string `query:"security_id_source,omitzero" json:"-"`
	// Field to sort by
	//
	// Any of "SYMBOL", "INSTRUMENT_TYPE", "QUANTITY", "MARKET_VALUE", "POSITION_TYPE",
	// "UNREALIZED_PNL", "DAILY_UNREALIZED_PNL".
	SortBy ActiveV1AccountPositionGetPositionsParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Sort direction
	//
	// Any of "ASC", "DESC".
	SortDirection ActiveV1AccountPositionGetPositionsParamsSortDirection `query:"sort_direction,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1AccountPositionGetPositionsParams]'s query
// parameters as `url.Values`.
func (r ActiveV1AccountPositionGetPositionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Field to sort by
type ActiveV1AccountPositionGetPositionsParamsSortBy string

const (
	ActiveV1AccountPositionGetPositionsParamsSortBySymbol             ActiveV1AccountPositionGetPositionsParamsSortBy = "SYMBOL"
	ActiveV1AccountPositionGetPositionsParamsSortByInstrumentType     ActiveV1AccountPositionGetPositionsParamsSortBy = "INSTRUMENT_TYPE"
	ActiveV1AccountPositionGetPositionsParamsSortByQuantity           ActiveV1AccountPositionGetPositionsParamsSortBy = "QUANTITY"
	ActiveV1AccountPositionGetPositionsParamsSortByMarketValue        ActiveV1AccountPositionGetPositionsParamsSortBy = "MARKET_VALUE"
	ActiveV1AccountPositionGetPositionsParamsSortByPositionType       ActiveV1AccountPositionGetPositionsParamsSortBy = "POSITION_TYPE"
	ActiveV1AccountPositionGetPositionsParamsSortByUnrealizedPnl      ActiveV1AccountPositionGetPositionsParamsSortBy = "UNREALIZED_PNL"
	ActiveV1AccountPositionGetPositionsParamsSortByDailyUnrealizedPnl ActiveV1AccountPositionGetPositionsParamsSortBy = "DAILY_UNREALIZED_PNL"
)

// Sort direction
type ActiveV1AccountPositionGetPositionsParamsSortDirection string

const (
	ActiveV1AccountPositionGetPositionsParamsSortDirectionAsc  ActiveV1AccountPositionGetPositionsParamsSortDirection = "ASC"
	ActiveV1AccountPositionGetPositionsParamsSortDirectionDesc ActiveV1AccountPositionGetPositionsParamsSortDirection = "DESC"
)
