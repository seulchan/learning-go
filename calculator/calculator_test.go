// Package calculator_test contains unit tests for the calculator package.
// Test files in Go are typically in the same package as the code they test,
// or in a separate package with a "_test" suffix (e.g., calculator_test).
// Using a separate package (calculator_test) allows testing only the exported
// functions, similar to how a user of the package would interact with it.
package calculator_test

import (
	"calculator" // The package we are testing.
	"math"       // Used for math.Abs in closeEnough.
	"testing"    // Go's built-in testing package.
)

// TestAdd tests the Add function of the calculator.
func TestAdd(t *testing.T) {
	// t.Parallel() marks this test function to be run in parallel with other parallelizable tests.
	// This can speed up test execution, especially for I/O-bound tests or on multi-core processors.
	t.Parallel()

	// Table-driven tests are a common pattern in Go.
	// We define a struct for our test cases and then a slice of these structs.
	type testCase struct {
		name string  // Optional: A descriptive name for the test case.
		a, b float64 // Inputs to the Add function.
		want float64 // The expected result.
	}

	testCases := []testCase{
		{name: "positive numbers", a: 2, b: 2, want: 4},
		{name: "adding one", a: 1, b: 1, want: 2},
		{name: "adding zero", a: 5, b: 0, want: 5},
		{name: "adding negative number", a: 5, b: -2, want: 3},
		{name: "adding two negative numbers", a: -5, b: -2, want: -7},
	}

	// We iterate over each test case.
	for _, tc := range testCases {
		// It's good practice to run each table entry as a subtest using t.Run.
		// This gives clearer output if a specific case fails and allows for focused re-running of failed subtests.
		// The first argument to t.Run is the name of the subtest.
		t.Run(tc.name, func(t *testing.T) {
			// Call the function we are testing.
			got := calculator.Add(tc.a, tc.b)
			// Check if the result ('got') is what we expected ('want').
			if tc.want != got {
				// t.Errorf logs an error message and marks the test as failed, but continues execution.
				t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
			}
		})
	}
}

// TestSubtract tests the Subtract function.
func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a, b float64
		want float64
	}
	testCases := []testCase{
		{name: "subtract equal numbers", a: 2, b: 2, want: 0},
		{name: "subtract one from one", a: 1, b: 1, want: 0},
		{name: "subtract zero", a: 5, b: 0, want: 5},
		{name: "subtract positive from positive", a: 5, b: 2, want: 3},
		{name: "subtract negative from positive", a: 5, b: -2, want: 7},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := calculator.Subtract(tc.a, tc.b)
			if tc.want != got {
				t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
			}
		})
	}
}

// TestMultiply tests the Multiply function.
func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a, b float64
		want float64
	}
	testCases := []testCase{
		{name: "multiply positive numbers", a: 2, b: 2, want: 4},
		{name: "multiply by one", a: 1, b: 1, want: 1},
		{name: "multiply by zero", a: 5, b: 0, want: 0},
		{name: "multiply by negative number", a: 5, b: -2, want: -10},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := calculator.Multiply(tc.a, tc.b)
			if tc.want != got {
				t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
			}
		})
	}
}

// TestDivide tests the Divide function for valid inputs.
func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a, b float64
		want float64
	}
	testCases := []testCase{
		{name: "divide two positives", a: 2, b: 2, want: 1},
		{name: "divide two negatives", a: -1, b: -1, want: 1},
		{name: "divide larger by smaller", a: 10, b: 2, want: 5},
		{name: "divide to get fraction", a: 1, b: 3, want: 0.333333}, // Note: floating point precision
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.Divide(tc.a, tc.b)
			// For valid inputs, we expect no error.
			if err != nil {
				// t.Fatalf logs an error and stops execution for this test case immediately.
				// Useful when a failure means subsequent checks are pointless.
				t.Fatalf("Divide(%f, %f): unexpected error: %v", tc.a, tc.b, err)
			}
			// Comparing floating-point numbers for exact equality can be problematic due to precision issues.
			// It's better to check if they are "close enough".
			if !closeEnough(tc.want, got, 0.000001) { // Using a smaller tolerance for more precision
				t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
			}
		})
	}
}

// TestDivideInvalid tests the Divide function for invalid inputs (division by zero).
func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	// We expect an error when dividing by zero.
	_, err := calculator.Divide(1, 0)
	if err == nil {
		// If err is nil, it means no error was returned, which is not what we want.
		t.Error("Divide(1,0): want error for division by zero, got nil")
	}
}

// TestSqrt tests the Sqrt function for valid inputs.
func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a    float64
		want float64
	}
	testCases := []testCase{
		{name: "sqrt of 2", a: 2, want: 1.41421356}, // More precise value
		{name: "sqrt of 1", a: 1, want: 1},
		{name: "sqrt of 10", a: 10, want: 3.16227766}, // More precise value
		{name: "sqrt of 16", a: 16, want: 4},
		{name: "sqrt of 0", a: 0, want: 0},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.Sqrt(tc.a)
			if err != nil {
				t.Fatalf("Sqrt(%f): unexpected error: %v", tc.a, err)
			}
			if !closeEnough(tc.want, got, 0.000001) { // Using a smaller tolerance
				t.Errorf("Sqrt(%f): want %f, got %f", tc.a, tc.want, got)
			}
		})
	}
}

// TestSqrtInvalid tests the Sqrt function for invalid inputs (negative numbers).
func TestSqrtInvalid(t *testing.T) {
	t.Parallel()
	// We expect an error when taking the square root of a negative number.
	_, err := calculator.Sqrt(-1)
	if err == nil {
		t.Error("Sqrt(-1): want error for negative input, got nil")
	}
}

// closeEnough checks if two floating-point numbers are within a certain tolerance of each other.
// This is necessary because floating-point arithmetic isn't always exact.
// Parameters:
//
//	a, b: the two float64 numbers to compare.
//	tolerance: the maximum allowed difference between a and b.
//
// Returns:
//
//	true if the absolute difference between a and b is less than or equal to tolerance, false otherwise.
func closeEnough(a, b, tolerance float64) bool {
	// math.Abs returns the absolute value of (a - b).
	return math.Abs(a-b) <= tolerance
}
