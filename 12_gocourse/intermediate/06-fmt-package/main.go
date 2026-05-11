package main

import "fmt"

func main() {
	// Printing Functions
	// fmt.Print("Hello")
	// fmt.Print("World!")
	// fmt.Print(12.456)

	// fmt.Println("Hello")
	// fmt.Println("World!")
	// fmt.Println(12.456)

	// name := "John"
	// age := 25
	// fmt.Printf("Name: %s, Age: %d\n", name, age)  // Name: John, Age: 25
	// fmt.Printf("Binary: %b, Hex: %X\n", age, age) // Binary: 11001, Hex: 19

	// Formatting Functions
	// fmt.Println("-----")
	// Doesn't add spaces between strings and oout put will be in one line
	// s := fmt.Sprint("Hello", "World", 123, 456)
	// fmt.Print(s) // HelloWorld123 456

	// s = fmt.Sprintln("Hello", "World!", 123, 456)
	// fmt.Print(s) // Hello World! 123 456

	// fmt.Spintf does not add \n at the end of the line
	// sf := fmt.Sprintf("Name: %s, Age: %d", name, age)
	// fmt.Print(sf) // Name: John, Age: 25

	// Scanning Functions
	var name string
	var age int
	// fmt.Println("-----")
	// fmt.Print("Enter your name and age:")
	// Name and age can be input separately with pressing enter button
	// fmt.Scan(&name, &age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)
	// Enter your name and age:John
	//
	// 25
	// Name: John, Age: 25

	// fmt.Print("Enter your name and age:")
	// Name and Age must be in one input, it stops scanning after press enter button
	// fmt.Scanln(&name, &age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)
	// Enter your name and age:John
	// Name: John, Age: 0

	fmt.Print("Enter your name and age:")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	// Enter your name and age:Peter 32
	// Name: Peter, Age: 32

	// Error Formatting Functions
	err := checkAge(15)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive", age)
	}

	return nil
}
