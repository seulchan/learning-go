package main

import (
	"fmt"
	"math" // Imported for math constants like MaxInt64, MaxFloat64, etc.
)

func main() {
	fmt.Println("--- Go Arithmetic Operators Tutorial ---")

	// --- Basic Arithmetic Operators ---
	// Let's define two integer variables for our examples.
	a := 10
	b := 3

	// Addition (+)
	// Adds two operands.
	// Example: 10 + 3 = 13
	fmt.Println("Addition (a + b):", a+b)

	// Subtraction (-)
	// Subtracts the second operand from the first.
	// Example: 10 - 3 = 7
	fmt.Println("Subtraction (a - b):", a-b)

	// Multiplication (*)
	// Multiplies two operands.
	// Example: 10 * 3 = 30
	fmt.Println("Multiplication (a * b):", a*b)

	// Division (/)
	// Divides the first operand by the second.
	// If both operands are integers, Go performs integer division,
	// meaning the result is also an integer, and any fractional part is truncated (not rounded).
	// Example: 10 / 3 = 3 (the .333... part is discarded)
	fmt.Println("Division (a / b - integer):", a/b)

	// Modulus (%)
	// Returns the remainder of an integer division.
	// Example: 10 % 3 = 1 (because 10 divided by 3 is 3 with a remainder of 1)
	fmt.Println("Modulus (a % b):", a%b)

	fmt.Println("\n--- Division with Floating-Point Numbers ---")
	// To get a floating-point result from a division, at least one of the operands must be a float.

	// Example 1: Integer division assigned to a float variable
	// Here, 22 / 7 is performed first as integer division (result is 3).
	// Then, the integer 3 is converted to a float64 (3.0) and assigned to p_int_div.
	var p_int_div float64 = 22 / 7 // 22 / 7 results in 3 (integer division)
	fmt.Printf("22 / 7 (integer division assigned to float): %.10f\n", p_int_div)

	// Example 2: Floating-point division
	// By making at least one number a float (e.g., 22.0), the division becomes a floating-point operation.
	var p_float_div float64 = 22.0 / 7
	fmt.Printf("22.0 / 7 (floating-point division): %.10f\n", p_float_div)

	// You can also convert integers to floats before division:
	x_val := 22
	y_val := 7
	p_float_cast := float64(x_val) / float64(y_val)
	fmt.Printf("float64(%d) / float64(%d): %.10f\n", x_val, y_val, p_float_cast)

	// --- Unary Operators ---
	// Unary operators work on a single operand.
	num := 5
	negativeNum := -num // Unary minus: negates the value of num
	positiveNum := +num // Unary plus: indicates a positive value (often implicit, +5 is same as 5)
	fmt.Printf("\nUnary Minus (-%d): %d\n", num, negativeNum)
	fmt.Printf("Unary Plus (+%d): %d\n", num, positiveNum)

	// --- Increment and Decrement Statements ---
	// Go has `++` (increment) and `--` (decrement) statements.
	// They modify the variable in place and do not return a value (i.e., they are not expressions).
	counter := 0
	fmt.Println("\nInitial counter:", counter)
	counter++ // Increment counter by 1 (counter is now 1)
	fmt.Println("After counter++:", counter)
	counter-- // Decrement counter by 1 (counter is now 0)
	fmt.Println("After counter--:", counter)
	// Note: You cannot do `newValue := counter++` because `counter++` is a statement, not an expression.

	fmt.Println("\n--- Integer Overflow and Underflow ---")
	// Overflow occurs when an arithmetic operation results in a value that is too large
	// to be represented by the variable's data type.
	// Underflow occurs when the result is too small (e.g., for signed integers, more negative than the minimum).
	// For standard integers in Go, overflow/underflow typically "wraps around".

	// Signed Integer (int64) Overflow
	var maxInt64Val int64 = math.MaxInt64 // Maximum value for int64
	fmt.Println("Max Int64:", maxInt64Val)
	maxInt64Val = maxInt64Val + 1 // Adding 1 causes it to overflow
	// It wraps around to the smallest possible int64 value.
	fmt.Println("After Overflow (maxInt64Val + 1):", maxInt64Val, "(Expected:", math.MinInt64, ")")

	// Unsigned Integer (uint64) Overflow
	var uMaxIntVal uint64 = math.MaxUint64 // Maximum value for uint64
	fmt.Println("Max Uint64:", uMaxIntVal)
	uMaxIntVal = uMaxIntVal + 1 // Adding 1 causes it to overflow
	// It wraps around to 0 for unsigned integers.
	fmt.Println("After Overflow (uMaxIntVal + 1):", uMaxIntVal, "(Expected: 0)")

	// Signed Integer (int64) Underflow
	var minInt64Val int64 = math.MinInt64 // Minimum value for int64
	fmt.Println("Min Int64:", minInt64Val)
	minInt64Val = minInt64Val - 1 // Subtracting 1 causes it to underflow
	// It wraps around to the largest possible int64 value.
	fmt.Println("After Underflow (minInt64Val - 1):", minInt64Val, "(Expected:", math.MaxInt64, ")")

	// Unsigned Integer (uint64) Underflow
	var uMinIntVal uint64 = 0 // Minimum value for uint64 (which is 0)
	fmt.Println("Min Uint64:", uMinIntVal)
	uMinIntVal = uMinIntVal - 1 // Subtracting 1 causes it to underflow
	// It wraps around to the largest possible uint64 value.
	fmt.Println("After Underflow (uMinIntVal - 1):", uMinIntVal, "(Expected:", uint64(math.MaxUint64), ")")

	fmt.Println("\n--- Floating-Point Underflow ---")
	// For floating-point numbers, underflow can occur when a calculation results
	// in a number smaller in magnitude than what can be represented, often becoming zero.
	var smallFloat float64 = 1.0e-323
	fmt.Println("Small Float:", smallFloat) // Small float value
	// Dividing by a very large number can cause underflow.
	smallFloat = smallFloat / math.MaxFloat64
	// The result becomes 0.0 because it's too small to be represented.
	fmt.Println("After Division by Max Float64 (underflow to 0):", smallFloat)

	// Note on Floating-Point Precision:
	// Floating-point arithmetic can sometimes lead to precision issues due to how
	// decimal numbers are represented in binary. This is a general concept in computing.
	f1 := 0.1
	f2 := 0.2
	f3 := f1 + f2
	fmt.Printf("\nFloating-point precision (0.1 + 0.2): %.17f\n", f3) // May not be exactly 0.3
	if f3 != 0.3 {
		fmt.Println("  0.1 + 0.2 is NOT exactly 0.3 due to floating point representation.")
	}

	fmt.Println("\nEnd of arithmetic operators demonstration.")
}
