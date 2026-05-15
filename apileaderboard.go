// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomusesapientgosdk

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/usesapient/go-sdk/internal/apiquery"
	"github.com/usesapient/go-sdk/internal/requestconfig"
	"github.com/usesapient/go-sdk/option"
	"github.com/usesapient/go-sdk/packages/param"
)

// APILeaderboardService contains methods and other services that help with
// interacting with the sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPILeaderboardService] method instead.
type APILeaderboardService struct {
	options   []option.RequestOption
	Companies APILeaderboardCompanyService
}

// NewAPILeaderboardService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPILeaderboardService(opts ...option.RequestOption) (r APILeaderboardService) {
	r = APILeaderboardService{}
	r.options = opts
	r.Companies = NewAPILeaderboardCompanyService(opts...)
	return
}

// List Companies
func (r *APILeaderboardService) ListCompanies(ctx context.Context, query APILeaderboardListCompaniesParams, opts ...option.RequestOption) (res *APILeaderboardListCompaniesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/api/leaderboard/companies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type APILeaderboardListCompaniesResponse = any

type APILeaderboardListCompaniesParams struct {
	Category param.Opt[string] `query:"category,omitzero" json:"-"`
	Limit    param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APILeaderboardListCompaniesParams]'s query parameters as
// `url.Values`.
func (r APILeaderboardListCompaniesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
