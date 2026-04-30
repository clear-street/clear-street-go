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

func TestActiveV1SavedScreenerNewScreenerWithOptionalParams(t *testing.T) {
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
	_, err := client.Active.V1.SavedScreeners.NewScreener(context.TODO(), clearstreet.ActiveV1SavedScreenerNewScreenerParams{
		FieldFilter: []clearstreet.FieldRefParam{{
			Name:      "market_cap",
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
		Name: clearstreet.String("name"),
		SortBy: clearstreet.FieldRefParam{
			Name:      "market_cap",
			Lookback:  clearstreet.FieldLookbackOneWeek,
			Period:    clearstreet.FieldPeriodQuarter,
			ValueType: clearstreet.FieldTypeDecimal,
		},
		SortDirection: clearstreet.ActiveV1SavedScreenerNewScreenerParamsSortDirectionAsc,
	})
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestActiveV1SavedScreenerDeleteScreener(t *testing.T) {
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
	err := client.Active.V1.SavedScreeners.DeleteScreener(context.TODO(), "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestActiveV1SavedScreenerGetScreenerByID(t *testing.T) {
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
	_, err := client.Active.V1.SavedScreeners.GetScreenerByID(context.TODO(), "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestActiveV1SavedScreenerGetScreeners(t *testing.T) {
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
	_, err := client.Active.V1.SavedScreeners.GetScreeners(context.TODO())
	if err != nil {
		var apierr *clearstreet.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestActiveV1SavedScreenerReplaceScreenerWithOptionalParams(t *testing.T) {
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
	_, err := client.Active.V1.SavedScreeners.ReplaceScreener(
		context.TODO(),
		"550e8400-e29b-41d4-a716-446655440000",
		clearstreet.ActiveV1SavedScreenerReplaceScreenerParams{
			FieldFilter: []clearstreet.FieldRefParam{{
				Name:      "market_cap",
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
			Name: clearstreet.String("name"),
			SortBy: clearstreet.FieldRefParam{
				Name:      "market_cap",
				Lookback:  clearstreet.FieldLookbackOneWeek,
				Period:    clearstreet.FieldPeriodQuarter,
				ValueType: clearstreet.FieldTypeDecimal,
			},
			SortDirection: clearstreet.ActiveV1SavedScreenerReplaceScreenerParamsSortDirectionAsc,
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
