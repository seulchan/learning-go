package main

import (
	"errors"
	"fmt"
)

// main is the entry point of our Go program.
// We'll use this function to demonstrate how other functions work.
func main() {
	fmt.Println("--- Go Functions Tutorial ---")

	// --- 1. Basic Function Call ---
	// Let's call our `addNumbers` function and print its result.
	fmt.Println("\n--- 1. Basic Function Call ---")
	sumOfTwoNumbers := addNumbers(100, 1)
	fmt.Printf("The sum of 100 and 1 is: %d\n", sumOfTwoNumbers)

	sumOfMoreNumbers := addNumbers(10, 20)
	fmt.Printf("The sum of 10 and 20 is: %d\n", sumOfMoreNumbers)

	// --- 2. Anonymous Functions ---
	// Functions without a name are called anonymous functions.
	// They can be defined and called inline, or assigned to variables.
	fmt.Println("\n--- 2. Anonymous Functions ---")
	// Define an anonymous function and assign it to the variable `sayHello`.
	sayHello := func() {
		fmt.Println("Hello Anonymous Function!")
	}
	sayHello() // Call the anonymous function.

	// Anonymous function with parameters and return value.
	multiply := func(a, b int) int {
		return a * b
	}
	product := multiply(5, 3)
	fmt.Printf("Product from anonymous function (5 * 3): %d\n", product)

	// --- 3. Functions as First-Class Citizens ---
	// In Go, functions can be treated like any other value:
	// assigned to variables, passed as arguments, or returned from other functions.
	fmt.Println("\n--- 3. Functions as First-Class Citizens ---")
	var mathOperation func(int, int) int // Declare a variable that can hold a function
	mathOperation = addNumbers           // Assign the `addNumbers` function to `mathOperation`

	resultFromVar := mathOperation(3, 5)
	fmt.Printf("Result of `mathOperation(3, 5)` (which is `addNumbers`): %d\n", resultFromVar)

	// --- 4. Higher-Order Functions (Passing Functions as Arguments) ---
	// A function that takes another function as an argument or returns a function.
	fmt.Println("\n--- 4. Higher-Order Functions (Passing Functions as Arguments) ---")
	sumResult := calculate(100, 10, addNumbers) // Pass `addNumbers` function as an argument
	fmt.Printf("Result of calculate(100, 10, addNumbers): %d\n", sumResult)

	subtractResult := calculate(100, 10, func(a, b int) int { return a - b }) // Pass an anonymous function
	fmt.Printf("Result of calculate(100, 10, func_subtract): %d\n", subtractResult)

	// --- 5. Functions Returning Functions (Closures) ---
	// A function can return another function. The returned function "closes over"
	// variables from its lexical scope (the `createMultiplier` function's scope).
	fmt.Println("\n--- 5. Functions Returning Functions (Closures) ---")
	multiplyBy100 := createMultiplier(100)                           // `multiplyBy100` is now a function
	fmt.Printf("Result of multiplyBy100(5): %d\n", multiplyBy100(5)) // Calls the returned function

	multiplyBy5 := createMultiplier(5)
	fmt.Printf("Result of multiplyBy5(10): %d\n", multiplyBy5(10))

	// --- 6. Multiple Return Values and Error Handling ---
	// Functions in Go can return multiple values. This is often used to return
	// a result and an error value.
	fmt.Println("\n--- 6. Multiple Return Values and Error Handling ---")
	quotient1, remainder1, err1 := divideWithValidation(10, 3)
	if err1 != nil {
		fmt.Printf("Error dividing 10 by 3: %s\n", err1)
	} else {
		fmt.Printf("10 divided by 3: Quotient = %d, Remainder = %d\n", quotient1, remainder1)
	}

	_, _, err2 := divideWithValidation(10, 0) // Attempting to divide by zero
	if err2 != nil {
		fmt.Printf("Error dividing 10 by 0: %s\n", err2)
	} else {
		// This block won't be reached in this case
		fmt.Println("Division by zero somehow succeeded (this should not happen).")
	}

	// --- 7. Simple Function with Single String Return ---
	fmt.Println("\n--- 7. Simple Function with Single String Return ---")
	comparisonResult1 := compareNumbers(5, 2)
	fmt.Printf("Comparing 5 and 2: %s\n", comparisonResult1)

	comparisonResult2 := compareNumbers(2, 5)
	fmt.Printf("Comparing 2 and 5: %s\n", comparisonResult2)

	comparisonResult3 := compareNumbers(5, 5)
	fmt.Printf("Comparing 5 and 5: %s\n", comparisonResult3)

	// --- 8. Variadic Functions ---
	// Functions that can accept a variable number of arguments of the same type.
	fmt.Println("\n--- 8. Variadic Functions ---")
	totalSum1 := sumNumbersVariadic(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1,2,3,4,5: %d\n", totalSum1)

	totalSum2 := sumNumbersVariadic(10, 20)
	fmt.Printf("Sum of 10,20: %d\n", totalSum2)

	totalSum3 := sumNumbersVariadic() // Calling with no arguments
	fmt.Printf("Sum of no numbers: %d\n", totalSum3)

	// You can also pass a slice to a variadic function using the `...` syntax.
	numbersToSum := []int{100, 200, 300}
	totalSumFromSlice := sumNumbersVariadic(numbersToSum...)
	fmt.Printf("Sum of slice [100, 200, 300]: %d\n", totalSumFromSlice)

	// --- 9. Variadic Functions with Other Parameters ---
	fmt.Println("\n--- 9. Variadic Functions with Other Parameters ---")
	message, totalForMessage := sumWithLabel("The sum of these numbers is: ", 1, 2, 3)
	fmt.Printf("%s%d\n", message, totalForMessage)

	fmt.Println("\n--- End of Functions Tutorial ---")
}

// --- Function Definitions ---

// addNumbers demonstrates a basic function in Go.
//
// Syntax:
//
//	func <functionName>(<parameterName> <parameterType>, ...) <returnType> {
//	    // function body
//	    return <value>
//	}
//
// - `func` keyword: Used to declare a function.
// - `addNumbers`: The name of our function.
// - `(num1 int, num2 int)`: The parameter list.
//   - `num1` and `num2` are parameter names.
//   - `int` is their type. If multiple parameters have the same type,
//     you can write `(num1, num2 int)`.
//
// - `int`: The type of the value this function will return.
// - The function body contains the logic.
// - `return num1 + num2`: The `return` statement specifies the value to be returned.
func addNumbers(num1 int, num2 int) int {
	return num1 + num2
}

// calculate is a higher-order function.
// It takes two integers (`operand1`, `operand2`) and a function (`operation`) as arguments.
// The `operation` function itself must take two integers and return an integer.
// `calculate` then applies the `operation` to `operand1` and `operand2`.
func calculate(operand1 int, operand2 int, operation func(int, int) int) int {
	// Call the passed-in `operation` function with the provided operands.
	return operation(operand1, operand2)
}

// createMultiplier is a function that returns another function (a closure).
// The returned function "remembers" the `factor` from its parent's scope.
//   - `factor int`: The number by which the returned function will multiply its input.
//   - `func(int) int`: This is the signature of the function that `createMultiplier` returns.
//     It means the returned function takes an `int` and returns an `int`.
func createMultiplier(factor int) func(int) int {
	// This is the anonymous function that gets returned.
	// It takes one argument `numberToMultiply` and multiplies it by `factor`.
	// `factor` is captured from the `createMultiplier` function's scope (this is a closure).
	return func(numberToMultiply int) int {
		return numberToMultiply * factor
	}
}

// divideWithValidation demonstrates returning multiple values, including an error.
// It also shows how to use named return values.
//
// Syntax for multiple (and named) return values:
//
//		func <functionName>(...) (returnValueName1 <type1>, returnValueName2 <type2>, ...) {
//		    // function body
//		    // Assign values to returnValueName1, returnValueName2, ...
//		    return // A "naked" return can be used if all named return values are set.
//		           // Or, explicitly: return returnValueName1, returnValueName2, ...
//		}
//
//	  - `dividend int, divisor int`: Input parameters.
//	  - `(quotient int, remainder int, err error)`: Named return values.
//	    This declares three return values: `quotient` (int), `remainder` (int), and `err` (error).
func divideWithValidation(dividend int, divisor int) (quotient int, remainder int, err error) {
	// It's crucial to handle potential errors, like division by zero.
	if divisor == 0 {
		// If divisor is zero, we can't perform the division.
		// We set `err` to a new error created using `errors.New()`.
		// The quotient and remainder will be their zero values (0 for int).
		err = errors.New("cannot divide by zero")
		// We use an explicit return here to make it clear what's being returned in the error case.
		// A naked `return` would also work as `err` is set and `quotient`/`remainder` are 0.
		return 0, 0, err
	}

	// If no error, perform the division and modulo operations.
	quotient = dividend / divisor
	remainder = dividend % divisor
	// `err` is already `nil` by default (its zero value), so we don't need to set it explicitly.
	// A "naked" return is used here. It returns the current values of `quotient`, `remainder`, and `err`.
	return
}

// compareNumbers compares two integers and returns a string describing their relationship.
// This is a simple function demonstrating a single string return value based on conditions.
func compareNumbers(num1, num2 int) string {
	if num1 > num2 {
		return fmt.Sprintf("%d is greater than %d", num1, num2)
	} else if num1 < num2 {
		return fmt.Sprintf("%d is less than %d", num1, num2)
	}
	// If neither greater nor less, they must be equal.
	return fmt.Sprintf("%d is equal to %d", num1, num2)
}

// sumNumbersVariadic demonstrates a variadic function.
// The `...int` syntax means this function can accept zero or more `int` arguments.
// Inside the function, `numbers` is treated as a slice of integers (`[]int`).
//
// Syntax for variadic functions:
//
//	func <functionName>(param1 type1, ..., <variadicParamName> ...<variadicParamType>) <returnType> {
//	    // function body
//	}
//
// The variadic parameter must be the last parameter in the function signature.
func sumNumbersVariadic(numbers ...int) int {
	total := 0
	// We can iterate over the `numbers` slice using a for...range loop.
	// `_` (blank identifier) is used because we don't need the index in this case.
	for _, num := range numbers {
		total += num
	}
	return total
}

// sumWithLabel demonstrates a function with a regular parameter followed by a variadic parameter.
// It also returns multiple values.
// - `label string`: A regular string parameter.
// - `numbers ...int`: A variadic integer parameter.
// - `(string, int)`: Returns the label and the sum of the numbers.
func sumWithLabel(label string, numbers ...int) (string, int) {
	total := 0
	for _, v := range numbers {
		total += v
	}
	// Return the provided label and the calculated total.
	return label, total
}
