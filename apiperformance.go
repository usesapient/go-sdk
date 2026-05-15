// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomusesapientgosdk

import (
	"github.com/usesapient/go-sdk/option"
)

// APIPerformanceService contains methods and other services that help with
// interacting with the Sapient API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIPerformanceService] method instead.
type APIPerformanceService struct {
	Options          []option.RequestOption
	Platforms        APIPerformancePlatformService
	Interfaces       APIPerformanceInterfaceService
	EvaluationConfig APIPerformanceEvaluationConfigService
	Operations       APIPerformanceOperationService
	OperationPrompts APIPerformanceOperationPromptService
	Runs             APIPerformanceRunService
	Prompts          APIPerformancePromptService
	UseCases         APIPerformanceUseCaseService
}

// NewAPIPerformanceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIPerformanceService(opts ...option.RequestOption) (r APIPerformanceService) {
	r = APIPerformanceService{}
	r.Options = opts
	r.Platforms = NewAPIPerformancePlatformService(opts...)
	r.Interfaces = NewAPIPerformanceInterfaceService(opts...)
	r.EvaluationConfig = NewAPIPerformanceEvaluationConfigService(opts...)
	r.Operations = NewAPIPerformanceOperationService(opts...)
	r.OperationPrompts = NewAPIPerformanceOperationPromptService(opts...)
	r.Runs = NewAPIPerformanceRunService(opts...)
	r.Prompts = NewAPIPerformancePromptService(opts...)
	r.UseCases = NewAPIPerformanceUseCaseService(opts...)
	return
}
