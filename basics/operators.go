package main

import "fmt" // Import fmt for printing output

func main() {
	// Go supports various operators to perform operations on variables and values.

	// --- Logical Operators ---
	// Logical operators are used to combine or modify boolean expressions.
	// !  (Logical NOT): Reverses the boolean value. If a condition is true, ! makes it false, and vice-versa.
	// || (Logical OR):  Returns true if at least one of the operands is true.
	// && (Logical AND): Returns true if both operands are true.

	isSunny := true
	isRaining := false
	fmt.Println("--- Logical Operators ---")
	fmt.Printf("isSunny: %t, isRaining: %t\n", isSunny, isRaining)
	fmt.Println("!isSunny (NOT sunny?):", !isSunny)                                       // Example of !
	fmt.Println("isSunny || isRaining (Sunny OR Raining?):", isSunny || isRaining)        // Example of ||
	fmt.Println("isSunny && !isRaining (Sunny AND NOT Raining?):", isSunny && !isRaining) // Example of &&
	fmt.Println()                                                                         // Add a blank line for readability

	// --- Bitwise Operators ---
	// Bitwise operators perform operations on the individual bits of integer types.
	// &  (Bitwise AND): Sets each bit to 1 if both corresponding bits are 1.
	// |  (Bitwise OR):  Sets each bit to 1 if at least one of the corresponding bits is 1.
	// ^  (Bitwise XOR): Sets each bit to 1 if the corresponding bits are different.
	// &^ (Bitwise AND NOT / Bit Clear): Clears (sets to 0) the bits in the first operand
	//                                   that are set to 1 in the second operand.
	// << (Left Shift):   Shifts bits to the left, filling new bits on the right with 0s.
	//                    Each left shift multiplies the number by 2.
	// >> (Right Shift):  Shifts bits to the right. For unsigned integers, new bits on the left
	//                    are filled with 0s (logical shift). For signed integers, the sign bit
	//                    is used to fill the new bits (arithmetic shift). Each right shift divides by 2 (integer division).

	a := 5 // Binary: 0101
	b := 3 // Binary: 0011

	fmt.Println("--- Bitwise Operators ---")
	fmt.Printf("a = %d (binary %04b), b = %d (binary %04b)\n", a, a, b, b)
	fmt.Printf("a & b  (AND)  = %d (binary %04b)\n", a&b, a&b)      // 0101 & 0011 = 0001 (1)
	fmt.Printf("a | b  (OR)   = %d (binary %04b)\n", a|b, a|b)      // 0101 | 0011 = 0111 (7)
	fmt.Printf("a ^ b  (XOR)  = %d (binary %04b)\n", a^b, a^b)      // 0101 ^ 0011 = 0110 (6)
	fmt.Printf("a &^ b (AND NOT) = %d (binary %04b)\n", a&^b, a&^b) // 0101 &^ 0011 = 0100 (4)

	numToShift := uint(4) // Binary: 00000100
	fmt.Printf("\nnumToShift = %d (binary %08b)\n", numToShift, numToShift)
	fmt.Printf("numToShift << 1 (Left Shift by 1) = %d (binary %08b)\n", numToShift<<1, numToShift<<1)  // 00000100 << 1 = 00001000 (8)
	fmt.Printf("numToShift >> 1 (Right Shift by 1) = %d (binary %08b)\n", numToShift>>1, numToShift>>1) // 00000100 >> 1 = 00000010 (2)
	fmt.Println()

	// --- Comparison Operators ---
	// Comparison operators are used to compare two values and return a boolean result (true or false).
	// == (Equal to):          Returns true if both operands are equal.
	// != (Not equal to):       Returns true if operands are not equal.
	// <  (Less than):          Returns true if the left operand is less than the right.
	// <= (Less than or equal to): Returns true if the left operand is less than or equal to the right.
	// >  (Greater than):       Returns true if the left operand is greater than the right.
	// >= (Greater than or equal to): Returns true if the left operand is greater than or equal to the right.

	x := 10
	y := 20
	z := 10
	fmt.Println("--- Comparison Operators ---")
	fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)
	fmt.Println("x == y (x equals y?):", x == y)
	fmt.Println("x == z (x equals z?):", x == z)
	fmt.Println("x != y (x not equal to y?):", x != y)
	fmt.Println("x < y  (x less than y?):", x < y)
	fmt.Println("x <= z (x less than or equal to z?):", x <= z)
	fmt.Println("y > x  (y greater than x?):", y > x)
	fmt.Println("x >= z (x greater than or equal to z?):", x >= z)
	fmt.Println()

	// --- Assignment Operators ---
	// Assignment operators are used to assign values to variables.
	// Go also provides compound assignment operators that combine an arithmetic or bitwise operation with assignment.
	// =   (Simple Assignment): Assigns the value of the right operand to the left operand.
	// +=  (Add and Assign): Adds the right operand to the left operand and assigns the result to the left operand. (e.g., c += 5 is c = c + 5)
	// -=  (Subtract and Assign): Subtracts the right operand from the left operand and assigns the result. (e.g., c -= 3 is c = c - 3)
	// *=  (Multiply and Assign): Multiplies the left operand by the right operand and assigns the result. (e.g., c *= 2 is c = c * 2)
	// /=  (Divide and Assign): Divides the left operand by the right operand and assigns the result. (e.g., c /= 4 is c = c / 4)
	// %=  (Modulo and Assign): Takes modulo using two operands and assigns the result to the left operand. (e.g., c %= 2 is c = c % 2)
	// &=  (Bitwise AND Assign): Performs Bitwise AND and assigns the result. (e.g., c &= 2 is c = c & 2)
	// |=  (Bitwise OR Assign): Performs Bitwise OR and assigns the result. (e.g., c |= 2 is c = c | 2)
	// ^=  (Bitwise XOR Assign): Performs Bitwise XOR and assigns the result. (e.g., c ^= 2 is c = c ^ 2)
	// &^= (Bitwise AND NOT Assign): Performs Bitwise AND NOT and assigns the result. (e.g., c &^= 2 is c = c &^ 2)
	// <<= (Left Shift Assign): Performs Left Shift and assigns the result. (e.g., c <<= 1 is c = c << 1)
	// >>= (Right Shift Assign): Performs Right Shift and assigns the result. (e.g., c >>= 1 is c = c >> 1)

	fmt.Println("--- Assignment Operators ---")
	var c int = 10 // Start with an initial value for c
	fmt.Printf("Initial c: %d\n", c)

	c += 5 // c is now 10 + 5 = 15
	fmt.Printf("c += 5: %d\n", c)

	c -= 3 // c is now 15 - 3 = 12
	fmt.Printf("c -= 3: %d\n", c)

	c *= 2 // c is now 12 * 2 = 24
	fmt.Printf("c *= 2: %d\n", c)

	c /= 4 // c is now 24 / 4 = 6
	fmt.Printf("c /= 4: %d\n", c)

	c %= 5                         // c is now 6 % 5 = 1
	fmt.Printf("c %%= 5: %d\n", c) // Note: %% is used to print a literal % in Printf
	fmt.Println()
}
