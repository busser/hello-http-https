package main

import (
  "log"
  "net/http"
  "sync"
)

func main() {
  var wg sync.WaitGroup
  wg.Add(2)

  helloHandler := func(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello, World!\n"))
  }

  http.HandleFunc("/hello-world", helloHandler)

  go func() {
    defer wg.Done()
    err := http.ListenAndServe(":80", nil)
    if err != nil {
      log.Fatal("Could not listen on port 80: ", err)
    }
  }()

  go func() {
    defer wg.Done()
    err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
    if err != nil {
      log.Fatal("Could not listen on port 443: ", err)
    }
  }()

  wg.Wait()
}
