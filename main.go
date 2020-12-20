package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

collection := helper.ConnectDB()
func main() {
	//Init Router
	r := mux.NewRouter()

  	// arrange our route
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", getID).Methods("GET")
	r.HandleFunc("/contacts", createContact).Methods("POST")
	r.HandleFunc("/contacts?user={id}&infection_timestamp={timestamp}", updateBook).Methods("GET")

  	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}