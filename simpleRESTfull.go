// this program will be used what internal training to show features of go
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Abdul2/readdata"
)

// Poeple holds records
var people []readdata.Person

//GetPersonEndpoint returns json record of requested id that
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, record := range people {
		if record.Personid == params["personid"] {
			json.NewEncoder(w).Encode(record)
			return
		}
	}
	json.NewEncoder(w).Encode(&readdata.Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person readdata.Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.Personid = params["personid"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for index, record := range people {
		if record.Personid == params["ipersonid"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = readdata.Readdata("data.json")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8082", router))
}
