package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	}))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
