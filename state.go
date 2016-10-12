package main

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
