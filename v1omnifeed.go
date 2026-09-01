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

// Personalized feed of market stories: upcoming earnings, dividends, and splits,
// plus market news. Served per caller in a stable order; item ids double as
// pagination cursors, so any previously returned page can be re-read.
//
// V1OmniFeedService contains methods and other services that help with interacting
// with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1OmniFeedService] method instead.
type V1OmniFeedService struct {
	options []option.RequestOption
}

// NewV1OmniFeedService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1OmniFeedService(opts ...option.RequestOption) (r V1OmniFeedService) {
	r = V1OmniFeedService{}
	r.options = opts
	return
}

// > **Alpha** — this endpoint is experimental and may change or be removed at any
// > time.
//
// Returns the caller's personalized feed of market stories: upcoming earnings,
// dividends, and splits, plus market news.
//
// The feed is a stable, append-only sequence per caller. Without a `cursor`, the
// response resumes from the caller's oldest unseen item and continues forward;
// passing the id of the last item received as `cursor` returns the items after it.
// Either way, when the known sequence runs short of `limit`, fresh stories are
// appended and included. Re-requesting an earlier cursor replays the same items in
// the same order.
func (r *V1OmniFeedService) GetFeed(ctx context.Context, query V1OmniFeedGetFeedParams, opts ...option.RequestOption) (res *V1OmniFeedGetFeedResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/omni-ai/feed"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// > **Alpha** — this endpoint is experimental and may change or be removed at any
// > time.
//
// Records one of the caller's interactions with a feed item: the item rendered on
// screen (`seen`), expanded from its headline to its summary (`click`), or voted
// on (`upvote`, `downvote`). Marking an item `seen` excludes it from the caller's
// next feed response.
//
// Each request records a new event, so a request that is retried after a failure
// of unknown outcome may be recorded twice. That is harmless for `seen` and for
// votes — the first is a yes-or-no exclusion and the latest vote is the one that
// counts — so retry freely for those. The event is wrapped in an `event` field so
// that a future request may carry several at once without breaking this one.
func (r *V1OmniFeedService) PostFeedEvent(ctx context.Context, body V1OmniFeedPostFeedEventParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/omni-ai/feed/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// One item in the caller's feed.
type FeedItem struct {
	// Unique item id. Also the pagination cursor: pass it as `cursor` to fetch the
	// items that follow it.
	ID string `json:"id" api:"required" format:"uuid"`
	// Headline text.
	Headline string `json:"headline" api:"required"`
	// What the item is about.
	//
	// Any of "earnings", "dividend", "split", "news", "omni".
	Kind FeedItemKind `json:"kind" api:"required"`
	// When the item's content was published.
	PublishedAt time.Time `json:"published_at" api:"required" format:"date-time"`
	// Summary text.
	Summary string `json:"summary" api:"required"`
	// The item's headline number. When a null/undefined value is observed, it
	// indicates that there is no available data.
	Metric FeedItemMetric `json:"metric" api:"nullable"`
	// When the underlying event is expected to occur, for items about an upcoming
	// event. When a null/undefined value is observed, it indicates it does not apply.
	OccursAt time.Time `json:"occurs_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Headline    respjson.Field
		Kind        respjson.Field
		PublishedAt respjson.Field
		Summary     respjson.Field
		Metric      respjson.Field
		OccursAt    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeedItem) RawJSON() string { return r.JSON.raw }
func (r *FeedItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What a feed item is about; fixes the meaning of the item's `metric`.
type FeedItemKind string

const (
	FeedItemKindEarnings FeedItemKind = "earnings"
	FeedItemKindDividend FeedItemKind = "dividend"
	FeedItemKindSplit    FeedItemKind = "split"
	FeedItemKindNews     FeedItemKind = "news"
	FeedItemKindOmni     FeedItemKind = "omni"
)

// A feed item's headline number.
type FeedItemMetric struct {
	// What the number measures.
	//
	// Any of "eps_estimate", "dividend_amount", "split_ratio".
	Type FeedMetricType `json:"type" api:"required"`
	// The number itself, as a decimal string.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeedItemMetric) RawJSON() string { return r.JSON.raw }
func (r *FeedItemMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What a feed item's number measures.
type FeedMetricType string

const (
	FeedMetricTypeEpsEstimate    FeedMetricType = "eps_estimate"
	FeedMetricTypeDividendAmount FeedMetricType = "dividend_amount"
	FeedMetricTypeSplitRatio     FeedMetricType = "split_ratio"
)

// One page of the caller's feed.
type FeedPage struct {
	// Feed items, in feed order.
	Items []FeedItem `json:"items" api:"required"`
	// Cursor for the page after this one: the last item's id. Absent only when there
	// are no items to serve. When a null/undefined value is observed, it indicates
	// that there is no available data.
	NextCursor string `json:"next_cursor" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeedPage) RawJSON() string { return r.JSON.raw }
func (r *FeedPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniFeedGetFeedResponse struct {
	// One page of the caller's feed.
	Data FeedPage `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1OmniFeedGetFeedResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniFeedGetFeedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniFeedGetFeedParams struct {
	// Trading account to serve as context. Optional — the feed works without one.
	AccountID param.Opt[int64] `query:"account_id,omitzero" json:"-"`
	// Id of the last item already received; the response continues from the item after
	// it. Omit to resume from the oldest unseen item.
	Cursor param.Opt[string] `query:"cursor,omitzero" format:"uuid" json:"-"`
	// Maximum number of items to return (1–100, default 20).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1OmniFeedGetFeedParams]'s query parameters as
// `url.Values`.
func (r V1OmniFeedGetFeedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1OmniFeedPostFeedEventParams struct {
	// The event to record.
	Event V1OmniFeedPostFeedEventParamsEvent `json:"event,omitzero" api:"required"`
	paramObj
}

func (r V1OmniFeedPostFeedEventParams) MarshalJSON() (data []byte, err error) {
	type shadow V1OmniFeedPostFeedEventParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1OmniFeedPostFeedEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event to record.
//
// The properties ItemID, Type are required.
type V1OmniFeedPostFeedEventParamsEvent struct {
	// The feed item the event concerns.
	ItemID string `json:"item_id" api:"required" format:"uuid"`
	// What happened.
	//
	// Any of "seen", "click", "upvote", "downvote".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r V1OmniFeedPostFeedEventParamsEvent) MarshalJSON() (data []byte, err error) {
	type shadow V1OmniFeedPostFeedEventParamsEvent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1OmniFeedPostFeedEventParamsEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1OmniFeedPostFeedEventParamsEvent](
		"type", "seen", "click", "upvote", "downvote",
	)
}
