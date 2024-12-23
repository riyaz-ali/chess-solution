package chess_test

import (
	"github.com/riyaz-ali/chess/pkg/chess"
	"reflect"
	"sort"
	"testing"
)

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
