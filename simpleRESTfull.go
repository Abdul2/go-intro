package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Personid string `json:"personid,omitempty"`
	Object   string `json:"object,omitempty"`
	Location string `json:"location,omitempty"`
	Event    *Event `json:"event,omitempty"`
}

type Event struct {
	Date      string `json:"city,omitempty"`
	Eventtype string `json:"state,omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, record := range people {
		if record.Personid == params["personid"] {
			json.NewEncoder(w).Encode(record)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
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
	//people = append(people, Person{Personid: "1", Object: "passport", Location: "London", Event: &Event{Date: "Dublin", Eventtype: "CA"}})
	people = Readdata("data.json")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8082", router))
}
