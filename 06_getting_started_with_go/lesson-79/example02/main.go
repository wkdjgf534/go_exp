package main

import (
	"fmt"
	"net/http"
)

func checkIfExist(done <-chan struct{}, urls ...string) (<-chan *http.Response, <-chan error) {
	responsec := make(chan *http.Response)
	errc := make(chan error)

	go func() {
		for _, url := range urls {
			select {
			case <-done:
				return
			default:
				res, err := http.Get(url)
				if err != nil {
					errc <- err
					continue
				}
				responsec <- res
			}
		}
		close(responsec)
		close(errc)
	}()

	return responsec, errc
}

func main() {
	done := make(chan struct{})

	responses, errc := checkIfExist(done, "https://google.com", "http:localhost:3000")

	for range 2 {
		select {
		case res := <-responses:
			fmt.Println("Status: ", res.Status)
		case err := <-errc:
			fmt.Println("Error: ", err)
		}
	}

	close(done)
}

// 200 OK
// Get "http:localhost:3000": http: no Host in request URL
