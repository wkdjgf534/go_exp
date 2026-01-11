package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Message string
	Wrapped error
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Wrapped)
}

func (e CustomError) Unwrap() error {
	return e.Wrapped
}

func firstFunction() error {
	return fmt.Errorf("original error: something went wrong in first function")
}

func secondFunction() error {
	firstErr := firstFunction()
	if firstErr != nil {
		secondErr := errors.New("failed in second function")
		return errors.Join(firstErr, secondErr) // print out 2 strings with errors
		// original error: something went wrong in first function
		// failed in second function

		// %w - wrap message
		// return fmt.Errorf("failed in the second function: %w", firstErr) // failed in the second function: original error: something went wrong in first function
	}
	return nil
}

func thirdFunction() error {
	err := firstFunction()
	if err != nil {
		// Wrap the error with additional context
		return fmt.Errorf("failed in Another function: %w", err)
	}

	return nil
}

func SomeFunction() error {
	return CustomError{
		Message: "original error: something went wrong",
		Wrapped: errors.New("wrapped error"),
	}
}

func main() {
	// err := secondFunction()
	// fmt.Println(err)

	// err := thirdFunction()
	// fmt.Println("Original error:", err)
	// innerError := errors.Unwrap(err)
	// fmt.Println(innerError)

	error := SomeFunction()
	fmt.Println("Error:", error) // Error: original error: something went wrong: wrapped error

	error = errors.Unwrap(error)

	innerError := fmt.Errorf("innermost: %w", error)
	fmt.Println("Innermost Error:", innerError) // Innermost Error: innermost: wrapped error
}
