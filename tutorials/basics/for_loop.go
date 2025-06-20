package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// --- Go `for` Loop Tutorial ---
	// The `for` loop is Go's only looping construct. It's versatile and can be used
	// in several ways, similar to `for`, `while`, and `do-while` loops in other languages.

	// 1. Basic `for` loop (C-style)
	// This form has three components separated by semicolons:
	// - init statement: `i := 1` (executed once before the first iteration)
	// - condition expression: `i <= 5` (evaluated before every iteration; loop continues if true)
	// - post statement: `i++` (executed at the end of every iteration)
	fmt.Println("--- Basic for loop (1 to 5) ---")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	fmt.Println() // Add a newline for better output separation

	// 2. `for...range` loop (iterating over collections)
	// This form is used to iterate over elements in data structures like slices, arrays,
	// strings, maps, and channels.
	fmt.Println("--- for...range with a slice ---")
	numbers := []int{10, 20, 30, 40, 50} // Example slice
	for index, value := range numbers {
		// `range numbers` returns two values for each iteration:
		// - `index`: The current element's index.
		// - `value`: A copy of the element at that index.
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	fmt.Println()

	// If you only need the value, you can ignore the index using the blank identifier `_`.
	fmt.Println("--- for...range with a slice (value only) ---")
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}
	fmt.Println()

	// If you only need the index, you can omit the second variable.
	fmt.Println("--- for...range with a slice (index only) ---")
	for index := range numbers { // Note: `value` is omitted
		fmt.Printf("Index: %d\n", index)
	}
	fmt.Println()

	// 2a. `for...range` with a string
	// When ranging over a string, it iterates over Unicode code points (runes).
	// The first value is the starting byte index of the rune, and the second is the rune itself.
	fmt.Println("--- for...range with a string ---")
	message := "Hi, Go! 👋"
	for index, charRune := range message {
		fmt.Printf("Byte Index: %d, Rune: %c (Unicode: %U)\n", index, charRune, charRune)
	}
	fmt.Println()

	// 2b. `for...range` with a map
	// When ranging over a map, it iterates over key-value pairs in an unspecified order.
	fmt.Println("--- for...range with a map ---")
	capitals := map[string]string{
		"USA":    "Washington D.C.",
		"France": "Paris",
		"Japan":  "Tokyo",
	}
	for country, capital := range capitals {
		fmt.Printf("The capital of %s is %s\n", country, capital)
	}
	// To iterate in a specific order, you'd typically get the keys, sort them, then iterate.
	fmt.Println()

	// 3. `for` loop with `continue` and `break` statements
	// These statements control the flow of the loop.
	fmt.Println("--- for loop with continue and break ---")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			// `continue` skips the rest of the current iteration
			// and proceeds to the next iteration.
			continue // Skip even numbers, go to next i
		}
		fmt.Println("Odd Number:", i) // This line is only reached for odd numbers

		if i == 7 { // Example condition to exit early
			// `break` terminates the loop entirely.
			fmt.Println("Breaking loop at 7.")
			break // Exit the for loop
		}
	}
	fmt.Println()

	// 4. Nested `for` loops
	// You can have loops inside other loops.
	// This example prints a simple triangle pattern.
	fmt.Println("--- Nested for loops (triangle pattern) ---")
	rows := 5
	for i := 1; i <= rows; i++ { // Outer loop: controls the number of rows
		// Inner loop 1: prints spaces for formatting the triangle
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		// Inner loop 2: prints asterisks for the triangle
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println() // Move to the next line after each row is printed
	}
	fmt.Println()

	// 5. `for...range` with an integer (Go 1.22+)
	// As of Go 1.22, `for...range` can iterate over an integer `n`,
	// yielding values from 0 up to (but not including) `n`.
	// If only one variable is used, it gets the iteration value (0, 1, 2,...).
	fmt.Println("--- for...range with an integer (Go 1.22+) ---")
	// This loop will iterate with `val` taking values 0, 1, 2, 3, 4.
	// Equivalent to: for val := 0; val < 5; val++
	for val := range 5 { // Iterates 0, 1, 2, 3, 4
		fmt.Printf("Range over integer, val = %d\n", val)
	}
	fmt.Println()

	// 6. `for` loop as a `while` loop
	// Go doesn't have a separate `while` keyword.
	// A `for` loop with only a condition acts like a `while` loop.
	fmt.Println("--- for loop as a while loop ---")
	count := 1
	for count <= 3 { // Loop continues as long as `count <= 3`
		fmt.Println("Iteration (while-style):", count)
		count++ // Increment is done inside the loop body
	}
	fmt.Println()

	// 7. Infinite `for` loop
	// A `for` loop without any condition is an infinite loop.
	// It must be terminated explicitly using `break`, `return`, or `panic`,
	// or by the program exiting.
	fmt.Println("--- Infinite for loop (with a break) ---")
	loopCounter := 0
	for {
		fmt.Println("Infinite loop iteration:", loopCounter)
		loopCounter++
		if loopCounter >= 2 { // Condition to exit the loop
			fmt.Println("Breaking the infinite loop.")
			break // Essential to prevent an actual infinite loop in this example
		}
	}
	fmt.Println()

	// 8. Practical Example: Guess the Number Game
	// This demonstrates a `for` loop used for interactive input until a condition is met.
	fmt.Println("--- Guess the Number Game ---")

	// Seed the random number generator.
	// Using time.Now().UnixNano() provides a different seed each time the program runs.
	// Note: As of Go 1.20, the global functions in "math/rand" (like rand.Intn) are
	// automatically seeded and are safe for concurrent use.
	// Creating a local Rand instance is useful if you need a specific seed for reproducibility
	// or a separate stream of random numbers.
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)

	// Generate a random number between 1 and 100 (inclusive).
	// randomGenerator.Intn(100) generates a number from 0 to 99.
	targetNumber := randomGenerator.Intn(100) + 1

	fmt.Println("I've picked a number between 1 and 100. Try to guess it!")
	var userGuess int

	// This is an infinite loop that will only terminate when the user guesses correctly.
	for {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanf("%d", &userGuess) // Read user input
		if err != nil {
			fmt.Println("Invalid input. Please enter a whole number.")
			// Attempt to clear the invalid input from the buffer.
			// A more robust solution might use bufio.NewReader(os.Stdin).ReadString('\n').
			var discard string
			fmt.Scanln(&discard) // Try to consume the rest of the bad input line
			continue             // Skip to the next iteration
		}

		if userGuess < targetNumber {
			fmt.Println("Too low! Try again.")
		} else if userGuess > targetNumber {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("Congratulations! You've guessed the number!")
			break // Exit the loop when the guess is correct
		}
	}
	fmt.Println("\nEnd of for loop examples.") // Added newline for cleaner final output
}
