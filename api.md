# Apikeys

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#APIKey">APIKey</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#APIKeyEdges">APIKeyEdges</a>

Methods:

- <code title="post /apikeys">client.Apikeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#ApikeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#ApikeyNewParams">ApikeyNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /apikeys">client.Apikeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#ApikeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([][]<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /apikeys/revoke/{id}">client.Apikeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#ApikeyService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) ([][]<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Environments

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Environment">Environment</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#EnvironmentEdges">EnvironmentEdges</a>

Methods:

- <code title="post /org/environments">client.Environments.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#EnvironmentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#EnvironmentNewParams">EnvironmentNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Environment">Environment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /org/environments">client.Environments.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#EnvironmentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Environment">Environment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /org/environments/{id}">client.Environments.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#EnvironmentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Environment">Environment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Jobs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Job">Job</a>

Methods:

- <code title="get /jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobListParams">JobListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /jobs/definition">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobService.Define">Define</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobDefineParams">JobDefineParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# JobExecution

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobExecution">JobExecution</a>

Methods:

- <code title="get /jobs/executions/{jobId}">client.JobExecution.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobExecutionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#int64">int64</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobExecutionListParams">JobExecutionListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobExecution">JobExecution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /jobs/executions/complete/{executionId}">client.JobExecution.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobExecutionService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, executionID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobExecutionCompleteParams">JobExecutionCompleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobExecution">JobExecution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/executions">client.JobExecution.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobExecutionService.Poll">Poll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobExecution">JobExecution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Org

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Org">Org</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#OrgEdges">OrgEdges</a>

Methods:

- <code title="get /org">client.Org.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#OrgService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Org">Org</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
