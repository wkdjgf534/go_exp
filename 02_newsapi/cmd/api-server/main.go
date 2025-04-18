package main

import (
	"log/slog"
	"net/http"
	"os"

	"newsapi/internal/logger"
	"newsapi/internal/router"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	r := router.New()
	wrappedRouter := logger.AddLoggerMid(log, logger.LoggerMid(r))

	log.Info("server starting on port 8080")
	if err := http.ListenAndServe(":8080", wrappedRouter); err != nil {
		log.Error("failed to start server", "error", err)
	}
}
