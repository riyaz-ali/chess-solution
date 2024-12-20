package chess

import "testing"

func TestBoard_WithinBound(t *testing.T) {
	var cases = []struct {
		N        int
		Pos      Position
		Expected bool
	}{
		{4, "A8", false},
		{8, "A8", true},
		{8, "a99", false}, // a99 is an invalid position identifier
		{8, "I4", false},  // column I is not valid on a 8x8 grid
		{26, "Z26", true},
	}

	for _, c := range cases {
		if got := New(c.N).WithinBound(c.Pos); got != c.Expected {
			t.Errorf("Board(n=%d), %s: got %v; want %v", c.N, c.Pos.String(), got, c.Expected)
		}
	}
}

func TestPosition(t *testing.T) {
	var cases = []struct {
		Pos      Position
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
