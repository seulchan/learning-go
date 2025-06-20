// Package money (continued) - this file focuses on currency conversion logic.
package money

import (
	"fmt"
)

// Convert takes an Amount in a source currency, a target Currency, and a ratesFetcher
// to get the exchange rate. It then returns the converted Amount in the target currency.
func Convert(amount Amount, to Currency, rates ratesFetcher) (Amount, error) {
	// Step 1: Fetch the exchange rate for the given source and target currencies.
	// The ratesFetcher interface allows for different ways to get rates (e.g., from a live API, a database, or a mock for testing).
	r, err := rates.FetchExchangeRate(amount.currency, to)
	if err != nil {
		// If fetching the rate fails, wrap the error and return.
		// %w is used to wrap the original error, allowing callers to inspect it using errors.Is or errors.As.
		return Amount{}, fmt.Errorf("failed to fetch exchange rate for %s to %s: %w", amount.currency.Code(), to.Code(), err)
	}

	// Step 2: Apply the fetched exchange rate to the original amount's quantity.
	// This calculation results in a new Decimal value representing the amount in the target currency,
	// but potentially with a precision that doesn't match the target currency yet.
	convertedValue := applyExchangeRate(amount, to, r)

	// Step 3: Validate the converted amount.
	// This checks if the new amount is within supported limits (e.g., not too large)
	// and if its precision is valid for the target currency.
	// Note: applyExchangeRate already adjusts precision, so this primarily checks for size.
	if err = convertedValue.validate(); err != nil {
		return Amount{}, fmt.Errorf("converted amount %s is invalid: %w", convertedValue.String(), err)
	}

	// If all steps are successful, return the new, converted Amount.
	return convertedValue, nil
}

// ratesFetcher is an interface that defines a method for fetching exchange rates.
// This abstraction allows the Convert function to be independent of how rates are obtained.
// For example, one implementation might call a web service, while another might read from a local cache or a mock for tests.
type ratesFetcher interface {
	// FetchExchangeRate retrieves the exchange rate for converting from a source Currency to a target Currency.
	FetchExchangeRate(source, target Currency) (ExchangeRate, error)
}

// ExchangeRate represents a rate to convert from a currency to another.
// It's an alias for Decimal, meaning an ExchangeRate is structurally identical to a Decimal
// but provides semantic distinction (it represents a rate, not a monetary quantity).
type ExchangeRate Decimal

// applyExchangeRate returns a new Amount representing the input multiplied by the rate.
// The precision of the returned value is that of the target Currency.
// This function assumes the multiplication itself doesn't cause an overflow that
// `multiply` can't handle before simplification. The `validate` call in `Convert`
// checks the final amount.
func applyExchangeRate(originalAmount Amount, targetCurrency Currency, exchangeRate ExchangeRate) Amount {
	// Multiply the original amount's quantity (a Decimal) by the exchange rate (also a Decimal).
	// The `multiply` function handles the arithmetic of scaled integers and their precisions.
	product := multiply(originalAmount.quantity, exchangeRate)

	// After multiplication, the product's precision (product.precision) might not match
	// the target currency's required precision (targetCurrency.precision).
	// We need to adjust it.
	switch {
	case product.precision > targetCurrency.precision:
		// The product is too precise (e.g., 1.2345 but target needs 2 decimal places).
		// We truncate the extra digits by dividing the subunits. This effectively floors the number.
		// Example: 12345 (prec 4) to prec 2 -> 12345 / 10^(4-2) = 12345 / 100 = 123.
		product.subunits = product.subunits / pow10(product.precision-targetCurrency.precision)
	case product.precision < targetCurrency.precision:
		// The product is not precise enough (e.g., 1.2 but target needs 3 decimal places).
		// We scale up the subunits by multiplying, effectively adding trailing zeros.
		// Example: 12 (prec 1) to prec 3 -> 12 * 10^(3-1) = 12 * 100 = 1200.
		product.subunits = product.subunits * pow10(targetCurrency.precision-product.precision)
	}
	// Set the product's precision to match the target currency's precision.
	product.precision = targetCurrency.precision

	return Amount{
		currency: targetCurrency,
		quantity: product, // The adjusted Decimal value
	}
}

// multiply performs decimal multiplication: (d.subunits * 10^-d.precision) * (er.subunits * 10^-er.precision).
// The result is (d.subunits * er.subunits) * 10^-(d.precision + er.precision).
func multiply(d Decimal, er ExchangeRate) Decimal {
	// Create a new Decimal for the product.
	// The new subunits value is the product of the original subunits.
	// The new precision is the sum of the original precisions.
	product := Decimal{
		subunits:  d.subunits * er.subunits,   // e.g., (150 [for 1.50]) * (20 [for 2.0]) = 3000
		precision: d.precision + er.precision, // e.g., 2 + 1 = 3. So, 3000 * 10^-3 = 3.000
	}

	// Simplify the product to its canonical form (e.g., remove trailing zeros from the fractional part).
	// For example, if product is {3000, 3} (representing 3.000), simplify changes it to {3, 0} (representing 3).
	product.simplify()

	return product
}
