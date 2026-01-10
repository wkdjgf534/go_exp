package main

import "fmt"

func main() {
	var a *int
	var i any

	fmt.Println(a == nil) // true
	fmt.Println(i == nil) // true

	i = a
	// interace has two pointers: type and value
	fmt.Printf("Value: %v\nType: %T\n", i, i)
	fmt.Println(i == nil) // false
}
