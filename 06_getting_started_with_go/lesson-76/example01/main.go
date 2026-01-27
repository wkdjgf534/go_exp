package main

import "fmt"

func main() {
	// readChannel := make(<-chan int)
	// writeChannel := make(chan<- int)

	intChannel := make(chan int)
	go func() {
		intChannel <- 10
	}()

	readValue, ok := <-intChannel
	fmt.Println(ok)
	fmt.Println(readValue)

	close(intChannel)

	readValue, ok = <-intChannel
	fmt.Println(ok)
	fmt.Println(readValue)
}
