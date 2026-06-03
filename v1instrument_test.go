// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/clear-street/clear-street-go"
	"github.com/clear-street/clear-street-go/internal/testutil"
	"github.com/clear-street/clear-street-go/option"
)

func TestV1InstrumentGetInstrumentByIDWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Instruments.GetInstrumentByID(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		clearstreet.V1InstrumentGetInstrumentByIDParams{
			IncludeOptionsExpiryDates: clearstreet.Bool(true),
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

func TestV1InstrumentGetInstrumentsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Instruments.GetInstruments(context.TODO(), clearstreet.V1InstrumentGetInstrumentsParams{
		EasyToBorrow:        clearstreet.Bool(true),
		InstrumentIDs:       []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		InstrumentType:      clearstreet.V1InstrumentGetInstrumentsParamsInstrumentTypeCommonStock,
		IsLiquidationOnly:   clearstreet.Bool(true),
		IsMarginable:        clearstreet.Bool(true),
		IsPtp:               clearstreet.Bool(true),
		IsShortProhibited:   clearstreet.Bool(true),
		IsThresholdSecurity: clearstreet.Bool(true),
		PageSize:            clearstreet.Int(1),
		PageToken:           clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1InstrumentGetOptionContractsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Instruments.GetOptionContracts(context.TODO(), clearstreet.V1InstrumentGetOptionContractsParams{
		ContractType:           clearstreet.ContractTypeCall,
		Expiry:                 clearstreet.Time(time.Now()),
		PageSize:               clearstreet.Int(1),
		PageToken:              clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
		Underlier:              clearstreet.String("underlier"),
		UnderlyingInstrumentID: clearstreet.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1InstrumentSearchInstrumentsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Instruments.SearchInstruments(context.TODO(), clearstreet.V1InstrumentSearchInstrumentsParams{
		Q:               "q",
		AssetClass:      clearstreet.String("asset_class"),
		Country:         clearstreet.String("country"),
		Currency:        clearstreet.String("currency"),
		IncludeInactive: clearstreet.Bool(true),
		IncludePtp:      clearstreet.Bool(true),
		PageSize:        clearstreet.Int(1),
		PageToken:       clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
