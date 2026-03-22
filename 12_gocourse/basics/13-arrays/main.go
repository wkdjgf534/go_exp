package main

import "fmt"

func main() {
	// var arrayName [size]elementType

	// var numbers [5]int
	// fmt.Println(numbers) // [0 0 0 0 0]

	// numbers[4] = 20
	// fmt.Println(numbers) // [0 0 0 0 20]

	// numbers[0] = 9
	// fmt.Println(numbers) // [9 0 0 0 20]

	// fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	// fmt.Println("Fruits array:", fruits)

	// fmt.Println("Third element:", fruits[2])

	// orignalArray := [3]int{1, 2, 3}
	// copiedArray := orignalArray

	// copiedArray[0] = 100

	// fmt.Println("Original array", orignalArray) // [1 2 3]
	// fmt.Println("Copied array", copiedArray)    // [100 2 3]

	// for i, v := range numbers {
	// 	fmt.Printf("Element: %d, Value: %d\n", i, v)
	// }

	// // underscore is blank identifier, used to store unused values
	// for _, v := range numbers {
	// 	fmt.Printf("Value: %d\n", v)
	// }

	// fmt.Println("The length of numbers array is", len(numbers))

	// // Comparing Arrays
	// array1 := [3]int{1, 2, 3}
	// array2 := [3]int{1, 2, 5}
	// fmt.Println("Array1 is equal Array2:", array1 == array2)

	// var matrix [3][3]int = [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	// fmt.Println(matrix) // [[1 2 3] [4 5 6] [7 8 9]]

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100

	fmt.Println("Original array:", originalArray)
	fmt.Println("Copied array:", *copiedArray)
}
