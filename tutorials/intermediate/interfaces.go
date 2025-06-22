// --- Go Interfaces Tutorial ---
//
// An interface in Go is a type that defines a set of method signatures. It acts as a
// "contract". Any type that implements all the methods listed in the interface
// is said to "satisfy" that interface. This allows for powerful, flexible, and
// decoupled code, a key feature of Go's design.
//
// This tutorial covers:
// 1. Defining an interface.
// 2. Implicitly implementing an interface with structs.
// 3. Using interfaces to write generic functions (polymorphism).
// 4. The special empty interface `interface{}`.
// 5. Using type switches to work with interface values.
package main

import (
	"fmt"
	"math"
)

// --- 1. Defining an Interface ---
// We define an interface named `Shape`. By convention, interface names in Go
// often end with "er" (like `Reader`, `Writer`), but naming them after the
// concept they represent (like `Shape`) is also very common and clear.
//
// This interface specifies that any type wanting to be considered a `Shape`
// MUST have both an `Area()` method and a `Perimeter()` method.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// --- 2. Implementing the Interface ---
// Now, let's create some concrete types that will implement our `Shape` interface.

// Rectangle is a struct representing a rectangle.
type Rectangle struct {
	width, height float64
}

// Circle is a struct representing a circle.
type Circle struct {
	radius float64
}

// In Go, interface implementation is implicit. There's no `implements` keyword.
// If a type has all the methods required by an interface, it automatically
// satisfies that interface.

// The `Area` method for `Rectangle`. Since `Rectangle` now has `Area()` and `Perimeter()`
// (see below), it satisfies the `Shape` interface.
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// The `Perimeter` method for `Rectangle`.
func (r Rectangle) Perimeter() float64 {
	return 2*r.width + 2*r.height
}

// The `Area` method for `Circle`.
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// The `Perimeter` method for `Circle`. Now `Circle` also satisfies the `Shape` interface.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// A type can have more methods than the interface requires.
// `Diameter` is a method specific to `Circle` and is not part of the `Shape` interface.
func (c Circle) Diameter() float64 {
	return 2 * c.radius
}

// --- 3. Using the Interface for Polymorphism ---

// `printShapeDetails` is a function that can accept any type that satisfies the `Shape`
// interface. This is polymorphism in action. The function doesn't need to know
// whether it's receiving a `Rectangle` or a `Circle`; it only cares that the value
// it receives has `Area()` and `Perimeter()` methods.
func printShapeDetails(s Shape) {
	fmt.Printf("--- Details for shape [%T] ---\n", s)
	fmt.Printf("Value: %+v\n", s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())

	// We can't call `s.Diameter()` directly because the `Shape` interface
	// contract doesn't include a `Diameter` method.
	// To access type-specific methods, we use a "type assertion".
	// We check if the shape `s` is, in fact, a `Circle`.
	if c, ok := s.(Circle); ok {
		// The `ok` boolean is true if the assertion succeeds.
		// `c` will be the underlying `Circle` value.
		fmt.Printf("This is a Circle! Its diameter is: %.2f\n", c.Diameter())
	}
}

// --- 4. The Empty Interface: `interface{}` ---

// The empty interface has zero methods. Since every type has at least zero methods,
// every type in Go satisfies the empty interface. This means a variable of type
// `interface{}` can hold a value of ANY type. It's Go's way of handling dynamic types.
func printAnything(values ...interface{}) {
	fmt.Println("\n--- Printing various types using an empty interface ---")
	for _, value := range values {
		fmt.Printf("Value: %-10v \t Type: %T\n", value, value)
	}
}

// --- 5. Type Switches ---

// A "type switch" is a construct that lets you perform several type assertions in a series.
// It's a clean and idiomatic way to handle a value of an unknown type (like `interface{}`).
func inspectType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("The integer is %d\n", v)
	case string:
		fmt.Printf("The string is \"%s\" and its length is %d\n", v, len(v))
	case bool:
		fmt.Printf("The boolean is %t\n", v)
	default:
		// The `v` variable here will have the same type and value as `i`.
		fmt.Printf("Unhandled type! Value: %+v, Type: %T\n", v, v)
	}
}

func main() {
	// Create instances of our structs.
	rectangle := Rectangle{width: 10, height: 5}
	circle := Circle{radius: 7}

	// Pass both `rectangle` and `circle` to our generic function.
	printShapeDetails(rectangle)
	printShapeDetails(circle)

	// Demonstrate the empty interface's ability to hold any type.
	printAnything(42, "hello", true, rectangle, 3.14)

	// Demonstrate the type switch.
	fmt.Println("\n--- Inspecting types with a type switch ---")
	inspectType(99)
	inspectType("Go is fun!")
	inspectType(false)
	inspectType(circle) // Pass a struct to the default case

	fmt.Println("\n--- End of Interfaces Tutorial ---")
}
