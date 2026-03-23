package main

import "fmt"

func main() {
	process(10)
}

func process(i int) {
	defer fmt.Println("Deffered i value:", i)               // Deffered i value: 10
	defer fmt.Println("First deferred statement executed")  // 4
	defer fmt.Println("Second deferred statement executed") // 3
	defer fmt.Println("Third deferred statement executed")  // 2
	i++
	fmt.Println("Normal execution statement") // 1
	fmt.Println("Value of i:", i)             // Value of i 11
}
