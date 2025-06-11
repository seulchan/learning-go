package termle

// corpusError defines a sentinel error.
// A sentinel error is a specific, predefined error value that can be checked for identity.
// By creating a custom type for our error, we can distinguish it from other generic errors.
type corpusError string

// Error is the implementation of the error interface by corpusError.
// The `error` interface in Go requires a single method: `Error() string`.
// This method returns the error message as a string.
func (e corpusError) Error() string {
	return string(e)
}
