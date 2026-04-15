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

func TestActiveV1InstrumentEventGetAllInstrumentEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.Active.V1.Instruments.Events.GetAllInstrumentEvents(context.TODO(), clearstreet.ActiveV1InstrumentEventGetAllInstrumentEventsParams{
		EventTypes:       []clearstreet.AllEventsEventType{clearstreet.AllEventsEventTypeEarnings},
		FromDate:         clearstreet.String("from_date"),
		InstrumentIDs:    []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		SecurityID:       []string{"string"},
		SecurityIDSource: []string{"string"},
		ToDate:           clearstreet.String("to_date"),
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestActiveV1InstrumentEventGetInstrumentEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.Active.V1.Instruments.Events.GetInstrumentEvents(
		context.TODO(),
		"security_id",
		clearstreet.ActiveV1InstrumentEventGetInstrumentEventsParams{
			SecurityIDSource: clearstreet.SecurityIDSourceCms,
			FromDate:         clearstreet.String("from_date"),
			ToDate:           clearstreet.String("to_date"),
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
