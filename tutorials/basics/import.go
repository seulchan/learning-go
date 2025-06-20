package main

// The `import` keyword is fundamental in Go for bringing in external code, known as packages.
// Packages are directories containing Go source files that provide specific functionalities.
// Think of them as libraries or modules in other languages.
// Every Go program starts execution in the `main` package, specifically the `main` function.

import (
	// === Standard Import ===
	// This is the most common way to import a package.
	// "fmt" is a core package from Go's standard library. It provides functions for formatted
	// input and output, like printing to the console.
	// After importing, you access its functions using the package name as a prefix, e.g., `fmt.Println()`.
	"fmt"

	// "math" is another standard library package, providing basic mathematical constants and functions.
	"math" // We'll use this for a simple math operation.

	// === Aliased Import ===
	// Sometimes, a package name might be too long, or it might conflict with another package name
	// or a local variable. In such cases, you can provide an alias.
	// "net/http" is a standard library package for making HTTP requests and building HTTP servers.
	// Here, we're aliasing "net/http" to `httpClient`.
	// So, instead of `http.Get()`, we'll use `httpClient.Get()`.
	httpClient "net/http"
	// === Dot Import (Use with Extreme Caution) ===
	// import . "strings"
	//
	// If you uncomment the line above, all exported names (functions, types, variables) from the "strings"
	// package would become available in the current file's scope directly, without needing the "strings." prefix.
	// For example, you could call `ToUpper()` instead of `strings.ToUpper()`.
	// WHY CAUTION? This can make code harder to read because it's not immediately clear where `ToUpper()`
	// comes from. It can also lead to name collisions if different packages export names that are identical.
	// It's generally discouraged in idiomatic Go.
	// === Blank Identifier Import (for Side Effects) ===
	// import _ "database/sql/driver" // Example for database driver registration
	// import _ "image/png"
	//
	// Sometimes, you need to import a package solely for its "side effects," without directly using any
	// of its exported code in your current file.
	// A common use case is when a package's `init()` function needs to run. For example:
	//   - Database drivers often register themselves with the `database/sql` package in their `init()` function.
	//   - Image packages (like "image/png" or "image/jpeg") register their respective image formats
	//     so the `image.Decode` function can understand them.
	// The blank identifier `_` tells the Go compiler that you are intentionally importing the package
	// but not using its name. This prevents an "unused import" error.
	// The package's `init()` function (if it has one) will be executed.
)

func main() {
	// Using the "fmt" package (standard import)
	fmt.Println("--- Go Import Statement Tutorial ---")
	fmt.Println("Exploring different ways to import packages.")

	// Using the "math" package (standard import)
	radius := 5.0
	area := math.Pi * math.Pow(radius, 2) // math.Pi and math.Pow come from the "math" package
	fmt.Printf("Area of a circle with radius %.1f: %.2f\n", radius, area)

	// Using the "net/http" package via its alias `httpClient`
	fmt.Println("\n--- Making an HTTP GET request using an aliased import ---")
	// We'll fetch some simple data from httpbin.org, a service for testing HTTP requests.
	apiURL := "https://catfact.ninja/fact"
	resp, err := httpClient.Get(apiURL) // Using the alias `httpClient`
	if err != nil {
		// `fmt.Errorf` is useful for formatting error messages.
		fmt.Printf("Error making HTTP request to %s: %v\n", apiURL, err)
		return // Exit main if there's an error
	}
	// `defer` ensures that `resp.Body.Close()` is called just before the `main` function exits.
	// This is crucial for releasing resources, especially network connections.
	defer resp.Body.Close()

	fmt.Printf("HTTP Response Status from %s: %s\n", apiURL, resp.Status)
	// For simplicity, we're not parsing the JSON body here, just showing the status.
	// You would typically read `resp.Body` to get the actual content.

	// --- Additional Important Notes on Imports ---
	fmt.Println("\n--- Key Takeaways about Imports ---")

	// 1. Package Paths and Sources:
	//    - Standard Library: Packages like "fmt", "math", "net/http", "os", "strings" are built into Go.
	//      You don't need to install them separately.
	//    - Third-Party Packages: You can use packages created by others, often hosted on platforms like GitHub.
	//      Example: `import "github.com/sirupsen/logrus"` (a popular logging library).
	//      To use third-party packages, you first need to download them using `go get <package-path>`.
	//      Go Modules (go.mod file) manage these dependencies in modern Go projects.

	// 2. Unused Imports are Errors:
	//    Go is strict about clean code. If you import a package but do not use any of its exported
	//    identifiers (functions, types, variables, constants) in your code, the Go compiler
	//    will report an "imported and not used" error, and your program will not compile.
	//    This helps keep your codebase tidy and dependencies explicit.
	//    The only exception is the blank identifier import (`_ "package/path"`), which explicitly
	//    signals that the import is for side effects only.

	// 3. Package `init` Functions:
	//    When a package is imported, if it contains one or more `init()` functions, those functions
	//    are executed automatically before the `main` function (or the importing package's code) runs.
	//    `init()` functions cannot be called directly. They are used for setup tasks within a package.
	//    This is particularly relevant for blank identifier imports, where the `init()` side effect is the primary goal.

	fmt.Println("\nEnd of import demonstration.")
}
