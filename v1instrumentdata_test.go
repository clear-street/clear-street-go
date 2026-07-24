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

func TestV1InstrumentDataGetAllInstrumentEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.InstrumentData.GetAllInstrumentEvents(context.TODO(), clearstreet.V1InstrumentDataGetAllInstrumentEventsParams{
		EventTypes:    []clearstreet.AllEventsEventType{clearstreet.AllEventsEventTypeEarnings},
		FromDate:      clearstreet.String("from_date"),
		InstrumentIDs: []clearstreet.InstrumentIDOrSymbol{"x"},
		ToDate:        clearstreet.String("to_date"),
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1InstrumentDataGetInstrumentAnalystConsensusWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.InstrumentData.GetInstrumentAnalystConsensus(
		context.TODO(),
		"x",
		clearstreet.V1InstrumentDataGetInstrumentAnalystConsensusParams{
			From: clearstreet.Time(time.Now()),
			To:   clearstreet.Time(time.Now()),
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

func TestV1InstrumentDataGetInstrumentBalanceSheetStatementsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.InstrumentData.GetInstrumentBalanceSheetStatements(
		context.TODO(),
		"x",
		clearstreet.V1InstrumentDataGetInstrumentBalanceSheetStatementsParams{
			FromDate:  clearstreet.String("from_date"),
			PageSize:  clearstreet.Int(1),
			PageToken: clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
			ToDate:    clearstreet.String("to_date"),
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

func TestV1InstrumentDataGetInstrumentCashFlowStatementsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.InstrumentData.GetInstrumentCashFlowStatements(
		context.TODO(),
		"x",
		clearstreet.V1InstrumentDataGetInstrumentCashFlowStatementsParams{
			FromDate:  clearstreet.String("from_date"),
			PageSize:  clearstreet.Int(1),
			PageToken: clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
			ToDate:    clearstreet.String("to_date"),
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

func TestV1InstrumentDataGetInstrumentEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.InstrumentData.GetInstrumentEvents(
		context.TODO(),
		"x",
		clearstreet.V1InstrumentDataGetInstrumentEventsParams{
			FromDate: clearstreet.String("from_date"),
			ToDate:   clearstreet.String("to_date"),
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

func TestV1InstrumentDataGetInstrumentFundamentals(t *testing.T) {
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
	_, err := client.V1.InstrumentData.GetInstrumentFundamentals(context.TODO(), "x")
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1InstrumentDataGetInstrumentIncomeStatementsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.InstrumentData.GetInstrumentIncomeStatements(
		context.TODO(),
		"x",
		clearstreet.V1InstrumentDataGetInstrumentIncomeStatementsParams{
			FromDate:  clearstreet.String("from_date"),
			PageSize:  clearstreet.Int(1),
			PageToken: clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
			ToDate:    clearstreet.String("to_date"),
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
