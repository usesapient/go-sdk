// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomusesapientgosdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/usesapient/go-sdk"
	"github.com/usesapient/go-sdk/internal/testutil"
	"github.com/usesapient/go-sdk/option"
)

func TestAPIPerformanceEvaluationConfigGetWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomusesapientgosdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.APIPerformance.EvaluationConfig.Get(context.TODO(), githubcomusesapientgosdk.APIPerformanceEvaluationConfigGetParams{
		InterfaceID: githubcomusesapientgosdk.String("interface_id"),
	})
	if err != nil {
		var apierr *githubcomusesapientgosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIPerformanceEvaluationConfigUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomusesapientgosdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.APIPerformance.EvaluationConfig.Update(context.TODO(), githubcomusesapientgosdk.APIPerformanceEvaluationConfigUpdateParams{
		APIBaseURL: githubcomusesapientgosdk.String("api_base_url"),
		EnvVars: map[string]string{
			"foo": "string",
		},
		EvalTypes:    []string{"string"},
		Framework:    githubcomusesapientgosdk.String("framework"),
		InterfaceID:  githubcomusesapientgosdk.String("interface_id"),
		OperationIDs: []string{"string"},
		Platforms:    []string{"string"},
		UseCaseIDs:   []string{"string"},
	})
	if err != nil {
		var apierr *githubcomusesapientgosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
