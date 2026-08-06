package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	seen := make(map[rune]bool)
	for _, r := range s1 {
		if !seen[r] {
			fmt.Printf("%c", r)
			seen[r] = true
		}
	}
	for _, r := range s2 {
		if !seen[r] {
			fmt.Printf("%c", r)
			seen[r] = true
		}
	}
	fmt.Println()
}
