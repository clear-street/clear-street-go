// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

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
// V1PrivateMarketService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1PrivateMarketService] method instead.
type V1PrivateMarketService struct {
	options []option.RequestOption
	Iois    V1PrivateMarketIoisService
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

// Create an IOI for a visible upcoming offering.
func (r *V1PrivateMarketService) NewIoi(ctx context.Context, params V1PrivateMarketNewIoiParams, opts ...option.RequestOption) (res *V1PrivateMarketNewIoiResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/private-markets/iois"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Withdraw a live IOI. Repeating a withdrawal returns 404.
func (r *V1PrivateMarketService) DeleteIoi(ctx context.Context, ioiID string, body V1PrivateMarketDeleteIoiParams, opts ...option.RequestOption) (err error) {
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
func (r *V1PrivateMarketService) GetIois(ctx context.Context, query V1PrivateMarketGetIoisParams, opts ...option.RequestOption) (res *V1PrivateMarketGetIoisResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/private-markets/iois"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Update an IOI's notional, accepting the current NDA revision when required.
func (r *V1PrivateMarketService) UpdateIoi(ctx context.Context, ioiID string, params V1PrivateMarketUpdateIoiParams, opts ...option.RequestOption) (res *V1PrivateMarketUpdateIoiResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if ioiID == "" {
		err = errors.New("missing required ioi_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/private-markets/iois/%s", ioiID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

type V1PrivateMarketNewIoiResponse struct {
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
func (r V1PrivateMarketNewIoiResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PrivateMarketNewIoiResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PrivateMarketGetIoisResponse struct {
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
func (r V1PrivateMarketGetIoisResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PrivateMarketGetIoisResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PrivateMarketUpdateIoiResponse struct {
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
func (r V1PrivateMarketUpdateIoiResponse) RawJSON() string { return r.JSON.raw }
func (r *V1PrivateMarketUpdateIoiResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PrivateMarketNewIoiParams struct {
	AccountID      int64  `query:"account_id" api:"required" json:"-"`
	NotionalAmount string `json:"notional_amount" api:"required"`
	OfferingID     string `json:"offering_id" api:"required" format:"uuid"`
	// Required only when the offering's attached SPV has an NDA agreement.
	NdaAcceptance V1PrivateMarketNewIoiParamsNdaAcceptance `json:"nda_acceptance,omitzero"`
	paramObj
}

func (r V1PrivateMarketNewIoiParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PrivateMarketNewIoiParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PrivateMarketNewIoiParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [V1PrivateMarketNewIoiParams]'s query parameters as
// `url.Values`.
func (r V1PrivateMarketNewIoiParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Required only when the offering's attached SPV has an NDA agreement.
//
// The properties Accepted, AgreementID, AuthorityConfirmed are required.
type V1PrivateMarketNewIoiParamsNdaAcceptance struct {
	// Must be true; confirms affirmative assent.
	Accepted bool `json:"accepted" api:"required"`
	// Exact agreement id returned by offering detail.
	AgreementID string `json:"agreement_id" api:"required" format:"uuid"`
	// Must be true; confirms the signer may bind the account-holder entity.
	AuthorityConfirmed bool `json:"authority_confirmed" api:"required"`
	paramObj
}

func (r V1PrivateMarketNewIoiParamsNdaAcceptance) MarshalJSON() (data []byte, err error) {
	type shadow V1PrivateMarketNewIoiParamsNdaAcceptance
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PrivateMarketNewIoiParamsNdaAcceptance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1PrivateMarketDeleteIoiParams struct {
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1PrivateMarketDeleteIoiParams]'s query parameters as
// `url.Values`.
func (r V1PrivateMarketDeleteIoiParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1PrivateMarketGetIoisParams struct {
	AccountID int64 `query:"account_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1PrivateMarketGetIoisParams]'s query parameters as
// `url.Values`.
func (r V1PrivateMarketGetIoisParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1PrivateMarketUpdateIoiParams struct {
	AccountID      int64  `query:"account_id" api:"required" json:"-"`
	NotionalAmount string `json:"notional_amount" api:"required"`
	// Required when the SPV's current NDA version is newer than the IOI's latest
	// acceptance. Irrelevant acceptances are rejected.
	NdaAcceptance V1PrivateMarketUpdateIoiParamsNdaAcceptance `json:"nda_acceptance,omitzero"`
	paramObj
}

func (r V1PrivateMarketUpdateIoiParams) MarshalJSON() (data []byte, err error) {
	type shadow V1PrivateMarketUpdateIoiParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PrivateMarketUpdateIoiParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [V1PrivateMarketUpdateIoiParams]'s query parameters as
// `url.Values`.
func (r V1PrivateMarketUpdateIoiParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Required when the SPV's current NDA version is newer than the IOI's latest
// acceptance. Irrelevant acceptances are rejected.
//
// The properties Accepted, AgreementID, AuthorityConfirmed are required.
type V1PrivateMarketUpdateIoiParamsNdaAcceptance struct {
	// Must be true; confirms affirmative assent.
	Accepted bool `json:"accepted" api:"required"`
	// Exact agreement id returned by offering detail.
	AgreementID string `json:"agreement_id" api:"required" format:"uuid"`
	// Must be true; confirms the signer may bind the account-holder entity.
	AuthorityConfirmed bool `json:"authority_confirmed" api:"required"`
	paramObj
}

func (r V1PrivateMarketUpdateIoiParamsNdaAcceptance) MarshalJSON() (data []byte, err error) {
	type shadow V1PrivateMarketUpdateIoiParamsNdaAcceptance
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1PrivateMarketUpdateIoiParamsNdaAcceptance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
