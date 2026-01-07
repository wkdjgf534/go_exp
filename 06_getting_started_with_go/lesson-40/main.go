package main

import (
	"errors"
	"fmt"
)

type student struct {
	firstName string
	lastName  string
}

func main() {
	a := 10
	increment(&a)
	fmt.Println(a) // 11

	s := student{
		firstName: "code",
		lastName:  "learn",
	}

	prevLastName, err := updateLastName(&s, "")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(prevLastName) // <nil> It is better to return a nil value when an empty string to minimize potential errors

	prettyPrintStudent(&s) // It is better to send pointer instead of strucutre
}

func updateLastName(s *student, newLastName string) (*string, error) {
	if newLastName == "" {
		return nil, errors.New("empty new last name")
	}

	previous := s.lastName
	s.lastName = newLastName
	return &previous, nil
}

func increment(x *int) {
	*x++
}

func prettyPrintStudent(s *student) {
	fmt.Printf("First name: %s\nLast Name: %s\n", s.firstName, s.lastName)
}
