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

	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/internal/apiquery"
	"github.com/stainless-sdks/clear-street-go/internal/requestconfig"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
	"github.com/stainless-sdks/clear-street-go/shared"
)

// ActiveV1InstrumentReportingService contains methods and other services that help
// with interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1InstrumentReportingService] method instead.
type ActiveV1InstrumentReportingService struct {
	Options []option.RequestOption
}

// NewActiveV1InstrumentReportingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewActiveV1InstrumentReportingService(opts ...option.RequestOption) (r ActiveV1InstrumentReportingService) {
	r = ActiveV1InstrumentReportingService{}
	r.Options = opts
	return
}

// Retrieves fundamental and financial reporting data for an instrument.
func (r *ActiveV1InstrumentReportingService) GetInstrumentReporting(ctx context.Context, securityID string, params ActiveV1InstrumentReportingGetInstrumentReportingParams, opts ...option.RequestOption) (res *ActiveV1InstrumentReportingGetInstrumentReportingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if securityID == "" {
		err = errors.New("missing required security_id parameter")
		return
	}
	path := fmt.Sprintf("active/v1/instruments/%v/%s/reporting", params.SecurityIDSource, securityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Fiscal period type for earnings reports
type FiscalPeriodType string

const (
	FiscalPeriodTypeQuarterly FiscalPeriodType = "QUARTERLY"
	FiscalPeriodTypeAnnual    FiscalPeriodType = "ANNUAL"
	FiscalPeriodTypeTtm       FiscalPeriodType = "TTM"
)

// Represents instrument earnings data
type InstrumentEarnings struct {
	// The date when the earnings report was published
	Date time.Time `json:"date,required" format:"date"`
	// The fiscal period (e.g., quarter) within the year
	Period int64 `json:"period,required"`
	// The type of fiscal period
	//
	// Any of "QUARTERLY", "ANNUAL", "TTM".
	PeriodType FiscalPeriodType `json:"period_type,required"`
	// The fiscal year of the earnings period
	Year int64 `json:"year,required"`
	// The actual earnings per share (EPS) for the period
	EpsActual string `json:"eps_actual,nullable"`
	// The estimated earnings per share (EPS) for the period
	EpsEstimate string `json:"eps_estimate,nullable"`
	// The percentage difference between actual and estimated EPS
	EpsSurprisePercent string `json:"eps_surprise_percent,nullable"`
	// The actual total revenue for the period
	RevenueActual int64 `json:"revenue_actual,nullable"`
	// The estimated total revenue for the period
	RevenueEstimate int64 `json:"revenue_estimate,nullable"`
	// The percentage difference between actual and estimated revenue
	RevenueSurprisePercent string `json:"revenue_surprise_percent,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date                   respjson.Field
		Period                 respjson.Field
		PeriodType             respjson.Field
		Year                   respjson.Field
		EpsActual              respjson.Field
		EpsEstimate            respjson.Field
		EpsSurprisePercent     respjson.Field
		RevenueActual          respjson.Field
		RevenueEstimate        respjson.Field
		RevenueSurprisePercent respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstrumentEarnings) RawJSON() string { return r.JSON.raw }
func (r *InstrumentEarnings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1InstrumentReportingGetInstrumentReportingResponse struct {
	// Represents instrument earnings data
	Data InstrumentEarnings `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r ActiveV1InstrumentReportingGetInstrumentReportingResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ActiveV1InstrumentReportingGetInstrumentReportingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActiveV1InstrumentReportingGetInstrumentReportingParams struct {
	// Security identifier source
	//
	// Any of "CMS", "CLST", "OPRA", "FIGI", "CUSIP", "OTHER".
	SecurityIDSource SecurityIDSource `path:"security_id_source,omitzero,required" json:"-"`
	// The start date for the query range, inclusive (YYYY-MM-DD)
	FromDate string `query:"from_date,required" json:"-"`
	// The end date for the query range, inclusive (YYYY-MM-DD)
	ToDate string `query:"to_date,required" json:"-"`
	paramObj
}

// URLQuery serializes [ActiveV1InstrumentReportingGetInstrumentReportingParams]'s
// query parameters as `url.Values`.
func (r ActiveV1InstrumentReportingGetInstrumentReportingParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
