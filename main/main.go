package main

import (
	"log"
	"net/http"
)

var hub = NewHub()

func main() {

	CreateGame(Player{})
	JoinGame(Player{}, 1)

	go hub.Run()
	router := NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../public/"))))

	log.Fatal(http.ListenAndServe(":8080", router))

}
