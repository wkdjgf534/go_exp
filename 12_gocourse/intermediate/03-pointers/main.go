package main

import "fmt"

func main() {
	var ptr *int
	var a int = 10
	ptr = &a

	fmt.Println(a)    // 10
	fmt.Println(&a)   // hexademical address to a
	fmt.Println(ptr)  // hexademical address to a
	fmt.Println(*ptr) // dereferencing a pointer - 10

	if ptr == nil {
		fmt.Println("Pointer is nil")
	}

	modifyValue(ptr)
	fmt.Println(a)

}

func modifyValue(ptr *int) {
	*ptr++ // it changes the a value by 1 - 11
}
