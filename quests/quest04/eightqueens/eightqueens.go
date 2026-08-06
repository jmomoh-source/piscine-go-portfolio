package main

import "fmt"

func main() {
    EightQueens()
}

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