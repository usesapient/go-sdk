// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sapient

import (
	"context"
	"net/http"
	"slices"

	"github.com/usesapient/go-sdk/internal/apijson"
	"github.com/usesapient/go-sdk/internal/requestconfig"
	"github.com/usesapient/go-sdk/option"
	"github.com/usesapient/go-sdk/packages/param"
	"github.com/usesapient/go-sdk/packages/respjson"
)

// APIPromptService contains methods and other services that help with interacting
// with the sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIPromptService] method instead.
type APIPromptService struct {
	options []option.RequestOption
}

// NewAPIPromptService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIPromptService(opts ...option.RequestOption) (r APIPromptService) {
	r = APIPromptService{}
	r.options = opts
	return
}

// Create Prompt Batch
func (r *APIPromptService) NewBatch(ctx context.Context, body APIPromptNewBatchParams, opts ...option.RequestOption) (res *APIPromptNewBatchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/api/prompts/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create Prompt Batch
func (r *APIPromptService) Batch(ctx context.Context, body APIPromptBatchParams, opts ...option.RequestOption) (res *APIPromptBatchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/api/prompts/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type APIPromptNewBatchResponse struct {
	Count   int64            `json:"count" api:"required"`
	Created []map[string]any `json:"created" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Created     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPromptNewBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPromptNewBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPromptBatchResponse struct {
	Count   int64            `json:"count" api:"required"`
	Created []map[string]any `json:"created" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Created     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPromptBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPromptBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPromptNewBatchParams struct {
	Prompts []APIPromptNewBatchParamsPrompt `json:"prompts,omitzero" api:"required"`
	paramObj
}

func (r APIPromptNewBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow APIPromptNewBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIPromptNewBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties LanguageID, RegionID, Text, TopicID are required.
type APIPromptNewBatchParamsPrompt struct {
	LanguageID  string   `json:"language_id" api:"required"`
	RegionID    string   `json:"region_id" api:"required"`
	Text        string   `json:"text" api:"required"`
	TopicID     string   `json:"topic_id" api:"required"`
	PlatformIDs []string `json:"platform_ids,omitzero"`
	paramObj
}

func (r APIPromptNewBatchParamsPrompt) MarshalJSON() (data []byte, err error) {
	type shadow APIPromptNewBatchParamsPrompt
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIPromptNewBatchParamsPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPromptBatchParams struct {
	Prompts []APIPromptBatchParamsPrompt `json:"prompts,omitzero" api:"required"`
	paramObj
}

func (r APIPromptBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow APIPromptBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIPromptBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties LanguageID, RegionID, Text, TopicID are required.
type APIPromptBatchParamsPrompt struct {
	LanguageID  string   `json:"language_id" api:"required"`
	RegionID    string   `json:"region_id" api:"required"`
	Text        string   `json:"text" api:"required"`
	TopicID     string   `json:"topic_id" api:"required"`
	PlatformIDs []string `json:"platform_ids,omitzero"`
	paramObj
}

func (r APIPromptBatchParamsPrompt) MarshalJSON() (data []byte, err error) {
	type shadow APIPromptBatchParamsPrompt
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIPromptBatchParamsPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
