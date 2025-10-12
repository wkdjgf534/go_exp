package router

import (
	"net/http"

	"newsapi/internal/handler"
)

// New creates a new router with all the handlers configured.
func New(ns handler.NewsStorer) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("POST /news", handler.PostNews(ns))
	r.HandleFunc("GET /news", handler.GetAllNews(ns))
	r.HandleFunc("GET /news/{news_id}", handler.GetNewsByID(ns))
	r.HandleFunc("PUT /news/{news_id}", handler.UpdateNewsByID(ns))
	r.HandleFunc("DELETE /news/{news_id}", handler.DeleteNewsByID(ns))

	return r
}
