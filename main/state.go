package main

import (
	"fmt"
)

var currentGameId int

var chessgames []Game

func FindGame(id int) *Game {
	for i, _ := range chessgames {
		if chessgames[i].Id == id {
			return &chessgames[i]
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
func GetBoardForID(i int) Board {
	return Board{}
}

func CreateGame(p Player) Game {
	g := Game{}
	currentGameId += 1
	g.Id = currentGameId
	g.P1 = p
	g.LastMove = &Move{currentGameId, 0, "A1", "A1"}
	chessgames = append(chessgames, g)
	return g
}

func ProcessMove(move Move) {
	g := FindGame(move.GameID)
	p := g.Board.getPieceAt(move.FromSquare)
	g.Board.setPieceAt(move.ToSquare, p)
	g.Board.setEmpty(move.FromSquare)
	fmt.Printf("piece at %s, is %v\n", move.FromSquare, p)
	//p2 := *g.Board.getPieceAt(move.ToSquare)
	//p2 = p

	//move.FromSquare
}

func (g *Game) initState() {
	gameBoard := BuildNewBoard()
	g.Board = gameBoard
}
