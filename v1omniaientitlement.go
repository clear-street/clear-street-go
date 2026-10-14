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

// Thread-centric AI assistant for conversational trading. Create threads to start
// conversations, poll response objects for in-progress output, and read finalized
// messages from thread history. Thread/message/response endpoints require an
// explicit account_id. Entitlement endpoints are caller-scoped and use
// account_ids.
//
// V1OmniAIEntitlementService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1OmniAIEntitlementService] method instead.
type V1OmniAIEntitlementService struct {
	options []option.RequestOption
}

// NewV1OmniAIEntitlementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1OmniAIEntitlementService(opts ...option.RequestOption) (r V1OmniAIEntitlementService) {
	r = V1OmniAIEntitlementService{}
	r.options = opts
	return
}

// Record consent and upsert one-or-more active grants.
func (r *V1OmniAIEntitlementService) NewEntitlements(ctx context.Context, body V1OmniAIEntitlementNewEntitlementsParams, opts ...option.RequestOption) (res *V1OmniAIEntitlementNewEntitlementsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/omni-ai/entitlements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Revoke one entitlement grant by id.
//
// Deprecated -- use `POST /omni-ai/entitlements/opt-out`.
//
// Deprecated: deprecated
func (r *V1OmniAIEntitlementService) DeleteEntitlement(ctx context.Context, entitlementID string, opts ...option.RequestOption) (res *V1OmniAIEntitlementDeleteEntitlementResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if entitlementID == "" {
		err = errors.New("missing required entitlement_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/omni-ai/entitlements/%s", url.PathEscape(entitlementID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// List current signable entitlement agreements for consent UX.
func (r *V1OmniAIEntitlementService) GetEntitlementAgreements(ctx context.Context, opts ...option.RequestOption) (res *V1OmniAIEntitlementGetEntitlementAgreementsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/omni-ai/entitlement-agreements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Whether each entitlement is enabled, per account. Accounts with no recorded
// choice are enabled by default. Without `account_id`, covers every account the
// caller can access.
func (r *V1OmniAIEntitlementService) GetEntitlementStatus(ctx context.Context, query V1OmniAIEntitlementGetEntitlementStatusParams, opts ...option.RequestOption) (res *V1OmniAIEntitlementGetEntitlementStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/omni-ai/entitlement-status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List caller's active entitlement grants.
//
// Deprecated -- use `GET /omni-ai/entitlement-status`; this lists explicit grants
// only and omits accounts that are enabled by default.
//
// Deprecated: deprecated
func (r *V1OmniAIEntitlementService) GetEntitlements(ctx context.Context, query V1OmniAIEntitlementGetEntitlementsParams, opts ...option.RequestOption) (res *V1OmniAIEntitlementGetEntitlementsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/omni-ai/entitlements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Turn off one-or-more entitlements for the given accounts. Idempotent: accounts
// that are already opted out are left as they are. To turn an entitlement back on,
// sign the agreement with `POST /omni-ai/entitlements`.
func (r *V1OmniAIEntitlementService) OptOutEntitlements(ctx context.Context, body V1OmniAIEntitlementOptOutEntitlementsParams, opts ...option.RequestOption) (res *V1OmniAIEntitlementOptOutEntitlementsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/omni-ai/entitlements/opt-out"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type DeleteEntitlementResponse struct {
	EntitlementID string `json:"entitlement_id" api:"required"`
	Revoked       bool   `json:"revoked" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EntitlementID respjson.Field
		Revoked       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeleteEntitlementResponse) RawJSON() string { return r.JSON.raw }
func (r *DeleteEntitlementResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntitlementAgreementResource struct {
	AgreementID string `json:"agreement_id" api:"required"`
	// Stable entitlement agreement family key.
	//
	// Any of "omni_account_data_access".
	AgreementKey     EntitlementAgreementKey `json:"agreement_key" api:"required"`
	DocumentContent  string                  `json:"document_content" api:"required"`
	DocumentSha256   string                  `json:"document_sha256" api:"required"`
	EntitlementCodes []EntitlementCode       `json:"entitlement_codes" api:"required"`
	Title            string                  `json:"title" api:"required"`
	Version          int64                   `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgreementID      respjson.Field
		AgreementKey     respjson.Field
		DocumentContent  respjson.Field
		DocumentSha256   respjson.Field
		EntitlementCodes respjson.Field
		Title            respjson.Field
		Version          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntitlementAgreementResource) RawJSON() string { return r.JSON.raw }
func (r *EntitlementAgreementResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntitlementAgreementResourceList []EntitlementAgreementResource

type EntitlementResource struct {
	AccountID   int64  `json:"account_id" api:"required"`
	AgreementID string `json:"agreement_id" api:"required"`
	// Stable entitlement code granted by an agreement.
	//
	// Any of "omni.account_data".
	EntitlementCode EntitlementCode `json:"entitlement_code" api:"required"`
	EntitlementID   string          `json:"entitlement_id" api:"required"`
	GrantedAt       string          `json:"granted_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID       respjson.Field
		AgreementID     respjson.Field
		EntitlementCode respjson.Field
		EntitlementID   respjson.Field
		GrantedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntitlementResource) RawJSON() string { return r.JSON.raw }
func (r *EntitlementResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntitlementResourceList []EntitlementResource

// Whether an entitlement is enabled for one account. Accounts the user never chose
// for are enabled by default.
type EntitlementStatusResource struct {
	AccountID int64 `json:"account_id" api:"required"`
	Enabled   bool  `json:"enabled" api:"required"`
	// Stable entitlement code granted by an agreement.
	//
	// Any of "omni.account_data".
	EntitlementCode EntitlementCode `json:"entitlement_code" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID       respjson.Field
		Enabled         respjson.Field
		EntitlementCode respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntitlementStatusResource) RawJSON() string { return r.JSON.raw }
func (r *EntitlementStatusResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntitlementStatusResourceList []EntitlementStatusResource

type V1OmniAIEntitlementNewEntitlementsResponse struct {
	Data EntitlementResourceList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1OmniAIEntitlementNewEntitlementsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIEntitlementNewEntitlementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIEntitlementDeleteEntitlementResponse struct {
	Data DeleteEntitlementResponse `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1OmniAIEntitlementDeleteEntitlementResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIEntitlementDeleteEntitlementResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIEntitlementGetEntitlementAgreementsResponse struct {
	Data EntitlementAgreementResourceList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1OmniAIEntitlementGetEntitlementAgreementsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIEntitlementGetEntitlementAgreementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIEntitlementGetEntitlementStatusResponse struct {
	Data EntitlementStatusResourceList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1OmniAIEntitlementGetEntitlementStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIEntitlementGetEntitlementStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIEntitlementGetEntitlementsResponse struct {
	Data EntitlementResourceList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1OmniAIEntitlementGetEntitlementsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIEntitlementGetEntitlementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIEntitlementOptOutEntitlementsResponse struct {
	Data EntitlementStatusResourceList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1OmniAIEntitlementOptOutEntitlementsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OmniAIEntitlementOptOutEntitlementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIEntitlementNewEntitlementsParams struct {
	AccountIDs       []int64           `json:"account_ids,omitzero" api:"required"`
	AgreementID      string            `json:"agreement_id" api:"required"`
	EntitlementCodes []EntitlementCode `json:"entitlement_codes,omitzero" api:"required"`
	paramObj
}

func (r V1OmniAIEntitlementNewEntitlementsParams) MarshalJSON() (data []byte, err error) {
	type shadow V1OmniAIEntitlementNewEntitlementsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1OmniAIEntitlementNewEntitlementsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OmniAIEntitlementGetEntitlementStatusParams struct {
	AccountID param.Opt[int64] `query:"account_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1OmniAIEntitlementGetEntitlementStatusParams]'s query
// parameters as `url.Values`.
func (r V1OmniAIEntitlementGetEntitlementStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1OmniAIEntitlementGetEntitlementsParams struct {
	AccountID param.Opt[int64] `query:"account_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1OmniAIEntitlementGetEntitlementsParams]'s query
// parameters as `url.Values`.
func (r V1OmniAIEntitlementGetEntitlementsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1OmniAIEntitlementOptOutEntitlementsParams struct {
	AccountIDs       []int64           `json:"account_ids,omitzero" api:"required"`
	EntitlementCodes []EntitlementCode `json:"entitlement_codes,omitzero" api:"required"`
	paramObj
}

func (r V1OmniAIEntitlementOptOutEntitlementsParams) MarshalJSON() (data []byte, err error) {
	type shadow V1OmniAIEntitlementOptOutEntitlementsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1OmniAIEntitlementOptOutEntitlementsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
