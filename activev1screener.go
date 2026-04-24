// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"encoding/json"
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

// Search and manage saved screeners.
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

// Screen instruments.
//
// Searches for instruments matching specified criteria.
func (r *ActiveV1ScreenerService) GetScreener(ctx context.Context, query ActiveV1ScreenerGetScreenerParams, opts ...option.RequestOption) (res *ActiveV1ScreenerGetScreenerResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/screener"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search instruments using structured filters.
//
// Returns a columnar response where each row is an array of column objects. Each
// column contains a human-readable name, a field reference, an optional type hint
// (e.g. `CURR_USD`, `PERCENT`), and the value.
//
// Use `field_filter` to select which columns appear in each row. When omitted, the
// default field set is returned.
func (r *ActiveV1ScreenerService) SearchScreener(ctx context.Context, body ActiveV1ScreenerSearchScreenerParams, opts ...option.RequestOption) (res *ActiveV1ScreenerSearchScreenerResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/screener"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Historical lookback window for price/change fields.
type FieldLookback string

const (
	FieldLookbackOneWeek     FieldLookback = "ONE_WEEK"
	FieldLookbackOneMonth    FieldLookback = "ONE_MONTH"
	FieldLookbackThreeMonths FieldLookback = "THREE_MONTHS"
	FieldLookbackSixMonths   FieldLookback = "SIX_MONTHS"
	FieldLookbackYtd         FieldLookback = "YTD"
	FieldLookbackOneYear     FieldLookback = "ONE_YEAR"
)

// Reporting period for financial data fields.
type FieldPeriod string

const (
	FieldPeriodQuarter FieldPeriod = "QUARTER"
	FieldPeriodTtm     FieldPeriod = "TTM"
)

// A reference to a screener field.
type FieldRef struct {
	// The field name.
	Name string `json:"name" api:"required"`
	// Optional historical lookback window.
	//
	// Any of "ONE_WEEK", "ONE_MONTH", "THREE_MONTHS", "SIX_MONTHS", "YTD", "ONE_YEAR".
	Lookback FieldLookback `json:"lookback" api:"nullable"`
	// Optional reporting period (e.g. quarter or TTM).
	//
	// Any of "QUARTER", "TTM".
	Period FieldPeriod `json:"period" api:"nullable"`
	// The data type of the field value. Present only in responses.
	//
	// Any of "DECIMAL", "INTEGER", "STRING", "ANALYST_RATING", "DATE".
	ValueType FieldType `json:"value_type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Lookback    respjson.Field
		Period      respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FieldRef) RawJSON() string { return r.JSON.raw }
func (r *FieldRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FieldRef to a FieldRefParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FieldRefParam.Overrides()
func (r FieldRef) ToParam() FieldRefParam {
	return param.Override[FieldRefParam](json.RawMessage(r.RawJSON()))
}

// A reference to a screener field.
//
// The property Name is required.
type FieldRefParam struct {
	// The field name.
	Name string `json:"name" api:"required"`
	// Optional historical lookback window.
	//
	// Any of "ONE_WEEK", "ONE_MONTH", "THREE_MONTHS", "SIX_MONTHS", "YTD", "ONE_YEAR".
	Lookback FieldLookback `json:"lookback,omitzero"`
	// Optional reporting period (e.g. quarter or TTM).
	//
	// Any of "QUARTER", "TTM".
	Period FieldPeriod `json:"period,omitzero"`
	// The data type of the field value. Present only in responses.
	//
	// Any of "DECIMAL", "INTEGER", "STRING", "ANALYST_RATING", "DATE".
	ValueType FieldType `json:"value_type,omitzero"`
	paramObj
}

func (r FieldRefParam) MarshalJSON() (data []byte, err error) {
	type shadow FieldRefParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FieldRefParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The data type of a screener field value.
type FieldType string

const (
	FieldTypeDecimal       FieldType = "DECIMAL"
	FieldTypeInteger       FieldType = "INTEGER"
	FieldTypeString        FieldType = "STRING"
	FieldTypeAnalystRating FieldType = "ANALYST_RATING"
	FieldTypeDate          FieldType = "DATE"
)

// A single column in the screener search response.
type ScreenerColumn struct {
	// Field reference (same shape as filter/sort field references)
	Field FieldRef `json:"field" api:"required"`
	// Human-readable display name for this field
	Name  string                   `json:"name" api:"required"`
	Value ScreenerColumnValueUnion `json:"value" api:"required"`
	// Value format hint: "CURR_USD", "PERCENT", etc. Omitted when not applicable.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Name        respjson.Field
		Value       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScreenerColumn) RawJSON() string { return r.JSON.raw }
func (r *ScreenerColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScreenerColumnValueUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type ScreenerColumnValueUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ScreenerColumnValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScreenerColumnValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScreenerColumnValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ScreenerColumnValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	// The trading symbol for the instrument
	Symbol string `json:"symbol" api:"required"`
	// The total count of analyst ratings
	TotalRatings int64 `json:"total_ratings" api:"required"`
	// The consensus analyst price target
	ConsensusPriceTarget string `json:"consensus_price_target" api:"nullable"`
	// The consensus analyst rating
	//
	// Any of "STRONG_BUY", "BUY", "HOLD", "SELL", "STRONG_SELL".
	ConsensusRating AnalystRating `json:"consensus_rating" api:"nullable"`
	// The ISO country code of the instrument's issue
	CountryOfIssue string `json:"country_of_issue" api:"nullable"`
	// The TTM debt-to-equity ratio
	DebtToEquityTtm string `json:"debt_to_equity_ttm" api:"nullable"`
	// A detailed description of the instrument or company
	Description string `json:"description" api:"nullable"`
	// The TTM dividend yield percent
	DividendYieldTtm string `json:"dividend_yield_ttm" api:"nullable"`
	// The TTM earnings per share
	EarningsPerShareTtm string `json:"earnings_per_share_ttm" api:"nullable"`
	// The MIC code of the primary listing exchange
	Exchange string `json:"exchange" api:"nullable"`
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
	// The TTM price-to-earnings ratio
	PriceToEarningsTtm string `json:"price_to_earnings_ttm" api:"nullable"`
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
		Price                respjson.Field
		SecurityID           respjson.Field
		SecurityIDSource     respjson.Field
		Symbol               respjson.Field
		TotalRatings         respjson.Field
		ConsensusPriceTarget respjson.Field
		ConsensusRating      respjson.Field
		CountryOfIssue       respjson.Field
		DebtToEquityTtm      respjson.Field
		Description          respjson.Field
		DividendYieldTtm     respjson.Field
		EarningsPerShareTtm  respjson.Field
		Exchange             respjson.Field
		FiftyTwoWeekHigh     respjson.Field
		FiftyTwoWeekLow      respjson.Field
		GapFrom52wHighPct    respjson.Field
		GapFrom52wLowPct     respjson.Field
		Industry             respjson.Field
		ListDate             respjson.Field
		MarketCap            respjson.Field
		MonthAvgVolume       respjson.Field
		Name                 respjson.Field
		OneMonthAgoClose     respjson.Field
		OneMonthAgoOpen      respjson.Field
		OneMonthChangePct    respjson.Field
		OneWeekAgoClose      respjson.Field
		OneWeekAgoOpen       respjson.Field
		OneWeekChangePct     respjson.Field
		OneYearAgoClose      respjson.Field
		OneYearAgoOpen       respjson.Field
		OneYearChangePct     respjson.Field
		PercentChange        respjson.Field
		PrevDayClose         respjson.Field
		PriceToEarningsTtm   respjson.Field
		Sector               respjson.Field
		SecurityType         respjson.Field
		SixMonthChangePct    respjson.Field
		SixMonthsAgoClose    respjson.Field
		SixMonthsAgoOpen     respjson.Field
		ThreeMonthChangePct  respjson.Field
		ThreeMonthsAgoClose  respjson.Field
		ThreeMonthsAgoOpen   respjson.Field
		Volume               respjson.Field
		WeekAvgVolume        respjson.Field
		YearToDateOpen       respjson.Field
		YtdChangePct         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScreenerItem) RawJSON() string { return r.JSON.raw }
func (r *ScreenerItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScreenerItemList []ScreenerItem

type ScreenerRow []ScreenerColumn

type ScreenerRowList []ScreenerRow

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

type ActiveV1ScreenerSearchScreenerResponse struct {
	Data ScreenerRowList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1ScreenerSearchScreenerResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1ScreenerSearchScreenerResponse) UnmarshalJSON(data []byte) error {
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

type ActiveV1ScreenerSearchScreenerParams struct {
	// Maximum number of results per page.
	PageSize param.Opt[int64] `json:"page_size,omitzero"`
	// Opaque token for cursor-based pagination.
	PageToken param.Opt[string] `json:"page_token,omitzero"`
	// Whether string sorts should be case-sensitive (default: false).
	SortCaseSensitive param.Opt[bool] `json:"sort_case_sensitive,omitzero"`
	// Subset of fields to include in the response.
	FieldFilter []FieldRefParam `json:"field_filter,omitzero"`
	// Filter conditions to apply.
	Filters []ActiveV1ScreenerSearchScreenerParamsFilter `json:"filters,omitzero"`
	// Multi-field sort specifications. When present, takes precedence over
	// sort_by/sort_direction.
	Sorts []ActiveV1ScreenerSearchScreenerParamsSort `json:"sorts,omitzero"`
	// Field to sort results by.
	SortBy FieldRefParam `json:"sort_by,omitzero"`
	// Sort direction (defaults to DESC).
	//
	// Any of "ASC", "DESC".
	SortDirection ActiveV1ScreenerSearchScreenerParamsSortDirection `json:"sort_direction,omitzero"`
	paramObj
}

func (r ActiveV1ScreenerSearchScreenerParams) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1ScreenerSearchScreenerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1ScreenerSearchScreenerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single filter condition.
//
// The properties Left, Op, Right are required.
type ActiveV1ScreenerSearchScreenerParamsFilter struct {
	// The field to filter on.
	Left FieldRefParam `json:"left,omitzero" api:"required"`
	// The operator and optional arguments.
	Op ActiveV1ScreenerSearchScreenerParamsFilterOp `json:"op,omitzero" api:"required"`
	// The value(s) to compare against.
	Right []ActiveV1ScreenerSearchScreenerParamsFilterRight `json:"right,omitzero" api:"required"`
	paramObj
}

func (r ActiveV1ScreenerSearchScreenerParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1ScreenerSearchScreenerParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1ScreenerSearchScreenerParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The operator and optional arguments.
//
// The property Name is required.
type ActiveV1ScreenerSearchScreenerParamsFilterOp struct {
	// The operator to apply.
	//
	// Any of "LT", "LTE", "GT", "GTE", "EQ", "BETWEEN", "NOT_BETWEEN", "ONE_OF",
	// "REGEX", "BEGINS_WITH", "ENDS_WITH", "CONTAINS", "IS_NULL", "IS_NOT_NULL".
	Name string `json:"name,omitzero" api:"required"`
	// Optional arguments that modify operator behavior.
	//
	// Any of "LEFT_INCLUSIVE", "RIGHT_INCLUSIVE", "LEFT_EXCLUSIVE", "RIGHT_EXCLUSIVE",
	// "CASE_INSENSITIVE".
	Args []string `json:"args,omitzero"`
	paramObj
}

func (r ActiveV1ScreenerSearchScreenerParamsFilterOp) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1ScreenerSearchScreenerParamsFilterOp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1ScreenerSearchScreenerParamsFilterOp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActiveV1ScreenerSearchScreenerParamsFilterOp](
		"name", "LT", "LTE", "GT", "GTE", "EQ", "BETWEEN", "NOT_BETWEEN", "ONE_OF", "REGEX", "BEGINS_WITH", "ENDS_WITH", "CONTAINS", "IS_NULL", "IS_NOT_NULL",
	)
}

// A filter value: either a literal or a variable reference.
type ActiveV1ScreenerSearchScreenerParamsFilterRight struct {
	Value ActiveV1ScreenerSearchScreenerParamsFilterRightValueUnion `json:"value,omitzero"`
	// A variable reference.
	Variable ActiveV1ScreenerSearchScreenerParamsFilterRightVariable `json:"variable,omitzero"`
	paramObj
}

func (r ActiveV1ScreenerSearchScreenerParamsFilterRight) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1ScreenerSearchScreenerParamsFilterRight
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1ScreenerSearchScreenerParamsFilterRight) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ActiveV1ScreenerSearchScreenerParamsFilterRightValueUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ActiveV1ScreenerSearchScreenerParamsFilterRightValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ActiveV1ScreenerSearchScreenerParamsFilterRightValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// A variable reference.
//
// The property Name is required.
type ActiveV1ScreenerSearchScreenerParamsFilterRightVariable struct {
	// The variable name.
	Name string `json:"name" api:"required"`
	// Optional arithmetic modifier.
	Modifier ActiveV1ScreenerSearchScreenerParamsFilterRightVariableModifier `json:"modifier,omitzero"`
	// Optional historical lookback window.
	//
	// Any of "ONE_WEEK", "ONE_MONTH", "THREE_MONTHS", "SIX_MONTHS", "YTD", "ONE_YEAR".
	Lookback FieldLookback `json:"lookback,omitzero"`
	// Optional reporting period.
	//
	// Any of "QUARTER", "TTM".
	Period FieldPeriod `json:"period,omitzero"`
	paramObj
}

func (r ActiveV1ScreenerSearchScreenerParamsFilterRightVariable) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1ScreenerSearchScreenerParamsFilterRightVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1ScreenerSearchScreenerParamsFilterRightVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional arithmetic modifier.
//
// The properties Args, Name are required.
type ActiveV1ScreenerSearchScreenerParamsFilterRightVariableModifier struct {
	Args []ActiveV1ScreenerSearchScreenerParamsFilterRightVariableModifierArgUnion `json:"args,omitzero" api:"required"`
	// The modifier operation.
	//
	// Any of "ADD", "SUB".
	Name string `json:"name,omitzero" api:"required"`
	paramObj
}

func (r ActiveV1ScreenerSearchScreenerParamsFilterRightVariableModifier) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1ScreenerSearchScreenerParamsFilterRightVariableModifier
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1ScreenerSearchScreenerParamsFilterRightVariableModifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActiveV1ScreenerSearchScreenerParamsFilterRightVariableModifier](
		"name", "ADD", "SUB",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ActiveV1ScreenerSearchScreenerParamsFilterRightVariableModifierArgUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ActiveV1ScreenerSearchScreenerParamsFilterRightVariableModifierArgUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ActiveV1ScreenerSearchScreenerParamsFilterRightVariableModifierArgUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Sort direction (defaults to DESC).
type ActiveV1ScreenerSearchScreenerParamsSortDirection string

const (
	ActiveV1ScreenerSearchScreenerParamsSortDirectionAsc  ActiveV1ScreenerSearchScreenerParamsSortDirection = "ASC"
	ActiveV1ScreenerSearchScreenerParamsSortDirectionDesc ActiveV1ScreenerSearchScreenerParamsSortDirection = "DESC"
)

// A sort specification pairing a field with a direction.
//
// The property Field is required.
type ActiveV1ScreenerSearchScreenerParamsSort struct {
	// The field to sort by.
	Field FieldRefParam `json:"field,omitzero" api:"required"`
	// Sort direction (defaults to DESC).
	//
	// Any of "ASC", "DESC".
	Direction string `json:"direction,omitzero"`
	paramObj
}

func (r ActiveV1ScreenerSearchScreenerParamsSort) MarshalJSON() (data []byte, err error) {
	type shadow ActiveV1ScreenerSearchScreenerParamsSort
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActiveV1ScreenerSearchScreenerParamsSort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActiveV1ScreenerSearchScreenerParamsSort](
		"direction", "ASC", "DESC",
	)
}
