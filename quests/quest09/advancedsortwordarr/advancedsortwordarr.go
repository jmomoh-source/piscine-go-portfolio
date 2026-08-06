package main

import (
    "fmt"
)

func Compare(a, b string) int {
    return int([]rune(a)[0] - []rune(b)[0])
}

func main() {
    result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
    AdvancedSortWordArr(result, Compare)
    fmt.Println(result)
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
    for i := 0; i < len(a)-1; i++ {
        for j := i + 1; j < len(a); j++ {
            if f(a[i], a[j]) > 0 {
                a[i], a[j] = a[j], a[i]
            }
        }
    }
}