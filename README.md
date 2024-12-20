# Chess ♟️

This repository contains solution to the coding assignment problem found in [problem.pdf](./problem.pdf)

## Overview

The program emulates an `8x8` chess board with six unique types of pieces - King, Queen, Rook, Bishop, Horse and Pawn - each with their own unique types of movements.

It expects, as input, a string of the form **`<type>, <position>`**, where `<type>` is the type of the pawn, and `<position>` is its position
on the board in the form of `<col><row>` (eg. `A8`).

Given the input, it generates a list of all possible cells in which the chess piece can move from its current position.

## Assumptions

1. The size of the board is fixed (an `8x8` grid) and not configurable at runtime
2. Movements for only three types of pieces - Pawn, King and Queen - are to be simulated
3. The tool does not need to implement a "chess solver" functionality
4. The tool does not need to maintain a "current state" or "scorecard-like" functionality
5. The input will only be the position of a single piece at a time, ie. it won't be an array of positions of multiple pieces