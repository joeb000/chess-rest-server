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
	Id       int    `json:"id"`
	P1       Player `json:"player1"`
	P2       Player `json:"player2"`
	Board    Board  `json:"board"`
	LastMove Move   `json:"last_move"`
}

type Player struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
	auth   string `json:"auth_string"`
}

type Piece struct {
	Code     string `json:"code"`
	PlayerId int    `json:"player_id"`
}

type Board struct {
	A8, B8, C8, D8, E8, F8, G8, H8,
	A1, B1, C1, D1, E1, F1, G1, H1,
	A2, B2, C2, D2, E2, F2, G2, H2,
	A3, B3, C3, D3, E3, F3, G3, H3,
	A4, B4, C4, D4, E4, F4, G4, H4,
	A5, B5, C5, D5, E5, F5, G5, H5,
	A6, B6, C6, D6, E6, F6, G6, H6,
	A7, B7, C7, D7, E7, F7, G7, H7 Piece
}
