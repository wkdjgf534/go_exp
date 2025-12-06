package main

import "fmt"

// Group similar variables
//var (
//	myInt    int    = 10
//	myString string = "Hello World!"
//)

func main() {
	// Explicit type assigment
	var myInt int = 10
	var myString string = "Hello World!"
	fmt.Println(myInt)
	fmt.Println(myString)

	// Implicit type assigment
	//var age = 25
	//fmt.Println(age)

	// Multiple variable declaration
	//var year, name = 2023, "code & learn"
	//fmt.Println(year, name)

	// Shorthand variable declaration
	age := 30
	year := 2023
	fmt.Println(age, year)

	year, name := 2024, "code & learn"
	fmt.Println(year, name)

	name = "Code & Learn"
	fmt.Println(year, name)

	// Examples:
	// var apartmentNumber int        // Zero value to the variable
	// var apartmentNumber int = 2000 // When you want to assign a value to variable in declaration
	// var apartmentNumber = 200
	// var apartmentNumber, streetName = 2000, "Bayshore" // Multiple variable declaration on the same line
	// apartmentNumber := 2000
	apartmentNumber, streetName := 2000, "Bayshore"
	fmt.Println(apartmentNumber, streetName)
}
