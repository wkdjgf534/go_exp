package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"newsapi/internal/logger"
	"newsapi/internal/store"
)

type NewsStorer interface {
	Create(store.News) (store.News, error)
	FindByID(uuid.UUID) (store.News, error)
	FindAll() ([]store.News, error)
	DeleteByID(uuid.UUID) error
	UpdateByID(store.News) error
}

func PostNews(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")

		var newsRequestBody NewsPostReqBody
		if err := json.NewDecoder(r.Body).Decode(&newsRequestBody); err != nil {
			logger.Error("failed to decode the request", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		n, err := newsRequestBody.Validate()
		if err != nil {
			logger.Error("request validation failed", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if _, err := ns.Create(n); err != nil {
			logger.Error("error creating news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetAllNews(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")
		news, err := ns.FindAll()
		if err != nil {
			logger.Error("failed to fetch all news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		allNewsResponse := allNewsResponse{News: news}
		if err := json.NewEncoder(w).Encode(allNewsResponse); err != nil {
			logger.Error("failed to write response", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func GetNewsByID(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")
		newsID := r.PathValue("news_id")
		newsUUID, err := uuid.Parse(newsID)
		if err != nil {
			logger.Error("news id not a valid uuid", "newsId", newsID, "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		news, err := ns.FindByID(newsUUID)
		if err != nil {
			logger.Error("news not found", "newsId", newsID)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(&news); err != nil {
			logger.Error("failed to encode", "newsId", newsID, "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func UpdateNewsByID(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")

		var newsRequestBody NewsPostReqBody
		if err := json.NewDecoder(r.Body).Decode(&newsRequestBody); err != nil {
			logger.Error("failed to decode the request", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		n, err := newsRequestBody.Validate()
		if err != nil {
			logger.Error("request validation failed", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if err := ns.UpdateByID(n); err != nil {
			logger.Error("error updating news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func DeleteNewsByID(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())

		newsID := r.PathValue("news_id")
		newsUUID, err := uuid.Parse(newsID)
		if err != nil {
			logger.Error("news id not a valid uuid", "newsId", newsID, "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := ns.DeleteByID(newsUUID); err != nil {
			logger.Error("news not found", "newsId", newsID, "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
