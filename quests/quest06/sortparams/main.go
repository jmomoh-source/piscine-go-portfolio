package main

import (
    "os"
    "fmt"
)

func main() {
    args := os.Args[1:]

    // Simple bubble sort for ASCII order
    for i := 0; i < len(args); i++ {
        for j := i + 1; j < len(args); j++ {
            if args[i] > args[j] {
                args[i], args[j] = args[j], args[i]
            }
        }
    }

    for _, arg := range args {
        for _, r := range arg {
            fmt.Print(r)
        }
        fmt.Print('\n')
    }
}