// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
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

// Manage trading accounts, balances, and portfolio history.
//
// V1AccountPortfolioHistoryService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1AccountPortfolioHistoryService] method instead.
type V1AccountPortfolioHistoryService struct {
	options []option.RequestOption
}

// NewV1AccountPortfolioHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1AccountPortfolioHistoryService(opts ...option.RequestOption) (r V1AccountPortfolioHistoryService) {
	r = V1AccountPortfolioHistoryService{}
	r.options = opts
	return
}

// Retrieves daily portfolio history for the specified account.
func (r *V1AccountPortfolioHistoryService) GetPortfolioHistory(ctx context.Context, accountID int64, query V1AccountPortfolioHistoryGetPortfolioHistoryParams, opts ...option.RequestOption) (res *V1AccountPortfolioHistoryGetPortfolioHistoryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/portfolio-history", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type PortfolioHistoryResponse struct {
	Segments []PortfolioHistorySegment `json:"segments" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Segments    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortfolioHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *PortfolioHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortfolioHistorySegment struct {
	// The date for this segment
	Date time.Time `json:"date" api:"required" format:"date"`
	// The equity at the end of the trading day.
	EodEquity string `json:"eod_equity" api:"required"`
	// Sum of the profit and loss realized from position closing trading activity.
	RealizedPnl string `json:"realized_pnl" api:"required"`
	// The equity at the start of the trading day.
	SodEquity string `json:"sod_equity" api:"required"`
	// Sum of the profit and loss from market changes.
	UnrealizedPnl string `json:"unrealized_pnl" api:"required"`
	// Amount bought MTM
	BoughtNotional string `json:"bought_notional" api:"nullable"`
	// Sum of the profit and loss from intraday trading activities for the trading day.
	DayPnl string `json:"day_pnl" api:"nullable"`
	// P&L after netting all realized and unrealized P&L, adjustments, dividends,
	// change in accruals, income and expenses
	NetPnl string `json:"net_pnl" api:"nullable"`
	// P&L attributable to start-of-day (carried) positions from market movement during
	// this trading day.
	PositionPnl string `json:"position_pnl" api:"nullable"`
	// Amount sold MTM
	SoldNotional string `json:"sold_notional" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date           respjson.Field
		EodEquity      respjson.Field
		RealizedPnl    respjson.Field
		SodEquity      respjson.Field
		UnrealizedPnl  respjson.Field
		BoughtNotional respjson.Field
		DayPnl         respjson.Field
		NetPnl         respjson.Field
		PositionPnl    respjson.Field
		SoldNotional   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortfolioHistorySegment) RawJSON() string { return r.JSON.raw }
func (r *PortfolioHistorySegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountPortfolioHistoryGetPortfolioHistoryResponse struct {
	Data PortfolioHistoryResponse `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AccountPortfolioHistoryGetPortfolioHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AccountPortfolioHistoryGetPortfolioHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AccountPortfolioHistoryGetPortfolioHistoryParams struct {
	StartDate time.Time `query:"start_date" api:"required" format:"date" json:"-"`
	// Defaults to today in America/New_York when omitted.
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [V1AccountPortfolioHistoryGetPortfolioHistoryParams]'s query
// parameters as `url.Values`.
func (r V1AccountPortfolioHistoryGetPortfolioHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
