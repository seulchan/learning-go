package main

import "fmt"

// --- Go Formatting Verbs Tutorial ---
//
// This file provides a comprehensive, beginner-friendly guide to the formatting "verbs"
// used with `fmt.Printf` and other functions in Go's `fmt` package.
//
// Formatting verbs are special placeholders (like %v, %d, %s) that tell `fmt.Printf`
// how to format and display data of different types (integers, strings, structs, etc.).
// Understanding these verbs is essential for producing clear, well-formatted output.

// To make the tutorial clearer, we'll define a simple struct.
// This will help demonstrate how formatting verbs work with custom data types.
type User struct {
	ID   int
	Name string
}

func main() {
	fmt.Println("--- Go `fmt` Package Formatting Verbs Tutorial ---")

	demonstrateGeneralVerbs()
	demonstrateIntegerVerbs()
	demonstrateFloatVerbs()
	demonstrateStringVerbs()
	demonstrateBooleanVerbs()
	demonstratePointerAndStructVerbs()

	fmt.Println("\n--- End of Formatting Verbs Tutorial ---")
}

// demonstrateGeneralVerbs shows verbs that can be used with almost any data type.
func demonstrateGeneralVerbs() {
	fmt.Println("\n--- 1. General Purpose Verbs ---")
	fmt.Println("These verbs work with most data types.")

	sampleString := "Hello Go!"
	sampleNumber := 42

	// %v: Prints the value in a default, natural format.
	fmt.Printf("%%v (default format) for string: %v\n", sampleString)
	fmt.Printf("%%v (default format) for number: %v\n", sampleNumber)

	// %#v: Prints the value in a Go-syntax representation. This is very useful for
	// debugging, as it shows you how you would write the value in Go code.
	fmt.Printf("%%#v (Go-syntax format) for string: %#v\n", sampleString) // Note the quotes
	fmt.Printf("%%#v (Go-syntax format) for number: %#v\n", sampleNumber)

	// %T: Prints the type of the value.
	fmt.Printf("%%T (type) for string: %T\n", sampleString)
	fmt.Printf("%%T (type) for number: %T\n", sampleNumber)

	// %%: To print a literal percent sign, you use two of them.
	fmt.Printf("To print a percent sign (%%), you write %%%%.\n")
}

// demonstrateIntegerVerbs covers the various ways to format integer numbers.
func demonstrateIntegerVerbs() {
	fmt.Println("\n--- 2. Integer Formatting Verbs ---")
	sampleInteger := 255

	fmt.Printf("Formatting the integer: %d\n", sampleInteger)

	// %d: Formats the integer in base-10 (decimal).
	fmt.Printf("%%d (base-10 decimal): %d\n", sampleInteger)

	// %+d: Always shows the sign (+ or -).
	fmt.Printf("%%+d (decimal with sign): %+d\n", sampleInteger)

	// %b: Formats the integer in base-2 (binary).
	fmt.Printf("%%b (base-2 binary): %b\n", sampleInteger)

	// %o: Formats the integer in base-8 (octal).
	fmt.Printf("%%o (base-8 octal): %o\n", sampleInteger)

	// %O: Formats in base-8 with a leading "0o".
	fmt.Printf("%%O (base-8 with 0o prefix): %O\n", sampleInteger)

	// %x, %X: Formats in base-16 (hexadecimal), with lowercase or uppercase letters.
	fmt.Printf("%%x (base-16 hex, lowercase): %x\n", sampleInteger)
	fmt.Printf("%%X (base-16 hex, uppercase): %X\n", sampleInteger)

	// %#x: Formats in base-16 with a leading "0x".
	fmt.Printf("%%#x (base-16 with 0x prefix): %#x\n", sampleInteger)

	// Padding: You can specify a width for the output.
	// The number will be padded with spaces to fill the width.
	// A positive number means right-justified, a negative number means left-justified.
	fmt.Printf("%%4d (width 4, right-justified): |%4d|\n", 42)
	fmt.Printf("%%-4d (width 4, left-justified): |%-4d|\n", 42)

	// You can also pad with leading zeros instead of spaces.
	fmt.Printf("%%04d (width 4, padded with zeros): %04d\n", 42)
}

