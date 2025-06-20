// Package money_test contains internal tests for the money package.
package money

import (
	"errors"
	"testing"
)

// TestParseCurrency_Success tests successful parsing of various currency codes.
func TestParseCurrency_Success(t *testing.T) {
	tt := map[string]struct {
		in       string
		expected Currency
	}{
		"majority EUR":     {in: "EUR", expected: Currency{code: "EUR", precision: 2}},
		"thousandth BHD":   {in: "BHD", expected: Currency{code: "BHD", precision: 3}},
		"tenth CNY":        {in: "CNY", expected: Currency{code: "CNY", precision: 1}},
		"zero decimal IRR": {in: "IRR", expected: Currency{code: "IRR", precision: 0}},
		"default USD":      {in: "USD", expected: Currency{code: "USD", precision: 2}}, // Handled by default case
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseCurrency(tc.in)
			if err != nil {
				t.Fatalf("ParseCurrency(%q) returned an unexpected error: %v", tc.in, err)
			}

			if got != tc.expected {
				t.Errorf("ParseCurrency(%q) = %v, want %v", tc.in, got, tc.expected)
			}
		})
	}
}

// TestParseCurrency_InvalidCode tests parsing of invalid currency codes.
func TestParseCurrency_InvalidCode(t *testing.T) {
	testCases := []string{"INVALID", "US", "EURO", ""}

	for _, tc := range testCases {
		t.Run(tc, func(t *testing.T) {
			_, err := ParseCurrency(tc)
			// errors.Is checks if the returned error 'err' is, or wraps, ErrInvalidCurrencyCode.
			if !errors.Is(err, ErrInvalidCurrencyCode) {
				t.Errorf("ParseCurrency(%q) expected error %v, got %v", tc, ErrInvalidCurrencyCode, err)
			}
		})
	}
}

func TestCurrency_StringAndCode(t *testing.T) {
	c := Currency{code: "XYZ", precision: 2}
	if c.String() != "XYZ" {
		t.Errorf("Currency.String() = %q, want %q", c.String(), "XYZ")
	}
	if c.Code() != "XYZ" {
		t.Errorf("Currency.Code() = %q, want %q", c.Code(), "XYZ")
	}
}
