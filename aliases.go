// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"github.com/clear-street/clear-street-go/internal/apierror"
	"github.com/clear-street/clear-street-go/packages/param"
	"github.com/clear-street/clear-street-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// A direct mapping of tonic::Status, for use in HTTP responses.
//
// This is an alias to an internal type.
type APIError = shared.APIError

// This is an alias to an internal type.
type BaseResponse = shared.BaseResponse

// Metadata for the response. This will always contain a request ID which can be
// used to identify the request to Clear Street for tracing, and optionally may
// include pagination data.
//
// This is an alias to an internal type.
type ResponseMetadata = shared.ResponseMetadata
