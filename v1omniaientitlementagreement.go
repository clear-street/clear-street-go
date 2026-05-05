// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"net/http"
	"slices"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Thread-centric AI assistant for conversational trading. Create threads to start
// conversations, poll response objects for in-progress output, and read finalized
// messages from thread history. Thread/message/response endpoints require an
// explicit account_id. Entitlement endpoints are caller-scoped and use
// trading_account_ids.
//
// V1OmniAIEntitlementAgreementService contains methods and other services that
// help with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1OmniAIEntitlementAgreementService] method instead.
type V1OmniAIEntitlementAgreementService struct {
	options []option.RequestOption
}

// NewV1OmniAIEntitlementAgreementService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1OmniAIEntitlementAgreementService(opts ...option.RequestOption) (r V1OmniAIEntitlementAgreementService) {
	r = V1OmniAIEntitlementAgreementService{}
	r.options = opts
	return
}

// List current signable entitlement agreements for consent UX.
func (r *V1OmniAIEntitlementAgreementService) GetEntitlementAgreements(ctx context.Context, opts ...option.RequestOption) (res *V1OmniAIEntitlementAgreementGetEntitlementAgreementsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/omni-ai/entitlement-agreements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type EntitlementAgreementResource struct {
	AgreementID      string   `json:"agreement_id" api:"required"`
	AgreementKey     string   `json:"agreement_key" api:"required"`
	DocumentContent  string   `json:"document_content" api:"required"`
	DocumentSha256   string   `json:"document_sha256" api:"required"`
	EntitlementCodes []string `json:"entitlement_codes" api:"required"`
	Title            string   `json:"title" api:"required"`
	Version          int64    `json:"version" api:"required"`
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

type V1OmniAIEntitlementAgreementGetEntitlementAgreementsResponse struct {
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
func (r V1OmniAIEntitlementAgreementGetEntitlementAgreementsResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *V1OmniAIEntitlementAgreementGetEntitlementAgreementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
