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

func TestAPIPerformanceOperationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.APIPerformance.Operations.New(context.TODO(), githubcomusesapientgosdk.APIPerformanceOperationNewParams{
		Method:             "x",
		Path:               "x",
		CategoryName:       githubcomusesapientgosdk.String("category_name"),
		CategorySlug:       githubcomusesapientgosdk.String("category_slug"),
		Description:        githubcomusesapientgosdk.String("description"),
		InterfaceID:        githubcomusesapientgosdk.String("interface_id"),
		OpenAPIOperationID: githubcomusesapientgosdk.String("openapi_operation_id"),
	})
	if err != nil {
		var apierr *githubcomusesapientgosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIPerformanceOperationGet(t *testing.T) {
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
	_, err := client.APIPerformance.Operations.Get(context.TODO(), "operation_id")
	if err != nil {
		var apierr *githubcomusesapientgosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIPerformanceOperationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.APIPerformance.Operations.Update(
		context.TODO(),
		"operation_id",
		githubcomusesapientgosdk.APIPerformanceOperationUpdateParams{
			CategoryName:       githubcomusesapientgosdk.String("category_name"),
			CategorySlug:       githubcomusesapientgosdk.String("category_slug"),
			Description:        githubcomusesapientgosdk.String("description"),
			InterfaceID:        githubcomusesapientgosdk.String("interface_id"),
			Method:             githubcomusesapientgosdk.String("x"),
			OpenAPIOperationID: githubcomusesapientgosdk.String("openapi_operation_id"),
			Path:               githubcomusesapientgosdk.String("x"),
		},
	)
	if err != nil {
		var apierr *githubcomusesapientgosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIPerformanceOperationListWithOptionalParams(t *testing.T) {
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
	_, err := client.APIPerformance.Operations.List(context.TODO(), githubcomusesapientgosdk.APIPerformanceOperationListParams{
		CategoryID:  githubcomusesapientgosdk.String("category_id"),
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

func TestAPIPerformanceOperationDelete(t *testing.T) {
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
	_, err := client.APIPerformance.Operations.Delete(context.TODO(), "operation_id")
	if err != nil {
		var apierr *githubcomusesapientgosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
