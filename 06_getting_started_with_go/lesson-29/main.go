package main

import "fmt"

var (
	// Package level variable
	a = 10
)

func main() {
	{
		// Shadowing
		// block scoped variable a
		a := 15
		fmt.Println(a)
	}

	fmt.Println(a)

	something()
}
