// Package money (continued) - this file defines custom error types for the package.
package money

// MoneyError is a custom error type for errors specific to the money package.
// This allows callers to use errors.Is or errors.As for more specific error handling.
type MoneyError string

// Error implements the error interface.
func (e MoneyError) Error() string {
	return string(e)
}
