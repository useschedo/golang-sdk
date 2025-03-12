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

func TestApikeyNew(t *testing.T) {
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
	_, err := client.Apikeys.New(context.TODO(), schedo.ApikeyNewParams{
		Name: schedo.F("First ApiKey"),
	})
	if err != nil {
		var apierr *schedo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApikeyList(t *testing.T) {
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
	_, err := client.Apikeys.List(context.TODO(), schedo.ApikeyListParams{
		XAPIEnvironment: schedo.F("X-API-ENVIRONMENT"),
	})
	if err != nil {
		var apierr *schedo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApikeyRevoke(t *testing.T) {
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
	_, err := client.Apikeys.Revoke(context.TODO(), int64(0))
	if err != nil {
		var apierr *schedo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
