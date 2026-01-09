package main

import "fmt"

type Calculator struct {
}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func AddFunction(a, b int) int {
	return a + b
}

func ArithmeticalOperation(f func(int, int) int, a, b int) int {
	return f(a, b)
}

//--------------------------------------------//

type Printer struct {
}

func (p *Printer) Print(m string) {
	fmt.Println(m)
}

//--------------------------------------------//

type Person struct {
	Name string
	Age  int
}

func (p Person) GetDetails() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func main() {
	c := Calculator{}
	fmt.Println(ArithmeticalOperation(c.Add, 10, 15))
	fmt.Println(ArithmeticalOperation(AddFunction, 10, 15))

	printer := &Printer{}
	printFunction := printer.Print
	printFunction("Hellow World!")

	f1 := Person.GetDetails
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(f1(p))
}
