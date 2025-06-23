// --- Go Custom Errors Tutorial ---
//
// While `errors.New` and `fmt.Errorf` are great for simple errors, sometimes you
// need to convey more structured information when something goes wrong. For example,
// you might want to include an operation name or a specific error code.
//
// This is where custom error types come in. By creating your own struct that
// implements the built-in `error` interface, you can create rich, informative
// errors that your program can inspect and act upon.
//
// This tutorial demonstrates:
// 1. How to define a custom error struct.
// 2. How to make it a valid error by implementing the `Error()` method.
// 3. How to support modern error wrapping by implementing the `Unwrap()` method.
// 4. How to inspect an error chain for your custom type using `errors.As`.
package main

import (
	"errors"
	"fmt"
	"time"
)

// --- 1. Defining the Custom Error Type ---
// `DatabaseError` is our custom error struct. It holds more context than a simple string.
type DatabaseError struct {
	Operation string    // The specific database operation that failed (e.g., "connect", "query")
	Code      int       // An internal error code for programmatic checks
	Timestamp time.Time // The time the error occurred
	Err       error     // The underlying, wrapped error
}

// --- 2. Implementing the `error` Interface ---
// To be a valid error, our type must have an `Error() string` method.
// This method defines the string representation of our error, which is what gets
// printed by functions like `fmt.Println(err)`.
func (e *DatabaseError) Error() string {
	// We format a descriptive error message. If there's a wrapped error, we include it.
	if e.Err != nil {
		return fmt.Sprintf("database error during '%s' (code %d): %v", e.Operation, e.Code, e.Err)
	}
	return fmt.Sprintf("database error during '%s' (code %d)", e.Operation, e.Code)
}

// --- 3. Supporting Error Unwrapping ---
// By implementing an `Unwrap() error` method, our custom error can be part of an
// "error chain". This allows functions like `errors.Is`, `errors.As`, and
// `errors.Unwrap` to inspect the errors that we have wrapped.
func (e *DatabaseError) Unwrap() error {
	return e.Err
}

// --- 4. Example Usage ---

// `connectToDB` simulates a low-level function that can fail.
func connectToDB() error {
	// In a real application, this would attempt a network connection.
	// Here, we just return a standard error.
	return errors.New("network timeout: could not reach database server")
}

// `fetchUserByID` is a higher-level function that uses `connectToDB`.
// If `connectToDB` fails, `fetchUserByID` wraps the error in our custom `DatabaseError` type.
func fetchUserByID(userID int) error {
	err := connectToDB()
	if err != nil {
		// Here we create and return our custom error.
		// We wrap the original `err` to preserve the root cause.
		return &DatabaseError{
			Operation: "connect",
			Code:      5001, // An arbitrary internal code for "connection failure"
			Timestamp: time.Now(),
			Err:       err, // Wrapping the original error
		}
	}
	// If connection succeeded, we would proceed to query for the user.
	// For this example, we'll just return nil for the success case.
	return nil
}

func main() {
	fmt.Println("--- Go Custom Errors Tutorial ---")

	// Let's call our function that can fail.
	err := fetchUserByID(123)

	if err != nil {
		// The default `fmt.Println` will call our custom Error() method.
		fmt.Printf("An error occurred: %v\n\n", err)

		// Now, let's inspect the error. We can check if the error (or any error
		// it wraps) is of our custom type `*DatabaseError`.
		// `errors.As` is the modern, idiomatic way to do this.
		var dbErr *DatabaseError
		if errors.As(err, &dbErr) {
			// The assertion succeeded! We can now access the fields of our custom error.
			fmt.Println("Error was identified as a DatabaseError. Details:")
			fmt.Printf("  - Operation: %s\n", dbErr.Operation)
			fmt.Printf("  - Code: %d\n", dbErr.Code)
			fmt.Printf("  - Timestamp: %s\n", dbErr.Timestamp.Format(time.RFC1123))
			fmt.Printf("  - Underlying Cause: %v\n", dbErr.Unwrap()) // Accessing the wrapped error
		}
	}
}
