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

// APIPerformanceRunService contains methods and other services that help with
// interacting with the Sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIPerformanceRunService] method instead.
type APIPerformanceRunService struct {
	Options []option.RequestOption
}

// NewAPIPerformanceRunService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIPerformanceRunService(opts ...option.RequestOption) (r APIPerformanceRunService) {
	r = APIPerformanceRunService{}
	r.Options = opts
	return
}

// Retrieve Run
func (r *APIPerformanceRunService) Get(ctx context.Context, runID string, opts ...option.RequestOption) (res *APIPerformanceRunGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api-performance/runs/%s", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Operation Runs
func (r *APIPerformanceRunService) List(ctx context.Context, query APIPerformanceRunListParams, opts ...option.RequestOption) (res *APIPerformanceRunListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api-performance/runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type APIPerformanceRunGetResponse struct {
	Data      APIPerformanceRunGetResponseData      `json:"data" api:"required"`
	Operation APIPerformanceRunGetResponseOperation `json:"operation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Operation   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceRunGetResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceRunGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceRunGetResponseData struct {
	ID                string           `json:"id" api:"required"`
	APICallLog        map[string]any   `json:"api_call_log" api:"nullable"`
	ConversationTurns []map[string]any `json:"conversation_turns" api:"nullable"`
	CostUsd           float64          `json:"cost_usd" api:"nullable"`
	CreatedAt         string           `json:"created_at" api:"nullable"`
	ErrorMessage      string           `json:"error_message" api:"nullable"`
	EvalID            string           `json:"eval_id" api:"nullable"`
	EvalType          string           `json:"eval_type" api:"nullable"`
	ExecutionStderr   string           `json:"execution_stderr" api:"nullable"`
	ExecutionStdout   string           `json:"execution_stdout" api:"nullable"`
	ExitCode          int64            `json:"exit_code" api:"nullable"`
	FailureReasons    []string         `json:"failure_reasons" api:"nullable"`
	GeneratedFiles    map[string]any   `json:"generated_files" api:"nullable"`
	GraderResults     map[string]any   `json:"grader_results" api:"nullable"`
	LatencyMs         int64            `json:"latency_ms" api:"nullable"`
	Model             string           `json:"model" api:"nullable"`
	ModelType         string           `json:"model_type" api:"nullable"`
	Passed            bool             `json:"passed" api:"nullable"`
	Platform          string           `json:"platform" api:"nullable"`
	Prompt            string           `json:"prompt" api:"nullable"`
	RawResponse       string           `json:"raw_response" api:"nullable"`
	RunDate           string           `json:"run_date" api:"nullable"`
	Score             float64          `json:"score" api:"nullable"`
	TokensUsed        int64            `json:"tokens_used" api:"nullable"`
	ToolCalls         []map[string]any `json:"tool_calls" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		APICallLog        respjson.Field
		ConversationTurns respjson.Field
		CostUsd           respjson.Field
		CreatedAt         respjson.Field
		ErrorMessage      respjson.Field
		EvalID            respjson.Field
		EvalType          respjson.Field
		ExecutionStderr   respjson.Field
		ExecutionStdout   respjson.Field
		ExitCode          respjson.Field
		FailureReasons    respjson.Field
		GeneratedFiles    respjson.Field
		GraderResults     respjson.Field
		LatencyMs         respjson.Field
		Model             respjson.Field
		ModelType         respjson.Field
		Passed            respjson.Field
		Platform          respjson.Field
		Prompt            respjson.Field
		RawResponse       respjson.Field
		RunDate           respjson.Field
		Score             respjson.Field
		TokensUsed        respjson.Field
		ToolCalls         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceRunGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceRunGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceRunGetResponseOperation struct {
	ID                 string                                        `json:"id" api:"required"`
	Method             string                                        `json:"method" api:"required"`
	Path               string                                        `json:"path" api:"required"`
	Category           APIPerformanceRunGetResponseOperationCategory `json:"category" api:"nullable"`
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
func (r APIPerformanceRunGetResponseOperation) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceRunGetResponseOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceRunGetResponseOperationCategory struct {
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
func (r APIPerformanceRunGetResponseOperationCategory) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceRunGetResponseOperationCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceRunListResponse struct {
	Data      []APIPerformanceRunListResponseData    `json:"data" api:"required"`
	Meta      APIPerformanceRunListResponseMeta      `json:"meta" api:"required"`
	Operation APIPerformanceRunListResponseOperation `json:"operation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		Operation   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceRunListResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceRunListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceRunListResponseData struct {
	ID             string   `json:"id" api:"required"`
	CostUsd        float64  `json:"cost_usd" api:"nullable"`
	CreatedAt      string   `json:"created_at" api:"nullable"`
	ErrorMessage   string   `json:"error_message" api:"nullable"`
	EvalID         string   `json:"eval_id" api:"nullable"`
	EvalType       string   `json:"eval_type" api:"nullable"`
	FailureReasons []string `json:"failure_reasons" api:"nullable"`
	LatencyMs      int64    `json:"latency_ms" api:"nullable"`
	Model          string   `json:"model" api:"nullable"`
	ModelType      string   `json:"model_type" api:"nullable"`
	Passed         bool     `json:"passed" api:"nullable"`
	Platform       string   `json:"platform" api:"nullable"`
	Prompt         string   `json:"prompt" api:"nullable"`
	RawResponse    string   `json:"raw_response" api:"nullable"`
	RunDate        string   `json:"run_date" api:"nullable"`
	Score          float64  `json:"score" api:"nullable"`
	TokensUsed     int64    `json:"tokens_used" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CostUsd        respjson.Field
		CreatedAt      respjson.Field
		ErrorMessage   respjson.Field
		EvalID         respjson.Field
		EvalType       respjson.Field
		FailureReasons respjson.Field
		LatencyMs      respjson.Field
		Model          respjson.Field
		ModelType      respjson.Field
		Passed         respjson.Field
		Platform       respjson.Field
		Prompt         respjson.Field
		RawResponse    respjson.Field
		RunDate        respjson.Field
		Score          respjson.Field
		TokensUsed     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceRunListResponseData) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceRunListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceRunListResponseMeta struct {
	Count int64 `json:"count" api:"required"`
	Limit int64 `json:"limit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Limit       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceRunListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceRunListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceRunListResponseOperation struct {
	ID                 string                                         `json:"id" api:"required"`
	Method             string                                         `json:"method" api:"required"`
	Path               string                                         `json:"path" api:"required"`
	Category           APIPerformanceRunListResponseOperationCategory `json:"category" api:"nullable"`
	CreatedAt          string                                         `json:"created_at" api:"nullable"`
	Description        string                                         `json:"description" api:"nullable"`
	InterfaceID        string                                         `json:"interface_id" api:"nullable"`
	OpenAPIOperationID string                                         `json:"openapi_operation_id" api:"nullable"`
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
func (r APIPerformanceRunListResponseOperation) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceRunListResponseOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceRunListResponseOperationCategory struct {
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
func (r APIPerformanceRunListResponseOperationCategory) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceRunListResponseOperationCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceRunListParams struct {
	OperationID string           `query:"operation_id" api:"required" json:"-"`
	IncludeRaw  param.Opt[bool]  `query:"include_raw,omitzero" json:"-"`
	Limit       param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIPerformanceRunListParams]'s query parameters as
// `url.Values`.
func (r APIPerformanceRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
