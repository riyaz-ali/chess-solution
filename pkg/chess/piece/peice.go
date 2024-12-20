// Package piece provides implementation of the various pieces in a chess game.
package piece

import (
	"fmt"
	"github.com/riyaz-ali/chess/pkg/chess"
	"sort"
	"sync"
)

// Pawn represents a pawn chess.Piece. A pawn can move one step in vertical forward direction only.
type Pawn struct{ pos chess.Position }

func (p *Pawn) String() string      { return fmt.Sprintf("pawn(%s)", p.pos.String()) }
func (p *Pawn) Pos() chess.Position { return p.pos }

func (p *Pawn) ListAll(board *chess.Board) []chess.Position {
	col, row := p.pos.Split()

	next := chess.Pos(col, row+1) // move to next vertical block
	if next.Valid() && board.WithinBound(next) {
		return []chess.Position{next}
	}

	return nil
}

// King represents a king chess.Piece. A king can move one step in all directions.
type King struct{ pos chess.Position }

func (k *King) String() string      { return fmt.Sprintf("king(%s)", k.pos.String()) }
func (k *King) Pos() chess.Position { return k.pos }

func (k *King) ListAll(board *chess.Board) []chess.Position {
	var result = make([]chess.Position, 0, 8) // max result could be 8 for king

	col, row := k.pos.Split()
	for i := col - 1; i <= col+1; i++ {
		for j := row - 1; j <= row+1; j++ {
			pos := chess.Pos(i, j)
			if pos.Valid() && pos != k.pos && board.WithinBound(pos) {
				result = append(result, pos)
			}
		}
	}

	return result
}

// Queen represents a queen chess.Piece. A queen can move multiple steps in all directions.
type Queen struct{ pos chess.Position }

func (q *Queen) String() string      { return fmt.Sprintf("queen(%s)", q.pos.String()) }
func (q *Queen) Pos() chess.Position { return q.pos }

func (q *Queen) ListAll(board *chess.Board) []chess.Position {
	var wg sync.WaitGroup
	var ch = make(chan chess.Position)

	// A queen can move multiple steps in all 8 directions.
	var moves = []func(chess.Position) chess.Position{
		// horizontal axis
		func(cur chess.Position) chess.Position { return cur.Add(0, 1) },
		func(cur chess.Position) chess.Position { return cur.Add(0, -1) },

		// vertical axis
		func(cur chess.Position) chess.Position { return cur.Add(1, 0) },
		func(cur chess.Position) chess.Position { return cur.Add(-1, 0) },

		// diagonal axis
		func(cur chess.Position) chess.Position { return cur.Add(-1, 1) },
		func(cur chess.Position) chess.Position { return cur.Add(-1, -1) },
		func(cur chess.Position) chess.Position { return cur.Add(1, 1) },
		func(cur chess.Position) chess.Position { return cur.Add(1, -1) },
	}

	for _, fn := range moves {
		wg.Add(1)
		go func() { defer wg.Done(); move(ch, board, q, fn) }()
	}

	// wait for all to finish and then close the channel
	// this will cause the for-loop below to end as well
	go func() { wg.Wait(); close(ch) }()

	var result []chess.Position
	for pos := range ch {
		result = append(result, pos)
	}

	sort.Sort(chess.ByPosition(result)) // sort the result to ensure consistency in output
	return result
}

// move is a utility method that emits valid positions the given piece can move into
// a given direction. The direction is determined by the next() function which returns the next
// position based on the piece's current position.
//
// This method is meant for use with pieces that can move multiple positions in a given direction, like Queen.
func move(result chan<- chess.Position, board *chess.Board, piece chess.Piece, next func(chess.Position) chess.Position) {
	// TODO(@riyaz): support non-empty board by adding a check to detect if
	//               we encounter another piece in our given trajectory.

	// we start loop from next(current) thereby skipping the current position
	for pos := next(piece.Pos()); pos.Valid() && board.WithinBound(pos); pos = next(pos) {
		result <- pos
	}
}
