# Quest04 — eightqueens

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **backtracking recursion** in Go.  
The task: write a function `EightQueens` that prints all solutions to the classic **Eight Queens puzzle**.

Rules:
- Place 8 queens on a chessboard so that none attack each other.
- Each solution is printed as a sequence of 8 digits:
  - Each digit represents the row position of the queen in that column.
  - Indexing starts at `1`.
  - Solutions are printed in ascending order.
- Recursion must be used.

## Instructions
- File to submit: `eightqueens.go` (inside the `piscine` package)
- Expected function signature:
```go
func EightQueens() {
}
```

Output must look like:
```
bash
15863724
16837425
17468253
...
```

## Implementation
eightqueens.go:
```
go
package piscine

import "fmt"

func EightQueens() {
    var solve func(col int, board []int)
    solve = func(col int, board []int) {
        if col == 8 {
            for _, v := range board {
                fmt.Print(v)
            }
            fmt.Println()
            return
        }
        for row := 1; row <= 8; row++ {
            if isSafe(board, col, row) {
                board[col] = row
                solve(col+1, board)
            }
        }
    }
    solve(0, make([]int, 8))
}

func isSafe(board []int, col, row int) bool {
    for c := 0; c < col; c++ {
        r := board[c]
        if r == row || abs(r-row) == abs(c-col) {
            return false
        }
    }
    return true
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
```

## Explanation
solve(col, board) → recursive function that tries to place queens column by column.

isSafe(board, col, row) → checks if placing a queen at (col, row) is valid:

No two queens share the same row.

No two queens share the same diagonal.

- Base case: when all 8 columns are filled, print the solution.

Solutions are printed in ascending order naturally by iterating rows from 1 to 8.

## Usage
Example test program:
```
go
package main

import "piscine"

func main() {
    piscine.EightQueens()
}
```

Run it:
```
bash
go run .
```

Expected output (first few lines):
```
text
15863724
16837425
17468253
...
```

## Skills Practiced
Recursive backtracking

Constraint satisfaction problems

Chessboard representation

Efficient pruning of invalid states

## Notes
The Eight Queens puzzle is a classic example of recursion and backtracking.

Solutions are printed in ascending order due to the natural iteration order.

Output must be exactly as specified, with no extra spaces or characters.

## Resources
Eight Queens Puzzle — MathWorld (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)

Go Recursion — Tour of Go (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)

Backtracking Algorithms Overview (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)