package main

import "fmt"

func main() {
	// if condition {
	//	block of code
	// }

	// age := 25
	// if age >= 18 {
	//	fmt.Println("You are an adult.")
	// }

	// if condition {
	//	block of code
	// } else if {
	//	block of code
	//} else {
	//	block of code
	// }

	temperature := 25
	if temperature >= 30 {
		fmt.Println("It's hot outside.")
	} else {
		fmt.Println("It's cool outside.")
	}

	score := 65
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grafe B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	// if condition1 {
	//	code of block
	//	if condition 2 {
	//		code of block
	//	}
	// }

	// num := 16
	// if num%2 == 0 {
	//	if num%3 == 0 {
	//		fmt.Println("Number is divisible by both 2 and 3.")
	//	} else {
	//		fmt.Println("Number is divisible by 2 but not 3.")
	//	}
	// } else {
	//	fmt.Println("Number is not divisible by 2.")
	// }

	// || OR
	// && AND

	if 10%2 == 0 || 5%2 == 0 {
		fmt.Println("Either 10 or 5 are even")
	}

	if 10%2 == 0 && 6%2 == 0 {
		fmt.Println("Both 10 and 6 are even")
	}
}
