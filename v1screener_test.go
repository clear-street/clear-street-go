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

func TestV1ScreenerNewScreenerWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Screener.NewScreener(context.TODO(), clearstreet.V1ScreenerNewScreenerParams{
		Columns: []clearstreet.FieldRefParam{{
			Name:      "market_cap",
			Lookback:  clearstreet.FieldLookbackOneDay,
			Period:    clearstreet.FieldPeriodQuarter,
			ValueType: clearstreet.FieldTypeDecimal,
		}},
		FieldFilter: []clearstreet.FieldRefParam{{
			Name:      "market_cap",
			Lookback:  clearstreet.FieldLookbackOneDay,
			Period:    clearstreet.FieldPeriodQuarter,
			ValueType: clearstreet.FieldTypeDecimal,
		}},
		Filters: []clearstreet.SearchFilterParam{{
			Left: clearstreet.FieldRefParam{
				Name:      "market_cap",
				Lookback:  clearstreet.FieldLookbackOneDay,
				Period:    clearstreet.FieldPeriodQuarter,
				ValueType: clearstreet.FieldTypeDecimal,
			},
			Op: clearstreet.FilterOpSpecParam{
				Name: clearstreet.FilterOperatorLessThan,
				Args: []clearstreet.OperatorArg{clearstreet.OperatorArgLeftInclusive},
			},
			Right: []clearstreet.FilterValueParam{{
				Value: clearstreet.FilterValueValueUnionParam{
					OfFloat: clearstreet.Float(1000000000),
				},
				Variable: clearstreet.VariableParam{
					Name:     "today",
					Lookback: clearstreet.FieldLookbackOneDay,
					Modifier: clearstreet.ModifierParam{
						Args: []clearstreet.ModifierArgUnionParam{{
							OfFloat: clearstreet.Float(30),
						}, {
							OfString: clearstreet.String("DAY"),
						}},
						Name: clearstreet.ModifierOpSubtract,
					},
					Period: clearstreet.FieldPeriodQuarter,
				},
			}},
		}},
		Name: clearstreet.String("name"),
		Sorts: []clearstreet.SortSpecParam{{
			Field: clearstreet.FieldRefParam{
				Name:      "market_cap",
				Lookback:  clearstreet.FieldLookbackOneDay,
				Period:    clearstreet.FieldPeriodQuarter,
				ValueType: clearstreet.FieldTypeDecimal,
			},
			Direction: clearstreet.SortDirectionDesc,
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

func TestV1ScreenerDeleteScreener(t *testing.T) {
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
	err := client.V1.Screener.DeleteScreener(context.TODO(), "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ScreenerGetScreenerByID(t *testing.T) {
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
	_, err := client.V1.Screener.GetScreenerByID(context.TODO(), "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ScreenerGetScreeners(t *testing.T) {
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
	_, err := client.V1.Screener.GetScreeners(context.TODO())
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ScreenerReplaceScreenerWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Screener.ReplaceScreener(
		context.TODO(),
		"550e8400-e29b-41d4-a716-446655440000",
		clearstreet.V1ScreenerReplaceScreenerParams{
			Columns: []clearstreet.FieldRefParam{{
				Name:      "market_cap",
				Lookback:  clearstreet.FieldLookbackOneDay,
				Period:    clearstreet.FieldPeriodQuarter,
				ValueType: clearstreet.FieldTypeDecimal,
			}},
			FieldFilter: []clearstreet.FieldRefParam{{
				Name:      "market_cap",
				Lookback:  clearstreet.FieldLookbackOneDay,
				Period:    clearstreet.FieldPeriodQuarter,
				ValueType: clearstreet.FieldTypeDecimal,
			}},
			Filters: []clearstreet.SearchFilterParam{{
				Left: clearstreet.FieldRefParam{
					Name:      "market_cap",
					Lookback:  clearstreet.FieldLookbackOneDay,
					Period:    clearstreet.FieldPeriodQuarter,
					ValueType: clearstreet.FieldTypeDecimal,
				},
				Op: clearstreet.FilterOpSpecParam{
					Name: clearstreet.FilterOperatorLessThan,
					Args: []clearstreet.OperatorArg{clearstreet.OperatorArgLeftInclusive},
				},
				Right: []clearstreet.FilterValueParam{{
					Value: clearstreet.FilterValueValueUnionParam{
						OfFloat: clearstreet.Float(1000000000),
					},
					Variable: clearstreet.VariableParam{
						Name:     "today",
						Lookback: clearstreet.FieldLookbackOneDay,
						Modifier: clearstreet.ModifierParam{
							Args: []clearstreet.ModifierArgUnionParam{{
								OfFloat: clearstreet.Float(30),
							}, {
								OfString: clearstreet.String("DAY"),
							}},
							Name: clearstreet.ModifierOpSubtract,
						},
						Period: clearstreet.FieldPeriodQuarter,
					},
				}},
			}},
			Name: clearstreet.String("name"),
			Sorts: []clearstreet.SortSpecParam{{
				Field: clearstreet.FieldRefParam{
					Name:      "market_cap",
					Lookback:  clearstreet.FieldLookbackOneDay,
					Period:    clearstreet.FieldPeriodQuarter,
					ValueType: clearstreet.FieldTypeDecimal,
				},
				Direction: clearstreet.SortDirectionDesc,
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

func TestV1ScreenerSearchScreenerWithOptionalParams(t *testing.T) {
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
		Columns: []clearstreet.FieldRefParam{{
			Name:      "market_cap",
			Lookback:  clearstreet.FieldLookbackOneDay,
			Period:    clearstreet.FieldPeriodQuarter,
			ValueType: clearstreet.FieldTypeDecimal,
		}, {
			Name:      "price",
			Lookback:  clearstreet.FieldLookbackOneDay,
			Period:    clearstreet.FieldPeriodQuarter,
			ValueType: clearstreet.FieldTypeDecimal,
		}, {
			Name:      "volume",
			Lookback:  clearstreet.FieldLookbackOneDay,
			Period:    clearstreet.FieldPeriodQuarter,
			ValueType: clearstreet.FieldTypeDecimal,
		}},
		FieldFilter: []clearstreet.FieldRefParam{{
			Name:      "market_cap",
			Lookback:  clearstreet.FieldLookbackOneDay,
			Period:    clearstreet.FieldPeriodQuarter,
			ValueType: clearstreet.FieldTypeDecimal,
		}},
		Filters: []clearstreet.SearchFilterParam{{
			Left: clearstreet.FieldRefParam{
				Name:      "market_cap",
				Lookback:  clearstreet.FieldLookbackOneDay,
				Period:    clearstreet.FieldPeriodQuarter,
				ValueType: clearstreet.FieldTypeDecimal,
			},
			Op: clearstreet.FilterOpSpecParam{
				Name: clearstreet.FilterOperatorLessThan,
				Args: []clearstreet.OperatorArg{clearstreet.OperatorArgLeftInclusive},
			},
			Right: []clearstreet.FilterValueParam{{
				Value: clearstreet.FilterValueValueUnionParam{
					OfFloat: clearstreet.Float(1000000000),
				},
				Variable: clearstreet.VariableParam{
					Name:     "today",
					Lookback: clearstreet.FieldLookbackOneDay,
					Modifier: clearstreet.ModifierParam{
						Args: []clearstreet.ModifierArgUnionParam{{
							OfFloat: clearstreet.Float(30),
						}, {
							OfString: clearstreet.String("DAY"),
						}},
						Name: clearstreet.ModifierOpSubtract,
					},
					Period: clearstreet.FieldPeriodQuarter,
				},
			}},
		}},
		PageSize:          clearstreet.Int(25),
		PageToken:         clearstreet.String("U3RhaW5sZXNzIHJvY2tz"),
		SortCaseSensitive: clearstreet.Bool(true),
		Sorts: []clearstreet.SortSpecParam{{
			Field: clearstreet.FieldRefParam{
				Name:      "market_cap",
				Lookback:  clearstreet.FieldLookbackOneDay,
				Period:    clearstreet.FieldPeriodQuarter,
				ValueType: clearstreet.FieldTypeDecimal,
			},
			Direction: clearstreet.SortDirectionDesc,
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
