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

// Browse private-market offerings and their indicative terms. Access requires the
// account holder to hold an accreditation attestation.
//
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

// Create an IOI for a visible upcoming offering.
func (r *V1PrivateMarketIoisService) NewIoi(ctx context.Context, params V1PrivateMarketIoisNewIoiParams, opts ...option.RequestOption) (res *V1PrivateMarketIoisNewIoiResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/private-markets/iois"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Withdraw a live IOI. Repeating a withdrawal returns 404.
func (r *V1PrivateMarketIoisService) DeleteIoi(ctx context.Context, ioiID string, body V1PrivateMarketIoisDeleteIoiParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if ioiID == "" {
		err = errors.New("missing required ioi_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/private-markets/iois/%s", ioiID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// List every live IOI for the caller's account-holder entity.
func (r *V1PrivateMarketIoisService) GetIois(ctx context.Context, query V1PrivateMarketIoisGetIoisParams, opts ...option.RequestOption) (res *V1PrivateMarketIoisGetIoisResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/private-markets/iois"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Update an IOI's notional, accepting the current NDA revision when required.
func (r *V1PrivateMarketIoisService) UpdateIoi(ctx context.Context, ioiID string, params V1PrivateMarketIoisUpdateIoiParams, opts ...option.RequestOption) (res *V1PrivateMarketIoisUpdateIoiResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if ioiID == "" {
		err = errors.New("missing required ioi_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/private-markets/iois/%s", ioiID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
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

type V1PrivateMarketIoisNewIoiResponse struct {
	// IOI list item with the campaign identity needed to render it.
	Data IoiListingResource `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1PrivateMarketIoisNewIoiResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PrivateMarketIoisNewIoiResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PrivateMarketIoisGetIoisResponse struct {
	Data IoiListingResourceList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1PrivateMarketIoisGetIoisResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PrivateMarketIoisGetIoisResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PrivateMarketIoisUpdateIoiResponse struct {
	// IOI list item with the campaign identity needed to render it.
	Data IoiListingResource `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1PrivateMarketIoisUpdateIoiResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PrivateMarketIoisUpdateIoiResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PrivateMarketIoisNewIoiParams struct {
	AccountID      int64  `query:"account_id" api:"required" json:"-"`
	NotionalAmount string `json:"notional_amount" api:"required"`
	OfferingID     string `json:"offering_id" api:"required" format:"uuid"`
	// Required only when the offering's attached SPV has an NDA agreement.
	NdaAcceptance V1PrivateMarketIoisNewIoiParamsNdaAcceptance `json:"nda_acceptance,omitzero"`
	paramObj
}

func (r V1PrivateMarketIoisNewIoiParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PrivateMarketIoisNewIoiParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PrivateMarketIoisNewIoiParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [V1PrivateMarketIoisNewIoiParams]'s query parameters as
// `url.Values`.
func (r V1PrivateMarketIoisNewIoiParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Required only when the offering's attached SPV has an NDA agreement.
//
// The properties Accepted, AgreementID, AuthorityConfirmed are required.
type V1PrivateMarketIoisNewIoiParamsNdaAcceptance struct {
	// Must be true; confirms affirmative assent.
	Accepted bool `json:"accepted" api:"required"`
	// Exact agreement id returned by offering detail.
	AgreementID string `json:"agreement_id" api:"required" format:"uuid"`
	// Must be true; confirms the signer may bind the account-holder entity.
	AuthorityConfirmed bool `json:"authority_confirmed" api:"required"`
	paramObj
}

func (r V1PrivateMarketIoisNewIoiParamsNdaAcceptance) MarshalJSON() (data []byte, err error) {
	type shadow V1PrivateMarketIoisNewIoiParamsNdaAcceptance
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PrivateMarketIoisNewIoiParamsNdaAcceptance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PrivateMarketIoisDeleteIoiParams struct {
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1PrivateMarketIoisDeleteIoiParams]'s query parameters as
// `url.Values`.
func (r V1PrivateMarketIoisDeleteIoiParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1PrivateMarketIoisGetIoisParams struct {
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1PrivateMarketIoisGetIoisParams]'s query parameters as
// `url.Values`.
func (r V1PrivateMarketIoisGetIoisParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1PrivateMarketIoisUpdateIoiParams struct {
	AccountID      int64  `query:"account_id" api:"required" json:"-"`
	NotionalAmount string `json:"notional_amount" api:"required"`
	// Required when the SPV's current NDA version is newer than the IOI's latest
	// acceptance. Irrelevant acceptances are rejected.
	NdaAcceptance V1PrivateMarketIoisUpdateIoiParamsNdaAcceptance `json:"nda_acceptance,omitzero"`
	paramObj
}

func (r V1PrivateMarketIoisUpdateIoiParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PrivateMarketIoisUpdateIoiParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PrivateMarketIoisUpdateIoiParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [V1PrivateMarketIoisUpdateIoiParams]'s query parameters as
// `url.Values`.
func (r V1PrivateMarketIoisUpdateIoiParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Required when the SPV's current NDA version is newer than the IOI's latest
// acceptance. Irrelevant acceptances are rejected.
//
// The properties Accepted, AgreementID, AuthorityConfirmed are required.
type V1PrivateMarketIoisUpdateIoiParamsNdaAcceptance struct {
	// Must be true; confirms affirmative assent.
	Accepted bool `json:"accepted" api:"required"`
	// Exact agreement id returned by offering detail.
	AgreementID string `json:"agreement_id" api:"required" format:"uuid"`
	// Must be true; confirms the signer may bind the account-holder entity.
	AuthorityConfirmed bool `json:"authority_confirmed" api:"required"`
	paramObj
}

func (r V1PrivateMarketIoisUpdateIoiParamsNdaAcceptance) MarshalJSON() (data []byte, err error) {
	type shadow V1PrivateMarketIoisUpdateIoiParamsNdaAcceptance
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PrivateMarketIoisUpdateIoiParamsNdaAcceptance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
