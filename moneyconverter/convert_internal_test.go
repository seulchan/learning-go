// Package money_test contains internal tests for the money package.
package money

import (
	"reflect"
	"testing"
)

// TestApplyExchangeRate tests the logic of applying an exchange rate to an amount.
func TestApplyExchangeRate(t *testing.T) {
	tt := map[string]struct {
		in             Amount
		rate           ExchangeRate
		targetCurrency Currency
		expected       Amount
	}{
		// Test cases demonstrate various scenarios of precision handling.
		"Amount(1.52) * rate(1)": {
			in: Amount{
				quantity: Decimal{
					subunits:  152,
					precision: 2,
				},
				currency: Currency{code: "TST", precision: 2},
			}, // 1.52 TST
			rate:           ExchangeRate{subunits: 1, precision: 0}, // Rate of 1
			targetCurrency: Currency{code: "TRG", precision: 4},
			expected: Amount{
				quantity: Decimal{
					subunits:  15200,
					precision: 4,
				},
				currency: Currency{code: "TRG", precision: 4},
			}, // Expected: 1.5200 TRG
		},
		"Amount(2.50) * rate(4)": {
			in: Amount{
				quantity: Decimal{
					subunits:  250,
					precision: 2,
				}}, // 2.50
			rate:           ExchangeRate{subunits: 4, precision: 0},
			targetCurrency: Currency{code: "TRG", precision: 2},
			expected: Amount{
				quantity: Decimal{
					subunits:  1000,
					precision: 2,
				},
				currency: Currency{code: "TRG", precision: 2},
			}, // Expected: 10.00 TRG
		},
		"Amount(4) * rate(2.5)": {
			in: Amount{
				quantity: Decimal{
					subunits:  4,
					precision: 0,
				},
			}, // 4
			rate:           ExchangeRate{subunits: 25, precision: 1}, // Rate of 2.5
			targetCurrency: Currency{code: "TRG", precision: 0},
			expected: Amount{
				quantity: Decimal{
					subunits:  10,
					precision: 0,
				},
				currency: Currency{code: "TRG", precision: 0},
			}, // Expected: 10 TRG
		},
		"Amount(3.14) * rate(2.52678)": {
			in: Amount{
				quantity: Decimal{
					subunits:  314,
					precision: 2,
				}}, // 3.14
			rate:           ExchangeRate{subunits: 252678, precision: 5}, // Rate of 2.52678
			targetCurrency: Currency{code: "TRG", precision: 2},
			expected: Amount{
				quantity: Decimal{
					subunits:  793,
					precision: 2,
				},
				currency: Currency{code: "TRG", precision: 2},
			}, // 3.14 * 2.52678 = 7.9340892. Floored to 2 decimal places: 7.93
			// Calculation: (314 * 252678) = 79340892. Precision 2+5=7. (7.9340892)
			// Target precision 2. 79340892 / 10^(7-2) = 79340892 / 10^5 = 793. Subunits 793, precision 2.
		},
		"Amount(1.1) * rate(10)": {
			in: Amount{
				quantity: Decimal{
					subunits:  11,
					precision: 1,
				}}, // 1.1
			rate:           ExchangeRate{subunits: 10, precision: 0}, // Rate of 10
			targetCurrency: Currency{code: "TRG", precision: 1},
			expected: Amount{
				quantity: Decimal{
					subunits:  110,
					precision: 1,
				},
				currency: Currency{code: "TRG", precision: 1},
			}, // Expected: 11.0 TRG
		},
		"Amount(1_000_000_000.01) * rate(2)": {
			in: Amount{
				quantity: Decimal{
					subunits:  1_000_000_001,
					precision: 2,
				}}, // 1,000,000,000.01
			rate:           ExchangeRate{subunits: 2, precision: 0}, // Rate of 2
			targetCurrency: Currency{code: "TRG", precision: 2},
			expected: Amount{
				quantity: Decimal{
					subunits:  2_000_000_002,
					precision: 2,
				},
				currency: Currency{code: "TRG", precision: 2},
			}, // Expected: 2,000,000,000.02 TRG
		},
		"Amount(265_413.87) * rate(5.05935e-5)": {
			in: Amount{
				quantity: Decimal{
					subunits:  265_413_87,
					precision: 2,
				}}, // 265,413.87
			rate:           ExchangeRate{subunits: 505935, precision: 10}, // Rate of 0.0000505935
			targetCurrency: Currency{code: "TRG", precision: 2},
			expected: Amount{
				quantity: Decimal{
					subunits:  13_42,
					precision: 2,
				},
				currency: Currency{code: "TRG", precision: 2},
			}, // 265413.87 * 0.0000505935 = 13.4269...
			// Calculation: (26541387 * 505935) = 1342690303395. Precision 2+10=12. (13.42690303395)
			// Target precision 2. 1342690303395 / 10^(12-2) = 1342690303395 / 10^10 = 1342. Subunits 1342, precision 2. (13.42)

		},
		"Amount(265_413) * rate(1)": {
			in: Amount{
				quantity: Decimal{
					subunits:  265_413,
					precision: 0,
				}}, // 265,413
			rate:           ExchangeRate{subunits: 1, precision: 0}, // Rate of 1
			targetCurrency: Currency{code: "TRG", precision: 3},
			expected: Amount{
				quantity: Decimal{
					subunits:  265_413_000,
					precision: 3,
				},
				currency: Currency{code: "TRG", precision: 3},
			}, // Expected: 265,413.000 TRG
		},
		"Amount(2) * rate(1.337)": {
			in: Amount{
				quantity: Decimal{
					subunits:  2,
					precision: 0,
				}}, // 2
			rate:           ExchangeRate{subunits: 1337, precision: 3}, // Rate of 1.337
			targetCurrency: Currency{code: "TRG", precision: 5},
			expected: Amount{
				quantity: Decimal{
					subunits:  267400,
					precision: 5,
				},
				currency: Currency{code: "TRG", precision: 5},
			}, // Expected: 2.67400 TRG
			// Calculation: (2 * 1337) = 2674. Precision 0+3=3. (2.674)
			// Target precision 5. Subunits 2674 * 10^(5-3) = 2674 * 100 = 267400. Precision 5.
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			// applyExchangeRate now returns only Amount, as its internal logic doesn't produce errors.
			got := applyExchangeRate(tc.in, tc.targetCurrency, tc.rate)
			// reflect.DeepEqual is used for comparing structs.
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		name     string
		d1       Decimal
		r1       ExchangeRate // ExchangeRate is an alias for Decimal
		expected Decimal
	}{
		{
			name:     "1.50 * 2.0",
			d1:       Decimal{subunits: 150, precision: 2},     // 1.50
			r1:       ExchangeRate{subunits: 20, precision: 1}, // 2.0
			expected: Decimal{subunits: 3, precision: 0},       // 1.50 * 2.0 = 3.000 -> simplify to 3
			// (150 * 20) = 3000. precision 2+1=3. {3000, 3} -> {3,0}
		},
		{
			name:     "0.1 * 0.1",
			d1:       Decimal{subunits: 1, precision: 1},      // 0.1
			r1:       ExchangeRate{subunits: 1, precision: 1}, // 0.1
			expected: Decimal{subunits: 1, precision: 2},      // 0.1 * 0.1 = 0.01
			// (1*1)=1. precision 1+1=2. {1,2}
		},
		{
			name:     "123 * 1",
			d1:       Decimal{subunits: 123, precision: 0},    // 123
			r1:       ExchangeRate{subunits: 1, precision: 0}, // 1
			expected: Decimal{subunits: 123, precision: 0},    // 123
		},
		{
			name:     "0.5 * 0.5",
			d1:       Decimal{subunits: 5, precision: 1},      // 0.5
			r1:       ExchangeRate{subunits: 5, precision: 1}, // 0.5
			expected: Decimal{subunits: 25, precision: 2},     // 0.25
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := multiply(tc.d1, tc.r1)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("multiply(%v, %v) = %v, want %v", tc.d1, tc.r1, got, tc.expected)
			}
		})
	}
}
