// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet

import (
	"github.com/stainless-sdks/clear-street-go/internal/apijson"
	"github.com/stainless-sdks/clear-street-go/option"
	"github.com/stainless-sdks/clear-street-go/packages/respjson"
)

// ActiveV1IrisService contains methods and other services that help with
// interacting with the clear-street API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActiveV1IrisService] method instead.
type ActiveV1IrisService struct {
	Options  []option.RequestOption
	Feedback ActiveV1IrisFeedbackService
	Runs     ActiveV1IrisRunService
	Threads  ActiveV1IrisThreadService
}

// NewActiveV1IrisService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActiveV1IrisService(opts ...option.RequestOption) (r ActiveV1IrisService) {
	r = ActiveV1IrisService{}
	r.Options = opts
	r.Feedback = NewActiveV1IrisFeedbackService(opts...)
	r.Runs = NewActiveV1IrisRunService(opts...)
	r.Threads = NewActiveV1IrisThreadService(opts...)
	return
}

type Message struct {
	// Denormalized text content for search/display
	ContentText string `json:"content_text,required"`
	CreatedAt   string `json:"created_at,required"`
	// Any of "UNSPECIFIED", "SYSTEM", "USER", "ASSISTANT", "TOOL".
	Role         MessageRole `json:"role,required"`
	Seq          int64       `json:"seq,required"`
	ID           string      `json:"id,nullable" format:"uuid"`
	AuthorUserID string      `json:"author_user_id,nullable"`
	// Parsed content parts (text and structured actions)
	Content  MessageContent `json:"content,nullable"`
	Metadata any            `json:"metadata,nullable"`
	RunID    string         `json:"run_id,nullable" format:"uuid"`
	ThreadID string         `json:"thread_id,nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentText  respjson.Field
		CreatedAt    respjson.Field
		Role         respjson.Field
		Seq          respjson.Field
		ID           respjson.Field
		AuthorUserID respjson.Field
		Content      respjson.Field
		Metadata     respjson.Field
		RunID        respjson.Field
		ThreadID     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Message) RawJSON() string { return r.JSON.raw }
func (r *Message) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Thread struct {
	AccountID   string `json:"account_id,required"`
	CreatedAt   string `json:"created_at,required"`
	Description string `json:"description,required"`
	OwnerUserID string `json:"owner_user_id,required"`
	Title       string `json:"title,required"`
	UpdatedAt   string `json:"updated_at,required"`
	ID          string `json:"id,nullable" format:"uuid"`
	Metadata    any    `json:"metadata,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID   respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		OwnerUserID respjson.Field
		Title       respjson.Field
		UpdatedAt   respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Thread) RawJSON() string { return r.JSON.raw }
func (r *Thread) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
