package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(message string, w http.ResponseWriter, statusCode int) {
	response := Response{message}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

type Response struct {
	Message string `json:"message"`
}
