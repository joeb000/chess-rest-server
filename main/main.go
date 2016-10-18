package main

import (
	"log"
	"net/http"
)

var hub = NewHub()

func main() {

	go hub.Run()
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
