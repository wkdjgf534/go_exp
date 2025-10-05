package router

import (
	"net/http"

	"newsapi/internal/handler"
)

// New creates a new router with all the handlers configured.
func New() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("POST /news", handler.PostNews())
	r.HandleFunc("GET /news", handler.GetAllNews())
	r.HandleFunc("GET /news/{news_id}", handler.GetNewsByID())
	r.HandleFunc("PUT /news/{news_id}", handler.UpdateNewsByID())
	r.HandleFunc("DELETE /news/{news_id}", handler.DeleteNewsByID())

	return r
}
