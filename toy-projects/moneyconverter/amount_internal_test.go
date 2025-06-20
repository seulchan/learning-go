// Package money_test contains internal tests for the money package.
package money

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestNewAmount(t *testing.T) {
	tt := map[string]struct {
		quantity Decimal
		currency Currency
		want     Amount
		err      error
	}{
		"1.50 €": {
			quantity: Decimal{subunits: 150, precision: 2}, // 1.50
			currency: Currency{code: "EUR", precision: 2},  // EUR expects 2 decimal places
			want: Amount{
				quantity: Decimal{subunits: 150, precision: 2}, // Stays 1.50
				currency: Currency{code: "EUR", precision: 2},
			},
		},
		"1.5 € to 2-precision currency": {
			quantity: Decimal{subunits: 15, precision: 1}, // 1.5
			currency: Currency{code: "USD", precision: 2}, // USD expects 2 decimal places
			want: Amount{
				quantity: Decimal{subunits: 150, precision: 2}, // Becomes 1.50
				currency: Currency{code: "USD", precision: 2},
			},
		},
		"1 € to 2-precision currency": {
			quantity: Decimal{subunits: 1, precision: 0},  // 1
			currency: Currency{code: "CAD", precision: 2}, // CAD expects 2 decimal places
			want: Amount{
				quantity: Decimal{subunits: 100, precision: 2}, // Becomes 1.00
				currency: Currency{code: "CAD", precision: 2},
			},
		},
		"1.500 € (too precise)": {
			quantity: Decimal{subunits: 1500, precision: 3}, // 1.500
			currency: Currency{code: "EUR", precision: 2},   // EUR expects 2 decimal places
			err:      ErrTooPrecise,
		},
		"0 precision currency, 0 precision amount": {
			quantity: Decimal{subunits: 100, precision: 0}, // 100
			currency: Currency{code: "JPY", precision: 0},  // JPY expects 0 decimal places
			want: Amount{
				quantity: Decimal{subunits: 100, precision: 0},
				currency: Currency{code: "JPY", precision: 0},
			},
		},
		"0 precision currency, 1 precision amount (too precise)": {
			quantity: Decimal{subunits: 105, precision: 1}, // 10.5
			currency: Currency{code: "JPY", precision: 0},  // JPY expects 0 decimal places
			err:      ErrTooPrecise,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := NewAmount(tc.quantity, tc.currency)
			// errors.Is checks if the returned error 'err' is, or wraps, tc.err.
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
			// reflect.DeepEqual is used for comparing complex types like structs.
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected %v, got %v", tc.want, got)
			}
		})
	}
}

func TestAmount_String(t *testing.T) {
	amount, _ := NewAmount(Decimal{subunits: 12345, precision: 2}, Currency{code: "EUR", precision: 2}) // 123.45 EUR
	expected := "123.45 EUR"
	if got := amount.String(); got != expected {
		t.Errorf("Amount.String() = %q, want %q", got, expected)
	}

	amountNoDecimal, _ := NewAmount(Decimal{subunits: 1500, precision: 0}, Currency{code: "JPY", precision: 0}) // 1500 JPY
	expectedNoDecimal := "1500 JPY"
	if got := amountNoDecimal.String(); got != expectedNoDecimal {
		t.Errorf("Amount.String() for no decimal currency = %q, want %q", got, expectedNoDecimal)
	}
}

func TestAmount_validate(t *testing.T) {
	eur := Currency{code: "EUR", precision: 2}

	t.Run("valid amount", func(t *testing.T) {
		amount, _ := NewAmount(Decimal{subunits: 100, precision: 2}, eur) // 1.00 EUR
		if err := amount.validate(); err != nil {
			t.Errorf("validate() returned unexpected error: %v", err)
		}
	})

	t.Run("too large", func(t *testing.T) {
		// Create an amount that exceeds maxDecimal
		amount := Amount{quantity: Decimal{subunits: maxDecimal + 1, precision: 2}, currency: eur}
		err := amount.validate()
		if !errors.Is(err, ErrTooLarge) {
			t.Errorf("validate() expected ErrTooLarge, got %v", err)
		}
	})

	t.Run("too precise (inconsistent state)", func(t *testing.T) {
		// Manually create an inconsistent amount where quantity.precision > currency.precision
		// This bypasses NewAmount's adjustment logic to test validate's own check.
		amount := Amount{quantity: Decimal{subunits: 1234, precision: 3}, currency: eur} // 1.234 EUR, but EUR precision is 2
		err := amount.validate()
		if !errors.Is(err, ErrTooPrecise) {
			t.Errorf("validate() expected ErrTooPrecise, got %v", err)
		}
	})

	// Example of how to use mustParseDecimal and mustParseCurrency if needed for setup
	t.Run("valid amount using helpers", func(t *testing.T) {
		qty := mustParseDecimal(t, "25.50")
		curr := mustParseCurrency(t, "USD")
		amount, err := NewAmount(qty, curr)
		if err != nil {
			t.Fatalf("NewAmount failed: %v", err)
		}
		if err := amount.validate(); err != nil {
			t.Errorf("validate() returned unexpected error: %v", err)
		}
	})
}

// Helper functions (can be defined in a _test.go file or a separate test utility file)

func mustParseCurrency(t *testing.T, code string) Currency {
	t.Helper()
	c, err := ParseCurrency(code)
	if err != nil {
		t.Fatalf("mustParseCurrency: failed to parse currency %q: %v", code, err)
	}
	return c
}

func mustParseDecimal(t *testing.T, val string) Decimal {
	t.Helper()
	d, err := ParseDecimal(val)
	if err != nil {
		t.Fatalf("mustParseDecimal: failed to parse decimal %q: %v", val, err)
	}
	return d
}

func mustNewAmount(t *testing.T, valStr string, currencyCode string) Amount {
	t.Helper()
	dec, err := ParseDecimal(valStr)
	if err != nil {
		t.Fatalf("mustNewAmount: failed to parse decimal %q: %v", valStr, err)
	}
	curr, err := ParseCurrency(currencyCode)
	if err != nil {
		t.Fatalf("mustNewAmount: failed to parse currency %q: %v", currencyCode, err)
	}
	amt, err := NewAmount(dec, curr)
	if err != nil {
		t.Fatalf("mustNewAmount: failed to create amount for %s %s: %v", valStr, currencyCode, err)
	}
	return amt
}

// Example usage of Stringer for Amount in a test (can be part of TestAmount_String)
func ExampleAmount_String() {
	// Assume ParseDecimal and ParseCurrency are available and work
	d, _ := ParseDecimal("19.99")
	c, _ := ParseCurrency("USD")
	amount, _ := NewAmount(d, c)
	fmt.Println(amount)
	// Output: 19.99 USD
}
