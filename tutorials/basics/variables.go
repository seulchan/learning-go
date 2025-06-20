// In Go, a variable is a named storage location that holds a value of a specific type.
// This file demonstrates various ways to declare, initialize, and use variables,
// along with explanations of their scope and Go's zero values.
package main

import "fmt"

// --- Package Level Variables & Constants ---
// Variables declared outside of any function have "package scope".
// They are visible and accessible to all functions within the same package.

// `defaultUserName` is declared at the package level using `var` with an initializer.
// Go infers its type as `string` because it's initialized with a string literal.
var defaultUserName = "GuestUser"

// Constants are declared using the `const` keyword.
// Their values are fixed at compile time and cannot be changed during program execution.
const AppVersion = "1.0.1"

// You can declare multiple package-level variables or constants in a block for better organization.
const (
	MaxLoginAttempts = 3
	DefaultPort      = 8080
)

var (
	// `IsProduction` is a package-level boolean variable, initialized to its zero value (false).
	IsProduction bool
)

// `demonstrateLocalScope` shows variables local to this function.
func demonstrateLocalScope() {
	// `sessionID` is declared using the short variable declaration `:=`.
	// This form is only available inside functions.
	// It infers the type as `string` and assigns the value.
	// `sessionID` is local to the `demonstrateLocalScope` function and cannot be accessed outside it.
	sessionID := "xyz123abc"
	fmt.Println("Inside demonstrateLocalScope:")
	fmt.Println("  Local sessionID:", sessionID)

	// We can access package-level variables like `defaultUserName` here.
	fmt.Println("  Accessing package-level defaultUserName:", defaultUserName)
	fmt.Println("  Accessing package-level constant AppVersion:", AppVersion)
}

func main() {
	fmt.Println("--- Go Variable Declaration & Usage Tutorial ---")

	// --- Function Level Variables (Local Scope) ---
	// Variables declared inside a function are local to that function.

	// 1. Declaration with explicit type, without initialization.
	// `userAge` is declared as an integer (`int`).
	// It will be automatically initialized to its "zero value", which is 0 for numeric types.
	var age int
	fmt.Println("Initial age (zero value for int):", age)
	age = 35 // Assigning a value
	fmt.Println("Assigned age:", age)

	// 2. Declaration with explicit type and initialization.
	// `userName` is declared as a string and initialized to "Alice".
	var userName string = "Alice"
	fmt.Println("User Name (explicit type):", userName)

	// 3. Declaration with type inference (Go figures out the type).
	// `userCity` is initialized with "London", so Go infers its type as `string`.
	var userCity = "London"
	fmt.Println("User City (type inferred):", userCity)

	// 4. Short variable declaration `:=` (most common inside functions).
	// This declares and initializes `itemCount` to 5, inferring its type as `int`.
	// This syntax can ONLY be used inside functions.
	// At least one variable on the left side of `:=` must be newly declared.
	itemCount := 5
	fmt.Println("Item Count (short declaration):", itemCount)

	// You can re-assign to a variable declared with `:=` using `=`, as long as it's in the same scope.
	itemCount = itemCount + 2
	fmt.Println("Updated Item Count:", itemCount)

	// 5. Multiple variable declarations in a single statement.
	// `country` and `postalCode` are both inferred as strings.
	var country, postalCode = "Canada", "K1A 0B1"
	fmt.Printf("Country: %s, Postal Code: %s\n", country, postalCode)

	// Multiple declarations using a block (often used for related variables).
	var (
		isActive   bool    // Zero value for bool is `false`
		userScore  float64 // Zero value for float64 is `0.0`
		userEmail  string  // Zero value for string is "" (empty string)
		userPoints *int    // Zero value for a pointer is `nil`
	)
	fmt.Printf("Initial isActive: %t, userScore: %.1f, userEmail: '%s', userPoints: %v\n",
		isActive, userScore, userEmail, userPoints)

	isActive = true
	userScore = 88.7
	userEmail = "test@example.com"
	fmt.Printf("Updated isActive: %t, userScore: %.1f, userEmail: '%s'\n",
		isActive, userScore, userEmail)

	// --- Zero Values Explained ---
	// When a variable is declared without an explicit initial value, Go assigns it a "zero value".
	// - Numeric types (int, float32, float64, etc.): 0
	// - Boolean type (`bool`): false
	// - String type (`string`): "" (an empty string)
	// - Pointers, functions, interfaces, slices, channels, and maps: nil

	// --- Scope Demonstration ---
	fmt.Println("\n--- Scope Demonstration ---")
	// Accessing package-level variables and constants from `main`:
	fmt.Println("Accessing package-level defaultUserName from main:", defaultUserName)
	fmt.Println("Accessing package-level constant MaxLoginAttempts from main:", MaxLoginAttempts)
	IsProduction = true // Modifying a package-level variable
	fmt.Println("IsProduction (modified in main):", IsProduction)

	demonstrateLocalScope() // Call another function to see its local scope.
	// Note: `sessionID` from `demonstrateLocalScope` is NOT accessible here in `main`.

	// Block Scope: Variables can also be scoped to blocks (e.g., within an `if` or `for` statement).
	if true {
		blockScopedMessage := "I'm only visible inside this if-block!"
		fmt.Println(blockScopedMessage)
		// `userName` (from main's scope) is accessible here.
		fmt.Println("  Accessing `userName` (from main's scope) inside block:", userName)
	}
	// fmt.Println(blockScopedMessage) // This would cause a compile error: undefined: blockScopedMessage

	fmt.Println("\nEnd of variable.")
}
