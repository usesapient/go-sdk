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

// APIPerformanceOperationPromptService contains methods and other services that
// help with interacting with the Sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIPerformanceOperationPromptService] method instead.
type APIPerformanceOperationPromptService struct {
	Options []option.RequestOption
}

// NewAPIPerformanceOperationPromptService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAPIPerformanceOperationPromptService(opts ...option.RequestOption) (r APIPerformanceOperationPromptService) {
	r = APIPerformanceOperationPromptService{}
	r.Options = opts
	return
}

// Create Operation Prompt
func (r *APIPerformanceOperationPromptService) New(ctx context.Context, operationID string, body APIPerformanceOperationPromptNewParams, opts ...option.RequestOption) (res *APIPerformanceOperationPromptNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api-performance/operations/%s/prompts", operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List Operation Prompts
func (r *APIPerformanceOperationPromptService) List(ctx context.Context, operationID string, opts ...option.RequestOption) (res *[]APIPerformanceOperationPromptListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api-performance/operations/%s/prompts", operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type APIPerformanceOperationPromptNewResponse struct {
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
func (r APIPerformanceOperationPromptNewResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceOperationPromptNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceOperationPromptListResponse struct {
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
func (r APIPerformanceOperationPromptListResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceOperationPromptListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceOperationPromptNewParams struct {
	Prompt           string            `json:"prompt" api:"required"`
	ExpectedBehavior param.Opt[string] `json:"expected_behavior,omitzero"`
	ReferenceAnswer  param.Opt[string] `json:"reference_answer,omitzero"`
	EvalType         param.Opt[string] `json:"eval_type,omitzero"`
	Graders          []map[string]any  `json:"graders,omitzero"`
	paramObj
}

func (r APIPerformanceOperationPromptNewParams) MarshalJSON() (data []byte, err error) {
	type shadow APIPerformanceOperationPromptNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIPerformanceOperationPromptNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
