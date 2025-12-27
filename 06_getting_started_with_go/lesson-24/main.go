package main

import "fmt"

// 1. A nil map.
// 2. A empty map.
// 3. Writing to a map.
// 4. Reading from map.
func main() {
	// Creation of a nil map
	var nameAge map[string]int
	fmt.Println(len(nameAge))   // 0
	fmt.Println(nameAge["foo"]) // 0

	// nameAge["foo"] = 21 // panic: assignment to entry in nil map

	// Varialbe shorthand
	nameAge1 := map[string]int{}
	fmt.Println(len(nameAge1)) // 0

	// Make function
	nameAge2 := make(map[string]int)
	fmt.Println(len(nameAge2)) // 0

	// Literal variable declaration.
	var nameAge3 map[string]int = map[string]int{}
	fmt.Println(len(nameAge3)) // 0

	nameAge4 := map[string]int{
		"a": 24,
		"b": 30,
		"c": 60,
	}

	nameAge4["foo"] = 25
	nameAge4["bar"] = 30
	nameAge4["foo bar"] = 45
	fmt.Println(len(nameAge4)) // 6

	// Reading from map.
	fmt.Println(nameAge4["c"])       // 60
	fmt.Println(nameAge4["foo bar"]) // 45
	fmt.Println(nameAge4["x"])       // 0 the key "x", does not exist in map, it returns a default value for int
	nameAge4["c"] = 70               // Somekind of the Set Class, keys in map stay unique.
	fmt.Println(nameAge4["c"])

	// comma ok idiom
	nameGrade := map[string]int{
		"foo":    10,
		"bar":    9,
		"foobar": 8,
		"x":      0,
	}

	gradeX, ok := nameGrade["y"]

	fmt.Println(gradeX, ok) // 0, false

	a := map[string][]int{
		"foo": {1, 2, 3, 4},
	}

	fmt.Println(a)
}
