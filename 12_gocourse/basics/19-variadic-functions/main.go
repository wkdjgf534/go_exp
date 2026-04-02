package main

import "fmt"

func main() {
	//... - Ellipsis
	// func functionName(param1 type1, param2 type2, param3 ...type3) returnType {
	// function body
	// }

	statement, total := sum("The sum of 1, 2, 3, 4, 5, 6 is", 1, 2, 3, 4, 5, 6)
	fmt.Println(statement, total)
}

func sum(returnString string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}

	return returnString, total
}
