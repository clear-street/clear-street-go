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

func TestV1InstrumentDataNewsGetNewsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.InstrumentData.News.GetNews(context.TODO(), clearstreet.V1InstrumentDataNewsGetNewsParams{
		ExcludePublishers: clearstreet.String("exclude_publishers"),
		From:              clearstreet.String("from"),
		IncludePublishers: clearstreet.String("include_publishers"),
		InstrumentIDs:     []string{"string"},
		NewsType:          clearstreet.V1InstrumentDataNewsGetNewsParamsNewsTypeNews,
		PageSize:          clearstreet.Int(1),
		PageToken:         clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
		SearchQuery:       clearstreet.String("search_query"),
		Sectors:           []string{"BASIC_MATERIALS"},
		To:                clearstreet.String("to"),
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
