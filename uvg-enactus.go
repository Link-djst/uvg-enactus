package main

import(
  "io"
  "os"
  "net/http"

  "github.com/gorilla/mux"
)

func main() {
  port := os.Getenv("PORT")

  r := mux.NewRouter()
  r.HandleFunc("/", hello)

  http.ListenAndServe(":" + port, r)
}

func hello(w http.ResponseWriter, r *http.Request){
  io.WriteString(w, "Hello Arch-Golang")
}
