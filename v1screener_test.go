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

func TestV1ScreenerGetScreenerWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Screener.GetScreener(context.TODO(), clearstreet.V1ScreenerGetScreenerParams{
		FieldFilter: []string{"string"},
		Filter: map[string]string{
			"foo": "string",
		},
		PageSize:      clearstreet.Int(1),
		PageToken:     clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
		SortBy:        clearstreet.String("sort_by"),
		SortDirection: clearstreet.V1ScreenerGetScreenerParamsSortDirectionAsc,
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ScreenerSearchScreenerWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Screener.SearchScreener(context.TODO(), clearstreet.V1ScreenerSearchScreenerParams{
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
		Filters: []clearstreet.SearchFilterParam{{
			Left: clearstreet.FieldRefParam{
				Name:      "market_cap",
				Lookback:  clearstreet.FieldLookbackOneWeek,
				Period:    clearstreet.FieldPeriodQuarter,
				ValueType: clearstreet.FieldTypeDecimal,
			},
			Op: clearstreet.FilterOpSpecParam{
				Name: clearstreet.FilterOperatorGte,
				Args: []clearstreet.OperatorArg{clearstreet.OperatorArgLeftInclusive},
			},
			Right: []clearstreet.FilterValueParam{{
				Value: clearstreet.FilterValueValueUnionParam{
					OfFloat: clearstreet.Float(1000000000),
				},
				Variable: clearstreet.VariableParam{
					Name:     "today",
					Lookback: clearstreet.FieldLookbackOneWeek,
					Modifier: clearstreet.ModifierParam{
						Args: []clearstreet.ModifierArgUnionParam{{
							OfFloat: clearstreet.Float(30),
						}, {
							OfString: clearstreet.String("DAY"),
						}},
						Name: clearstreet.ModifierOpSub,
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
		SortDirection:     clearstreet.V1ScreenerSearchScreenerParamsSortDirectionAsc,
		Sorts: []clearstreet.V1ScreenerSearchScreenerParamsSort{{
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
