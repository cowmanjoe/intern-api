package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

type Internship struct {
  ID        string
  Title     string
}

var internships []Internship

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/internships/{id}", GetInternship).Methods("GET")

  internships = append(internships, Internship{ID: "1", Title: "Amazon summer internship"})
  internships = append(internships, Internship{ID: "2", Title: "Facebook winter internship"})

  log.Fatal(http.ListenAndServe(":8000", router))
}

func GetInternship(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for _, item := range internships {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Internship{})
}