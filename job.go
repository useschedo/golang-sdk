// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schedo

import (
	"github.com/stainless-sdks/Schedo-go/internal/apijson"
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

type Job struct {
	// ID of the ent.
	ID int64 `json:"id"`
	// Command to execute
	Command string `json:"command"`
	// Time when the job was created
	CreatedAt string `json:"created_at"`
	// Cron expression for job scheduling
	CronExpression string `json:"cron_expression"`
	// Edges holds the relations/edges for other nodes in the graph. The values are
	// being populated by the JobQuery when eager-loading is set.
	Edges *JobEdges `json:"edges"`
	// ID of the user who owns this job
	EnvironmentID int64 `json:"environment_id"`
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
	Edges          apijson.Field
	EnvironmentID  apijson.Field
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

type JobEdges struct {
	// Job execution history
	Executions []interface{} `json:"executions"`
	JSON       jobEdgesJSON  `json:"-"`
}

// jobEdgesJSON contains the JSON metadata for the struct [JobEdges]
type jobEdgesJSON struct {
	Executions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobEdges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobEdgesJSON) RawJSON() string {
	return r.raw
}
