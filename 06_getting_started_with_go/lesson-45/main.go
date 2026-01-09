package main

import "fmt"

type Age int

type Human struct {
	Age int
}

type Student Human

func (h Human) printAge() {
	fmt.Println(h.Age)
}

type Size int

const (
	ExtraSmall Size = iota // 0
	Small
	Medium
	Large
	ExtraLarge
)

func printSize(s Size) {
	switch s {
	case ExtraSmall:
		fmt.Println("Extra Small")
	case Small:
		fmt.Println("Small")
	case Medium:
		fmt.Println("Medium")
	case Large:
		fmt.Println("Large")
	case ExtraLarge:
		fmt.Println("Extra Large")
	}
}

func main() {
	var young Age = 10
	var old Age = 60
	fmt.Println(young + old) // 70

	s := Student{Age: 10}
	fmt.Println(s.Age) // 10

	fmt.Println(ExtraSmall, Small, Medium, Large, ExtraLarge) // 0 1 2 3 4
	printSize(Medium)                                         // Medium
}
