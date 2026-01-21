package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	router.HandleFunc("POST /hello/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("user_id")

		fmt.Fprintf(w, "hello %s", userID)
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
	}
}
