// Package money_test contains external tests for the money package.
// These tests import the "money" package just like an external user would.
package money_test

import (
	"errors"
	"fmt"
	money "learning-go/moneyconverter"
	"reflect"
	"strings"
	"testing"
)

// TestConvert demonstrates testing the main Convert function using a stub for the ratesFetcher interface.
func TestConvert(t *testing.T) {
	tt := map[string]struct {
		amount      money.Amount
		to          money.Currency
		stub        stubRateFetcher // Using our stub implementation of ratesFetcher
		expected    money.Amount
		expectedErr error // Use a specific error type if possible, or check with errors.Is
	}{
		"34.98 USD to EUR": {
			amount:      mustNewAmount(t, "34.98", "USD"), // USD has precision 2
			to:          mustParseCurrency(t, "EUR"),      // EUR has precision 2
			stub:        stubRateFetcher{rateStr: "2"},    // Rate of 2 (e.g., 1 USD = 2 EUR for this test)
			expected:    mustNewAmount(t, "69.96", "EUR"), // 34.98 * 2 = 69.96
			expectedErr: nil,
		},
		"100 JPY to USD with rate 0.0075": {
			amount:      mustNewAmount(t, "100", "JPY"),     // JPY has precision 0
			to:          mustParseCurrency(t, "USD"),        // USD has precision 2
			stub:        stubRateFetcher{rateStr: "0.0075"}, // 1 JPY = 0.0075 USD
			expected:    mustNewAmount(t, "0.75", "USD"),    // 100 * 0.0075 = 0.75
			expectedErr: nil,
		},
		"Error fetching rate": {
			amount:      mustNewAmount(t, "10.00", "CAD"),
			to:          mustParseCurrency(t, "GBP"),
			stub:        stubRateFetcher{err: fmt.Errorf("network unavailable")}, // Simulate an error from the fetcher
			expected:    money.Amount{},                                          // Expect zero Amount on error
			expectedErr: fmt.Errorf("failed to fetch exchange rate"),             // Check if the error is wrapped or of a specific type
		},
		"Conversion results in value too large": {
			amount:      mustNewAmount(t, "1000000000", "USD"), // 1 Billion USD
			to:          mustParseCurrency(t, "EUR"),
			stub:        stubRateFetcher{rateStr: "2000"}, // Rate that makes result > maxDecimal (10^12)
			expected:    money.Amount{},
			expectedErr: money.ErrTooLarge, // Expecting the validation error from Amount
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := money.Convert(tc.amount, tc.to, tc.stub)

			if tc.expectedErr != nil {
				if err == nil {
					t.Errorf("expected error satisfying %v, but got nil", tc.expectedErr)
				} else if !errors.Is(err, tc.expectedErr) && !strings.Contains(err.Error(), tc.expectedErr.Error()) {
					// Check if 'err' wraps 'tc.expectedErr' OR if the error message contains the expected message.
					// This provides flexibility if the exact error type/instance isn't available for errors.Is.
					t.Errorf("expected error satisfying %v, got %v", tc.expectedErr, err)
				}
			} else if err != nil {
				t.Errorf("expected no error, but got %v", err)
			}

			// Only compare amounts if no error was expected or if the error matched.
			// If an error is expected, 'got' might be a zero value and not meaningful to compare.
			if tc.expectedErr == nil && !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("expected amount %v, got %v", tc.expected, got)
			}
		})
	}
}

// stubRateFetcher is a simple stub implementation of the ratesFetcher interface,
// used for testing the Convert function without making real network calls.
type stubRateFetcher struct {
	rateStr string // The exchange rate to return, as a string (to be parsed into Decimal).
	err     error  // An error to return, if simulating a failure.
}

// FetchExchangeRate implements the ratesFetcher interface for stubRateFetcher.
// It ignores the source and target currency arguments and returns the pre-configured rate or error.
func (s stubRateFetcher) FetchExchangeRate(_, _ money.Currency) (money.ExchangeRate, error) {
	if s.err != nil {
		return money.ExchangeRate{}, s.err
	}
	// Parse the rate string into a Decimal.
	// In a real test, you might want to handle this parsing error too,
	// but for a simple stub, we can assume rateStr is valid if err is nil.
	rateDecimal, parseErr := money.ParseDecimal(s.rateStr)
	if parseErr != nil {
		// This would be an error in setting up the stub itself.
		return money.ExchangeRate{}, fmt.Errorf("stubRateFetcher: error parsing rateStr %q: %w", s.rateStr, parseErr)
	}
	return money.ExchangeRate(rateDecimal), nil
}

func mustParseCurrency(t *testing.T, code string) money.Currency {
	t.Helper()

	currency, err := money.ParseCurrency(code)
	if err != nil {
		t.Fatalf("mustParseCurrency: cannot parse currency code %q: %v", code, err)
	}

	return currency
}

// mustNewAmount is a test helper to create a money.Amount, failing the test on any error.
// This simplifies test case setup.
func mustNewAmount(t *testing.T, valueStr string, currencyCode string) money.Amount {
	t.Helper()

	decimalValue, err := money.ParseDecimal(valueStr)
	if err != nil {
		t.Fatalf("mustNewAmount: invalid decimal number %q: %v", valueStr, err)
	}

	currency, err := money.ParseCurrency(currencyCode)
	if err != nil {
		t.Fatalf("mustNewAmount: invalid currency code %q: %v", currencyCode, err)
	}

	amount, err := money.NewAmount(decimalValue, currency)
	if err != nil {
		t.Fatalf("mustNewAmount: cannot create amount with value %s and currency %s: %v", valueStr, currencyCode, err)
	}
	return amount
}
