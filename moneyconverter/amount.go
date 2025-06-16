// Package money provides types and functions for handling monetary values,
// including currencies, decimal amounts, and currency conversion.
package money

// Amount defines a decimal of money in a given currency.
// It combines a Decimal value with a Currency type.
type Amount struct {
	// quantity stores the monetary value as a Decimal.
	quantity Decimal
	// currency stores the currency information (code and precision).
	currency Currency
}

// Predefined error for amounts that are too precise for their currency.
const (
	// ErrTooPrecise is returned if the number is too precise for the currency.
	// For example, trying to represent 1.234 EUR when EUR only supports 2 decimal places.
	ErrTooPrecise = MoneyError("amount quantity is too precise for its currency")
)

// NewAmount returns an Amount of money.
// It takes a Decimal quantity and a Currency.
// It ensures that the quantity's precision matches the currency's precision.
// If the quantity is more precise than the currency allows, it returns ErrTooPrecise.
// If the quantity is less precise, it's adjusted to match the currency's precision
// (e.g., 1.5 EUR becomes 1.50 EUR if EUR precision is 2).
func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	switch {
	case quantity.precision > currency.precision:
		// The provided quantity has more decimal places than the currency supports.
		// For example, quantity is 1.234 (precision 3) but currency is EUR (precision 2).
		return Amount{}, ErrTooPrecise
	case quantity.precision < currency.precision:
		// The provided quantity has fewer decimal places than the currency requires.
		// We need to scale it up by adding trailing zeros.
		// Example: quantity is 1.5 (precision 1), currency is EUR (precision 2).
		// We change quantity to 1.50 (subunits 150, precision 2).
		quantity.subunits *= pow10(currency.precision - quantity.precision)
		quantity.precision = currency.precision
	}
	// If quantity.precision == currency.precision, no adjustment is needed.
	return Amount{quantity: quantity, currency: currency}, nil
}

// validate checks if an Amount is internally consistent and within supported limits.
// It's typically used after calculations to ensure the result is valid.
func (a Amount) validate() error {
	switch {
	case a.quantity.subunits > maxDecimal:
		// The underlying value (subunits) exceeds the maximum supported decimal value.
		return ErrTooLarge
	case a.quantity.precision > a.currency.precision:
		// This case should ideally not be reached if NewAmount is used correctly,
		// but it's a safeguard. It means the amount's precision somehow became
		// greater than what its currency allows.
		return ErrTooPrecise
	}
	return nil
}

// String implements the fmt.Stringer interface for the Amount type.
// It returns a string representation like "123.45 EUR".
func (a Amount) String() string {
	return a.quantity.String() + " " + a.currency.Code()
}
