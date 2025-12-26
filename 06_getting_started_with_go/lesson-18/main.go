package main

import "fmt"

// 1. Creating a slice.
// 2. The make function.
// 3. The append function.
// 4. nil and empty slice.
func main() {
	// Nil slice
	var sl1 []int
	fmt.Println(sl1 == nil) // true

	// Emprt slice
	sl1 = []int{} // an empty slice, it is not nil anymore
	fmt.Println(sl1)
	fmt.Println(sl1 == nil) // false

	// Different way of declaring and initializing slices
	sl2 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl2) // [1, 2, 3, 4, 5]
	sl2 = []int{5, 3: 10, 50}
	fmt.Println(sl2) // [5 0 0 10 50]

	// Make function
	sl3 := make([]int, 5)
	fmt.Println(sl3)      // [0, 0, 0, 0, 0]
	fmt.Println(len(sl3)) // 5
	fmt.Println(cap(sl3)) // 5

	sl4 := make([]string, 5, 10)
	fmt.Println(sl4)      // [     ]
	fmt.Println(len(sl4)) // 5
	fmt.Println(cap(sl4)) // 10

	// Append function
	sl5 := make([]int, 0)
	sl5 = append(sl5, 10)
	fmt.Println(sl5)      // [10]
	fmt.Println(cap(sl5)) // 1
	sl5 = append(sl5, 10, 15, 20)
	fmt.Println(sl5)      // [10, 10, 15, 20]
	fmt.Println(cap(sl5)) // 4
	sl5 = append(sl5, 25, 30)
	fmt.Println(cap(sl5)) // 8

	sl6 := []int{1, 2, 4, 5}
	sl7 := []int{6, 7, 8}
	sl6 = append(sl6, sl7...)
	fmt.Println(sl6) // [1, 2, 3, 4, 5, 6, 7, 8]
}
