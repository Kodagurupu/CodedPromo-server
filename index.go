package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var peopleCount = 0

type servResponce struct {
	People int `json:"people"`
}

func printPeople(w http.ResponseWriter, r *http.Request) {
	var output servResponce
	w.Header().Set("Content-Type", "application/json")

	output.People = peopleCount
	json.NewEncoder(w).Encode(output)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	var input servResponce

	json.NewDecoder(r.Body).Decode(&input)
	peopleCount = input.People
}

func main() {
	log.Printf("Starting server")

	router := mux.NewRouter()
	router.HandleFunc("/getPeople", printPeople).Methods("GET")
	router.HandleFunc("/sendPeople", getPeople).Methods("POST")

	defer func() {
		log.Fatal(http.ListenAndServe(":1080", router))
	}()
}
