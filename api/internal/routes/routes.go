packages routes

import (
  "net/http"

  "upcycleconnect/internal/handlers"

  )

func NewRouter () http.Handler {
  mux := http.NewServeMux ()

  mux.HandleFunc("/health", handlers.Health)

  return mux

  }
