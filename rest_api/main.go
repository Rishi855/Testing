package main

import "net/http"

func main() {

	http.HandleFunc("/users", userHandler)     // get , post
	http.HandleFunc("/users", userByIdHandler) // get , put,  delete

	http.ListenAndServe(":8080", nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
	}
}

func userByIdHandler(w http.ResponseWriter, r *http.Request){
	
}