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

func TestActiveV1CalendarDividendGetDividendsCalendar(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Active.V1.Calendars.Dividends.GetDividendsCalendar(context.TODO(), clearstreet.ActiveV1CalendarDividendGetDividendsCalendarParams{
		FromDate: "from_date",
		ToDate:   "to_date",
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
