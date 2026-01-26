package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sender(ch)
	go sender(ch)
	go sender(ch)

	go receiver(ch)
	go receiver(ch)
	go receiver(ch)

	fmt.Scanln()
}

// send-only channel
func sender(ch chan<- string) {
	for i := range 5 {
		ch <- fmt.Sprintf("message%d", i)
	}

}

// receive-only channel
func receiver(ch <-chan string) {
	for msg := range ch {
		time.Sleep(time.Second)
		fmt.Println(msg)
	}
}
