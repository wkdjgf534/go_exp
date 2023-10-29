package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/hello/", hello)
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Home\n"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.Path, "/")[2]
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}
