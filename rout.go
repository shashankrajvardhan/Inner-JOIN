package main

import "github.com/gorilla/mux"

func rout(r *mux.Router) {
	r.HandleFunc("/home/fetch", getUsers).Methods("GET")
}
