// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomusesapientgosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/usesapient/go-sdk/internal/apijson"
	"github.com/usesapient/go-sdk/internal/apiquery"
	"github.com/usesapient/go-sdk/internal/requestconfig"
	"github.com/usesapient/go-sdk/option"
	"github.com/usesapient/go-sdk/packages/param"
	"github.com/usesapient/go-sdk/packages/respjson"
)

// PromptService contains methods and other services that help with interacting
// with the Sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPromptService] method instead.
type PromptService struct {
	Options   []option.RequestOption
	Topics    PromptTopicService
	Platforms PromptPlatformService
}

// NewPromptService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPromptService(opts ...option.RequestOption) (r PromptService) {
	r = PromptService{}
	r.Options = opts
	r.Topics = NewPromptTopicService(opts...)
	r.Platforms = NewPromptPlatformService(opts...)
	return
}

// Create Prompt
func (r *PromptService) New(ctx context.Context, body PromptNewParams, opts ...option.RequestOption) (res *PromptNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/prompts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve Prompt
func (r *PromptService) Get(ctx context.Context, promptID string, opts ...option.RequestOption) (res *PromptGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if promptID == "" {
		err = errors.New("missing required prompt_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/prompts/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update Prompt
func (r *PromptService) Update(ctx context.Context, promptID string, body PromptUpdateParams, opts ...option.RequestOption) (res *PromptUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if promptID == "" {
		err = errors.New("missing required prompt_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/prompts/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Prompts
func (r *PromptService) List(ctx context.Context, query PromptListParams, opts ...option.RequestOption) (res *PromptListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/prompts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete Prompt
func (r *PromptService) Delete(ctx context.Context, promptID string, opts ...option.RequestOption) (res *PromptDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if promptID == "" {
		err = errors.New("missing required prompt_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/prompts/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type PromptNewResponse struct {
	ID         string                      `json:"id" api:"required"`
	Text       string                      `json:"text" api:"required"`
	TopicID    string                      `json:"topic_id" api:"required"`
	CreatedAt  string                      `json:"created_at" api:"nullable"`
	IsActive   bool                        `json:"is_active"`
	LanguageID string                      `json:"language_id" api:"nullable"`
	Platforms  []PromptNewResponsePlatform `json:"platforms"`
	RegionID   string                      `json:"region_id" api:"nullable"`
	// Any of "active", "inactive", "archived".
	Status    PromptNewResponseStatus `json:"status" api:"nullable"`
	Tags      []PromptNewResponseTag  `json:"tags"`
	Topic     PromptNewResponseTopic  `json:"topic" api:"nullable"`
	UpdatedAt string                  `json:"updated_at" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Text        respjson.Field
		TopicID     respjson.Field
		CreatedAt   respjson.Field
		IsActive    respjson.Field
		LanguageID  respjson.Field
		Platforms   respjson.Field
		RegionID    respjson.Field
		Status      respjson.Field
		Tags        respjson.Field
		Topic       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PromptNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptNewResponsePlatform struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptNewResponsePlatform) RawJSON() string { return r.JSON.raw }
func (r *PromptNewResponsePlatform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptNewResponseStatus string

const (
	PromptNewResponseStatusActive   PromptNewResponseStatus = "active"
	PromptNewResponseStatusInactive PromptNewResponseStatus = "inactive"
	PromptNewResponseStatusArchived PromptNewResponseStatus = "archived"
)

type PromptNewResponseTag struct {
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
func (r PromptNewResponseTag) RawJSON() string { return r.JSON.raw }
func (r *PromptNewResponseTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptNewResponseTopic struct {
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
func (r PromptNewResponseTopic) RawJSON() string { return r.JSON.raw }
func (r *PromptNewResponseTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptGetResponse struct {
	ID         string                      `json:"id" api:"required"`
	Text       string                      `json:"text" api:"required"`
	TopicID    string                      `json:"topic_id" api:"required"`
	CreatedAt  string                      `json:"created_at" api:"nullable"`
	IsActive   bool                        `json:"is_active"`
	LanguageID string                      `json:"language_id" api:"nullable"`
	Platforms  []PromptGetResponsePlatform `json:"platforms"`
	RegionID   string                      `json:"region_id" api:"nullable"`
	// Any of "active", "inactive", "archived".
	Status    PromptGetResponseStatus `json:"status" api:"nullable"`
	Tags      []PromptGetResponseTag  `json:"tags"`
	Topic     PromptGetResponseTopic  `json:"topic" api:"nullable"`
	UpdatedAt string                  `json:"updated_at" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Text        respjson.Field
		TopicID     respjson.Field
		CreatedAt   respjson.Field
		IsActive    respjson.Field
		LanguageID  respjson.Field
		Platforms   respjson.Field
		RegionID    respjson.Field
		Status      respjson.Field
		Tags        respjson.Field
		Topic       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PromptGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptGetResponsePlatform struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptGetResponsePlatform) RawJSON() string { return r.JSON.raw }
func (r *PromptGetResponsePlatform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptGetResponseStatus string

const (
	PromptGetResponseStatusActive   PromptGetResponseStatus = "active"
	PromptGetResponseStatusInactive PromptGetResponseStatus = "inactive"
	PromptGetResponseStatusArchived PromptGetResponseStatus = "archived"
)

type PromptGetResponseTag struct {
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
func (r PromptGetResponseTag) RawJSON() string { return r.JSON.raw }
func (r *PromptGetResponseTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptGetResponseTopic struct {
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
func (r PromptGetResponseTopic) RawJSON() string { return r.JSON.raw }
func (r *PromptGetResponseTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptUpdateResponse struct {
	ID         string                         `json:"id" api:"required"`
	Text       string                         `json:"text" api:"required"`
	TopicID    string                         `json:"topic_id" api:"required"`
	CreatedAt  string                         `json:"created_at" api:"nullable"`
	IsActive   bool                           `json:"is_active"`
	LanguageID string                         `json:"language_id" api:"nullable"`
	Platforms  []PromptUpdateResponsePlatform `json:"platforms"`
	RegionID   string                         `json:"region_id" api:"nullable"`
	// Any of "active", "inactive", "archived".
	Status    PromptUpdateResponseStatus `json:"status" api:"nullable"`
	Tags      []PromptUpdateResponseTag  `json:"tags"`
	Topic     PromptUpdateResponseTopic  `json:"topic" api:"nullable"`
	UpdatedAt string                     `json:"updated_at" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Text        respjson.Field
		TopicID     respjson.Field
		CreatedAt   respjson.Field
		IsActive    respjson.Field
		LanguageID  respjson.Field
		Platforms   respjson.Field
		RegionID    respjson.Field
		Status      respjson.Field
		Tags        respjson.Field
		Topic       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PromptUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptUpdateResponsePlatform struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptUpdateResponsePlatform) RawJSON() string { return r.JSON.raw }
func (r *PromptUpdateResponsePlatform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptUpdateResponseStatus string

const (
	PromptUpdateResponseStatusActive   PromptUpdateResponseStatus = "active"
	PromptUpdateResponseStatusInactive PromptUpdateResponseStatus = "inactive"
	PromptUpdateResponseStatusArchived PromptUpdateResponseStatus = "archived"
)

type PromptUpdateResponseTag struct {
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
func (r PromptUpdateResponseTag) RawJSON() string { return r.JSON.raw }
func (r *PromptUpdateResponseTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptUpdateResponseTopic struct {
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
func (r PromptUpdateResponseTopic) RawJSON() string { return r.JSON.raw }
func (r *PromptUpdateResponseTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptListResponse struct {
	Count int64                    `json:"count" api:"required"`
	Data  []PromptListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptListResponse) RawJSON() string { return r.JSON.raw }
func (r *PromptListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptListResponseData struct {
	ID         string                           `json:"id" api:"required"`
	Text       string                           `json:"text" api:"required"`
	TopicID    string                           `json:"topic_id" api:"required"`
	CreatedAt  string                           `json:"created_at" api:"nullable"`
	IsActive   bool                             `json:"is_active"`
	LanguageID string                           `json:"language_id" api:"nullable"`
	Platforms  []PromptListResponseDataPlatform `json:"platforms"`
	RegionID   string                           `json:"region_id" api:"nullable"`
	// Any of "active", "inactive", "archived".
	Status    string                      `json:"status" api:"nullable"`
	Tags      []PromptListResponseDataTag `json:"tags"`
	Topic     PromptListResponseDataTopic `json:"topic" api:"nullable"`
	UpdatedAt string                      `json:"updated_at" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Text        respjson.Field
		TopicID     respjson.Field
		CreatedAt   respjson.Field
		IsActive    respjson.Field
		LanguageID  respjson.Field
		Platforms   respjson.Field
		RegionID    respjson.Field
		Status      respjson.Field
		Tags        respjson.Field
		Topic       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PromptListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptListResponseDataPlatform struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptListResponseDataPlatform) RawJSON() string { return r.JSON.raw }
func (r *PromptListResponseDataPlatform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptListResponseDataTag struct {
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
func (r PromptListResponseDataTag) RawJSON() string { return r.JSON.raw }
func (r *PromptListResponseDataTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptListResponseDataTopic struct {
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
func (r PromptListResponseDataTopic) RawJSON() string { return r.JSON.raw }
func (r *PromptListResponseDataTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptDeleteResponse = any

type PromptNewParams struct {
	Text        string            `json:"text" api:"required"`
	TopicID     string            `json:"topic_id" api:"required"`
	LanguageID  param.Opt[string] `json:"language_id,omitzero"`
	RegionID    param.Opt[string] `json:"region_id,omitzero"`
	PlatformIDs []string          `json:"platform_ids,omitzero"`
	paramObj
}

func (r PromptNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PromptNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PromptNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptUpdateParams struct {
	IsActive    param.Opt[bool]   `json:"is_active,omitzero"`
	LanguageID  param.Opt[string] `json:"language_id,omitzero"`
	RegionID    param.Opt[string] `json:"region_id,omitzero"`
	Text        param.Opt[string] `json:"text,omitzero"`
	TopicID     param.Opt[string] `json:"topic_id,omitzero"`
	PlatformIDs []string          `json:"platform_ids,omitzero"`
	paramObj
}

func (r PromptUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PromptUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PromptUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptListParams struct {
	IsActive        param.Opt[bool]   `query:"is_active,omitzero" json:"-"`
	Status          param.Opt[string] `query:"status,omitzero" json:"-"`
	TopicID         param.Opt[string] `query:"topic_id,omitzero" json:"-"`
	IncludeArchived param.Opt[bool]   `query:"include_archived,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PromptListParams]'s query parameters as `url.Values`.
func (r PromptListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
