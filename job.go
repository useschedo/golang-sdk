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

// Retrieve a job by its ID
func (r *JobService) Get(ctx context.Context, jobID int64, query JobGetParams, opts ...option.RequestOption) (res *Output, err error) {
	if query.XAPIEnvironment.Present {
		opts = append(opts, option.WithHeader("X-API-ENVIRONMENT", fmt.Sprintf("%s", query.XAPIEnvironment)))
	}
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("jobs/%v", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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

// After you delete a job, you can't recover it, but if you have services still
// running with that job reference, they will re-create and re-schedule a new job
// automatically.
func (r *JobService) Delete(ctx context.Context, jobID int64, body JobDeleteParams, opts ...option.RequestOption) (res *string, err error) {
	if body.XAPIEnvironment.Present {
		opts = append(opts, option.WithHeader("X-API-ENVIRONMENT", fmt.Sprintf("%s", body.XAPIEnvironment)))
	}
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("jobs/%v", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Tries to create a new Job Definition
func (r *JobService) Define(ctx context.Context, body JobDefineParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = append(r.Options[:], opts...)
	path := "jobs/definition"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Temporary stops a job from running
func (r *JobService) Pause(ctx context.Context, jobID int64, body JobPauseParams, opts ...option.RequestOption) (res *JobExecution, err error) {
	if body.XAPIEnvironment.Present {
		opts = append(opts, option.WithHeader("X-API-ENVIRONMENT", fmt.Sprintf("%s", body.XAPIEnvironment)))
	}
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("jobs/pause/%v", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Resumes job execution
func (r *JobService) Resume(ctx context.Context, jobID int64, body JobResumeParams, opts ...option.RequestOption) (res *JobExecution, err error) {
	if body.XAPIEnvironment.Present {
		opts = append(opts, option.WithHeader("X-API-ENVIRONMENT", fmt.Sprintf("%s", body.XAPIEnvironment)))
	}
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("jobs/resume/%v", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Immediately triggers a job
func (r *JobService) Trigger(ctx context.Context, jobID int64, body JobTriggerParams, opts ...option.RequestOption) (res *JobExecution, err error) {
	if body.XAPIEnvironment.Present {
		opts = append(opts, option.WithHeader("X-API-ENVIRONMENT", fmt.Sprintf("%s", body.XAPIEnvironment)))
	}
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("jobs/trigger/%v", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type Job struct {
	// ID of the ent.
	ID int64 `json:"id"`
	// Blocking holds the value of the "blocking" field.
	Blocking bool `json:"blocking"`
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
	// Paused holds the value of the "paused" field.
	Paused bool `json:"paused"`
	// Number of retry attempts
	RetryCount int64 `json:"retry_count"`
	// Current job status (pending, running, completed, failed)
	Status string `json:"status"`
	// Maximum execution time before job is terminated
	Timeout string `json:"timeout"`
	// Maximum execution time before job is terminated
	TimeoutSeconds int64 `json:"timeout_seconds"`
	// Time when the job was last updated
	UpdatedAt string  `json:"updated_at"`
	JSON      jobJSON `json:"-"`
}

// jobJSON contains the JSON metadata for the struct [Job]
type jobJSON struct {
	ID             apijson.Field
	Blocking       apijson.Field
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
	Paused         apijson.Field
	RetryCount     apijson.Field
	Status         apijson.Field
	Timeout        apijson.Field
	TimeoutSeconds apijson.Field
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

type Output struct {
	Job     Job          `json:"job"`
	LastRun JobExecution `json:"last_run"`
	JSON    outputJSON   `json:"-"`
}

// outputJSON contains the JSON metadata for the struct [Output]
type outputJSON struct {
	Job         apijson.Field
	LastRun     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Output) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outputJSON) RawJSON() string {
	return r.raw
}

type JobGetParams struct {
	XAPIEnvironment param.Field[int64] `header:"X-API-ENVIRONMENT,required"`
}

type JobListParams struct {
	XAPIEnvironment param.Field[int64] `header:"X-API-ENVIRONMENT,required"`
}

type JobDeleteParams struct {
	XAPIEnvironment param.Field[int64] `header:"X-API-ENVIRONMENT,required"`
}

type JobDefineParams struct {
	Name           param.Field[string]                 `json:"name,required"`
	Schedule       param.Field[string]                 `json:"schedule,required"`
	Blocking       param.Field[bool]                   `json:"blocking"`
	Metadata       param.Field[map[string]interface{}] `json:"metadata"`
	TimeoutSeconds param.Field[int64]                  `json:"timeout_seconds"`
}

func (r JobDefineParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type JobPauseParams struct {
	XAPIEnvironment param.Field[int64] `header:"X-API-ENVIRONMENT,required"`
}

type JobResumeParams struct {
	XAPIEnvironment param.Field[int64] `header:"X-API-ENVIRONMENT,required"`
}

type JobTriggerParams struct {
	XAPIEnvironment param.Field[int64] `header:"X-API-ENVIRONMENT,required"`
}
