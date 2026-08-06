package main

import (
	"fmt"
	"os"
)

func validChar(r rune) bool {
	return r == '.' || (r >= '1' && r <= '9')
}

func isSafe(grid *[9][9]int, row, column, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num || grid[i][column] == num {
			return false
		}
	}
	boxRow := row / 3 * 3
	boxCol := column / 3 * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[boxRow+i][boxCol+j] == num {
				return false
			}
		}
	}
	return true
}

func solve(grid *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if grid[row][column] == 0 {
				for num := 1; num <= 9; num++ {
					if isSafe(grid, row, column, num) {
						grid[row][column] = num
						if solve(grid) {
							return true
						}
						grid[row][column] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	var grid [9][9]int

	for row := 0; row < 9; row++ {
		line := os.Args[row+1]
		if len(line) != 9 {
			fmt.Println("Error")
			return
		}
		for column, r := range line {
			if !validChar(r) {
				fmt.Println("Error")
				return
			}
			if r != '.' {
				grid[row][column] = int(r - '0')
			}
		}
	}

	if !solve(&grid) {
		fmt.Println("Error")
		return
	}

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			fmt.Print(grid[row][column])
			if column < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
