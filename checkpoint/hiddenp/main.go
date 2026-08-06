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

	if s1 == "" {
		fmt.Println(1)
		return
	}

	i := 0
	for j := 0; j < len(s2); j++ {
		if s2[j] == s1[i] {
			i++
			if i == len(s1) {
				break
			}
		}
	}

	if i == len(s1) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
