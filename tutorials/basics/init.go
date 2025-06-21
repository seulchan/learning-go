package main

import "fmt"

// --- Go Package Initialization (init function) Tutorial - Single File Example ---
//
// In Go, the `init` function is a special function that is automatically
// executed when a package is initialized. It's run before the `main` function
// (if the package is `main`) or before any exported functions are called from
// an imported package.
//
// Key points about `init` functions:
// 1. A single Go source file can have AT MOST ONE `init` function.
//    The original version of this file attempted to define `init()` twice,
//    which causes a compile-time error: `init redeclared in this block`.
//    We have corrected this by having only one `init` function below.
//
// 2. A package (which can consist of multiple files) CAN have multiple `init`
//    functions, but each `init` function must reside in a *different file*
//    within that package. (See the `main.go` and `another_init.go` examples
//    for how multiple `init` functions work across different files in the same package).
//
// 3. All `init` functions in a package are executed *before* the `main` function
//    (if it's the `main` package) or before the package is used by another.
//
// 4. `init` functions are typically used for:
//    - Setting up package-specific state (e.g., initializing global variables).
//    - Registering components with a central registry (e.g., database drivers, image formats).
//    - Performing one-time computations or validations needed before the package is used.
//
// 5. `init` functions cannot be called directly from your code. They take no
//    arguments and return no values.
//
// 6. If a package imports other packages, the `init` functions of the imported
//    packages are guaranteed to run before the `init` function of the importing package.

// This `init` function will be executed automatically when the `main` package
// (which this file is part of) is initialized. This happens right before
// the `main` function runs.
func init() {
	fmt.Println("Executing the init function in init.go...")
	// You can perform various setup tasks here.
	// For example, initializing a package-level variable:
	// globalConfig = loadConfigurationFromFile("config.json")
	fmt.Println("Package-level setup in init.go is complete.")
}

// The `main` function is the special entry point for an executable Go program.
// Execution of your program's logic begins here, but only *after* all `init`
// functions in the `main` package (and any imported packages) have finished.
func main() {
	fmt.Println("Executing the main function.")
	// Your program's primary logic would start here.
}
