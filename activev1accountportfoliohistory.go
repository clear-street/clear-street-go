// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/apiquery"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// Manage trading accounts and view balances.
//
// ActiveV1AccountPortfolioHistoryService contains methods and other services that
// help with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1AccountPortfolioHistoryService] method instead.
type ActiveV1AccountPortfolioHistoryService struct {
	Options []option.RequestOption
}

// NewActiveV1AccountPortfolioHistoryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1AccountPortfolioHistoryService(opts ...option.RequestOption) (r ActiveV1AccountPortfolioHistoryService) {
	r = ActiveV1AccountPortfolioHistoryService{}
	r.Options = opts
	return
}

// Retrieves daily portfolio history for the specified account.
func (r *ActiveV1AccountPortfolioHistoryService) GetPortfolioHistory(ctx context.Context, accountID int64, query ActiveV1AccountPortfolioHistoryGetPortfolioHistoryParams, opts ...option.RequestOption) (res *ActiveV1AccountPortfolioHistoryGetPortfolioHistoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("active/v1/accounts/%v/portfolio-history", accountID)
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
	EndingEquity string `json:"ending_equity" api:"required"`
	// Sum of the profit and loss realized from position closing trading activity.
	RealizedPnl string `json:"realized_pnl" api:"required"`
	// The equity at the start of the trading day.
	StartingEquity string `json:"starting_equity" api:"required"`
	// Sum of the profit and loss from market changes.
	UnrealizedPnl string `json:"unrealized_pnl" api:"required"`
	// Amount bought MTM
	BoughtNotional string `json:"bought_notional" api:"nullable"`
	// Quantity bought MTM
	BoughtQuantity string `json:"bought_quantity" api:"nullable"`
	// Sum of the profit and loss from intraday trading activities for the trading day.
	DayPnl string `json:"day_pnl" api:"nullable"`
	// P&L after netting all realized and unrealized P&L, adjustments, dividends,
	// change in accruals, income and expenses
	NetPnl string `json:"net_pnl" api:"nullable"`
	// Sum of the profit and loss from the previous trading day.
	PositionPnl string `json:"position_pnl" api:"nullable"`
	// Amount sold MTM
	SoldNotional string `json:"sold_notional" api:"nullable"`
	// Quantity sold MTM
	SoldQuantity string `json:"sold_quantity" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date           respjson.Field
		EndingEquity   respjson.Field
		RealizedPnl    respjson.Field
		StartingEquity respjson.Field
		UnrealizedPnl  respjson.Field
		BoughtNotional respjson.Field
		BoughtQuantity respjson.Field
		DayPnl         respjson.Field
		NetPnl         respjson.Field
		PositionPnl    respjson.Field
		SoldNotional   respjson.Field
		SoldQuantity   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortfolioHistorySegment) RawJSON() string { return r.JSON.raw }
func (r *PortfolioHistorySegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountPortfolioHistoryGetPortfolioHistoryResponse struct {
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
func (r ActiveV1AccountPortfolioHistoryGetPortfolioHistoryResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ActiveV1AccountPortfolioHistoryGetPortfolioHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1AccountPortfolioHistoryGetPortfolioHistoryParams struct {
	EndDate   time.Time `query:"end_date" api:"required" format:"date" json:"-"`
	StartDate time.Time `query:"start_date" api:"required" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1AccountPortfolioHistoryGetPortfolioHistoryParams]'s
// query parameters as `url.Values`.
func (r ActiveV1AccountPortfolioHistoryGetPortfolioHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
