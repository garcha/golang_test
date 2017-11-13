package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  //"github.com/joho/godotenv"
)

func newRouter() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/hello", handler).Methods("GET")

  staticFileDirectory := http.Dir("./assets/")
  staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
  r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

  r.HandleFunc("/bird", getBirdHandler).Methods("GET")
  r.HandleFunc("/bird", createBirdHandler).Methods("POST")
  return r
}

func main() {
  // http.HandleFunc("/", handler)
  r := newRouter()
  http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello World!")
}
