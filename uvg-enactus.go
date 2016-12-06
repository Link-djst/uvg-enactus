package main

import(
  "io"
  "os"
  "net/http"
)

func main() {
  port := os.Getenv("PORT")

  http.HandleFunc("/", hello)
  http.ListenAndServe(":" + port, nil)
}

func hello(w http.ResponseWriter, r *http.Request){
  io.WriteString(w, "Hello Arch-Golang")
}
