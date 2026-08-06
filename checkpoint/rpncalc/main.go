package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	tokens := strings.Fields(os.Args[1])
	if len(tokens) == 0 {
		fmt.Println("Error")
		return
	}

	var stack []int
	for _, tok := range tokens {
		if val, err := strconv.Atoi(tok); err == nil {
			stack = append(stack, val)
		} else if tok == "+" || tok == "-" || tok == "*" || tok == "/" || tok == "%" {
			if len(stack) < 2 {
				fmt.Println("Error")
				return
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			res := 0
			switch tok {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				if b == 0 {
					fmt.Println("Error")
					return
				}
				res = a / b
			case "%":
				if b == 0 {
					fmt.Println("Error")
					return
				}
				res = a % b
			}
			stack = append(stack, res)
		} else {
			fmt.Println("Error")
			return
		}
	}

	if len(stack) != 1 {
		fmt.Println("Error")
		return
	}

	fmt.Println(stack[0])
}
