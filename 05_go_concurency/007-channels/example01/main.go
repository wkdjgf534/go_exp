package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	// sender
	go func() {
		time.Sleep(time.Second)
		ch <- "message123"
	}()

	// receiver
	go func() {
		message := <-ch
		fmt.Printf("message: %v\n", message)
	}()

	fmt.Scanln()
}
