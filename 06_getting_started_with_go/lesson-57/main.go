package main

import (
	"fmt"

	_ "lesson-57/init"
)

func init() {
	fmt.Println("init 1 called")
}

func init() {
	fmt.Println("init 2 called")
}

func init() {
	fmt.Println("init 3 called")
}

func main() {
	fmt.Println("main function called")
}

// init packadge init function called
// init 1 called
// init 2 called
// init 3 called
// main function called
