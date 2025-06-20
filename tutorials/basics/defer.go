// Package main is the entry point for our program.
package main

// Import the "fmt" package, which provides functions for formatted input and output,
// such as printing to the console.
import "fmt"

// main is the function where program execution begins.
func main() {
	// We'll call our function that demonstrates defer with an initial value.
	fmt.Println("--- Starting Defer Demonstration ---")
	demonstrateDefer(10)
	fmt.Println("--- Defer Demonstration Finished ---")
}

// demonstrateDefer shows how the defer statement works in Go.
// It takes an integer `initialValue` as input.
func demonstrateDefer(initialValue int) {
	// The `defer` keyword schedules a function call to be executed just before
	// the surrounding function (in this case, `demonstrateDefer`) returns.

	// IMPORTANT POINT 1: Argument Evaluation
	// The arguments of a deferred function call are evaluated when the `defer`
	// statement itself is executed, NOT when the actual call is performed.
	// So, `initialValue` here will be its value at this point in the code (e.g., 10 if 10 was passed).
	defer fmt.Println("(Deferred Call 4) Value of 'initialValue' captured at defer time:", initialValue)

	// IMPORTANT POINT 2: LIFO (Last-In, First-Out) Order
	// If multiple `defer` statements are used in a function, they are pushed onto a stack.
	// When the function returns, the deferred calls are executed in LIFO order.
	// The last `defer` statement encountered will be the first one to execute.

	// This is the first `defer` statement we encounter among the next three.
	// It will be pushed onto the stack first, so it will execute LAST among these three.
	defer fmt.Println("(Deferred Call 3) This was deferred FIRST, so it runs THIRD (LIFO).")

	// This is the second `defer` statement.
	// It will be pushed onto the stack second, so it will execute SECOND among these three.
	defer fmt.Println("(Deferred Call 2) This was deferred SECOND, so it runs SECOND (LIFO).")

	// This is the third `defer` statement.
	// It will be pushed onto the stack last, so it will execute FIRST among these three.
	defer fmt.Println("(Deferred Call 1) This was deferred THIRD, so it runs FIRST (LIFO).")

	// Now, let's modify `initialValue` *after* the first defer statement
	// (Deferred Call 4) has already captured its original value.
	initialValue++ // If initialValue was 10, it's now 11.

	// These are normal print statements that execute in the order they appear.
	fmt.Println("Inside demonstrateDefer - Regular execution continues...")
	fmt.Println("Inside demonstrateDefer - Current value of 'initialValue':", initialValue)

	// When `demonstrateDefer` is about to return (after this line),
	// all the deferred fmt.Println calls will be executed in LIFO order.
	fmt.Println("Inside demonstrateDefer - Reaching the end of the function.")
}
