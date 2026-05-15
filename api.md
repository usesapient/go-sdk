# Status

Response Types:

- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#StatusGetResponse">StatusGetResponse</a>

Methods:

- <code title="get /v1/status">client.Status.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#StatusService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#StatusGetResponse">StatusGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#AuthStatusResponse">AuthStatusResponse</a>

Methods:

- <code title="get /v1/auth/status">client.Auth.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#AuthService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#AuthStatusResponse">AuthStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Prompts

Response Types:

- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptNewResponse">PromptNewResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptGetResponse">PromptGetResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptUpdateResponse">PromptUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptListResponse">PromptListResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptDeleteResponse">PromptDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptEstimateCostResponse">PromptEstimateCostResponse</a>

Methods:

- <code title="post /v1/prompts">client.Prompts.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptNewParams">PromptNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptNewResponse">PromptNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/prompts/{prompt_id}">client.Prompts.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptGetResponse">PromptGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/prompts/{prompt_id}">client.Prompts.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptUpdateParams">PromptUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptUpdateResponse">PromptUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/prompts">client.Prompts.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptListParams">PromptListParams</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptListResponse">PromptListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/prompts/{prompt_id}">client.Prompts.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptDeleteResponse">PromptDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/prompts/estimate-cost">client.Prompts.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptService.EstimateCost">EstimateCost</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptEstimateCostResponse">PromptEstimateCostResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Topics

Response Types:

- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptTopicNewResponse">PromptTopicNewResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptTopicUpdateResponse">PromptTopicUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptTopicListResponse">PromptTopicListResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptTopicDeleteResponse">PromptTopicDeleteResponse</a>

Methods:

- <code title="post /v1/prompts/topics">client.Prompts.Topics.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptTopicService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptTopicNewParams">PromptTopicNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptTopicNewResponse">PromptTopicNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/prompts/topics/{topic_id}">client.Prompts.Topics.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptTopicService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, topicID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptTopicUpdateParams">PromptTopicUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptTopicUpdateResponse">PromptTopicUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/prompts/topics">client.Prompts.Topics.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptTopicService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptTopicListResponse">PromptTopicListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/prompts/topics/{topic_id}">client.Prompts.Topics.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptTopicService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, topicID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptTopicDeleteResponse">PromptTopicDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Platforms

Response Types:

- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptPlatformListResponse">PromptPlatformListResponse</a>

Methods:

- <code title="get /v1/prompts/platforms">client.Prompts.Platforms.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptPlatformService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#PromptPlatformListResponse">PromptPlatformListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# APIPerformance

## Platforms

Response Types:

- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePlatformListResponse">APIPerformancePlatformListResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePlatformEstimateCostResponse">APIPerformancePlatformEstimateCostResponse</a>

Methods:

- <code title="get /v1/api-performance/platforms">client.APIPerformance.Platforms.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePlatformService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePlatformListResponse">APIPerformancePlatformListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api-performance/platforms/estimate-cost">client.APIPerformance.Platforms.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePlatformService.EstimateCost">EstimateCost</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePlatformEstimateCostResponse">APIPerformancePlatformEstimateCostResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Interfaces

Response Types:

- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceInterfaceListResponse">APIPerformanceInterfaceListResponse</a>

Methods:

- <code title="get /v1/api-performance/interfaces">client.APIPerformance.Interfaces.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceInterfaceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceInterfaceListResponse">APIPerformanceInterfaceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## EvaluationConfig

Response Types:

- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceEvaluationConfigGetResponse">APIPerformanceEvaluationConfigGetResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceEvaluationConfigUpdateResponse">APIPerformanceEvaluationConfigUpdateResponse</a>

Methods:

- <code title="get /v1/api-performance/evaluation-config">client.APIPerformance.EvaluationConfig.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceEvaluationConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceEvaluationConfigGetParams">APIPerformanceEvaluationConfigGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceEvaluationConfigGetResponse">APIPerformanceEvaluationConfigGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/api-performance/evaluation-config">client.APIPerformance.EvaluationConfig.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceEvaluationConfigService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceEvaluationConfigUpdateParams">APIPerformanceEvaluationConfigUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceEvaluationConfigUpdateResponse">APIPerformanceEvaluationConfigUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Operations

Response Types:

- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationNewResponse">APIPerformanceOperationNewResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationGetResponse">APIPerformanceOperationGetResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationUpdateResponse">APIPerformanceOperationUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationListResponse">APIPerformanceOperationListResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationDeleteResponse">APIPerformanceOperationDeleteResponse</a>

Methods:

- <code title="post /v1/api-performance/operations">client.APIPerformance.Operations.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationNewParams">APIPerformanceOperationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationNewResponse">APIPerformanceOperationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api-performance/operations/{operation_id}">client.APIPerformance.Operations.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationGetResponse">APIPerformanceOperationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/api-performance/operations/{operation_id}">client.APIPerformance.Operations.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationUpdateParams">APIPerformanceOperationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationUpdateResponse">APIPerformanceOperationUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api-performance/operations">client.APIPerformance.Operations.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationListParams">APIPerformanceOperationListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationListResponse">APIPerformanceOperationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/api-performance/operations/{operation_id}">client.APIPerformance.Operations.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationDeleteResponse">APIPerformanceOperationDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## OperationPrompts

Response Types:

- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationPromptNewResponse">APIPerformanceOperationPromptNewResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationPromptListResponse">APIPerformanceOperationPromptListResponse</a>

Methods:

- <code title="post /v1/api-performance/operations/{operation_id}/prompts">client.APIPerformance.OperationPrompts.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationPromptService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationPromptNewParams">APIPerformanceOperationPromptNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationPromptNewResponse">APIPerformanceOperationPromptNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api-performance/operations/{operation_id}/prompts">client.APIPerformance.OperationPrompts.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationPromptService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*[]<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceOperationPromptListResponse">APIPerformanceOperationPromptListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Runs

Response Types:

- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceRunGetResponse">APIPerformanceRunGetResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceRunListResponse">APIPerformanceRunListResponse</a>

Methods:

- <code title="get /v1/api-performance/runs/{run_id}">client.APIPerformance.Runs.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceRunService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, runID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceRunGetResponse">APIPerformanceRunGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api-performance/runs">client.APIPerformance.Runs.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceRunService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceRunListParams">APIPerformanceRunListParams</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceRunListResponse">APIPerformanceRunListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Prompts

Response Types:

- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePromptGetResponse">APIPerformancePromptGetResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePromptUpdateResponse">APIPerformancePromptUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePromptDeleteResponse">APIPerformancePromptDeleteResponse</a>

Methods:

- <code title="get /v1/api-performance/prompts/{prompt_id}">client.APIPerformance.Prompts.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePromptService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePromptGetResponse">APIPerformancePromptGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/api-performance/prompts/{prompt_id}">client.APIPerformance.Prompts.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePromptService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePromptUpdateParams">APIPerformancePromptUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePromptUpdateResponse">APIPerformancePromptUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/api-performance/prompts/{prompt_id}">client.APIPerformance.Prompts.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePromptService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformancePromptDeleteResponse">APIPerformancePromptDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## UseCases

Response Types:

- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseNewResponse">APIPerformanceUseCaseNewResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseGetResponse">APIPerformanceUseCaseGetResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseUpdateResponse">APIPerformanceUseCaseUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseListResponse">APIPerformanceUseCaseListResponse</a>
- <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseDeleteResponse">APIPerformanceUseCaseDeleteResponse</a>

Methods:

- <code title="post /v1/api-performance/use-cases">client.APIPerformance.UseCases.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseNewParams">APIPerformanceUseCaseNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseNewResponse">APIPerformanceUseCaseNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api-performance/use-cases/{use_case_id}">client.APIPerformance.UseCases.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, useCaseID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseGetResponse">APIPerformanceUseCaseGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/api-performance/use-cases/{use_case_id}">client.APIPerformance.UseCases.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, useCaseID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseUpdateParams">APIPerformanceUseCaseUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseUpdateResponse">APIPerformanceUseCaseUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api-performance/use-cases">client.APIPerformance.UseCases.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseListResponse">APIPerformanceUseCaseListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/api-performance/use-cases/{use_case_id}">client.APIPerformance.UseCases.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, useCaseID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/usesapient/go-sdk">githubcomusesapientgosdk</a>.<a href="https://pkg.go.dev/github.com/usesapient/go-sdk#APIPerformanceUseCaseDeleteResponse">APIPerformanceUseCaseDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