// demonstrateFloatVerbs covers formatting for floating-point numbers.
func demonstrateFloatVerbs() {
	fmt.Println("\n--- 3. Floating-Point Formatting Verbs ---")
	sampleFloat := 12345.6789

	fmt.Printf("Formatting the float: %f\n", sampleFloat)

	// %f: Standard decimal point notation.
	fmt.Printf("%%f (standard decimal): %f\n", sampleFloat)

	// %e, %E: Scientific notation (e.g., 1.2345e+04).
	fmt.Printf("%%e (scientific notation): %e\n", sampleFloat)
	fmt.Printf("%%E (scientific notation, uppercase E): %E\n", sampleFloat)

	// Precision: You can control the number of digits after the decimal point.
	fmt.Printf("%%.2f (precision 2): %.2f\n", sampleFloat) // Rounds the number

	// Width and Precision: You can specify both total width and precision.
	// %9.2f means a total width of 9 characters, with 2 digits after the decimal.
	fmt.Printf("%%9.2f (width 9, precision 2): |%9.2f|\n", sampleFloat)
	fmt.Printf("%%9.f (width 9, precision 0): |%9.f|\n", sampleFloat)

	// %g, %G: Chooses the more compact representation: %f or %e.
	// It's useful for general-purpose float printing.
	fmt.Printf("%%g (compact format): %g\n", sampleFloat)
	fmt.Printf("%%g on a small number: %g\n", 0.000012345)
}

// demonstrateStringVerbs covers formatting for strings.
func demonstrateStringVerbs() {
	fmt.Println("\n--- 4. String and Byte Slice Formatting Verbs ---")
	sampleString := "Go World"

	fmt.Printf("Formatting the string: \"%s\"\n", sampleString)

	// %s: Prints the string's contents.
	fmt.Printf("%%s (plain string): %s\n", sampleString)

	// %q: Prints a double-quoted string, safely escaping any special characters.
	// This is very useful for showing the exact string value.
	fmt.Printf("%%q (quoted string): %q\n", "A string with a \"quote\" and a \nnewline.")

	// Padding and justification work for strings too.
	fmt.Printf("%%12s (width 12, right-justified): |%12s|\n", sampleString)
	fmt.Printf("%%-12s (width 12, left-justified): |%-12s|\n", sampleString)

	// %x, %X: For strings and byte slices, these print the hexadecimal representation of the bytes.
	fmt.Printf("%%x (hex dump of bytes): %x\n", sampleString)
	// The space flag adds spaces between the hex bytes, which can improve readability.
	fmt.Printf("%% x (hex dump with spaces): % x\n", sampleString)
}

// demonstrateBooleanVerbs shows how to format boolean values.
func demonstrateBooleanVerbs() {
	fmt.Println("\n--- 5. Boolean Formatting Verbs ---")

	// %t: Prints the word "true" or "false".
	// This is the standard way to format booleans.
	fmt.Printf("%%t for true: %t\n", true)
	fmt.Printf("%%t for false: %t\n", false)
}

// demonstratePointerAndStructVerbs shows verbs for pointers and custom structs.
func demonstratePointerAndStructVerbs() {
	fmt.Println("\n--- 6. Pointer and Struct Formatting Verbs ---")

	// Create an instance of our User struct.
	user := User{ID: 101, Name: "Alice"}

	// %p: Prints the memory address of a pointer in hexadecimal format.
	// You must pass a pointer (or the address of a variable using `&`).
	fmt.Printf("%%p (pointer address): %p\n", &user)

	// Let's see how general verbs work on a struct.
	fmt.Println("\nFormatting a struct variable:")

	// %v on a struct prints the field values, enclosed in {}.
	fmt.Printf("%%v (default struct format): %v\n", user)

	// %+v is like %v, but it also includes the field names. Very handy for debugging!
	fmt.Printf("%%+v (struct with field names): %+v\n", user)

	// %#v prints the struct in Go-syntax, including the type name.
	// This shows you exactly how to declare this struct instance in code.
	fmt.Printf("%%#v (Go-syntax for struct): %#v\n", user)
}
