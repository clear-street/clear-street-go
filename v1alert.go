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

// Create and manage alerts that watch market and portfolio conditions on an
// account and notify when they trigger.
//
// V1AlertService contains methods and other services that help with interacting
// with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1AlertService] method instead.
type V1AlertService struct {
	options []option.RequestOption
}

// NewV1AlertService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1AlertService(opts ...option.RequestOption) (r V1AlertService) {
	r = V1AlertService{}
	r.options = opts
	return
}

// Create an alert that watches a market or portfolio condition on the account and
// notifies when it triggers.
//
// The alert starts evaluating immediately. A `once` alert triggers a single time
// and then completes. Instrument references in the condition accept a ticker
// symbol or an OEMS instrument id; they are stored and returned as instrument ids.
//
// `account_id` is optional: an alert without one may only watch market conditions,
// so a condition that reads account data (an `account.*` signal or a holdings
// scope) is rejected without it.
func (r *V1AlertService) NewAlert(ctx context.Context, body V1AlertNewAlertParams, opts ...option.RequestOption) (res *V1AlertNewAlertResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/alerts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Delete an alert. It stops evaluating and disappears from this API; its trigger
// history is retained server-side.
//
// Only `active` and `paused` alerts can be deleted; `completed` and `expired`
// alerts are immutable history. Repeating a delete reports 404, matching what GET
// shows.
func (r *V1AlertService) DeleteAlert(ctx context.Context, alertID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if alertID == "" {
		err = errors.New("missing required alert_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/alerts/%s", alertID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get one alert by id.
func (r *V1AlertService) GetAlertByID(ctx context.Context, alertID string, opts ...option.RequestOption) (res *V1AlertGetAlertByIDResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if alertID == "" {
		err = errors.New("missing required alert_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/alerts/%s", alertID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the caller's alerts, newest first.
//
// `status` narrows the result to a comma-separated set of statuses; when absent,
// alerts of every status are returned. Deleted alerts are never returned.
func (r *V1AlertService) GetAlerts(ctx context.Context, query V1AlertGetAlertsParams, opts ...option.RequestOption) (res *V1AlertGetAlertsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/alerts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A stored alert: the spec it was created with plus its lifecycle facts.
type Alert struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The boolean condition tree, with instrument references resolved to OEMS
	// instrument ids.
	Condition any    `json:"condition" api:"required"`
	CreatedAt string `json:"created_at" api:"required"`
	// How often an alert's condition is evaluated.
	//
	// Any of "every_1m".
	Schedule Schedule `json:"schedule" api:"required"`
	// Where an alert came from.
	//
	// Any of "api", "ui", "internal", "omni".
	Source AlertSource `json:"source" api:"required"`
	// Lifecycle status of an alert. Soft-deleted alerts are invisible on this API, so
	// there is no `deleted` value.
	//
	// Any of "active", "paused", "completed", "expired".
	Status AlertStatus `json:"status" api:"required"`
	// How an alert triggers. `once` alerts complete after their first trigger.
	//
	// Any of "once".
	Trigger   TriggerMode `json:"trigger" api:"required"`
	AccountID int64       `json:"account_id" api:"nullable"`
	ExpiresAt string      `json:"expires_at" api:"nullable"`
	// The originating natural-language text, for alerts compiled from one.
	OmniText string `json:"omni_text" api:"nullable"`
	// When the alert last triggered; absent if it never has.
	TriggeredAt string `json:"triggered_at" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Condition   respjson.Field
		CreatedAt   respjson.Field
		Schedule    respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		Trigger     respjson.Field
		AccountID   respjson.Field
		ExpiresAt   respjson.Field
		OmniText    respjson.Field
		TriggeredAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Alert) RawJSON() string { return r.JSON.raw }
func (r *Alert) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlertList []Alert

// Where an alert came from.
type AlertSource string

const (
	AlertSourceAPI      AlertSource = "api"
	AlertSourceUi       AlertSource = "ui"
	AlertSourceInternal AlertSource = "internal"
	AlertSourceOmni     AlertSource = "omni"
)

// Lifecycle status of an alert. Soft-deleted alerts are invisible on this API, so
// there is no `deleted` value.
type AlertStatus string

const (
	AlertStatusActive    AlertStatus = "active"
	AlertStatusPaused    AlertStatus = "paused"
	AlertStatusCompleted AlertStatus = "completed"
	AlertStatusExpired   AlertStatus = "expired"
)

// Response payload for alert creation.
type CreateAlertResponse struct {
	AlertID string `json:"alert_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AlertID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateAlertResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateAlertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How often an alert's condition is evaluated.
type Schedule string

const (
	ScheduleEvery1m Schedule = "every_1m"
)

// How an alert triggers. `once` alerts complete after their first trigger.
type TriggerMode string

const (
	TriggerModeOnce TriggerMode = "once"
)

type V1AlertNewAlertResponse struct {
	// Response payload for alert creation.
	Data CreateAlertResponse `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AlertNewAlertResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AlertNewAlertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AlertGetAlertByIDResponse struct {
	// A stored alert: the spec it was created with plus its lifecycle facts.
	Data Alert `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AlertGetAlertByIDResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AlertGetAlertByIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AlertGetAlertsResponse struct {
	Data AlertList `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.BaseResponse
}

// Returns the unmodified JSON received from the API
func (r V1AlertGetAlertsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1AlertGetAlertsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AlertNewAlertParams struct {
	// The boolean condition tree, in the condition grammar. `"instrument_id"`
	// references accept a ticker or an OEMS instrument id.
	Condition any `json:"condition,omitzero" api:"required"`
	// How often an alert's condition is evaluated.
	//
	// Any of "every_1m".
	Schedule Schedule `json:"schedule,omitzero" api:"required"`
	// How an alert triggers. `once` alerts complete after their first trigger.
	//
	// Any of "once".
	Trigger TriggerMode `json:"trigger,omitzero" api:"required"`
	// The account whose `account.*` signals and holdings scopes the condition reads.
	// Optional: a market-only alert needs no account.
	AccountID param.Opt[int64] `json:"account_id,omitzero"`
	paramObj
}

func (r V1AlertNewAlertParams) MarshalJSON() (data []byte, err error) {
	type shadow V1AlertNewAlertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1AlertNewAlertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1AlertGetAlertsParams struct {
	// Comma-separated status filter (`active`, `paused`, `completed`, `expired`).
	// Unknown values are rejected. Absent = every status.
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// The number of items to return per page. Only used when page_token is not
	// provided.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for retrieving the next or previous page of results. Contains encoded
	// pagination state; when provided, page_size is ignored.
	PageToken param.Opt[string] `query:"page_token,omitzero" format:"byte" json:"-"`
	paramObj
}

// URLQuery serializes [V1AlertGetAlertsParams]'s query parameters as `url.Values`.
func (r V1AlertGetAlertsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
