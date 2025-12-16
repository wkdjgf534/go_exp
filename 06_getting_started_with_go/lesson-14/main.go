package main

import (
	"fmt"
	"strconv"
)

// 1. Right shift operator >>
// 2. Left shift operator <<
// 3. Creating a bit mask
// 4. Set bit a given position
// 5. Unset a bit at a given position
// 6. Invert a bit at a given position
func main() {
	var a uint8 = 20 // 10100
	fmt.Println(strconv.FormatUint(uint64(a), 2), a)

	// Right shift operation
	a >>= 2 // Remove the last two bits -> 101 or 5
	fmt.Println(strconv.FormatUint(uint64(a), 2), a)

	// Left shift operation
	a <<= 2 // Append two zero bits -> 10100
	fmt.Println(strconv.FormatUint(uint64(a), 2), a)

	// Create a bit mask
	a = 1
	a <<= 3 // Bitmask where 3rd bit is set to 1
	fmt.Println(strconv.FormatUint(uint64(a), 2), a)

	// Set a bit a given position
	a = 5 | (1 << 1) // 101 -> 111 set bit a position 1 to 1
	// bitwise OR
	// 101 (5)
	// 010 (2) (1 << 1)
	// 111 (7)
	fmt.Println(strconv.FormatUint(uint64(a), 2), a)

	// Unset bit at a given position
	a = a &^ (1 << 1) // 111 -> 101 clear bit at postion 1
	fmt.Println(strconv.FormatUint(uint64(a), 2), a)

	// Toggle a bit at a given position
	a = a ^ (1 << 0) // 101 -> 100
	fmt.Println(strconv.FormatUint(uint64(a), 2), a)
}
