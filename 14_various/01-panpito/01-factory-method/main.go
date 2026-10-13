package main

import "fmt"

// Example 1

type Person struct {
	age int
}

func NewPerson(yearOfBirth int) Person {
	return Person{age: 2026 - yearOfBirth}
}

// Example 2

// Interface
type Animal interface {
	Sound() string
}

// Cat
type Cat struct {
}

func (c Cat) Sound() string {
	return "meow"
}

func NewCat() Cat {
	return Cat{}
}

// Dog
type Dog struct {
	name string
}

func (d Dog) Sound() string {
	return "woof"
}

func NewDog() Dog {
	return Dog{name: "Doug"}
}

func farm(x int) {
	var a Animal

	if x > 42 {
		a = NewCat()
	} else {
		a = NewDog()
	}
	fmt.Println(a.Sound())
}

func main() {
	p := NewPerson(1996)
	fmt.Println(p.age)

	farm(52)
}
