package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func ChessIndex(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "welcome to chess")
}

func ChessState(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gid := vars["gameid"]
	fmt.Fprintln(w, "Game ID:", gid)
}

func ChessMove(w http.ResponseWriter, r *http.Request) {
	var move Move
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &move); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(move); err != nil {
		panic(err)
	}
}
