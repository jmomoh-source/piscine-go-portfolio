package main

import (
    "os"
    "fmt"
)

func main() {
    for _, arg := range os.Args[1:] {
        for _, r := range arg {
            fmt.Print(r)
        }
        fmt.Print('\n')
    }
}