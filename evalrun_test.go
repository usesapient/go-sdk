// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sapient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/usesapient/go-sdk"
	"github.com/usesapient/go-sdk/internal/testutil"
	"github.com/usesapient/go-sdk/option"
)

func TestEvalRunGet(t *testing.T) {
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
	_, err := client.EvalRuns.Get(context.TODO(), "run_id")
	if err != nil {
		var apierr *sapient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEvalRunListWithOptionalParams(t *testing.T) {
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
	_, err := client.EvalRuns.List(context.TODO(), sapient.EvalRunListParams{
		Company: "company",
		Limit:   sapient.Int(1),
		Since:   sapient.String("since"),
	})
	if err != nil {
		var apierr *sapient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEvalRunDiagnoseWithOptionalParams(t *testing.T) {
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
	_, err := client.EvalRuns.Diagnose(context.TODO(), sapient.EvalRunDiagnoseParams{
		Company:         "company",
		Format:          sapient.EvalRunDiagnoseParamsFormatJson,
		IncludeExamples: sapient.Bool(true),
		MaxExamples:     sapient.Int(1),
		Since:           sapient.String("since"),
	})
	if err != nil {
		var apierr *sapient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
