package main

import (
	"fmt"  // For printing to the console
	"maps" // Go 1.21+ package for map utility functions (like maps.Equal)
)

func main() {
	fmt.Println("--- Go Maps Tutorial ---")

	// --- What is a Map? ---
	// A map is an unordered collection of key-value pairs.
	// Keys must be of a comparable type (e.g., string, int, float, bool, pointers, structs if all fields are comparable).
	// Values can be of any type.
	// Maps are reference types, like slices and channels.

	// --- 1. Declaring and Initializing Maps ---
	fmt.Println("\n--- 1. Declaring and Initializing Maps ---")

	// a. Declaring a map (it will be nil)
	// A nil map has no keys and cannot have keys added to it directly.
	var studentAges map[string]int
	fmt.Printf("Declared map (studentAges): %v, Is nil? %t\n", studentAges, studentAges == nil)
	// Attempting to add to a nil map will cause a runtime panic:
	// studentAges["Alice"] = 30 // This would panic!

	// b. Initializing an empty map using `make`
	// This is the common way to create a map that's ready to use.
	fruitPrices := make(map[string]float64)
	fmt.Printf("Initialized empty map (fruitPrices) using make: %v, Is nil? %t\n", fruitPrices, fruitPrices == nil)

	// c. Initializing a map with a map literal
	// This creates and initializes the map with some key-value pairs.
	capitals := map[string]string{
		"USA":    "Washington D.C.",
		"France": "Paris",
		"Japan":  "Tokyo",
	}
	fmt.Printf("Initialized map (capitals) with literal: %v\n", capitals)

	// --- 2. Basic Map Operations ---
	fmt.Println("\n--- 2. Basic Map Operations ---")

	// a. Adding or Updating elements
	// If the key exists, its value is updated. If not, a new key-value pair is added.
	fruitPrices["apple"] = 0.75
	fruitPrices["banana"] = 0.50
	fruitPrices["orange"] = 0.60
	fmt.Println("fruitPrices after adding elements:", fruitPrices)

	fruitPrices["apple"] = 0.80 // Update the price of "apple"
	fmt.Println("fruitPrices after updating 'apple':", fruitPrices)

	// b. Accessing elements
	// Accessing a key returns its value.
	priceOfApple := fruitPrices["apple"]
	fmt.Printf("Price of apple: %.2f\n", priceOfApple)

	// Accessing a non-existent key
	// If a key doesn't exist, accessing it returns the zero value for the map's value type.
	// For float64, the zero value is 0.0. For string, it's "". For int, it's 0.
	priceOfMango := fruitPrices["mango"] // "mango" is not in the map
	fmt.Printf("Price of mango (non-existent key, zero value): %.2f\n", priceOfMango)

	// c. Checking for key existence (the "comma ok" idiom)
	// To distinguish between a key that exists with a zero value and a key that doesn't exist,
	// use the two-value assignment.
	value, ok := fruitPrices["orange"]
	if ok {
		fmt.Printf("Price of orange: %.2f (found)\n", value)
	} else {
		fmt.Println("Orange not found in fruitPrices.")
	}

	value, ok = fruitPrices["grape"] // "grape" is not in the map
	if ok {
		fmt.Printf("Price of grape: %.2f (found)\n", value)
	} else {
		fmt.Println("Grape not found in fruitPrices. 'value' is zero:", value, "'ok' is false:", ok)
	}

	// d. Getting the length of a map (number of key-value pairs)
	fmt.Printf("Number of items in fruitPrices: %d\n", len(fruitPrices))
	fmt.Printf("Number of items in capitals: %d\n", len(capitals))

	// e. Deleting elements
	// The `delete` built-in function removes a key-value pair from a map.
	// Deleting a non-existent key does nothing (no error).
	delete(fruitPrices, "banana")
	fmt.Println("fruitPrices after deleting 'banana':", fruitPrices)
	delete(fruitPrices, "grape") // "grape" doesn't exist, no error
	fmt.Println("fruitPrices after attempting to delete 'grape':", fruitPrices)

	// --- 3. Iterating Over Maps ---
	fmt.Println("\n--- 3. Iterating Over Maps ---")
	// Use a `for...range` loop to iterate over a map.
	// IMPORTANT: The order of iteration over a map is not guaranteed and can vary.
	// If you need a specific order, you must extract keys, sort them, then iterate.

	fmt.Println("Iterating over capitals (key and value):")
	for country, capitalCity := range capitals {
		fmt.Printf("  The capital of %s is %s\n", country, capitalCity)
	}

	// Iterating and using only keys
	fmt.Println("Iterating over capitals (keys only):")
	for country := range capitals {
		fmt.Printf("  Country: %s\n", country)
	}

	// Iterating and using only values
	fmt.Println("Iterating over capitals (values only):")
	for _, capitalCity := range capitals {
		fmt.Printf("  Capital City: %s\n", capitalCity)
	}

	// --- 4. Nil Maps Revisited ---
	fmt.Println("\n--- 4. Nil Maps Revisited ---")
	var preferences map[string]string // This is a nil map

	if preferences == nil {
		fmt.Println("The 'preferences' map is currently nil.")
	}

	// Reading from a nil map is safe and returns the zero value for the value type.
	colorPreference := preferences["color"]
	fmt.Printf("Reading 'color' from nil 'preferences' map: '%s' (zero value for string)\n", colorPreference)

	// Writing to a nil map will cause a runtime panic.
	// preferences["color"] = "blue" // Uncommenting this line would cause a panic!

	// To use a nil map, it must first be initialized using `make`.
	preferences = make(map[string]string)
	preferences["color"] = "blue"
	preferences["font"] = "Arial"
	fmt.Println("Initialized 'preferences' map and added items:", preferences)

	// --- 5. Clearing a Map (Go 1.21+) ---
	fmt.Println("\n--- 5. Clearing a Map (Go 1.21+) ---")
	// The `clear` built-in function removes all entries from a map, leaving it empty.
	// This feature was added in Go 1.21.
	fmt.Println("Capitals map before clear:", capitals)
	if len(capitals) > 0 { // Check if map is not empty before clearing, for older Go versions this check would be more relevant
		clear(capitals) // Requires Go 1.21+
		fmt.Println("Capitals map after clear:", capitals)
		fmt.Printf("Length of capitals map after clear: %d\n", len(capitals))
	} else {
		fmt.Println("Capitals map is already empty or clear() is not available (pre-Go 1.21).")
	}

	// --- 6. Comparing Maps (Go 1.21+) ---
	fmt.Println("\n--- 6. Comparing Maps (Go 1.21+) ---")
	// Maps can be compared for equality using the `maps.Equal` function from the `maps` package.
	// This requires Go 1.21+.
	// Two maps are equal if they are both nil, or if they are non-nil, have the same length,
	// and contain the same key-value pairs.

	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 2, "a": 1} // Same elements, different order in literal
	map3 := map[string]int{"a": 1, "c": 3}

	if maps.Equal(map1, map2) { // Requires Go 1.21+
		fmt.Println("map1 and map2 are equal.")
	} else {
		fmt.Println("map1 and map2 are NOT equal (or maps.Equal not available).")
	}

	if maps.Equal(map1, map3) { // Requires Go 1.21+
		fmt.Println("map1 and map3 are equal.")
	} else {
		fmt.Println("map1 and map3 are NOT equal (or maps.Equal not available).")
	}
	// Note: Before Go 1.21, you'd have to write a custom function to compare maps by
	// checking lengths and then iterating over one map to see if all key-value pairs exist in the other.

	// --- 7. Nested Maps (Map of Maps) ---
	fmt.Println("\n--- 7. Nested Maps (Map of Maps) ---")
	// Map values can themselves be maps, creating nested structures.
	userPermissions := make(map[string]map[string]bool)

	// Initialize the inner map for "alice"
	userPermissions["alice"] = make(map[string]bool)
	userPermissions["alice"]["read"] = true
	userPermissions["alice"]["write"] = false

	// Initialize the inner map for "bob"
	userPermissions["bob"] = make(map[string]bool)
	userPermissions["bob"]["read"] = true
	userPermissions["bob"]["write"] = true

	fmt.Println("User Permissions (nested map):", userPermissions)
	fmt.Println("Alice's write permission:", userPermissions["alice"]["write"])
	fmt.Println("Bob's read permission:", userPermissions["bob"]["read"])

	fmt.Println("\n--- End of Maps Tutorial ---")
}
