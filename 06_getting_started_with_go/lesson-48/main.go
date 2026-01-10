package main

import "fmt"

func printValue(value any) {
	fmt.Println(value)
}

func checkType(value any) {
	switch value := value.(type) {
	case int:
		fmt.Println("type is an int: ", value)
	case string:
		fmt.Println("type is a string: ", value)
	case float64:
		fmt.Println("type is a float64: ", value)
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	misedSlice := []any{42, "Hello", 3.144, true}

	for _, value := range misedSlice {
		printValue(value)
	}

	var emptyInterface any
	emptyInterface = "Hello World!"

	if str, ok := emptyInterface.(string); ok {
		fmt.Println("The underlying value is a  tring:", str)
	}

	checkType(42)
	checkType("Hello World!")
	checkType(3.14)
	checkType(true)
}
