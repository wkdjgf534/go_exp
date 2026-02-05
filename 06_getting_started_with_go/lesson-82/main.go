package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type result struct {
	url    string
	exists bool
}

func checkIfExists(ctx context.Context, urls <-chan string) <-chan result {
	resultsc := make(chan result)

	go func() {
		defer close(resultsc)

		for {
			select {
			case <-ctx.Done():
				err := ctx.Err()
				fmt.Println(err)
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

func run(ctx context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	urls := make(chan string, 4)

	urls <- "https://google.com"
	urls <- "https://amazon.com"
	urls <- "https://in-valid-url.invalid"
	urls <- "https://facebook.com"
	close(urls)

	c := checkIfExists(ctxWithTimeout, urls)

	now := time.Now()
	for result := range c {
		fmt.Printf("url: %s, exists: %v\n", result.url, result.exists)
	}
	fmt.Println(time.Since(now))
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	run(ctx)
}
