package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 10, 12}
	s := sum(nums...)
	fmt.Println(s, "a", "x", 10) // Println also accepts variadic parameters
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}
