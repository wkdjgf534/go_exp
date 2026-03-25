package main

import (
	"fmt"
	"os"
)

func main() {

	// This will never be executed
	defer fmt.Println("Deferred staement")

	fmt.Println("Starting the main function")
	// Exit with status code of 1
	os.Exit(1)

	// This will never be executed too
	fmt.Println("End of main function")
}
