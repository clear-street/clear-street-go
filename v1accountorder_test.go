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

func TestV1AccountOrderCancelAllOpenOrdersWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Accounts.Orders.CancelAllOpenOrders(
		context.TODO(),
		0,
		clearstreet.V1AccountOrderCancelAllOpenOrdersParams{
			InstrumentIDs:  []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			InstrumentType: clearstreet.V1AccountOrderCancelAllOpenOrdersParamsInstrumentTypeCommonStock,
			Side:           clearstreet.V1AccountOrderCancelAllOpenOrdersParamsSideBuy,
			Type:           clearstreet.V1AccountOrderCancelAllOpenOrdersParamsTypeMarket,
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

func TestV1AccountOrderCancelOpenOrder(t *testing.T) {
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
	_, err := client.V1.Accounts.Orders.CancelOpenOrder(
		context.TODO(),
		"order_id",
		clearstreet.V1AccountOrderCancelOpenOrderParams{
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

func TestV1AccountOrderGetOrderByID(t *testing.T) {
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
	_, err := client.V1.Accounts.Orders.GetOrderByID(
		context.TODO(),
		"order_id",
		clearstreet.V1AccountOrderGetOrderByIDParams{
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

func TestV1AccountOrderGetOrdersWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Accounts.Orders.GetOrders(
		context.TODO(),
		0,
		clearstreet.V1AccountOrderGetOrdersParams{
			From:                    clearstreet.Time(time.Now()),
			InstrumentIDs:           []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			InstrumentType:          clearstreet.V1AccountOrderGetOrdersParamsInstrumentTypeCommonStock,
			PageSize:                clearstreet.Int(1),
			PageToken:               clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
			Status:                  []string{"PENDING_NEW"},
			Symbol:                  clearstreet.String("symbol"),
			To:                      clearstreet.Time(time.Now()),
			UnderlyingInstrumentIDs: clearstreet.String("underlying_instrument_ids"),
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

func TestV1AccountOrderReplaceOrderWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Accounts.Orders.ReplaceOrder(
		context.TODO(),
		"order_id",
		clearstreet.V1AccountOrderReplaceOrderParams{
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

func TestV1AccountOrderSubmitOrders(t *testing.T) {
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
	_, err := client.V1.Accounts.Orders.SubmitOrders(
		context.TODO(),
		0,
		clearstreet.V1AccountOrderSubmitOrdersParams{
			Orders: []clearstreet.V1AccountOrderSubmitOrdersParamsOrderUnion{{
				OfV1AccountOrderSubmitOrderssOrderNewOrderMultilegRequest: &clearstreet.V1AccountOrderSubmitOrdersParamsOrderNewOrderMultilegRequest{
					Legs: []clearstreet.V1AccountOrderSubmitOrdersParamsOrderNewOrderMultilegRequestLeg{{
						InstrumentType: clearstreet.SecurityTypeOption,
						Ratio:          "ratio",
						Security:       "0193bb84-447a-706f-996f-097254663f02",
						Side:           clearstreet.SideBuy,
						ID:             clearstreet.String("1"),
						PositionEffect: "OPEN",
					}, {
						InstrumentType: clearstreet.SecurityTypeOption,
						Ratio:          "ratio",
						Security:       "0193bb84-4db4-78ec-b4fd-cba8be61cf8a",
						Side:           clearstreet.SideSell,
						ID:             clearstreet.String("2"),
						PositionEffect: "OPEN",
					}, {
						InstrumentType: clearstreet.SecurityTypeOption,
						Ratio:          "ratio",
						Security:       "0193bb84-5264-7f20-8fd3-35df82cd6ef0",
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
