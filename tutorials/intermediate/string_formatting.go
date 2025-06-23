package main

import "fmt"

// --- Go String Formatting and Literals Tutorial ---
//
// This tutorial focuses on specific aspects of string formatting using `fmt.Printf`
// such as padding and alignment, and clarifies the different types of string literals
// available in Go.
//
// For a more comprehensive guide on all `fmt.Printf` verbs, please refer to
// the `formatting_verbs.go` tutorial.
// For a deeper dive into Go's string representation and runes, see
// the `strings_runes.go` tutorial.

func main() {
	fmt.Println("--- Go String Formatting and Literals Tutorial ---")

	demonstrateIntegerPadding()
	demonstrateStringAlignment()
	demonstrateStringLiterals()

	fmt.Println("\n--- End of String Formatting and Literals Tutorial ---")
}

// demonstrateIntegerPadding shows how to pad integers with leading zeros or spaces.
func demonstrateIntegerPadding() {
	fmt.Println("\n--- 1. Integer Padding with fmt.Printf ---")

	// %d: Standard decimal integer.
	// %05d: Formats an integer to a width of 5 characters, padding with leading zeros if necessary.
	// If the number is shorter than 5 digits, zeros are added to the left.
	// If the number is longer, it's printed fully.
	number := 42
	fmt.Printf("Original number: %d\n", number)
	fmt.Printf("Padded with leading zeros (%%05d): %05d\n", number) // Output: 00042

	largeNumber := 123456
	fmt.Printf("Large number (%%05d): %05d\n", largeNumber) // Output: 123456 (width is minimum)

	// %5d: Formats an integer to a width of 5 characters, padding with leading spaces.
	fmt.Printf("Padded with leading spaces (%%5d): %5d\n", number) // Output:    42
}

// demonstrateStringAlignment shows how to align strings within a fixed width.
func demonstrateStringAlignment() {
	fmt.Println("\n--- 2. String Alignment with fmt.Printf ---")

	text := "Hello"
	fmt.Printf("Original string: \"%s\"\n", text)

	// %10s: Formats a string to a width of 10 characters, right-justified (padded with spaces on the left).
	fmt.Printf("Right-justified (%%10s): |%10s|\n", text) // Output: |     Hello|

	// %-10s: Formats a string to a width of 10 characters, left-justified (padded with spaces on the right).
	// The hyphen (`-`) flag indicates left justification.
	fmt.Printf("Left-justified (%%-10s): |%-10s|\n", text) // Output: |Hello     |

	longText := "GoLang"
	// If the string is longer than the specified width, it's printed fully.
	fmt.Printf("Long string right-justified (%%5s): |%5s|\n", longText)  // Output: |GoLang|
	fmt.Printf("Long string left-justified (%%-5s): |%-5s|\n", longText) // Output: |GoLang|
}

// demonstrateStringLiterals explains the two types of string literals in Go.
func demonstrateStringLiterals() {
	fmt.Println("\n--- 3. String Literals: Interpreted vs. Raw ---")

	// a) Interpreted String Literals (using double quotes "")
	// These strings process escape sequences like `\n` (newline), `\t` (tab), `\"` (double quote).
	// They are the most common type of string literal.
	interpretedString := "Hello,\n\tWorld with a \"quote\"!"
	fmt.Println("Interpreted String Literal (double quotes):")
	fmt.Println(interpretedString)

	// b) Raw String Literals (using backticks ``)
	// These strings treat all characters literally. No escape sequences are processed.
	// They are useful for multi-line strings, regular expressions, or file paths
	// where backslashes should be treated as literal characters.
	rawString := `This is a raw string.
It can span multiple lines.
Backslashes like \n or \t are treated literally.
No "quotes" need escaping either.`
	fmt.Println("\nRaw String Literal (backticks):")
	fmt.Println(rawString)

	// Example of a common use case for raw strings: file paths on Windows or regex patterns.
	windowsPath := `C:\Program Files\Go\bin`
	fmt.Printf("\nRaw string for a Windows path: %s\n", windowsPath)

	regexPattern := `\d{3}-\d{2}-\d{4}` // Matches a Social Security Number pattern
	fmt.Printf("Raw string for a regex pattern: %s\n", regexPattern)
}
