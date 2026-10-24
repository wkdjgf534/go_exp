package main

import (
	"fmt"
)

type Object interface {
	Perimeter() float64
	Height() float64
}

type VolumeCalculator struct {
	object Object
}

func (calculator VolumeCalculator) Volume() float64 {
	return calculator.object.Perimeter() * calculator.object.Height()
}

// Implementation of an Object: triangular
type TriangularObject struct {
	height float64
	base   float64
	lenght float64
}

func (t TriangularObject) Perimeter() float64 { return t.height * t.base / 2 }
func (t TriangularObject) Height() float64    { return t.lenght }

// Implementation of an Object: rectangular
type RectangularObject struct {
	x float64
	y float64
	z float64
}

func (r RectangularObject) Perimeter() float64 { return r.x * r.y }
func (r RectangularObject) Height() float64    { return r.z }

func main() {
	triangle := TriangularObject{
		height: 1,
		base:   2,
		lenght: 3,
	}

	rectangle := RectangularObject{
		x: 23,
		y: 4,
		z: 6,
	}

	x := 3
	var object Object
	if x > 2 {
		object = triangle
	} else {
		object = rectangle
	}

	VolumeCalculator := VolumeCalculator{
		object: object,
	}

	volume := VolumeCalculator.Volume()
	fmt.Println(volume)
}
