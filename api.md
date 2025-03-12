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

- <code title="get /org/environments">client.Environments.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#EnvironmentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Environment">Environment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /org/environments/{id}">client.Environments.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#EnvironmentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Environment">Environment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /org/environments">client.Environments.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#EnvironmentService.Created">Created</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#EnvironmentCreatedParams">EnvironmentCreatedParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Environment">Environment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Jobs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Job">Job</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobEdges">JobEdges</a>

Methods:

- <code title="get /jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#JobListParams">JobListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Org

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Org">Org</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#OrgEdges">OrgEdges</a>

Methods:

- <code title="get /org">client.Org.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#OrgService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go">schedo</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/Schedo-go#Org">Org</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
