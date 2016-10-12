package main

import (
	"fmt"
	"strconv"
)

var currentGameId int

var games Games

func FindGame(id int) Game {
	for _, g := range games {
		if g.Id == id {
			return g
		}
	}
	return Game{}
}

func JoinGame(p Player, i int) Game {
	g := FindGame(i)
	g.P2 = p
	g.initState()
	//todo
	pi := Piece{true, "Rook", 2}
	g.placePiece("a2", pi)
	return g
}

func CreateGame(p Player) Game {
	g := Game{}
	currentGameId += 1
	g.Id = currentGameId
	g.P1 = p
	games = append(games, g)
	return g
}

func (g *Game) initState() {
	for col := range g.CurrentState {
		for row := range g.CurrentState {
			g.CurrentState[col][row] = Piece{false, "piece", col + row}
		}
	}
}

func (g *Game) placePiece(square string, p Piece) {
	if len(square) != 2 {
		fmt.Println("error placing piece")
	}
	col := convertColString(square[0:1])
	row, _ := strconv.Atoi(square[1:2])
	g.CurrentState[row][col] = p

}

func convertColString(s string) int {
	switch s {
	case "a":
		return 0
	case "b":
		return 1
	case "c":
		return 2
	case "d":
		return 3
	case "e":
		return 4
	case "f":
		return 5
	case "g":
		return 6
	case "h":
		return 7
	}
	return 99
}
