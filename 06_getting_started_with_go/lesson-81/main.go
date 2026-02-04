package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Why fain-in and fan-out?

// fan-out: distribution of work
// Starting multiple goroutines to distribute work.

// fan-in: combining results
// Merging results from multiple goroutines into a channel.

// When to use fan-out and fan-in?
// Independence: stage should be independent of previous stages in terms of data.
// Computational intensity: stage should be coputationally expensive or time consuming.
// Consierations
// Ordering:  order of results may not be gauranteed.
// Error handling: handle error appropriately to prevent unexpected behaviour.

type result struct {
	url    string
	exists bool
}

func checkIfExists(done <-chan struct{}, urls <-chan string) <-chan result {
	resultsc := make(chan result)

	go func() {
		defer close(resultsc)

		for {
			select {
			case <-done:
				return
			case url, ok := <-urls:
				if !ok {
					return
				}
				res, err := http.Get(url)
				if err != nil {
					resultsc <- result{url: url, exists: false}
				} else if res.StatusCode == http.StatusOK {
					resultsc <- result{url: url, exists: true}
				} else {
					resultsc <- result{url: url, exists: false}
				}
			}
		}
	}()

	return resultsc
}

func merge[T any](done <-chan struct{}, channels ...<-chan T) <-chan T {
	results := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, c := range channels {
		go func(c <-chan T) {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				case i, ok := <-c:
					if !ok {
						return
					}
					results <- i
				}
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

func main() {
	done := make(chan struct{})
	defer close(done)

	urls := make(chan string, 4)

	urls <- "https://google.com"
	urls <- "https://amazon.com"
	urls <- "https://in-valid-url.invalid"
	urls <- "https://facebook.com"
	close(urls)

	c1 := checkIfExists(done, urls)
	c2 := checkIfExists(done, urls)
	c3 := checkIfExists(done, urls)

	now := time.Now()
	for result := range merge(done, c1, c2, c3) {
		fmt.Printf("url: %s, exists: %v\n", result.url, result.exists)
	}
	fmt.Println(time.Since(now))
}
