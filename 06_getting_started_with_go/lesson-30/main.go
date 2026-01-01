package main

import "fmt"

func main() {
	a := 10

	for {
		fmt.Println(a)
		a--
		if a == 0 {
			break
		}
	}

	// old variant
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// new variant
	for i := range 10 {
		fmt.Println(i)
	}

	b := []int{1, 2, 100, 4, 5}

	for index, value := range b {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	str := "大五码"

	for index, runeValue := range str {
		fmt.Printf("Index: %d, Rune Value: %c\n", index, runeValue)
	}

	ages := map[string]int{
		"Alice": 30,
		"Bob":   35,
	}

	for name, age := range ages {
		fmt.Printf("%s is %d years old.\n", name, age)
	}
}
