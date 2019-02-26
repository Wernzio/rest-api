package Views

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	db "../Database"

	"github.com/gorilla/mux"
)

func RegisterPeopleView(r *mux.Router) {
	r.HandleFunc("/people", GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	r.HandleFunc("/people", CreatePerson).Methods("POST")
	r.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	people := db.GetPeople()
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}
	person, err := db.GetPerson(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person db.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(person, r.Body)
	person = *db.CreatePerson(&person)
	json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_ = db.DeletePerson(id)
	w.WriteHeader(204)
}
