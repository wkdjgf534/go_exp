package main

import "fmt"

func panicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("Something went wrong")
}

func main() {
	fmt.Println("Start of the program")

	panicExample()

	fmt.Println("Exited function")
}
