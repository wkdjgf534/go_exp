package goasteroids

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v Vector) Normalize() Vector {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)
	if magnitude == 0 {
		return Vector{0, 0}
	}

	return Vector{v.X / magnitude, v.Y / magnitude}
}
