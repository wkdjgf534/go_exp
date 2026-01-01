package main

import "fmt"

func main() {
	firstDay := "Friday"

	switch firstDay {
	case "Monday":
		fmt.Println("It's Monday")
	case "Thuesday":
		fmt.Println("It's Tuesday")
	default:
		fmt.Println("It's not Monday or Tuesday")
	}

	word := "Test1"

	switch wordLen := len(word); wordLen {
	case 3:
		b := 4 // This variable is not accessible in the rest of the programm only here
		fmt.Println(b)
		fmt.Println("word length is 3")
	case 4:
		fmt.Println("word length is 4")
	default:
		fmt.Println("word length is neither 3 nor 4 but", wordLen)
	}

	newWord := "a"

	switch newWord {
	case "a":
		fallthrough
	case "b":
		fmt.Println("a or b")
	case "c", "d": // can accept several values
		fmt.Println("d")
	default:
		fmt.Println("not a, b, c or d")
	}
}
