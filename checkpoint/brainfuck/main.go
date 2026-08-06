package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	code := os.Args[1]
	tape := make([]byte, 2048)
	ptr := 0

	match := make(map[int]int)
	var stack []int
	for i, ch := range code {
		if ch == '[' {
			stack = append(stack, i)
		} else if ch == ']' {
			if len(stack) > 0 {
				start := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				match[start] = i
				match[i] = start
			}
		}
	}

	pc := 0
	for pc < len(code) {
		ch := code[pc]
		switch ch {
		case '>':
			ptr = (ptr + 1) % 2048
		case '<':
			ptr = (ptr - 1 + 2048) % 2048
		case '+':
			tape[ptr]++
		case '-':
			tape[ptr]--
		case '.':
			fmt.Printf("%c", tape[ptr])
		case '[':
			if tape[ptr] == 0 {
				pc = match[pc]
			}
		case ']':
			if tape[ptr] != 0 {
				pc = match[pc]
			}
		}
		pc++
	}
}
