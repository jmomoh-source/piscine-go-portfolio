package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arrStr := os.Args[1]
	targetStr := os.Args[2]

	if !strings.HasPrefix(arrStr, "[") || !strings.HasSuffix(arrStr, "]") {
		fmt.Println("Invalid input.")
		return
	}

	target, err := strconv.Atoi(targetStr)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	content := strings.TrimSpace(arrStr[1 : len(arrStr)-1])
	if content == "" {
		fmt.Println("No pairs found.")
		return
	}

	parts := strings.Split(content, ",")
	nums := make([]int, len(parts))
	for i, p := range parts {
		token := strings.TrimSpace(p)
		v, err := strconv.Atoi(token)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", token)
			return
		}
		nums[i] = v
	}

	var pairs [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				pairs = append(pairs, []int{i, j})
			}
		}
	}

	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
		return
	}

	fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
}
