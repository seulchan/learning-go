package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// --- Go String Functions Tutorial ---
//
// This tutorial provides a comprehensive overview of common and useful functions for
// working with strings in Go. It covers everything from basic operations like
// concatenation and length to more advanced topics like splitting, joining, searching,
// replacing, and efficient string building with `strings.Builder`.

func main() {
	fmt.Println("--- Go String Functions Tutorial ---")

	demonstrateBasicOperations()
	demonstrateConversion()
	demonstrateSplittingAndJoining()
	demonstrateSearchingAndReplacing()
	demonstrateCaseAndTrimming()
	demonstrateAdvancedTechniques()

	fmt.Println("\n--- End of String Functions Tutorial ---")
}

// demonstrateBasicOperations covers length, concatenation, and slicing.
func demonstrateBasicOperations() {
	fmt.Println("\n--- 1. Basic Operations: Length, Concatenation, Slicing ---")

	greeting := "Hello, Go!"

	// `len()` returns the number of bytes in a string, not necessarily the number of characters.
	// This is important for multi-byte characters (like in UTF-8).
	fmt.Printf("The string is: \"%s\"\n", greeting)
	fmt.Printf("Length in bytes (len()): %d\n", len(greeting))

	// Strings can be concatenated using the `+` operator.
	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName
	fmt.Printf("Concatenated name: %s\n", fullName)

	// Slicing a string creates a new string from a range of byte indices.
	// `greeting[0:5]` extracts bytes from index 0 up to (but not including) 5.
	helloSlice := greeting[0:5]
	fmt.Printf("Slice `greeting[0:5]`: %s\n", helloSlice)
}

// demonstrateConversion shows how to convert other types to and from strings.
func demonstrateConversion() {
	fmt.Println("\n--- 2. String Conversion ---")

	// The `strconv` package is the standard way to convert types.
	// `Itoa` stands for "Integer to ASCII" (i.e., integer to string).
	age := 30
	ageAsString := strconv.Itoa(age)
	fmt.Printf("The integer %d as a string is \"%s\"\n", age, ageAsString)

	// `Atoi` (ASCII to Integer) does the reverse. It can return an error.
	ageFromString, err := strconv.Atoi("30")
	if err == nil {
		fmt.Printf("The string \"%s\" as an integer is %d\n", "30", ageFromString)
	}
}

// demonstrateSplittingAndJoining shows how to break strings apart and put them together.
func demonstrateSplittingAndJoining() {
	fmt.Println("\n--- 3. Splitting and Joining Strings ---")

	// `strings.Split` divides a string into a slice of substrings based on a separator.
	csvData := "apple,orange,banana"
	fruitSlice := strings.Split(csvData, ",")
	fmt.Printf("Splitting \"%s\" by \",\": %v\n", csvData, fruitSlice)
	fmt.Printf("The second fruit is: %s\n", fruitSlice[1])

	// `strings.Join` does the opposite: it concatenates elements of a string slice
	// into a single string, with a separator between each element.
	countries := []string{"Germany", "France", "Italy"}
	joinedCountries := strings.Join(countries, " | ")
	fmt.Printf("Joining a slice with \" | \": %s\n", joinedCountries)
}

