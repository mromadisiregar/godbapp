package main

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"myapp/api"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/counter", api.checkCounter).Methods("GET")
	r.HandleFunc("/api/users", api.createUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", api.getUser).Methods("GET")
	r.HandleFunc("/api/users/{id}", api.deleteUser).Methods("DELETE")
	r.HandleFunc("/api/users", api.getAllUser).Methods("GET")
	if err := http.ListenAndServe(":8088", r); err != nil {
		log.Fatal(err)
	}
}	
