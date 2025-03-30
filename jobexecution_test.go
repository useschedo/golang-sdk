// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schedo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/Schedo-go"
	"github.com/stainless-sdks/Schedo-go/internal/testutil"
	"github.com/stainless-sdks/Schedo-go/option"
)

func TestJobExecutionListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := schedo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.JobExecution.List(
		context.TODO(),
		int64(0),
		schedo.JobExecutionListParams{
			XAPIEnvironment: schedo.F(int64(0)),
			Cursor:          schedo.F(int64(0)),
			Limit:           schedo.F(int64(0)),
		},
	)
	if err != nil {
		var apierr *schedo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobExecutionCompleteWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := schedo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.JobExecution.Complete(
		context.TODO(),
		int64(0),
		schedo.JobExecutionCompleteParams{
			Success:               schedo.F(true),
			CompleteServerTimeUtc: schedo.F("2023-01-01T00:00:00Z"),
			Error:                 schedo.F("Error message"),
			Output:                schedo.F("Output message"),
			StartServerTimeUtc:    schedo.F("2023-01-01T00:00:00Z"),
		},
	)
	if err != nil {
		var apierr *schedo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobExecutionPoll(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := schedo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.JobExecution.Poll(context.TODO())
	if err != nil {
		var apierr *schedo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
