package main

import (
	"fmt"
	"os"
)

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	word := os.Args[1]
	vIdx := -1
	for i := 0; i < len(word); i++ {
		if isVowel(word[i]) {
			vIdx = i
			break
		}
	}

	if vIdx == -1 {
		fmt.Println("No vowels")
		return
	}

	if vIdx == 0 {
		fmt.Println(word + "ay")
	} else {
		fmt.Println(word[vIdx:] + word[:vIdx] + "ay")
	}
}
