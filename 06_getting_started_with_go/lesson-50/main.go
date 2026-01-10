package main

import (
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", c.Code, c.Message)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, CustomError{Code: -1, Message: "cannot divide by zero"}
		// return 0, errors.New("cannot divide by zero")
		// return 0, fmt.Errorf("cannot divide by zero %f", b) // This variant works slightly slower in new go version
	}

	return a / b, nil
}

func main() {
	result, err := Divide(10, 0)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println(result)
}
