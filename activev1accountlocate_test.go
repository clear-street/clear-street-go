// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/clear-street/clear-street-go"
	"github.com/clear-street/clear-street-go/internal/testutil"
	"github.com/clear-street/clear-street-go/option"
)

func TestActiveV1AccountLocateNewLocateRequest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := clearstreet.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Active.V1.Accounts.Locates.NewLocateRequest(
		context.TODO(),
		0,
		clearstreet.ActiveV1AccountLocateNewLocateRequestParams{
			Body: []clearstreet.ActiveV1AccountLocateNewLocateRequestParamsBody{{
				Quantity:    500,
				Symbol:      "AAPL",
				Comments:    clearstreet.String("Locate for earnings play"),
				ReferenceID: clearstreet.String("my-locate-batch-001"),
			}},
		},
	)
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestActiveV1AccountLocateGetLocateRequestsWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := clearstreet.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Active.V1.Accounts.Locates.GetLocateRequests(
		context.TODO(),
		0,
		clearstreet.ActiveV1AccountLocateGetLocateRequestsParams{
			PageSize:    clearstreet.Int(1),
			PageToken:   clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
			ReferenceID: clearstreet.String("reference_id"),
			Status:      clearstreet.LocateOrderStatusPending,
		},
	)
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestActiveV1AccountLocateUpdateLocateRequest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := clearstreet.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Active.V1.Accounts.Locates.UpdateLocateRequest(
		context.TODO(),
		0,
		clearstreet.ActiveV1AccountLocateUpdateLocateRequestParams{
			Accept: true,
		},
	)
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
