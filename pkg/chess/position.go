package chess

import "fmt"

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
