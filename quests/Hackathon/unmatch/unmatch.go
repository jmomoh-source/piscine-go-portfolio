package main

import (
    "fmt"
)

func main() {
    a := []int{1, 2, 3, 1, 2, 3, 4}
    unmatch := Unmatch(a)
    fmt.Println(unmatch) // 4
}

func Unmatch(a []int) int {
    counts := make(map[int]int)
    for _, v := range a {
        counts[v]++
    }
    for k, v := range counts {
        if v%2 != 0 {
            return k
        }
    }
    return -1
}
