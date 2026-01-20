package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time.
	now := time.Now()
	fmt.Println(now)

	// Calculate time 2 hours from now
	twoHours := 2 * time.Hour
	futureTime := now.Add(twoHours)
	fmt.Println("Time 2 hours from now:", futureTime)

	// Measure the execution time of a code block.
	start := time.Now()
	a := 1
	for i := 0; i < 1000000000; i++ {
		a++
	}
	elapsed := time.Since(start)
	fmt.Println("Loop execution time:", elapsed)
}
