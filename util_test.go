package main

import (
	"github.com/riyaz-ali/chess/pkg/chess"
	"testing"
)

func TestJoin(t *testing.T) {
	var cases = []struct {
		Input  []chess.Position
		Output string
	}{
		{[]chess.Position{"A1", "B2", "C3", "D4", "E5", "F6", "G7", "H8"}, "A1, B2, C3, D4, E5, F6, G7, H8"},
		{[]chess.Position{"A1"}, "A1"},
		{nil, ""},
	}

	for _, c := range cases {
		if got := Join(c.Input, ", "); got != c.Output {
			t.Errorf("Join(%v, \", \"): %v", c.Input, got)
		}
	}
}
