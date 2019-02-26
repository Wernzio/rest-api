package main

import (
	"log"
	"net/http"

	views "./Views"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	views.RegisterPeopleView(router)
	log.Println("Listing to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
