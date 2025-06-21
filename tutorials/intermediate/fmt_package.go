// --- Go `fmt` Package Tutorial ---
//
// The `fmt` package is one of the most frequently used packages in Go. It implements
// functions for formatted I/O (input/output), similar to C's `printf` and `scanf`.
// This tutorial covers the most common functions for printing to the console,
// formatting strings, and reading user input.
package main

import "fmt"

// main is the entry point of our program. It calls various demonstration functions.
func main() {
	fmt.Println("--- Go `fmt` Package Tutorial ---")

	// Part 1: Printing to the console.
	demonstratePrintingFunctions()

	// Part 2: Formatting data into strings (without printing).
	demonstrateStringFormattingFunctions()

	// Part 3: Reading input from the user.
	demonstrateScanningFunctions()

	// Part 4: Creating formatted error messages.
	demonstrateErrorFormatting()

	fmt.Println("\n--- End of `fmt` Package Tutorial ---")
}

// demonstratePrintingFunctions shows the use of Print, Println, and Printf.
func demonstratePrintingFunctions() {
	fmt.Println("\n--- 1. Printing to the Console ---")
	fmt.Println("These functions write data directly to standard output (your terminal).")

	// fmt.Print: Writes its arguments to the console. It does NOT add spaces between
	// arguments or a newline character at the end.
	fmt.Println("\n* Using fmt.Print:")
	fmt.Print("Hello,")
	fmt.Print(" world!")
	fmt.Print(" Numbers:", 123, 456, "\n") // We must add `\n` manually for a new line.

	// fmt.Println: Similar to Print, but it adds spaces between its arguments
	// and ALWAYS adds a newline character at the end. This is very common for quick debugging.
	fmt.Println("\n* Using fmt.Println:")
	fmt.Println("Hello,", "world!")
	fmt.Println("Numbers:", 123, 456)

	// fmt.Printf: "Print Formatted". This is the most powerful printing function.
	// It takes a "format string" with "verbs" (like %s, %d) that act as placeholders
	// for the values that follow. It does NOT add a newline unless you specify `\n`.
	// For a detailed guide on formatting verbs, see the `formatting_verbs.go` tutorial.
	fmt.Println("\n* Using fmt.Printf:")
	name := "Alice"
	age := 30
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("The number %d in binary is %b, and in hexadecimal is %X.\n", age, age, age)
}

// demonstrateStringFormattingFunctions shows how to create formatted strings without printing them.
func demonstrateStringFormattingFunctions() {
	fmt.Println("\n--- 2. Formatting to a String (Sprint, Sprintln, Sprintf) ---")
	fmt.Println("These functions are counterparts to the Print functions, but instead of")
	fmt.Println("writing to the console, they return the formatted data as a string.")

	// fmt.Sprint: Formats using default formats and returns the resulting string.
	// It adds spaces between arguments when neither are strings.
	s1 := fmt.Sprint("Hello, ", "world! ", 123, 456)
	fmt.Println("Result of fmt.Sprint:", s1)

	// fmt.Sprintln: Formats using default formats, adds spaces between arguments,
	// adds a newline at the end, and returns the resulting string.
	s2 := fmt.Sprintln("Hello,", "world!", 123, 456)
	fmt.Print("Result of fmt.Sprintln (contains a newline):", s2) // Using Print to show the newline from s2

	// fmt.Sprintf: Formats according to a format specifier and returns the resulting string.
	// This is extremely useful for building dynamic strings.
	name := "Bob"
	age := 42
	formattedString := fmt.Sprintf("User Profile -> Name: %s, Age: %d", name, age)
	fmt.Println("Result of fmt.Sprintf:", formattedString)
}

// demonstrateScanningFunctions shows how to read user input from the console.
func demonstrateScanningFunctions() {
	fmt.Println("\n--- 3. Scanning User Input (Scan, Scanln, Scanf) ---")
	fmt.Println("These functions read input from standard input (your keyboard).")
	fmt.Println("NOTE: This part of the tutorial is interactive. You will be prompted for input.")

	var userName string
	var userAge int

	// There are three main scanning functions:
	// - fmt.Scan: Stops scanning at the first space. Reads "John Doe" as "John".
	// - fmt.Scanln: Stops scanning at a newline. Reads the whole line.
	// - fmt.Scanf: Reads input according to a format string.
	// We will use Scanln here as it's often intuitive for simple inputs.

	fmt.Print("\nPlease enter your name: ")
	// We use the address-of operator `&` to pass the memory address of our variables
	// to fmt.Scanln. This allows the function to modify the original variables.
	_, err := fmt.Scanln(&userName)
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}

	fmt.Print("Please enter your age: ")
	_, err = fmt.Scanln(&userAge)
	if err != nil {
		// This handles cases where the user enters text instead of a number.
		fmt.Println("Invalid input for age. Please enter a number. Error:", err)
		return
	}

	fmt.Printf("\nHello, %s! You are %d years old.\n", userName, userAge)
}

// demonstrateErrorFormatting shows how to create custom, formatted error messages.
func demonstrateErrorFormatting() {
	fmt.Println("\n--- 4. Creating Formatted Errors (Errorf) ---")
	fmt.Println("`fmt.Errorf` is like `fmt.Sprintf`, but it returns an `error` type.")
	fmt.Println("This is the standard way to create descriptive, dynamic error messages in Go.")

	err := checkEligibility(15)
	if err != nil {
		// The `err` variable holds the error returned by the function.
		fmt.Println("Eligibility Check Failed. Error:", err)
	}

	err = checkEligibility(21)
	if err == nil {
		// `nil` is the zero value for an error, meaning "no error occurred".
		fmt.Println("Eligibility Check Passed for age 21.")
	}
}

// checkEligibility checks if an age meets a minimum requirement.
// It returns an error if the age is too low, and `nil` otherwise.
func checkEligibility(age int) error {
	const minimumAge = 18
	if age < minimumAge {
		// fmt.Errorf creates an error value with a formatted message.
		// This is much more informative than a generic error.
		return fmt.Errorf("age %d is below the minimum requirement of %d", age, minimumAge)
	}
	return nil // Return nil to indicate success (no error).
}
