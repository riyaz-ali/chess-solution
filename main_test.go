package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	var cases = []struct {
		In, Out string
	}{
		{"pawn, E3", "E4"},
		{"king, D5", "C4, C5, C6, D4, D6, E4, E5, E6"},
		{"queen, E4", "A4, A8, B1, B4, B7, C2, C4, C6, D3, D4, D5, E1, E2, E3, E5, E6, E7, E8, F3, F4, F5, G2, G4, G6, H1, H4, H7"},
	}

	for _, c := range cases {
		var buf bytes.Buffer
		if err := run(c.In, &buf); err != nil {
			t.Errorf("failed to run test: %s", err.Error())
		} else if buf.String() != c.Out {
			t.Errorf("got %q, expected %q", buf.String(), c.Out)
		}
	}
}

func TestRun_invalidInput(t *testing.T) {
	var cases = []string{"king, D-5", "queen E4"}
	for _, c := range cases {
		var buf bytes.Buffer
		if err := run(c, &buf); err == nil {
			t.Logf("output: %s", buf.String())
			t.Error("expected error")
		}
	}
}
