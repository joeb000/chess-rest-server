package main

//import "time"

type Move struct {
	GameID     int    `json:"game_id"`
	PlayerID   int    `json:"player_id"`
	FromSquare string `json:"from_square"`
	ToSquare   string `json:"to_square"`
	//Time     time.Time `json:"time"`
}
type Both struct {
	B Board
	G Game
}

type OldGame struct {
	Id           int `json:"id"`
	P1           Player
	P2           Player
	CurrentState [8][8]OldPiece
}

type Game struct {
	Id    int `json:"id"`
	P1    Player
	P2    Player
	Board Board
}

type Games []OldGame

type Player struct {
	Id     int
	Name   string
	Wins   int
	Losses int
	auth   string
}

type OldPiece struct {
	Color bool
	Type  string
	Id    int
}
type Piece struct {
	Code     string
	PlayerId int
}

type Column struct {
	a, b, c, d, e, f, g, h string
}

type Board struct {
	A0, B0, C0, D0, E0, F0, G0, H0,
	A1, B1, C1, D1, E1, F1, G1, H1,
	A2, B2, C2, D2, E2, F2, G2, H2,
	A3, B3, C3, D3, E3, F3, G3, H3,
	A4, B4, C4, D4, E4, F4, G4, H4,
	A5, B5, C5, D5, E5, F5, G5, H5,
	A6, B6, C6, D6, E6, F6, G6, H6,
	A7, B7, C7, D7, E7, F7, G7, H7 Piece
}
