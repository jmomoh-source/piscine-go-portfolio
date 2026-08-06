package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	inS2 := make(map[rune]bool)
	for _, r := range s2 {
		inS2[r] = true
	}

	seen := make(map[rune]bool)
	for _, r := range s1 {
		if inS2[r] && !seen[r] {
			fmt.Printf("%c", r)
			seen[r] = true
		}
	}
	fmt.Println()
}