// demonstrateSearchingAndReplacing covers finding and modifying substrings.
func demonstrateSearchingAndReplacing() {
	fmt.Println("\n--- 4. Searching, Replacing, and Checking Substrings ---")

	sentence := "A quick brown fox jumps over the lazy fox."
	fmt.Printf("Original sentence: \"%s\"\n", sentence)

	// `strings.Contains` checks if a substring is present.
	fmt.Printf("Contains \"brown\"? %t\n", strings.Contains(sentence, "brown"))
	fmt.Printf("Contains \"cat\"? %t\n", strings.Contains(sentence, "cat"))

	// `strings.HasPrefix` and `strings.HasSuffix` check the start and end of a string.
	fmt.Printf("Has prefix \"A quick\"? %t\n", strings.HasPrefix(sentence, "A quick"))
	fmt.Printf("Has suffix \"lazy dog.\"? %t\n", strings.HasSuffix(sentence, "lazy dog."))

	// `strings.Count` counts non-overlapping instances of a substring.
	fmt.Printf("Count of \"fox\": %d\n", strings.Count(sentence, "fox"))

	// `strings.Replace` replaces instances of a substring.
	// The last argument `n` specifies the max number of replacements. If `n < 0`, all are replaced.
	replacedOnce := strings.Replace(sentence, "fox", "wolf", 1)
	fmt.Printf("Replaced \"fox\" once: \"%s\"\n", replacedOnce)
	replacedAll := strings.Replace(sentence, "fox", "wolf", -1)
	fmt.Printf("Replaced all \"fox\": \"%s\"\n", replacedAll)
}

// demonstrateCaseAndTrimming shows text manipulation functions.
func demonstrateCaseAndTrimming() {
	fmt.Println("\n--- 5. Case Conversion, Trimming, and Repeating ---")

	messyData := "   Hello World!   "
	fmt.Printf("Original messy data: \"%s\"\n", messyData)

	// `strings.TrimSpace` removes all leading and trailing white space.
	trimmed := strings.TrimSpace(messyData)
	fmt.Printf("Trimmed: \"%s\"\n", trimmed)

	// `strings.ToLower` and `strings.ToUpper` convert the case of the string.
	fmt.Printf("Lowercase: \"%s\"\n", strings.ToLower(trimmed))
	fmt.Printf("Uppercase: \"%s\"\n", strings.ToUpper(trimmed))

	// `strings.Repeat` creates a new string by repeating a given string a number of times.
	repeated := strings.Repeat("Go! ", 3)
	fmt.Printf("Repeated string: \"%s\"\n", repeated)
}

// demonstrateAdvancedTechniques covers runes, regex, and efficient building.
func demonstrateAdvancedTechniques() {
	fmt.Println("\n--- 6. Advanced Techniques: Runes, Regex, and Builder ---")

	// a) Counting Runes (Characters)
	// As mentioned, `len()` counts bytes. For multi-byte characters (like '안' or '녕'),
	// you need `utf8.RuneCountInString` to get the actual character count.
	koreanGreeting := "Hello, 안녕"
	fmt.Printf("\nString with multi-byte chars: \"%s\"\n", koreanGreeting)
	fmt.Printf("Length in bytes (len()): %d\n", len(koreanGreeting))
	fmt.Printf("Length in runes/characters (utf8.RuneCountInString()): %d\n", utf8.RuneCountInString(koreanGreeting))

	// b) Regular Expressions
	// The `regexp` package allows for powerful pattern matching.
	textWithNumbers := "Order items: 101, 205, and 303."
	// `regexp.MustCompile` compiles a regular expression. `\d+` matches one or more digits.
	// It will panic if the expression is invalid, which is fine for fixed patterns at compile time.
	digitRegex := regexp.MustCompile(`\d+`)
	// `FindAllString` finds all non-overlapping matches. The `-1` means find all.
	matches := digitRegex.FindAllString(textWithNumbers, -1)
	fmt.Printf("\nFound numbers in \"%s\": %v\n", textWithNumbers, matches)

	// c) Efficient String Building with `strings.Builder`
	// Repeatedly concatenating strings with `+` can be inefficient because it creates
	// a new string in memory for each operation. `strings.Builder` is much more efficient.
	fmt.Println("\nUsing strings.Builder for efficient concatenation:")
	var builder strings.Builder
	builder.WriteString("Building a string")
	builder.WriteString(" piece")
	builder.WriteString(" by piece.")
	builder.WriteRune('!') // Use WriteRune for single characters

	finalString := builder.String() // Get the final result.
	fmt.Println("Built string:", finalString)

	// You can reset the builder to reuse it without allocating new memory.
	builder.Reset()
	builder.WriteString("Starting fresh!")
	fmt.Println("Builder after reset:", builder.String())
}
