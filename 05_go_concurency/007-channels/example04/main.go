package main

import "fmt"

func main() {
	ch := make(chan string, 1)

	ch <- "message" // send

	// ch <- "message1" // the channel is full, deadlock

	fmt.Println(<-ch) // read

	// fmt.Println(<-ch) // the channel is empty, deadlock
}
