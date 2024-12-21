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

func TestPosition_sort(t *testing.T) {
	var positions = []chess.Position{"B5", "A2", "A6", "D1"}
	sort.Sort(chess.ByPosition(positions))

	var expected = []chess.Position{"A2", "A6", "B5", "D1"}
	if !reflect.DeepEqual(positions, expected) {
		t.Errorf("got %v; want %v", positions, expected)
	}
}
