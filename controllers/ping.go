package controllers

import (
	"encoding/json"
	"net/http"
)

func Pong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(StandardResponse{
		true, "pong",
	})
}
