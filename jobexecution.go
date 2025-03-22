// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schedo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/Schedo-go/internal/apijson"
	"github.com/stainless-sdks/Schedo-go/internal/apiquery"
	"github.com/stainless-sdks/Schedo-go/internal/param"
	"github.com/stainless-sdks/Schedo-go/internal/requestconfig"
	"github.com/stainless-sdks/Schedo-go/option"
)

// JobExecutionService contains methods and other services that help with
// interacting with the Schedo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobExecutionService] method instead.
type JobExecutionService struct {
	Options []option.RequestOption
}

// NewJobExecutionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJobExecutionService(opts ...option.RequestOption) (r *JobExecutionService) {
	r = &JobExecutionService{}
	r.Options = opts
	return
}

// Returns a list of executions for a job
func (r *JobExecutionService) List(ctx context.Context, jobID int64, params JobExecutionListParams, opts ...option.RequestOption) (res *[]JobExecution, err error) {
	if params.XAPIEnvironment.Present {
		opts = append(opts, option.WithHeader("X-API-ENVIRONMENT", fmt.Sprintf("%s", params.XAPIEnvironment)))
	}
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("jobs/executions/%v", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Marks pending job execution as complete
func (r *JobExecutionService) Complete(ctx context.Context, executionID int64, body JobExecutionCompleteParams, opts ...option.RequestOption) (res *JobExecution, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("jobs/executions/complete/%v", executionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns list of jobs that must be executed
func (r *JobExecutionService) Poll(ctx context.Context, opts ...option.RequestOption) (res *[]JobExecution, err error) {
	opts = append(r.Options[:], opts...)
	path := "jobs/executions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Streams job executions over websockets
func (r *JobExecutionService) Updates(ctx context.Context, query JobExecutionUpdatesParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "jobs/executions/stream"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type JobExecution struct {
	// ID of the ent.
	ID int64 `json:"id"`
	// Duration holds the value of the "duration" field.
	Duration int64 `json:"duration"`
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
	// Time when execution was picked up by a worker
	PickUpTime string `json:"pick_up_time"`
	// Time when execution started
	StartTime string `json:"start_time"`
	// Execution status (running, completed, failed)
	Status string           `json:"status"`
	JSON   jobExecutionJSON `json:"-"`
}

// jobExecutionJSON contains the JSON metadata for the struct [JobExecution]
type jobExecutionJSON struct {
	ID          apijson.Field
	Duration    apijson.Field
	EndTime     apijson.Field
	Error       apijson.Field
	ExitCode    apijson.Field
	JobCode     apijson.Field
	Output      apijson.Field
	PickUpTime  apijson.Field
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

type JobExecutionListParams struct {
	XAPIEnvironment param.Field[int64] `header:"X-API-ENVIRONMENT,required"`
	// 1
	Cursor param.Field[int64] `query:"cursor"`
	// 1
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [JobExecutionListParams]'s query parameters as `url.Values`.
func (r JobExecutionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type JobExecutionCompleteParams struct {
	Success param.Field[bool]   `json:"success,required"`
	Error   param.Field[string] `json:"error"`
	Output  param.Field[string] `json:"output"`
}

func (r JobExecutionCompleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type JobExecutionUpdatesParams struct {
	// List of Job Codes to stream
	Jobs param.Field[[]string] `query:"jobs"`
}

// URLQuery serializes [JobExecutionUpdatesParams]'s query parameters as
// `url.Values`.
func (r JobExecutionUpdatesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
