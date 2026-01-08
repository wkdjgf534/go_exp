package main

import "fmt"

type Rectangle struct {
	Length float64
	Width  float64
}

// value receiver
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// pointer receiver
func (r *Rectangle) SetLength(l float64) {
	r.Length = l
}

func main() {
	rect := Rectangle{Length: 5.0, Width: 3.0}
	fmt.Println(rect.Area())

	rect.SetLength(10)
	fmt.Println(rect.Area())
}
