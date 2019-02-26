package main

import (
	"log"
	"net/http"
)

func main() {
	seedData()
	log.Fatal(http.ListenAndServe(":8000", router))
}
