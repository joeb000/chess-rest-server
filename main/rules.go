package main

//"fmt"

func IsValidMove(m Move) bool {
	game := GetBoardForID(m.GameID)
	piece := game.getPieceAt(m.FromSquare)
	switch piece.Code
	case WhiteBishop:
		return g.A8
	case "A1":
		return g.A1

	return false
}
