// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomusesapientgosdk

import (
	"context"
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

// APIPerformanceEvaluationConfigService contains methods and other services that
// help with interacting with the Sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIPerformanceEvaluationConfigService] method instead.
type APIPerformanceEvaluationConfigService struct {
	Options []option.RequestOption
}

// NewAPIPerformanceEvaluationConfigService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAPIPerformanceEvaluationConfigService(opts ...option.RequestOption) (r APIPerformanceEvaluationConfigService) {
	r = APIPerformanceEvaluationConfigService{}
	r.Options = opts
	return
}

// Retrieve Evaluation Config
func (r *APIPerformanceEvaluationConfigService) Get(ctx context.Context, query APIPerformanceEvaluationConfigGetParams, opts ...option.RequestOption) (res *APIPerformanceEvaluationConfigGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api-performance/evaluation-config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Update Evaluation Config
func (r *APIPerformanceEvaluationConfigService) Update(ctx context.Context, body APIPerformanceEvaluationConfigUpdateParams, opts ...option.RequestOption) (res *APIPerformanceEvaluationConfigUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api-performance/evaluation-config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type APIPerformanceEvaluationConfigGetResponse struct {
	OperationIDs []string `json:"operation_ids" api:"required"`
	APIBaseURL   string   `json:"api_base_url" api:"nullable"`
	EnvVarKeys   []string `json:"env_var_keys"`
	EvalTypes    []string `json:"eval_types"`
	Framework    string   `json:"framework" api:"nullable"`
	Platforms    []string `json:"platforms"`
	UseCaseIDs   []string `json:"use_case_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OperationIDs respjson.Field
		APIBaseURL   respjson.Field
		EnvVarKeys   respjson.Field
		EvalTypes    respjson.Field
		Framework    respjson.Field
		Platforms    respjson.Field
		UseCaseIDs   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceEvaluationConfigGetResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceEvaluationConfigGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceEvaluationConfigUpdateResponse struct {
	OperationIDs []string `json:"operation_ids" api:"required"`
	APIBaseURL   string   `json:"api_base_url" api:"nullable"`
	EnvVarKeys   []string `json:"env_var_keys"`
	EvalTypes    []string `json:"eval_types"`
	Framework    string   `json:"framework" api:"nullable"`
	Platforms    []string `json:"platforms"`
	UseCaseIDs   []string `json:"use_case_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OperationIDs respjson.Field
		APIBaseURL   respjson.Field
		EnvVarKeys   respjson.Field
		EvalTypes    respjson.Field
		Framework    respjson.Field
		Platforms    respjson.Field
		UseCaseIDs   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceEvaluationConfigUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceEvaluationConfigUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceEvaluationConfigGetParams struct {
	InterfaceID param.Opt[string] `query:"interface_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIPerformanceEvaluationConfigGetParams]'s query parameters
// as `url.Values`.
func (r APIPerformanceEvaluationConfigGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIPerformanceEvaluationConfigUpdateParams struct {
	APIBaseURL   param.Opt[string] `json:"api_base_url,omitzero"`
	Framework    param.Opt[string] `json:"framework,omitzero"`
	InterfaceID  param.Opt[string] `json:"interface_id,omitzero"`
	EnvVars      map[string]string `json:"env_vars,omitzero"`
	EvalTypes    []string          `json:"eval_types,omitzero"`
	OperationIDs []string          `json:"operation_ids,omitzero"`
	Platforms    []string          `json:"platforms,omitzero"`
	UseCaseIDs   []string          `json:"use_case_ids,omitzero"`
	paramObj
}

func (r APIPerformanceEvaluationConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow APIPerformanceEvaluationConfigUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIPerformanceEvaluationConfigUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
