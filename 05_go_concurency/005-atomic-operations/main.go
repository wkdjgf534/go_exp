package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// first example
//func main() {
//	var count int32
//	var wg sync.WaitGroup
//
//	for range 1000 {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			atomic.AddInt32(&count, 1)
//		}()
//
//	}
//
//	wg.Wait()
//	fmt.Println(atomic.LoadInt32(&count))
//}

type Person struct {
	Name string
	Age  int32
}

func main() {
	var wg sync.WaitGroup
	var person atomic.Value

	person.Store(&Person{Name: "Peter", Age: 40})

	for range 10 {
		wg.Go(func() { // wg.Go - new format for waitgroups go version >= 1.25
			p := person.Load().(*Person)
			atomic.AddInt32(&p.Age, 1)
		})

	}

	wg.Wait()
	fmt.Println(person.Load().(*Person))
}
