package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw(
		"internal server error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf(
		"bad request error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf(
		"not found error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusNotFound, "not found")
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorf(
		"conflict response",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error(),
	)

	writeJSONError(w, http.StatusConflict, err.Error())
}
