package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)

	// sending
	go func() {
		ch1 <- "message to channel 1"
	}()

	go func() {
		ch2 <- "message to channel 2"
	}()

	go func() {
		ch3 <- "message to channel 3"
	}()

	go func() {
		ch4 <- "message to channel 4"
	}()

	// receiving
	timeout := time.After(5 * time.Second)
	go func() {
		for {
			select {
			case v := <-ch1:
				fmt.Printf("message on ch1: %v\n", v)
			case v := <-ch2:
				fmt.Printf("message on ch2: %v\n", v)
			case v := <-ch3:
				fmt.Printf("message on ch3: %v\n", v)
			case v := <-ch4:
				fmt.Printf("message on ch4: %v\n", v)
			case <-timeout:
				fmt.Println("five seconds without messages, panic")
				panic("no messages")
			default:
				time.Sleep(time.Second)
				fmt.Println("waiting")
			}
		}
	}()

	fmt.Scanln()
}
