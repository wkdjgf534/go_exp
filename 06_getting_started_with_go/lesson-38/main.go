package main

import "fmt"

func main() {
	var intPtr *int
	fmt.Println(intPtr) // nil

	age := 10
	intPtr = &age
	fmt.Println(intPtr)  // Print address
	fmt.Println(*intPtr) // 10

	name := "code & learn"
	namePtr := &name
	fmt.Println(*namePtr) // code & learn
}
