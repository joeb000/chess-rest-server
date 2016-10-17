package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
)

type Fields struct{ val string }

const (
	innerhtml = `<div class="chessboard">
<!-- 1st -->
<div id="a0" class="white">&#9820;</div>
<div id="b0" class="black">&#9822;</div>
<div id="c0" class="white">&#9821;</div>
<div id="d0" class="black">&#9819;</div>
<div id="e0" class="white">&#9818;</div>
<div id="f0" class="black">&#9821;</div>
<div id="g0" class="white">&#9822;</div>
<div id="h0" class="black">&#9820;</div>
<!-- 2nd -->
<div id="a1" class="black">&#9823;</div>
<div id="b1" class="white">&#9823;</div>
<div id="c1" class="black">&#9823;</div>
<div id="d1" class="white">&#9823;</div>
<div id="e1" class="black">&#9823;</div>
<div id="f1" class="white">&#9823;</div>
<div id="g1" class="black">&#9823;</div>
<div id="h1" class="white">&#9823;</div>
<!-- 3th -->
<div id="a2" class="white"></div>
<div id="b2" class="black"></div>
<div id="c2" class="white"></div>
<div id="d2" class="black"></div>
<div id="e2" class="white"></div>
<div id="f2" class="black"></div>
<div id="g2" class="white"></div>
<div id="h2" class="black"></div>
<!-- 4st -->
<div id="a3" class="black"></div>
<div id="b3" class="white"></div>
<div id="c3" class="black"></div>
<div id="d3" class="white"></div>
<div id="e3" class="black"></div>
<div id="f3" class="white"></div>
<div id="g3" class="black"></div>
<div id="h3" class="white"></div>
<!-- 5th -->
<div id="a4" class="white"></div>
<div id="b4" class="black"></div>
<div id="c4" class="white"></div>
<div id="d4" class="black"></div>
<div id="e4" class="white"></div>
<div id="f4" class="black"></div>
<div id="g4" class="white"></div>
<div id="h4" class="black"></div>
<!-- 6th -->
<div id="a5" class="black"></div>
<div id="b5" class="white"></div>
<div id="c5" class="black"></div>
<div id="d5" class="white"></div>
<div id="e5" class="black"></div>
<div id="f5" class="white"></div>
<div id="g5" class="black"></div>
<div id="h5" class="white"></div>
<!-- 7th -->
<div id="a6" class="white">&#9817;</div>
<div id="b6" class="black">&#9817;</div>
<div id="c6" class="white">&#9817;</div>
<div id="d6" class="black">&#9817;</div>
<div id="e6" class="white">&#9817;</div>
<div id="f6" class="black">&#9817;</div>
<div id="g6" class="white">&#9817;</div>
<div id="h6" class="black">&#9817;</div>
<!-- 8th -->
<div id="a7" class="black">&#9814;</div>
<div id="b7" class="white">&#9816;</div>
<div id="c7" class="black">&#9815;</div>
<div id="d7" class="white">&#9813;</div>
<div id="e7" class="black">&#9812;</div>
<div id="f7" class="white">&#9815;</div>
<div id="g7" class="black">&#9816;</div>
<div id="h7" class="white">&#9814;</div>
</div>`
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

func ChessFormMove(w http.ResponseWriter, r *http.Request) {
	var move Move
	move.FromSquare = r.FormValue("from_square")
	move.ToSquare = r.FormValue("to_square")

	move.GameID, _ = strconv.Atoi(r.FormValue("game_id"))

	ProcessMove(move)

	http.Redirect(w, r, "/chess/"+strconv.Itoa(move.GameID), http.StatusFound)

}

func ShowBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sgid := vars["gameid"]
	gid, err := strconv.Atoi(sgid)

	t := template.New("emptyboard.html")
	temp, err := t.ParseFiles("../public/emptyboard.html")
	if err != nil {
		log.Fatal(err)
	}
	game := *FindGame(gid)
	board := game.Board

	//bo := Both{B: board, G: game}

	err = temp.Execute(w,
		struct {
			B  Board
			G  Game
			Ip string
		}{board, game, r.Host})
	if err != nil {
		log.Fatal(err)
	}
}

func TemplateBoard(w http.ResponseWriter, r *http.Request) {
	t := template.New("emptyboard.html")
	temp, err := t.ParseFiles("../public/emptyboard.html")
	if err != nil {
		log.Fatal(err)
	}

	board := BuildNewBoard()

	err = temp.Execute(w, board)
	if err != nil {
		log.Fatal(err)
	}
}

func GetSandaloneChat(w http.ResponseWriter, r *http.Request) {
	t := template.New("home.html")
	homeTemp, err := t.ParseFiles("../public/home.html")
	if err != nil {
		log.Fatal(err)
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Println(r.Host)
	homeTemp.Execute(w, r.Host)
}

func ServeSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hersse!")

	ServeWs(hub, w, r)
	fmt.Println("Hereaaa!")

}
