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

// Retrieve details and lists of tradable instruments.
//
// ActiveV1NewsService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1NewsService] method instead.
type ActiveV1NewsService struct {
	Options []option.RequestOption
}

// NewActiveV1NewsService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActiveV1NewsService(opts ...option.RequestOption) (r ActiveV1NewsService) {
	r = ActiveV1NewsService{}
	r.Options = opts
	return
}

// Retrieves news items with optional filtering by security IDs, time range,
// publisher, type, and text query.
func (r *ActiveV1NewsService) GetNews(ctx context.Context, query ActiveV1NewsGetNewsParams, opts ...option.RequestOption) (res *ActiveV1NewsGetNewsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "active/v1/news"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A single news item and its associated instruments.
type NewsItem struct {
	// Instruments associated with this news item.
	Instruments []NewsItemInstrument `json:"instruments" api:"required"`
	// Classification of the item.
	//
	// Any of "NEWS", "PRESS_RELEASE".
	NewsType NewsItemNewsType `json:"news_type" api:"required"`
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

// Instrument associated with a news item.
type NewsItemInstrument struct {
	// Security identifier value.
	SecurityID string `json:"security_id" api:"required"`
	// Security identifier source.
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
	// OEMS instrument UUID, if available from instrument cache enrichment.
	InstrumentID string `json:"instrument_id" api:"nullable" format:"uuid"`
	// Instrument name/description, if available from instrument cache enrichment.
	Name string `json:"name" api:"nullable"`
	// Trading symbol, if available from instrument cache enrichment.
	Symbol string `json:"symbol" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SecurityID       respjson.Field
		SecurityIDSource respjson.Field
		InstrumentID     respjson.Field
		Name             respjson.Field
		Symbol           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsItemInstrument) RawJSON() string { return r.JSON.raw }
func (r *NewsItemInstrument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Classification of the item.
type NewsItemNewsType string

const (
	NewsItemNewsTypeNews         NewsItemNewsType = "NEWS"
	NewsItemNewsTypePressRelease NewsItemNewsType = "PRESS_RELEASE"
)

type NewsItemList []NewsItem

type ActiveV1NewsGetNewsResponse struct {
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
func (r ActiveV1NewsGetNewsResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1NewsGetNewsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1NewsGetNewsParams struct {
	// Comma-separated list of publishers to exclude (mutually exclusive with
	// include_publishers).
	ExcludePublishers param.Opt[string] `query:"exclude_publishers,omitzero" json:"-"`
	// Inclusive start timestamp. Accepts `YYYY-MM-DD` or RFC3339 datetime.
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// Comma-separated list of publishers to include (mutually exclusive with
	// exclude_publishers).
	IncludePublishers param.Opt[string] `query:"include_publishers,omitzero" json:"-"`
	// The number of items to return per page (only used when page_token is not
	// provided)
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Free-text query matched against title/text and associated security IDs.
	SearchQuery param.Opt[string] `query:"search_query,omitzero" json:"-"`
	// Inclusive end timestamp. Accepts `YYYY-MM-DD` or RFC3339 datetime.
	To param.Opt[string] `query:"to,omitzero" json:"-"`
	// Comma-delimited OEMS instrument UUIDs to filter by.
	InstrumentIDs []string `query:"instrument_ids,omitzero" json:"-"`
	// Filter by news type.
	//
	// Any of "NEWS", "PRESS_RELEASE".
	NewsType ActiveV1NewsGetNewsParamsNewsType `query:"news_type,omitzero" json:"-"`
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
	paramObj
}

// URLQuery serializes [ActiveV1NewsGetNewsParams]'s query parameters as
// `url.Values`.
func (r ActiveV1NewsGetNewsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by news type.
type ActiveV1NewsGetNewsParamsNewsType string

const (
	ActiveV1NewsGetNewsParamsNewsTypeNews         ActiveV1NewsGetNewsParamsNewsType = "NEWS"
	ActiveV1NewsGetNewsParamsNewsTypePressRelease ActiveV1NewsGetNewsParamsNewsType = "PRESS_RELEASE"
)
