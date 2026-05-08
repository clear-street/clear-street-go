// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"context"
	"net/http"
	"slices"

	"github.com/clear-street/clear-street-go/internal/requestconfig"
	"github.com/clear-street/clear-street-go/option"
)

// Active Websocket.
//
// V1WebsocketService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1WebsocketService] method instead.
type V1WebsocketService struct {
	options []option.RequestOption
}

// NewV1WebsocketService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1WebsocketService(opts ...option.RequestOption) (r V1WebsocketService) {
	r = V1WebsocketService{}
	r.options = opts
	return
}

// Upgrade the HTTP connection to a WebSocket and echo incoming messages.
func (r *V1WebsocketService) WebsocketHandler(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/ws"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}
