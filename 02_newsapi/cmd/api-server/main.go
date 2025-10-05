package main

import (
	"log/slog"
	"net/http"
	"newsapi/internal/router"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	logger.Info("server starting on port 8080")

	r := router.New()
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("failed to start server", "error", err)
	}

}
