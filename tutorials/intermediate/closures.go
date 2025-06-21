package main

import "fmt"

// --- Go Closures Tutorial ---
//
// A closure is a function value that "remembers" the environment in which it was created.
// This means it has access to variables from its surrounding scope (the "enclosing" function),
// *even after* the surrounding function has finished executing.
//
// Closures are a powerful concept used for:
// - Creating "factory" functions that generate other functions with specific configurations.
// - Managing state without using global variables.
// - Implementing concepts like iterators, generators, and private variables.

func main() {
	fmt.Println("--- Go Closures Demonstration ---")

	// --- Example 1: Simple Counter Closure ---
	fmt.Println("\n--- Creating the first counter (counterA) ---")
	// When we call `createCounter()`, it creates a new environment.
	// `counterA` is now a closure with its own `sum` variable, initialized to 0.
	counterA := createCounter()

	// Calling `counterA` multiple times demonstrates that it maintains its own state.
	fmt.Println("counterA call 1:", counterA()) // Output: 1
	fmt.Println("counterA call 2:", counterA()) // Output: 2
	fmt.Println("counterA call 3:", counterA()) // Output: 3

	fmt.Println("\n--- Creating a second, independent counter (counterB) ---")
	// Calling `createCounter()` again creates a *completely separate* environment.
	// `counterB` is a new closure with its own, independent `sum` variable, also starting at 0.
	counterB := createCounter()

	// The state of `counterB` is not affected by `counterA`.
	fmt.Println("counterB call 1:", counterB()) // Output: 1
	fmt.Println("counterB call 2:", counterB()) // Output: 2
	fmt.Println("counterA call 4:", counterA()) // Output: 4 (Shows counterA's state was preserved)

	// --- Example 2: Closure with Parameters ---
	fmt.Println("\n--- Creating a decrementer starting at 100 ---")
	// We create a new decrementer function that starts counting down from 100.
	// `countdown` is a closure that remembers its `currentValue` is 100.
	countdown := createCustomDecrementer(100)

	// Now we call the `countdown` closure, passing the amount to subtract.
	fmt.Println("Subtracting 5:", countdown(5))   // Output: 95
	fmt.Println("Subtracting 10:", countdown(10)) // Output: 85
	fmt.Println("Subtracting 20:", countdown(20)) // Output: 65

	fmt.Println("\nEnd of closures demonstration.")
}

// createCounter is a "factory" function. It doesn't return a value directly,
// but instead returns another function. The returned function is a closure.
func createCounter() func() int {
	// `sum` is a local variable within the `createCounter` function's scope.
	// When we return the inner anonymous function, this function "closes over" `sum`.
	// This means the inner function will have access to its own unique `sum` variable.
	sum := 0

	// We return an anonymous function (a function without a name).
	// This is the closure.
	return func() int {
		// Each time this closure is called, it increments its own `sum` and returns it.
		// It "remembers" the value of `sum` from the previous call.
		sum++
		return sum
	}
}

// createCustomDecrementer is another factory function that demonstrates a closure
// with a parameter. It creates a function that subtracts a given amount from a starting value.
func createCustomDecrementer(startValue int) func(int) int {
	// `currentValue` is the variable that the closure will "remember".
	// It's initialized with the `startValue` passed to the factory.
	currentValue := startValue

	// Return the closure.
	return func(decrementBy int) int {
		// This function subtracts `decrementBy` from its own `currentValue`.
		currentValue -= decrementBy
		return currentValue
	}
}
