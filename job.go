// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schedo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/Schedo-go/internal/apijson"
	"github.com/stainless-sdks/Schedo-go/internal/param"
	"github.com/stainless-sdks/Schedo-go/internal/requestconfig"
	"github.com/stainless-sdks/Schedo-go/option"
)

// JobService contains methods and other services that help with interacting with
// the Schedo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobService] method instead.
type JobService struct {
	Options []option.RequestOption
}

// NewJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJobService(opts ...option.RequestOption) (r *JobService) {
	r = &JobService{}
	r.Options = opts
	return
}

// List all jobs
func (r *JobService) List(ctx context.Context, query JobListParams, opts ...option.RequestOption) (res *Job, err error) {
	if query.XAPIEnvironment.Present {
		opts = append(opts, option.WithHeader("X-API-ENVIRONMENT", fmt.Sprintf("%s", query.XAPIEnvironment)))
	}
	opts = append(r.Options[:], opts...)
	path := "jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Marks pending job execution as complete
func (r *JobService) Complete(ctx context.Context, executionID int64, opts ...option.RequestOption) (res *JobExecution, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("jobs/executions/complete/%v", executionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Tries to create a new Job Definition
func (r *JobService) Define(ctx context.Context, body JobDefineParams, opts ...option.RequestOption) (res *JobExecution, err error) {
	opts = append(r.Options[:], opts...)
	path := "jobs/definition"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns list of jobs that must be executed
func (r *JobService) Executions(ctx context.Context, opts ...option.RequestOption) (res *JobExecution, err error) {
	opts = append(r.Options[:], opts...)
	path := "jobs/executions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Job struct {
	// ID of the ent.
	ID int64 `json:"id"`
	// Command to execute
	Command string `json:"command"`
	// Time when the job was created
	CreatedAt string `json:"created_at"`
	// Cron expression for job scheduling
	CronExpression string `json:"cron_expression"`
	// ID of the user who owns this job
	EnvironmentID int64 `json:"environment_id"`
	// Key holds the value of the "key" field.
	Key string `json:"key"`
	// Time when the job was last executed
	LastRunAt string `json:"last_run_at"`
	// Maximum number of retry attempts
	MaxRetries int64 `json:"max_retries"`
	// Additional metadata for the job
	Metadata map[string]interface{} `json:"metadata"`
	// Name of the job
	Name string `json:"name"`
	// Scheduled time for next execution
	NextRunAt string `json:"next_run_at"`
	// Number of retry attempts
	RetryCount int64 `json:"retry_count"`
	// Current job status (pending, running, completed, failed)
	Status string `json:"status"`
	// Maximum execution time before job is terminated
	Timeout string `json:"timeout"`
	// Time when the job was last updated
	UpdatedAt string  `json:"updated_at"`
	JSON      jobJSON `json:"-"`
}

// jobJSON contains the JSON metadata for the struct [Job]
type jobJSON struct {
	ID             apijson.Field
	Command        apijson.Field
	CreatedAt      apijson.Field
	CronExpression apijson.Field
	EnvironmentID  apijson.Field
	Key            apijson.Field
	LastRunAt      apijson.Field
	MaxRetries     apijson.Field
	Metadata       apijson.Field
	Name           apijson.Field
	NextRunAt      apijson.Field
	RetryCount     apijson.Field
	Status         apijson.Field
	Timeout        apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Job) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobJSON) RawJSON() string {
	return r.raw
}

type JobExecution struct {
	// ID of the ent.
	ID int64 `json:"id"`
	// Time when execution completed
	EndTime string `json:"end_time"`
	// Error message if execution failed
	Error string `json:"error"`
	// Exit code of the executed command
	ExitCode int64 `json:"exit_code"`
	// JobCode holds the value of the "job_code" field.
	JobCode string `json:"job_code"`
	// Output of the executed command
	Output string `json:"output"`
	// Time when execution started
	StartTime string `json:"start_time"`
	// Execution status (running, completed, failed)
	Status string           `json:"status"`
	JSON   jobExecutionJSON `json:"-"`
}

// jobExecutionJSON contains the JSON metadata for the struct [JobExecution]
type jobExecutionJSON struct {
	ID          apijson.Field
	EndTime     apijson.Field
	Error       apijson.Field
	ExitCode    apijson.Field
	JobCode     apijson.Field
	Output      apijson.Field
	StartTime   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobExecutionJSON) RawJSON() string {
	return r.raw
}

type JobListParams struct {
	XAPIEnvironment param.Field[int64] `header:"X-API-ENVIRONMENT,required"`
}

type JobDefineParams struct {
	Name       param.Field[string]                 `json:"name,required"`
	Schedule   param.Field[string]                 `json:"schedule,required"`
	MaxRetries param.Field[int64]                  `json:"max_retries"`
	Metadata   param.Field[map[string]interface{}] `json:"metadata"`
	Timeout    param.Field[string]                 `json:"timeout"`
}

func (r JobDefineParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
