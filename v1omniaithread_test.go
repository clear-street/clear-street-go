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

func TestV1OmniAIThreadNewMessageWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.OmniAI.Threads.NewMessage(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		clearstreet.V1OmniAIThreadNewMessageParams{
			AccountID:    19816,
			Text:         "Compare that to AMD.",
			Capabilities: []string{"PREFILL_ORDER"},
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

func TestV1OmniAIThreadNewThreadWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.OmniAI.Threads.NewThread(context.TODO(), clearstreet.V1OmniAIThreadNewThreadParams{
		AccountID:    19816,
		Type:         clearstreet.V1OmniAIThreadNewThreadParamsTypeInstant,
		Capabilities: []string{"PREFILL_ORDER"},
		Target: clearstreet.V1OmniAIThreadNewThreadParamsTarget{
			Ticker: "ticker",
			Type:   "ticker",
		},
		Text:   clearstreet.String("What changed in NVDA today?"),
		Thesis: clearstreet.String("thesis"),
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1OmniAIThreadGetMessagesWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.OmniAI.Threads.GetMessages(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		clearstreet.V1OmniAIThreadGetMessagesParams{
			AccountID: 0,
			PageSize:  clearstreet.Int(1),
			PageToken: clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
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

func TestV1OmniAIThreadGetThreadByID(t *testing.T) {
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
	_, err := client.V1.OmniAI.Threads.GetThreadByID(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		clearstreet.V1OmniAIThreadGetThreadByIDParams{
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

func TestV1OmniAIThreadGetThreadResponse(t *testing.T) {
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
	_, err := client.V1.OmniAI.Threads.GetThreadResponse(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		clearstreet.V1OmniAIThreadGetThreadResponseParams{
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

func TestV1OmniAIThreadGetThreadsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.OmniAI.Threads.GetThreads(context.TODO(), clearstreet.V1OmniAIThreadGetThreadsParams{
		AccountID: 0,
		PageSize:  clearstreet.Int(1),
		PageToken: clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
