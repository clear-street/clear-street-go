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
// ActiveV1InstrumentOptionService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1InstrumentOptionService] method instead.
type ActiveV1InstrumentOptionService struct {
	options []option.RequestOption
}

// NewActiveV1InstrumentOptionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1InstrumentOptionService(opts ...option.RequestOption) (r ActiveV1InstrumentOptionService) {
	r = ActiveV1InstrumentOptionService{}
	r.options = opts
	return
}

// List options contracts.
//
// Returns options contracts for a given underlier with options-specific metadata.
// Exactly one underlier identifier must be provided.
func (r *ActiveV1InstrumentOptionService) Contracts(ctx context.Context, query ActiveV1InstrumentOptionContractsParams, opts ...option.RequestOption) (res *ActiveV1InstrumentOptionContractsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "active/v1/instruments/options/contracts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ActiveV1InstrumentOptionContractsResponse struct {
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
func (r ActiveV1InstrumentOptionContractsResponse) RawJSON() string { return r.JSON.raw }
func (r *ActiveV1InstrumentOptionContractsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1InstrumentOptionContractsParams struct {
	// Filter to contracts expiring on this date (YYYY-MM-DD)
	Expiry   param.Opt[time.Time] `query:"expiry,omitzero" format:"date" json:"-"`
	PageSize param.Opt[int64]     `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	// Underlier symbol (e.g., AAPL, SPX)
	Underlier param.Opt[string] `query:"underlier,omitzero" json:"-"`
	// OEMS instrument UUID of the underlying equity/index
	UnderlierInstrumentID param.Opt[string] `query:"underlier_instrument_id,omitzero" format:"uuid" json:"-"`
	// Security identifier of the underlying (e.g., CUSIP, ISIN). Must be paired with
	// underlier_security_id_source.
	UnderlierSecurityID param.Opt[string] `query:"underlier_security_id,omitzero" json:"-"`
	// Filter by contract type: CALL or PUT
	//
	// Any of "CALL", "PUT".
	ContractType ContractType `query:"contract_type,omitzero" json:"-"`
	// Security ID source for the underlier (e.g., CMS, CUSIP). Must be paired with
	// underlier_security_id.
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "CURRENCY", "FMP", "OEMS",
	// "SEDOL", "QUIK", "ISIN", "RIC", "COUNTRY", "EXCHANGE", "CTA", "BLOOMBERG",
	// "WERTPAPIER", "DUTCH", "VALOREN", "SICOVAM", "BELGIAN", "COMMON",
	// "CLEARING_HOUSE", "ISDA_FPML_SPECIFICATION", "ISDA_FPML_URL",
	// "LETTER_OF_CREDIT", "MARKETPLACE_ASSIGNED_IDENTIFIER", "MARKIT_RED_ENTITY_CLIP",
	// "MARKIT_RED_PAIR_CLIP", "CFTC", "ISDA_COMMODITY_REFERENCE_PRICE",
	// "LEGAL_ENTITY_IDENTIFIER", "SYNTHETIC", "FIDESSA_INSTRUMENT_MNEMONIC",
	// "INDEX_NAME", "UNIFORM_SYMBOL", "DIGITAL_TOKEN_IDENTIFIER", "MASSIVE", "OTHER".
	UnderlierSecurityIDSource SecurityIDSource `query:"underlier_security_id_source,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1InstrumentOptionContractsParams]'s query parameters
// as `url.Values`.
func (r ActiveV1InstrumentOptionContractsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatIndices,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
