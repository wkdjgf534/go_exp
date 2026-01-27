package main

import "fmt"

func main() {
	intChannel := make(chan int)

	go func() {
		for i := range 5 {
			intChannel <- i
		}
		close(intChannel)
	}()

	for j := range intChannel {
		fmt.Println(j)
	}
}
