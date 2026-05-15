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

// APIPerformanceInterfaceService contains methods and other services that help
// with interacting with the Sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIPerformanceInterfaceService] method instead.
type APIPerformanceInterfaceService struct {
	Options []option.RequestOption
}

// NewAPIPerformanceInterfaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIPerformanceInterfaceService(opts ...option.RequestOption) (r APIPerformanceInterfaceService) {
	r = APIPerformanceInterfaceService{}
	r.Options = opts
	return
}

// List Interfaces
func (r *APIPerformanceInterfaceService) List(ctx context.Context, opts ...option.RequestOption) (res *APIPerformanceInterfaceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api-performance/interfaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type APIPerformanceInterfaceListResponse struct {
	Interfaces []APIPerformanceInterfaceListResponseInterface `json:"interfaces" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Interfaces  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceInterfaceListResponse) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceInterfaceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIPerformanceInterfaceListResponseInterface struct {
	ID             string `json:"id" api:"required"`
	Label          string `json:"label" api:"required"`
	Type           string `json:"type" api:"required"`
	CreatedAt      string `json:"created_at" api:"nullable"`
	OperationCount int64  `json:"operation_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Label          respjson.Field
		Type           respjson.Field
		CreatedAt      respjson.Field
		OperationCount respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPerformanceInterfaceListResponseInterface) RawJSON() string { return r.JSON.raw }
func (r *APIPerformanceInterfaceListResponseInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
