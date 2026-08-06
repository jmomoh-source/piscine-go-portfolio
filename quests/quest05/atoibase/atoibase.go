package main

import (
    "fmt"
)

func main() {
    fmt.Println(AtoiBase("125", "0123456789"))
    fmt.Println(AtoiBase("1111101", "01"))
    fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
    fmt.Println(AtoiBase("uoi", "choumi"))
    fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
    if !isValidBase(base) {
        return 0
    }

    baseLen := len(base)
    indexMap := make(map[rune]int)
    for i, r := range base {
        indexMap[r] = i
    }

    result := 0
    for _, r := range s {
        value, ok := indexMap[r]
        if !ok {
            return 0
        }
        result = result*baseLen + value
    }
    return result
}

func isValidBase(base string) bool {
    if len(base) < 2 {
        return false
    }
    seen := map[rune]bool{}
    for _, r := range base {
        if r == '+' || r == '-' || seen[r] {
            return false
        }
        seen[r] = true
    }
    return true
}
