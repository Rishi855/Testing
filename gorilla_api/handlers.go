package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var users = make(map[int]User)

func GetUsers(w http.ResponseWriter, r *http.Request){
	var result []User
	for _,u:= range users{
		result = append(result,u)
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(result)
}

func GetUser(w http.ResponseWriter, r *http.Request){
	id,_ := strconv.Atoi(mux.Vars(r)["id"])
	user, exits := users[id]
	if !exits{
		http.Error(w, "user not found",http.StatusNotFound)
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	var u User
	json.NewDecoder(r.Body).Decode(u)
	users[u.Id] = u
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}