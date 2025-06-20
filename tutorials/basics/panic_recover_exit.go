package main

import (
	"fmt"
	"os"
)

// --- Go Panic, Recover, and os.Exit Tutorial ---
// This program demonstrates three ways to handle or cause abrupt program termination
// and how deferred functions interact with them.

func main() {
	fmt.Println("--- Starting Go Termination Concepts Tutorial ---")

	// Section 1: Demonstrating Panic and Defer
	// Panic is used for unrecoverable errors. It stops the normal flow,
	// runs deferred functions, and then crashes the program (unless recovered).
	demonstratePanicAndDefer()
	fmt.Println("--- Finished Panic Demonstration ---") // This line is reached because demonstratePanicAndDefer recovers internally for the tutorial

	fmt.Println() // Add a newline for clarity

	// Section 2: Demonstrating Recover
	// Recover is used to regain control after a panic. It only works inside a deferred function.
	demonstrateRecovery()
	fmt.Println("--- Finished Recover Demonstration ---") // This line is reached because demonstrateRecovery handles the panic

	fmt.Println() // Add a newline for clarity

	// Section 3: Demonstrating os.Exit
	// os.Exit terminates the program immediately without running deferred functions.
	demonstrateExit()
	// Note: The line below will NEVER be reached because os.Exit terminates the program.
	fmt.Println("--- Finished os.Exit Demonstration ---") // This line is unreachable

	// Note: Any code here after demonstrateExit() will not execute.
}

// simulateOperationWithPanic simulates a function that might encounter a critical error
// and uses panic. It also shows how deferred functions behave during a panic.
func simulateOperationWithPanic(input int) {
	// Deferred functions are pushed onto a stack. They are executed in LIFO (Last-In, First-Out) order
	// just before the surrounding function returns, OR when a panic occurs.
	defer fmt.Println("  [simulateOperationWithPanic] Deferred 1: This runs before returning or panicking.")
	defer fmt.Println("  [simulateOperationWithPanic] Deferred 2: This runs before Deferred 1.")

	fmt.Printf("  [simulateOperationWithPanic] Processing input: %d\n", input)

	// Check for a condition that we consider unrecoverable for this operation.
	if input < 0 {
		fmt.Println("  [simulateOperationWithPanic] Input is negative. Initiating panic!")
		// Calling panic stops the normal execution flow immediately.
		// The line below 'panic' is never reached.
		panic("Negative input is not allowed.")
		// fmt.Println("  [simulateOperationWithPanic] After Panic: This line is never reached.")
		// defer fmt.Println("  [simulateOperationWithPanic] Deferred 3: This defer is never registered because the line is unreachable.")
	}

	// This line is only reached if the input is non-negative and no panic occurs.
	fmt.Println("  [simulateOperationWithPanic] Input processed successfully.")
	// Deferred functions (Deferred 2, then Deferred 1) will run now before the function returns normally.
}

// demonstratePanicAndDefer shows what happens when a panic occurs and is NOT recovered
// at the point of the panic. It highlights that defers DO run.
// For the purpose of this tutorial, we wrap the panicking call in a local
// recover block so the program can continue to the next demonstration.
// In a real application, an unrecovered panic in main would crash the program.
func demonstratePanicAndDefer() {
	fmt.Println("Section 1: Demonstrating Panic and Defer...")

	// Example 1: Call with valid input (no panic).
	fmt.Println("Calling simulateOperationWithPanic with positive input (10)...")
	simulateOperationWithPanic(10)
	fmt.Println("simulateOperationWithPanic with 10 finished normally.")
	fmt.Println() // Newline

	// Example 2: Call with invalid input (will panic).
	fmt.Println("Calling simulateOperationWithPanic with negative input (-5)...")

	// We use an anonymous function and a local defer+recover here *only* to catch
	// the panic for this demonstration section, allowing the rest of main to run.
	// In a typical program, this unrecovered panic would stop the entire application.
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  [demonstratePanicAndDefer] Caught panic for tutorial purposes: %v\n", r)
				fmt.Println("  [demonstratePanicAndDefer] Execution continues after recovery within this block.")
			}
		}() // The deferred function is defined and immediately scheduled.

		// This call will cause a panic.
		simulateOperationWithPanic(-5)

		// This line is inside the anonymous function but after the panicking call.
		// It will NOT be reached because the panic stops execution flow within this function.
		fmt.Println("  [demonstratePanicAndDefer] This line is after the panicking call and is not reached.")
	}() // The anonymous function is immediately executed.

	// Execution continues here after the anonymous function's deferred recover handles the panic.
	fmt.Println("Back in demonstratePanicAndDefer after the panicking call block.")
}

// demonstrateRecovery shows how to use recover within a deferred function
// to gracefully handle a panic and prevent program termination.
func demonstrateRecovery() {
	fmt.Println("Section 2: Demonstrating Recover...")

	// Schedule a deferred function that attempts to recover from a panic.
	defer func() {
		// recover() is called inside the deferred function.
		// If a panic occurred, recover() stops the panicking sequence and returns the value passed to panic().
		// If no panic occurred, recover() returns nil.
		if r := recover(); r != nil {
			// If r is not nil, it means we successfully caught a panic.
			fmt.Printf("  [demonstrateRecovery] Successfully recovered from panic: %v\n", r)
			fmt.Println("  [demonstrateRecovery] Execution continues after the deferred function finishes.")
		} else {
			// This part would execute if no panic happened in this function.
			fmt.Println("  [demonstrateRecovery] Deferred function ran, but no panic occurred.")
		}
	}() // Define and immediately schedule the deferred function.

	fmt.Println("  [demonstrateRecovery] Starting operation that will panic...")

	// This line will cause a panic.
	panic("Something critical went wrong during operation.")

	// This line is immediately after the panic call and will NOT be reached.
	fmt.Println("  [demonstrateRecovery] This line is after the panic call and is not reached.")

	// After the panic, the deferred function runs, recover() is called,
	// the panic is stopped, and execution resumes *after* the deferred function.
}

// demonstrateExit shows how os.Exit terminates the program immediately.
// It highlights that deferred functions are NOT executed with os.Exit.
func demonstrateExit() {
	fmt.Println("Section 3: Demonstrating os.Exit...")

	// Schedule a deferred function.
	defer fmt.Println("  [demonstrateExit] This deferred function will NOT run because of os.Exit.")

	fmt.Println("  [demonstrateExit] About to call os.Exit(0)...") // Using 0 for success exit code in this demo

	// os.Exit terminates the program immediately.
	// It does NOT run deferred functions.
	// It does NOT unwind the stack.
	// The integer argument is the exit code (0 typically means success, non-zero means error).
	os.Exit(0)

	// This line is immediately after os.Exit and will NOT be reached.
	fmt.Println("  [demonstrateExit] This line is after os.Exit and is not reached.")
}
