// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sapient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/sapient-go/internal/apijson"
	"github.com/stainless-sdks/sapient-go/internal/requestconfig"
	"github.com/stainless-sdks/sapient-go/option"
	"github.com/stainless-sdks/sapient-go/packages/param"
	"github.com/stainless-sdks/sapient-go/packages/respjson"
)

// WorkflowService contains methods and other services that help with interacting
// with the sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowService] method instead.
type WorkflowService struct {
	options []option.RequestOption
}

// NewWorkflowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowService(opts ...option.RequestOption) (r WorkflowService) {
	r = WorkflowService{}
	r.options = opts
	return
}

// Create Workflow
func (r *WorkflowService) New(ctx context.Context, body WorkflowNewParams, opts ...option.RequestOption) (res *WorkflowResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/workflows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve Workflow
func (r *WorkflowService) Get(ctx context.Context, workflowID string, opts ...option.RequestOption) (res *WorkflowResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workflows/%s", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Run Workflow
func (r *WorkflowService) Run(ctx context.Context, workflowID string, body WorkflowRunParams, opts ...option.RequestOption) (res *WorkflowRunResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workflows/%s/run", url.PathEscape(workflowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type WorkflowResponse struct {
	ID          string           `json:"id" api:"required"`
	Name        string           `json:"name" api:"required"`
	CreatedAt   string           `json:"created_at" api:"nullable"`
	Description string           `json:"description" api:"nullable"`
	Metadata    map[string]any   `json:"metadata"`
	Steps       []map[string]any `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Metadata    respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowRunResponse struct {
	RunID      string    `json:"run_id" api:"required"`
	Status     string    `json:"status" api:"required"`
	WorkflowID string    `json:"workflow_id" api:"required"`
	Job        PublicJob `json:"job" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RunID       respjson.Field
		Status      respjson.Field
		WorkflowID  respjson.Field
		Job         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRunResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowNewParams struct {
	Name        string            `json:"name" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	Metadata    map[string]any    `json:"metadata,omitzero"`
	Steps       []map[string]any  `json:"steps,omitzero"`
	paramObj
}

func (r WorkflowNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowRunParams struct {
	// Workflow ID. Optional when passed in the URL.
	WorkflowID param.Opt[string] `json:"workflow_id,omitzero"`
	Watch      param.Opt[bool]   `json:"watch,omitzero"`
	Input      map[string]any    `json:"input,omitzero"`
	paramObj
}

func (r WorkflowRunParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
