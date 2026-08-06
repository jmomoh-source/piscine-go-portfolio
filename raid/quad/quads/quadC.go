package quads

import "fmt"

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 && col == 0 {
				// Top-left corner
				fmt.Print("A")
			} else if row == 0 && col == x-1 {
				// Top-right corner
				fmt.Print("A")
			} else if row == y-1 && col == 0 {
				// Bottom-left corner
				fmt.Print("C")
			} else if row == y-1 && col == x-1 {
				// Bottom-right corner
				fmt.Print("C")
			} else if row == 0 || row == y-1 || col == 0 || col == x-1 {
				// Edges
				fmt.Print("B")
			} else {
				// Interior
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
