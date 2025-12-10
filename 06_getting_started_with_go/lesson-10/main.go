package main

import "fmt"

func main() {
	a := 10
	b := 5
	c := 15.50

	complex1 := complex(10, 15)
	complex2 := complex(5, 10)

	// Sum operator
	sum := a + b
	fmt.Println("Sum: ", sum)

	firstName := "Peter"
	lastName := "Peterson"
	fmt.Printf("%s %s\n", firstName, lastName)

	complexSum := complex1 + complex2
	fmt.Println("Complex sum: ", complexSum)

	// Difference operator
	difference := a - b
	fmt.Println("Diffenrece: ", difference)

	complexDif := complex1 - complex2
	fmt.Println("Complex dif: ", complexDif)

	// Product operator
	product := a * b
	fmt.Println("Product: ", product)

	complexProd := complex1 * complex2
	fmt.Println("Complex prod: ", complexProd)

	// Division operator
	division := a / b
	fmt.Println("Division: ", division)

	complexDiv := complex1 / complex2
	fmt.Println("Complex div: ", complexDiv)

	// Remainder operator
	remainder := a % b
	fmt.Println("Remainder ", remainder)

	newSum := a + int(c) // float part is dropped by int func
	fmt.Println(newSum)

	// Shorthand computation
	d := 12

	d += 5 // d = d + 5
	fmt.Println("Shorthand sum: ", d)

	d -= 5 // d = d - 5
	fmt.Println("Shorthand Dif: ", d)

	d *= 5 // d = d * 5
	fmt.Println("Shorthand prod: ", d)

	d /= 5 // d = d / 5
	fmt.Println("Shorthand div: ", d)

	d %= 5 // d = d % 5
	fmt.Println("Shorthand rem: ", d)

	d++
	fmt.Println("Inc", d)
	d--
	fmt.Println("Dec", d)
}
