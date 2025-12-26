package main

import "fmt"

// 1. different slice expressions.
// 2. The problem with slicing.
// 3. Passing a capacity to the slice expression.
// 4. The copy function.
func main() {
	exp1 := []int{1, 2, 3, 4, 5, 6}
	exp2 := exp1[2:5] // a range from 3 to 5
	fmt.Println(exp2) // [3, 4, 5]

	exp2 = exp1[:5]   // a range from the beginnig of values to 5
	fmt.Println(exp2) // [1, 2, 3, 4, 5]

	exp2 = exp1[3:]   // a range from 4 to the end of values
	fmt.Println(exp2) // [4, 5, 6]

	exp2 = exp1[:]    // a whole range
	fmt.Println(exp2) // [1, 2, 3, 4, 5, 6]

	exp2 = exp1[:3] // a whole range
	exp2[0] = 10
	fmt.Println(exp2) // [1, 2, 3]
	fmt.Println(exp1) // [10, 2, 3, 4, 5, 6]

	exp1 = []int{1, 2, 3, 4, 5, 6}
	exp2 = exp1[:3]
	exp2 = append(exp2, 15)
	fmt.Println(exp1) // [1, 2, 3, 15, 5, 6]
	fmt.Println(exp2) // [1, 2, 3, 15] // Keep the same array

	exp1 = []int{1, 2, 3, 4, 5, 6}
	exp2 = exp1[2:]
	exp2[0] = 100           // This operation update the original array
	exp2 = append(exp2, 15) // This operation create a new array
	fmt.Println(exp1)       // [1, 2, 100, 15, 5, 6]
	fmt.Println(exp2)       // [100, 4, 5, 6, 15] // A new array

	exp1 = []int{1, 2, 3, 4, 5, 6}
	exp2 = exp1[:6:6]
	exp2 = append(exp2, 15)
	fmt.Println(exp1)
	fmt.Println(exp2)

	exp1 = []int{1, 2, 3, 4, 5, 6}
	exp2 = make([]int, 10)
	fmt.Println(exp1) // [1, 2, 3, 4 ,5 ,6]
	fmt.Println(exp2) // [0, 0, 0, 0, 0, 0]
	copy(exp2, exp1)  // Copy that possible
	fmt.Println(exp1) // [1, 2, 3, 4,5, 6]
	fmt.Println(exp2) // [1, 2, 3, 4, 5, 6, 0, 0, 0, 0]

	arr := [6]int{1, 2, 3, 4, 5, 6}
	exp2 = make([]int, 6)
	copy(exp2, arr[:]) // Copy all elements from an array
	fmt.Println(arr)   // [1, 2, 3, 4,5, 6]
	fmt.Println(exp2)  // [1, 2, 3, 4, 5, 6]
}
