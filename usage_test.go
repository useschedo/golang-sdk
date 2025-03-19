// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schedo_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/Schedo-go"
	"github.com/stainless-sdks/Schedo-go/internal/testutil"
	"github.com/stainless-sdks/Schedo-go/option"
)

func TestUsage(t *testing.T) {
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
	apiKey, err := client.Apikeys.New(context.TODO(), schedo.ApikeyNewParams{
		EnvironmentID: schedo.F(int64(1)),
		Name:          schedo.F("First ApiKey"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", apiKey.ID)
}
