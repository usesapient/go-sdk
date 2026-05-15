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

// APIPerformanceOperationService contains methods and other services that help
// with interacting with the Sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIPerformanceOperationService] method instead.
type APIPerformanceOperationService struct {
	Options []option.RequestOption
}

// NewAPIPerformanceOperationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIPerformanceOperationService(opts ...option.RequestOption) (r APIPerformanceOperationService) {
	r = APIPerformanceOperationService{}
	r.Options = opts
	return
}

// Create Operation
func (r *APIPerformanceOperationService) New(ctx context.Context, body APIPerformanceOperationNewParams, opts ...option.RequestOption) (res *APIPerformanceOperationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api-performance/operations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve Operation
func (r *APIPerformanceOperationService) Get(ctx context.Context, operationID string, opts ...option.RequestOption) (res *APIPerformanceOperationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api-performance/operations/%s", operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update Operation
func (r *APIPerformanceOperationService) Update(ctx context.Context, operationID string, body APIPerformanceOperationUpdateParams, opts ...option.RequestOption) (res *APIPerformanceOperationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api-performance/operations/%s", operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Operations
func (r *APIPerformanceOperationService) List(ctx context.Context, query APIPerformanceOperationListParams, opts ...option.RequestOption) (res *[]APIPerformanceOperationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api-performance/operations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete Operation
func (r *APIPerformanceOperationService) Delete(ctx context.Context, operationID string, opts ...option.RequestOption) (res *APIPerformanceOperationDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api-performance/operations/%s", operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type APIPerformanceOperationNewResponse struct {
	ID                 string                                     `json:"id" api:"required"`
	Method             string                                     `json:"method" api:"required"`
	Path               string                                     `json:"path" api:"required"`
	Category           APIPerformanceOperationNewResponseCategory `json:"category" api:"nullable"`
	CreatedAt          string                                     `json:"created_at" api:"nullable"`
	Description        string                                     `json:"description" api:"nullable"`
	InterfaceID        string                                     `json:"interface_id" api:"nullable"`
	OpenAPIOperationID string                                     `json:"openapi_operation_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Method             respjson.Field
		Path               respjson.Field
		Category           respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		InterfaceID        respjson.Field
		OpenAPIOperationID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceOperationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceOperationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceOperationNewResponseCategory struct {
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
func (r APIPerformanceOperationNewResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceOperationNewResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceOperationGetResponse struct {
	ID                 string                                     `json:"id" api:"required"`
	Method             string                                     `json:"method" api:"required"`
	Path               string                                     `json:"path" api:"required"`
	Category           APIPerformanceOperationGetResponseCategory `json:"category" api:"nullable"`
	CreatedAt          string                                     `json:"created_at" api:"nullable"`
	Description        string                                     `json:"description" api:"nullable"`
	InterfaceID        string                                     `json:"interface_id" api:"nullable"`
	OpenAPIOperationID string                                     `json:"openapi_operation_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Method             respjson.Field
		Path               respjson.Field
		Category           respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		InterfaceID        respjson.Field
		OpenAPIOperationID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceOperationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceOperationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceOperationGetResponseCategory struct {
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
func (r APIPerformanceOperationGetResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceOperationGetResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceOperationUpdateResponse struct {
	ID                 string                                        `json:"id" api:"required"`
	Method             string                                        `json:"method" api:"required"`
	Path               string                                        `json:"path" api:"required"`
	Category           APIPerformanceOperationUpdateResponseCategory `json:"category" api:"nullable"`
	CreatedAt          string                                        `json:"created_at" api:"nullable"`
	Description        string                                        `json:"description" api:"nullable"`
	InterfaceID        string                                        `json:"interface_id" api:"nullable"`
	OpenAPIOperationID string                                        `json:"openapi_operation_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Method             respjson.Field
		Path               respjson.Field
		Category           respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		InterfaceID        respjson.Field
		OpenAPIOperationID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceOperationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceOperationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceOperationUpdateResponseCategory struct {
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
func (r APIPerformanceOperationUpdateResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceOperationUpdateResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceOperationListResponse struct {
	ID                 string                                      `json:"id" api:"required"`
	Method             string                                      `json:"method" api:"required"`
	Path               string                                      `json:"path" api:"required"`
	Category           APIPerformanceOperationListResponseCategory `json:"category" api:"nullable"`
	CreatedAt          string                                      `json:"created_at" api:"nullable"`
	Description        string                                      `json:"description" api:"nullable"`
	InterfaceID        string                                      `json:"interface_id" api:"nullable"`
	OpenAPIOperationID string                                      `json:"openapi_operation_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Method             respjson.Field
		Path               respjson.Field
		Category           respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		InterfaceID        respjson.Field
		OpenAPIOperationID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceOperationListResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceOperationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceOperationListResponseCategory struct {
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
func (r APIPerformanceOperationListResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceOperationListResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceOperationDeleteResponse = any

type APIPerformanceOperationNewParams struct {
	Method             string            `json:"method" api:"required"`
	Path               string            `json:"path" api:"required"`
	CategoryName       param.Opt[string] `json:"category_name,omitzero"`
	CategorySlug       param.Opt[string] `json:"category_slug,omitzero"`
	Description        param.Opt[string] `json:"description,omitzero"`
	InterfaceID        param.Opt[string] `json:"interface_id,omitzero"`
	OpenAPIOperationID param.Opt[string] `json:"openapi_operation_id,omitzero"`
	paramObj
}

func (r APIPerformanceOperationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow APIPerformanceOperationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIPerformanceOperationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceOperationUpdateParams struct {
	CategoryName       param.Opt[string] `json:"category_name,omitzero"`
	CategorySlug       param.Opt[string] `json:"category_slug,omitzero"`
	Description        param.Opt[string] `json:"description,omitzero"`
	InterfaceID        param.Opt[string] `json:"interface_id,omitzero"`
	Method             param.Opt[string] `json:"method,omitzero"`
	OpenAPIOperationID param.Opt[string] `json:"openapi_operation_id,omitzero"`
	Path               param.Opt[string] `json:"path,omitzero"`
	paramObj
}

func (r APIPerformanceOperationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow APIPerformanceOperationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIPerformanceOperationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceOperationListParams struct {
	CategoryID  param.Opt[string] `query:"category_id,omitzero" json:"-"`
	InterfaceID param.Opt[string] `query:"interface_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIPerformanceOperationListParams]'s query parameters as
// `url.Values`.
func (r APIPerformanceOperationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
