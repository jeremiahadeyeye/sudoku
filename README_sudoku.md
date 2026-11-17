# Sudoku Solver

A command-line Sudoku solver written in Go. It takes a 9x9 puzzle as nine command-line arguments, solves it with backtracking, and prints the completed grid. If the input is invalid or the puzzle has no solution, it prints `Error`.

## Requirements

- Go 1.22.2 or newer
- No external dependencies (standard library only: `fmt`, `os`)

## Usage

Pass the puzzle as 9 arguments, one per row. Each row must be exactly 9 characters: digits `1`-`9` for given cells and `.` for empty cells.

```bash
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
```

Or build a binary first:

```bash
go build -o sudoku
./sudoku ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
```

### Output

The solved grid is printed as 9 lines, with digits separated by single spaces:

```
3 9 6 2 4 5 7 8 1
1 7 8 3 6 9 5 2 4
5 2 4 8 1 7 3 9 6
2 8 7 9 5 1 6 4 3
9 3 1 4 8 6 2 7 5
4 6 5 7 2 3 9 1 8
7 1 2 6 3 8 4 5 9
6 5 9 1 7 4 8 3 2
8 4 3 5 9 2 1 6 7
```

### Error cases

The program prints `Error` (and nothing else) when:

| Case | Example |
|------|---------|
| Not exactly 9 arguments | `./sudoku 123` |
| A row is not exactly 9 characters | `"12345678"` |
| A row has a character other than `1`-`9` or `.` | `"12345678a"`, `"123456780"` |
| The given digits repeat in a row, column or 3x3 box | two `5`s in the same row |
| The puzzle is valid but has no solution | contradictory clues |

## How it works

1. **Argument check:** exactly 9 arguments are required.
2. **Validation (`isValidInput`):** checks row length and allowed characters, then checks that the starting digits have no duplicates in any row, column or 3x3 box.
3. **Conversion (`convertToGrid`):** turns the strings into a 9x9 `int` grid, using `-1` for empty cells.
4. **Solving (`solveSudoku`):** classic recursive backtracking. It finds the first empty cell (`findEmptyCell`), tries digits 1-9 that pass `isValidMove`, recurses, and undoes the placement if it hits a dead end.
5. **Output (`printGrid`):** prints the solved grid.

If a puzzle has more than one solution, the solver prints the first one it finds (it does not report that the solution is non-unique).

## Project structure

```
sudoku/
├── go.mod       # module "sudoku", Go 1.22.2
└── sudoku.go    # entire program (package main)
```
