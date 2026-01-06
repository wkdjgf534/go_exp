package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Write a function that takes an arbitrary number of input channels and returns a
// single output channel, into which all values from the provided channels will be
// sent. Once all input channels are closed, the output channel should also be
// closed. The function must be non-blocking.

// See fan-in pattern
func mergeNChannels(channels ...chan string) chan string {
	out := make(chan string)
	wg := sync.WaitGroup{}

	for ch := range channels {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()
			for val := range channels[c] {
				out <- val // Forward value to output channel
			}
		}(ch)
	}

	// Start a goroutine to close the output channel when all input channels are done
	go func() {
		wg.Wait()  // Wait for all goroutines to finish
		close(out) // Close output channel
	}()

	return out // Return output channel immediately (non-blocking)
}

func main() {
	nChannels := 2
	nMessages := 5

	channels := make([]chan string, nChannels)
	for i := range channels {
		channels[i] = make(chan string)
		go func(i int) {
			for j := range nMessages {
				milliseconds := rand.Intn(2000)
				time.Sleep(time.Duration(milliseconds) * time.Millisecond)
				channels[i] <- fmt.Sprintf("[%02d:%02d:%02d] Channel %d: %d",
					time.Now().Hour(), time.Now().Minute(), time.Now().Second(), i, j,
				)
			}
			close(channels[i])
		}(i)
	}

	out := mergeNChannels(channels...)
	for val := range out {
		fmt.Println(val)
	}
}
