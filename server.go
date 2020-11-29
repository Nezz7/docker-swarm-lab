package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
)



func main() {
	
	fmt.Printf("Starting server at port 8080\n")
	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() 
	
	
	s.HandleFunc("/user", createProfile).Methods("POST")
	s.HandleFunc("/users", getAllUsers).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8080", s))

}