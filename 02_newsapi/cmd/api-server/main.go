package main

import (
	"log"
	"net/http"

	"newsapi/internal/router"
)

func main() {
	r := router.New()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("failed to start server: ", err)
	}

}
