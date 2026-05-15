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

// APIPerformanceUseCaseService contains methods and other services that help with
// interacting with the Sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIPerformanceUseCaseService] method instead.
type APIPerformanceUseCaseService struct {
	Options []option.RequestOption
}

// NewAPIPerformanceUseCaseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIPerformanceUseCaseService(opts ...option.RequestOption) (r APIPerformanceUseCaseService) {
	r = APIPerformanceUseCaseService{}
	r.Options = opts
	return
}

// Create Use Case
func (r *APIPerformanceUseCaseService) New(ctx context.Context, body APIPerformanceUseCaseNewParams, opts ...option.RequestOption) (res *APIPerformanceUseCaseNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api-performance/use-cases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve Use Case
func (r *APIPerformanceUseCaseService) Get(ctx context.Context, useCaseID string, opts ...option.RequestOption) (res *APIPerformanceUseCaseGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if useCaseID == "" {
		err = errors.New("missing required use_case_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api-performance/use-cases/%s", useCaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update Use Case
func (r *APIPerformanceUseCaseService) Update(ctx context.Context, useCaseID string, body APIPerformanceUseCaseUpdateParams, opts ...option.RequestOption) (res *APIPerformanceUseCaseUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if useCaseID == "" {
		err = errors.New("missing required use_case_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api-performance/use-cases/%s", useCaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Use Cases
func (r *APIPerformanceUseCaseService) List(ctx context.Context, opts ...option.RequestOption) (res *[]APIPerformanceUseCaseListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api-performance/use-cases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete Use Case
func (r *APIPerformanceUseCaseService) Delete(ctx context.Context, useCaseID string, opts ...option.RequestOption) (res *APIPerformanceUseCaseDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if useCaseID == "" {
		err = errors.New("missing required use_case_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api-performance/use-cases/%s", useCaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type APIPerformanceUseCaseNewResponse struct {
	ID          string                                   `json:"id" api:"required"`
	Prompt      string                                   `json:"prompt" api:"required"`
	Category    APIPerformanceUseCaseNewResponseCategory `json:"category" api:"nullable"`
	CreatedAt   string                                   `json:"created_at" api:"nullable"`
	Description string                                   `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Prompt      respjson.Field
		Category    respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceUseCaseNewResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceUseCaseNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceUseCaseNewResponseCategory struct {
	ID        string `json:"id" api:"required"`
	Name      string `json:"name" api:"required"`
	Slug      string `json:"slug" api:"required"`
	SortOrder int64  `json:"sort_order"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		SortOrder   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceUseCaseNewResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceUseCaseNewResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceUseCaseGetResponse struct {
	ID          string                                   `json:"id" api:"required"`
	Prompt      string                                   `json:"prompt" api:"required"`
	Category    APIPerformanceUseCaseGetResponseCategory `json:"category" api:"nullable"`
	CreatedAt   string                                   `json:"created_at" api:"nullable"`
	Description string                                   `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Prompt      respjson.Field
		Category    respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceUseCaseGetResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceUseCaseGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceUseCaseGetResponseCategory struct {
	ID        string `json:"id" api:"required"`
	Name      string `json:"name" api:"required"`
	Slug      string `json:"slug" api:"required"`
	SortOrder int64  `json:"sort_order"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		SortOrder   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceUseCaseGetResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceUseCaseGetResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceUseCaseUpdateResponse struct {
	ID          string                                      `json:"id" api:"required"`
	Prompt      string                                      `json:"prompt" api:"required"`
	Category    APIPerformanceUseCaseUpdateResponseCategory `json:"category" api:"nullable"`
	CreatedAt   string                                      `json:"created_at" api:"nullable"`
	Description string                                      `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Prompt      respjson.Field
		Category    respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceUseCaseUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceUseCaseUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceUseCaseUpdateResponseCategory struct {
	ID        string `json:"id" api:"required"`
	Name      string `json:"name" api:"required"`
	Slug      string `json:"slug" api:"required"`
	SortOrder int64  `json:"sort_order"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		SortOrder   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceUseCaseUpdateResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceUseCaseUpdateResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceUseCaseListResponse struct {
	ID          string                                    `json:"id" api:"required"`
	Prompt      string                                    `json:"prompt" api:"required"`
	Category    APIPerformanceUseCaseListResponseCategory `json:"category" api:"nullable"`
	CreatedAt   string                                    `json:"created_at" api:"nullable"`
	Description string                                    `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Prompt      respjson.Field
		Category    respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceUseCaseListResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceUseCaseListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceUseCaseListResponseCategory struct {
	ID        string `json:"id" api:"required"`
	Name      string `json:"name" api:"required"`
	Slug      string `json:"slug" api:"required"`
	SortOrder int64  `json:"sort_order"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		SortOrder   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceUseCaseListResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceUseCaseListResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceUseCaseDeleteResponse struct {
	Ok bool `json:"ok" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ok          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceUseCaseDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceUseCaseDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceUseCaseNewParams struct {
	Prompt       string            `json:"prompt" api:"required"`
	CategoryName param.Opt[string] `json:"category_name,omitzero"`
	Description  param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r APIPerformanceUseCaseNewParams) MarshalJSON() (data []byte, err error) {
	type shadow APIPerformanceUseCaseNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIPerformanceUseCaseNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceUseCaseUpdateParams struct {
	CategoryName param.Opt[string] `json:"category_name,omitzero"`
	Description  param.Opt[string] `json:"description,omitzero"`
	Prompt       param.Opt[string] `json:"prompt,omitzero"`
	paramObj
}

func (r APIPerformanceUseCaseUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow APIPerformanceUseCaseUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIPerformanceUseCaseUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
