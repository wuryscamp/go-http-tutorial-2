package main

import (
  "fmt"
  "net/http"
  "log"
)


func main() {

  h := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Hello Golang,,,,,,,")
  })

  log.Println("Server running on port 3000")

  err := http.ListenAndServe(":3000", h)

  if err != nil{
    log.Fatal(err)
  }


}
