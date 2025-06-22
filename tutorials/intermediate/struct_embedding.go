// --- Go Struct Embedding Tutorial ---
//
// Struct embedding is a powerful feature in Go that allows you to compose structs
// together. It's Go's way of achieving code reuse and is often compared to
// inheritance in other languages, though it's fundamentally different.
//
// When you "embed" a struct, its fields and methods are "promoted" to the
// outer struct, meaning you can access them directly as if they were part of
// the outer struct itself.
//
// This tutorial demonstrates:
// 1. How to define and embed a struct.
// 2. How fields and methods are promoted.
// 3. How to override an embedded method.
// 4. The difference between embedding and standard composition.
package main

import "fmt"

// --- 1. Defining the Base Struct to be Embedded ---
// `BasicInfo` is a struct containing common fields that could be shared
// across different types like `Employee`, `Customer`, etc.
type BasicInfo struct {
	Name string
	Age  int
}

// `Greet` is a method on the `BasicInfo` struct.
// When `BasicInfo` is embedded, this method will be promoted.
func (bi BasicInfo) Greet() {
	fmt.Printf("Hi, my name is %s and I am %d years old.\n", bi.Name, bi.Age)
}

// --- 2. The Outer Struct with Embedding ---
// `Employee` is a struct that represents an employee.
// Instead of redefining Name and Age, we embed `BasicInfo`.
type Employee struct {
	BasicInfo  // Anonymous field: This is struct embedding.
	EmployeeID string
	Salary     float64
}

// --- 3. Method Overriding ---
// The `Employee` struct also has a `Greet` method. This "overrides" the
// `Greet` method from the embedded `BasicInfo` struct.
// When you call `emp.Greet()`, this version will be executed.
func (e Employee) Greet() {
	// We can still access the embedded struct's method explicitly if needed.
	// This is useful for extending behavior rather than completely replacing it.
	fmt.Print("As an employee, I'd like to say: ")
	e.BasicInfo.Greet()
}

// --- 4. Contrast: Struct with Composition (Not Embedding) ---
// `Manager` is a struct that uses standard composition.
// It has a *named field* `Info` of type `BasicInfo`.
type Manager struct {
	Info     BasicInfo // Named field: This is composition.
	TeamSize int
}

func main() {
	fmt.Println("--- Go Struct Embedding Tutorial ---")

	// --- Creating and Using an Embedded Struct ---
	fmt.Println("\n--- 1. Demonstrating Struct Embedding ---")

	// Initialize an Employee instance.
	// Note that we initialize the embedded `BasicInfo` struct as a field.
	emp := Employee{
		BasicInfo:  BasicInfo{Name: "Alice", Age: 30},
		EmployeeID: "E42",
		Salary:     75000,
	}

	// --- Accessing Promoted Fields ---
	// Because `BasicInfo` is embedded, we can access its fields directly on `emp`.
	// `emp.Name` is a shortcut for `emp.BasicInfo.Name`.
	fmt.Printf("Employee Name (promoted field): %s\n", emp.Name)
	fmt.Printf("Employee Age (promoted field): %d\n", emp.Age)
	fmt.Println("Employee ID:", emp.EmployeeID)

	// --- Calling Promoted and Overridden Methods ---
	fmt.Println("\nCalling methods:")
	// When we call `Greet`, we get the `Employee`'s version (the override).
	emp.Greet()

	// To call the original method from the embedded struct, we must access it explicitly.
	fmt.Print("Calling the embedded struct's method directly: ")
	emp.BasicInfo.Greet()

	// --- Comparison with Standard Composition ---
	fmt.Println("\n--- 2. Demonstrating Standard Composition (for comparison) ---")
	mgr := Manager{
		Info:     BasicInfo{Name: "Bob", Age: 45},
		TeamSize: 10,
	}

	// With composition, there are no promoted fields. You MUST go through the named field.
	// `mgr.Name` would cause a compile error. You must use `mgr.Info.Name`.
	fmt.Printf("Manager Name (via named field): %s\n", mgr.Info.Name)

	// Similarly, methods are not promoted. You must call them on the named field.
	// `mgr.Greet()` would cause a compile error.
	fmt.Print("Calling the manager's info method: ")
	mgr.Info.Greet()

	fmt.Println("\n--- End of Struct Embedding Tutorial ---")
}
