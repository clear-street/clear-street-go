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

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/internal/apiquery"
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/param"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Retrieve core details and discovery endpoints for tradable instruments.
//
// V1InstrumentService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1InstrumentService] method instead.
type V1InstrumentService struct {
	options []option.RequestOption
}

// NewV1InstrumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1InstrumentService(opts ...option.RequestOption) (r V1InstrumentService) {
	r = V1InstrumentService{}
	r.options = opts
	return
}

// Retrieves detailed information for a specific instrument.
func (r *V1InstrumentService) GetInstrumentByID(ctx context.Context, instrumentID string, query V1InstrumentGetInstrumentByIDParams, opts ...option.RequestOption) (res *V1InstrumentGetInstrumentByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/instruments/%s", url.PathEscape(instrumentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves a list of tradeable instruments.
func (r *V1InstrumentService) GetInstruments(ctx context.Context, query V1InstrumentGetInstrumentsParams, opts ...option.RequestOption) (res *V1InstrumentGetInstrumentsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/instruments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List options contracts.
//
// Returns options contracts with options-specific metadata. Exactly one identifier
// must be provided: `root`/`underlying_instrument_id` (list all contracts for that
// underlier) or `contract_ids` (look up specific contracts directly).
// `expiry`/`contract_type` apply as filters in either case.
func (r *V1InstrumentService) GetOptionContracts(ctx context.Context, query V1InstrumentGetOptionContractsParams, opts ...option.RequestOption) (res *V1InstrumentGetOptionContractsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/instruments/options/contracts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// > **Alpha** — this endpoint is experimental and may change or be removed at any
// > time.
//
// Lists the distinct categories, sorted alphabetically, of every prediction market
// contract that can currently be listed: contracts that are initialized or active
// and have not yet closed.
//
// The list refreshes every 15 minutes. Returns `503` until the categories first
// load after a restart.
func (r *V1InstrumentService) GetPredictionCategories(ctx context.Context, opts ...option.RequestOption) (res *V1InstrumentGetPredictionCategoriesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/instruments/predictions/categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// > **Alpha** — this endpoint is experimental and may change or be removed at any
// > time.
//
// Lists the prediction market events that have at least one market that can
// currently be listed: initialized or active and not yet closed. A market that is
// paused does not count, so an event whose markets are all paused is not listed.
//
// Events are ordered by code unless `sort_by` names another key or `search` is
// set; ties break on code ascending. `series_codes` and `categories` keep events
// of those series or with a listable market in those categories, ignoring case.
//
// `search` keeps events whose code, series code, or a market's question or title
// contains every word of the text. Words match by prefix and English stem,
// ignoring case and punctuation; text with no searchable word, such as `the`,
// matches nothing. A word of four or more letters that matches nothing is replaced
// by up to three close spellings. Without `sort_by`, the best match comes first,
// closest spellings first. Search covers listable events unless
// `include_settled=true`.
//
// With `include_settled=true`, settled and expired events are listed too: every
// market counts as listable for `categories`, `title`, `open_at.*`, `close_at.*`,
// `search`, and `sort_by`. Such a request that takes longer than 30 seconds fails
// with `504`; narrow `search` or the filters and try again.
//
// `open_at.*` and `close_at.*` bound an event's open and close time, which are the
// earliest across its listable markets, and `sort_by` uses the same times.
//
// An event's `title` is the question of its first listable market by symbol, and
// its `categories` are those of its listable markets.
//
// With `include_markets=true`, each event carries `markets`: a summary of every
// market of the event, ordered by symbol, in any status. This includes paused
// markets of a listed event and settled markets. Pages carry no totals: follow
// `next_page_token` until it is absent.
//
// `market_ids` (1-100 market ids or tickers, in any status) keeps the events those
// markets belong to, each once, and the other filters still apply. Ids that match
// no market are reported: `207` with the events found and the missing ids in
// `error.details`, or `404` when none matched.
func (r *V1InstrumentService) GetPredictionEvents(ctx context.Context, query V1InstrumentGetPredictionEventsParams, opts ...option.RequestOption) (res *V1InstrumentGetPredictionEventsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/instruments/predictions/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// > **Alpha** — this endpoint is experimental and may change or be removed at any
// > time.
//
// Looks up prediction markets in any status, by `market_ids` (1-100 market ids or
// tickers) or by `event_code`. Exactly one of the two is required. Markets are
// ordered by symbol.
//
// `statuses` keeps markets in any listed status. It defaults to `OPEN` with
// `event_code` and to every status with `market_ids`.
//
// With `event_code`, pages carry no totals: follow `next_page_token` until it is
// absent. `market_ids` lookups are not paginated and ignore page parameters.
//
// With `market_ids`, ids that match no market are reported: `207` with the markets
// found and the missing ids in `error.details`, or `404` when none matched. A
// market removed only by `statuses` is not reported as missing.
//
// Each market trades one YES-denominated instrument; there is no NO instrument.
// Prices, including those from `/market-data/snapshot`, are YES prices in dollars,
// and `no_price = 1 - yes_price`. A negative position quantity is NO exposure.
func (r *V1InstrumentService) GetPredictionMarkets(ctx context.Context, query V1InstrumentGetPredictionMarketsParams, opts ...option.RequestOption) (res *V1InstrumentGetPredictionMarketsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/instruments/predictions/markets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// > **Alpha** — this endpoint is experimental and may change or be removed at any
// > time.
//
// Lists the distinct series codes of prediction markets that can currently be
// listed: markets that are initialized or active and have not yet closed. Series
// are ordered by code.
//
// `categories` keeps series with a market in any listed category, ignoring case.
// Pages carry no totals: follow `next_page_token` until it is absent.
func (r *V1InstrumentService) GetPredictionSeries(ctx context.Context, query V1InstrumentGetPredictionSeriesParams, opts ...option.RequestOption) (res *V1InstrumentGetPredictionSeriesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/instruments/predictions/series"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search instruments by symbol, alternate identifier, or company name.
//
// The `q` parameter is case-insensitive and supports ticker symbols, alternate
// identifiers such as CUSIP, ISIN, and OPRA root, and company names for non-option
// instruments. Results are ranked by match quality plus instrument quality signals
// and relevance. Defaults to the `EQUITY` asset class (common stocks, preferred
// shares, ADRs, ETFs, and exchange-traded mutual funds). Pass `asset_class=OPTION`
// to search option contracts: by full OSI symbol, by an OSI prefix (root +
// `YYMMDD` expiry, e.g. `AAPL 261217`), or by a root-scoped phrase such as
// `AAPL Dec 250 call`.
func (r *V1InstrumentService) SearchInstruments(ctx context.Context, query V1InstrumentSearchInstrumentsParams, opts ...option.RequestOption) (res *V1InstrumentSearchInstrumentsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/instruments/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// The type of options contract
type ContractType string

const (
	ContractTypeCall ContractType = "CALL"
	ContractTypePut  ContractType = "PUT"
)

// The exercise style of an options contract
type ExerciseStyle string

const (
	ExerciseStyleAmerican ExerciseStyle = "AMERICAN"
	ExerciseStyleEuropean ExerciseStyle = "EUROPEAN"
)

// Represents a tradable financial instrument.
type Instrument struct {
	// Unique instrument identifier (UUID)
	ID string `json:"id" api:"required" format:"uuid"`
	// The ISO country code of the instrument's issue
	CountryOfIssue string `json:"country_of_issue" api:"required"`
	// The ISO currency code in which the instrument is traded
	Currency string `json:"currency" api:"required"`
	// Indicates if the instrument is classified as Easy-To-Borrow
	EasyToBorrow bool `json:"easy_to_borrow" api:"required"`
	// Indicates if the instrument supports fractional-quantity orders
	IsFractionable bool `json:"is_fractionable" api:"required"`
	// Indicates if the instrument is liquidation only and cannot be bought
	IsLiquidationOnly bool `json:"is_liquidation_only" api:"required"`
	// Indicates if the instrument is marginable
	IsMarginable bool `json:"is_marginable" api:"required"`
	// Indicates if the instrument is a publicly traded partnership (PTP). PTP sales
	// are subject to a 10% withholding tax for non-US tax residents.
	IsPtp bool `json:"is_ptp" api:"required"`
	// Indicates if short selling is prohibited for the instrument. This is a standing
	// property of the security. For the live Rule 201 circuit breaker, see
	// `short_sale_restricted` on the market-data snapshot.
	IsShortProhibited bool `json:"is_short_prohibited" api:"required"`
	// Indicates if the instrument is on the Regulation SHO Threshold Security List
	IsThresholdSecurity bool `json:"is_threshold_security" api:"required"`
	// Indicates if the instrument is tradable
	IsTradable bool `json:"is_tradable" api:"required"`
	// The trading symbol for the instrument
	Symbol string `json:"symbol" api:"required"`
	// The MIC code of the primary listing venue
	Venue string `json:"venue" api:"required"`
	// Average daily share volume from the security definition. When a null/undefined
	// value is observed, it indicates that there is no available data.
	Adv string `json:"adv" api:"nullable"`
	// Corporate-action-adjusted last close; present only when an adjustment exists for
	// the previous_close date. When a null/undefined value is observed, it indicates
	// that there is no available data.
	CaxAdjustedPreviousClose string `json:"cax_adjusted_previous_close" api:"nullable"`
	// The type of security (e.g., Common Stock, ETF) When a null/undefined value is
	// observed, it indicates that there is no available data.
	//
	// Any of "COMMON_STOCK", "INDEX", "OPTION", "CASH", "EVENT_CONTRACT".
	InstrumentType SecurityType `json:"instrument_type" api:"nullable"`
	// The percent of a long position's value you must post as margin When a
	// null/undefined value is observed, it indicates that there is no available data.
	LongMarginRate string `json:"long_margin_rate" api:"nullable"`
	// The full name of the instrument or its issuer When a null/undefined value is
	// observed, it indicates that there is no available data.
	Name string `json:"name" api:"nullable"`
	// Why the instrument is `NOT_SHORTABLE`: every reason that applies, using the same
	// codes as the `type` of the equivalent order-preview errors. Omitted unless
	// `shortable_attribute` is `NOT_SHORTABLE`. When a null/undefined value is
	// observed, it indicates it does not apply.
	NotShortableReasons []NotShortableReason `json:"not_shortable_reasons" api:"nullable"`
	// Notional average daily volume (ADV multiplied by the cax-adjusted close when
	// present, the raw previous close otherwise). When a null/undefined value is
	// observed, it indicates that there is no available data.
	NotionalAdv string `json:"notional_adv" api:"nullable"`
	// Available options expiration dates for this instrument, each annotated with
	// which settlement cycles have listed contracts on it. Present only when
	// `include_options_expiry_dates=true` in the request. When a null/undefined value
	// is observed, it indicates it does not apply.
	OptionsContractExpiryDates []OptionExpiryDate `json:"options_contract_expiry_dates" api:"nullable"`
	// Available options expiration dates for this instrument. Present only when
	// `include_options_expiry_dates=true` in the request.
	//
	// Deprecated: use `options_contract_expiry_dates`, which carries the same dates
	// annotated with settlement-cycle information. When a null/undefined value is
	// observed, it indicates it does not apply.
	//
	// Deprecated: deprecated
	OptionsExpiryDates []time.Time `json:"options_expiry_dates" api:"nullable" format:"date"`
	// Last close price from the security definition. When a null/undefined value is
	// observed, it indicates that there is no available data.
	PreviousClose string `json:"previous_close" api:"nullable"`
	// The percent of a short position's value you must post as margin When a
	// null/undefined value is observed, it indicates that there is no available data.
	ShortMarginRate string `json:"short_margin_rate" api:"nullable"`
	// How readily the instrument can be shorted: `EASY_TO_BORROW`, `HARD_TO_BORROW`,
	// or `NOT_SHORTABLE`. Easy-to-borrow instruments can be shorted without a locate.
	// A hard-to-borrow instrument is not easy to borrow but has locate inventory
	// remaining and supports short selling; anything else is not shortable. Omitted
	// for instruments it does not apply to. When a null/undefined value is observed,
	// it indicates it does not apply.
	//
	// Any of "EASY_TO_BORROW", "HARD_TO_BORROW", "NOT_SHORTABLE".
	ShortableAttribute ShortableAttribute `json:"shortable_attribute" api:"nullable"`
	// Price bands this instrument quotes on, ascending. A price outside every band is
	// not valid. Absent when we have no schedule for it, which includes an option
	// whose penny-program status our reference data never supplied.
	TickRules []TickRule `json:"tick_rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		CountryOfIssue             respjson.Field
		Currency                   respjson.Field
		EasyToBorrow               respjson.Field
		IsFractionable             respjson.Field
		IsLiquidationOnly          respjson.Field
		IsMarginable               respjson.Field
		IsPtp                      respjson.Field
		IsShortProhibited          respjson.Field
		IsThresholdSecurity        respjson.Field
		IsTradable                 respjson.Field
		Symbol                     respjson.Field
		Venue                      respjson.Field
		Adv                        respjson.Field
		CaxAdjustedPreviousClose   respjson.Field
		InstrumentType             respjson.Field
		LongMarginRate             respjson.Field
		Name                       respjson.Field
		NotShortableReasons        respjson.Field
		NotionalAdv                respjson.Field
		OptionsContractExpiryDates respjson.Field
		OptionsExpiryDates         respjson.Field
		PreviousClose              respjson.Field
		ShortMarginRate            respjson.Field
		ShortableAttribute         respjson.Field
		TickRules                  respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Instrument) RawJSON() string { return r.JSON.raw }
func (r *Instrument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstrumentCore struct {
	// Unique instrument identifier (UUID)
	ID string `json:"id" api:"required" format:"uuid"`
	// The ISO country code of the instrument's issue
	CountryOfIssue string `json:"country_of_issue" api:"required"`
	// The ISO currency code in which the instrument is traded
	Currency string `json:"currency" api:"required"`
	// Indicates if the instrument is classified as Easy-To-Borrow
	EasyToBorrow bool `json:"easy_to_borrow" api:"required"`
	// Indicates if the instrument supports fractional-quantity orders
	IsFractionable bool `json:"is_fractionable" api:"required"`
	// Indicates if the instrument is liquidation only and cannot be bought
	IsLiquidationOnly bool `json:"is_liquidation_only" api:"required"`
	// Indicates if the instrument is marginable
	IsMarginable bool `json:"is_marginable" api:"required"`
	// Indicates if the instrument is a publicly traded partnership (PTP). PTP sales
	// are subject to a 10% withholding tax for non-US tax residents.
	IsPtp bool `json:"is_ptp" api:"required"`
	// Indicates if short selling is prohibited for the instrument. This is a standing
	// property of the security. For the live Rule 201 circuit breaker, see
	// `short_sale_restricted` on the market-data snapshot.
	IsShortProhibited bool `json:"is_short_prohibited" api:"required"`
	// Indicates if the instrument is on the Regulation SHO Threshold Security List
	IsThresholdSecurity bool `json:"is_threshold_security" api:"required"`
	// Indicates if the instrument is tradable
	IsTradable bool `json:"is_tradable" api:"required"`
	// The trading symbol for the instrument
	Symbol string `json:"symbol" api:"required"`
	// The MIC code of the primary listing venue
	Venue string `json:"venue" api:"required"`
	// Average daily share volume from the security definition. When a null/undefined
	// value is observed, it indicates that there is no available data.
	Adv string `json:"adv" api:"nullable"`
	// Corporate-action-adjusted last close; present only when an adjustment exists for
	// the previous_close date. When a null/undefined value is observed, it indicates
	// that there is no available data.
	CaxAdjustedPreviousClose string `json:"cax_adjusted_previous_close" api:"nullable"`
	// The type of security (e.g., Common Stock, ETF) When a null/undefined value is
	// observed, it indicates that there is no available data.
	//
	// Any of "COMMON_STOCK", "INDEX", "OPTION", "CASH", "EVENT_CONTRACT".
	InstrumentType SecurityType `json:"instrument_type" api:"nullable"`
	// The percent of a long position's value you must post as margin When a
	// null/undefined value is observed, it indicates that there is no available data.
	LongMarginRate string `json:"long_margin_rate" api:"nullable"`
	// The full name of the instrument or its issuer When a null/undefined value is
	// observed, it indicates that there is no available data.
	Name string `json:"name" api:"nullable"`
	// Why the instrument is `NOT_SHORTABLE`: every reason that applies, using the same
	// codes as the `type` of the equivalent order-preview errors. Omitted unless
	// `shortable_attribute` is `NOT_SHORTABLE`. When a null/undefined value is
	// observed, it indicates it does not apply.
	NotShortableReasons []NotShortableReason `json:"not_shortable_reasons" api:"nullable"`
	// Notional average daily volume (ADV multiplied by the cax-adjusted close when
	// present, the raw previous close otherwise). When a null/undefined value is
	// observed, it indicates that there is no available data.
	NotionalAdv string `json:"notional_adv" api:"nullable"`
	// Last close price from the security definition. When a null/undefined value is
	// observed, it indicates that there is no available data.
	PreviousClose string `json:"previous_close" api:"nullable"`
	// The percent of a short position's value you must post as margin When a
	// null/undefined value is observed, it indicates that there is no available data.
	ShortMarginRate string `json:"short_margin_rate" api:"nullable"`
	// How readily the instrument can be shorted: `EASY_TO_BORROW`, `HARD_TO_BORROW`,
	// or `NOT_SHORTABLE`. Easy-to-borrow instruments can be shorted without a locate.
	// A hard-to-borrow instrument is not easy to borrow but has locate inventory
	// remaining and supports short selling; anything else is not shortable. Omitted
	// for instruments it does not apply to. When a null/undefined value is observed,
	// it indicates it does not apply.
	//
	// Any of "EASY_TO_BORROW", "HARD_TO_BORROW", "NOT_SHORTABLE".
	ShortableAttribute ShortableAttribute `json:"shortable_attribute" api:"nullable"`
	// Price bands this instrument quotes on, ascending. A price outside every band is
	// not valid. Absent when we have no schedule for it, which includes an option
	// whose penny-program status our reference data never supplied.
	TickRules []TickRule `json:"tick_rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		CountryOfIssue           respjson.Field
		Currency                 respjson.Field
		EasyToBorrow             respjson.Field
		IsFractionable           respjson.Field
		IsLiquidationOnly        respjson.Field
		IsMarginable             respjson.Field
		IsPtp                    respjson.Field
		IsShortProhibited        respjson.Field
		IsThresholdSecurity      respjson.Field
		IsTradable               respjson.Field
		Symbol                   respjson.Field
		Venue                    respjson.Field
		Adv                      respjson.Field
		CaxAdjustedPreviousClose respjson.Field
		InstrumentType           respjson.Field
		LongMarginRate           respjson.Field
		Name                     respjson.Field
		NotShortableReasons      respjson.Field
		NotionalAdv              respjson.Field
		PreviousClose            respjson.Field
		ShortMarginRate          respjson.Field
		ShortableAttribute       respjson.Field
		TickRules                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentCore) RawJSON() string { return r.JSON.raw }
func (r *InstrumentCore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstrumentCoreList []InstrumentCore

// The listing type of an options contract
type ListingType string

const (
	ListingTypeStandard ListingType = "STANDARD"
	ListingTypeFlex     ListingType = "FLEX"
	ListingTypeOtc      ListingType = "OTC"
)

// Why an instrument cannot be shorted.
type NotShortableReason string

const (
	NotShortableReasonInsufficientLocateInventory NotShortableReason = "INSUFFICIENT_LOCATE_INVENTORY"
	NotShortableReasonUnsupportedSecurity         NotShortableReason = "UNSUPPORTED_SECURITY"
)

// An options expiry date, annotated with which settlement cycles have listed
// contracts on it.
type OptionExpiryDate struct {
	// The expiration date.
	Date time.Time `json:"date" api:"required" format:"date"`
	// Whether this date has at least one listed contract that settles at the close (PM
	// settlement) -- the standard cycle.
	HasSettlesOnClose bool `json:"has_settles_on_close" api:"required"`
	// Whether this date has at least one contract that settles on the opening print
	// (AM settlement) and can still be traded. AM-settled contracts stop trading at
	// the close of the business day before settlement, so this turns false before the
	// expiration date arrives. A date leaves the list once no contract on it can be
	// traded in either settlement cycle.
	HasSettlesOnOpen bool `json:"has_settles_on_open" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date              respjson.Field
		HasSettlesOnClose respjson.Field
		HasSettlesOnOpen  respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptionExpiryDate) RawJSON() string { return r.JSON.raw }
func (r *OptionExpiryDate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An options contract with options-specific metadata
type OptionsContract struct {
	// Instrument identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether this is a CALL or PUT
	//
	// Any of "CALL", "PUT".
	ContractType ContractType `json:"contract_type" api:"required"`
	// ISO currency code
	Currency string `json:"currency" api:"required"`
	// MIC code of the primary listing venue
	Exchange string `json:"exchange" api:"required"`
	// Exercise style
	//
	// Any of "AMERICAN", "EUROPEAN".
	ExerciseStyle ExerciseStyle `json:"exercise_style" api:"required"`
	// Expiration date
	Expiry time.Time `json:"expiry" api:"required" format:"date"`
	// Whether the contract is liquidation-only
	IsLiquidationOnly bool `json:"is_liquidation_only" api:"required"`
	// Whether the contract is marginable
	IsMarginable bool `json:"is_marginable" api:"required"`
	// Whether the contract is tradable
	IsTradable bool `json:"is_tradable" api:"required"`
	// Listing type
	//
	// Any of "STANDARD", "FLEX", "OTC".
	ListingType ListingType `json:"listing_type" api:"required"`
	// Contract multiplier (100 for standard options)
	Multiplier string `json:"multiplier" api:"required"`
	// Strike price
	StrikePrice string `json:"strike_price" api:"required"`
	// OSI symbol (e.g. "AAPL 251219C00150000")
	Symbol string `json:"symbol" api:"required"`
	// Whether the option settles on the opening price (AM settlement), if known When a
	// null/undefined value is observed, it indicates that there is no available data.
	IsSettleOnOpen bool `json:"is_settle_on_open" api:"nullable"`
	// Last moment the option can trade (UTC), if known When a null/undefined value is
	// observed, it indicates that there is no available data.
	LastTradeCutoff time.Time `json:"last_trade_cutoff" api:"nullable" format:"date-time"`
	// Open interest (number of outstanding contracts), if available When a
	// null/undefined value is observed, it indicates that there is no available data.
	OpenInterest int64 `json:"open_interest" api:"nullable"`
	// Price bands this contract quotes on, ascending. Absent when our reference data
	// never supplied the contract's penny-program status.
	TickRules []TickRule `json:"tick_rules"`
	// Instrument ID of the underlying instrument, when available When a null/undefined
	// value is observed, it indicates that there is no available data.
	UnderlyingInstrumentID string `json:"underlying_instrument_id" api:"nullable" format:"uuid"`
	// Type of the underlying instrument, alongside `underlying_instrument_id` When a
	// null/undefined value is observed, it indicates that there is no available data.
	//
	// Any of "COMMON_STOCK", "INDEX", "OPTION", "CASH", "EVENT_CONTRACT".
	UnderlyingInstrumentType SecurityType `json:"underlying_instrument_type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		ContractType             respjson.Field
		Currency                 respjson.Field
		Exchange                 respjson.Field
		ExerciseStyle            respjson.Field
		Expiry                   respjson.Field
		IsLiquidationOnly        respjson.Field
		IsMarginable             respjson.Field
		IsTradable               respjson.Field
		ListingType              respjson.Field
		Multiplier               respjson.Field
		StrikePrice              respjson.Field
		Symbol                   respjson.Field
		IsSettleOnOpen           respjson.Field
		LastTradeCutoff          respjson.Field
		OpenInterest             respjson.Field
		TickRules                respjson.Field
		UnderlyingInstrumentID   respjson.Field
		UnderlyingInstrumentType respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OptionsContract) RawJSON() string { return r.JSON.raw }
func (r *OptionsContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OptionsContractList []OptionsContract

// Categories of the prediction market contracts that can currently be listed.
type PredictionCategories struct {
	// Distinct categories, sorted alphabetically.
	Categories []string `json:"categories" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionCategories) RawJSON() string { return r.JSON.raw }
func (r *PredictionCategories) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A prediction market event: a set of related markets that resolve together.
type PredictionEvent struct {
	// Distinct categories of the event's listable markets, sorted alphabetically.
	Categories []string `json:"categories" api:"required"`
	// Event code.
	Code string `json:"code" api:"required"`
	// Every market of the event in any status, ordered by symbol. Present only when
	// `include_markets` is `true`.
	Markets []PredictionMarketSummary `json:"markets" api:"nullable"`
	// The question of the event's first listable market, by symbol.
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories  respjson.Field
		Code        respjson.Field
		Markets     respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionEvent) RawJSON() string { return r.JSON.raw }
func (r *PredictionEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction market events with at least one market that can currently be listed.
type PredictionEvents struct {
	// Events in the requested order.
	Events []PredictionEvent `json:"events" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionEvents) RawJSON() string { return r.JSON.raw }
func (r *PredictionEvents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A prediction market.
type PredictionMarket struct {
	// Instrument id of the market.
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether the market may close before `close_at`.
	CanCloseEarly bool `json:"can_close_early" api:"required"`
	// Categories of the market.
	Categories []string `json:"categories" api:"required"`
	// Code of the event the market belongs to.
	EventCode string `json:"event_code" api:"required"`
	// Whether only orders that reduce a position are accepted.
	IsLiquidationOnly bool `json:"is_liquidation_only" api:"required"`
	// Whether the market accepts orders now.
	IsTradable bool `json:"is_tradable" api:"required"`
	// Trading restrictions of the market.
	Restrictions []string `json:"restrictions" api:"required"`
	// Code of the series the market belongs to.
	SeriesCode string `json:"series_code" api:"required"`
	// Current status of the market.
	//
	// Any of "OPEN", "UNOPENED", "PAUSED", "CLOSED", "SETTLED".
	Status PredictionMarketStatus `json:"status" api:"required"`
	// Market ticker.
	Symbol string `json:"symbol" api:"required"`
	// When trading is scheduled to close.
	CloseAt time.Time `json:"close_at" api:"nullable" format:"date-time"`
	// Link to the contract terms document.
	ContractTermsURL string `json:"contract_terms_url" api:"nullable"`
	// Condition under which the market closes early.
	EarlyCloseCondition string `json:"early_close_condition" api:"nullable"`
	// When settlement is expected after the market closes.
	ExpectedSettlementAt time.Time `json:"expected_settlement_at" api:"nullable" format:"date-time"`
	// Full title of the market, event and market together.
	FullTitle string `json:"full_title" api:"nullable"`
	// When trading opens.
	OpenAt time.Time `json:"open_at" api:"nullable" format:"date-time"`
	// Winning side. Set only when `status` is `SETTLED` and the market resolved `YES`
	// or `NO`. Absent otherwise, including for voided and scalar markets.
	//
	// Any of "YES", "NO".
	Outcome PredictionMarketOutcome `json:"outcome" api:"nullable"`
	// The question the market resolves.
	Predicate string `json:"predicate" api:"nullable"`
	// Primary rules of the market.
	RulesPrimary string `json:"rules_primary" api:"nullable"`
	// Secondary rules of the market.
	RulesSecondary string `json:"rules_secondary" api:"nullable"`
	// Settlement value of the YES side, in dollars per contract. Set only when
	// `status` is `SETTLED`, and absent otherwise.
	SettlementPrice string `json:"settlement_price" api:"nullable"`
	// Short title of the market, for example the name of the candidate.
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CanCloseEarly        respjson.Field
		Categories           respjson.Field
		EventCode            respjson.Field
		IsLiquidationOnly    respjson.Field
		IsTradable           respjson.Field
		Restrictions         respjson.Field
		SeriesCode           respjson.Field
		Status               respjson.Field
		Symbol               respjson.Field
		CloseAt              respjson.Field
		ContractTermsURL     respjson.Field
		EarlyCloseCondition  respjson.Field
		ExpectedSettlementAt respjson.Field
		FullTitle            respjson.Field
		OpenAt               respjson.Field
		Outcome              respjson.Field
		Predicate            respjson.Field
		RulesPrimary         respjson.Field
		RulesSecondary       respjson.Field
		SettlementPrice      respjson.Field
		Title                respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionMarket) RawJSON() string { return r.JSON.raw }
func (r *PredictionMarket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Winning side of a settled prediction market.
type PredictionMarketOutcome string

const (
	PredictionMarketOutcomeYes PredictionMarketOutcome = "YES"
	PredictionMarketOutcomeNo  PredictionMarketOutcome = "NO"
)

// Coarse status of a prediction market.
type PredictionMarketStatus string

const (
	PredictionMarketStatusOpen     PredictionMarketStatus = "OPEN"
	PredictionMarketStatusUnopened PredictionMarketStatus = "UNOPENED"
	PredictionMarketStatusPaused   PredictionMarketStatus = "PAUSED"
	PredictionMarketStatusClosed   PredictionMarketStatus = "CLOSED"
	PredictionMarketStatusSettled  PredictionMarketStatus = "SETTLED"
)

// Summary of a prediction market inside an event listing. Request the market by
// `id` or `symbol` for its full detail.
type PredictionMarketSummary struct {
	// Instrument id of the market.
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether only orders that reduce a position are accepted.
	IsLiquidationOnly bool `json:"is_liquidation_only" api:"required"`
	// Whether the market accepts orders now.
	IsTradable bool `json:"is_tradable" api:"required"`
	// Current status of the market.
	//
	// Any of "OPEN", "UNOPENED", "PAUSED", "CLOSED", "SETTLED".
	Status PredictionMarketStatus `json:"status" api:"required"`
	// Market ticker.
	Symbol string `json:"symbol" api:"required"`
	// When trading is scheduled to close.
	CloseAt time.Time `json:"close_at" api:"nullable" format:"date-time"`
	// When trading opens.
	OpenAt time.Time `json:"open_at" api:"nullable" format:"date-time"`
	// Winning side. Set only when `status` is `SETTLED` and the market resolved `YES`
	// or `NO`. Absent otherwise, including for voided and scalar markets.
	//
	// Any of "YES", "NO".
	Outcome PredictionMarketOutcome `json:"outcome" api:"nullable"`
	// Settlement value of the YES side, in dollars per contract. Set only when
	// `status` is `SETTLED`, and absent otherwise.
	SettlementPrice string `json:"settlement_price" api:"nullable"`
	// Short title of the market, for example the name of the candidate.
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		IsLiquidationOnly respjson.Field
		IsTradable        respjson.Field
		Status            respjson.Field
		Symbol            respjson.Field
		CloseAt           respjson.Field
		OpenAt            respjson.Field
		Outcome           respjson.Field
		SettlementPrice   respjson.Field
		Title             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionMarketSummary) RawJSON() string { return r.JSON.raw }
func (r *PredictionMarketSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction markets matching the request.
type PredictionMarkets struct {
	// Markets ordered by symbol.
	Markets []PredictionMarket `json:"markets" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Markets     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionMarkets) RawJSON() string { return r.JSON.raw }
func (r *PredictionMarkets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A prediction market series: a family of related markets that share a code.
type PredictionSeries struct {
	// Series code, shared by the markets in the family.
	Code string `json:"code" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionSeries) RawJSON() string { return r.JSON.raw }
func (r *PredictionSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction market series with at least one market that can currently be listed.
type PredictionSeriesList struct {
	// Series ordered by code.
	Series []PredictionSeries `json:"series" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Series      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionSeriesList) RawJSON() string { return r.JSON.raw }
func (r *PredictionSeriesList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How readily a tradable equity can be shorted.
type ShortableAttribute string

const (
	ShortableAttributeEasyToBorrow ShortableAttribute = "EASY_TO_BORROW"
	ShortableAttributeHardToBorrow ShortableAttribute = "HARD_TO_BORROW"
	ShortableAttributeNotShortable ShortableAttribute = "NOT_SHORTABLE"
)

// One band of an instrument's tick schedule. A price in the band is valid only if
// it is a whole multiple of `tick_size`. Bands describe the instrument itself: on
// an equity they say nothing about that equity's option chain.
type TickRule struct {
	// Lowest price in the band, inclusive.
	StartPrice string `json:"start_price" api:"required"`
	// Minimum price increment within the band.
	TickSize string `json:"tick_size" api:"required"`
	// Upper bound of the band, exclusive. Absent when the band runs to infinity. When
	// a null/undefined value is observed, it indicates it does not apply.
	EndPrice string `json:"end_price" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StartPrice  respjson.Field
		TickSize    respjson.Field
		EndPrice    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TickRule) RawJSON() string { return r.JSON.raw }
func (r *TickRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentGetInstrumentByIDResponse struct {
	// Represents a tradable financial instrument.
	Data Instrument `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentGetInstrumentByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentGetInstrumentByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentGetInstrumentsResponse struct {
	Data InstrumentCoreList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentGetInstrumentsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentGetInstrumentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentGetOptionContractsResponse struct {
	Data OptionsContractList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentGetOptionContractsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentGetOptionContractsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentGetPredictionCategoriesResponse struct {
	// Categories of the prediction market contracts that can currently be listed.
	Data PredictionCategories `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentGetPredictionCategoriesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentGetPredictionCategoriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentGetPredictionEventsResponse struct {
	// Prediction market events with at least one market that can currently be listed.
	Data PredictionEvents `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentGetPredictionEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentGetPredictionEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentGetPredictionMarketsResponse struct {
	// Prediction markets matching the request.
	Data PredictionMarkets `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentGetPredictionMarketsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentGetPredictionMarketsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentGetPredictionSeriesResponse struct {
	// Prediction market series with at least one market that can currently be listed.
	Data PredictionSeriesList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentGetPredictionSeriesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentGetPredictionSeriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentSearchInstrumentsResponse struct {
	Data InstrumentCoreList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1InstrumentSearchInstrumentsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentSearchInstrumentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentGetInstrumentByIDParams struct {
	// When true, include unique options expiry dates for this instrument
	IncludeOptionsExpiryDates param.Opt[bool] `query:"include_options_expiry_dates,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentGetInstrumentByIDParams]'s query parameters as
// `url.Values`.
func (r V1InstrumentGetInstrumentByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1InstrumentGetInstrumentsParams struct {
	// Filter by easy to borrow status
	EasyToBorrow param.Opt[bool] `query:"easy_to_borrow,omitzero" json:"-"`
	// Filter by liquidation only status
	IsLiquidationOnly param.Opt[bool] `query:"is_liquidation_only,omitzero" json:"-"`
	// Filter by marginable status
	IsMarginable param.Opt[bool] `query:"is_marginable,omitzero" json:"-"`
	// Filter by publicly traded partnership (PTP) status
	IsPtp param.Opt[bool] `query:"is_ptp,omitzero" json:"-"`
	// Filter by short prohibited status
	IsShortProhibited param.Opt[bool] `query:"is_short_prohibited,omitzero" json:"-"`
	// Filter by threshold security status
	IsThresholdSecurity param.Opt[bool] `query:"is_threshold_security,omitzero" json:"-"`
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Comma-separated instrument identifiers: unique identifiers or symbols (symbol
	// for equities, OSI for options)
	InstrumentIDs []string `query:"instrument_ids,omitzero" json:"-"`
	// Filter by instrument type (e.g. COMMON_STOCK, OPTION)
	//
	// Any of "COMMON_STOCK", "INDEX", "OPTION", "CASH", "EVENT_CONTRACT".
	InstrumentType V1InstrumentGetInstrumentsParamsInstrumentType `query:"instrument_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentGetInstrumentsParams]'s query parameters as
// `url.Values`.
func (r V1InstrumentGetInstrumentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by instrument type (e.g. COMMON_STOCK, OPTION)
type V1InstrumentGetInstrumentsParamsInstrumentType string

const (
	V1InstrumentGetInstrumentsParamsInstrumentTypeCommonStock   V1InstrumentGetInstrumentsParamsInstrumentType = "COMMON_STOCK"
	V1InstrumentGetInstrumentsParamsInstrumentTypeIndex         V1InstrumentGetInstrumentsParamsInstrumentType = "INDEX"
	V1InstrumentGetInstrumentsParamsInstrumentTypeOption        V1InstrumentGetInstrumentsParamsInstrumentType = "OPTION"
	V1InstrumentGetInstrumentsParamsInstrumentTypeCash          V1InstrumentGetInstrumentsParamsInstrumentType = "CASH"
	V1InstrumentGetInstrumentsParamsInstrumentTypeEventContract V1InstrumentGetInstrumentsParamsInstrumentType = "EVENT_CONTRACT"
)

type V1InstrumentGetOptionContractsParams struct {
	// Filter to contracts expiring on this date (YYYY-MM-DD)
	Expiry param.Opt[time.Time] `query:"expiry,omitzero" format:"date" json:"-"`
	// Filter by settlement cycle: true for early-settling (AM, settle-on-open)
	// contracts, false for normal (PM) contracts. Omit to return both.
	IsSettleOnOpen param.Opt[bool] `query:"is_settle_on_open,omitzero" json:"-"`
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// OSI option root symbol (e.g., AAPL, SPXW). Index roots also include their weekly
	// series (e.g., SPX includes SPXW).
	Root param.Opt[string] `query:"root,omitzero" json:"-"`
	// Deprecated: use `root`. If both are sent, the values must match.
	Underlier param.Opt[string] `query:"underlier,omitzero" json:"-"`
	// Instrument identifier or symbol of the underlying equity/index
	UnderlyingInstrumentID param.Opt[string] `query:"underlying_instrument_id,omitzero" json:"-"`
	// Comma-separated contract instrument IDs (UUID) or OSI option symbols to look up
	// directly, bypassing underlier expansion. Mutually exclusive with
	// root/underlying_instrument_id; up to 100 values.
	ContractIDs []InstrumentIDOrSymbol `query:"contract_ids,omitzero" json:"-"`
	// Filter by contract type: CALL or PUT
	//
	// Any of "CALL", "PUT".
	ContractType V1InstrumentGetOptionContractsParamsContractType `query:"contract_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentGetOptionContractsParams]'s query parameters as
// `url.Values`.
func (r V1InstrumentGetOptionContractsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by contract type: CALL or PUT
type V1InstrumentGetOptionContractsParamsContractType string

const (
	V1InstrumentGetOptionContractsParamsContractTypeCall V1InstrumentGetOptionContractsParamsContractType = "CALL"
	V1InstrumentGetOptionContractsParamsContractTypePut  V1InstrumentGetOptionContractsParamsContractType = "PUT"
)

type V1InstrumentGetPredictionEventsParams struct {
	// Comma-separated categories. Keeps events with a listable market in any of them,
	// ignoring case.
	Categories param.Opt[string] `query:"categories,omitzero" json:"-"`
	// Embed a summary of every market of each event, in any status. Default false.
	IncludeMarkets param.Opt[bool] `query:"include_markets,omitzero" json:"-"`
	// Also list settled and expired events, not only events with a listable market.
	// Every market then counts as listable for the other filters, `title`,
	// `categories`, and the open and close times. Default false. A request that takes
	// longer than 30 seconds fails with `504`; narrow `search` or the filters and try
	// again.
	IncludeSettled param.Opt[bool] `query:"include_settled,omitzero" json:"-"`
	// Comma-separated market ids (UUID) or tickers, 1-100. Keeps events of those
	// markets, in any status; the other filters still apply. Ids that match no market
	// are reported: `207`, or `404` when none match.
	MarketIDs param.Opt[string] `query:"market_ids,omitzero" json:"-"`
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Keeps events whose code, series code, or a market's question or title contains
	// every word, matched by prefix and English stem, ignoring case. Searches listable
	// events, or every event with `include_settled`. A word of four or more letters
	// that matches nothing is replaced by up to three close spellings. Without
	// `sort_by`, orders by relevance, closest spellings first.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Comma-separated series codes. Keeps events of those series, ignoring case.
	SeriesCodes param.Opt[string]                            `query:"series_codes,omitzero" json:"-"`
	CloseAt     V1InstrumentGetPredictionEventsParamsCloseAt `query:"close_at,omitzero" json:"-"`
	OpenAt      V1InstrumentGetPredictionEventsParamsOpenAt  `query:"open_at,omitzero" json:"-"`
	// Sort key. Defaults to relevance when `search` is set, else the event code. Ties
	// break on event code ascending.
	//
	// Any of "OPEN_AT", "CLOSE_AT".
	SortBy V1InstrumentGetPredictionEventsParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Sort direction. Defaults to ascending. Ignored when sorting by relevance.
	//
	// Any of "ASC", "DESC".
	SortDirection V1InstrumentGetPredictionEventsParamsSortDirection `query:"sort_direction,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentGetPredictionEventsParams]'s query parameters
// as `url.Values`.
func (r V1InstrumentGetPredictionEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1InstrumentGetPredictionEventsParamsCloseAt struct {
	// Return only rows where `close_at` is strictly after the given value. A bare
	// `YYYY-MM-DD` date expands to the end of that day (UTC), so this matches from the
	// start of the following day. See
	// [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters)
	// for accepted formats, bare-date expansion, and combining bounds. Returns 400 if
	// the resulting range is inverted.
	Gt param.Opt[string] `query:"gt,omitzero" json:"-"`
	// Return only rows where `close_at` is on or after the given value. A bare
	// `YYYY-MM-DD` date expands to the start of that day (UTC). See
	// [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters)
	// for accepted formats, bare-date expansion, and combining bounds. Returns 400 if
	// the resulting range is inverted.
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	// Return only rows where `close_at` is strictly before the given value. A bare
	// `YYYY-MM-DD` date expands to the start of that day (UTC). See
	// [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters)
	// for accepted formats, bare-date expansion, and combining bounds. Returns 400 if
	// the resulting range is inverted.
	Lt param.Opt[string] `query:"lt,omitzero" json:"-"`
	// Return only rows where `close_at` is on or before the given value. A bare
	// `YYYY-MM-DD` date expands to the end of that day (UTC), so this matches through
	// the end of that day. See
	// [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters)
	// for accepted formats, bare-date expansion, and combining bounds. Returns 400 if
	// the resulting range is inverted.
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentGetPredictionEventsParamsCloseAt]'s query
// parameters as `url.Values`.
func (r V1InstrumentGetPredictionEventsParamsCloseAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1InstrumentGetPredictionEventsParamsOpenAt struct {
	// Return only rows where `open_at` is strictly after the given value. A bare
	// `YYYY-MM-DD` date expands to the end of that day (UTC), so this matches from the
	// start of the following day. See
	// [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters)
	// for accepted formats, bare-date expansion, and combining bounds. Returns 400 if
	// the resulting range is inverted.
	Gt param.Opt[string] `query:"gt,omitzero" json:"-"`
	// Return only rows where `open_at` is on or after the given value. A bare
	// `YYYY-MM-DD` date expands to the start of that day (UTC). See
	// [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters)
	// for accepted formats, bare-date expansion, and combining bounds. Returns 400 if
	// the resulting range is inverted.
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	// Return only rows where `open_at` is strictly before the given value. A bare
	// `YYYY-MM-DD` date expands to the start of that day (UTC). See
	// [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters)
	// for accepted formats, bare-date expansion, and combining bounds. Returns 400 if
	// the resulting range is inverted.
	Lt param.Opt[string] `query:"lt,omitzero" json:"-"`
	// Return only rows where `open_at` is on or before the given value. A bare
	// `YYYY-MM-DD` date expands to the end of that day (UTC), so this matches through
	// the end of that day. See
	// [Range filters](https://docs.clearstreet.com/guides/api-fundamentals#range-filters)
	// for accepted formats, bare-date expansion, and combining bounds. Returns 400 if
	// the resulting range is inverted.
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentGetPredictionEventsParamsOpenAt]'s query
// parameters as `url.Values`.
func (r V1InstrumentGetPredictionEventsParamsOpenAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort key. Defaults to relevance when `search` is set, else the event code. Ties
// break on event code ascending.
type V1InstrumentGetPredictionEventsParamsSortBy string

const (
	V1InstrumentGetPredictionEventsParamsSortByOpenAt  V1InstrumentGetPredictionEventsParamsSortBy = "OPEN_AT"
	V1InstrumentGetPredictionEventsParamsSortByCloseAt V1InstrumentGetPredictionEventsParamsSortBy = "CLOSE_AT"
)

// Sort direction. Defaults to ascending. Ignored when sorting by relevance.
type V1InstrumentGetPredictionEventsParamsSortDirection string

const (
	V1InstrumentGetPredictionEventsParamsSortDirectionAsc  V1InstrumentGetPredictionEventsParamsSortDirection = "ASC"
	V1InstrumentGetPredictionEventsParamsSortDirectionDesc V1InstrumentGetPredictionEventsParamsSortDirection = "DESC"
)

type V1InstrumentGetPredictionMarketsParams struct {
	// Event code; returns the markets of the event. Mutually exclusive with
	// `market_ids`.
	EventCode param.Opt[string] `query:"event_code,omitzero" json:"-"`
	// Comma-separated market ids (UUID) or tickers, 1-100, in any status. Mutually
	// exclusive with `event_code`.
	MarketIDs param.Opt[string] `query:"market_ids,omitzero" json:"-"`
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Comma-separated statuses to keep: `OPEN`, `UNOPENED`, `PAUSED`, `CLOSED`,
	// `SETTLED`. Defaults to `OPEN` with `event_code` and to every status with
	// `market_ids`. Unrecognized values are ignored.
	Statuses param.Opt[string] `query:"statuses,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentGetPredictionMarketsParams]'s query parameters
// as `url.Values`.
func (r V1InstrumentGetPredictionMarketsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1InstrumentGetPredictionSeriesParams struct {
	// Comma-separated categories. Keeps series with a market in any of them, ignoring
	// case.
	Categories param.Opt[string] `query:"categories,omitzero" json:"-"`
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentGetPredictionSeriesParams]'s query parameters
// as `url.Values`.
func (r V1InstrumentGetPredictionSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1InstrumentSearchInstrumentsParams struct {
	// Search term applied case-insensitively to ticker symbols, alternate identifiers
	// (CUSIP, ISIN, OPRA root), and company names for non-option instruments. Option
	// searches match symbols and alternate identifiers.
	Q string `query:"q" api:"required" json:"-"`
	// Comma-separated asset classes (EQUITY|OPTION|WARRANT|BOND|FX|OTHER). Defaults to
	// EQUITY.
	AssetClass param.Opt[string] `query:"asset_class,omitzero" json:"-"`
	// Optional listing-country filter (e.g., US).
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// Optional ISO currency filter (e.g., USD).
	Currency param.Opt[string] `query:"currency,omitzero" json:"-"`
	// Include inactive instruments. Default false.
	IncludeInactive param.Opt[bool] `query:"include_inactive,omitzero" json:"-"`
	// Include publicly traded partnership (PTP) instruments. Default true (penalized
	// in ranking).
	IncludePtp param.Opt[bool] `query:"include_ptp,omitzero" json:"-"`
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentSearchInstrumentsParams]'s query parameters as
// `url.Values`.
func (r V1InstrumentSearchInstrumentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
