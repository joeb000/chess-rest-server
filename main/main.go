package main

import (
	"fmt"
	"log"
	"net/http"
)

var hub = NewHub()

func main() {

	go hub.Run()

	fmt.Println("ddd!")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
