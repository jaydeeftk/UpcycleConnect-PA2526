package handlers

import (
  "net/http"

)

func Health(w http.ResponseWriter, r *http.Request) {
  w.Header() .Set("Content-Type", "application/json")
  r.WriteHeader(http.StatusOK)
  _, _ = w.Write([]byte(`{"status":"ok"}`))
}
