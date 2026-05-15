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

// APIAPIPerformanceService contains methods and other services that help with
// interacting with the sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIAPIPerformanceService] method instead.
type APIAPIPerformanceService struct {
	options []option.RequestOption
}

// NewAPIAPIPerformanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIAPIPerformanceService(opts ...option.RequestOption) (r APIAPIPerformanceService) {
	r = APIAPIPerformanceService{}
	r.options = opts
	return
}

// Failure Analysis
func (r *APIAPIPerformanceService) FailureAnalysis(ctx context.Context, brandID string, query APIAPIPerformanceFailureAnalysisParams, opts ...option.RequestOption) (res *FailureAnalysisSnapshot, err error) {
	opts = slices.Concat(r.options, opts)
	if brandID == "" {
		err = errors.New("missing required brand_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api/api-performance/failure-analysis/%s", url.PathEscape(brandID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Latest Runs
func (r *APIAPIPerformanceService) LatestRuns(ctx context.Context, brandID string, query APIAPIPerformanceLatestRunsParams, opts ...option.RequestOption) (res *ApiapiPerformanceLatestRunsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if brandID == "" {
		err = errors.New("missing required brand_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api/api-performance/latest-runs/%s", url.PathEscape(brandID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type FailureAnalysisSnapshot struct {
	EndpointsByFailureCount []map[string]any                 `json:"endpoints_by_failure_count"`
	FailedRuns              int64                            `json:"failed_runs"`
	FailureRate             float64                          `json:"failure_rate" api:"nullable"`
	Patterns                []FailureAnalysisSnapshotPattern `json:"patterns"`
	TopFailureReasons       []map[string]any                 `json:"top_failure_reasons"`
	TotalRuns               int64                            `json:"total_runs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndpointsByFailureCount respjson.Field
		FailedRuns              respjson.Field
		FailureRate             respjson.Field
		Patterns                respjson.Field
		TopFailureReasons       respjson.Field
		TotalRuns               respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FailureAnalysisSnapshot) RawJSON() string { return r.JSON.raw }
func (r *FailureAnalysisSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FailureAnalysisSnapshotPattern struct {
	Count             int64            `json:"count" api:"required"`
	Label             string           `json:"label" api:"required"`
	Pattern           string           `json:"pattern" api:"required"`
	AffectedEndpoints []map[string]any `json:"affected_endpoints"`
	AffectedModels    []string         `json:"affected_models"`
	Description       string           `json:"description" api:"nullable"`
	ExampleReasons    []string         `json:"example_reasons"`
	Percentage        float64          `json:"percentage" api:"nullable"`
	RunIDs            []string         `json:"run_ids"`
	Severity          string           `json:"severity" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count             respjson.Field
		Label             respjson.Field
		Pattern           respjson.Field
		AffectedEndpoints respjson.Field
		AffectedModels    respjson.Field
		Description       respjson.Field
		ExampleReasons    respjson.Field
		Percentage        respjson.Field
		RunIDs            respjson.Field
		Severity          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FailureAnalysisSnapshotPattern) RawJSON() string { return r.JSON.raw }
func (r *FailureAnalysisSnapshotPattern) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApiapiPerformanceLatestRunsResponse struct {
	Data []EvalRunSummary `json:"data" api:"required"`
	Meta PublicListMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApiapiPerformanceLatestRunsResponse) RawJSON() string { return r.JSON.raw }
func (r *ApiapiPerformanceLatestRunsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIAPIPerformanceFailureAnalysisParams struct {
	IntegrationID param.Opt[string] `query:"integration_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIAPIPerformanceFailureAnalysisParams]'s query parameters
// as `url.Values`.
func (r APIAPIPerformanceFailureAnalysisParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIAPIPerformanceLatestRunsParams struct {
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIAPIPerformanceLatestRunsParams]'s query parameters as
// `url.Values`.
func (r APIAPIPerformanceLatestRunsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
