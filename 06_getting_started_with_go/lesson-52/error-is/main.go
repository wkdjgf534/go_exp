package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

var ErrFirstError = CustomError{Message: "first error"}

type CustomError struct {
	Message string
}

func (c CustomError) Error() string {
	return c.Message
}

func someFunction() error {
	return fmt.Errorf("some function err: %w", ErrFirstError)
}

func main() {
	err := someFunction()
	if err != nil {
		if errors.Is(err, ErrFirstError) {
			fmt.Println("sentinel error found")
		}
	}

	if _, err := os.Open("non-existing"); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}
}
