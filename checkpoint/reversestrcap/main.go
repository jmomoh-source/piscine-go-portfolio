package main

import (
	"fmt"
	"os"
)

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}
	return c
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func process(s string) string {
	res := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if isLetter(ch) {
			isLast := (i+1 == len(s)) || s[i+1] == ' '
			if isLast {
				res[i] = toUpper(ch)
			} else {
				res[i] = toLower(ch)
			}
		} else {
			res[i] = ch
		}
	}
	return string(res)
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	for _, arg := range os.Args[1:] {
		fmt.Println(process(arg))
	}
}
