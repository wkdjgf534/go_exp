package main

import "fmt"

const (
	x                    = 10 // It works for int, float etc
	y              int32 = 15 // Only for int32
	applcationName       = "Lesson 7"

	// All values that can be used as a constant
	isRunning = true
	character = 'a'

	// complex, real, imag, len and cap
	isTRue = 1 < 2
)

func main() {
	var a int
	a = x
	fmt.Println(a)

	var b float64
	b = x
	fmt.Println(b)

	var c int
	c = int(y)
	fmt.Println(c)

	var d int32
	d = y
	fmt.Println(d)

	// const z = a + int(b) // a + int(b) (value of type int) is not constant

	// it is possible to create const with complex, real, imag, len, cap
	const z = complex(10.2, 100.9)
	const l = imag(z)
	fmt.Println(z, l)
}
