// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schedo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/Schedo-go/internal/apijson"
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

// Marks pending job execution as complete
func (r *JobExecutionService) Complete(ctx context.Context, executionID int64, opts ...option.RequestOption) (res *JobExecution, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("jobs/executions/complete/%v", executionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Returns list of jobs that must be executed
func (r *JobExecutionService) Poll(ctx context.Context, opts ...option.RequestOption) (res *JobExecution, err error) {
	opts = append(r.Options[:], opts...)
	path := "jobs/executions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
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
