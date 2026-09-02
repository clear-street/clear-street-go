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

func TestV1PrivateMarketNewIoiWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.PrivateMarkets.NewIoi(context.TODO(), clearstreet.V1PrivateMarketNewIoiParams{
		AccountID:      0,
		NotionalAmount: "100000.00",
		OfferingID:     "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		NdaAcceptance: clearstreet.V1PrivateMarketNewIoiParamsNdaAcceptance{
			Accepted:           true,
			AgreementID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			AuthorityConfirmed: true,
		},
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1PrivateMarketDeleteIoi(t *testing.T) {
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
	err := client.V1.PrivateMarkets.DeleteIoi(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		clearstreet.V1PrivateMarketDeleteIoiParams{
			AccountID: 0,
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

func TestV1PrivateMarketGetCompanyByID(t *testing.T) {
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
	_, err := client.V1.PrivateMarkets.GetCompanyByID(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		clearstreet.V1PrivateMarketGetCompanyByIDParams{
			AccountID: 0,
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

func TestV1PrivateMarketGetIois(t *testing.T) {
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
	_, err := client.V1.PrivateMarkets.GetIois(context.TODO(), clearstreet.V1PrivateMarketGetIoisParams{
		AccountID: 0,
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1PrivateMarketGetSpvByID(t *testing.T) {
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
	_, err := client.V1.PrivateMarkets.GetSpvByID(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		clearstreet.V1PrivateMarketGetSpvByIDParams{
			AccountID: 0,
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

func TestV1PrivateMarketUpdateIoiWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.PrivateMarkets.UpdateIoi(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		clearstreet.V1PrivateMarketUpdateIoiParams{
			AccountID:      0,
			NotionalAmount: "125000.00",
			NdaAcceptance: clearstreet.V1PrivateMarketUpdateIoiParamsNdaAcceptance{
				Accepted:           true,
				AgreementID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				AuthorityConfirmed: true,
			},
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
