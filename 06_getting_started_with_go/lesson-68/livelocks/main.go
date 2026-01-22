package main

import (
	"fmt"
	"sync"
)

var (
	mutexA sync.Mutex
	mutexB sync.Mutex
)

func routineA(name string) {
	mutexA.Lock()
	defer mutexA.Unlock()
	fmt.Println(name, "has mutex A")
	fmt.Println(name, "trying to acquire mutex B...")
	if !mutexB.TryLock() {
		fmt.Println(name, "cannot acquire mutex B. Backing off.")
		// Simulate some backoff time
		for range 100 {
		}
		routineA(name) // Recursive call to simulate retry (core of the livelock)
	} else {
		defer mutexB.Unlock()
		fmt.Println(name, "acquired mutex B")
		fmt.Println(name, "doing some work...")
	}
}

func routineB(name string) {
	mutexB.Lock()
	defer mutexB.Unlock()
	fmt.Println(name, "has mutex B")
	fmt.Println(name, "trying to acquire mutex A...")
	if !mutexA.TryLock() {
		fmt.Println(name, "cannot acquire mutex A. Backing off.")
		// Simulate some backoff time
		for range 100 {
		}
		routineB(name) // Recursive call to simulate retry (core of the livelock)
	} else {
		defer mutexA.Unlock()
		fmt.Println(name, "acquire mutex A")
		fmt.Println(name, "doing some work...")
	}
}

func main() {
	go routineB("Routine B")
	go routineA("Routine A")

	// Wait for routines (although they might not finish due to the livelock)
	fmt.Println("Waiting...")
	select {}

	// Waiting...
	// Routine A has mutex A
	// Routine B has mutex B
	// Routine B trying to acquire mutex A...
	// Routine A trying to acquire mutex B...
	// Routine A cannot acquire mutex B. Backing off.
	// Routine B cannot acquire mutex A. Backing off.
	// fatal error: all goroutines are asleep - deadlock!
}
