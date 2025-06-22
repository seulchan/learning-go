// --- Go Structs Tutorial ---
//
// A `struct` (short for structure) is a composite data type that groups together
// zero or more variables (called fields) under a single name. Structs are used to
// represent real-world entities or complex data records. For example, you could
// use a struct to represent an employee, a user profile, or a geometric shape.
//
// This tutorial covers how to define, initialize, and use structs, including
// nested structs, embedded structs, methods, and pointers.
package main

import (
	"fmt"
)

// Address is a struct that holds components of a physical address.
// This is an example of a "nested struct" when used inside another struct.
// Note: Field names like `City` and `Country` start with an uppercase letter,
// which makes them "exported". Exported fields can be accessed from other packages.
type Address struct {
	City    string
	Country string
}

// ContactInfo is a struct that holds contact details.
// This will be used as an "embedded struct".
type ContactInfo struct {
	HomePhone string
	CellPhone string
}

// Person is a struct representing a person's profile.
// It demonstrates how to group different data types together.
type Person struct {
	// --- Basic Fields ---
	FirstName string
	LastName  string
	Age       int

	// --- Nested Struct ---
	// The `HomeAddress` field is of type `Address`. This is called composition or nesting.
	// To access the city, you would use `person.HomeAddress.City`.
	HomeAddress Address

	// --- Embedded Struct ---
	// We can "embed" the `ContactInfo` struct directly by just specifying its type.
	// This promotes the fields of `ContactInfo` (`HomePhone`, `CellPhone`) to the `Person`
	// struct. This means you can access them directly, like `person.HomePhone`,
	// as if they were fields of `Person` itself. This is a form of composition, not inheritance.
	ContactInfo
}

// --- Methods ---
// Methods are functions that are associated with a specific type.

// FullName is a "value receiver" method for the Person struct.
// It receives a *copy* of the Person instance (`p`).
// This means it can read the data but cannot modify the original Person struct.
// It's efficient for methods that don't need to change state.
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// IncrementAge is a "pointer receiver" method for the Person struct.
// It receives a *pointer* to the Person instance (`*p`).
// This allows the method to modify the original struct's data.
// Use pointer receivers when you need to change the state of the struct.
func (p *Person) IncrementAge() {
	p.Age++ // This modifies the original Person's age.
}

func main() {
	fmt.Println("--- Go Structs Tutorial ---")

	// --- 1. Initializing a Struct ---
	// You can initialize a struct by providing values for its fields.
	// This is called a "struct literal".
	fmt.Println("\n--- 1. Initializing Structs ---")
	p1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		HomeAddress: Address{ // Initializing the nested struct
			City:    "New York",
			Country: "USA",
		},
		ContactInfo: ContactInfo{ // Initializing the embedded struct
			HomePhone: "123-456-7890",
			CellPhone: "987-654-3210",
		},
	}
	fmt.Printf("Initialized Person (p1): %+v\n", p1) // %+v prints field names, great for structs!

	// --- 2. Accessing Struct Fields ---
	fmt.Println("\n--- 2. Accessing Struct Fields ---")
	fmt.Println("First Name:", p1.FirstName)
	fmt.Println("Nested Field (City):", p1.HomeAddress.City)
	// Accessing fields from the embedded struct directly.
	fmt.Println("Embedded Field (Cell Phone):", p1.CellPhone)

	// --- 3. Zero Value of a Struct ---
	// If you declare a struct variable without initializing it, its fields
	// will be set to their respective "zero values" (0 for numbers, "" for strings, etc.).
	fmt.Println("\n--- 3. Zero Value of a Struct ---")
	var p2 Person
	fmt.Printf("Zero-valued Person (p2): %+v\n", p2)

	// --- 4. Calling Methods ---
	fmt.Println("\n--- 4. Calling Methods ---")
	fmt.Println("Full Name (from value receiver method):", p1.FullName())

	fmt.Println("Original Age:", p1.Age)
	p1.IncrementAge() // Calling the pointer receiver method to modify p1.
	fmt.Println("Age after IncrementAge():", p1.Age)

	// --- 5. Struct Pointers ---
	// It's common to work with pointers to structs to avoid copying large structs
	// and to allow functions/methods to modify them.
	fmt.Println("\n--- 5. Struct Pointers ---")
	p3 := &Person{FirstName: "Jane", LastName: "Smith", Age: 25}
	// Go automatically dereferences the pointer for you when accessing fields.
	// So, `p3.FirstName` is a convenient shorthand for `(*p3).FirstName`.
	fmt.Printf("Person from pointer (p3): Name: %s, Age: %d\n", p3.FullName(), p3.Age)
	p3.IncrementAge()
	fmt.Printf("Age of p3 after increment: %d\n", p3.Age)

	// --- 6. Struct Comparison ---
	// Structs are comparable if all their fields are comparable.
	// They are equal if their corresponding fields are equal.
	fmt.Println("\n--- 6. Struct Comparison ---")
	personA := Person{FirstName: "Alex", Age: 42}
	personB := Person{FirstName: "Alex", Age: 42}
	personC := Person{FirstName: "Maria", Age: 35}

	fmt.Println("Is personA == personB?", personA == personB) // true
	fmt.Println("Is personA == personC?", personA == personC) // false

	// --- 7. Anonymous Structs ---
	// Sometimes you need a struct for a short-lived, temporary purpose without
	// giving it a formal type name. This is an "anonymous struct".
	fmt.Println("\n--- 7. Anonymous Structs ---")
	// The struct is defined and instantiated in the same statement.
	tempUser := struct {
		Username string
		IsActive bool
	}{
		Username: "temp_user_123",
		IsActive: true,
	}
	fmt.Printf("Anonymous struct user: %+v\n", tempUser)
	fmt.Println("Username from anonymous struct:", tempUser.Username)

	fmt.Println("\n--- End of Structs Tutorial ---")
}
