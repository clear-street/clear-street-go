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

// Query locate requests, granted locates, and cached locate borrow rates.
//
// V1LocateService contains methods and other services that help with interacting
// with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1LocateService] method instead.
type V1LocateService struct {
	options []option.RequestOption
}

// NewV1LocateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1LocateService(opts ...option.RequestOption) (r V1LocateService) {
	r = V1LocateService{}
	r.options = opts
	return
}

// > **Beta** — this endpoint may change before general availability.
//
// Fetch one locate by id, matched against the locate request id, locate id, and
// client request id, across all days in force.
func (r *V1LocateService) GetLocateByID(ctx context.Context, id string, query V1LocateGetLocateByIDParams, opts ...option.RequestOption) (res *V1LocateGetLocateByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%v/locates/%s", query.AccountID, url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// > **Beta** — this endpoint may change before general availability.
//
// Fetch the cached locate borrow rate for one instrument. Rates refresh
// periodically; `404` means no rate is currently cached for the instrument.
func (r *V1LocateService) GetLocateRateByID(ctx context.Context, instrumentID InstrumentIDOrSymbol, opts ...option.RequestOption) (res *V1LocateGetLocateRateByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if instrumentID == "" {
		err = errors.New("missing required instrument_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/locates/rates/%s", url.PathEscape(instrumentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// > **Beta** — this endpoint may change before general availability.
//
// Fetch the cached locate borrow rates for up to 1,000 instruments. The response
// returns one row per requested instrument, in request order. Each row carries
// either the `rate` or an `error` — a `404` when the instrument is unknown or has
// no cached rate. The status is `207` when any row has an error, `200` when every
// instrument resolved to a rate.
func (r *V1LocateService) GetLocateRates(ctx context.Context, query V1LocateGetLocateRatesParams, opts ...option.RequestOption) (res *V1LocateGetLocateRatesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/locates/rates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// > **Beta** — this endpoint may change before general availability.
//
// List locates for the account on one effective date, which defaults to today's
// New York calendar date. Filter by comma-separated statuses and instrument ids or
// symbols.
func (r *V1LocateService) GetLocates(ctx context.Context, accountID int64, query V1LocateGetLocatesParams, opts ...option.RequestOption) (res *V1LocateGetLocatesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("v1/accounts/%v/locates", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A locate request joined with the locate granted for it, if any.
type Locate struct {
	// Account the locate is booked to.
	AccountID int64 `json:"account_id" api:"required"`
	// Caller-assigned id of the locate request.
	ClientRequestID string `json:"client_request_id" api:"required"`
	// New York calendar date the locate is in force for.
	EffectiveDate time.Time `json:"effective_date" api:"required" format:"date"`
	// Instrument id.
	InstrumentID string `json:"instrument_id" api:"required" format:"uuid"`
	// Engine-assigned locate request id.
	LocateRequestID string `json:"locate_request_id" api:"required" format:"uuid"`
	// Number of shares requested.
	RequestedQuantity int64 `json:"requested_quantity" api:"required"`
	// Presented status of the locate.
	//
	// Any of "REQUESTED", "ACCEPTED", "REJECTED", "EXPIRED".
	Status LocateStatus `json:"status" api:"required"`
	// Engine-assigned locate id; set once the locate is granted. When a null/undefined
	// value is observed, it indicates it does not apply.
	LocateID string `json:"locate_id" api:"nullable" format:"uuid"`
	// Number of shares granted; set once the locate is granted. When a null/undefined
	// value is observed, it indicates it does not apply.
	Quantity int64 `json:"quantity" api:"nullable"`
	// Client order id the request relates to, if any. When a null/undefined value is
	// observed, it indicates it does not apply.
	RelatedOrderID string `json:"related_order_id" api:"nullable" format:"uuid"`
	// Number of granted shares reserved. When a null/undefined value is observed, it
	// indicates it does not apply.
	ReservedQuantity int64 `json:"reserved_quantity" api:"nullable"`
	// Number of granted shares already used. When a null/undefined value is observed,
	// it indicates it does not apply.
	UsedQuantity int64 `json:"used_quantity" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID         respjson.Field
		ClientRequestID   respjson.Field
		EffectiveDate     respjson.Field
		InstrumentID      respjson.Field
		LocateRequestID   respjson.Field
		RequestedQuantity respjson.Field
		Status            respjson.Field
		LocateID          respjson.Field
		Quantity          respjson.Field
		RelatedOrderID    respjson.Field
		ReservedQuantity  respjson.Field
		UsedQuantity      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Locate) RawJSON() string { return r.JSON.raw }
func (r *Locate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocateList []Locate

// A locate borrow rate for an instrument.
type LocateRate struct {
	// Instrument id the rate applies to.
	InstrumentID string `json:"instrument_id" api:"required" format:"uuid"`
	// Borrow rate, expressed as a decimal (e.g. 0.037012 for 3.7012%).
	Rate string `json:"rate" api:"required"`
	// When the rate was requested from the locates venue.
	AsOf time.Time `json:"as_of" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InstrumentID respjson.Field
		Rate         respjson.Field
		AsOf         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocateRate) RawJSON() string { return r.JSON.raw }
func (r *LocateRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A requested instrument with either its cached borrow rate or an error.
type LocateRateResult struct {
	// The instrument id (UUID) or symbol, echoed exactly as requested, so each result
	// maps back to its request.
	Instrument InstrumentIDOrSymbol `json:"instrument" api:"required"`
	// Why no rate is present — a `404` when the instrument is unknown or has no cached
	// rate. Omitted when `rate` is present.
	Error shared.APIError `json:"error" api:"nullable"`
	// The cached borrow rate. Present when a rate was found; omitted when `error` is
	// present.
	Rate LocateRate `json:"rate" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instrument  respjson.Field
		Error       respjson.Field
		Rate        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocateRateResult) RawJSON() string { return r.JSON.raw }
func (r *LocateRateResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocateRateResultList []LocateRateResult

// The presented status of a locate: the locate-request lifecycle plus `EXPIRED`
// for accepted locates whose effective date has passed.
type LocateStatus string

const (
	LocateStatusRequested LocateStatus = "REQUESTED"
	LocateStatusAccepted  LocateStatus = "ACCEPTED"
	LocateStatusRejected  LocateStatus = "REJECTED"
	LocateStatusExpired   LocateStatus = "EXPIRED"
)

type V1LocateGetLocateByIDResponse struct {
	// A locate request joined with the locate granted for it, if any.
	Data Locate `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1LocateGetLocateByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1LocateGetLocateByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1LocateGetLocateRateByIDResponse struct {
	// A locate borrow rate for an instrument.
	Data LocateRate `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1LocateGetLocateRateByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1LocateGetLocateRateByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1LocateGetLocateRatesResponse struct {
	Data LocateRateResultList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1LocateGetLocateRatesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1LocateGetLocateRatesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1LocateGetLocatesResponse struct {
	Data LocateList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1LocateGetLocatesResponse) RawJSON() string { return r.JSON.raw }
func (r *V1LocateGetLocatesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1LocateGetLocateByIDParams struct {
	AccountID int64 `path:"account_id" api:"required" json:"-"`
	paramObj
}

type V1LocateGetLocateRatesParams struct {
	// Comma-separated instrument IDs (UUID) or symbols; at least one, at most 1,000.
	InstrumentIDs []InstrumentIDOrSymbol `query:"instrument_ids,omitzero" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1LocateGetLocateRatesParams]'s query parameters as
// `url.Values`.
func (r V1LocateGetLocateRatesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1LocateGetLocatesParams struct {
	// Effective date to list; defaults to today's New York calendar date. Explicit
	// dates are used unchanged, including weekends and holidays.
	EffectiveDate param.Opt[time.Time] `query:"effective_date,omitzero" format:"date" json:"-"`
	// Comma-separated instrument IDs (UUID) or symbols.
	InstrumentIDs []InstrumentIDOrSymbol `query:"instrument_ids,omitzero" json:"-"`
	// Comma-separated locate statuses to filter by
	//
	// Any of "REQUESTED", "ACCEPTED", "REJECTED", "EXPIRED".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1LocateGetLocatesParams]'s query parameters as
// `url.Values`.
func (r V1LocateGetLocatesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
