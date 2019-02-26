package Views

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func registerPeopleUrl(r *mux.Router) {
	r.HandleFunc("/people", GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	r.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	r.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
}

func GetPeople(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(people)
}

func GetPerson(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for _, item := range people {
		if item.ID == params["id"] {
			_ = json.NewEncoder(response).Encode(item)
			break
		}
	}
}

func CreatePerson(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var person Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(response).Encode(people)
}

func DeletePerson(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1]...)
			break
		}
	}
}
