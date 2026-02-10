package main

import "fmt"

// pipelines

// The first stage
func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

// The second stage
func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func main() {
	c := gen(2, 3, 4, 5)
	out := sq(c)

	for n := range out {
		fmt.Println(n)
	}
}
