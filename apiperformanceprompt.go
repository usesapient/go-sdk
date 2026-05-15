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

// APIPerformancePromptService contains methods and other services that help with
// interacting with the Sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIPerformancePromptService] method instead.
type APIPerformancePromptService struct {
	Options []option.RequestOption
}

// NewAPIPerformancePromptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIPerformancePromptService(opts ...option.RequestOption) (r APIPerformancePromptService) {
	r = APIPerformancePromptService{}
	r.Options = opts
	return
}

// Retrieve Operation Prompt
func (r *APIPerformancePromptService) Get(ctx context.Context, promptID string, opts ...option.RequestOption) (res *APIPerformancePromptGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if promptID == "" {
		err = errors.New("missing required prompt_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api-performance/prompts/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update Operation Prompt
func (r *APIPerformancePromptService) Update(ctx context.Context, promptID string, body APIPerformancePromptUpdateParams, opts ...option.RequestOption) (res *APIPerformancePromptUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if promptID == "" {
		err = errors.New("missing required prompt_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api-performance/prompts/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete Operation Prompt
func (r *APIPerformancePromptService) Delete(ctx context.Context, promptID string, opts ...option.RequestOption) (res *APIPerformancePromptDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if promptID == "" {
		err = errors.New("missing required prompt_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api-performance/prompts/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type APIPerformancePromptGetResponse struct {
	ID               string           `json:"id" api:"required"`
	EvalType         string           `json:"eval_type" api:"required"`
	OperationID      string           `json:"operation_id" api:"required"`
	Prompt           string           `json:"prompt" api:"required"`
	CreatedAt        string           `json:"created_at" api:"nullable"`
	Enabled          bool             `json:"enabled"`
	ExpectedBehavior string           `json:"expected_behavior" api:"nullable"`
	Graders          []map[string]any `json:"graders"`
	ReferenceAnswer  string           `json:"reference_answer" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		EvalType         respjson.Field
		OperationID      respjson.Field
		Prompt           respjson.Field
		CreatedAt        respjson.Field
		Enabled          respjson.Field
		ExpectedBehavior respjson.Field
		Graders          respjson.Field
		ReferenceAnswer  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformancePromptGetResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformancePromptGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformancePromptUpdateResponse struct {
	ID               string           `json:"id" api:"required"`
	EvalType         string           `json:"eval_type" api:"required"`
	OperationID      string           `json:"operation_id" api:"required"`
	Prompt           string           `json:"prompt" api:"required"`
	CreatedAt        string           `json:"created_at" api:"nullable"`
	Enabled          bool             `json:"enabled"`
	ExpectedBehavior string           `json:"expected_behavior" api:"nullable"`
	Graders          []map[string]any `json:"graders"`
	ReferenceAnswer  string           `json:"reference_answer" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		EvalType         respjson.Field
		OperationID      respjson.Field
		Prompt           respjson.Field
		CreatedAt        respjson.Field
		Enabled          respjson.Field
		ExpectedBehavior respjson.Field
		Graders          respjson.Field
		ReferenceAnswer  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformancePromptUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformancePromptUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformancePromptDeleteResponse struct {
	Ok bool `json:"ok" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ok          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformancePromptDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformancePromptDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformancePromptUpdateParams struct {
	Enabled          param.Opt[bool]   `json:"enabled,omitzero"`
	ExpectedBehavior param.Opt[string] `json:"expected_behavior,omitzero"`
	Prompt           param.Opt[string] `json:"prompt,omitzero"`
	ReferenceAnswer  param.Opt[string] `json:"reference_answer,omitzero"`
	Graders          []map[string]any  `json:"graders,omitzero"`
	paramObj
}

func (r APIPerformancePromptUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow APIPerformancePromptUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIPerformancePromptUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
