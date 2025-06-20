// Package money_test contains internal tests for the money package.
package money

import (
	"errors"
	"testing"
)

// TestParseDecimal tests the parsing of string values into Decimal structs.
func TestParseDecimal(t *testing.T) {
	tt := map[string]struct {
		decimal  string
		expected Decimal
		err      error
	}{
		"2 decimal digits": {
			decimal:  "1.52", // Standard case
			expected: Decimal{subunits: 152, precision: 2},
			err:      nil,
		},
		"no decimal digits": {
			decimal:  "1", // Integer
			expected: Decimal{1, 0},
			err:      nil,
		},
		"suffix 0 as decimal digits": {
			decimal:  "1.50", // Should simplify to 1.5
			expected: Decimal{15, 1},
			err:      nil,
		},
		"prefix 0 as decimal digits": {
			decimal:  "1.02", // Leading zero in fractional part
			expected: Decimal{102, 2},
			err:      nil,
		},
		"multiple of 10": {
			decimal:  "150", // Integer ending in zero
			expected: Decimal{150, 0},
			err:      nil,
		},
		"only fractional part": {
			decimal:  ".25", // e.g., 0.25
			expected: Decimal{25, 2},
			err:      nil,
		},
		"only fractional part with trailing zero": {
			decimal:  ".50", // e.g., 0.50, simplifies to 0.5
			expected: Decimal{5, 1},
			err:      nil,
		},
		"zero value": {
			decimal:  "0.00",
			expected: Decimal{0, 0}, // Simplifies to 0
			err:      nil,
		},
		"zero integer": {
			decimal:  "0",
			expected: Decimal{0, 0},
			err:      nil,
		},
		"invalid decimal part": {
			decimal: "65.pocket", // Non-numeric fractional part
			err:     ErrInvalidDecimal,
		},
		"invalid integer part": {
			decimal: "pocket.65", // Non-numeric integer part
			err:     ErrInvalidDecimal,
		},
		"multiple decimal points": {
			decimal: "1.2.3",
			err:     ErrInvalidDecimal,
		},
		"Not a number": {
			decimal: "NaN", // Special float string, not a valid decimal for this parser
			err:     ErrInvalidDecimal,
		},
		"empty string": {
			decimal: "", // Empty input
			err:     ErrInvalidDecimal,
		},
		"too large": {
			decimal: "1234567890123", // Exceeds maxDecimal (10^12)
			err:     ErrTooLarge,
		},
		"too large with decimals": {
			decimal: "1234567890123.45", // Integer part alone exceeds maxDecimal
			err:     ErrTooLarge,
		},
		"just at maxDecimal": {
			decimal:  "1000000000000", // 10^12
			expected: Decimal{subunits: 1000000000000, precision: 0},
			err:      nil,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseDecimal(tc.decimal)
			// errors.Is checks if the returned error 'err' is, or wraps, tc.err.
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
			// Only compare 'got' with 'tc.expected' if no error was expected.
			// If an error is expected, 'got' will be the zero value for Decimal.
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestDecimal_String(t *testing.T) {
	testCases := []struct {
		name     string
		decimal  Decimal
		expected string
	}{
		{"integer", Decimal{subunits: 123, precision: 0}, "123"},
		{"two decimal places", Decimal{subunits: 12345, precision: 2}, "123.45"},
		{"one decimal place", Decimal{subunits: 1235, precision: 1}, "123.5"},
		{"trailing zero in subunits, one decimal place", Decimal{subunits: 120, precision: 1}, "12.0"}, // e.g. from 12.00 simplified to {120,1}
		{"three decimal places", Decimal{subunits: 12305, precision: 3}, "12.305"},
		{"zero value", Decimal{subunits: 0, precision: 0}, "0"},
		{"zero with precision", Decimal{subunits: 0, precision: 2}, "0.00"}, // e.g. from 0.00
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.decimal.String(); got != tc.expected {
				t.Errorf("Decimal.String() for %v: got %q, want %q", tc.decimal, got, tc.expected)
			}
		})
	}
}

func TestDecimal_simplify(t *testing.T) {
	testCases := []struct {
		name     string
		input    Decimal
		expected Decimal
	}{
		{"1.50 -> 1.5", Decimal{150, 2}, Decimal{15, 1}},
		{"1.00 -> 1", Decimal{100, 2}, Decimal{1, 0}},
		{"1.23 -> 1.23 (no change)", Decimal{123, 2}, Decimal{123, 2}},
		{"1500 -> 1500 (no change, integer)", Decimal{1500, 0}, Decimal{1500, 0}},
		{"0.00 -> 0", Decimal{0, 2}, Decimal{0, 0}},
		{"0.50 -> 0.5", Decimal{50, 2}, Decimal{5, 1}},
		{"12000, prec 3 (12.000) -> 12, prec 0", Decimal{12000, 3}, Decimal{12, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := tc.input // Make a copy to modify
			d.simplify()
			if d != tc.expected {
				t.Errorf("simplify() for %v: got %v, want %v", tc.input, d, tc.expected)
			}
		})
	}
}

func TestPow10(t *testing.T) {
	testCases := []struct {
		power    byte
		expected int64
	}{
		{0, 1},
		{1, 10},
		{2, 100},
		{3, 1000},
		{4, 10000}, // Test default case in switch
		{5, 100000},
	}

	for _, tc := range testCases {
		t.Run(string(tc.power), func(t *testing.T) {
			if got := pow10(tc.power); got != tc.expected {
				t.Errorf("pow10(%d) = %d, want %d", tc.power, got, tc.expected)
			}
		})
	}
}
