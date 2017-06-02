package main

import (
	"fmt"

	"encoding/json"
	"io/ioutil"
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

func Readdata(f string) []Person {

	fmt.Println("Hello")

	content, err := ioutil.ReadFile(f)

	if err != nil {

		fmt.Println(err)

	}

	var a []Person

	json.Unmarshal(content, &a)

	for _, value := range a {

		fmt.Println(value.Personid)

	}

	return a

}

func main() {

	Readdata("data.json")
}
