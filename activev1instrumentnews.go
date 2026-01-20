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

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/apiquery"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// ActiveV1InstrumentNewsService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1InstrumentNewsService] method instead.
type ActiveV1InstrumentNewsService struct {
	Options []option.RequestOption
}

// NewActiveV1InstrumentNewsService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewActiveV1InstrumentNewsService(opts ...option.RequestOption) (r ActiveV1InstrumentNewsService) {
	r = ActiveV1InstrumentNewsService{}
	r.Options = opts
	return
}

// Retrieves recent news articles related to an instrument.
func (r *ActiveV1InstrumentNewsService) GetInstrumentNews(ctx context.Context, securityID string, params ActiveV1InstrumentNewsGetInstrumentNewsParams, opts ...option.RequestOption) (res *ActiveV1InstrumentNewsGetInstrumentNewsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if securityID == "" {
		err = errors.New("missing required security_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/instruments/%v/%s/news", params.SecurityIDSource, securityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// A news or press release item related to an instrument
type InstrumentNews struct {
	// The published date/time of the article in UTC
	PublishedAt time.Time `json:"published_at,required" format:"date-time"`
	// The trading symbol associated with the news item
	Symbol string `json:"symbol,required"`
	// The headline/title of the article
	Title string `json:"title,required"`
	// Classification of the item
	//
	// Any of "NEWS", "PRESS_RELEASE".
	Type InstrumentNewsType `json:"type,required"`
	// Canonical URL to the full article
	URL string `json:"url,required"`
	// URL of an associated image if provided by the source
	ImageURL string `json:"image_url,nullable"`
	// The publisher or newswire source
	Publisher string `json:"publisher,nullable"`
	// The primary domain/site of the publisher
	Site string `json:"site,nullable"`
	// The full or excerpted article body
	Text string `json:"text,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PublishedAt respjson.Field
		Symbol      respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ImageURL    respjson.Field
		Publisher   respjson.Field
		Site        respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentNews) RawJSON() string { return r.JSON.raw }
func (r *InstrumentNews) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Classification of the item
type InstrumentNewsType string

const (
	InstrumentNewsTypeNews         InstrumentNewsType = "NEWS"
	InstrumentNewsTypePressRelease InstrumentNewsType = "PRESS_RELEASE"
)

type InstrumentNewsList []InstrumentNews

type ActiveV1InstrumentNewsGetInstrumentNewsResponse struct {
	Data InstrumentNewsList `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1InstrumentNewsGetInstrumentNewsResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1InstrumentNewsGetInstrumentNewsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1InstrumentNewsGetInstrumentNewsParams struct {
	// Security identifier source
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "OTHER".
	SecurityIDSource SecurityIDSource `path:"security_id_source,omitzero,required" json:"-"`
	// The start date for the query range, inclusive (YYYY-MM-DD)
	FromDate string `query:"from_date,required" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD)
	ToDate string `query:"to_date,required" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1InstrumentNewsGetInstrumentNewsParams]'s query
// parameters as `url.Values`.
func (r ActiveV1InstrumentNewsGetInstrumentNewsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
