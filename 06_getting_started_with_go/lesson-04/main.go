package main

import "fmt"

func main() {
	var smallFloat float32
	fmt.Println(smallFloat)
	smallFloat = 23.023433232
	fmt.Println(smallFloat) // 23.023434 Truncase the rest of the numbers

	var bigFloat float64
	fmt.Println(bigFloat)
	bigFloat = 23.023433232
	fmt.Println(bigFloat) // 23.023433232 All numbers

	var myComplex complex128
	myComplex = complex(bigFloat, bigFloat)
	fmt.Println(myComplex)

	var myRealPart, myImaginaryPart float64
	myRealPart = real(myComplex)
	myImaginaryPart = imag(myComplex)
	fmt.Println(myRealPart, myImaginaryPart)
}
