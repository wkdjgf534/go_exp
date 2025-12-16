package main

import "fmt"

// In this lesson we will cover:
// 1. Different ways of array declaration.
// 2. The len function.
// 3. Accessing array elements by their index.
// 4. Trying to access array elements outside range.
// 5. Two dimensional array.
func main() {
	// Array declaration basic form
	var arr1 [5]int
	fmt.Println(arr1) // -> [0, 0, 0, 0, 0]

	// Array declaration with elements
	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2) // -> [1, 2, 3, 4, 5]

	// Sparse array declaration
	arr3 := [5]int{5, 2: 10, 50}
	fmt.Println(arr3) // -> [5, 0, 10, 50, 0]

	// Implicit length declaration
	arr4 := [...]int{3, 4, 5, 1, 2}
	fmt.Println(arr4)
	arr4[1] = 10
	fmt.Println(arr4) // -> [3, 10, 5, 1, 2]
	fmt.Println(len(arr4))

	// Accessing array elements
	arr5 := arr4[4]
	fmt.Println(arr5)

	// Two dimensional array
	arr6 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr6)
}
