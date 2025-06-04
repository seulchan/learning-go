package main

import "fmt"

// --- Go Conditional Statements Tutorial ---
// This program demonstrates how to use conditional statements (`if`, `else if`, `else`, and `switch`)
// in Go to control the flow of execution based on different conditions.

func main() {
	fmt.Println("--- `if`, `else if`, `else` Statements ---")

	// --- Basic `if-else` statement ---
	// The `if` statement executes a block of code if a specified condition is true.
	// If the condition is false, the `else` block (if present) is executed.
	age := 25
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}
	fmt.Println() // Newline for readability

	// --- `if-else if-else` statement ---
	// This structure allows checking multiple conditions in sequence.
	// Go evaluates conditions from top to bottom.
	// The first `true` condition's block is executed, and the rest are skipped.
	// If no conditions are true, the `else` block (if present) is executed.
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		// This block is executed because score (85) is not >= 90, but it is >= 80.
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D")
	}
	fmt.Println()

	// --- Nested `if` statements ---
	// You can place `if` statements inside other `if` statements.
	// This is useful for checking sub-conditions.
	num := 15
	fmt.Printf("Checking number: %d\n", num)
	if num%2 == 0 { // Check if the number is even
		fmt.Println("The number is even.")
		if num%3 == 0 { // If even, check if it's also divisible by 3
			fmt.Println("The number is even and divisible by 3.")
		} else {
			fmt.Println("The number is even but not divisible by 3.")
		}
	} else { // The number is odd
		fmt.Println("The number is odd.")
		if num%3 == 0 { // If odd, check if it's also divisible by 3
			fmt.Println("The number is odd and divisible by 3.")
		} else {
			fmt.Println("The number is odd and not divisible by 3.")
		}
	}
	fmt.Println()

	// --- `if` with a short statement ---
	// An `if` statement can include a short variable declaration statement before the condition.
	// Variables declared here are scoped only to the `if` and any `else if`/`else` blocks.
	// This is often used for functions that return a value and an error.
	fmt.Println("--- `if` with a short statement ---")
	if n := 10; n%2 == 0 {
		fmt.Printf("%d is even (declared in if-statement).\n", n)
	} else {
		fmt.Printf("%d is odd (declared in if-statement).\n", n)
	}
	// fmt.Println(n) // This would cause a compile error: n is not defined in this scope.
	fmt.Println()

	// --- `switch` statement ---
	// The `switch` statement provides a multi-way branch.
	// It's often a cleaner alternative to a long series of `if-else if` statements.
	fmt.Println("--- `switch` statement ---")

	// Basic `switch` with an expression
	// The expression `day` is evaluated, and its value is compared against each `case`.
	day := "Monday"
	fmt.Printf("Today is %s.\n", day)
	switch day {
	case "Monday":
		fmt.Println("Start of the work week.")
	case "Friday":
		fmt.Println("End of the work week.")
	case "Saturday", "Sunday": // Multiple values can be listed in a single case
		fmt.Println("It's the weekend!")
	default: // The `default` case is executed if no other case matches. It's optional.
		fmt.Println("It's a regular day.")
	}
	fmt.Println()

	// `switch` without an expression (acts like `if/else if`)
	// If no expression is provided after `switch`, it defaults to `switch true`.
	// Each `case` is then an expression that evaluates to true or false.
	anotherNum := 2
	fmt.Printf("Checking anotherNum: %d\n", anotherNum)
	switch {
	case anotherNum < 0:
		fmt.Println("The number is negative.")
	case anotherNum == 0:
		fmt.Println("The number is zero.")
	case anotherNum > 0 && anotherNum <= 5:
		fmt.Println("The number is positive and between 1 and 5.")
		// `fallthrough` is used to execute the next case block unconditionally.
		// Use `fallthrough` with caution as it can make logic harder to follow.
		// In Go, cases do not automatically fall through like in C/C++.
		if anotherNum == 2 {
			fmt.Println("Specifically, the number is 2, demonstrating fallthrough target if it were used above.")
			// If fallthrough were uncommented in a case above that matched `anotherNum == 2`,
			// this block would also execute.
		}
	default:
		fmt.Println("The number is greater than 5 or some other condition.")
	}
	fmt.Println()

	// `switch` with `fallthrough`
	// The `fallthrough` statement transfers control to the next case block
	// *without* evaluating the case's condition.
	fmt.Println("--- `switch` with `fallthrough` ---")
	checkValue := 2
	fmt.Printf("Checking value with fallthrough: %d\n", checkValue)
	switch checkValue {
	case 1:
		fmt.Println("Value is 1.")
	case 2:
		fmt.Println("Value is 2.")
		fallthrough // Execution will continue to the next case (case 3)
	case 3:
		fmt.Println("Value is 3 (reached due to fallthrough from case 2 or if value was 3).")
	default:
		fmt.Println("Value is something else.")
	}
	fmt.Println()

	// Type `switch` (also known as a type assertion switch)
	// This is used to determine the dynamic type of an interface variable.
	fmt.Println("--- Type `switch` ---")
	checkType(10)
	checkType(3.14)
	checkType("Hello, Go!")
	checkType(true)
	fmt.Println("\nEnd of conditional statements demonstration.")
}

// `checkType` demonstrates a type switch.
// The `x.(type)` syntax is only allowed within a `switch` statement.
func checkType(x interface{}) {
	// `x` is an interface{}, meaning it can hold a value of any type.
	switch x.(type) {
	case int:
		fmt.Println("x is an int")
		// `fallthrough` is not permitted in a type switch.
	case float64:
		fmt.Println("x is a float64")
	case string:
		fmt.Println("x is a string")
	default:
		fmt.Printf("x is of an unhandled type: %T\n", x) // %T prints the type of the variable
	}
}
