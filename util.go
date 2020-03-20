package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func dbContext(i time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), i*time.Second)
	return ctx, cancel
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	writeResponse(w, statusCode, map[string]string{"message": message})
}

func writeResponse(w http.ResponseWriter, statusCode int, v interface{}) {
	response, _ := json.Marshal(v)

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
