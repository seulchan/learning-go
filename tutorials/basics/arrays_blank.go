// In Go, an array is a numbered sequence of elements of a single type with a fixed length.
// This file demonstrates how to declare, initialize, access, and iterate over arrays,
// highlighting key concepts like zero values and arrays being value types.
package main

import "fmt"

func main() {
	fmt.Println("--- Go Arrays Tutorial ---")

	// --- 1. Array Declaration and Initialization ---

	// Declare an array of 5 integers.
	// When declared without explicit initialization, elements are set to their "zero value".
	// The zero value for integers is 0.
	var integerArray [5]int
	fmt.Println("\n1a. Declared array (zero values):", integerArray) // Output: [0 0 0 0 0]

	// Assign values to specific elements using their index (position).
	// Array indices start from 0.
	integerArray[4] = 20                                         // Assign 20 to the 5th element (index 4)
	fmt.Println("1b. After assigning to index 4:", integerArray) // Output: [0 0 0 0 20]

	integerArray[0] = 9                                          // Assign 9 to the 1st element (index 0)
	fmt.Println("1c. After assigning to index 0:", integerArray) // Output: [9 0 0 0 20]

	// Declare and initialize an array using an array literal.
	// Go counts the elements to determine the fixed length (4 in this case).
	stringArray := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("\n1d. Initialized string array:", stringArray) // Output: [Apple Banana Orange Grapes]

	// Declare and initialize with inferred length using `...`.
	// Go determines the length based on the number of elements provided.
	inferredLengthArray := [...]float64{1.1, 2.2, 3.3}
	fmt.Printf("\n1e. Array with inferred length (%d): %v\n", len(inferredLengthArray), inferredLengthArray)

	// --- 2. Accessing Array Elements ---

	// Access elements using their index within square brackets `[]`.
	// Remember indices are 0-based.
	fmt.Println("\n2a. Accessing the third element (index 2) of stringArray:", stringArray[2]) // Output: Orange

	// The length of an array is fixed and can be obtained using the `len()` function.
	fmt.Printf("2b. Length of integerArray: %d\n", len(integerArray)) // Output: 5
	fmt.Printf("2c. Length of stringArray: %d\n", len(stringArray))   // Output: 4

	// --- 3. Iterating Over Arrays ---

	// 3a. Using a standard `for` loop with index.
	fmt.Println("\n3a. Iterating using a standard for loop:")
	for i := 0; i < len(integerArray); i++ {
		// Access element using the index `i`.
		fmt.Printf("  Element at index %d: %d\n", i, integerArray[i])
	}

	// 3b. Using the `for...range` loop (most common for iteration).
	// `range` returns two values for each element: the index and a copy of the value.
	fmt.Println("\n3b. Iterating using for...range (index and value):")
	for index, value := range stringArray {
		fmt.Printf("  Index: %d, Value: %s\n", index, value)
	}

	// 3c. Using `for...range` when you only need the value.
	// Use the blank identifier `_` to ignore the index.
	fmt.Println("\n3c. Iterating using for...range (value only, ignoring index):")
	for _, value := range integerArray {
		fmt.Printf("  Value: %d\n", value)
	}

	// --- 4. Arrays are Value Types (Copying) ---

	// When you assign one array to another, Go creates a *copy* of the original array.
	// They are independent in memory.
	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray // This creates a new array, a copy of originalArray

	fmt.Println("\n4a. Original Array before modification:", originalArray) // Output: [1 2 3]
	fmt.Println("4b. Copied Array before modification:", copiedArray)       // Output: [1 2 3]

	// Modifying the copied array does NOT affect the original array.
	copiedArray[0] = 100

	fmt.Println("4c. Original Array after modifying copiedArray:", originalArray) // Output: [1 2 3] (unchanged!)
	fmt.Println("4d. Copied Array after modification:", copiedArray)              // Output: [100 2 3] (changed)

	// --- 5. Array Comparison ---

	// Arrays can be compared using `==` and `!=` if their elements are comparable
	// and they have the exact same type (same element type AND same fixed length).
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	array3 := [3]int{100, 2, 3}
	// array4 := [4]int{1, 2, 3, 4} // Cannot compare array1 ([3]int) with array4 ([4]int) - different types

	fmt.Println("\n5a. Array1:", array1)
	fmt.Println("5b. Array2:", array2)
	fmt.Println("5c. Array3:", array3)

	fmt.Println("5d. Array1 == Array2:", array1 == array2) // Output: true (elements and type/length match)
	fmt.Println("5e. Array1 == Array3:", array1 == array3) // Output: false (elements differ)
	fmt.Println("5f. Array1 != Array3:", array1 != array3) // Output: true

	// --- 6. Multidimensional Arrays ---

	// Arrays can contain other arrays, creating multidimensional structures (like matrices).
	// This declares a 3x3 integer matrix.
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3}, // First row (index 0)
		{4, 5, 6}, // Second row (index 1)
		{7, 8, 9}, // Third row (index 2)
	}
	fmt.Println("\n6a. Multidimensional Array (Matrix):", matrix)
	fmt.Println("6b. Accessing element at [1][2] (row 1, column 2):", matrix[1][2]) // Output: 6

	// --- 7. Pointers to Arrays ---

	// You can get a pointer to an array using the `&` operator.
	// Modifying the array through the pointer *does* affect the original array.
	checkArray := [3]int{1, 2, 3}
	var pointerToArray *[3]int = &checkArray // pointerToArray now holds the memory address of checkArray

	fmt.Println("\n7a. Original array before pointer modification:", checkArray)
	fmt.Println("7b. Pointer to array (address):", pointerToArray)              // Prints memory address
	fmt.Println("7c. Dereferenced pointer (*pointerToArray):", *pointerToArray) // Prints the array value the pointer points to

	// Modify the array using the pointer. Go allows `pointerToArray[index]` as a shorthand for `(*pointerToArray)[index]`.
	pointerToArray[0] = 100

	fmt.Println("7d. Original array after pointerToArray[0] = 100:", checkArray) // Output: [100 2 3] (changed!)
	fmt.Println("7e. Pointer to array after modification:", *pointerToArray)     // Output: [100 2 3] (changed)

	// --- 8. Blank Identifier with Function Returns ---
	// The blank identifier `_` is also used to ignore unwanted return values from functions.
	value1, _ := returnMultipleValues()                                               // Call the function, ignore the second return value
	fmt.Println("\n8a. Calling returnMultipleValues, ignoring second value:", value1) // Output: 1

	value1Again, value2 := returnMultipleValues()                                         // Call again, capture both values
	fmt.Println("8b. Calling returnMultipleValues, capturing both:", value1Again, value2) // Output: 1 2

	fmt.Println("\n--- End of Arrays Tutorial ---")
}

// returnMultipleValues is a helper function that returns two integers.
// Used to demonstrate ignoring return values with the blank identifier.
func returnMultipleValues() (int, int) {
	return 1, 2
}

// Note: Slices are a more flexible and commonly used data structure in Go
// compared to arrays. Slices are built on top of arrays but have dynamic length.
// This tutorial focuses specifically on the fixed-length nature of arrays.
