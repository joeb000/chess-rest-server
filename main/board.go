package main

import (
	"fmt"
	"strings"
)

func BuildNewBoard() Board {
	newBoard := Board{}
	newBoard.A0.Code = BlackRook
	newBoard.B0.Code = BlackKnight
	newBoard.C0.Code = BlackBishop
	newBoard.D0.Code = BlackQueen
	newBoard.E0.Code = BlackKing
	newBoard.F0.Code = BlackBishop
	newBoard.G0.Code = BlackKnight
	newBoard.H0.Code = BlackRook

	newBoard.A1.Code = BlackPawn
	newBoard.B1.Code = BlackPawn
	newBoard.C1.Code = BlackPawn
	newBoard.D1.Code = BlackPawn
	newBoard.E1.Code = BlackPawn
	newBoard.F1.Code = BlackPawn
	newBoard.G1.Code = BlackPawn
	newBoard.H1.Code = BlackPawn

	newBoard.A6.Code = WhitePawn
	newBoard.B6.Code = WhitePawn
	newBoard.C6.Code = WhitePawn
	newBoard.D6.Code = WhitePawn
	newBoard.E6.Code = WhitePawn
	newBoard.F6.Code = WhitePawn
	newBoard.G6.Code = WhitePawn
	newBoard.H6.Code = WhitePawn

	newBoard.A7.Code = WhiteRook
	newBoard.B7.Code = WhiteKnight
	newBoard.C7.Code = WhiteBishop
	newBoard.D7.Code = WhiteQueen
	newBoard.E7.Code = WhiteKing
	newBoard.F7.Code = WhiteBishop
	newBoard.G7.Code = WhiteKnight
	newBoard.H7.Code = WhiteRook

	return newBoard
}

const (
	BlackRook   = "&#9820;"
	BlackKnight = "&#9822;"
	BlackBishop = "&#9821;"
	BlackKing   = "&#9818;"
	BlackQueen  = "&#9819;"
	BlackPawn   = "&#9823;"

	WhiteRook   = "&#9814;"
	WhiteKnight = "&#9816;"
	WhiteBishop = "&#9815;"
	WhiteKing   = "&#9812;"
	WhiteQueen  = "&#9813;"
	WhitePawn   = "&#9817;"
)

func (b *Board) setEmpty(square string) {
	b.setPieceAt(square, Piece{})
}

func (g *Board) getPieceAt(square string) Piece {
	switch strings.ToUpper(square) {
	case "A0":
		return g.A0
	case "A1":
		return g.A1
	case "A2":
		return g.A2
	case "A3":
		return g.A3
	case "A4":
		return g.A4
	case "A5":
		return g.A5
	case "A6":
		return g.A6
	case "A7":
		return g.A7

	case "B0":
		return g.B0
	case "B1":
		return g.B1
	case "B2":
		return g.B2
	case "B3":
		return g.B3
	case "B4":
		return g.B4
	case "B5":
		return g.B5
	case "B6":
		return g.B6
	case "B7":
		return g.B7

	case "C0":
		return g.C0
	case "C1":
		return g.C1
	case "C2":
		return g.C2

	case "C3":
		return g.C3
	case "C4":
		return g.C4
	case "C5":
		return g.C5
	case "C6":
		return g.C6
	case "C7":
		return g.C7

	case "D0":
		return g.D0
	case "D1":
		return g.D1
	case "D2":
		return g.D2
	case "D3":
		return g.D3
	case "D4":
		return g.D4
	case "D5":
		return g.D5
	case "D6":
		return g.D6
	case "D7":
		return g.D7

	case "E0":
		return g.E0
	case "E1":
		return g.E1
	case "E2":
		return g.E2
	case "E3":
		return g.E3
	case "E4":
		return g.E4
	case "E5":
		return g.E5
	case "E6":
		return g.E6
	case "E7":
		return g.E7

	case "F0":
		return g.F0
	case "F1":
		return g.F1
	case "F2":
		return g.F2
	case "F3":
		return g.F3
	case "F4":
		return g.F4
	case "F5":
		return g.F5
	case "F6":
		return g.F6
	case "F7":
		return g.F7

	case "G0":
		return g.G0
	case "G1":
		return g.G1
	case "G2":
		return g.G2
	case "G3":
		return g.G3
	case "G4":
		return g.G4
	case "G5":
		return g.G5
	case "G6":
		return g.G6
	case "G7":
		return g.G7

	case "H0":
		return g.H0
	case "H1":
		return g.H1
	case "H2":
		return g.H2
	case "H3":
		return g.H3
	case "H4":
		return g.H4
	case "H5":
		return g.H5
	case "H6":
		return g.H6
	case "H7":
		return g.H7
	}

	fmt.Errorf("SOMETHING WENT WRONG \n")

	return Piece{}
}

func (g *Board) setPieceAt(square string, p Piece) {
	switch strings.ToUpper(square) {
	case "A0":
		g.A0 = p
	case "A1":
		g.A1 = p
	case "A2":
		g.A2 = p
	case "A3":
		g.A3 = p
	case "A4":
		g.A4 = p
	case "A5":
		g.A5 = p
	case "A6":
		g.A6 = p
	case "A7":
		g.A7 = p

	case "B0":
		g.B0 = p
	case "B1":
		g.B1 = p
	case "B2":
		g.B2 = p
	case "B3":
		g.B3 = p
	case "B4":
		g.B4 = p
	case "B5":
		g.B5 = p
	case "B6":
		g.B6 = p
	case "B7":
		g.B7 = p

	case "C0":
		g.C0 = p
	case "C1":
		g.C1 = p
	case "C2":
		g.C2 = p

	case "C3":
		g.C3 = p
	case "C4":
		g.C4 = p
	case "C5":
		g.C5 = p
	case "C6":
		g.C6 = p
	case "C7":
		g.C7 = p

	case "D0":
		g.D0 = p
	case "D1":
		g.D1 = p
	case "D2":
		g.D2 = p
	case "D3":
		g.D3 = p
	case "D4":
		g.D4 = p
	case "D5":
		g.D5 = p
	case "D6":
		g.D6 = p
	case "D7":
		g.D7 = p

	case "E0":
		g.E0 = p
	case "E1":
		g.E1 = p
	case "E2":
		g.E2 = p
	case "E3":
		g.E3 = p
	case "E4":
		g.E4 = p
	case "E5":
		g.E5 = p
	case "E6":
		g.E6 = p
	case "E7":
		g.E7 = p

	case "F0":
		g.F0 = p
	case "F1":
		g.F1 = p
	case "F2":
		g.F2 = p
	case "F3":
		g.F3 = p
	case "F4":
		g.F4 = p
	case "F5":
		g.F5 = p
	case "F6":
		g.F6 = p
	case "F7":
		g.F7 = p

	case "G0":
		g.G0 = p
	case "G1":
		g.G1 = p
	case "G2":
		g.G2 = p
	case "G3":
		g.G3 = p
	case "G4":
		g.G4 = p
	case "G5":
		g.G5 = p
	case "G6":
		g.G6 = p
	case "G7":
		g.G7 = p

	case "H0":
		g.H0 = p
	case "H1":
		g.H1 = p
	case "H2":
		g.H2 = p
	case "H3":
		g.H3 = p
	case "H4":
		g.H4 = p
	case "H5":
		g.H5 = p
	case "H6":
		g.H6 = p
	case "H7":
		g.H7 = p
	}

	fmt.Errorf("SOMETHING WENT WRONG \n")

}
