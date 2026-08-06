package main

import (
	"fmt"
	"os"
)

func isBracketed(s string) bool {
	var stack []rune
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	for _, arg := range os.Args[1:] {
		if isBracketed(arg) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}
