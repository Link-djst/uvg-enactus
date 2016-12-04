package main

import (
  "net/http"
  "log"
  "os"
  "github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Gorilla!\n"))
}

func main() {
  port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT must be set. Running localhost:8000.")
    port = ":8000"
	}

  r := mux.NewRouter()
  // Routes consist of a path and a handler function.
  r.HandleFunc("/", YourHandler)

  // Bind to a port and pass our router in
  log.Fatal(http.ListenAndServe(":8000", r))
}
