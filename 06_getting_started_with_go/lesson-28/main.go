package main

import "fmt"

// 1. The if statement
// 2. The if-else statement
// 3. The if-else-if statement
// 4. Variable declaration within if.
func main() {
	age := 14

	if age >= 18 {
		fmt.Println("You are an adult!")
	} else if age >= 13 {
		fmt.Println("You are teenager!")
	} else {
		fmt.Println("You are chiled!")
	}

	if even := isEven(age); even { // variable even is reachable only in this block of code
		fmt.Println("Age is even")
	}
}

// isEven - function return true if number is even
func isEven(n int) bool {
	return n&1 == 0
}
