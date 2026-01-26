package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		done <- true
	}()

	<-done

	fmt.Println("main thread unblocked")
}
