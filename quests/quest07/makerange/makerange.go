package main

import (
    "fmt"
)

func main() {
    fmt.Println(MakeRange(5, 10))
    fmt.Println(MakeRange(10, 5))
}

func MakeRange(min, max int) []int {
    if min >= max {
        return nil
    }
    result := make([]int, max-min)
    for i := min; i < max; i++ {
        result[i-min] = i
    }
    return result
}