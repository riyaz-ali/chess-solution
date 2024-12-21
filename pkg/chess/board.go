// Package chess provide services to build and emulate a chess board.
package chess

import "fmt"

// Board represent a MxN size chess board.
type Board struct {
	col, row int // the board's grid dimensions

	// current state of the board
	pieces map[Position]Piece // map of position to pieces
}

// NewBoard returns a new Board of size n*n
func NewBoard(n int) *Board { return &Board{col: n, row: n, pieces: make(map[Position]Piece)} }

// WithinBound returns true if the given position is within this board's bounding grid.
func (board *Board) WithinBound(pos Position) bool {
	if !pos.Valid() {
		return false
	}

	col, row := pos.Split()
	if (col < 1 || col > board.col) || (row < 1 || row > board.row) {
		return false
	}

	return true
}

// Dimension return the board's dimensions represented by number of columns and rows.
func (board *Board) Dimension() (col, row int) { return board.col, board.row }

// HasPiece returns true if there is a Piece at given position
func (board *Board) HasPiece(pos Position) bool { _, has := board.pieces[pos]; return has }

// Move moves the given piece to the provided location.
//
// If the piece was not previously on the board, it is added.
// If the piece was already on the board, its position is updated.
//
// If any other piece is there on the given position it is overridden.
//
// The caller MUST ensure that the move being made is valid, i.e. the pos appears
// in board.ListMove(piece) results.
func (board *Board) Move(piece Piece, pos Position) error {
	if !pos.Valid() || !board.WithinBound(pos) {
		return fmt.Errorf("invalid position: %s", pos.String())
	}

	if _, ok := board.pieces[piece.Pos()]; ok {
		delete(board.pieces, piece.Pos()) // remove from old location if piece is there
	}

	// TODO(@riyaz): update scoring after piece capture
	// if captured, ok := board.pieces[pos]; ok {
	//
	// }

	board.pieces[pos] = piece // overrides what was there at pos (old piece got captured)
	piece.SetPos(pos)

	return nil
}

// PieceAt returns the piece at the given position, or nil if none exists
func (board *Board) PieceAt(pos Position) Piece { return board.pieces[pos] }

// ListMoves returns a list of possible moves for the given Piece
func (board *Board) ListMoves(piece Piece) []Position {
	if piece != board.PieceAt(piece.Pos()) {
		return nil // given piece and piece on board at the location are not the same!
	}

	return piece.ListAll(board)
}

// Reset resets the state of the board
func (board *Board) Reset() { board.pieces = make(map[Position]Piece) }

// Piece represents a single piece on the chess board.
type Piece interface {
	fmt.Stringer

	// Pos returns the piece's current position
	Pos() Position

	// SetPos update the piece's current position
	SetPos(Position)

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

// Add adds the given delta for col and row and returns a new position
func (pos Position) Add(dc, dr int) Position { col, row := pos.Split(); return Pos(col+dc, row+dr) }

// ByPosition is sort.Interface implementation that allows sorting
// positions by (col, row) in ascending order.
type ByPosition []Position

func (p ByPosition) Len() int      { return len(p) }
func (p ByPosition) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p ByPosition) Less(i, j int) bool {
	colA, rowA := p[i].Split()
	colB, rowB := p[j].Split()

	if colA == colB {
		return rowA < rowB
	} else {
		return colA < colB
	}
}
