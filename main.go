package main

import (
	"bufio"
	"fmt"
	"github.com/riyaz-ali/chess/pkg/chess"
	"github.com/riyaz-ali/chess/pkg/chess/piece"
	"io"
	"log"
	"os"
	"strings"
)

func run(input string, out io.Writer) (err error) {
	var board = chess.NewBoard(8) // create a new 8x8 size board

	parts := strings.SplitN(input, ",", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid input: %s", input)
	}

	kind, pos := piece.Kind(strings.TrimSpace(parts[0])), chess.Position(strings.TrimSpace(parts[1]))

	// create new piece and place it on the board
	if chessPiece, err := piece.New(kind, pos); err != nil {
		return fmt.Errorf("failed to create piece: %s", err.Error())
	} else {
		if err = board.Move(chessPiece, pos); err != nil {
			return fmt.Errorf("failed to move piece: %s", err.Error())
		}
	}

	var positions = board.ListMoves(board.PieceAt(pos))
	var result = Join(positions, ", ")

	_, _ = fmt.Fprint(out, result) // write to output stream
	return nil
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	_, _ = fmt.Fprint(os.Stdout, "Enter (Type, Position) of the piece: ")

	input, err := stdin.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read input: %s", err.Error())
	}

	if err = run(input, os.Stdout); err != nil {
		log.Fatalf("failed to run: %s", err.Error())
	}
}
