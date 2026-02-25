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

func JSON(w http.ResponseWriter, status int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func JSONOK(w http.ResponseWriter, status int, data interface{}) {
	JSON(w, status, APIResponse{
		Success: true,
		Data:    data,
		Error:   nil,
	})
}

func JSONError(w http.ResponseWriter, status int, message string) {
	JSON(w, status, APIResponse{
		Success: false,
		Data:    nil,
		Error:   message,
	})
}
