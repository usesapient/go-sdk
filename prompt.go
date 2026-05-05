// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sapient

import (
	"context"
	"net/http"
	"slices"

	"github.com/usesapient/go-sdk/internal/apijson"
	"github.com/usesapient/go-sdk/internal/requestconfig"
	"github.com/usesapient/go-sdk/option"
	"github.com/usesapient/go-sdk/packages/param"
	"github.com/usesapient/go-sdk/packages/respjson"
)

// PromptService contains methods and other services that help with interacting
// with the sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPromptService] method instead.
type PromptService struct {
	options []option.RequestOption
}

// NewPromptService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPromptService(opts ...option.RequestOption) (r PromptService) {
	r = PromptService{}
	r.options = opts
	return
}

// Generate From List
func (r *PromptService) GenerateFromList(ctx context.Context, body PromptGenerateFromListParams, opts ...option.RequestOption) (res *PromptGenerateFromListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/prompts/generate-from-list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PromptGenerateFromListResponse struct {
	Count   int64                                  `json:"count" api:"required"`
	Prompts []PromptGenerateFromListResponsePrompt `json:"prompts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Prompts     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptGenerateFromListResponse) RawJSON() string { return r.JSON.raw }
func (r *PromptGenerateFromListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptGenerateFromListResponsePrompt struct {
	SourceName   string         `json:"source_name" api:"required"`
	Text         string         `json:"text" api:"required"`
	Goal         string         `json:"goal" api:"nullable"`
	Metadata     map[string]any `json:"metadata"`
	SourceDomain string         `json:"source_domain" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SourceName   respjson.Field
		Text         respjson.Field
		Goal         respjson.Field
		Metadata     respjson.Field
		SourceDomain respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptGenerateFromListResponsePrompt) RawJSON() string { return r.JSON.raw }
func (r *PromptGenerateFromListResponsePrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptGenerateFromListParams struct {
	Items []PromptGenerateFromListParamsItem `json:"items,omitzero" api:"required"`
	// Prompt goal or evaluation theme.
	Goal              param.Opt[string] `json:"goal,omitzero"`
	CountPerItem      param.Opt[int64]  `json:"count_per_item,omitzero"`
	IncludeAgentTasks param.Opt[bool]   `json:"include_agent_tasks,omitzero"`
	Language          param.Opt[string] `json:"language,omitzero"`
	Region            param.Opt[string] `json:"region,omitzero"`
	paramObj
}

func (r PromptGenerateFromListParams) MarshalJSON() (data []byte, err error) {
	type shadow PromptGenerateFromListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PromptGenerateFromListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type PromptGenerateFromListParamsItem struct {
	Name     string            `json:"name" api:"required"`
	Category param.Opt[string] `json:"category,omitzero"`
	Domain   param.Opt[string] `json:"domain,omitzero"`
	Notes    param.Opt[string] `json:"notes,omitzero"`
	paramObj
}

func (r PromptGenerateFromListParamsItem) MarshalJSON() (data []byte, err error) {
	type shadow PromptGenerateFromListParamsItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PromptGenerateFromListParamsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
