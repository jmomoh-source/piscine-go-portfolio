package main

import (
    "fmt"
)

func main() {
    fmt.Println(IterativePower(4, 3))
}

func IterativePower(nb int, power int) int {
    if power < 0 {
        return 0
    }
    result := 1
    for i := 0; i < power; i++ {
        result *= nb
    }
    return result
}