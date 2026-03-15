package main

import (
	"fmt"
	"math"
)

func main() {
	// Variable declaration
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtaraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	const p float64 = 22 / 7.0
	fmt.Println(p)

	// Overflow with signed integers
	var maxInt int64 = 9223372036854775807 // max value that int64 can hold
	fmt.Println(maxInt)                    // 9223372036854775807

	maxInt += 1
	fmt.Println(maxInt) // -9223372036854775808

	// Overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615 // max value for uint64 type
	fmt.Println(uMaxInt)                      // 18446744073709551615

	uMaxInt += 1
	fmt.Println(uMaxInt) // 0

	// Underflow with floating point numbers
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat) // 1.0e-323

	smallFloat = smallFloat / math.MaxFloat32
	fmt.Println(smallFloat) // 0
}
