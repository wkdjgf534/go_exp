package main

import "fmt"

func main() {
	var myString string
	// Zero value for string is ""
	fmt.Println(myString)
	myString = "Hello World"
	fmt.Println(myString)

	// in is an escape sequence for new line
	myString = "Hello\nWorld"
	fmt.Println(myString)

	myString = `Welcome to

Golang course`
	fmt.Println(myString)

	var firstName, lastName string
	firstName = "code"
	lastName = "learn"

	var fullName string
	fullName = firstName + " " + lastName
	fmt.Println(fullName)

	fmt.Printf("%s %s\n", firstName, lastName)

	fullName = fmt.Sprintf("%s %s", firstName, lastName)
	fmt.Println(fullName)
}
