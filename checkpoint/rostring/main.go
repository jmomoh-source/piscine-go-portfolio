package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	words := strings.Fields(os.Args[1])
	if len(words) == 0 {
		fmt.Println()
		return
	}

	if len(words) > 1 {
		words = append(words[1:], words[0])
	}

	fmt.Println(strings.Join(words, " "))
}
