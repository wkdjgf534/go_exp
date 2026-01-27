package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := range 3 {
			ch1 <- i
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := range 3 {
			ch2 <- i
			time.Sleep(1 * time.Second)
		}
	}()

	for range 8 {
		select {
		case value := <-ch1:
			fmt.Println("Received form ch1:", value)
		case value := <-ch2:
			fmt.Println("Received form ch2:", value)
		case <-time.After(2 * time.Second):
			fmt.Println("Returning after timeout")
			return
		}
	}
}
