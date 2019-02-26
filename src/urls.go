package main

import (
	views "./Views"
	"github.com/gorilla/mux"
)

func importRoutes() *mux.Router {
	router := mux.NewRouter()
	views.RegisterPeopleView(&router)
}
