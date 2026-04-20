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

func TestActiveV1ScreenerGetScreenerWithOptionalParams(t *testing.T) {
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
	_, err := client.Active.V1.Screener.GetScreener(context.TODO(), clearstreet.ActiveV1ScreenerGetScreenerParams{
		FieldFilter: []string{"string"},
		Filter: map[string]string{
			"foo": "string",
		},
		PageSize:      clearstreet.Int(1),
		PageToken:     clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
		SortBy:        clearstreet.String("sort_by"),
		SortDirection: clearstreet.ActiveV1ScreenerGetScreenerParamsSortDirectionAsc,
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestActiveV1ScreenerSearchScreenerWithOptionalParams(t *testing.T) {
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
	_, err := client.Active.V1.Screener.SearchScreener(context.TODO(), clearstreet.ActiveV1ScreenerSearchScreenerParams{
		FieldFilter: []clearstreet.FieldRefParam{{
			Name:      "market_cap",
			Lookback:  clearstreet.FieldLookbackOneWeek,
			Period:    clearstreet.FieldPeriodQuarter,
			ValueType: clearstreet.FieldTypeDecimal,
		}, {
			Name:      "price",
			Lookback:  clearstreet.FieldLookbackOneWeek,
			Period:    clearstreet.FieldPeriodQuarter,
			ValueType: clearstreet.FieldTypeDecimal,
		}, {
			Name:      "volume",
			Lookback:  clearstreet.FieldLookbackOneWeek,
			Period:    clearstreet.FieldPeriodQuarter,
			ValueType: clearstreet.FieldTypeDecimal,
		}},
		Filters: []clearstreet.ActiveV1ScreenerSearchScreenerParamsFilter{{
			Left: clearstreet.FieldRefParam{
				Name:      "market_cap",
				Lookback:  clearstreet.FieldLookbackOneWeek,
				Period:    clearstreet.FieldPeriodQuarter,
				ValueType: clearstreet.FieldTypeDecimal,
			},
			Op: clearstreet.ActiveV1ScreenerSearchScreenerParamsFilterOp{
				Name: "GTE",
				Args: []string{"LEFT_INCLUSIVE"},
			},
			Right: []clearstreet.ActiveV1ScreenerSearchScreenerParamsFilterRight{{
				Value: clearstreet.ActiveV1ScreenerSearchScreenerParamsFilterRightValueUnion{
					OfFloat: clearstreet.Float(1000000000),
				},
				Variable: clearstreet.ActiveV1ScreenerSearchScreenerParamsFilterRightVariable{
					Name:     "today",
					Lookback: clearstreet.FieldLookbackOneWeek,
					Modifier: clearstreet.ActiveV1ScreenerSearchScreenerParamsFilterRightVariableModifier{
						Args: []clearstreet.ActiveV1ScreenerSearchScreenerParamsFilterRightVariableModifierArgUnion{{
							OfFloat: clearstreet.Float(30),
						}, {
							OfString: clearstreet.String("DAY"),
						}},
						Name: "SUB",
					},
					Period: clearstreet.FieldPeriodQuarter,
				},
			}},
		}},
		PageSize:  clearstreet.Int(25),
		PageToken: clearstreet.String("page_token"),
		SortBy: clearstreet.FieldRefParam{
			Name:      "market_cap",
			Lookback:  clearstreet.FieldLookbackOneWeek,
			Period:    clearstreet.FieldPeriodQuarter,
			ValueType: clearstreet.FieldTypeDecimal,
		},
		SortCaseSensitive: clearstreet.Bool(true),
		SortDirection:     clearstreet.ActiveV1ScreenerSearchScreenerParamsSortDirectionAsc,
		Sorts: []clearstreet.ActiveV1ScreenerSearchScreenerParamsSort{{
			Field: clearstreet.FieldRefParam{
				Name:      "market_cap",
				Lookback:  clearstreet.FieldLookbackOneWeek,
				Period:    clearstreet.FieldPeriodQuarter,
				ValueType: clearstreet.FieldTypeDecimal,
			},
			Direction: "DESC",
		}},
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
