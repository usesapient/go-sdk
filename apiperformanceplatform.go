// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomusesapientgosdk

import (
	"context"
	"net/http"
	"slices"

	"github.com/usesapient/go-sdk/internal/apijson"
	"github.com/usesapient/go-sdk/internal/requestconfig"
	"github.com/usesapient/go-sdk/option"
	"github.com/usesapient/go-sdk/packages/respjson"
)

// APIPerformancePlatformService contains methods and other services that help with
// interacting with the Sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIPerformancePlatformService] method instead.
type APIPerformancePlatformService struct {
	Options []option.RequestOption
}

// NewAPIPerformancePlatformService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIPerformancePlatformService(opts ...option.RequestOption) (r APIPerformancePlatformService) {
	r = APIPerformancePlatformService{}
	r.Options = opts
	return
}

// List Platforms
func (r *APIPerformancePlatformService) List(ctx context.Context, opts ...option.RequestOption) (res *APIPerformancePlatformListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api-performance/platforms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Estimate Platform Cost
func (r *APIPerformancePlatformService) EstimateCost(ctx context.Context, opts ...option.RequestOption) (res *APIPerformancePlatformEstimateCostResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api-performance/platforms/estimate-cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type APIPerformancePlatformListResponse struct {
	AgentPlatforms []APIPerformancePlatformListResponseAgentPlatform `json:"agent_platforms" api:"required"`
	TextPlatforms  []APIPerformancePlatformListResponseTextPlatform  `json:"text_platforms" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentPlatforms respjson.Field
		TextPlatforms  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformancePlatformListResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformancePlatformListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformancePlatformListResponseAgentPlatform struct {
	ID             string `json:"id" api:"required"`
	CreditsPerEval int64  `json:"credits_per_eval" api:"required"`
	Label          string `json:"label" api:"required"`
	Provider       string `json:"provider" api:"required"`
	Type           string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreditsPerEval respjson.Field
		Label          respjson.Field
		Provider       respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformancePlatformListResponseAgentPlatform) RawJSON() string { return r.JSON.raw }
func (r *APIPerformancePlatformListResponseAgentPlatform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformancePlatformListResponseTextPlatform struct {
	ID             string `json:"id" api:"required"`
	CreditsPerEval int64  `json:"credits_per_eval" api:"required"`
	Label          string `json:"label" api:"required"`
	Provider       string `json:"provider" api:"required"`
	Type           string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreditsPerEval respjson.Field
		Label          respjson.Field
		Provider       respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformancePlatformListResponseTextPlatform) RawJSON() string { return r.JSON.raw }
func (r *APIPerformancePlatformListResponseTextPlatform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformancePlatformEstimateCostResponse struct {
	Estimates    map[string]APIPerformancePlatformEstimateCostResponseEstimate `json:"estimates" api:"required"`
	TotalCredits int64                                                         `json:"total_credits" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Estimates    respjson.Field
		TotalCredits respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformancePlatformEstimateCostResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformancePlatformEstimateCostResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformancePlatformEstimateCostResponseEstimate struct {
	CreditsPerEval int64 `json:"credits_per_eval" api:"required"`
	EvalCount      int64 `json:"eval_count" api:"required"`
	TotalCredits   int64 `json:"total_credits" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditsPerEval respjson.Field
		EvalCount      respjson.Field
		TotalCredits   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformancePlatformEstimateCostResponseEstimate) RawJSON() string { return r.JSON.raw }
func (r *APIPerformancePlatformEstimateCostResponseEstimate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
