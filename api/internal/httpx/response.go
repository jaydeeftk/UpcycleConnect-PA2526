package httpx

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

// JSON writes a JSON response with a given status code.
func JSON(w http.ResponseWriter, status int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(payload)
}

// JSONOK writes a success response.
func JSONOK(w http.ResponseWriter, status int, data interface{}) {
	JSON(w, status, APIResponse{
		Success: true,
		Data:    data,
		Error:   nil,
	})
}

// JSONError writes an error response.
func JSONError(w http.ResponseWriter, status int, message string) {
	JSON(w, status, APIResponse{
		Success: false,
		Data:    nil,
		Error:   message,
	})
}
