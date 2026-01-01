package main

import "fmt"

func main() {
	sayHi()
	sayHiToSomeone("Peter")
	fn, fnLen := fullNameWithLength("Peter", "Peterson")
	fmt.Println(fn, fnLen)
}

func sayHi() {
	fmt.Println("Hi!")
}

func sayHiToSomeone(name string) {
	fmt.Println("Hi!", name)
}

func fullName(firstName, lastName string) string {
	return fmt.Sprintf("%s %s", firstName, lastName)
}

func fullNameWithLength(firstName, lastName string) (string, int) {
	fn := fullName(firstName, lastName)
	return fn, len(fn)
}
