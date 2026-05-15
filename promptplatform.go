// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomusesapientgosdk

import (
	"context"
	"net/http"
	"slices"

	"github.com/usesapient/go-sdk/internal/requestconfig"
	"github.com/usesapient/go-sdk/option"
)

// PromptPlatformService contains methods and other services that help with
// interacting with the Sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPromptPlatformService] method instead.
type PromptPlatformService struct {
	Options []option.RequestOption
}

// NewPromptPlatformService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPromptPlatformService(opts ...option.RequestOption) (r PromptPlatformService) {
	r = PromptPlatformService{}
	r.Options = opts
	return
}

// List Platforms
func (r *PromptPlatformService) List(ctx context.Context, opts ...option.RequestOption) (res *PromptPlatformListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/prompts/platforms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PromptPlatformListResponse = any
