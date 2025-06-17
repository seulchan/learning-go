// Package main is the entry point for the calculator application.
// This simple program demonstrates how to use the 'calculator' package.
package main

import (
	"calculator" // Importing our custom calculator package.
	"fmt"        // For formatted I/O, like printing to the console.
	"log"        // For logging errors.
)

// main is the function where program execution begins.
func main() {
	fmt.Println("Calculator Demo")
	fmt.Println("---------------")

	// Demonstrate Add
	sum := calculator.Add(5, 3)
	fmt.Printf("5 + 3 = %f\n", sum)

	// Demonstrate Subtract
	difference := calculator.Subtract(10, 4)
	fmt.Printf("10 - 4 = %f\n", difference)

	// Demonstrate Multiply
	product := calculator.Multiply(6, 7)
	fmt.Printf("6 * 7 = %f\n", product)

	// Demonstrate Divide (successful case)
	quotient, err := calculator.Divide(20, 4)
	if err != nil {
		// This should not happen with these inputs.
		log.Fatalf("Error dividing 20 by 4: %v", err)
	}
	fmt.Printf("20 / 4 = %f\n", quotient)

	// Demonstrate Divide (division by zero)
	_, err = calculator.Divide(10, 0)
	if err != nil {
		fmt.Printf("Attempting to divide 10 by 0: Error: %v\n", err)
	}

	// Demonstrate Sqrt (successful case)
	root, err := calculator.Sqrt(25)
	if err != nil {
		log.Fatalf("Error calculating sqrt of 25: %v", err)
	}
	fmt.Printf("Square root of 25 = %f\n", root)

	// Demonstrate Sqrt (negative input)
	_, err = calculator.Sqrt(-9)
	if err != nil {
		fmt.Printf("Attempting square root of -9: Error: %v\n", err)
	}
}
