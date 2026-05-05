// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sapient

import (
	"context"
	"encoding/json"
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

// EvalRunService contains methods and other services that help with interacting
// with the sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvalRunService] method instead.
type EvalRunService struct {
	options []option.RequestOption
}

// NewEvalRunService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEvalRunService(opts ...option.RequestOption) (r EvalRunService) {
	r = EvalRunService{}
	r.options = opts
	return
}

// Retrieve Eval Run
func (r *EvalRunService) Get(ctx context.Context, runID string, opts ...option.RequestOption) (res *EvalRunGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/eval-runs/%s", url.PathEscape(runID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Eval Runs
func (r *EvalRunService) List(ctx context.Context, query EvalRunListParams, opts ...option.RequestOption) (res *EvalRunListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/eval-runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Diagnose Eval Runs
func (r *EvalRunService) Diagnose(ctx context.Context, body EvalRunDiagnoseParams, opts ...option.RequestOption) (res *EvalRunDiagnoseResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/eval-runs/diagnose"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type EvalRunDetail struct {
	ID              string                     `json:"id" api:"required"`
	Endpoint        EvalRunDetailEndpointUnion `json:"endpoint" api:"nullable"`
	EvalType        string                     `json:"eval_type" api:"nullable"`
	ExecutionStderr string                     `json:"execution_stderr" api:"nullable"`
	ExecutionStdout string                     `json:"execution_stdout" api:"nullable"`
	ExitCode        int64                      `json:"exit_code" api:"nullable"`
	FailureReasons  []string                   `json:"failure_reasons" api:"nullable"`
	LatencyMs       int64                      `json:"latency_ms" api:"nullable"`
	Model           string                     `json:"model" api:"nullable"`
	ModelType       string                     `json:"model_type" api:"nullable"`
	Passed          bool                       `json:"passed" api:"nullable"`
	Platform        string                     `json:"platform" api:"nullable"`
	PromptSent      string                     `json:"prompt_sent" api:"nullable"`
	Raw             map[string]any             `json:"raw" api:"nullable"`
	ResponseText    string                     `json:"response_text" api:"nullable"`
	RunAt           string                     `json:"run_at" api:"nullable"`
	Score           float64                    `json:"score" api:"nullable"`
	ToolCalls       []map[string]any           `json:"tool_calls" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Endpoint        respjson.Field
		EvalType        respjson.Field
		ExecutionStderr respjson.Field
		ExecutionStdout respjson.Field
		ExitCode        respjson.Field
		FailureReasons  respjson.Field
		LatencyMs       respjson.Field
		Model           respjson.Field
		ModelType       respjson.Field
		Passed          respjson.Field
		Platform        respjson.Field
		PromptSent      respjson.Field
		Raw             respjson.Field
		ResponseText    respjson.Field
		RunAt           respjson.Field
		Score           respjson.Field
		ToolCalls       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalRunDetail) RawJSON() string { return r.JSON.raw }
func (r *EvalRunDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EvalRunDetailEndpointUnion contains all possible properties and values from
// [PublicEndpoint], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfEvalRunDetailEndpointMapItem]
type EvalRunDetailEndpointUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfEvalRunDetailEndpointMapItem any `json:",inline"`
	// This field is from variant [PublicEndpoint].
	ID string `json:"id"`
	// This field is from variant [PublicEndpoint].
	Category map[string]any `json:"category"`
	// This field is from variant [PublicEndpoint].
	Description string `json:"description"`
	// This field is from variant [PublicEndpoint].
	Method string `json:"method"`
	// This field is from variant [PublicEndpoint].
	Path string `json:"path"`
	JSON struct {
		OfEvalRunDetailEndpointMapItem respjson.Field
		ID                             respjson.Field
		Category                       respjson.Field
		Description                    respjson.Field
		Method                         respjson.Field
		Path                           respjson.Field
		raw                            string
	} `json:"-"`
}

func (u EvalRunDetailEndpointUnion) AsPublicEndpoint() (v PublicEndpoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvalRunDetailEndpointUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EvalRunDetailEndpointUnion) RawJSON() string { return u.JSON.raw }

func (r *EvalRunDetailEndpointUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicListMeta struct {
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
func (r PublicListMeta) RawJSON() string { return r.JSON.raw }
func (r *PublicListMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalRunGetResponse struct {
	Data EvalRunDetail `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalRunGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EvalRunGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalRunListResponse struct {
	Company PublicCompany    `json:"company" api:"required"`
	Data    []EvalRunSummary `json:"data" api:"required"`
	Meta    PublicListMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company     respjson.Field
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalRunListResponse) RawJSON() string { return r.JSON.raw }
func (r *EvalRunListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalRunDiagnoseResponse struct {
	Company         PublicCompany                  `json:"company" api:"required"`
	FailureAnalysis FailureAnalysisSnapshot        `json:"failure_analysis" api:"required"`
	Since           string                         `json:"since" api:"required"`
	Summary         EvalRunDiagnoseResponseSummary `json:"summary" api:"required"`
	Examples        []EvalRunDetail                `json:"examples"`
	Markdown        string                         `json:"markdown" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company         respjson.Field
		FailureAnalysis respjson.Field
		Since           respjson.Field
		Summary         respjson.Field
		Examples        respjson.Field
		Markdown        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalRunDiagnoseResponse) RawJSON() string { return r.JSON.raw }
func (r *EvalRunDiagnoseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalRunDiagnoseResponseSummary struct {
	FailedRuns         int64            `json:"failed_runs"`
	FailureRate        float64          `json:"failure_rate" api:"nullable"`
	PassedRuns         int64            `json:"passed_runs"`
	TopFailedEndpoints []map[string]any `json:"top_failed_endpoints"`
	TopFailedModels    []map[string]any `json:"top_failed_models"`
	TopFailureReasons  []map[string]any `json:"top_failure_reasons"`
	TotalRuns          int64            `json:"total_runs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FailedRuns         respjson.Field
		FailureRate        respjson.Field
		PassedRuns         respjson.Field
		TopFailedEndpoints respjson.Field
		TopFailedModels    respjson.Field
		TopFailureReasons  respjson.Field
		TotalRuns          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalRunDiagnoseResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *EvalRunDiagnoseResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalRunListParams struct {
	Company string            `query:"company" api:"required" json:"-"`
	Since   param.Opt[string] `query:"since,omitzero" json:"-"`
	Limit   param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EvalRunListParams]'s query parameters as `url.Values`.
func (r EvalRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EvalRunDiagnoseParams struct {
	// Company name, domain, brand ID, or org brand ID.
	Company         string           `json:"company" api:"required"`
	IncludeExamples param.Opt[bool]  `json:"include_examples,omitzero"`
	MaxExamples     param.Opt[int64] `json:"max_examples,omitzero"`
	// Lookback window, such as 20d, 72h, or an ISO timestamp.
	Since param.Opt[string] `json:"since,omitzero"`
	// Report format to include in the response.
	//
	// Any of "json", "markdown", "md".
	Format EvalRunDiagnoseParamsFormat `json:"format,omitzero"`
	paramObj
}

func (r EvalRunDiagnoseParams) MarshalJSON() (data []byte, err error) {
	type shadow EvalRunDiagnoseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalRunDiagnoseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Report format to include in the response.
type EvalRunDiagnoseParamsFormat string

const (
	EvalRunDiagnoseParamsFormatJson     EvalRunDiagnoseParamsFormat = "json"
	EvalRunDiagnoseParamsFormatMarkdown EvalRunDiagnoseParamsFormat = "markdown"
	EvalRunDiagnoseParamsFormatMd       EvalRunDiagnoseParamsFormat = "md"
)
