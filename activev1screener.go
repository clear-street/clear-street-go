// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
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

// Retrieve details and lists of tradable instruments.
//
// ActiveV1ScreenerService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1ScreenerService] method instead.
type ActiveV1ScreenerService struct {
	options []option.RequestOption
}

// NewActiveV1ScreenerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1ScreenerService(opts ...option.RequestOption) (r ActiveV1ScreenerService) {
	r = ActiveV1ScreenerService{}
	r.options = opts
	return
}

// Searches for instruments matching specified criteria.
func (r *ActiveV1ScreenerService) GetScreener(ctx context.Context, query ActiveV1ScreenerGetScreenerParams, opts ...option.RequestOption) (res *ActiveV1ScreenerGetScreenerResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/screener"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A single filter criterion for the screener.
type ScreenerFilter struct {
	// Field to filter on (e.g., "market_cap", "sector", "price")
	Field string `json:"field" api:"required"`
	// Comparison operator (e.g., "eq", "gte", "lte", "in")
	Operator string `json:"operator" api:"required"`
	// Filter value
	Value any `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScreenerFilter) RawJSON() string { return r.JSON.raw }
func (r *ScreenerFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An instrument returned by the screener
type ScreenerItem struct {
	// The count of buy analyst ratings
	BuyRatings int64 `json:"buy_ratings" api:"required"`
	// The count of hold analyst ratings
	HoldRatings int64 `json:"hold_ratings" api:"required"`
	// The latest price for the instrument
	Price string `json:"price" api:"required"`
	// The identifier for the instrument
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
	// The count of sell analyst ratings
	SellRatings int64 `json:"sell_ratings" api:"required"`
	// The count of strong buy analyst ratings
	StrongBuyRatings int64 `json:"strong_buy_ratings" api:"required"`
	// The count of strong sell analyst ratings
	StrongSellRatings int64 `json:"strong_sell_ratings" api:"required"`
	// The trading symbol for the instrument
	Symbol string `json:"symbol" api:"required"`
	// The total count of analyst ratings
	TotalRatings int64 `json:"total_ratings" api:"required"`
	// The consensus analyst price target
	ConsensusPriceTarget string `json:"consensus_price_target" api:"nullable"`
	// The highest analyst price target
	ConsensusPriceTargetHigh string `json:"consensus_price_target_high" api:"nullable"`
	// The lowest analyst price target
	ConsensusPriceTargetLow string `json:"consensus_price_target_low" api:"nullable"`
	// The consensus analyst rating
	//
	// Any of "STRONG_BUY", "BUY", "HOLD", "SELL", "STRONG_SELL".
	ConsensusRating AnalystRating `json:"consensus_rating" api:"nullable"`
	// The ISO country code of the instrument's issue
	CountryOfIssue string `json:"country_of_issue" api:"nullable"`
	// A detailed description of the instrument or company
	Description string `json:"description" api:"nullable"`
	// The highest price over the last 52 weeks
	FiftyTwoWeekHigh string `json:"fifty_two_week_high" api:"nullable"`
	// The lowest price over the last 52 weeks
	FiftyTwoWeekLow string `json:"fifty_two_week_low" api:"nullable"`
	// Percent gap from 52-week high to previous day close (negative = below high)
	GapFrom52wHighPct string `json:"gap_from_52w_high_pct" api:"nullable"`
	// Percent gap from 52-week low to previous day close (positive = above low)
	GapFrom52wLowPct string `json:"gap_from_52w_low_pct" api:"nullable"`
	// The specific industry of the instrument's issuer
	Industry string `json:"industry" api:"nullable"`
	// The date the instrument was first listed
	ListDate time.Time `json:"list_date" api:"nullable" format:"date"`
	// The total market capitalization
	MarketCap string `json:"market_cap" api:"nullable"`
	// The average trading volume over the past month
	MonthAvgVolume string `json:"month_avg_volume" api:"nullable"`
	// The full name of the instrument or its issuer
	Name string `json:"name" api:"nullable"`
	// The closing price approximately one month ago
	OneMonthAgoClose string `json:"one_month_ago_close" api:"nullable"`
	// The opening price approximately one month ago
	OneMonthAgoOpen string `json:"one_month_ago_open" api:"nullable"`
	// Percent change from one month ago close to previous day close
	OneMonthChangePct string `json:"one_month_change_pct" api:"nullable"`
	// The closing price approximately one week ago
	OneWeekAgoClose string `json:"one_week_ago_close" api:"nullable"`
	// The opening price approximately one week ago
	OneWeekAgoOpen string `json:"one_week_ago_open" api:"nullable"`
	// Percent change from one week ago close to previous day close
	OneWeekChangePct string `json:"one_week_change_pct" api:"nullable"`
	// The closing price approximately one year ago
	OneYearAgoClose string `json:"one_year_ago_close" api:"nullable"`
	// The opening price approximately one year ago
	OneYearAgoOpen string `json:"one_year_ago_open" api:"nullable"`
	// Percent change from one year ago close to previous day close
	OneYearChangePct string `json:"one_year_change_pct" api:"nullable"`
	// The percent change from previous close to current price
	PercentChange string `json:"percent_change" api:"nullable"`
	// The previous day's closing price
	PrevDayClose string `json:"prev_day_close" api:"nullable"`
	// The business sector of the instrument's issuer
	Sector string `json:"sector" api:"nullable"`
	// The type of security
	SecurityType string `json:"security_type" api:"nullable"`
	// Percent change from six months ago close to previous day close
	SixMonthChangePct string `json:"six_month_change_pct" api:"nullable"`
	// The closing price approximately six months ago
	SixMonthsAgoClose string `json:"six_months_ago_close" api:"nullable"`
	// The opening price approximately six months ago
	SixMonthsAgoOpen string `json:"six_months_ago_open" api:"nullable"`
	// Percent change from three months ago close to previous day close
	ThreeMonthChangePct string `json:"three_month_change_pct" api:"nullable"`
	// The closing price approximately three months ago
	ThreeMonthsAgoClose string `json:"three_months_ago_close" api:"nullable"`
	// The opening price approximately three months ago
	ThreeMonthsAgoOpen string `json:"three_months_ago_open" api:"nullable"`
	// The TTM debt-to-equity ratio
	TtmDebtToEquity string `json:"ttm_debt_to_equity" api:"nullable"`
	// The TTM dividend yield percent
	TtmDividendYield string `json:"ttm_dividend_yield" api:"nullable"`
	// The TTM earnings per share
	TtmEarningsPerShare string `json:"ttm_earnings_per_share" api:"nullable"`
	// The TTM price-to-earnings ratio
	TtmPriceToEarnings string `json:"ttm_price_to_earnings" api:"nullable"`
	// The MIC code of the primary listing venue
	Venue string `json:"venue" api:"nullable"`
	// The latest trading volume for the instrument
	Volume string `json:"volume" api:"nullable"`
	// The average trading volume over the past week
	WeekAvgVolume string `json:"week_avg_volume" api:"nullable"`
	// The opening price on the first trading day of the current year
	YearToDateOpen string `json:"year_to_date_open" api:"nullable"`
	// Percent change from year-to-date open to previous day close
	YtdChangePct string `json:"ytd_change_pct" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BuyRatings               respjson.Field
		HoldRatings              respjson.Field
		Price                    respjson.Field
		SecurityID               respjson.Field
		SecurityIDSource         respjson.Field
		SellRatings              respjson.Field
		StrongBuyRatings         respjson.Field
		StrongSellRatings        respjson.Field
		Symbol                   respjson.Field
		TotalRatings             respjson.Field
		ConsensusPriceTarget     respjson.Field
		ConsensusPriceTargetHigh respjson.Field
		ConsensusPriceTargetLow  respjson.Field
		ConsensusRating          respjson.Field
		CountryOfIssue           respjson.Field
		Description              respjson.Field
		FiftyTwoWeekHigh         respjson.Field
		FiftyTwoWeekLow          respjson.Field
		GapFrom52wHighPct        respjson.Field
		GapFrom52wLowPct         respjson.Field
		Industry                 respjson.Field
		ListDate                 respjson.Field
		MarketCap                respjson.Field
		MonthAvgVolume           respjson.Field
		Name                     respjson.Field
		OneMonthAgoClose         respjson.Field
		OneMonthAgoOpen          respjson.Field
		OneMonthChangePct        respjson.Field
		OneWeekAgoClose          respjson.Field
		OneWeekAgoOpen           respjson.Field
		OneWeekChangePct         respjson.Field
		OneYearAgoClose          respjson.Field
		OneYearAgoOpen           respjson.Field
		OneYearChangePct         respjson.Field
		PercentChange            respjson.Field
		PrevDayClose             respjson.Field
		Sector                   respjson.Field
		SecurityType             respjson.Field
		SixMonthChangePct        respjson.Field
		SixMonthsAgoClose        respjson.Field
		SixMonthsAgoOpen         respjson.Field
		ThreeMonthChangePct      respjson.Field
		ThreeMonthsAgoClose      respjson.Field
		ThreeMonthsAgoOpen       respjson.Field
		TtmDebtToEquity          respjson.Field
		TtmDividendYield         respjson.Field
		TtmEarningsPerShare      respjson.Field
		TtmPriceToEarnings       respjson.Field
		Venue                    respjson.Field
		Volume                   respjson.Field
		WeekAvgVolume            respjson.Field
		YearToDateOpen           respjson.Field
		YtdChangePct             respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScreenerItem) RawJSON() string { return r.JSON.raw }
func (r *ScreenerItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScreenerItemList []ScreenerItem

type ActiveV1ScreenerGetScreenerResponse struct {
	Data ScreenerItemList `json:"data" api:"required"`
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
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
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
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort direction (ASC or DESC, defaults to DESC)
type ActiveV1ScreenerGetScreenerParamsSortDirection string

const (
	ActiveV1ScreenerGetScreenerParamsSortDirectionAsc  ActiveV1ScreenerGetScreenerParamsSortDirection = "ASC"
	ActiveV1ScreenerGetScreenerParamsSortDirectionDesc ActiveV1ScreenerGetScreenerParamsSortDirection = "DESC"
)
