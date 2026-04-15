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

func TestActiveV1AccountOrderCancelAllOrdersWithOptionalParams(t *testing.T) {
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
	_, err := client.Active.V1.Accounts.Orders.CancelAllOrders(
		context.TODO(),
		0,
		clearstreet.ActiveV1AccountOrderCancelAllOrdersParams{
			SecurityID:       []string{"string"},
			SecurityIDSource: []string{"string"},
			SecurityType:     clearstreet.ActiveV1AccountOrderCancelAllOrdersParamsSecurityTypeCommonStock,
			Side:             clearstreet.ActiveV1AccountOrderCancelAllOrdersParamsSideBuy,
			Type:             clearstreet.ActiveV1AccountOrderCancelAllOrdersParamsTypeMarket,
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
	_, err := client.Active.V1.Accounts.Orders.GetOrders(
		context.TODO(),
		0,
		clearstreet.ActiveV1AccountOrderGetOrdersParams{
			From:             clearstreet.Time(time.Now()),
			PageSize:         clearstreet.Int(1),
			PageToken:        clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
			SecurityID:       []string{"string"},
			SecurityIDSource: []string{"string"},
			SecurityType:     clearstreet.ActiveV1AccountOrderGetOrdersParamsSecurityTypeCommonStock,
			Status:           []string{"PENDING_NEW"},
			Symbol:           clearstreet.String("symbol"),
			To:               clearstreet.Time(time.Now()),
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
	_, err := client.Active.V1.Accounts.Orders.SubmitOrders(
		context.TODO(),
		0,
		clearstreet.ActiveV1AccountOrderSubmitOrdersParams{
			Body: []clearstreet.ActiveV1AccountOrderSubmitOrdersParamsBodyUnion{{
				OfActiveV1AccountOrderSubmitOrderssBodyNewOrderMultilegRequest: &clearstreet.ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequest{
					Legs: []clearstreet.ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLeg{{
						Ratio: "ratio",
						Security: clearstreet.ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLegSecurityUnion{
							OfString: clearstreet.String("0193bb84-447a-706f-996f-097254663f02"),
						},
						SecurityType:   clearstreet.SecurityTypeOption,
						Side:           clearstreet.SideBuy,
						ID:             clearstreet.String("1"),
						PositionEffect: "OPEN",
					}, {
						Ratio: "ratio",
						Security: clearstreet.ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLegSecurityUnion{
							OfString: clearstreet.String("0193bb84-4db4-78ec-b4fd-cba8be61cf8a"),
						},
						SecurityType:   clearstreet.SecurityTypeOption,
						Side:           clearstreet.SideSell,
						ID:             clearstreet.String("2"),
						PositionEffect: "OPEN",
					}, {
						Ratio: "ratio",
						Security: clearstreet.ActiveV1AccountOrderSubmitOrdersParamsBodyNewOrderMultilegRequestLegSecurityUnion{
							OfString: clearstreet.String("0193bb84-5264-7f20-8fd3-35df82cd6ef0"),
						},
						SecurityType:   clearstreet.SecurityTypeOption,
						Side:           clearstreet.SideBuy,
						ID:             clearstreet.String("3"),
						PositionEffect: "OPEN",
					}},
					OrderType:   clearstreet.OrderTypeLimit,
					TimeInForce: clearstreet.TimeInForceDay,
					ID:          clearstreet.String("my-mleg-ref-20251001-001"),
					LimitPrice:  clearstreet.String("0.50"),
					Quantity:    clearstreet.String("1"),
				},
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
