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
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Browse private-market offerings and their indicative terms. Access requires the
// account holder to hold an accreditation attestation.
//
// V1PrivateMarketOfferingService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1PrivateMarketOfferingService] method instead.
type V1PrivateMarketOfferingService struct {
	options []option.RequestOption
}

// NewV1PrivateMarketOfferingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1PrivateMarketOfferingService(opts ...option.RequestOption) (r V1PrivateMarketOfferingService) {
	r = V1PrivateMarketOfferingService{}
	r.options = opts
	return
}

// Fetch one visible private-market offering with its documents, participants, and
// any attached SPV. Requires the account holder to have attested. Returns `404`
// when the offering does not exist or is not currently visible.
func (r *V1PrivateMarketOfferingService) GetOfferingByID(ctx context.Context, offeringID string, query V1PrivateMarketOfferingGetOfferingByIDParams, opts ...option.RequestOption) (res *V1PrivateMarketOfferingGetOfferingByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if offeringID == "" {
		err = errors.New("missing required offering_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/private-markets/offerings/%s", offeringID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List every visible private-market offering as a card, with its derived class,
// company and SPV identity, and indicative terms. Requires the account holder to
// have attested.
func (r *V1PrivateMarketOfferingService) GetOfferings(ctx context.Context, query V1PrivateMarketOfferingGetOfferingsParams, opts ...option.RequestOption) (res *V1PrivateMarketOfferingGetOfferingsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/private-markets/offerings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Terms currency.
type Currency string

const (
	CurrencyUsd Currency = "USD"
)

// Unit for a resolved highlight metric.
type MetricUnit string

const (
	MetricUnitUsd     MetricUnit = "USD"
	MetricUnitPercent MetricUnit = "PERCENT"
	MetricUnitCount   MetricUnit = "COUNT"
	MetricUnitRank    MetricUnit = "RANK"
)

// Whether a resolved highlight value is observed or estimated.
type MetricValueType string

const (
	MetricValueTypeHistorical MetricValueType = "HISTORICAL"
	MetricValueTypeEstimated  MetricValueType = "ESTIMATED"
)

// One offering as it appears in a list: its derived class, indicative terms, a
// company identity summary, and any attached SPV.
type OfferingCard struct {
	// Stable public identifier; IOIs and history hang off it.
	ID string `json:"id" api:"required" format:"uuid"`
	// Derived classification.
	//
	// Any of "UPCOMING", "ACTIVE".
	Class OfferingClass `json:"class" api:"required"`
	// Owning company identity.
	Company OfferingCompany `json:"company" api:"required"`
	// Terms currency.
	//
	// Any of "USD".
	Currency Currency `json:"currency" api:"required"`
	// Card/detail headline.
	Headline string `json:"headline" api:"required"`
	// Top opportunity paragraph.
	Summary string `json:"summary" api:"required"`
	// Indicative price-per-share range, high endpoint.
	IndicativePriceHigh string `json:"indicative_price_high" api:"nullable"`
	// Indicative price-per-share range, low endpoint.
	IndicativePriceLow string `json:"indicative_price_low" api:"nullable"`
	// Meaning of the indicative valuation range.
	//
	// Any of "PRE_MONEY", "POST_MONEY", "REFERENCE", "IMPLIED".
	IndicativeValuationBasis ValuationBasis `json:"indicative_valuation_basis" api:"nullable"`
	// Indicative valuation range, high endpoint.
	IndicativeValuationHigh string `json:"indicative_valuation_high" api:"nullable"`
	// Indicative valuation range, low endpoint.
	IndicativeValuationLow string `json:"indicative_valuation_low" api:"nullable"`
	// Deadline for indications of interest.
	IoiDeadline time.Time `json:"ioi_deadline" api:"nullable" format:"date-time"`
	// Minimum indication-of-interest amount.
	MinimumIoiAmount string `json:"minimum_ioi_amount" api:"nullable"`
	// Attached SPV identity and lifecycle, once one exists.
	Spv OfferingSpv `json:"spv" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		Class                    respjson.Field
		Company                  respjson.Field
		Currency                 respjson.Field
		Headline                 respjson.Field
		Summary                  respjson.Field
		IndicativePriceHigh      respjson.Field
		IndicativePriceLow       respjson.Field
		IndicativeValuationBasis respjson.Field
		IndicativeValuationHigh  respjson.Field
		IndicativeValuationLow   respjson.Field
		IoiDeadline              respjson.Field
		MinimumIoiAmount         respjson.Field
		Spv                      respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OfferingCard) RawJSON() string { return r.JSON.raw }
func (r *OfferingCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OfferingCardList []OfferingCard

// Derived offering classification.
type OfferingClass string

const (
	OfferingClassUpcoming OfferingClass = "UPCOMING"
	OfferingClassActive   OfferingClass = "ACTIVE"
)

// Company identity carried on an offering card/detail.
type OfferingCompany struct {
	// Stable company identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Short card/search description.
	ShortDescription string `json:"short_description" api:"required"`
	// Lowercase URL slug.
	Slug string `json:"slug" api:"required"`
	// Company logo URL, when known.
	LogoURL string `json:"logo_url" api:"nullable"`
	// Canonical lowercase domain, when known.
	PrimaryDomain string `json:"primary_domain" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Name             respjson.Field
		ShortDescription respjson.Field
		Slug             respjson.Field
		LogoURL          respjson.Field
		PrimaryDomain    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OfferingCompany) RawJSON() string { return r.JSON.raw }
func (r *OfferingCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One offering with everything needed to render its detail payload.
type OfferingDetail struct {
	// Important disclosures.
	Disclosures string `json:"disclosures" api:"nullable"`
	// Campaign documents in display order.
	Documents []OfferingDocumentResource `json:"documents"`
	// Ordered resolved highlights.
	Highlights []OfferingHighlight `json:"highlights"`
	// Campaign-specific investment framing.
	InvestmentThesis string `json:"investment_thesis" api:"nullable"`
	// Ordered key risks.
	KeyRisks []OfferingKeyRisk `json:"key_risks"`
	// Campaign participants in display order.
	Participants []OfferingParticipantResource `json:"participants"`
	// Vehicle/structure framing shown before typed SPV terms exist.
	StructureDescription string `json:"structure_description" api:"nullable"`
	// Why-now framing.
	WhyNow string `json:"why_now" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Disclosures          respjson.Field
		Documents            respjson.Field
		Highlights           respjson.Field
		InvestmentThesis     respjson.Field
		KeyRisks             respjson.Field
		Participants         respjson.Field
		StructureDescription respjson.Field
		WhyNow               respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
	OfferingCard
}

// Returns the unmodified JSON received from the API
func (r OfferingDetail) RawJSON() string { return r.JSON.raw }
func (r *OfferingDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A campaign document's display metadata. Exactly one of `url`/`object_key` is
// set; an object key is resolved and signed elsewhere.
type OfferingDocumentResource struct {
	// Stable identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// Stable display position.
	DisplayOrder int64 `json:"display_order" api:"required"`
	// Document kind.
	//
	// Any of "TEARSHEET", "KEY_TERMS", "RISK_FACTORS", "PPM", "OTHER".
	DocumentType OfferingDocumentType `json:"document_type" api:"required"`
	// Display title.
	Title string `json:"title" api:"required"`
	// Object-store key, when the document is stored internally.
	ObjectKey string `json:"object_key" api:"nullable"`
	// Publication time, when known.
	PublishedAt time.Time `json:"published_at" api:"nullable" format:"date-time"`
	// Source publisher/provider.
	Source string `json:"source" api:"nullable"`
	// Source URL.
	SourceURL string `json:"source_url" api:"nullable"`
	// Externally reachable URL, when the document lives at one.
	URL string `json:"url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DisplayOrder respjson.Field
		DocumentType respjson.Field
		Title        respjson.Field
		ObjectKey    respjson.Field
		PublishedAt  respjson.Field
		Source       respjson.Field
		SourceURL    respjson.Field
		URL          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OfferingDocumentResource) RawJSON() string { return r.JSON.raw }
func (r *OfferingDocumentResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Kind of campaign document.
type OfferingDocumentType string

const (
	OfferingDocumentTypeTearsheet   OfferingDocumentType = "TEARSHEET"
	OfferingDocumentTypeKeyTerms    OfferingDocumentType = "KEY_TERMS"
	OfferingDocumentTypeRiskFactors OfferingDocumentType = "RISK_FACTORS"
	OfferingDocumentTypePpm         OfferingDocumentType = "PPM"
	OfferingDocumentTypeOther       OfferingDocumentType = "OTHER"
)

// A curated highlight, resolved against the company profile's metric series.
type OfferingHighlight struct {
	// Display label (the highlight's override, else the series' own label).
	Label string `json:"label" api:"required"`
	// Canonical metric key selected by the highlight (e.g. `REVENUE_GROWTH`).
	MetricKey string `json:"metric_key" api:"required"`
	// Value unit.
	//
	// Any of "USD", "PERCENT", "COUNT", "RANK".
	Unit MetricUnit `json:"unit" api:"required"`
	// Observation time of the latest value.
	ObservedAt time.Time `json:"observed_at" api:"nullable" format:"date-time"`
	// Latest observed value, when the series carries any points.
	Value string `json:"value" api:"nullable"`
	// Whether the latest value is historical or estimated.
	//
	// Any of "HISTORICAL", "ESTIMATED".
	ValueType MetricValueType `json:"value_type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		MetricKey   respjson.Field
		Unit        respjson.Field
		ObservedAt  respjson.Field
		Value       respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OfferingHighlight) RawJSON() string { return r.JSON.raw }
func (r *OfferingHighlight) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One ordered key-risk block.
type OfferingKeyRisk struct {
	// Plain-text risk body.
	Body string `json:"body" api:"required"`
	// Risk heading.
	Title string `json:"title" api:"required"`
	// Profile-local citation ids supporting the risk.
	CitationIDs []string `json:"citation_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Title       respjson.Field
		CitationIDs respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OfferingKeyRisk) RawJSON() string { return r.JSON.raw }
func (r *OfferingKeyRisk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An offering participant's display data.
type OfferingParticipantResource struct {
	// Stable identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// Stable display position.
	DisplayOrder int64 `json:"display_order" api:"required"`
	// Display name.
	Name string `json:"name" api:"required"`
	// Presentation role.
	//
	// Any of "LEAD_INVESTOR", "CO_LEAD", "FUND_MANAGER", "PLACEMENT_AGENT".
	Role ParticipantRole `json:"role" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DisplayOrder respjson.Field
		Name         respjson.Field
		Role         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OfferingParticipantResource) RawJSON() string { return r.JSON.raw }
func (r *OfferingParticipantResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The attached SPV's identity and lifecycle. Exact economics surface once the SPV
// opens; an upcoming offering's indicative ranges describe the terms until then.
type OfferingSpv struct {
	// Stable SPV identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// Legal/display name.
	Name string `json:"name" api:"required"`
	// Lifecycle state.
	//
	// Any of "DRAFT", "OPEN", "CLOSED", "LIQUIDATING", "DISSOLVED".
	Status SpvStatus `json:"status" api:"required"`
	// Custodian.
	CustodianName string `json:"custodian_name" api:"nullable"`
	// SPV manager.
	ManagerName string `json:"manager_name" api:"nullable"`
	// Underlying share class, when specified.
	ShareClass string `json:"share_class" api:"nullable"`
	// Plain-text vehicle structure.
	StructureDescription string `json:"structure_description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Name                 respjson.Field
		Status               respjson.Field
		CustodianName        respjson.Field
		ManagerName          respjson.Field
		ShareClass           respjson.Field
		StructureDescription respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OfferingSpv) RawJSON() string { return r.JSON.raw }
func (r *OfferingSpv) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Presentation role of an offering participant.
type ParticipantRole string

const (
	ParticipantRoleLeadInvestor   ParticipantRole = "LEAD_INVESTOR"
	ParticipantRoleCoLead         ParticipantRole = "CO_LEAD"
	ParticipantRoleFundManager    ParticipantRole = "FUND_MANAGER"
	ParticipantRolePlacementAgent ParticipantRole = "PLACEMENT_AGENT"
)

// SPV lifecycle state.
type SpvStatus string

const (
	SpvStatusDraft       SpvStatus = "DRAFT"
	SpvStatusOpen        SpvStatus = "OPEN"
	SpvStatusClosed      SpvStatus = "CLOSED"
	SpvStatusLiquidating SpvStatus = "LIQUIDATING"
	SpvStatusDissolved   SpvStatus = "DISSOLVED"
)

// Meaning of an indicative valuation range or an SPV valuation.
type ValuationBasis string

const (
	ValuationBasisPreMoney  ValuationBasis = "PRE_MONEY"
	ValuationBasisPostMoney ValuationBasis = "POST_MONEY"
	ValuationBasisReference ValuationBasis = "REFERENCE"
	ValuationBasisImplied   ValuationBasis = "IMPLIED"
)

type V1PrivateMarketOfferingGetOfferingByIDResponse struct {
	// One offering with everything needed to render its detail payload.
	Data OfferingDetail `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1PrivateMarketOfferingGetOfferingByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PrivateMarketOfferingGetOfferingByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PrivateMarketOfferingGetOfferingsResponse struct {
	Data OfferingCardList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1PrivateMarketOfferingGetOfferingsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PrivateMarketOfferingGetOfferingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PrivateMarketOfferingGetOfferingByIDParams struct {
	// Account whose account-holder entity must hold an accreditation attestation to
	// browse private-market offerings.
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1PrivateMarketOfferingGetOfferingByIDParams]'s query
// parameters as `url.Values`.
func (r V1PrivateMarketOfferingGetOfferingByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1PrivateMarketOfferingGetOfferingsParams struct {
	// Account whose account-holder entity must hold an accreditation attestation to
	// browse private-market offerings.
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1PrivateMarketOfferingGetOfferingsParams]'s query
// parameters as `url.Values`.
func (r V1PrivateMarketOfferingGetOfferingsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
