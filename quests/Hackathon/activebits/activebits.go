package main

import (
    "fmt"
)

func main() {
    fmt.Println(ActiveBits(7))
}

func ActiveBits(n int) int {
    count := 0
    for n != 0 {
        if n&1 == 1 {
            count++
        }
        n >>= 1
    }
    return count
}
