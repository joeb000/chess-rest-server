package main

//import "time"

type Move struct {
	GameID     int    `json:"game_id"`
	PlayerID   int    `json:"player_id"`
	FromSquare string `json:"from_square"`
	ToSquare   string `json:"to_square"`
	//Time     time.Time `json:"time"`
}

type Game struct {
	Id           int `json:"id"`
	P1           Player
	P2           Player
	CurrentState [8][8]Piece
}

type Games []Game

type Player struct {
	Id     int
	Name   string
	Wins   int
	Losses int
	auth   string
}

type Piece struct {
	Color bool
	Type  string
	Id    int
}
