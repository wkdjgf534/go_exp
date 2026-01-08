package main

import "fmt"

type Rectangle struct {
	Length float64
	Width  float64
}

// value receiver
func (r *Rectangle) Area() float64 {
	if r == nil {
		return 0
	}
	return r.Length * r.Width
}

// pointer receiver
func (r *Rectangle) SetLength(l float64) {
	r.Length = l
}

// function
func updateLength(r *Rectangle, l float64) {
	r.SetLength(l)
}

func main() {
	rect := Rectangle{Length: 5, Width: 5}
	fmt.Println(rect.Area())
	updateLength(&rect, 10)
	fmt.Println(rect.Area())
}
