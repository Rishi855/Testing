package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users",GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}",GetUser).Methods("GET")
	r.HandleFunc("/users",CreateUser).Methods("POST")

	http.ListenAndServe(":8080",r)
}