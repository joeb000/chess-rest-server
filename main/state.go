package main

import (
	"fmt"
	"strconv"
	"strings"
)

var currentGameId int

var games Games

func FindGame(id int) *Game {
	for i, _ := range games {
		if games[i].Id == id {
			return &games[i]
		}
	}
	emptyGame := Game{}
	return &emptyGame
}

func JoinGame(p Player, i int) Game {
	g := FindGame(i)
	g.P2 = p
	g.initState()
	//todo

	return *g
}

func CreateGame(p Player) Game {
	g := Game{}
	currentGameId += 1
	g.Id = currentGameId
	g.P1 = p
	games = append(games, g)
	return g
}

func ProcessMove(move Move) {
	g := FindGame(move.GameID)
	fmt.Println("found game")
	//get piece at from
	piece := g.getPieceAt(move.FromSquare)
	//move piece to
	if g.movePiece(piece, move.ToSquare) {
		g.makeEmpty(move.FromSquare)
	}

}

func (g *Game) initState() {
	for col := range g.CurrentState {
		for row := range g.CurrentState {
			g.CurrentState[col][row] = Piece{}
		}
	}
	wPawn := Piece{true, "Pawn", 1}
	wKing := Piece{true, "King", 1}
	wQueen := Piece{true, "Queen", 1}
	wRook := Piece{true, "Rook", 1}
	wBishop := Piece{true, "Bishop", 1}
	wKnight := Piece{true, "Knight", 1}

	bPawn := Piece{false, "Pawn", 2}
	bKing := Piece{false, "King", 2}
	bQueen := Piece{false, "Queen", 2}
	bRook := Piece{false, "Rook", 2}
	bBishop := Piece{false, "Bishop", 2}
	bKnight := Piece{false, "Knight", 2}

	g.placePiece("A0", bRook)
	g.placePiece("B0", bKnight)
	g.placePiece("C0", bBishop)
	g.placePiece("D0", bKing)
	g.placePiece("E0", bQueen)
	g.placePiece("F0", bBishop)
	g.placePiece("G0", bKnight)
	g.placePiece("H0", bRook)
	g.placePiece("A1", bPawn)
	g.placePiece("B1", bPawn)
	g.placePiece("C1", bPawn)
	g.placePiece("D1", bPawn)
	g.placePiece("E1", bPawn)
	g.placePiece("F1", bPawn)
	g.placePiece("G1", bPawn)
	g.placePiece("H1", bPawn)

	g.placePiece("A7", wRook)
	g.placePiece("B7", wKnight)
	g.placePiece("C7", wBishop)
	g.placePiece("D7", wKing)
	g.placePiece("E7", wQueen)
	g.placePiece("F7", wBishop)
	g.placePiece("G7", wKnight)
	g.placePiece("H7", wRook)
	g.placePiece("A6", wPawn)
	g.placePiece("B6", wPawn)
	g.placePiece("C6", wPawn)
	g.placePiece("D6", wPawn)
	g.placePiece("E6", wPawn)
	g.placePiece("F6", wPawn)
	g.placePiece("G6", wPawn)
	g.placePiece("H6", wPawn)
}

func (g *Game) placePiece(square string, p Piece) {
	if len(square) != 2 {
		fmt.Println("error placing piece")
	}
	col := convertColString(square[0:1])
	row, _ := strconv.Atoi(square[1:2])
	g.CurrentState[row][col] = p
}

func (g *Game) getPieceAt(square string) Piece {
	col := convertColString(square[0:1])
	row, _ := strconv.Atoi(square[1:2])
	return g.CurrentState[row][col]
}

func (g *Game) movePiece(piece Piece, square string) bool {
	col := convertColString(square[0:1])
	row, _ := strconv.Atoi(square[1:2])
	g.CurrentState[row][col] = piece
	return true
}

func (g *Game) makeEmpty(square string) {
	col := convertColString(square[0:1])
	row, _ := strconv.Atoi(square[1:2])
	g.CurrentState[row][col] = Piece{}
}

func convertColString(s string) int {
	s = strings.ToLower(s)
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
