module upcycleconnect
package main

import (
  "log"
  "net/http"

  "upcycleconnect/internal/routes"

  )

func main () {
  mux := routes.NewRouter()

  addr := ":8080"
  log.Printf("API listening on %s", addr)

  if err := http.ListenAndServe(addr, mux); err != nil {
    log.Fatal(err)
    }
}
