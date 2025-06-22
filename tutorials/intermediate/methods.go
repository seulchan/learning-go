// --- Go Methods Tutorial ---
//
// In Go, a method is a function that is associated with a particular type.
// Methods are a key part of Go's approach to object-oriented programming, allowing
// you to define behaviors for your custom data types (like structs).
//
// This tutorial covers:
// - Defining methods on structs.
// - The difference between "value receivers" and "pointer receivers".
// - Defining methods on non-struct types.
// - Method promotion through struct embedding.
package main

import "fmt"

// --- 1. Defining a Struct ---
// We'll start with a `Rectangle` struct. This will be the type we define methods on.
type Rectangle struct {
	length float64
	width  float64
}

// --- 2. Method with a Value Receiver ---
// `Area` is a method defined on the `Rectangle` type.
// The `(r Rectangle)` part is called the "receiver". It specifies which type this method belongs to.
//
// This is a VALUE receiver because `r` is a copy of the `Rectangle` instance the method is called on.
// This means `Area` can read the rectangle's data, but any changes it makes to `r`
// will NOT affect the original `Rectangle` variable. Value receivers are safe and efficient
// for methods that don't need to modify the original data.
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// --- 3. Method with a Pointer Receiver ---
// `Scale` is also a method on `Rectangle`, but it uses a POINTER receiver `(r *Rectangle)`.
// The `*` indicates that `r` is a pointer to the original `Rectangle` instance.
//
// Using a pointer receiver allows the method to MODIFY the original struct's data.
// This is necessary when a method needs to change the state of the object it's called on.
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

// --- 4. Methods on Non-Struct Types ---
// You can define methods on any type you declare in your package, not just structs.
// Here, we define a new type `MyInt` which is based on Go's built-in `int` type.
type MyInt int

// `IsPositive` is a method on our custom `MyInt` type.
func (m MyInt) IsPositive() bool {
	return m > 0
}

// The receiver name can be omitted if it's not used inside the method.
// This is done using the blank identifier `_`.
func (_ MyInt) Describe() string {
	return "This is a custom integer type named MyInt."
}

// --- 5. Method Promotion via Embedding ---
// `Figure` is a struct that "embeds" the `Rectangle` struct.
// Embedding is a form of composition that allows for "method promotion".
// This means the `Figure` type automatically gets all the methods of `Rectangle`.
type Figure struct {
	Rectangle // Embedded struct
	name      string
}

func main() {
	fmt.Println("--- Go Methods Tutorial ---")

	// --- Using Value and Pointer Receiver Methods ---
	fmt.Println("\n--- Part 1: Value vs. Pointer Receivers ---")
	// Create an instance of Rectangle.
	rect := Rectangle{length: 10, width: 5}
	fmt.Printf("Original Rectangle: %+v\n", rect)

	// Call the `Area` method (value receiver).
	area := rect.Area()
	fmt.Printf("Area of rectangle: %.2f\n", area)

	// Call the `Scale` method (pointer receiver).
	// This will modify the original `rect` variable.
	// Note: Go automatically converts `rect` to `&rect` when calling a pointer receiver method.
	// So, `rect.Scale(2)` is a convenient shorthand for `(&rect).Scale(2)`.
	rect.Scale(2)
	fmt.Printf("Rectangle after scaling by a factor of 2: %+v\n", rect)

	// Calculate the area again to see the result of the modification.
	newArea := rect.Area()
	fmt.Printf("New area after scaling: %.2f\n", newArea)

	// --- Using Methods on a Non-Struct Type ---
	fmt.Println("\n--- Part 2: Methods on Non-Struct Types ---")
	positiveNum := MyInt(10)
	negativeNum := MyInt(-5)

	fmt.Printf("Is %d positive? %t\n", positiveNum, positiveNum.IsPositive())
	fmt.Printf("Is %d positive? %t\n", negativeNum, negativeNum.IsPositive())
	fmt.Println(positiveNum.Describe()) // Calling the method with the blank identifier receiver.

	// --- Demonstrating Method Promotion ---
	fmt.Println("\n--- Part 3: Method Promotion via Embedding ---")
	fig := Figure{
		Rectangle: Rectangle{length: 8, width: 8},
		name:      "Square",
	}

	// Even though `fig` is a `Figure`, we can call the `Area` method directly on it.
	// This is "method promotion". The `Area` method from the embedded `Rectangle` is promoted to `Figure`.
	fmt.Printf("The figure named '%s' has an area of: %.2f\n", fig.name, fig.Area())

	// We can also call the pointer receiver method, which modifies the embedded Rectangle.
	fig.Scale(0.5)
	fmt.Printf("After scaling, the figure '%s' has a new area of: %.2f\n", fig.name, fig.Area())
	fmt.Printf("The figure's dimensions are now: %+v\n", fig.Rectangle)

	fmt.Println("\n--- End of Methods Tutorial ---")
}
