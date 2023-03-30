package http

import (
	"encoding/json"
	"net/http"
)

func InternalError(message string, w http.ResponseWriter, r *http.Request) {
	respondWithError(message, http.StatusInternalServerError, w, r)
}

func Unauthorized(message string, w http.ResponseWriter, r *http.Request) {
	respondWithError(message, http.StatusUnauthorized, w, r)
}

func BadRequest(message string, w http.ResponseWriter, r *http.Request) {
	respondWithError(message, http.StatusBadRequest, w, r)
}

func NotFound(message string, w http.ResponseWriter, r *http.Request) {
	respondWithError(message, http.StatusNotFound, w, r)
}

func respondWithError(message string, status int, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(status)
	payload, _ := json.Marshal(map[string]string{"message": message})
	_, _ = w.Write(payload)
}
