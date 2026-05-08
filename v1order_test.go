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

func TestV1OrderGetOrderByID(t *testing.T) {
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
	_, err := client.V1.Orders.GetOrders(
		context.TODO(),
		0,
		clearstreet.V1OrderGetOrdersParams{
			From:                    clearstreet.Time(time.Now()),
			InstrumentIDs:           []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			InstrumentType:          clearstreet.V1OrderGetOrdersParamsInstrumentTypeCommonStock,
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

func TestV1OrderReplaceOrderWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Orders.ReplaceOrder(
		context.TODO(),
		"order_id",
		clearstreet.V1OrderReplaceOrderParams{
			AccountID:   0,
			LimitPrice:  clearstreet.String("150.50"),
			Quantity:    clearstreet.String("125"),
			StopPrice:   clearstreet.String("148.00"),
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
	_, err := client.V1.Orders.SubmitOrders(
		context.TODO(),
		0,
		clearstreet.V1OrderSubmitOrdersParams{
			Orders: []clearstreet.V1OrderSubmitOrdersParamsOrderUnion{{
				OfV1OrderSubmitOrderssOrderNewOrderMultilegRequest: &clearstreet.V1OrderSubmitOrdersParamsOrderNewOrderMultilegRequest{
					Legs: []clearstreet.V1OrderSubmitOrdersParamsOrderNewOrderMultilegRequestLeg{{
						InstrumentType: clearstreet.SecurityTypeOption,
						Ratio:          "ratio",
						Security:       "0193bb84-447a-706f-996f-097254663f02",
						Side:           clearstreet.SideBuy,
						ID:             clearstreet.String("1"),
						PositionEffect: clearstreet.PositionEffectOpen,
					}, {
						InstrumentType: clearstreet.SecurityTypeOption,
						Ratio:          "ratio",
						Security:       "0193bb84-4db4-78ec-b4fd-cba8be61cf8a",
						Side:           clearstreet.SideSell,
						ID:             clearstreet.String("2"),
						PositionEffect: clearstreet.PositionEffectOpen,
					}, {
						InstrumentType: clearstreet.SecurityTypeOption,
						Ratio:          "ratio",
						Security:       "0193bb84-5264-7f20-8fd3-35df82cd6ef0",
						Side:           clearstreet.SideBuy,
						ID:             clearstreet.String("3"),
						PositionEffect: clearstreet.PositionEffectOpen,
					}},
					OrderType:   clearstreet.RequestOrderTypeLimit,
					TimeInForce: clearstreet.RequestTimeInForceDay,
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
