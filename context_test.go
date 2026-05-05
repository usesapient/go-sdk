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

func TestContextGetCompanyWithOptionalParams(t *testing.T) {
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
	_, err := client.Context.GetCompany(
		context.TODO(),
		"company",
		sapient.ContextGetCompanyParams{
			For:    sapient.String("for"),
			Format: sapient.ContextGetCompanyParamsFormatJson,
			Since:  sapient.String("since"),
		},
	)
	if err != nil {
		var apierr *sapient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContextCompanyWithOptionalParams(t *testing.T) {
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
	_, err := client.Context.Company(
		context.TODO(),
		"company",
		sapient.ContextCompanyParams{
			For:    sapient.String("for"),
			Format: sapient.ContextCompanyParamsFormatJson,
			Since:  sapient.String("since"),
		},
	)
	if err != nil {
		var apierr *sapient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
