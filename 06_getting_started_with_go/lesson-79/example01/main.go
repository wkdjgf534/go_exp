package main

import (
	"fmt"
	"net/http"
)

type result struct {
	resp *http.Response
	err  error
}

func checkIfExist(done <-chan struct{}, urls ...string) <-chan result {
	results := make(chan result)

	go func() {
		for _, url := range urls {
			select {
			case <-done:
				return
			default:
				res, err := http.Get(url)
				results <- result{resp: res, err: err}
			}
		}
		close(results)
	}()

	return results
}

func main() {
	done := make(chan struct{})

	results := checkIfExist(done, "https://google.com", "http:localhost:3000")

	for r := range results {
		if r.resp != nil {
			fmt.Println(r.resp.Status)
		}
		if r.err != nil {
			fmt.Println(r.err)
		}
	}

	close(done)
}

// 200 OK
// Get "http:localhost:3000": http: no Host in request URL
