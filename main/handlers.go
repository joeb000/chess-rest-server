package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

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

func ChessFind(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sgid := vars["gameid"]
	gid, _ := strconv.Atoi(sgid)

	g := FindGame(gid)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(g); err != nil {
		panic(err)
	}

	//fmt.Fprintln(w, "Game ID:", gid)
}

func ChessCreate(w http.ResponseWriter, r *http.Request) {
	var player Player
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &player); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	g := CreateGame(player)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(g); err != nil {
		panic(err)
	}

}

func ChessJoin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sgid := vars["gameid"]
	gid, err := strconv.Atoi(sgid)
	var player Player
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &player); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	g := JoinGame(player, gid)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(g); err != nil {
		panic(err)
	}

}

func DisplayBoard(w http.ResponseWriter, r *http.Request) {
	path := "../public/chessboard.html"

	//	vars := mux.Vars(r)
	//	sgid := vars["gameid"]
	//	gid, err := strconv.Atoi(sgid)

	f, err := os.Open(path)

	if err == nil {
		bufferedReader := bufio.NewReader(f)
		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "text/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".mp4") {
			contentType = "video/mp4"
		} else {
			contentType = "text/plain"
		}
		w.Header().Add("Content-Type", contentType)
		bufferedReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("shit: 404 - " + http.StatusText(404)))
	}
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
	ProcessMove(move)

	state := FindGame(move.GameID)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(state); err != nil {
		panic(err)
	}
}
