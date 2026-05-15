// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomusesapientgosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/usesapient/go-sdk/internal/apijson"
	"github.com/usesapient/go-sdk/internal/requestconfig"
	"github.com/usesapient/go-sdk/option"
	"github.com/usesapient/go-sdk/packages/param"
	"github.com/usesapient/go-sdk/packages/respjson"
)

// PromptTopicService contains methods and other services that help with
// interacting with the Sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPromptTopicService] method instead.
type PromptTopicService struct {
	Options []option.RequestOption
}

// NewPromptTopicService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPromptTopicService(opts ...option.RequestOption) (r PromptTopicService) {
	r = PromptTopicService{}
	r.Options = opts
	return
}

// Create Topic
func (r *PromptTopicService) New(ctx context.Context, body PromptTopicNewParams, opts ...option.RequestOption) (res *PromptTopicNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/prompts/topics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update Topic
func (r *PromptTopicService) Update(ctx context.Context, topicID string, body PromptTopicUpdateParams, opts ...option.RequestOption) (res *PromptTopicUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if topicID == "" {
		err = errors.New("missing required topic_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/prompts/topics/%s", topicID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Topics
func (r *PromptTopicService) List(ctx context.Context, opts ...option.RequestOption) (res *[]PromptTopicListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/prompts/topics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete Topic
func (r *PromptTopicService) Delete(ctx context.Context, topicID string, opts ...option.RequestOption) (res *PromptTopicDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if topicID == "" {
		err = errors.New("missing required topic_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/prompts/topics/%s", topicID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type PromptTopicNewResponse struct {
	ID        string `json:"id" api:"required"`
	Name      string `json:"name" api:"required"`
	CreatedAt string `json:"created_at" api:"nullable"`
	UpdatedAt string `json:"updated_at" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptTopicNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PromptTopicNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptTopicUpdateResponse struct {
	ID        string `json:"id" api:"required"`
	Name      string `json:"name" api:"required"`
	CreatedAt string `json:"created_at" api:"nullable"`
	UpdatedAt string `json:"updated_at" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptTopicUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PromptTopicUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptTopicListResponse struct {
	ID        string `json:"id" api:"required"`
	Name      string `json:"name" api:"required"`
	CreatedAt string `json:"created_at" api:"nullable"`
	UpdatedAt string `json:"updated_at" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptTopicListResponse) RawJSON() string { return r.JSON.raw }
func (r *PromptTopicListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptTopicDeleteResponse struct {
	Ok bool `json:"ok" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ok          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptTopicDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *PromptTopicDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptTopicNewParams struct {
	Name string `json:"name" api:"required"`
	paramObj
}

func (r PromptTopicNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PromptTopicNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PromptTopicNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptTopicUpdateParams struct {
	Name string `json:"name" api:"required"`
	paramObj
}

func (r PromptTopicUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PromptTopicUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PromptTopicUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
