package main

import "time"

type Move struct {
	GameID   int       `json:"game_id"`
	PlayerID int       `json:"player_id"`
	PieceID  string    `json:"piece_id"`
	Square   string    `json:"square"`
	Time     time.Time `json:"time"`
}

type Game struct {
	Id int `json:"id"`
	P1 Player
	P2 Player
}

type Games []Game

type Player struct {
	Id     int
	Name   string
	Wins   int
	Losses int
	auth   string
}
