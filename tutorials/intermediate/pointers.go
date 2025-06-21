// Package main provides a beginner-friendly tutorial on pointers in Go.
// Pointers are a fundamental concept in many programming languages, including Go.
// They hold the memory address of a variable, allowing for more efficient and
// flexible ways to handle data.
package main

import "fmt"

// --- Function to Modify Value Directly (Pass-by-Value) ---
// This function takes an integer `val` as an argument.
// In Go, arguments are passed by value, which means the function receives a *copy*
// of the original value. Any modifications inside this function will not affect
// the original variable outside of it.
func incrementValueDirectly(val int) {
	val++ // Increment the copy of the value
	fmt.Printf("Inside incrementValueDirectly, the value is now: %d\n", val)
}

// --- Function to Modify Value via a Pointer (Simulating Pass-by-Reference) ---
// This function takes a pointer to an integer `ptr *int` as an argument.
// By passing a pointer, we are giving the function the memory address of the
// original variable. This allows the function to directly access and modify
// the original value.
func incrementValueViaPointer(ptr *int) {
	// The `*` operator here is used for "dereferencing". It means "go to the
	// memory address that ptr is holding and get the value stored there."
	// We then increment that value.
	*ptr++
	fmt.Printf("Inside incrementValueViaPointer, the value at address %v is now: %d\n", ptr, *ptr)
}

func main() {
	fmt.Println("--- Go Pointers Tutorial ---")

	// --- 1. Declaring a variable ---
	// We start with a simple integer variable.
	number := 42
	fmt.Printf("1. Original variable 'number':\n   Value: %d\n   Memory Address: %v\n\n", number, &number)

	// --- 2. Declaring a pointer ---
	// A pointer is a variable that stores the memory address of another variable.
	// `var pointerToNumber *int` declares a pointer named `pointerToNumber` that can
	// hold the address of an integer (`int`).
	// Initially, it's a "nil pointer" because it doesn't point to anything yet.
	var pointerToNumber *int
	fmt.Printf("2. Declared pointer 'pointerToNumber':\n   Value (address it holds): %v\n", pointerToNumber)

	// A nil pointer has a value of `nil`.
	if pointerToNumber == nil {
		fmt.Println("   Status: The pointer is currently nil (it points to nothing).\n")
	}

	// --- 3. Assigning an address to a pointer ---
	// The `&` operator (the "address-of" operator) gives us the memory address of a variable.
	// We assign the memory address of `number` to our pointer.
	pointerToNumber = &number
	fmt.Printf("3. Pointer after assignment (`pointerToNumber = &number`):\n")
	fmt.Printf("   Value (address it holds): %v\n", pointerToNumber)
	fmt.Printf("   It now points to the address of 'number'.\n\n")

	// --- 4. Dereferencing a pointer ---
	// The `*` operator (the "dereferencing" operator) allows us to see the value
	// stored at the memory address the pointer is holding.
	// `*pointerToNumber` means "get the value that `pointerToNumber` points to".
	fmt.Printf("4. Dereferencing the pointer (`*pointerToNumber`):\n")
	fmt.Printf("   Value at the address the pointer holds: %d\n\n", *pointerToNumber)

	// --- 5. Modifying a value through a pointer ---
	// We can also use the dereferencing operator to change the original value.
	fmt.Println("5. Modifying the original value through the pointer:")
	fmt.Printf("   Original 'number' value before modification: %d\n", number)
	*pointerToNumber = 100 // Go to the address and change the value to 100
	fmt.Printf("   Value of 'number' after `*pointerToNumber = 100`: %d\n\n", number)

	// --- 6. Pointers in Functions: Pass-by-Value vs. Pass-by-Reference Simulation ---
	fmt.Println("6. Demonstrating pointers with functions:")

	// a. Trying to modify with a pass-by-value function
	currentValue := 10
	fmt.Printf("   a) Before calling incrementValueDirectly, currentValue is: %d\n", currentValue)
	incrementValueDirectly(currentValue)
	fmt.Printf("   a) After calling incrementValueDirectly, currentValue is: %d (Unchanged!)\n\n", currentValue)

	// b. Modifying with a function that accepts a pointer
	fmt.Printf("   b) Before calling incrementValueViaPointer, currentValue is: %d\n", currentValue)
	incrementValueViaPointer(&currentValue) // We pass the memory address
	fmt.Printf("   b) After calling incrementValueViaPointer, currentValue is: %d (Changed!)\n\n", currentValue)

	fmt.Println("--- End of Pointers Tutorial ---")
}
