package main

import (
	"fmt"
	"unicode/utf8"
)

// --- Go Strings and Runes Tutorial ---
//
// This tutorial explores the fundamental concepts of strings and runes in Go.
// Understanding how Go handles text is crucial for any Go developer.

func main() {
	fmt.Println("--- Go Strings and Runes Tutorial ---")

	// --- 1. String Literals ---
	// Go has two types of string literals: interpreted and raw.

	// a) Interpreted String Literals (using double quotes "")
	// These strings allow for "escape sequences" like `\n` (newline) and `\t` (tab).
	fmt.Println("\n--- 1. Interpreted vs. Raw String Literals ---")
	interpretedString := "Hello,\n\tGo developers!"
	fmt.Println("Interpreted String:")
	fmt.Println(interpretedString)

	// b) Raw String Literals (using backticks ``)
	// These strings treat all characters literally. Escape sequences are not processed.
	// They are useful for multi-line strings or strings containing backslashes (e.g., file paths, regex).
	rawString := `Hello,\n\tGo developers!`
	fmt.Println("\nRaw String (backslashes are treated as literal characters):")
	fmt.Println(rawString)

	// --- 2. Strings are Slices of Bytes ---
	// In Go, a string is an immutable slice of bytes. This has important implications.
	fmt.Println("\n--- 2. Strings are Slices of Bytes (and immutable) ---")

	// Let's use a string with a multi-byte character. 'é' takes 2 bytes in UTF-8.
	cafe := "café"

	// `len()` returns the number of bytes, NOT the number of characters (runes).
	fmt.Printf("String: \"%s\"\n", cafe)
	fmt.Printf("Length in bytes (len()): %d\n", len(cafe)) // Expected: 5 (c=1, a=1, f=1, é=2)

	// `utf8.RuneCountInString()` returns the number of characters (runes).
	fmt.Printf("Length in characters (RuneCountInString): %d\n", utf8.RuneCountInString(cafe)) // Expected: 4

	// Accessing a string by index gives you the byte at that position, not the character.
	fmt.Printf("Byte at index 3: %v (This is the first byte of 'é')\n", cafe[3])
	// fmt.Printf("Character at index 3: %c\n", cafe[3]) // This would print a garbled character.

	// Strings are immutable. You cannot change a character in a string directly.
	// The following line would cause a compile-time error:
	// cafe[0] = 'C' // Error: cannot assign to cafe[0]

	// --- 3. String Concatenation and Comparison ---
	fmt.Println("\n--- 3. String Concatenation and Comparison ---")
	greeting := "Hello, "
	name := "Alice"
	welcomeMessage := greeting + name // The `+` operator concatenates strings.
	fmt.Println("Concatenated string:", welcomeMessage)

	// Strings can be compared lexicographically (like in a dictionary).
	// Comparison is case-sensitive. 'A' comes before 'a'.
	fmt.Println("'Apple' < 'Banana':", "Apple" < "Banana") // true
	fmt.Println("'apple' > 'Apple':", "apple" > "Apple")   // true
	fmt.Println("'app' < 'apple':", "app" < "apple")       // true

	// --- 4. Runes: The Go Way of Handling Characters ---
	// A `rune` is Go's type for a single character. It's an alias for `int32`.
	// It can represent any Unicode code point.
	fmt.Println("\n--- 4. Understanding Runes ---")

	var englishRune rune = 'A'
	var koreanRune rune = '헬' // A Korean character
	var emojiRune rune = '😂'  // An emoji

	fmt.Printf("Rune: %c, Unicode Code Point: %U, Type: %T\n", englishRune, englishRune, englishRune)
	fmt.Printf("Rune: %c, Unicode Code Point: %U, Type: %T\n", koreanRune, koreanRune, koreanRune)
	fmt.Printf("Rune: %c, Unicode Code Point: %U, Type: %T\n", emojiRune, emojiRune, emojiRune)

	// You can convert a rune to a string.
	runeAsString := string(koreanRune)
	fmt.Printf("The rune '%c' converted to a string is \"%s\" of type %T\n", koreanRune, runeAsString, runeAsString)

	// --- 5. Iterating Over Strings with `for...range` ---
	// The `for...range` loop is the idiomatic way to iterate over characters in a string.
	// It decodes one UTF-8 encoded rune on each iteration.
	fmt.Println("\n--- 5. Iterating Over a String with for...range ---")
	multiByteString := "Go, 헬로우, 😂"
	fmt.Printf("Iterating over: \"%s\"\n", multiByteString)

	for index, r := range multiByteString {
		// `index` is the starting byte position of the rune.
		// `r` is the rune itself (an int32 value).
		fmt.Printf("Byte Index: %2d, Rune: %c, Unicode: %U\n", index, r, r)
	}
	// Notice how the byte index jumps by more than 1 for multi-byte characters.
	// '헬' is 3 bytes, '😂' is 4 bytes.

	fmt.Println("\nEnd of Strings and Runes Tutorial.")
}
