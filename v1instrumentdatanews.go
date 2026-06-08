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

// Retrieve instrument analytics, market data, news, and related reference data.
//
// V1InstrumentDataNewsService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1InstrumentDataNewsService] method instead.
type V1InstrumentDataNewsService struct {
	options []option.RequestOption
}

// NewV1InstrumentDataNewsService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1InstrumentDataNewsService(opts ...option.RequestOption) (r V1InstrumentDataNewsService) {
	r = V1InstrumentDataNewsService{}
	r.options = opts
	return
}

// Retrieves news items with optional filtering by security IDs, time range,
// publisher, type, and text query.
func (r *V1InstrumentDataNewsService) GetNews(ctx context.Context, query V1InstrumentDataNewsGetNewsParams, opts ...option.RequestOption) (res *V1InstrumentDataNewsGetNewsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/news"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Instrument associated with a news item.
type NewsInstrument struct {
	// Instrument identifier.
	InstrumentID string `json:"instrument_id" api:"required" format:"uuid"`
	// Instrument name/description, if available.
	Name string `json:"name" api:"nullable"`
	// Trading symbol, if available.
	Symbol string `json:"symbol" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InstrumentID respjson.Field
		Name         respjson.Field
		Symbol       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsInstrument) RawJSON() string { return r.JSON.raw }
func (r *NewsInstrument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single news item and its associated instruments.
type NewsItem struct {
	// Instruments associated with this news item.
	Instruments []NewsInstrument `json:"instruments" api:"required"`
	// Classification of the item.
	//
	// Any of "NEWS", "PRESS_RELEASE".
	NewsType NewsType `json:"news_type" api:"required"`
	// The published date/time of the article in UTC.
	PublishedAt time.Time `json:"published_at" api:"required" format:"date-time"`
	// The publisher or newswire source.
	Publisher string `json:"publisher" api:"required"`
	// The headline/title of the article.
	Title string `json:"title" api:"required"`
	// Canonical URL to the full article.
	URL string `json:"url" api:"required"`
	// URL of an associated image if provided by the source.
	ImageURL string `json:"image_url" api:"nullable"`
	// The primary domain/site of the publisher.
	Site string `json:"site" api:"nullable"`
	// The full or excerpted article body.
	Text string `json:"text" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instruments respjson.Field
		NewsType    respjson.Field
		PublishedAt respjson.Field
		Publisher   respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ImageURL    respjson.Field
		Site        respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsItem) RawJSON() string { return r.JSON.raw }
func (r *NewsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsItemList []NewsItem

// News item classification.
type NewsType string

const (
	NewsTypeNews         NewsType = "NEWS"
	NewsTypePressRelease NewsType = "PRESS_RELEASE"
)

type V1InstrumentDataNewsGetNewsResponse struct {
	Data NewsItemList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentDataNewsGetNewsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentDataNewsGetNewsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentDataNewsGetNewsParams struct {
	// Comma-separated list of publishers to exclude (mutually exclusive with
	// include_publishers).
	ExcludePublishers param.Opt[string] `query:"exclude_publishers,omitzero" json:"-"`
	// Inclusive start timestamp. Accepts `YYYY-MM-DD` or RFC3339 datetime.
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// Comma-separated list of publishers to include (mutually exclusive with
	// exclude_publishers).
	IncludePublishers param.Opt[string] `query:"include_publishers,omitzero" json:"-"`
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Free-text query matched against title/text and associated security IDs.
	SearchQuery param.Opt[string] `query:"search_query,omitzero" json:"-"`
	// Inclusive end timestamp. Accepts `YYYY-MM-DD` or RFC3339 datetime.
	To param.Opt[string] `query:"to,omitzero" json:"-"`
	// Comma-delimited instrument identifiers to filter by.
	InstrumentIDs []string `query:"instrument_ids,omitzero" json:"-"`
	// Filter by news type.
	//
	// Any of "NEWS", "PRESS_RELEASE".
	NewsType V1InstrumentDataNewsGetNewsParamsNewsType `query:"news_type,omitzero" json:"-"`
	// Comma-separated sector values to filter by.
	//
	// Any of "BASIC_MATERIALS", "COMMUNICATION_SERVICES", "CONSUMER_CYCLICAL",
	// "CONSUMER_DEFENSIVE", "ENERGY", "FINANCIAL_SERVICES", "HEALTHCARE",
	// "INDUSTRIALS", "REAL_ESTATE", "TECHNOLOGY", "UTILITIES".
	Sectors []string `query:"sectors,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentDataNewsGetNewsParams]'s query parameters as
// `url.Values`.
func (r V1InstrumentDataNewsGetNewsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by news type.
type V1InstrumentDataNewsGetNewsParamsNewsType string

const (
	V1InstrumentDataNewsGetNewsParamsNewsTypeNews         V1InstrumentDataNewsGetNewsParamsNewsType = "NEWS"
	V1InstrumentDataNewsGetNewsParamsNewsTypePressRelease V1InstrumentDataNewsGetNewsParamsNewsType = "PRESS_RELEASE"
)
