package main

import "net/http"

type MyHandler struct{}

func (MyHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello world!\n"))
}

func main() {
	err := http.ListenAndServe(":3000", MyHandler{})
	if err != nil {
		panic(err)
	}
}
