// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/clear-street-go"
	"github.com/stainless-sdks/clear-street-go/internal/testutil"
	"github.com/stainless-sdks/clear-street-go/option"
)

func TestActiveV1InstrumentGetInstrumentByIDWithOptionalParams(t *testing.T) {
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
	_, err := client.Active.V1.Instruments.GetInstrumentByID(
		context.TODO(),
		"security_id",
		clearstreet.ActiveV1InstrumentGetInstrumentByIDParams{
			SecurityIDSource:          clearstreet.SecurityIDSourceCms,
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

func TestActiveV1InstrumentGetInstrumentsWithOptionalParams(t *testing.T) {
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
	_, err := client.Active.V1.Instruments.GetInstruments(context.TODO(), clearstreet.ActiveV1InstrumentGetInstrumentsParams{
		EasyToBorrow:        clearstreet.Bool(true),
		IDFilter:            clearstreet.String("id_filter"),
		IsLiquidationOnly:   clearstreet.Bool(true),
		IsMarginable:        clearstreet.Bool(true),
		IsRestricted:        clearstreet.Bool(true),
		IsShortProhibited:   clearstreet.Bool(true),
		IsThresholdSecurity: clearstreet.Bool(true),
		PageSize:            clearstreet.Int(1),
		PageToken:           clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
		SecurityID:          []string{"string"},
		SecurityIDSource:    []string{"string"},
		SecurityType:        clearstreet.ActiveV1InstrumentGetInstrumentsParamsSecurityTypeCommonStock,
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
