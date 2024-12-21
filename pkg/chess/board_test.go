package chess_test

import (
	"github.com/riyaz-ali/chess/pkg/chess"
	"github.com/riyaz-ali/chess/pkg/chess/piece"
	"reflect"
	"sort"
	"testing"
)

func TestBoard_WithinBound(t *testing.T) {
	var cases = []struct {
		N        int
		Pos      chess.Position
		Expected bool
	}{
		{4, "A8", false},
		{8, "A8", true},
		{8, "a99", false}, // a99 is an invalid position identifier
		{8, "I4", false},  // column I is not valid on a 8x8 grid
		{26, "Z26", true},
	}

	for _, c := range cases {
		if got := chess.NewBoard(c.N).WithinBound(c.Pos); got != c.Expected {
			t.Errorf("Board(n=%d), %s: got %v; want %v", c.N, c.Pos.String(), got, c.Expected)
		}
	}
}

func TestBoard_Move(t *testing.T) {
	var board = chess.NewBoard(8)

	pawn, _ := piece.New(piece.KindPawn, "D2")

	if err := board.Move(pawn, "D10"); err == nil {
		t.Errorf("board.Move(%s, %s) expected error, got nil", pawn, "D10")
	}

	pos := chess.Position("D3")
	if err := board.Move(pawn, pos); err != nil {
		t.Errorf("board.Move(%s, %s): %v", pawn, pos, err)
	}

	if p := board.PieceAt(pos); p != pawn {
		t.Errorf("board.PieceAt(%s): got %v; want %v", pos, p, pawn)
	}

	if pawn.Pos() != pos {
		t.Errorf("%s position not updated", pawn)
	}

	newPosition := pos.Add(0, 1)
	if err := board.Move(pawn, newPosition); err != nil {
		t.Errorf("board.Move(%s, %s): %v", pawn, newPosition, err)
	}

	if p := board.PieceAt(newPosition); p != pawn {
		t.Errorf("board.PieceAt(%s): got %v; want %v", newPosition, p, pawn)
	}

	if pawn.Pos() != newPosition {
		t.Errorf("%s position not updated", pawn)
	}

	// previous position must be cleared
	if board.HasPiece(pos) {
		t.Errorf("board.HasPiece(%s): returned true, expected false", pos)
	}
}

func TestBoard_ListMoves(t *testing.T) {
	var board = chess.NewBoard(8)

	pawn, _ := piece.New(piece.KindPawn, "E5")
	king, _ := piece.New(piece.KindKing, "E5")

	_ = board.Move(pawn, "E5")
	if moves := board.ListMoves(king); len(moves) != 0 {
		t.Errorf("expected no moves for %s, got %d", king, len(moves))
	}

	if moves := board.ListMoves(pawn); len(moves) != 1 {
		t.Errorf("expected 1 moves for %s, got %d", pawn, len(moves))
	}
}

func TestPosition(t *testing.T) {
	var cases = []struct {
		Pos      chess.Position
		Expected bool
	}{
		{"A8", true},
		{"b10", false},
		{"AA9", false},
		{"Z-9", false},
		{"-F10", false},
		{"E99", true}, // technically valid
		{"44", false},
		{"A", false},
	}

	for _, c := range cases {
		if got := c.Pos.Valid(); got != c.Expected {
			t.Errorf("%s: got %v; want %v", c.Pos.String(), got, c.Expected)
		}
	}
}

func TestPosition_Add(t *testing.T) {
	var cases = []struct {
		P, Exp   chess.Position
		Col, Row int
	}{
		{"A1", "A2", 0, 1},
		{"A1", "B1", 1, 0},
		{"A1", "B2", 1, 1},

		{"D4", "D3", 0, -1},
		{"D4", "C4", -1, 0},
		{"D4", "C3", -1, -1},
	}

	for _, c := range cases {
		if got := c.P.Add(c.Col, c.Row); got != c.Exp {
			t.Errorf("%q.Add(%d, %d): got %v; want %v", c.P.String(), c.Col, c.Row, got, c.Exp)
		}
	}
}

func TestPosition_sort(t *testing.T) {
	var positions = []chess.Position{"B5", "A2", "A6", "D1"}
	sort.Sort(chess.ByPosition(positions))

	var expected = []chess.Position{"A2", "A6", "B5", "D1"}
	if !reflect.DeepEqual(positions, expected) {
		t.Errorf("got %v; want %v", positions, expected)
	}
}
