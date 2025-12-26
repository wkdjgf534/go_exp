package main

import "fmt"

// 1. Accessing byte data of a string.
// 2. Trying to update byte data of string.
// 3. Coverting byte, rune to string.
// 4. Slicing a string.
func main() {
	name := "Code & String"
	firstLetter := name[0]
	fmt.Println(firstLetter) // 67
	lastLetter := name[len(name)-1]
	fmt.Println(lastLetter) // 103

	// can't modify the string
	// name[0] = 'q'

	fmt.Println(string(firstLetter))
	fmt.Println(string(lastLetter))

	var character rune = 'x'
	// var character byte = 'x' - also works
	fmt.Println(string(character))

	// Convert string to bytes
	runeName := []rune(name)
	fmt.Println(runeName)

	byteName := []byte(name)
	fmt.Println(byteName)

	// Slicing string
	firstName := name[:4]
	apersand := name[4:6]
	fmt.Println(firstName)
	fmt.Println(apersand)

	emogi := "😜"
	firstCaracter := emogi[0]
	fmt.Println(string(firstCaracter)) // ð - first byte only, it is incorrect
	unicodeString := []rune(emogi)
	fmt.Println(string(unicodeString)) // 😜 - correct way
}
