package quads

import "fmt"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if (row == 0 || row == y-1) && (col == 0 || col == x-1) {
				// Corners
				fmt.Print("o")
			} else if row == 0 || row == y-1 {
				// Top and bottom edges
				fmt.Print("-")
			} else if col == 0 || col == x-1 {
				// Left and right edges
				fmt.Print("|")
			} else {
				// Interior
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
