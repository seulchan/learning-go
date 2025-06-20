// Package money (continued) - this file defines the Currency type and related logic.
package money

// Currency defines the code of a currency and its decimal precision.
// Precision indicates how many decimal places are typically used for this currency.
// For example, USD has precision 2 (e.g., $1.23), JPY has precision 0 (e.g., ¥123).
type Currency struct {
	// code is the 3-letter ISO 4217 currency code (e.g., "USD", "EUR", "JPY").
	code string
	// precision is the number of decimal places this currency uses.
	precision byte
}

// ErrInvalidCurrencyCode is returned when a currency code is not a valid 3-letter string.
const ErrInvalidCurrencyCode = MoneyError("invalid currency code: must be 3 letters")

// ParseCurrency attempts to parse a 3-letter currency code string and returns a Currency struct.
// It determines the currency's standard precision based on common conventions.
// If the code is not 3 letters long, it returns ErrInvalidCurrencyCode.
// For unrecognized but validly formatted 3-letter codes, it defaults to a precision of 2.
func ParseCurrency(code string) (Currency, error) {
	// ISO 4217 currency codes are always 3 letters.
	if len(code) != 3 {
		return Currency{}, ErrInvalidCurrencyCode
	}

	// Determine precision based on known currency conventions.
	// This list is not exhaustive but covers common cases.
	// Many currencies use 2 decimal places (e.g., hundredths like cents or pence).
	switch code {
	case "IRR": // Iranian Rial often has 0 decimal places in practice, though ISO might differ.
		return Currency{code: code, precision: 0}, nil
	case "MGA", "MRU": // Malagasy Ariary, Mauritanian Ouguiya have non-decimal subdivisions (1/5th). Precision 1 is a common simplification.
		return Currency{code: code, precision: 1}, nil
	case "CNY", "VND": // Chinese Yuan, Vietnamese Dong often use 1 decimal place in some contexts.
		return Currency{code: code, precision: 1}, nil
	case "BHD", "IQD", "KWD", "LYD", "OMR", "TND": // Bahraini Dinar, etc., use 3 decimal places.
		return Currency{code: code, precision: 3}, nil
	default:
		// For most other currencies, a precision of 2 (e.g., cents) is standard.
		// This includes EUR, USD, GBP, CAD, AUD, etc.
		return Currency{code: code, precision: 2}, nil
	}
}

// String implements the fmt.Stringer interface for the Currency type.
// It returns the 3-letter currency code.
func (c Currency) String() string {
	return c.code
}

// Code returns the ISO 4217 code for the currency (e.g., "USD").
func (c Currency) Code() string {
	return c.code
}
