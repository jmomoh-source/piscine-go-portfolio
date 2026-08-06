package main

import (
    "os"
    "fmt"
)

func main() {
    progName := os.Args[0]
    // Extract only the last part of the path
    nameRunes := []rune(progName)
    lastSlash := -1
    for i, r := range nameRunes {
        if r == '/' {
            lastSlash = i
        }
    }
    name := nameRunes[lastSlash+1:]

    for _, r := range name {
        fmt.Print(r)
    }
    fmt.Print('\n')
}