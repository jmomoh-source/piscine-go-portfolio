package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 1 {
		return
	}

	var factors []int
	divisor := 2
	for n > 1 {
		if n%divisor == 0 {
			factors = append(factors, divisor)
			n /= divisor
		} else {
			divisor++
		}
	}

	for i, f := range factors {
		if i > 0 {
			fmt.Print("*")
		}
		fmt.Print(f)
	}
	fmt.Println()
}
