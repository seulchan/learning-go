package main

import (
	"fmt"    // For formatted I/O, like printing to the console.
	"slices" // Go 1.18+ package providing utility functions for slices.
)

func main() {
	fmt.Println("--- Go Slices Tutorial ---")

	// --- What is a Slice? ---
	// A slice is a dynamically-sized, flexible view into the elements of an array.
	// Slices are much more common in Go than arrays because of their flexibility.
	// A slice has three components:
	// 1. Pointer: Points to the first element of the array accessible through the slice.
	// 2. Length: The number of elements in the slice.
	// 3. Capacity: The number of elements in the underlying array from the start of the slice.

	// --- 1. Declaring and Initializing Slices ---
	fmt.Println("\n--- 1. Declaring and Initializing Slices ---")

	// a. Declaring a nil slice
	// A nil slice has a length and capacity of 0 and does not point to an underlying array.
	var nilSlice []int
	fmt.Printf("nilSlice: %v, Length: %d, Capacity: %d, Is nil? %t\n", nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)

	// b. Using a slice literal
	// This creates a slice and an underlying array to hold the elements.
	literalSlice := []int{10, 20, 30}
	fmt.Printf("literalSlice: %v, Length: %d, Capacity: %d\n", literalSlice, len(literalSlice), cap(literalSlice))

	// c. Using the `make` function
	// `make` allocates an underlying array and returns a slice that refers to it.
	// `make([]type, length, capacity)`
	madeSlice := make([]int, 3, 5) // Length 3, Capacity 5
	fmt.Printf("madeSlice (initially): %v, Length: %d, Capacity: %d\n", madeSlice, len(madeSlice), cap(madeSlice))
	// Elements are zero-initialized (0 for int).

	// d. Slicing an existing array or slice
	// This creates a new slice that shares the underlying array.
	sourceArray := [5]int{1, 2, 3, 4, 5} // An array
	fmt.Printf("sourceArray: %v\n", sourceArray)

	// `sliceFromArray` will include elements from index 1 up to (but not including) index 4.
	// So, elements at index 1, 2, 3 of `sourceArray`.
	sliceFromArray := sourceArray[1:4] // Elements: sourceArray[1], sourceArray[2], sourceArray[3]
	// Length of sliceFromArray is 4 - 1 = 3.
	// Capacity is from sourceArray[1] to the end of sourceArray, which is 5 - 1 = 4.
	fmt.Printf("sliceFromArray (from sourceArray[1:4]): %v, Length: %d, Capacity: %d\n", sliceFromArray, len(sliceFromArray), cap(sliceFromArray))

	// --- 2. Length and Capacity ---
	// `len(s)`: Returns the number of elements in slice `s`.
	// `cap(s)`: Returns the capacity of slice `s` (elements available in underlying array).
	fmt.Printf("Recap - sliceFromArray: Length %d, Capacity %d\n", len(sliceFromArray), cap(sliceFromArray))

	// --- 3. Accessing and Modifying Slice Elements ---
	fmt.Println("\n--- 3. Accessing and Modifying Slice Elements ---")
	if len(sliceFromArray) > 1 {
		fmt.Printf("Element at index 1 of sliceFromArray: %d\n", sliceFromArray[1]) // Accesses the second element (value 3)
		sliceFromArray[1] = 33                                                      // Modify the element
		fmt.Printf("Modified sliceFromArray: %v\n", sliceFromArray)
		// IMPORTANT: Since sliceFromArray shares its underlying array with sourceArray,
		// modifications to sliceFromArray can affect sourceArray.
		// sliceFromArray[1] corresponds to sourceArray[2].
		fmt.Printf("sourceArray after modifying sliceFromArray: %v\n", sourceArray) // sourceArray is now [1 2 33 4 5]
	}

	// --- 4. Appending to a Slice using `append` ---
	fmt.Println("\n--- 4. Appending to a Slice ---")
	fmt.Printf("Original sliceFromArray: %v, Len: %d, Cap: %d\n", sliceFromArray, len(sliceFromArray), cap(sliceFromArray))

	// `append` adds elements to the end of a slice.
	// If the capacity is sufficient, the underlying array is reused.
	// If not, a new, larger underlying array is allocated, and elements are copied.
	sliceFromArray = append(sliceFromArray, 60) // Appending one element. Capacity was 4, len was 3. Now len is 4, cap is 4.
	fmt.Printf("After append(60): %v, Len: %d, Cap: %d\n", sliceFromArray, len(sliceFromArray), cap(sliceFromArray))
	fmt.Printf("sourceArray after first append on sliceFromArray: %v (modified because capacity was sufficient)\n", sourceArray)

	// Appending more elements, which might exceed current capacity
	sliceFromArray = append(sliceFromArray, 70, 80) // Now len will be 6. Original cap was 4.
	// A new underlying array will be allocated for sliceFromArray.
	// sourceArray will no longer be affected by changes to sliceFromArray.
	fmt.Printf("After append(70, 80): %v, Len: %d, Cap: %d (Cap likely doubled)\n", sliceFromArray, len(sliceFromArray), cap(sliceFromArray))
	fmt.Printf("sourceArray after second append (new array for slice): %v (no longer shares with sliceFromArray)\n", sourceArray)

	// --- 5. Copying Slices using `copy` ---
	// To create a truly independent slice (with its own underlying data), use `copy`.
	fmt.Println("\n--- 5. Copying Slices ---")
	sliceToCopyFrom := []string{"a", "b", "c"}
	// Create a destination slice with the same length.
	copiedSlice := make([]string, len(sliceToCopyFrom))
	// `copy` returns the number of elements copied.
	numCopied := copy(copiedSlice, sliceToCopyFrom)
	fmt.Printf("sliceToCopyFrom: %v\n", sliceToCopyFrom)
	fmt.Printf("copiedSlice: %v, Elements copied: %d\n", copiedSlice, numCopied)

	// Modifying copiedSlice will not affect sliceToCopyFrom.
	copiedSlice[0] = "X"
	fmt.Printf("After modifying copiedSlice: %v\n", copiedSlice)
	fmt.Printf("sliceToCopyFrom remains unchanged: %v\n", sliceToCopyFrom)

	// --- 6. Iterating Over Slices using `for...range` ---
	fmt.Println("\n--- 6. Iterating Over Slices ---")
	fmt.Println("Iterating over sliceFromArray:")
	for index, value := range sliceFromArray {
		fmt.Printf("  Index: %d, Value: %d\n", index, value)
	}

	// If you only need the value:
	fmt.Println("Iterating (value only):")
	for _, value := range sliceFromArray {
		fmt.Printf("  Value: %d\n", value)
	}

	// --- 7. Comparing Slices using `slices.Equal` (Go 1.18+) ---
	fmt.Println("\n--- 7. Comparing Slices ---")
	sliceOne := []int{1, 2, 3}
	sliceTwo := []int{1, 2, 3}
	sliceThree := []int{1, 2, 4}

	fmt.Printf("sliceOne: %v, sliceTwo: %v, sliceThree: %v\n", sliceOne, sliceTwo, sliceThree)
	if slices.Equal(sliceOne, sliceTwo) {
		fmt.Println("sliceOne and sliceTwo are equal.")
	} else {
		fmt.Println("sliceOne and sliceTwo are NOT equal.")
	}
	if slices.Equal(sliceOne, sliceThree) {
		fmt.Println("sliceOne and sliceThree are equal.")
	} else {
		fmt.Println("sliceOne and sliceThree are NOT equal.")
	}

	// --- 8. Reslicing (Creating Sub-slices) ---
	fmt.Println("\n--- 8. Reslicing ---")
	// `sliceFromArray` currently is: [2 33 4 60 70 80]
	fmt.Printf("Current sliceFromArray: %v, Len: %d, Cap: %d\n", sliceFromArray, len(sliceFromArray), cap(sliceFromArray))

	// Create a new slice `subSlice` from `sliceFromArray`.
	// It will share the same underlying array.
	// `subSlice` will contain elements from index 2 up to (but not including) index 4 of `sliceFromArray`.
	// Elements: sliceFromArray[2], sliceFromArray[3] which are 4 and 60.
	subSlice := sliceFromArray[2:4]
	// Length of subSlice is 4 - 2 = 2.
	// Capacity of subSlice is from sliceFromArray[2] to the end of sliceFromArray's capacity.
	// If sliceFromArray has len 6, cap 8, then cap(subSlice) = cap(sliceFromArray) - 2 = 8 - 2 = 6.
	fmt.Printf("subSlice (sliceFromArray[2:4]): %v, Len: %d, Cap: %d\n", subSlice, len(subSlice), cap(subSlice))

	// Modifying subSlice affects sliceFromArray because they share the underlying array.
	subSlice[0] = 400 // This is sliceFromArray[2]
	fmt.Printf("After modifying subSlice[0]: %v\n", subSlice)
	fmt.Printf("sliceFromArray after subSlice modification: %v\n", sliceFromArray)

	// --- 9. Multi-dimensional Slices (Slice of Slices) ---
	fmt.Println("\n--- 9. Multi-dimensional Slices ---")
	// You can create slices of slices, similar to 2D arrays but more flexible.
	rows := 3
	twoDSlice := make([][]int, rows) // Create a slice of `rows` nil slices

	fmt.Println("Initializing a 2D slice (jagged array):")
	for i := 0; i < rows; i++ {
		innerLength := i + 1 // Each inner slice will have a different length
		twoDSlice[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoDSlice[i][j] = i*10 + j
			// fmt.Printf("  twoDSlice[%d][%d] = %d\n", i, j, twoDSlice[i][j]) // Optional: print each assignment
		}
	}
	fmt.Println("Final 2D Slice (twoDSlice):", twoDSlice)
	// Example: twoDSlice might look like [[0] [10 11] [20 21 22]]

	fmt.Println("\n--- End of Slices Tutorial ---")
}
