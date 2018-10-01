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
  router.HandleFunc("/internships", GetInternships).Methods("GET")
  router.HandleFunc("/internships/{id}", GetInternship).Methods("GET")
  router.HandleFunc("/internships/{id}", CreateInternship).Methods("POST")

  internships = append(internships, Internship{ID: "1", Title: "Amazon summer internship"})
  internships = append(internships, Internship{ID: "2", Title: "Facebook winter internship"})

  log.Fatal(http.ListenAndServe(":8000", router))
}

func GetInternships(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(internships)
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

func CreateInternship(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  var internship Internship
  _ = json.NewDecoder(r.Body).Decode(&internship)
  internship.ID = params["id"]

  for _, item := range internships {
    if item.ID == internship.ID {
      http.Error(w, "Internship with this ID already exists", http.StatusBadRequest)
      return
    }
  }

  internships = append(internships, internship)
  json.NewEncoder(w).Encode(internships)
}