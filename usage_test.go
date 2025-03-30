// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schedo_test

import (
	"context"
	"os"
	"testing"

	"github.com/useschedo/golang-sdk"
	"github.com/useschedo/golang-sdk/internal/testutil"
	"github.com/useschedo/golang-sdk/option"
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
		return
	}
	t.Logf("%+v\n", apiKey.ID)
}
