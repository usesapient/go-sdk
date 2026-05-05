# Status

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#StatusGetResponse">StatusGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#StatusGetResponse">StatusGetResponse</a>

Methods:

- <code title="get /v1/status">client.Status.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#StatusService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#StatusGetResponse">StatusGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/status">client.Status.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#StatusService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#StatusGetResponse">StatusGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#AuthGetStatusResponse">AuthGetStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#AuthStatusResponse">AuthStatusResponse</a>

Methods:

- <code title="get /v1/auth/status">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#AuthService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#AuthGetStatusResponse">AuthGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/auth/status">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#AuthService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#AuthStatusResponse">AuthStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Context

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#EvalRunSummary">EvalRunSummary</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#PublicCompany">PublicCompany</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#PublicEndpoint">PublicEndpoint</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#ContextGetCompanyResponse">ContextGetCompanyResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#ContextCompanyResponse">ContextCompanyResponse</a>

Methods:

- <code title="get /v1/context/company/{company}">client.Context.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#ContextService.GetCompany">GetCompany</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, company <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#ContextGetCompanyParams">ContextGetCompanyParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#ContextGetCompanyResponse">ContextGetCompanyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/context/company/{company}">client.Context.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#ContextService.Company">Company</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, company <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#ContextCompanyParams">ContextCompanyParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#ContextCompanyResponse">ContextCompanyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# EvalRuns

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#EvalRunDetail">EvalRunDetail</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#PublicListMeta">PublicListMeta</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#EvalRunGetResponse">EvalRunGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#EvalRunListResponse">EvalRunListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#EvalRunDiagnoseResponse">EvalRunDiagnoseResponse</a>

Methods:

- <code title="get /v1/eval-runs/{run_id}">client.EvalRuns.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#EvalRunService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, runID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#EvalRunGetResponse">EvalRunGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/eval-runs">client.EvalRuns.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#EvalRunService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#EvalRunListParams">EvalRunListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#EvalRunListResponse">EvalRunListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/eval-runs/diagnose">client.EvalRuns.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#EvalRunService.Diagnose">Diagnose</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#EvalRunDiagnoseParams">EvalRunDiagnoseParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#EvalRunDiagnoseResponse">EvalRunDiagnoseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Prompts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#PromptGenerateFromListResponse">PromptGenerateFromListResponse</a>

Methods:

- <code title="post /v1/prompts/generate-from-list">client.Prompts.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#PromptService.GenerateFromList">GenerateFromList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#PromptGenerateFromListParams">PromptGenerateFromListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#PromptGenerateFromListResponse">PromptGenerateFromListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Workflows

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#WorkflowResponse">WorkflowResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#WorkflowRunResponse">WorkflowRunResponse</a>

Methods:

- <code title="post /v1/workflows">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#WorkflowService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#WorkflowNewParams">WorkflowNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#WorkflowResponse">WorkflowResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/workflows/{workflow_id}">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#WorkflowService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#WorkflowResponse">WorkflowResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/workflows/{workflow_id}/run">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#WorkflowService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#WorkflowRunParams">WorkflowRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#WorkflowRunResponse">WorkflowRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WorkflowRuns

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#PublicJob">PublicJob</a>

Methods:

- <code title="get /v1/workflow-runs/{run_id}">client.WorkflowRuns.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#WorkflowRunService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, runID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#PublicJob">PublicJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# API

## Leaderboard

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APILeaderboardListCompaniesResponse">APILeaderboardListCompaniesResponse</a>

Methods:

- <code title="get /v1/api/leaderboard/companies">client.API.Leaderboard.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APILeaderboardService.ListCompanies">ListCompanies</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APILeaderboardListCompaniesParams">APILeaderboardListCompaniesParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APILeaderboardListCompaniesResponse">APILeaderboardListCompaniesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Companies

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APILeaderboardCompanyListResponse">APILeaderboardCompanyListResponse</a>

Methods:

- <code title="get /v1/api/leaderboard/companies">client.API.Leaderboard.Companies.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APILeaderboardCompanyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APILeaderboardCompanyListParams">APILeaderboardCompanyListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APILeaderboardCompanyListResponse">APILeaderboardCompanyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## APIPerformance

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#FailureAnalysisSnapshot">FailureAnalysisSnapshot</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#ApiapiPerformanceLatestRunsResponse">ApiapiPerformanceLatestRunsResponse</a>

Methods:

- <code title="get /v1/api/api-performance/failure-analysis/{brand_id}">client.API.APIPerformance.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIAPIPerformanceService.FailureAnalysis">FailureAnalysis</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, brandID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIAPIPerformanceFailureAnalysisParams">APIAPIPerformanceFailureAnalysisParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#FailureAnalysisSnapshot">FailureAnalysisSnapshot</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api/api-performance/latest-runs/{brand_id}">client.API.APIPerformance.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIAPIPerformanceService.LatestRuns">LatestRuns</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, brandID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIAPIPerformanceLatestRunsParams">APIAPIPerformanceLatestRunsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#ApiapiPerformanceLatestRunsResponse">ApiapiPerformanceLatestRunsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Prompts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIPromptNewBatchResponse">APIPromptNewBatchResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIPromptBatchResponse">APIPromptBatchResponse</a>

Methods:

- <code title="post /v1/api/prompts/batch">client.API.Prompts.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIPromptService.NewBatch">NewBatch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIPromptNewBatchParams">APIPromptNewBatchParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIPromptNewBatchResponse">APIPromptNewBatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/api/prompts/batch">client.API.Prompts.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIPromptService.Batch">Batch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIPromptBatchParams">APIPromptBatchParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIPromptBatchResponse">APIPromptBatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Jobs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIJobListJobsResponse">APIJobListJobsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIJobListResponse">APIJobListResponse</a>

Methods:

- <code title="get /v1/api/jobs">client.API.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIJobService.ListJobs">ListJobs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIJobListJobsParams">APIJobListJobsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIJobListJobsResponse">APIJobListJobsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api/jobs/{job_id}">client.API.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIJobService.GetJob">GetJob</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#PublicJob">PublicJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api/jobs">client.API.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIJobService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIJobListParams">APIJobListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIJobListResponse">APIJobListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api/jobs/{job_id}">client.API.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#APIJobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go">sapient</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sapient-go#PublicJob">PublicJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
