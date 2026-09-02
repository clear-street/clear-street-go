// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"time"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/respjson"
)

// V1PrivateMarketCompanyService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1PrivateMarketCompanyService] method instead.
type V1PrivateMarketCompanyService struct {
	options []option.RequestOption
}

// NewV1PrivateMarketCompanyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1PrivateMarketCompanyService(opts ...option.RequestOption) (r V1PrivateMarketCompanyService) {
	r = V1PrivateMarketCompanyService{}
	r.options = opts
	return
}

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
