// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomusesapientgosdk

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

// ContextService contains methods and other services that help with interacting
// with the sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContextService] method instead.
type ContextService struct {
	options []option.RequestOption
}

// NewContextService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewContextService(opts ...option.RequestOption) (r ContextService) {
	r = ContextService{}
	r.options = opts
	return
}

// Get Company Context
func (r *ContextService) GetCompany(ctx context.Context, company string, query ContextGetCompanyParams, opts ...option.RequestOption) (res *ContextGetCompanyResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if company == "" {
		err = errors.New("missing required company parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/context/company/%s", url.PathEscape(company))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get Company Context
func (r *ContextService) Company(ctx context.Context, company string, query ContextCompanyParams, opts ...option.RequestOption) (res *ContextCompanyResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if company == "" {
		err = errors.New("missing required company parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/context/company/%s", url.PathEscape(company))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type EvalRunSummary struct {
	ID             string                      `json:"id" api:"required"`
	Endpoint       EvalRunSummaryEndpointUnion `json:"endpoint" api:"nullable"`
	EvalType       string                      `json:"eval_type" api:"nullable"`
	FailureReasons []string                    `json:"failure_reasons" api:"nullable"`
	LatencyMs      int64                       `json:"latency_ms" api:"nullable"`
	Model          string                      `json:"model" api:"nullable"`
	ModelType      string                      `json:"model_type" api:"nullable"`
	Passed         bool                        `json:"passed" api:"nullable"`
	Platform       string                      `json:"platform" api:"nullable"`
	RunAt          string                      `json:"run_at" api:"nullable"`
	Score          float64                     `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Endpoint       respjson.Field
		EvalType       respjson.Field
		FailureReasons respjson.Field
		LatencyMs      respjson.Field
		Model          respjson.Field
		ModelType      respjson.Field
		Passed         respjson.Field
		Platform       respjson.Field
		RunAt          respjson.Field
		Score          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalRunSummary) RawJSON() string { return r.JSON.raw }
func (r *EvalRunSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EvalRunSummaryEndpointUnion contains all possible properties and values from
// [PublicEndpoint], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfEvalRunSummaryEndpointMapItem]
type EvalRunSummaryEndpointUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfEvalRunSummaryEndpointMapItem any `json:",inline"`
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
		OfEvalRunSummaryEndpointMapItem respjson.Field
		ID                              respjson.Field
		Category                        respjson.Field
		Description                     respjson.Field
		Method                          respjson.Field
		Path                            respjson.Field
		raw                             string
	} `json:"-"`
}

func (u EvalRunSummaryEndpointUnion) AsPublicEndpoint() (v PublicEndpoint) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EvalRunSummaryEndpointUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EvalRunSummaryEndpointUnion) RawJSON() string { return u.JSON.raw }

func (r *EvalRunSummaryEndpointUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicCompany struct {
	// Sapient brand ID.
	BrandID string `json:"brand_id" api:"required"`
	Name    string `json:"name" api:"required"`
	Domain  string `json:"domain" api:"nullable"`
	LogoURL string `json:"logo_url" api:"nullable"`
	// Organization-scoped brand relationship ID.
	OrgBrandID   string `json:"org_brand_id" api:"nullable"`
	Relationship string `json:"relationship" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandID      respjson.Field
		Name         respjson.Field
		Domain       respjson.Field
		LogoURL      respjson.Field
		OrgBrandID   respjson.Field
		Relationship respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicCompany) RawJSON() string { return r.JSON.raw }
func (r *PublicCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicEndpoint struct {
	ID          string         `json:"id" api:"required"`
	Category    map[string]any `json:"category" api:"nullable"`
	Description string         `json:"description" api:"nullable"`
	Method      string         `json:"method" api:"nullable"`
	Path        string         `json:"path" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Category    respjson.Field
		Description respjson.Field
		Method      respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicEndpoint) RawJSON() string { return r.JSON.raw }
func (r *PublicEndpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextGetCompanyResponse struct {
	Company         PublicCompany           `json:"company" api:"required"`
	Endpoints       []PublicEndpoint        `json:"endpoints"`
	FailureAnalysis FailureAnalysisSnapshot `json:"failure_analysis" api:"nullable"`
	LatestRuns      []EvalRunSummary        `json:"latest_runs"`
	Markdown        string                  `json:"markdown" api:"nullable"`
	// Requested context use case.
	Purpose string `json:"purpose" api:"nullable"`
	// Requested lookback window.
	Since   string         `json:"since" api:"nullable"`
	Summary map[string]any `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company         respjson.Field
		Endpoints       respjson.Field
		FailureAnalysis respjson.Field
		LatestRuns      respjson.Field
		Markdown        respjson.Field
		Purpose         respjson.Field
		Since           respjson.Field
		Summary         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextGetCompanyResponse) RawJSON() string { return r.JSON.raw }
func (r *ContextGetCompanyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextCompanyResponse struct {
	Company         PublicCompany           `json:"company" api:"required"`
	Endpoints       []PublicEndpoint        `json:"endpoints"`
	FailureAnalysis FailureAnalysisSnapshot `json:"failure_analysis" api:"nullable"`
	LatestRuns      []EvalRunSummary        `json:"latest_runs"`
	Markdown        string                  `json:"markdown" api:"nullable"`
	// Requested context use case.
	Purpose string `json:"purpose" api:"nullable"`
	// Requested lookback window.
	Since   string         `json:"since" api:"nullable"`
	Summary map[string]any `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company         respjson.Field
		Endpoints       respjson.Field
		FailureAnalysis respjson.Field
		LatestRuns      respjson.Field
		Markdown        respjson.Field
		Purpose         respjson.Field
		Since           respjson.Field
		Summary         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextCompanyResponse) RawJSON() string { return r.JSON.raw }
func (r *ContextCompanyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextGetCompanyParams struct {
	For   param.Opt[string] `query:"for,omitzero" json:"-"`
	Since param.Opt[string] `query:"since,omitzero" json:"-"`
	// Any of "json", "markdown", "md".
	Format ContextGetCompanyParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ContextGetCompanyParams]'s query parameters as
// `url.Values`.
func (r ContextGetCompanyParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContextGetCompanyParamsFormat string

const (
	ContextGetCompanyParamsFormatJson     ContextGetCompanyParamsFormat = "json"
	ContextGetCompanyParamsFormatMarkdown ContextGetCompanyParamsFormat = "markdown"
	ContextGetCompanyParamsFormatMd       ContextGetCompanyParamsFormat = "md"
)

type ContextCompanyParams struct {
	For   param.Opt[string] `query:"for,omitzero" json:"-"`
	Since param.Opt[string] `query:"since,omitzero" json:"-"`
	// Any of "json", "markdown", "md".
	Format ContextCompanyParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ContextCompanyParams]'s query parameters as `url.Values`.
func (r ContextCompanyParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContextCompanyParamsFormat string

const (
	ContextCompanyParamsFormatJson     ContextCompanyParamsFormat = "json"
	ContextCompanyParamsFormatMarkdown ContextCompanyParamsFormat = "markdown"
	ContextCompanyParamsFormatMd       ContextCompanyParamsFormat = "md"
)
