// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/apiquery"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/param"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// ActiveV1ScreenerService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1ScreenerService] method instead.
type ActiveV1ScreenerService struct {
	Options []option.RequestOption
}

// NewActiveV1ScreenerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1ScreenerService(opts ...option.RequestOption) (r ActiveV1ScreenerService) {
	r = ActiveV1ScreenerService{}
	r.Options = opts
	return
}

// Searches for instruments matching specified criteria.
func (r *ActiveV1ScreenerService) GetScreener(ctx context.Context, query ActiveV1ScreenerGetScreenerParams, opts ...option.RequestOption) (res *ActiveV1ScreenerGetScreenerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/screener"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// An instrument returned by the screener
type ScreenerItem struct {
	// The latest price for the instrument
	Price string `json:"price,required"`
	// The identifier for the instrument
	SecurityID string `json:"security_id,required"`
	// The source of the security identifier
	SecurityIDSource string `json:"security_id_source,required"`
	// The trading symbol for the instrument
	Symbol string `json:"symbol,required"`
	// The latest trading volume for the instrument
	Volume string `json:"volume,required"`
	// The ISO country code of the instrument's issue
	CountryOfIssue string `json:"country_of_issue,nullable"`
	// A detailed description of the instrument or company
	Description string `json:"description,nullable"`
	// The specific industry of the instrument's issuer
	Industry string `json:"industry,nullable"`
	// The date the instrument was first listed
	ListDate time.Time `json:"list_date,nullable" format:"date"`
	// The total market capitalization
	MarketCap string `json:"market_cap,nullable"`
	// The average trading volume over the past month
	MonthAvgVolume string `json:"month_avg_volume,nullable"`
	// The full name of the instrument or its issuer
	Name string `json:"name,nullable"`
	// The percent change from previous close to current price
	PercentChange string `json:"percent_change,nullable"`
	// The business sector of the instrument's issuer
	Sector string `json:"sector,nullable"`
	// The type of security
	SecurityType string `json:"security_type,nullable"`
	// The TTM debt-to-equity ratio
	TtmDebtToEquity string `json:"ttm_debt_to_equity,nullable"`
	// The TTM dividend yield percent
	TtmDividendYield string `json:"ttm_dividend_yield,nullable"`
	// The TTM earnings per share
	TtmEarningsPerShare string `json:"ttm_earnings_per_share,nullable"`
	// The TTM price-to-earnings ratio
	TtmPriceToEarnings string `json:"ttm_price_to_earnings,nullable"`
	// The MIC code of the primary listing venue
	Venue string `json:"venue,nullable"`
	// The average trading volume over the past week
	WeekAvgVolume string `json:"week_avg_volume,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Price               respjson.Field
		SecurityID          respjson.Field
		SecurityIDSource    respjson.Field
		Symbol              respjson.Field
		Volume              respjson.Field
		CountryOfIssue      respjson.Field
		Description         respjson.Field
		Industry            respjson.Field
		ListDate            respjson.Field
		MarketCap           respjson.Field
		MonthAvgVolume      respjson.Field
		Name                respjson.Field
		PercentChange       respjson.Field
		Sector              respjson.Field
		SecurityType        respjson.Field
		TtmDebtToEquity     respjson.Field
		TtmDividendYield    respjson.Field
		TtmEarningsPerShare respjson.Field
		TtmPriceToEarnings  respjson.Field
		Venue               respjson.Field
		WeekAvgVolume       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScreenerItem) RawJSON() string { return r.JSON.raw }
func (r *ScreenerItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScreenerItemList []ScreenerItem

type ActiveV1ScreenerGetScreenerResponse struct {
	Data ScreenerItemList `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1ScreenerGetScreenerResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1ScreenerGetScreenerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1ScreenerGetScreenerParams struct {
	// Number of items to return per page (default: 100, max: 10000)
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination
	// state.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Field to sort by
	SortBy param.Opt[string] `query:"sort_by,omitzero" json:"-"`
	// Comma-separated list of field names to include in the response
	FieldFilter []string `query:"field_filter,omitzero" json:"-"`
	// Dynamic filters with dot notation (e.g., filter[price.gte]=50,
	// filter[symbol.bw]=A)
	Filter map[string]string `query:"filter,omitzero" json:"-"`
	// Sort direction (ASC or DESC, defaults to DESC)
	//
	// Any of "ASC", "DESC".
	SortDirection ActiveV1ScreenerGetScreenerParamsSortDirection `query:"sort_direction,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1ScreenerGetScreenerParams]'s query parameters as
// `url.Values`.
func (r ActiveV1ScreenerGetScreenerParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort direction (ASC or DESC, defaults to DESC)
type ActiveV1ScreenerGetScreenerParamsSortDirection string

const (
	ActiveV1ScreenerGetScreenerParamsSortDirectionAsc  ActiveV1ScreenerGetScreenerParamsSortDirection = "ASC"
	ActiveV1ScreenerGetScreenerParamsSortDirectionDesc ActiveV1ScreenerGetScreenerParamsSortDirection = "DESC"
)
