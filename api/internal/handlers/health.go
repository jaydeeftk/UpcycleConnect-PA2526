package handlers

import (
	"net/http"

	"upcycleconnect/internal/httpx"
)

func Health(w http.ResponseWriter, r *http.Request) {
	httpx.JSONOK(w, http.StatusOK, map[string]string{
		"service": "upcycleconnect-api",
		"status":  "healthy",
	})
}
