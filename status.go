// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomusesapientgosdk

import (
	"context"
	"net/http"
	"slices"

	"github.com/usesapient/go-sdk/internal/requestconfig"
	"github.com/usesapient/go-sdk/option"
)

// StatusService contains methods and other services that help with interacting
// with the sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatusService] method instead.
type StatusService struct {
	options []option.RequestOption
}

// NewStatusService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStatusService(opts ...option.RequestOption) (r StatusService) {
	r = StatusService{}
	r.options = opts
	return
}

// Get Status
func (r *StatusService) Get(ctx context.Context, opts ...option.RequestOption) (res *StatusGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get Status
func (r *StatusService) Get(ctx context.Context, opts ...option.RequestOption) (res *StatusGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type StatusGetResponse = any
