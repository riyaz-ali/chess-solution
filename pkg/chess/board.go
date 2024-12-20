// Package chess provide services to build and emulate a chess board.
package chess

import "fmt"

// Board represent a MxN size chess board.
type Board struct{ Col, Row int }

// New returns a new Board of size n*n
func New(n int) *Board { return &Board{Col: n, Row: n} }

// WithinBound returns true if the given position is within this board's bounding grid.
func (board *Board) WithinBound(pos Position) bool {
	if !pos.Valid() {
		return false
	}

	col, row := pos.Split()

	// we convert col from rune to an integer in the range 1-26 before comparing
	//
	// NOTE: because Position only supports a single character for column, the
	// effective size limit for a board is 26xN grid.
	if (col < 1 || col > board.Col) || (row < 1 || row > board.Row) {
		return false
	}

	return true
}

// Piece represents a single piece on the chess board.
type Piece interface {
	fmt.Stringer

	// Pos returns the piece's current position
	Pos() Position

	// ListAll returns a list of all possible position that this
	// piece can move to on the given board
	ListAll(board *Board) []Position
}

// Position represents a piece's position on the board.
// It is a string of the form <col><row> (eg. A8)
//
// NOTE: there is an hard-coded assumption that col will always be a single-character
// and within the A-Z range only, and row is a non-zero, positive integer number.
type Position string

func Pos(col, row int) Position     { return Position(fmt.Sprintf("%c%d", col-1+'A', row)) }
func (pos Position) String() string { return string(pos) }

// Split splits the position into col-row pair values. It does not guarantee
// that the given pos is valid, use Valid() to check if the position is valid.
func (pos Position) Split() (col, row int) {
	_, _ = fmt.Sscanf(string(pos), "%c%d", &col, &row)
	return (col - 'A') + 1, row
}

// Valid returns true if the given position is valid.
func (pos Position) Valid() bool {
	col, row := pos.Split()
	return col >= 1 && col <= 26 && row > 0
}
