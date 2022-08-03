package main

import (
  "log"
  "net/http"
)

var content = []byte(`hello world`)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write(content)
  })

  log.Fatal(http.ListenAndServe(":7777", nil))
}
