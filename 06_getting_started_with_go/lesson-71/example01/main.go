package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go sayHello(&wg)
	go sayHello(&wg)

	wg.Wait()
	fmt.Println("Both goroutines completed")
}
