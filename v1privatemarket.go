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
// V1PrivateMarketService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1PrivateMarketService] method instead.
type V1PrivateMarketService struct {
	options []option.RequestOption
	// Browse private-market offerings and their indicative terms. Access requires the
	// account holder to hold an accreditation attestation.
	Iois V1PrivateMarketIoisService
	// Browse private-market offerings and their indicative terms. Access requires the
	// account holder to hold an accreditation attestation.
	Offerings V1PrivateMarketOfferingService
}

// NewV1PrivateMarketService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1PrivateMarketService(opts ...option.RequestOption) (r V1PrivateMarketService) {
	r = V1PrivateMarketService{}
	r.options = opts
	r.Iois = NewV1PrivateMarketIoisService(opts...)
	r.Offerings = NewV1PrivateMarketOfferingService(opts...)
	return
}

// Fetch one published private-market company with its complete versioned profile.
// Requires the account holder to have attested. Returns `404` when the company
// does not exist or is not yet published.
func (r *V1PrivateMarketService) GetCompanyByID(ctx context.Context, companyID string, query V1PrivateMarketGetCompanyByIDParams, opts ...option.RequestOption) (res *V1PrivateMarketGetCompanyByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/private-markets/companies/%s", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Fetch one private-market SPV's complete economics and fee schedule. Requires the
// account holder to have attested. Returns `404` unless the SPV is `OPEN` and
// attached to a currently visible `ACTIVE` offering.
func (r *V1PrivateMarketService) GetSpvByID(ctx context.Context, spvID string, query V1PrivateMarketGetSpvByIDParams, opts ...option.RequestOption) (res *V1PrivateMarketGetSpvByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if spvID == "" {
		err = errors.New("missing required spv_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/private-markets/spvs/%s", spvID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Party charging a fee.
type ChargedBy string

const (
	ChargedByFundManager ChargedBy = "FUND_MANAGER"
	ChargedByClearStreet ChargedBy = "CLEAR_STREET"
	ChargedByThirdParty  ChargedBy = "THIRD_PARTY"
)

// A company category.
type CompanyCategory struct {
	// Display name.
	Name string `json:"name" api:"required"`
	// Stable lowercase category slug.
	Slug string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyCategory) RawJSON() string { return r.JSON.raw }
func (r *CompanyCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A cited source.
type CompanyCitation struct {
	// Stable profile-local citation identifier.
	ID string `json:"id" api:"required"`
	// Source publisher or provider.
	Source string `json:"source" api:"required"`
	// Human-readable source title.
	Title string `json:"title" api:"required"`
	// Source URL.
	URL string `json:"url" api:"required"`
	// Source publication time, when known.
	PublishedAt time.Time `json:"published_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Source      respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		PublishedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyCitation) RawJSON() string { return r.JSON.raw }
func (r *CompanyCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A named company customer.
type CompanyCustomer struct {
	// Customer name.
	Name string `json:"name" api:"required"`
	// Customer logo, when supplied.
	LogoURL string `json:"logo_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		LogoURL     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyCustomer) RawJSON() string { return r.JSON.raw }
func (r *CompanyCustomer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A company's identity and its complete published profile.
type CompanyDetail struct {
	// Stable company identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// Display name.
	Name string `json:"name" api:"required"`
	// The complete versioned company profile.
	Profile CompanyProfileResource `json:"profile" api:"required"`
	// Profile schema version discriminator.
	ProfileSchemaVersion int64 `json:"profile_schema_version" api:"required"`
	// Short card/search description.
	ShortDescription string `json:"short_description" api:"required"`
	// Lowercase URL slug.
	Slug string `json:"slug" api:"required"`
	// Company logo URL, when known.
	LogoURL string `json:"logo_url" api:"nullable"`
	// Canonical lowercase domain, when known.
	PrimaryDomain string `json:"primary_domain" api:"nullable"`
	// Publication time.
	PublishedAt time.Time `json:"published_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Name                 respjson.Field
		Profile              respjson.Field
		ProfileSchemaVersion respjson.Field
		ShortDescription     respjson.Field
		Slug                 respjson.Field
		LogoURL              respjson.Field
		PrimaryDomain        respjson.Field
		PublishedAt          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyDetail) RawJSON() string { return r.JSON.raw }
func (r *CompanyDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional document card preview.
type CompanyDocumentPreview struct {
	// Preview description.
	Description string `json:"description" api:"nullable"`
	// Preview image URL.
	ImageURL string `json:"image_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		ImageURL    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyDocumentPreview) RawJSON() string { return r.JSON.raw }
func (r *CompanyDocumentPreview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How a document relates to the company.
type CompanyDocumentRelation string

const (
	CompanyDocumentRelationSubject   CompanyDocumentRelation = "SUBJECT"
	CompanyDocumentRelationConnected CompanyDocumentRelation = "CONNECTED"
)

// A company-level research or source document.
type CompanyDocumentResource struct {
	// Typed document kind.
	//
	// Any of "COMPANY_PROFILE", "MARKET_RESEARCH", "INTERVIEW", "DEAL_SHEET",
	// "PRESS_RELEASE", "NEWS", "OTHER".
	DocumentType CompanyDocumentType `json:"document_type" api:"required"`
	// Relationship to this company.
	//
	// Any of "SUBJECT", "CONNECTED".
	Relation CompanyDocumentRelation `json:"relation" api:"required"`
	// Display title.
	Title string `json:"title" api:"required"`
	// Document URL.
	URL string `json:"url" api:"required"`
	// Optional source identifier retained for reconciliation.
	ExternalID string `json:"external_id" api:"nullable"`
	// Optional card preview.
	Preview CompanyDocumentPreview `json:"preview" api:"nullable"`
	// Publication time, when known.
	PublishedAt time.Time `json:"published_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentType respjson.Field
		Relation     respjson.Field
		Title        respjson.Field
		URL          respjson.Field
		ExternalID   respjson.Field
		Preview      respjson.Field
		PublishedAt  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyDocumentResource) RawJSON() string { return r.JSON.raw }
func (r *CompanyDocumentResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Company document kind.
type CompanyDocumentType string

const (
	CompanyDocumentTypeCompanyProfile CompanyDocumentType = "COMPANY_PROFILE"
	CompanyDocumentTypeMarketResearch CompanyDocumentType = "MARKET_RESEARCH"
	CompanyDocumentTypeInterview      CompanyDocumentType = "INTERVIEW"
	CompanyDocumentTypeDealSheet      CompanyDocumentType = "DEAL_SHEET"
	CompanyDocumentTypePressRelease   CompanyDocumentType = "PRESS_RELEASE"
	CompanyDocumentTypeNews           CompanyDocumentType = "NEWS"
	CompanyDocumentTypeOther          CompanyDocumentType = "OTHER"
)

// Company headquarters.
type CompanyHeadquarters struct {
	// City.
	City string `json:"city" api:"required"`
	// Country.
	Country string `json:"country" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyHeadquarters) RawJSON() string { return r.JSON.raw }
func (r *CompanyHeadquarters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A legal entity associated with the company.
type CompanyLegalEntity struct {
	// Country name or ISO country code supplied by the source.
	Country string `json:"country" api:"required"`
	// Legal name.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyLegalEntity) RawJSON() string { return r.JSON.raw }
func (r *CompanyLegalEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One metric observation.
type CompanyMetricPoint struct {
	// Observation time.
	ObservedAt time.Time `json:"observed_at" api:"required" format:"date-time"`
	// Exact decimal value, serialized as a string.
	Value string `json:"value" api:"required"`
	// Historical or estimated classification.
	//
	// Any of "HISTORICAL", "ESTIMATED".
	ValueType MetricValueType `json:"value_type" api:"required"`
	// Profile-local citation ids supporting this point.
	CitationIDs []string `json:"citation_ids"`
	// Optional source event identifier.
	SourceEventID string `json:"source_event_id" api:"nullable"`
	// Optional provider reconciliation metadata.
	SourceMetadata map[string]string `json:"source_metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObservedAt     respjson.Field
		Value          respjson.Field
		ValueType      respjson.Field
		CitationIDs    respjson.Field
		SourceEventID  respjson.Field
		SourceMetadata respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyMetricPoint) RawJSON() string { return r.JSON.raw }
func (r *CompanyMetricPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A historical or estimated company metric series.
type CompanyMetricSeries struct {
	// Observation cadence.
	//
	// Any of "YEAR", "QUARTER", "MONTH", "POINT_IN_TIME".
	Frequency MetricFrequency `json:"frequency" api:"required"`
	// Display label.
	Label string `json:"label" api:"required"`
	// Canonical metric key.
	//
	// Any of "ANNUALIZED_REVENUE", "REVENUE_GROWTH", "VALUATION", "ISSUE_PRICE",
	// "PRICE_PER_SHARE", "AMOUNT_RAISED", "ORDER_VOLUME", "PIPELINE_VALUE",
	// "GROSS_MARGIN", "EBIT_MARGIN", "FCF_CONVERSION", "CONTRACTED_REVENUE_PERCENT",
	// "NET_REVENUE_RETENTION", "CUSTOMER_COUNT", "MARKET_POSITION".
	MetricKey MetricKey `json:"metric_key" api:"required"`
	// Publisher/provider name.
	Source string `json:"source" api:"required"`
	// Value unit.
	//
	// Any of "USD", "PERCENT", "COUNT", "RANK".
	Unit MetricUnit `json:"unit" api:"required"`
	// Optional source identifier retained for reconciliation.
	ExternalID string `json:"external_id" api:"nullable"`
	// Ordered observations.
	Points []CompanyMetricPoint `json:"points"`
	// Source URL, when available.
	SourceURL string `json:"source_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Frequency   respjson.Field
		Label       respjson.Field
		MetricKey   respjson.Field
		Source      respjson.Field
		Unit        respjson.Field
		ExternalID  respjson.Field
		Points      respjson.Field
		SourceURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyMetricSeries) RawJSON() string { return r.JSON.raw }
func (r *CompanyMetricSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One ordered durable narrative block.
type CompanyNarrativeSection struct {
	// Plain-text section body.
	Body string `json:"body" api:"required"`
	// Stable display position within the profile.
	DisplayOrder int64 `json:"display_order" api:"required"`
	// Section heading.
	Title string `json:"title" api:"required"`
	// Profile-local citation ids supporting this block.
	CitationIDs []string `json:"citation_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body         respjson.Field
		DisplayOrder respjson.Field
		Title        respjson.Field
		CitationIDs  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyNarrativeSection) RawJSON() string { return r.JSON.raw }
func (r *CompanyNarrativeSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key person associated with the company.
type CompanyPerson struct {
	// Display name.
	Name string `json:"name" api:"required"`
	// Optional source identifier retained for reconciliation.
	ExternalID string `json:"external_id" api:"nullable"`
	// One or more curated company roles.
	Roles []CompanyPersonRole `json:"roles"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExternalID  respjson.Field
		Roles       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyPerson) RawJSON() string { return r.JSON.raw }
func (r *CompanyPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key person's relationship to the company.
type CompanyPersonRole string

const (
	CompanyPersonRoleFounder CompanyPersonRole = "FOUNDER"
	CompanyPersonRoleCeo     CompanyPersonRole = "CEO"
	CompanyPersonRoleOther   CompanyPersonRole = "OTHER"
)

// The complete versioned company profile (schema version one).
type CompanyProfileResource struct {
	// Company categories.
	Categories []CompanyCategory `json:"categories"`
	// Sources referenced by narrative sections and metrics.
	Citations []CompanyCitation `json:"citations"`
	// Named customers evidenced by the source material.
	Customers []CompanyCustomer `json:"customers"`
	// Company-level research and source documents.
	Documents []CompanyDocumentResource `json:"documents"`
	// Company headquarters, when known.
	Headquarters CompanyHeadquarters `json:"headquarters" api:"nullable"`
	// Known legal entities associated with the company.
	LegalEntities []CompanyLegalEntity `json:"legal_entities"`
	// Historical and estimated metric series.
	MetricSeries []CompanyMetricSeries `json:"metric_series"`
	// Ordered durable company fact and thesis blocks.
	NarrativeSections []CompanyNarrativeSection `json:"narrative_sections"`
	// Long company overview.
	Overview string `json:"overview" api:"nullable"`
	// Key people and their roles.
	People []CompanyPerson `json:"people"`
	// Social/profile links.
	Social []CompanySocialLink `json:"social"`
	// Short durable positioning line used with the company name.
	Tagline string `json:"tagline" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories        respjson.Field
		Citations         respjson.Field
		Customers         respjson.Field
		Documents         respjson.Field
		Headquarters      respjson.Field
		LegalEntities     respjson.Field
		MetricSeries      respjson.Field
		NarrativeSections respjson.Field
		Overview          respjson.Field
		People            respjson.Field
		Social            respjson.Field
		Tagline           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyProfileResource) RawJSON() string { return r.JSON.raw }
func (r *CompanyProfileResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A company social/profile link.
type CompanySocialLink struct {
	// Link type.
	//
	// Any of "WEBSITE", "LINKEDIN", "X", "FACEBOOK", "OTHER".
	Type CompanySocialType `json:"type" api:"required"`
	// Link URL.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanySocialLink) RawJSON() string { return r.JSON.raw }
func (r *CompanySocialLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Kind of company social/profile link.
type CompanySocialType string

const (
	CompanySocialTypeWebsite  CompanySocialType = "WEBSITE"
	CompanySocialTypeLinkedin CompanySocialType = "LINKEDIN"
	CompanySocialTypeX        CompanySocialType = "X"
	CompanySocialTypeFacebook CompanySocialType = "FACEBOOK"
	CompanySocialTypeOther    CompanySocialType = "OTHER"
)

// Fee timing/cadence.
type FeeFrequency string

const (
	FeeFrequencyOneTime     FeeFrequency = "ONE_TIME"
	FeeFrequencyAnnual      FeeFrequency = "ANNUAL"
	FeeFrequencyAtExit      FeeFrequency = "AT_EXIT"
	FeeFrequencyPassThrough FeeFrequency = "PASS_THROUGH"
)

// Kind of SPV fee.
type FeeType string

const (
	FeeTypeManagement     FeeType = "MANAGEMENT"
	FeeTypeCarry          FeeType = "CARRY"
	FeeTypePlacement      FeeType = "PLACEMENT"
	FeeTypeAdministrative FeeType = "ADMINISTRATIVE"
	FeeTypeOther          FeeType = "OTHER"
)

// Observation cadence for a metric series.
type MetricFrequency string

const (
	MetricFrequencyYear        MetricFrequency = "YEAR"
	MetricFrequencyQuarter     MetricFrequency = "QUARTER"
	MetricFrequencyMonth       MetricFrequency = "MONTH"
	MetricFrequencyPointInTime MetricFrequency = "POINT_IN_TIME"
)

// Canonical company metric key.
type MetricKey string

const (
	MetricKeyAnnualizedRevenue        MetricKey = "ANNUALIZED_REVENUE"
	MetricKeyRevenueGrowth            MetricKey = "REVENUE_GROWTH"
	MetricKeyValuation                MetricKey = "VALUATION"
	MetricKeyIssuePrice               MetricKey = "ISSUE_PRICE"
	MetricKeyPricePerShare            MetricKey = "PRICE_PER_SHARE"
	MetricKeyAmountRaised             MetricKey = "AMOUNT_RAISED"
	MetricKeyOrderVolume              MetricKey = "ORDER_VOLUME"
	MetricKeyPipelineValue            MetricKey = "PIPELINE_VALUE"
	MetricKeyGrossMargin              MetricKey = "GROSS_MARGIN"
	MetricKeyEbitMargin               MetricKey = "EBIT_MARGIN"
	MetricKeyFcfConversion            MetricKey = "FCF_CONVERSION"
	MetricKeyContractedRevenuePercent MetricKey = "CONTRACTED_REVENUE_PERCENT"
	MetricKeyNetRevenueRetention      MetricKey = "NET_REVENUE_RETENTION"
	MetricKeyCustomerCount            MetricKey = "CUSTOMER_COUNT"
	MetricKeyMarketPosition           MetricKey = "MARKET_POSITION"
)

// An OPEN SPV's identity, exact economics, and typed fee schedule.
type SpvDetail struct {
	// Stable SPV identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// Company whose shares the vehicle holds.
	CompanyID string `json:"company_id" api:"required" format:"uuid"`
	// Terms currency.
	//
	// Any of "USD".
	Currency Currency `json:"currency" api:"required"`
	// Legal/display name.
	Name string `json:"name" api:"required"`
	// Lifecycle state.
	//
	// Any of "DRAFT", "OPEN", "CLOSED", "LIQUIDATING", "DISSOLVED".
	Status SpvStatus `json:"status" api:"required"`
	// Price per share including fees.
	AllInPricePerShare string `json:"all_in_price_per_share" api:"nullable"`
	// Custodian.
	CustodianName string `json:"custodian_name" api:"nullable"`
	// Per-share fee.
	FeePerShare string `json:"fee_per_share" api:"nullable"`
	// Typed fee schedule.
	FeeTerms []SpvFeeTermResource `json:"fee_terms"`
	// Percentage of dollar allocation funded, derived from the allocation pair.
	FundedPercent string `json:"funded_percent" api:"nullable"`
	// Funding deadline.
	FundingDeadline time.Time `json:"funding_deadline" api:"nullable" format:"date-time"`
	// SPV manager.
	ManagerName string `json:"manager_name" api:"nullable"`
	// Minimum investment amount.
	MinimumInvestmentAmount string `json:"minimum_investment_amount" api:"nullable"`
	// Time the vehicle opened.
	OpenedAt time.Time `json:"opened_at" api:"nullable" format:"date-time"`
	// Price per share excluding fees.
	PricePerShare string `json:"price_per_share" api:"nullable"`
	// Remaining dollar allocation.
	RemainingAllocationAmount string `json:"remaining_allocation_amount" api:"nullable"`
	// Remaining share allocation.
	RemainingShareAllocation string `json:"remaining_share_allocation" api:"nullable"`
	// Underlying share class, when specified.
	ShareClass string `json:"share_class" api:"nullable"`
	// Plain-text vehicle structure.
	StructureDescription string `json:"structure_description" api:"nullable"`
	// Total dollar allocation.
	TotalAllocationAmount string `json:"total_allocation_amount" api:"nullable"`
	// Total share allocation.
	TotalShareAllocation string `json:"total_share_allocation" api:"nullable"`
	// Exact company valuation.
	Valuation string `json:"valuation" api:"nullable"`
	// Meaning of `valuation`.
	//
	// Any of "PRE_MONEY", "POST_MONEY", "REFERENCE", "IMPLIED".
	ValuationBasis ValuationBasis `json:"valuation_basis" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		CompanyID                 respjson.Field
		Currency                  respjson.Field
		Name                      respjson.Field
		Status                    respjson.Field
		AllInPricePerShare        respjson.Field
		CustodianName             respjson.Field
		FeePerShare               respjson.Field
		FeeTerms                  respjson.Field
		FundedPercent             respjson.Field
		FundingDeadline           respjson.Field
		ManagerName               respjson.Field
		MinimumInvestmentAmount   respjson.Field
		OpenedAt                  respjson.Field
		PricePerShare             respjson.Field
		RemainingAllocationAmount respjson.Field
		RemainingShareAllocation  respjson.Field
		ShareClass                respjson.Field
		StructureDescription      respjson.Field
		TotalAllocationAmount     respjson.Field
		TotalShareAllocation      respjson.Field
		Valuation                 respjson.Field
		ValuationBasis            respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpvDetail) RawJSON() string { return r.JSON.raw }
func (r *SpvDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One typed SPV fee term.
type SpvFeeTermResource struct {
	// Charging party.
	//
	// Any of "FUND_MANAGER", "CLEAR_STREET", "THIRD_PARTY".
	ChargedBy ChargedBy `json:"charged_by" api:"required"`
	// Terms currency.
	//
	// Any of "USD".
	Currency Currency `json:"currency" api:"required"`
	// Plain-text fee disclosure.
	Description string `json:"description" api:"required"`
	// Fee kind.
	//
	// Any of "MANAGEMENT", "CARRY", "PLACEMENT", "ADMINISTRATIVE", "OTHER".
	FeeType FeeType `json:"fee_type" api:"required"`
	// Timing/cadence.
	//
	// Any of "ONE_TIME", "ANNUAL", "AT_EXIT", "PASS_THROUGH".
	Frequency FeeFrequency `json:"frequency" api:"required"`
	// Exact fixed amount, when amount-based.
	Amount string `json:"amount" api:"nullable"`
	// Charge duration in years, when specified.
	DurationYears string `json:"duration_years" api:"nullable"`
	// Carry hurdle as a decimal fraction, when specified.
	HurdleRate string `json:"hurdle_rate" api:"nullable"`
	// Decimal fraction between zero and one, when percentage-based.
	Rate string `json:"rate" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChargedBy     respjson.Field
		Currency      respjson.Field
		Description   respjson.Field
		FeeType       respjson.Field
		Frequency     respjson.Field
		Amount        respjson.Field
		DurationYears respjson.Field
		HurdleRate    respjson.Field
		Rate          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpvFeeTermResource) RawJSON() string { return r.JSON.raw }
func (r *SpvFeeTermResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PrivateMarketGetCompanyByIDResponse struct {
	// A company's identity and its complete published profile.
	Data CompanyDetail `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1PrivateMarketGetCompanyByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PrivateMarketGetCompanyByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PrivateMarketGetSpvByIDResponse struct {
	// An OPEN SPV's identity, exact economics, and typed fee schedule.
	Data SpvDetail `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1PrivateMarketGetSpvByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PrivateMarketGetSpvByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PrivateMarketGetCompanyByIDParams struct {
	// Account whose account-holder entity must hold an accreditation attestation to
	// browse private-market offerings.
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1PrivateMarketGetCompanyByIDParams]'s query parameters as
// `url.Values`.
func (r V1PrivateMarketGetCompanyByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1PrivateMarketGetSpvByIDParams struct {
	// Account whose account-holder entity must hold an accreditation attestation to
	// browse private-market offerings.
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1PrivateMarketGetSpvByIDParams]'s query parameters as
// `url.Values`.
func (r V1PrivateMarketGetSpvByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
