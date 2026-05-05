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

// Retrieve details and lists of tradable instruments.
//
// V1InstrumentOptionService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1InstrumentOptionService] method instead.
type V1InstrumentOptionService struct {
	options []option.RequestOption
}

// NewV1InstrumentOptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1InstrumentOptionService(opts ...option.RequestOption) (r V1InstrumentOptionService) {
	r = V1InstrumentOptionService{}
	r.options = opts
	return
}

// List options contracts.
//
// Returns options contracts for a given underlier with options-specific metadata.
// Exactly one underlier identifier must be provided.
func (r *V1InstrumentOptionService) GetOptionContracts(ctx context.Context, query V1InstrumentOptionGetOptionContractsParams, opts ...option.RequestOption) (res *V1InstrumentOptionGetOptionContractsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/instruments/options/contracts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type V1InstrumentOptionGetOptionContractsResponse struct {
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
func (r V1InstrumentOptionGetOptionContractsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1InstrumentOptionGetOptionContractsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1InstrumentOptionGetOptionContractsParams struct {
	// Filter to contracts expiring on this date (YYYY-MM-DD)
	Expiry   param.Opt[time.Time] `query:"expiry,omitzero" format:"date" json:"-"`
	PageSize param.Opt[int64]     `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Underlier symbol (e.g., AAPL, SPX)
	Underlier param.Opt[string] `query:"underlier,omitzero" json:"-"`
	// OEMS instrument UUID or symbol of the underlying equity/index
	UnderlyingInstrumentID param.Opt[string] `query:"underlying_instrument_id,omitzero" format:"uuid" json:"-"`
	// Filter by contract type: CALL or PUT
	//
	// Any of "CALL", "PUT".
	ContractType ContractType `query:"contract_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1InstrumentOptionGetOptionContractsParams]'s query
// parameters as `url.Values`.
func (r V1InstrumentOptionGetOptionContractsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
