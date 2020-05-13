package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}
func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Person{})

}
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	var person Person
	params := mux.Vars(r)
	person.ID = params["id"] //person.ID = len(people) + 1 ESTO SIRVE PARA QUE SEA AUTOINCREMENTABLE
	_ = json.NewDecoder(r.Body).Decode(&person)
	people = append(people, person)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range people {
		if item.ID == params["id"] {
			people = append(people[:i], people[i+1:]...)
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}
func UpdatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	var person Person
	params := mux.Vars(r)
	for i, item := range people {
		if item.ID == params["id"] {
			_ = json.NewDecoder(r.Body).Decode(&person)
			person.ID = item.ID
			buffer := append(people[:i], person)
			people = append(buffer, people[i+1:]...)
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}
func main() {
	router := mux.NewRouter()
	p := Person{ID: "6", FirstName: "Luis", LastName: "Rodriguez"}
	people = append(people, Person{ID: "1", FirstName: "Pedro", LastName: "Gonzalez", Address: &Address{City: "CABA", State: "Buenos Aires"}})
	people = append(people, Person{ID: "2", FirstName: "Roberto", LastName: "Sanchez", Address: &Address{City: "Rio Cuarto", State: "Cordoba"}})
	people = append(people, Person{ID: "3", FirstName: "Mario", LastName: "Rodriguez", Address: &Address{City: "GBA", State: "Buenos Aires"}})
	people = append(people, Person{ID: "4", FirstName: "Rolando", LastName: "Schiavi", Address: &Address{City: "Obera", State: "Misiones"}})
	people = append(people, Person{ID: "5", FirstName: "Arturo", LastName: "Perez", Address: &Address{City: "Gral. Alvear", State: "Mendoza"}})
	people = append(people, p)
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/update/{id}", UpdatePersonEndpoint).Methods("POST") //Podria haber sido peticion PUT para no tener que agregar el /update/
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
