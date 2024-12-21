package piece_test

import (
	"github.com/riyaz-ali/chess/pkg/chess"
	"github.com/riyaz-ali/chess/pkg/chess/piece"
	"slices"
	"testing"
)

func assert(t *testing.T, board *chess.Board, piece chess.Piece, expected ...chess.Position) {
	t.Helper()
	board.Reset()
	_ = board.Move(piece, piece.Pos()) // move piece to its position on the board
	var positions = board.ListMoves(piece)

	if len(positions) != len(expected) {
		t.Errorf("expected %d positions, got %d", len(expected), len(positions))
		return
	}

	for _, exp := range expected {
		if i := slices.IndexFunc(positions, func(p chess.Position) bool { return exp == p }); i == -1 {
			t.Errorf("expected piece: %v, position: %v", piece, exp)
		}
	}
}

func TestPawn(t *testing.T) {
	var board = chess.NewBoard(8)

	assert(t, board, must(piece.New(piece.KindPawn, "D4")), "D5")
	assert(t, board, must(piece.New(piece.KindPawn, "A4")), "A5")
	assert(t, board, must(piece.New(piece.KindPawn, "E8"))) // cannot move further
}

func TestKing(t *testing.T) {
	var board = chess.NewBoard(8)

	assert(t, board, must(piece.New(piece.KindKing, "D5")), "C4", "C5", "C6", "D4", "D6", "E4", "E5", "E6")
	assert(t, board, must(piece.New(piece.KindKing, "D1")), "C1", "C2", "D2", "E2", "E1")
	assert(t, board, must(piece.New(piece.KindKing, "A1")), "A2", "B2", "B1")
	assert(t, board, must(piece.New(piece.KindKing, "A8")), "A7", "B8", "B7")
}

func TestQueen(t *testing.T) {
	var board = chess.NewBoard(8)

	assert(t, board, must(piece.New(piece.KindQueen, "E4")), "A4", "B4", "C4", "D4", "F4", "G4", "H4", "E1", "E2", "E3", "E5", "E6",
		"E7", "E8", "A8", "B7", "C6", "D5", "F3", "G2", "H1", "B1", "C2", "D3", "F5", "G6", "H7")

	assert(t, board, must(piece.New(piece.KindQueen, "A1")), "A2", "A3", "A4", "A5", "A6", "A7", "A8",
		"B1", "C1", "D1", "E1", "F1", "G1", "H1", "B2", "C3", "D4", "E5", "F6", "G7", "H8")
}

func must(p chess.Piece, err error) chess.Piece {
	if err != nil {
		panic(err)
	}

	return p
}
