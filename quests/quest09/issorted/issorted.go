package main

import (
    "fmt"
)

func f(a, b int) int {
    return a - b
}

func main() {
    a1 := []int{0, 1, 2, 3, 4, 5}
    a2 := []int{0, 2, 1, 3}

    result1 := IsSorted(f, a1)
    result2 := IsSorted(f, a2)

    fmt.Println(result1)
    fmt.Println(result2)
}

func IsSorted(f func(a, b int) int, a []int) bool {
    if len(a) < 2 {
        return true
    }

    ascending := true
    descending := true

    for i := 0; i < len(a)-1; i++ {
        cmp := f(a[i], a[i+1])
        if cmp > 0 {
            ascending = false
        }
        if cmp < 0 {
            descending = false
        }
    }

    return ascending || descending
}