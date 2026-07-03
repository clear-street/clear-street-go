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

func TestV1OrderCancelAllOpenOrdersWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Orders.CancelAllOpenOrders(
		context.TODO(),
		0,
		clearstreet.V1OrderCancelAllOpenOrdersParams{
			InstrumentIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			InstrumentType: clearstreet.V1OrderCancelAllOpenOrdersParamsInstrumentTypeCommonStock,
			Side:           clearstreet.V1OrderCancelAllOpenOrdersParamsSideBuy,
			Type:           clearstreet.V1OrderCancelAllOpenOrdersParamsTypeMarket,
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

func TestV1OrderCancelOpenOrder(t *testing.T) {
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
	_, err := client.V1.Orders.CancelOpenOrder(
		context.TODO(),
		"order_id",
		clearstreet.V1OrderCancelOpenOrderParams{
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

func TestV1OrderGetExecutionsWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Orders.GetExecutions(
		context.TODO(),
		0,
		clearstreet.V1OrderGetExecutionsParams{
			From:          clearstreet.Time(time.Now()),
			InstrumentIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			PageSize:      clearstreet.Int(1),
			PageToken:     clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
			To:            clearstreet.Time(time.Now()),
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

func TestV1OrderGetOrderByID(t *testing.T) {
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
	_, err := client.V1.Orders.GetOrderByID(
		context.TODO(),
		"order_id",
		clearstreet.V1OrderGetOrderByIDParams{
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

func TestV1OrderGetOrdersWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Orders.GetOrders(
		context.TODO(),
		0,
		clearstreet.V1OrderGetOrdersParams{
			From:                    clearstreet.Time(time.Now()),
			InstrumentIDs:           []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			InstrumentType:          clearstreet.V1OrderGetOrdersParamsInstrumentTypeCommonStock,
			OrderIDs:                []string{"string"},
			PageSize:                clearstreet.Int(1),
			PageToken:               clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
			Status:                  []string{"PENDING_NEW"},
			Symbol:                  clearstreet.String("symbol"),
			To:                      clearstreet.Time(time.Now()),
			UnderlyingInstrumentIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
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

func TestV1OrderReplaceOrderWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Orders.ReplaceOrder(
		context.TODO(),
		"order_id",
		clearstreet.V1OrderReplaceOrderParams{
			AccountID:   0,
			LimitPrice:  clearstreet.String("49.00"),
			Quantity:    clearstreet.String("1"),
			StopPrice:   clearstreet.String("52.00"),
			TimeInForce: clearstreet.RequestTimeInForceDay,
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

func TestV1OrderSubmitOrders(t *testing.T) {
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
	_, err := client.V1.Orders.SubmitOrders(
		context.TODO(),
		0,
		clearstreet.V1OrderSubmitOrdersParams{
			Orders: []clearstreet.NewOrderRequestParam{{
				OrderType:          clearstreet.RequestOrderTypeLimit,
				Quantity:           "1",
				Side:               clearstreet.SideBuy,
				TimeInForce:        clearstreet.RequestTimeInForceDay,
				ID:                 clearstreet.String("my-ref-id-20251001-002"),
				ExpiresAt:          clearstreet.Time(time.Now()),
				ExtendedHours:      clearstreet.Bool(true),
				InstrumentID:       clearstreet.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				LimitOffset:        clearstreet.String("0.50"),
				LimitPrice:         clearstreet.String("48.00"),
				PositionEffect:     clearstreet.PositionEffectOpen,
				StopPrice:          clearstreet.String("52.00"),
				Symbol:             clearstreet.String("TSLA"),
				TrailingOffset:     clearstreet.String("2.00"),
				TrailingOffsetType: clearstreet.TrailingOffsetTypePrice,
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
