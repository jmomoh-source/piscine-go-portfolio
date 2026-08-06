package main

import (
    "fmt"
)

func main() {
    s := "Hello 78 World!    4455 /"
    nb := AlphaCount(s)
    fmt.Println(nb)
}

func AlphaCount(s string) int {
    count := 0
    for _, r := range s {
        if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
            count++
        }
    }
    return count
}