package main

import (
	"fmt"
)

func main() {
	message := "Hello, "
	greetingFn := func(name string) {
		fmt.Println(message + name)
	}

	// LIFO
	defer greetingFn("Alice") // 3 Hi, Bob
	defer greetingFn("Bob")   // 4 Hi, Alice

	defer func(name string) {
		fmt.Println(message + name)
	}("Peter") // 2 Hi, Peter

	// os.Exit(1) // in this case defer wont print anything. It exists before the main func ends.

	fmt.Println("Test") // 1

	message = "Hi, "
}
