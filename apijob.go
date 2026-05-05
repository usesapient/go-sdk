// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sapient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/sapient-go/internal/apiquery"
	"github.com/stainless-sdks/sapient-go/internal/requestconfig"
	"github.com/stainless-sdks/sapient-go/option"
	"github.com/stainless-sdks/sapient-go/packages/param"
)

// APIJobService contains methods and other services that help with interacting
// with the sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIJobService] method instead.
type APIJobService struct {
	options []option.RequestOption
}

// NewAPIJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIJobService(opts ...option.RequestOption) (r APIJobService) {
	r = APIJobService{}
	r.options = opts
	return
}

// List Jobs
func (r *APIJobService) ListJobs(ctx context.Context, query APIJobListJobsParams, opts ...option.RequestOption) (res *APIJobListJobsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/api/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve Job
func (r *APIJobService) GetJob(ctx context.Context, jobID string, opts ...option.RequestOption) (res *PublicJob, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api/jobs/%s", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Jobs
func (r *APIJobService) List(ctx context.Context, query APIJobListParams, opts ...option.RequestOption) (res *APIJobListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/api/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve Job
func (r *APIJobService) Get(ctx context.Context, jobID string, opts ...option.RequestOption) (res *PublicJob, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/api/jobs/%s", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type APIJobListJobsResponse = any

type APIJobListResponse = any

type APIJobListJobsParams struct {
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	Type   param.Opt[string] `query:"type,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIJobListJobsParams]'s query parameters as `url.Values`.
func (r APIJobListJobsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIJobListParams struct {
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	Type   param.Opt[string] `query:"type,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIJobListParams]'s query parameters as `url.Values`.
func (r APIJobListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
