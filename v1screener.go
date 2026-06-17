// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/clear-street/clear-street-go/internal/apijson"
	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
	"github.com/clear-street/clear-street-go/packages/param"
	"github.com/clear-street/clear-street-go/packages/respjson"
	"github.com/clear-street/clear-street-go/shared"
)

// Search instruments and manage saved screeners.
//
// V1ScreenerService contains methods and other services that help with interacting
// with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ScreenerService] method instead.
type V1ScreenerService struct {
	options []option.RequestOption
}

// NewV1ScreenerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1ScreenerService(opts ...option.RequestOption) (r V1ScreenerService) {
	r = V1ScreenerService{}
	r.options = opts
	return
}

// Create a saved screener configuration.
//
// Persists a screener configuration for the authenticated user.
func (r *V1ScreenerService) NewScreener(ctx context.Context, body V1ScreenerNewScreenerParams, opts ...option.RequestOption) (res *V1ScreenerNewScreenerResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/saved-screeners"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete a saved screener configuration.
//
// Deletes the screener configuration for the authenticated user.
func (r *V1ScreenerService) DeleteScreener(ctx context.Context, screenerID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if screenerID == "" {
		err = errors.New("missing required screener_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/saved-screeners/%s", screenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get a saved screener configuration by ID.
//
// Returns a single screener configuration for the authenticated user.
func (r *V1ScreenerService) GetScreenerByID(ctx context.Context, screenerID string, opts ...option.RequestOption) (res *V1ScreenerGetScreenerByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if screenerID == "" {
		err = errors.New("missing required screener_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/saved-screeners/%s", screenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List saved screener configurations.
//
// Returns all screener configurations for the authenticated user.
func (r *V1ScreenerService) GetScreeners(ctx context.Context, opts ...option.RequestOption) (res *V1ScreenerGetScreenersResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/saved-screeners"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a saved screener configuration.
//
// Replaces the screener configuration for the authenticated user. If `name` is
// null, the existing name is preserved.
func (r *V1ScreenerService) ReplaceScreener(ctx context.Context, screenerID string, body V1ScreenerReplaceScreenerParams, opts ...option.RequestOption) (res *V1ScreenerReplaceScreenerResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if screenerID == "" {
		err = errors.New("missing required screener_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/saved-screeners/%s", screenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Search instruments using structured filters.
//
// Returns a columnar response where each row is an array of column objects. Each
// column contains a human-readable name, a field reference, an optional type hint
// (e.g. `CURR_USD`, `PERCENT`), and the value.
//
// Use `columns` to select which columns appear in each row. When omitted, the
// default field set is returned.
func (r *V1ScreenerService) SearchScreener(ctx context.Context, body V1ScreenerSearchScreenerParams, opts ...option.RequestOption) (res *V1ScreenerSearchScreenerResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/screener"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Historical lookback window for price/change fields.
type FieldLookback string

const (
	FieldLookbackOneDay      FieldLookback = "ONE_DAY"
	FieldLookbackOneWeek     FieldLookback = "ONE_WEEK"
	FieldLookbackOneMonth    FieldLookback = "ONE_MONTH"
	FieldLookbackThreeMonths FieldLookback = "THREE_MONTHS"
	FieldLookbackSixMonths   FieldLookback = "SIX_MONTHS"
	FieldLookbackYearToDate  FieldLookback = "YEAR_TO_DATE"
	FieldLookbackOneYear     FieldLookback = "ONE_YEAR"
)

// Reporting period for financial data fields.
type FieldPeriod string

const (
	FieldPeriodQuarter              FieldPeriod = "QUARTER"
	FieldPeriodTrailingTwelveMonths FieldPeriod = "TRAILING_TWELVE_MONTHS"
)

// A reference to a screener field.
type FieldRef struct {
	// The field name.
	Name string `json:"name" api:"required"`
	// Optional historical lookback window.
	//
	// Any of "ONE_DAY", "ONE_WEEK", "ONE_MONTH", "THREE_MONTHS", "SIX_MONTHS",
	// "YEAR_TO_DATE", "ONE_YEAR".
	Lookback FieldLookback `json:"lookback" api:"nullable"`
	// Optional reporting period (e.g. quarter or TTM).
	//
	// Any of "QUARTER", "TRAILING_TWELVE_MONTHS".
	Period FieldPeriod `json:"period" api:"nullable"`
	// The data type of the field value. Present only in responses.
	//
	// Any of "DECIMAL", "INTEGER", "STRING", "ANALYST_RATING", "DATE".
	ValueType FieldType `json:"value_type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Lookback    respjson.Field
		Period      respjson.Field
		ValueType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FieldRef) RawJSON() string { return r.JSON.raw }
func (r *FieldRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FieldRef to a FieldRefParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FieldRefParam.Overrides()
func (r FieldRef) ToParam() FieldRefParam {
	return param.Override[FieldRefParam](json.RawMessage(r.RawJSON()))
}

// A reference to a screener field.
//
// The property Name is required.
type FieldRefParam struct {
	// The field name.
	Name string `json:"name" api:"required"`
	// Optional historical lookback window.
	//
	// Any of "ONE_DAY", "ONE_WEEK", "ONE_MONTH", "THREE_MONTHS", "SIX_MONTHS",
	// "YEAR_TO_DATE", "ONE_YEAR".
	Lookback FieldLookback `json:"lookback,omitzero"`
	// Optional reporting period (e.g. quarter or TTM).
	//
	// Any of "QUARTER", "TRAILING_TWELVE_MONTHS".
	Period FieldPeriod `json:"period,omitzero"`
	// The data type of the field value. Present only in responses.
	//
	// Any of "DECIMAL", "INTEGER", "STRING", "ANALYST_RATING", "DATE".
	ValueType FieldType `json:"value_type,omitzero"`
	paramObj
}

func (r FieldRefParam) MarshalJSON() (data []byte, err error) {
	type shadow FieldRefParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FieldRefParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The data type of a screener field value.
type FieldType string

const (
	FieldTypeDecimal       FieldType = "DECIMAL"
	FieldTypeInteger       FieldType = "INTEGER"
	FieldTypeString        FieldType = "STRING"
	FieldTypeAnalystRating FieldType = "ANALYST_RATING"
	FieldTypeDate          FieldType = "DATE"
)

// Operator specification with optional behavioral arguments.
type FilterOpSpec struct {
	// The operator to apply.
	//
	// Any of "LESS_THAN", "LESS_OR_EQUAL", "GREATER_THAN", "GREATER_OR_EQUAL",
	// "EQUAL", "BETWEEN", "NOT_BETWEEN", "ONE_OF", "REGEX", "BEGINS_WITH",
	// "ENDS_WITH", "CONTAINS", "IS_NULL", "IS_NOT_NULL".
	Name FilterOperator `json:"name" api:"required"`
	// Optional arguments that modify operator behavior.
	Args []OperatorArg `json:"args"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Args        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilterOpSpec) RawJSON() string { return r.JSON.raw }
func (r *FilterOpSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FilterOpSpec to a FilterOpSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FilterOpSpecParam.Overrides()
func (r FilterOpSpec) ToParam() FilterOpSpecParam {
	return param.Override[FilterOpSpecParam](json.RawMessage(r.RawJSON()))
}

// Operator specification with optional behavioral arguments.
//
// The property Name is required.
type FilterOpSpecParam struct {
	// The operator to apply.
	//
	// Any of "LESS_THAN", "LESS_OR_EQUAL", "GREATER_THAN", "GREATER_OR_EQUAL",
	// "EQUAL", "BETWEEN", "NOT_BETWEEN", "ONE_OF", "REGEX", "BEGINS_WITH",
	// "ENDS_WITH", "CONTAINS", "IS_NULL", "IS_NOT_NULL".
	Name FilterOperator `json:"name,omitzero" api:"required"`
	// Optional arguments that modify operator behavior.
	Args []OperatorArg `json:"args,omitzero"`
	paramObj
}

func (r FilterOpSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow FilterOpSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilterOpSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter operators supported by the screener.
//
// Abbreviated and lowercase forms are accepted as serde aliases for backward
// compatibility with earlier API revisions; the canonical wire form is the
// SCREAMING_SNAKE_CASE rendering.
type FilterOperator string

const (
	FilterOperatorLessThan       FilterOperator = "LESS_THAN"
	FilterOperatorLessOrEqual    FilterOperator = "LESS_OR_EQUAL"
	FilterOperatorGreaterThan    FilterOperator = "GREATER_THAN"
	FilterOperatorGreaterOrEqual FilterOperator = "GREATER_OR_EQUAL"
	FilterOperatorEqual          FilterOperator = "EQUAL"
	FilterOperatorBetween        FilterOperator = "BETWEEN"
	FilterOperatorNotBetween     FilterOperator = "NOT_BETWEEN"
	FilterOperatorOneOf          FilterOperator = "ONE_OF"
	FilterOperatorRegex          FilterOperator = "REGEX"
	FilterOperatorBeginsWith     FilterOperator = "BEGINS_WITH"
	FilterOperatorEndsWith       FilterOperator = "ENDS_WITH"
	FilterOperatorContains       FilterOperator = "CONTAINS"
	FilterOperatorIsNull         FilterOperator = "IS_NULL"
	FilterOperatorIsNotNull      FilterOperator = "IS_NOT_NULL"
)

// A filter value: either a literal or a variable reference.
type FilterValue struct {
	Value FilterValueValueUnion `json:"value" api:"nullable"`
	// A variable reference.
	Variable Variable `json:"variable" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		Variable    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilterValue) RawJSON() string { return r.JSON.raw }
func (r *FilterValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FilterValue to a FilterValueParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FilterValueParam.Overrides()
func (r FilterValue) ToParam() FilterValueParam {
	return param.Override[FilterValueParam](json.RawMessage(r.RawJSON()))
}

// FilterValueValueUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type FilterValueValueUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u FilterValueValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FilterValueValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FilterValueValueUnion) RawJSON() string { return u.JSON.raw }

func (r *FilterValueValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A filter value: either a literal or a variable reference.
type FilterValueParam struct {
	Value FilterValueValueUnionParam `json:"value,omitzero"`
	// A variable reference.
	Variable VariableParam `json:"variable,omitzero"`
	paramObj
}

func (r FilterValueParam) MarshalJSON() (data []byte, err error) {
	type shadow FilterValueParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilterValueParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FilterValueValueUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u FilterValueValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *FilterValueValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Arithmetic modifier applied to a variable value.
type Modifier struct {
	Args []ModifierArgUnion `json:"args" api:"required"`
	// The modifier operation.
	//
	// Any of "ADD", "SUBTRACT".
	Name ModifierOp `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Modifier) RawJSON() string { return r.JSON.raw }
func (r *Modifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Modifier to a ModifierParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ModifierParam.Overrides()
func (r Modifier) ToParam() ModifierParam {
	return param.Override[ModifierParam](json.RawMessage(r.RawJSON()))
}

// ModifierArgUnion contains all possible properties and values from [float64],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type ModifierArgUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ModifierArgUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModifierArgUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ModifierArgUnion) RawJSON() string { return u.JSON.raw }

func (r *ModifierArgUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Arithmetic modifier applied to a variable value.
//
// The properties Args, Name are required.
type ModifierParam struct {
	Args []ModifierArgUnionParam `json:"args,omitzero" api:"required"`
	// The modifier operation.
	//
	// Any of "ADD", "SUBTRACT".
	Name ModifierOp `json:"name,omitzero" api:"required"`
	paramObj
}

func (r ModifierParam) MarshalJSON() (data []byte, err error) {
	type shadow ModifierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModifierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ModifierArgUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ModifierArgUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ModifierArgUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Modifier operation applied to a variable.
type ModifierOp string

const (
	ModifierOpAdd      ModifierOp = "ADD"
	ModifierOpSubtract ModifierOp = "SUBTRACT"
)

// Argument that modifies operator behavior.
type OperatorArg string

const (
	OperatorArgLeftInclusive   OperatorArg = "LEFT_INCLUSIVE"
	OperatorArgRightInclusive  OperatorArg = "RIGHT_INCLUSIVE"
	OperatorArgLeftExclusive   OperatorArg = "LEFT_EXCLUSIVE"
	OperatorArgRightExclusive  OperatorArg = "RIGHT_EXCLUSIVE"
	OperatorArgCaseInsensitive OperatorArg = "CASE_INSENSITIVE"
)

// A single column in the screener search response.
type ScreenerColumn struct {
	// Field reference (same shape as filter/sort field references)
	Field FieldRef `json:"field" api:"required"`
	// Human-readable display name for this field
	Name  string                   `json:"name" api:"required"`
	Value ScreenerColumnValueUnion `json:"value" api:"required"`
	// Value format hint: "CURR_USD", "PERCENT", etc. Omitted when not applicable. When
	// a null/undefined value is observed, it indicates it does not apply.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Name        respjson.Field
		Value       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScreenerColumn) RawJSON() string { return r.JSON.raw }
func (r *ScreenerColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScreenerColumnValueUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type ScreenerColumnValueUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ScreenerColumnValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScreenerColumnValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScreenerColumnValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ScreenerColumnValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A saved screener configuration entry
type ScreenerEntry struct {
	ID        string         `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time      `json:"created_at" api:"required" format:"date-time"`
	Filters   []SearchFilter `json:"filters" api:"required"`
	Name      string         `json:"name" api:"required"`
	UpdatedAt time.Time      `json:"updated_at" api:"required" format:"date-time"`
	// Field references included when running this screener.
	Columns []FieldRef `json:"columns" api:"nullable"`
	// Deprecated: use `columns` instead. Mirrors `columns`.
	//
	// Deprecated: deprecated
	FieldFilter []FieldRef `json:"field_filter" api:"nullable"`
	Sorts       []SortSpec `json:"sorts" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Filters     respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		Columns     respjson.Field
		FieldFilter respjson.Field
		Sorts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScreenerEntry) RawJSON() string { return r.JSON.raw }
func (r *ScreenerEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScreenerEntryList []ScreenerEntry

// A single filter criterion for the screener.
type ScreenerFilter struct {
	// Field to filter on (e.g., "market_cap", "sector", "price")
	Field string `json:"field" api:"required"`
	// Comparison operator (e.g., "eq", "gte", "lte", "in")
	Operator string `json:"operator" api:"required"`
	// Filter value
	Value any `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScreenerFilter) RawJSON() string { return r.JSON.raw }
func (r *ScreenerFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScreenerRow []ScreenerColumn

type ScreenerRowList []ScreenerRow

// A single filter condition.
//
// When `op` and `right` are both absent, the filter is "unenabled": it persists a
// `left` field reference without applying any predicate. Unenabled filters are
// skipped during search execution but still round-trip through save/load so
// callers can preserve draft state.
type SearchFilter struct {
	// The field to filter on.
	Left FieldRef `json:"left" api:"required"`
	// The operator and optional arguments. Omit together with `right` for an unenabled
	// filter.
	Op FilterOpSpec `json:"op" api:"nullable"`
	// The value(s) to compare against. Omit together with `op` for an unenabled
	// filter.
	Right []FilterValue `json:"right" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Left        respjson.Field
		Op          respjson.Field
		Right       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchFilter) RawJSON() string { return r.JSON.raw }
func (r *SearchFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SearchFilter to a SearchFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SearchFilterParam.Overrides()
func (r SearchFilter) ToParam() SearchFilterParam {
	return param.Override[SearchFilterParam](json.RawMessage(r.RawJSON()))
}

// A single filter condition.
//
// When `op` and `right` are both absent, the filter is "unenabled": it persists a
// `left` field reference without applying any predicate. Unenabled filters are
// skipped during search execution but still round-trip through save/load so
// callers can preserve draft state.
//
// The property Left is required.
type SearchFilterParam struct {
	// The field to filter on.
	Left FieldRefParam `json:"left,omitzero" api:"required"`
	// The value(s) to compare against. Omit together with `op` for an unenabled
	// filter.
	Right []FilterValueParam `json:"right,omitzero"`
	// The operator and optional arguments. Omit together with `right` for an unenabled
	// filter.
	Op FilterOpSpecParam `json:"op,omitzero"`
	paramObj
}

func (r SearchFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow SearchFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sort specification pairing a field with a direction.
type SortSpec struct {
	// The field to sort by.
	Field FieldRef `json:"field" api:"required"`
	// Sort direction (defaults to DESC).
	//
	// Any of "ASC", "DESC".
	Direction SortDirection `json:"direction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Direction   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SortSpec) RawJSON() string { return r.JSON.raw }
func (r *SortSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SortSpec to a SortSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SortSpecParam.Overrides()
func (r SortSpec) ToParam() SortSpecParam {
	return param.Override[SortSpecParam](json.RawMessage(r.RawJSON()))
}

// A sort specification pairing a field with a direction.
//
// The property Field is required.
type SortSpecParam struct {
	// The field to sort by.
	Field FieldRefParam `json:"field,omitzero" api:"required"`
	// Sort direction (defaults to DESC).
	//
	// Any of "ASC", "DESC".
	Direction SortDirection `json:"direction,omitzero"`
	paramObj
}

func (r SortSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow SortSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SortSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A variable reference (field or built-in like `today`).
type Variable struct {
	// The variable name.
	Name string `json:"name" api:"required"`
	// Optional historical lookback window.
	//
	// Any of "ONE_DAY", "ONE_WEEK", "ONE_MONTH", "THREE_MONTHS", "SIX_MONTHS",
	// "YEAR_TO_DATE", "ONE_YEAR".
	Lookback FieldLookback `json:"lookback" api:"nullable"`
	// Optional arithmetic modifier.
	Modifier Modifier `json:"modifier" api:"nullable"`
	// Optional reporting period.
	//
	// Any of "QUARTER", "TRAILING_TWELVE_MONTHS".
	Period FieldPeriod `json:"period" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Lookback    respjson.Field
		Modifier    respjson.Field
		Period      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Variable) RawJSON() string { return r.JSON.raw }
func (r *Variable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Variable to a VariableParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VariableParam.Overrides()
func (r Variable) ToParam() VariableParam {
	return param.Override[VariableParam](json.RawMessage(r.RawJSON()))
}

// A variable reference (field or built-in like `today`).
//
// The property Name is required.
type VariableParam struct {
	// The variable name.
	Name string `json:"name" api:"required"`
	// Optional historical lookback window.
	//
	// Any of "ONE_DAY", "ONE_WEEK", "ONE_MONTH", "THREE_MONTHS", "SIX_MONTHS",
	// "YEAR_TO_DATE", "ONE_YEAR".
	Lookback FieldLookback `json:"lookback,omitzero"`
	// Optional arithmetic modifier.
	Modifier ModifierParam `json:"modifier,omitzero"`
	// Optional reporting period.
	//
	// Any of "QUARTER", "TRAILING_TWELVE_MONTHS".
	Period FieldPeriod `json:"period,omitzero"`
	paramObj
}

func (r VariableParam) MarshalJSON() (data []byte, err error) {
	type shadow VariableParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VariableParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ScreenerNewScreenerResponse struct {
	// A saved screener configuration entry
	Data ScreenerEntry `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1ScreenerNewScreenerResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ScreenerNewScreenerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ScreenerGetScreenerByIDResponse struct {
	// A saved screener configuration entry
	Data ScreenerEntry `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1ScreenerGetScreenerByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ScreenerGetScreenerByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ScreenerGetScreenersResponse struct {
	Data ScreenerEntryList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1ScreenerGetScreenersResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ScreenerGetScreenersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ScreenerReplaceScreenerResponse struct {
	// A saved screener configuration entry
	Data ScreenerEntry `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1ScreenerReplaceScreenerResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ScreenerReplaceScreenerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ScreenerSearchScreenerResponse struct {
	Data ScreenerRowList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1ScreenerSearchScreenerResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ScreenerSearchScreenerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ScreenerNewScreenerParams struct {
	// The name for this screener configuration
	Name param.Opt[string] `json:"name,omitzero"`
	// Structured field references to include when running this screener
	Columns []FieldRefParam `json:"columns,omitzero"`
	// Deprecated: use `columns` instead. Ignored when `columns` is provided.
	FieldFilter []FieldRefParam `json:"field_filter,omitzero"`
	// Structured search filter criteria
	Filters []SearchFilterParam `json:"filters,omitzero"`
	// Multi-field sort specifications
	Sorts []SortSpecParam `json:"sorts,omitzero"`
	paramObj
}

func (r V1ScreenerNewScreenerParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ScreenerNewScreenerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ScreenerNewScreenerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ScreenerReplaceScreenerParams struct {
	// The name for this screener configuration
	Name param.Opt[string] `json:"name,omitzero"`
	// Structured field references to include when running this screener
	Columns []FieldRefParam `json:"columns,omitzero"`
	// Deprecated: use `columns` instead. Ignored when `columns` is provided.
	FieldFilter []FieldRefParam `json:"field_filter,omitzero"`
	// Structured search filter criteria
	Filters []SearchFilterParam `json:"filters,omitzero"`
	// Multi-field sort specifications
	Sorts []SortSpecParam `json:"sorts,omitzero"`
	paramObj
}

func (r V1ScreenerReplaceScreenerParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ScreenerReplaceScreenerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ScreenerReplaceScreenerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ScreenerSearchScreenerParams struct {
	// The number of items to return per page (only used when page_token is not
	// provided)
	PageSize param.Opt[int64] `json:"page_size,omitzero"`
	// Token for retrieving the next page of results. Contains encoded pagination state
	// (limit + offset). When provided, page_size is ignored.
	PageToken param.Opt[string] `json:"page_token,omitzero" format:"byte"`
	// Whether string sorts should be case-sensitive (default: false).
	SortCaseSensitive param.Opt[bool] `json:"sort_case_sensitive,omitzero"`
	// Subset of fields to include in the response.
	Columns []FieldRefParam `json:"columns,omitzero"`
	// Deprecated: use `columns` instead. Ignored when `columns` is provided.
	FieldFilter []FieldRefParam `json:"field_filter,omitzero"`
	// Filter conditions to apply.
	Filters []SearchFilterParam `json:"filters,omitzero"`
	// Multi-field sort specifications.
	Sorts []SortSpecParam `json:"sorts,omitzero"`
	paramObj
}

func (r V1ScreenerSearchScreenerParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ScreenerSearchScreenerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ScreenerSearchScreenerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
