package main

import "fmt"

func function1() {
	defer func() {
		fmt.Println("function1 deffered function called")
	}()

	function2()
}

func function2() {
	defer func() {
		fmt.Println("function2 deffered function called")
	}()

	panic("function2 panicked")
}

func main() {
	defer func() {
		fmt.Println("main function deffered")
	}()

	function1()
}
