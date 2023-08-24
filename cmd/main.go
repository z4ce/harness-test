package main

import (
  "fmt"
  "net/http"
  "time"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "The current time is: %s", time.Now())
  })
  http.ListenAndServe(":8080", nil)
}
