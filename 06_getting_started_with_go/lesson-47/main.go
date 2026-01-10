package main

import "fmt"

type Shape interface {
	Area() float64
}

type Object interface {
	Name() string
	Shape
}

type Reactangle struct {
	Width  float64
	Height float64
}

func (r Reactangle) Area() float64 {
	return r.Width * r.Height
}

func (r Reactangle) Name() string {
	return "Rectangle"
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Square struct {
	Width float64
}

func (s Square) Area() float64 {
	return s.Width * s.Width
}

func printArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}

func printObject(o Object) {
	fmt.Printf("Area: %f, Name: %s", o.Area(), o.Name())
}

func main() {
	rectangle := Reactangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}
	square := Square{Width: 10}

	shapes := []Shape{rectangle, circle, square}

	// Calculate and print the area of each shape
	for _, shape := range shapes {
		printArea(shape)
	}

	printObject(rectangle)
}
