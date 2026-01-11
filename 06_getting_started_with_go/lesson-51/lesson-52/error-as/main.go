package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// CustomeError is a custom type that implements the error interface.
type CustomError struct {
	Message string
	Status  int
}

// Error method make CustomError implement the error interface.
func (c CustomError) Error() string {
	return c.Message
}

func SomeFunction() error {
	return CustomError{Message: "original error: something went wrong", Status: 400}
}

func main() {
	err := SomeFunction()

	var customErr CustomError
	if errors.As(err, &customErr) {
		fmt.Println("Extracted CustomError:", customErr.Status) // Extracted CustomError: 400
	}

	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Op) // Failed at path: open
		}
	}
}
