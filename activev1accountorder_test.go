// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clearstreet_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/clear-street-go"
	"github.com/stainless-sdks/clear-street-go/internal/testutil"
	"github.com/stainless-sdks/clear-street-go/option"
)

func TestActiveV1AccountOrderCancelAllOrdersWithOptionalParams(t *testing.T) {
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
	_, err := client.Active.V1.Accounts.Orders.CancelAllOrders(
		context.TODO(),
		0,
		clearstreet.ActiveV1AccountOrderCancelAllOrdersParams{
			SecurityID:       clearstreet.String("security_id"),
			SecurityIDSource: clearstreet.SecurityIDSourceCms,
			SecurityType:     clearstreet.SecurityTypeCommonStock,
			Side:             clearstreet.SideBuy,
			Type:             clearstreet.OrderTypeMarket,
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

func TestActiveV1AccountOrderCancelOrder(t *testing.T) {
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
	_, err := client.Active.V1.Accounts.Orders.CancelOrder(
		context.TODO(),
		"order_id",
		clearstreet.ActiveV1AccountOrderCancelOrderParams{
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

func TestActiveV1AccountOrderGetOrderByID(t *testing.T) {
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
	_, err := client.Active.V1.Accounts.Orders.GetOrderByID(
		context.TODO(),
		"order_id",
		clearstreet.ActiveV1AccountOrderGetOrderByIDParams{
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

func TestActiveV1AccountOrderGetOrdersWithOptionalParams(t *testing.T) {
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
	_, err := client.Active.V1.Accounts.Orders.GetOrders(
		context.TODO(),
		0,
		clearstreet.ActiveV1AccountOrderGetOrdersParams{
			From:             "from",
			To:               "to",
			PageSize:         clearstreet.Int(1),
			PageToken:        clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
			SecurityID:       clearstreet.String("security_id"),
			SecurityIDSource: clearstreet.SecurityIDSourceCms,
			SecurityType:     clearstreet.SecurityTypeCommonStock,
			Status:           clearstreet.OrderStatusPendingNew,
			Symbol:           clearstreet.String("symbol"),
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

func TestActiveV1AccountOrderReplaceOrderWithOptionalParams(t *testing.T) {
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
	_, err := client.Active.V1.Accounts.Orders.ReplaceOrder(
		context.TODO(),
		"order_id",
		clearstreet.ActiveV1AccountOrderReplaceOrderParams{
			AccountID:   0,
			LimitPrice:  clearstreet.String("150.50"),
			Quantity:    clearstreet.String("125"),
			StopPrice:   clearstreet.String("148.00"),
			TimeInForce: clearstreet.TimeInForceDay,
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

func TestActiveV1AccountOrderSubmitOrders(t *testing.T) {
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
	_, err := client.Active.V1.Accounts.Orders.SubmitOrders(
		context.TODO(),
		0,
		clearstreet.ActiveV1AccountOrderSubmitOrdersParams{
			Body: []clearstreet.ActiveV1AccountOrderSubmitOrdersParamsBody{{
				OrderID:          "my-ref-id-20251001-002",
				OrderType:        clearstreet.OrderTypeLimit,
				Quantity:         "25",
				SecurityType:     clearstreet.SecurityTypeCommonStock,
				Side:             clearstreet.SideBuy,
				TimeInForce:      clearstreet.TimeInForceDay,
				ExpireAt:         clearstreet.Time(time.Now()),
				ExtendedHours:    clearstreet.Bool(true),
				LimitPrice:       clearstreet.String("140.50"),
				PositionEffect:   "OPEN",
				SecurityID:       clearstreet.String("AAPL"),
				SecurityIDSource: clearstreet.SecurityIDSourceCms,
				StopPrice:        clearstreet.String("135.00"),
				Strategy: clearstreet.OrderStrategyUnionParam{
					OfSor: &clearstreet.OrderStrategySorParam{
						SorStrategyParam: clearstreet.SorStrategyParam{
							EndAt:   clearstreet.Time(time.Now()),
							StartAt: clearstreet.Time(time.Now()),
							Urgency: clearstreet.UrgencySuperPassive,
						},
						Type: "SOR",
					},
				},
				Symbol: clearstreet.String("AAPL"),
				Venue:  clearstreet.String("CDRG"),
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
