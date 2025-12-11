package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint8 = 202
	var b uint8 = 141

	// The OR operator
	c := a | b
	fmt.Println(c)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(b), 2))
	fmt.Println(strconv.FormatUint(uint64(c), 2))

	// The AND operator
	d := a & b
	fmt.Println(d)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(b), 2))
	fmt.Println(strconv.FormatUint(uint64(d), 2))

	// The XOR operator
	e := a ^ b
	fmt.Println(e)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(b), 2))
	fmt.Println(strconv.FormatUint(uint64(e), 2))

	// The AND NOT or Bit clear operator
	f := a &^ b
	fmt.Println(f)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(b), 2))
	fmt.Println(strconv.FormatUint(uint64(f), 2))

	// The INVERT or the NOT operator
	g := ^a
	fmt.Println(g)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(g), 2))

	// To check if values is odd or even number
	fmt.Println(10 & 1) // 0
	fmt.Println(11 & 1) // 1
}
