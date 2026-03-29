package main

import (
	"fmt"
	"slices"
)

func main() {
	// var sliceName []ElementType

	// var numbers []int
	// var numbers1 = []int{1, 2, 3}
	// numbers2 := []int{9, 8, 7}
	// slice := make([]int, 5)

	a := [5]int{1, 2, 3, 4, 5}
	slice := a[1:4]
	fmt.Println(slice) // [2 3 4]

	slice = append(slice, 6, 7)
	fmt.Println("Slice:", slice) // [2 3 4 6 7]

	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	fmt.Println("SliceCopy:", sliceCopy) // [2 3 4 6 7]

	var nilSlice []int
	fmt.Println(nilSlice) // []

	for i, v := range slice {
		fmt.Println("Index:", i, "Value:", v)
	}
	fmt.Println("Element at index 3 of slice:", slice[3]) // 6

	// slice[3] = 50
	// fmt.Println("Element at index 3 of slice:", slice[3]) // 50

	if slices.Equal(slice, sliceCopy) {
		fmt.Println("slice is equal to sliceCopy")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
			fmt.Printf("Adding value %d in outer slice at index %d, and in inner slice index of %d\n", i+j, i, j)
		}
	}

	fmt.Println(twoD)

	// slice[low:high]
	slice2 := slice[2:4] // the last is excluded from the range
	fmt.Println(slice2)  // [4 6]

	fmt.Println("The capacity of slice2 is", cap(slice2)) // 6
	fmt.Println("The length of slice2 is", len(slice2))   // 2

}
