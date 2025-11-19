package main

import (
	"fmt"
	"time"
)

func main() {
	go printMessage("done")
	go printMessage("done1")
	go printMessage("done2")
	go printMessage("done3")

	// main goroutine
	printMessage("done4")

	go func() {
		fmt.Println("additional goroutine")
	}()

	time.Sleep(time.Second)
}

func printMessage(msg string) {
	for range 5 {
		fmt.Println(msg)
	}
}
