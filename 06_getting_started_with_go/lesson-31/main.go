package main

import "fmt"

func main() {
	a := 100

	for {
		fmt.Println(a)
		a--
		if a == 90 {
			break
		}
	}

	for i := range 5 {
		if i == 2 {
			continue
		}

		fmt.Println(i)
	}

outerLoop:
	for i := range 3 {
		for j := range 3 {
			if i == 1 && j == 1 {
				fmt.Println("skipped")
				continue outerLoop
			}
			fmt.Println(i, j)
		}
	}

	b := 10

	if b == 10 {
		goto end // move to on label end
	}
	fmt.Println(b) // Never print this text
end:
	fmt.Println("The end of the program.")
}
