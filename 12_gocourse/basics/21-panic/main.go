package main

import "fmt"

func main() {
	// panic(interface{})

	// Example of a valid input
	process(10)

	// Example of a invalid input
	process(-3)
}

func process(input int) {

	defer fmt.Println("Deferred 1") // 3
	defer fmt.Println("Deferred 2") // 2

	if input < 0 {
		fmt.Println("Before Panic") // 1
		panic("input must be a non-negative number")

		// defer fmt.Println("Deffered 3") unreachable after panic
	}
	fmt.Println("Processing input:", input)
}
