package main

import (
  "fmt"
  "net/http"
)

func main() {
  var h Hello
  err := http.ListenAndServe("localhost:4000", h)
  checkError(err)
}

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "<h1>Hello from Go web server!</h1>")
}

func checkError(err error) {

  if err != nil {
    panic(err)
  }
}
