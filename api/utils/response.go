package utils

import (
	"encoding/json"
	"net/http"
)


// ErrorResponse represents an error response format
// @Description Standard error response
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid data input"`
}

// SuccessResponse is a generic successful response
// @Description Generic success response (any JSON object)
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

func SendError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func SendSucces(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}