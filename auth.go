// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sapient

import (
	"context"
	"net/http"
	"slices"

	"github.com/usesapient/go-sdk/internal/requestconfig"
	"github.com/usesapient/go-sdk/option"
)

// AuthService contains methods and other services that help with interacting with
// the sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthService] method instead.
type AuthService struct {
	options []option.RequestOption
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r AuthService) {
	r = AuthService{}
	r.options = opts
	return
}

// Get Auth Status
func (r *AuthService) GetStatus(ctx context.Context, opts ...option.RequestOption) (res *AuthGetStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/auth/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get Auth Status
func (r *AuthService) Status(ctx context.Context, opts ...option.RequestOption) (res *AuthStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/auth/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AuthGetStatusResponse = any

type AuthStatusResponse = any
