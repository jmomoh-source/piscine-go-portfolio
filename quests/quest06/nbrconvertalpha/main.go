package main

import (
    "os"
    "strconv"
    "fmt"
)

func main() {
    args := os.Args[1:]
    upper := false

    if len(args) > 0 && args[0] == "--upper" {
        upper = true
        args = args[1:]
    }

    for _, arg := range args {
        n, err := strconv.Atoi(arg)
        if err != nil || n < 1 || n > 26 {
            fmt.Print(' ')
            continue
        }
        r := rune('a' + n - 1)
        if upper {
            r = rune('A' + n - 1)
        }
        fmt.Print(r)
    }
    fmt.Print('\n')
}