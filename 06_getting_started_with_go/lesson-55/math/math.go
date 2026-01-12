package math

// PI exported
const PI = 3.14159

// Add exported
func Add(x, y int) int {
	return x + y
}

// subtract unexported
func subtract(x, y int) int {
	return x - y
}
