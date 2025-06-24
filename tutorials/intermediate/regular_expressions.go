package main

import (
	"fmt"
	"regexp" // Import the "regexp" package for regular expression operations
)

// main is the entry point of the program.
func main() {
	fmt.Println("--- Go Regular Expressions Tutorial ---")

	// 1. Basic Matching: Check if a string matches a pattern
	demonstrateBasicMatching()

	// 2. Extracting Data: Use capturing groups to extract specific parts of a string
	demonstrateCapturingGroups()

	// 3. Replacing Substrings: Replace parts of a string that match a pattern
	demonstrateStringReplacement()

	// 4. Flags: Modify the behavior of regular expressions (e.g., case-insensitive matching)
	demonstrateRegexFlags()

	fmt.Println("\n--- End of Regular Expressions Tutorial ---")
}

// demonstrateBasicMatching shows how to check if a string matches a regular expression.
func demonstrateBasicMatching() {
	fmt.Println("\n--- 1. Basic Matching ---")

	// Define a regular expression pattern to match email addresses.
	// This pattern looks for:
	// - One or more characters that can be letters, numbers, dots, underscores, plus or minus signs, or percent symbols
	// - Followed by an "@" symbol
	// - Followed by one or more characters that can be letters, numbers, dots, or hyphens
	// - Followed by a dot "."
	// - Followed by two or more letters (for the top-level domain like "com", "org", etc.)
	emailPattern := `[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`

	// Compile the regular expression. `MustCompile` will panic (crash) if the pattern is invalid.
	// For patterns known at compile time, this is often preferred over `Compile`, which returns an error.
	emailRegex := regexp.MustCompile(emailPattern)

	// Test strings
	validEmail := "user@email.com"
	invalidEmail := "invalid_email"

	// Use `MatchString` to check if a string matches the regular expression.
	fmt.Printf("Is '%s' a valid email? %t\n", validEmail, emailRegex.MatchString(validEmail))     // Output: true
	fmt.Printf("Is '%s' a valid email? %t\n", invalidEmail, emailRegex.MatchString(invalidEmail)) // Output: false
}

// demonstrateCapturingGroups shows how to extract parts of a string using capturing groups.
func demonstrateCapturingGroups() {
	fmt.Println("\n--- 2. Capturing Groups ---")

	// Define a regular expression to match dates in YYYY-MM-DD format and capture the year, month, and day.
	// Parentheses `()` create "capturing groups" that allow us to extract matched parts of the string.
	datePattern := `(\d{4})-(\d{2})-(\d{2})`
	dateRegex := regexp.MustCompile(datePattern)

	// Test string
	dateString := "2024-07-30"

	// Use `FindStringSubmatch` to find the first match and capture the groups.
	// It returns a slice where the first element is the full match, and subsequent elements are the captured groups.
	submatches := dateRegex.FindStringSubmatch(dateString)

	// Check if a match was found
	if len(submatches) > 0 {
		fmt.Println("Full Date:", submatches[0]) // The entire matched string
		fmt.Println("Year:", submatches[1])      // The first capturing group: year
		fmt.Println("Month:", submatches[2])     // The second capturing group: month
		fmt.Println("Day:", submatches[3])       // The third capturing group: day
	}
}

// demonstrateStringReplacement shows how to replace substrings that match a pattern.
func demonstrateStringReplacement() {
	fmt.Println("\n--- 3. String Replacement ---")

	text := "Hello World"
	fmt.Println("Original string:", text)

	// Create a regex to match vowels (a, e, i, o, u).
	vowelRegex := regexp.MustCompile(`[aeiou]`)

	// Replace all vowels with an asterisk "*".
	// `ReplaceAllString` returns a new string with all matches replaced.
	replacedText := vowelRegex.ReplaceAllString(text, "*")
	fmt.Println("String with vowels replaced:", replacedText) // Output: H*ll* W*rld
}

// demonstrateRegexFlags shows how to use flags to modify the behavior of regular expressions.
func demonstrateRegexFlags() {
	fmt.Println("\n--- 4. Regular Expression Flags ---")

	text := "Golang is great"
	fmt.Println("Text:", text)

	// Create a case-insensitive regex to match "go" (or "Go", "gO", "GO").
	caseInsensitiveRegex := regexp.MustCompile(`(?i)go`)

	fmt.Printf("Case-insensitive match for 'go': %t\n", caseInsensitiveRegex.MatchString(text)) // Output: true
}
