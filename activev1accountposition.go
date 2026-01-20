// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/apiquery"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/param"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// ActiveV1AccountPositionService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1AccountPositionService] method instead.
type ActiveV1AccountPositionService struct {
	Options []option.RequestOption
}

// NewActiveV1AccountPositionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1AccountPositionService(opts ...option.RequestOption) (r ActiveV1AccountPositionService) {
	r = ActiveV1AccountPositionService{}
	r.Options = opts
	return
}

// Retrieves all positions for the specified trading account.
func (r *ActiveV1AccountPositionService) ClosePosition(ctx context.Context, securityID string, params ActiveV1AccountPositionClosePositionParams, opts ...option.RequestOption) (res *ActiveV1AccountPositionClosePositionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if securityID == "" {
		err = errors.New("missing required security_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/accounts/%v/positions/%v/%s", params.AccountID, params.SecurityIDSource, securityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Retrieves all positions for the specified trading account.
func (r *ActiveV1AccountPositionService) GetPositions(ctx context.Context, accountID int64, query ActiveV1AccountPositionGetPositionsParams, opts ...option.RequestOption) (res *ActiveV1AccountPositionGetPositionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v/positions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Represents a holding of a particular instrument in an account
type Position struct {
	// The account this position belongs to
	AccountID int64 `json:"account_id,required"`
	// The quantity of a position that is free to be operated on.
	AvailableQuantity string `json:"available_quantity,required"`
	// Type of security
	//
	// Any of "COMMON_STOCK", "PREFERRED_STOCK", "CORPORATE_BOND", "OPTION", "FUTURE",
	// "WARRANT", "CASH", "OTHER".
	InstrumentType SecurityType `json:"instrument_type,required"`
	// The current market value of the position
	MarketValue string `json:"market_value,required"`
	// The type of position
	//
	// Any of "LONG", "SHORT", "LONG_CALL", "SHORT_CALL", "LONG_PUT", "SHORT_PUT".
	PositionType PositionPositionType `json:"position_type,required"`
	// The number of shares or contracts. Can be positive (long) or negative (short)
	Quantity string `json:"quantity,required"`
	// An identifier for the instrument which, when paired with `security_id_source`,
	// identifies one or more financial instruments.
	SecurityID string `json:"security_id,required"`
	// The source of the security identifier
	SecurityIDSource string `json:"security_id_source,required"`
	// The trading symbol for the instrument
	Symbol string `json:"symbol,required"`
	// The MIC code of the primary listing venue
	Venue string `json:"venue,required"`
	// The average price paid per share or contract for this position
	AvgPrice string `json:"avg_price,nullable"`
	// The closing price used to value the position for the last trading day
	ClosingPrice string `json:"closing_price,nullable"`
	// The total cost basis for this position
	CostBasis string `json:"cost_basis,nullable"`
	// The unrealized profit or loss for this position relative to the previous close
	DailyUnrealizedPnl string `json:"daily_unrealized_pnl,nullable"`
	// The current market price of the instrument
	MarketPrice string `json:"market_price,nullable"`
	// The total unrealized profit or loss for this position based on current market
	// value
	UnrealizedPnl string `json:"unrealized_pnl,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID          respjson.Field
		AvailableQuantity  respjson.Field
		InstrumentType     respjson.Field
		MarketValue        respjson.Field
		PositionType       respjson.Field
		Quantity           respjson.Field
		SecurityID         respjson.Field
		SecurityIDSource   respjson.Field
		Symbol             respjson.Field
		Venue              respjson.Field
		AvgPrice           respjson.Field
		ClosingPrice       respjson.Field
		CostBasis          respjson.Field
		DailyUnrealizedPnl respjson.Field
		MarketPrice        respjson.Field
		UnrealizedPnl      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Position) RawJSON() string { return r.JSON.raw }
func (r *Position) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of position
type PositionPositionType string

const (
	PositionPositionTypeLong      PositionPositionType = "LONG"
	PositionPositionTypeShort     PositionPositionType = "SHORT"
	PositionPositionTypeLongCall  PositionPositionType = "LONG_CALL"
	PositionPositionTypeShortCall PositionPositionType = "SHORT_CALL"
	PositionPositionTypeLongPut   PositionPositionType = "LONG_PUT"
	PositionPositionTypeShortPut  PositionPositionType = "SHORT_PUT"
)

type PositionList []Position

type ActiveV1AccountPositionClosePositionResponse struct {
	Data OrderList `json:"data,required"`
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

type ActiveV1AccountPositionGetPositionsResponse struct {
	Data PositionList `json:"data,required"`
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
	AccountID int64 `path:"account_id,required" json:"-"`
	// Security identifier source
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "OTHER".
	SecurityIDSource SecurityIDSource `path:"security_id_source,omitzero,required" json:"-"`
	// The number of items to return per page (only used when page_token is not
	// provided)
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1AccountPositionClosePositionParams]'s query
// parameters as `url.Values`.
func (r ActiveV1AccountPositionClosePositionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ActiveV1AccountPositionGetPositionsParams struct {
	// The number of items to return per page (only used when page_token is not
	// provided)
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1AccountPositionGetPositionsParams]'s query
// parameters as `url.Values`.
func (r ActiveV1AccountPositionGetPositionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
