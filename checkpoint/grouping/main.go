package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func cleanWord(w string) string {
	return strings.Trim(w, ",.!?…:;\"'")
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	pattern := os.Args[1]
	text := os.Args[2]

	if text == "" || !strings.HasPrefix(pattern, "(") || !strings.HasSuffix(pattern, ")") {
		return
	}

	inner := pattern[1 : len(pattern)-1]
	if inner == "" {
		return
	}

	re, err := regexp.Compile(inner)
	if err != nil {
		return
	}

	rawWords := strings.Fields(text)
	matchCount := 0

	for _, rw := range rawWords {
		word := cleanWord(rw)
		if word == "" {
			continue
		}
		matches := re.FindAllString(word, -1)
		for range matches {
			matchCount++
			fmt.Printf("%d: %s\n", matchCount, word)
		}
	}
}
