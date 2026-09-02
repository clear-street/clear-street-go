// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"time"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/respjson"
)

// V1PrivateMarketIoisService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1PrivateMarketIoisService] method instead.
type V1PrivateMarketIoisService struct {
	options []option.RequestOption
}

// NewV1PrivateMarketIoisService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1PrivateMarketIoisService(opts ...option.RequestOption) (r V1PrivateMarketIoisService) {
	r = V1PrivateMarketIoisService{}
	r.options = opts
	return
}

// Company identity embedded in an IOI list item.
type IoiCompanyResource struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IoiCompanyResource) RawJSON() string { return r.JSON.raw }
func (r *IoiCompanyResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// IOI list item with the campaign identity needed to render it.
type IoiListingResource struct {
	// Company identity embedded in an IOI list item.
	Company IoiCompanyResource `json:"company" api:"required"`
	// Offering identity embedded in an IOI list item.
	Offering IoiOfferingResource `json:"offering" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company     respjson.Field
		Offering    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	IoiResource
}

// Returns the unmodified JSON received from the API
func (r IoiListingResource) RawJSON() string { return r.JSON.raw }
func (r *IoiListingResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IoiListingResourceList []IoiListingResource

// Offering identity embedded in an IOI list item.
type IoiOfferingResource struct {
	ID       string `json:"id" api:"required" format:"uuid"`
	Headline string `json:"headline" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Headline    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IoiOfferingResource) RawJSON() string { return r.JSON.raw }
func (r *IoiOfferingResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One live indication of interest.
type IoiResource struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	AccountID int64     `json:"account_id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Terms currency.
	//
	// Any of "USD".
	Currency       Currency  `json:"currency" api:"required"`
	NotionalAmount string    `json:"notional_amount" api:"required"`
	OfferingID     string    `json:"offering_id" api:"required" format:"uuid"`
	UpdatedAt      time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Most recent NDA acceptance linked to this IOI, if any.
	NdaAcceptance NdaAcceptanceResource `json:"nda_acceptance" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AccountID      respjson.Field
		CreatedAt      respjson.Field
		Currency       respjson.Field
		NotionalAmount respjson.Field
		OfferingID     respjson.Field
		UpdatedAt      respjson.Field
		NdaAcceptance  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IoiResource) RawJSON() string { return r.JSON.raw }
func (r *IoiResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public evidence that an NDA version was accepted. Signing IP and other
// provenance remain audit-only and are never returned by this API.
type NdaAcceptanceResource struct {
	AcceptedAt  time.Time `json:"accepted_at" api:"required" format:"date-time"`
	AgreementID string    `json:"agreement_id" api:"required" format:"uuid"`
	Version     int64     `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptedAt  respjson.Field
		AgreementID respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NdaAcceptanceResource) RawJSON() string { return r.JSON.raw }
func (r *NdaAcceptanceResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
