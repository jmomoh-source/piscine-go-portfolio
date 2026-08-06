package main

import (
    "os"
    "fmt"
)

func main() {
    args := os.Args[1:]
    for i := len(args) - 1; i >= 0; i-- {
        for _, r := range args[i] {
            fmt.Print(r)
        }
        fmt.Print('\n')
    }
}