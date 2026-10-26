package main

import "fmt"

type human struct {
	age      int
	height   float32
	eyeColor string
}

// Empty constructor
func newHuman() human {
	return human{}
}

func newHumanWithFields(age int, height float32, eyeColor string) human {
	h := human{
		age:      age,
		height:   height,
		eyeColor: eyeColor,
	}

	return h
}

// Builders
func (h human) withAge(age int) human {
	h.age = age

	return h
}

func (h human) withHeight(height float32) human {
	h.height = height

	return h
}

func (h human) withEyeColour(color string) human {
	h.eyeColor = color

	return h
}

// Reset
func (h human) reset() human {
	h = newHuman()

	return h
}

// Pre-defined builders
func giant() human {
	return newHuman().withHeight(2.5).withEyeColour("green")
}

func main() {
	me := newHuman().withAge(33).withHeight(1.7).withEyeColour("blue")
	fmt.Printf("%+v\n", me)
	fmt.Printf("%+v\n", me.reset())

	you := newHumanWithFields(33, 1.7, "blue")
	fmt.Printf("%+v\n", you)

	youngGiant := giant().withAge(14)
	oldGiant := giant().withAge(44)
	fmt.Printf("%+v\n", youngGiant)
	fmt.Printf("%+v\n", oldGiant)
}
