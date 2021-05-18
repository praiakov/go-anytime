package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	entities "github.com/praiakov/webapi/core"
)

var people []entities.Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&entities.Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	person := entities.Person{}
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func LoadData() {
	people = append(people, entities.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &entities.Address{City: "City X", State: "State X"}})
	people = append(people, entities.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &entities.Address{City: "City Z", State: "State Y"}})
	people = append(people, entities.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

}
