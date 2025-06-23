// --- Go Error Handling Tutorial ---
//
// In Go, errors are not exceptions; they are values. This is a fundamental concept.
// Functions that can fail return an `error` value as their last return value.
// The calling code is then responsible for checking this error value to see if the
// operation was successful. The standard way to do this is with an `if err != nil` check.
//
// This tutorial covers the essential patterns for working with errors in Go:
// 1. The basic `if err != nil` pattern.
// 2. Creating simple errors with `errors.New`.
// 3. Creating formatted errors and wrapping them with `fmt.Errorf` and the `%w` verb.
// 4. Defining and using custom error types.
// 5. Inspecting error chains with `errors.As` and `errors.Unwrap`.
package main

import (
	"errors"
	"fmt"
	"math"
)

// --- 1. The Basic Error Handling Pattern ---

// calculateSquareRoot is a function that can fail.
// It returns a `float64` and an `error`. If the operation is successful,
// the error is `nil`. If it fails, it returns a non-nil error.
func calculateSquareRoot(x float64) (float64, error) {
	if x < 0 {
		// `errors.New` creates a simple error with a static message.
		// This is the most basic way to create an error.
		return 0, errors.New("math: cannot calculate square root of a negative number")
	}
	return math.Sqrt(x), nil // Success: return the result and a nil error.
}

// --- 2. Custom Error Types ---

// Sometimes, a simple string isn't enough. A custom error type lets you
// attach more structured data to your errors.
// OperationError is a custom error struct.
type OperationError struct {
	Op      string // The operation that failed
	Message string // A descriptive message
}

// To be a valid error, our type must implement the Error() method.
// This makes it satisfy the built-in `error` interface.
func (e *OperationError) Error() string {
	return fmt.Sprintf("operation '%s' failed: %s", e.Op, e.Message)
}

// processData is a function that can return our custom error.
func processData(data []byte) error {
	if len(data) == 0 {
		// Return an instance of our custom error type.
		return &OperationError{
			Op:      "processData",
			Message: "input data cannot be empty",
		}
	}
	// Simulate successful processing.
	return nil
}

// --- 3. Error Wrapping ---

// Error wrapping allows you to add context to an error without losing the
// original error information. This creates a "chain" of errors.

// loadConfig simulates a low-level operation that can fail.
func loadConfig() error {
	// For demonstration, we'll create a new error. In a real app, this might
	// come from a file I/O operation, like `os.ErrPermission`.
	return errors.New("permission denied while reading config file")
}

// startup calls loadConfig and wraps the error if it occurs.
func startup() error {
	err := loadConfig()
	if err != nil {
		// `fmt.Errorf` with the `%w` verb wraps the original error.
		// This adds context ("failed to start application") while preserving `err`.
		return fmt.Errorf("failed to start application: %w", err)
	}
	return nil
}

func main() {
	fmt.Println("--- Go Error Handling Tutorial ---")

	// --- Part 1: Basic Error Checking ---
	fmt.Println("\n--- 1. Basic Error Checking ---")
	result, err := calculateSquareRoot(16.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err) // Handle the error
	} else {
		fmt.Printf("Square root of 16 is %.1f\n", result)
	}

	result, err = calculateSquareRoot(-4.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Square root of -4 is %.1f\n", result)
	}

	// --- Part 2: Handling Custom Errors ---
	fmt.Println("\n--- 2. Handling Custom Errors ---")
	err = processData([]byte{}) // Pass empty data to trigger the error
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		// We can check if the error is our specific custom type.
		// `errors.As` checks the error chain for a `*OperationError` and,
		// if found, assigns it to `opErr`.
		var opErr *OperationError
		if errors.As(err, &opErr) {
			fmt.Println("This is an OperationError. We can access its fields:")
			fmt.Printf("  Failing Operation: %s\n", opErr.Op)
		}
	}

	// --- Part 3: Handling Wrapped Errors ---
	fmt.Println("\n--- 3. Handling Wrapped Errors ---")
	err = startup()
	if err != nil {
		// The printed error includes both our context and the original error message.
		fmt.Printf("Startup Error: %v\n", err)

		// We can "unwrap" the error to get the original cause.
		originalErr := errors.Unwrap(err)
		if originalErr != nil {
			fmt.Printf("Original (wrapped) error: %v\n", originalErr)
		}
	}
}
