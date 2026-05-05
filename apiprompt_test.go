// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sapient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/sapient-go"
	"github.com/stainless-sdks/sapient-go/internal/testutil"
	"github.com/stainless-sdks/sapient-go/option"
)

func TestAPIPromptNewBatch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sapient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.API.Prompts.NewBatch(context.TODO(), sapient.APIPromptNewBatchParams{
		Prompts: []sapient.APIPromptNewBatchParamsPrompt{{
			LanguageID:  "language_id",
			RegionID:    "region_id",
			Text:        "text",
			TopicID:     "topic_id",
			PlatformIDs: []string{"string"},
		}},
	})
	if err != nil {
		var apierr *sapient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIPromptBatch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sapient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.API.Prompts.Batch(context.TODO(), sapient.APIPromptBatchParams{
		Prompts: []sapient.APIPromptBatchParamsPrompt{{
			LanguageID:  "language_id",
			RegionID:    "region_id",
			Text:        "text",
			TopicID:     "topic_id",
			PlatformIDs: []string{"string"},
		}},
	})
	if err != nil {
		var apierr *sapient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
