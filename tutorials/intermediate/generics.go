// --- Go Generics Tutorial ---
//
// Generics, introduced in Go 1.18, allow you to write code that works with
// multiple different types, without sacrificing type safety. Before generics,
// you would often need to write duplicate code for different types or resort
// to using the empty interface (`interface{}`), which loses type information.
//
// This tutorial demonstrates:
// 1. How to write a generic function.
// 2. How to define and use a generic data structure (a Stack).
// 3. How to use type constraints to limit the types a generic function can accept.

package main

import (
	"fmt"
	"strings"
)

// --- 1. A Simple Generic Function ---

// `Swap` is a generic function that can swap two values of any type.
//
// The syntax `[T any]` is the "type parameter list".
//   - `T` is the "type parameter". It's a placeholder for a real type (like int, string, etc.).
//   - `any` is the "type constraint". `any` is an alias for `interface{}`, meaning
//     `T` can be any type.
func Swap[T any](a, b T) (T, T) {
	return b, a
}

// --- 2. A Generic Data Structure ---

// `Stack[T]` is a generic type. It represents a last-in, first-out (LIFO) stack
// that can hold elements of any single type `T`.
// By defining `Stack` with a type parameter `T`, we can create a `Stack[int]`,
// a `Stack[string]`, or a stack of any other type, all using the same code.
type Stack[T any] struct {
	elements []T
}

// `Push` is a method on our generic `Stack`. It adds an element to the top of the stack.
// The receiver is `*Stack[T]`, indicating that the method will modify the stack.
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// `Pop` removes and returns the top element of the stack.
// It returns two values:
// 1. The popped element of type `T`.
// 2. A boolean `ok` which is `true` if an element was popped, and `false` if the stack was empty.
// This "comma, ok" idiom is very common in Go.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		// If the stack is empty, we can't return a real element.
		// We return the "zero value" for the type `T`.
		// `var zero T` creates a variable `zero` of type `T` initialized to its zero value
		// (e.g., 0 for int, "" for string, nil for pointers).
		var zero T
		return zero, false
	}
	// Get the index of the last element.
	lastIndex := len(s.elements) - 1
	// Get the element itself.
	element := s.elements[lastIndex]
	// Slice the underlying slice to remove the last element.
	s.elements = s.elements[:lastIndex]
	return element, true
}

// `IsEmpty` checks if the stack contains any elements.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// `String` makes our Stack implement the `fmt.Stringer` interface.
// This allows us to print the stack's contents in a clean, readable format
// just by using `fmt.Println(myStack)`.
func (s *Stack[T]) String() string {
	if s.IsEmpty() {
		return "Stack: []"
	}
	// This part is a bit more advanced but shows a nice way to format output.
	// It converts each element to a string and joins them with spaces.
	stringElements := make([]string, len(s.elements))
	for i, v := range s.elements {
		stringElements[i] = fmt.Sprint(v)
	}
	return fmt.Sprintf("Stack: [%s]", strings.Join(stringElements, " "))
}

// --- 3. Generic Functions with Constraints ---

// Sometimes, `any` is too broad. We need to guarantee that a type `T` supports
// certain operations (like `+` for addition). We do this with a custom constraint.
//
// `Number` is an interface used as a type constraint. The `|` symbol creates a
// "union" of types. This constraint says that any type that is an `int`, `int64`,
// `float32`, or `float64` satisfies the `Number` constraint.
type Number interface {
	int | int64 | float32 | float64
}

// `SumNumbers` is a generic function that can sum a slice of any type that
// satisfies the `Number` constraint.
// We can use the `+` operator on `T` because our constraint guarantees it's a numeric type.
func SumNumbers[T Number](numbers []T) T {
	var sum T // `sum` will be the zero value of T (e.g., 0 or 0.0)
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println("--- Go Generics Tutorial ---")

	// --- Using a generic function ---
	fmt.Println("\n--- 1. Generic Swap Function ---")
	// Swap integers
	a, b := 10, 20
	fmt.Printf("Before swap (int): a = %d, b = %d\n", a, b)
	a, b = Swap(a, b) // The compiler infers that T is int
	fmt.Printf("After swap (int):  a = %d, b = %d\n", a, b)

	// Swap strings
	str1, str2 := "Hello", "World"
	fmt.Printf("\nBefore swap (string): str1 = %s, str2 = %s\n", str1, str2)
	str1, str2 = Swap(str1, str2) // The compiler infers that T is string
	fmt.Printf("After swap (string):  str1 = %s, str2 = %s\n", str1, str2)

	// --- Using a generic type ---
	fmt.Println("\n--- 2. Generic Stack Data Structure ---")
	// Create and use a stack of integers
	fmt.Println("\n* Integer Stack *")
	intStack := &Stack[int]{}
	fmt.Println("Is empty?", intStack.IsEmpty())
	intStack.Push(100)
	intStack.Push(200)
	intStack.Push(300)
	fmt.Println(intStack) // Uses the String() method we defined
	fmt.Println("Is empty?", intStack.IsEmpty())

	if val, ok := intStack.Pop(); ok {
		fmt.Printf("Popped: %d\n", val)
	}
	fmt.Println(intStack)

	// Create and use a stack of strings
	fmt.Println("\n* String Stack *")
	stringStack := &Stack[string]{}
	stringStack.Push("Go")
	stringStack.Push("is")
	stringStack.Push("fun!")
	fmt.Println(stringStack)

	if val, ok := stringStack.Pop(); ok {
		fmt.Printf("Popped: %s\n", val)
	}
	if val, ok := stringStack.Pop(); ok {
		fmt.Printf("Popped: %s\n", val)
	}
	if val, ok := stringStack.Pop(); ok {
		fmt.Printf("Popped: %s\n", val)
	}
	// Try to pop from an empty stack
	if _, ok := stringStack.Pop(); !ok {
		fmt.Println("Attempted to pop from an empty stack, as expected.")
	}
	fmt.Println(stringStack)

	// --- Using a generic function with constraints ---
	fmt.Println("\n--- 3. Generic Function with a Number Constraint ---")
	intSlice := []int{1, 2, 3, 4, 5}
	floatSlice := []float64{1.1, 2.2, 3.3}

	intSum := SumNumbers(intSlice)
	floatSum := SumNumbers(floatSlice)

	fmt.Printf("Sum of integers %v: %d\n", intSlice, intSum)
	fmt.Printf("Sum of floats %v: %.2f\n", floatSlice, floatSum)

	// The following line would cause a compile error because `string` does not
	// satisfy the `Number` constraint. This is the power of type-safe generics!
	// stringSum := SumNumbers([]string{"a", "b"})
}
