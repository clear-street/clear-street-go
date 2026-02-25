// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/packages/param"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// A direct mapping of tonic::Status, for use in HTTP responses.
type APIError struct {
	// The error code is used to identify the nature of the error. It corresponds to an
	// HTTP status code.
	Code int64 `json:"code" api:"required"`
	// A human-readable message providing more details about the error.
	Message string `json:"message" api:"required"`
	// Additional error details, if any. This can include structured information such
	// as field violations or error metadata.
	Details []map[string]any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Details     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIError) RawJSON() string { return r.JSON.raw }
func (r *APIError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseResponse struct {
	// Response metadata, including the request ID and optional pagination info.
	Metadata ResponseMetadata `json:"metadata" api:"required"`
	// Structured error details when the request is unsuccessful.
	Error APIError `json:"error" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseResponse) RawJSON() string { return r.JSON.raw }
func (r *BaseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for the response. This will always contain a request ID which can be
// used to identify the request to Clear Street for tracing, and optionally may
// include pagination data.
type ResponseMetadata struct {
	// A unique ID for this request, generated upon ingestion of the request.
	RequestID string `json:"request_id" api:"required"`
	// Base64URL-encoded pagination token containing limit and offset
	NextPageToken string `json:"next_page_token" api:"nullable" format:"byte"`
	// Pagination. Included if this was a GET (list) response
	PageNumber int64 `json:"page_number" api:"nullable"`
	// Base64URL-encoded pagination token containing limit and offset
	PreviousPageToken string `json:"previous_page_token" api:"nullable" format:"byte"`
	// Total number of items available (not just in this page).
	TotalItems int64 `json:"total_items" api:"nullable"`
	// Total number of pages available.
	TotalPages int64 `json:"total_pages" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID         respjson.Field
		NextPageToken     respjson.Field
		PageNumber        respjson.Field
		PreviousPageToken respjson.Field
		TotalItems        respjson.Field
		TotalPages        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *ResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
