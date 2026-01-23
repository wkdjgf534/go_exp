package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumber(wg *sync.WaitGroup, number int) {
	defer wg.Done()
	fmt.Println("Printing number:", number)
	time.Sleep(1 * time.Second)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go printNumber(&wg, i)
	}

	wg.Wait()

	fmt.Println("All numbers printed")
}
