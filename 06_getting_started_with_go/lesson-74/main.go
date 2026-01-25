package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var initialized bool

func initSomething() {
	fmt.Println("Initializing...")
	initialized = true
}

func doSomething(wg *sync.WaitGroup) {
	once.Do(initSomething)
	fmt.Println("Doing something...")
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)

	for range 3 {
		go doSomething(wg)
	}
	wg.Wait()

	// Initializing...
	// Doing something...
	// Doing something...
	// Doing something...
}
