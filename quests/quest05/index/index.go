package main

import (
    "fmt"
)

func main() {
    fmt.Println(Index("Hello!", "l"))
    fmt.Println(Index("Salut!", "alu"))
    fmt.Println(Index("Ola!", "hOl"))
}

func Index(s string, toFind string) int {
    if toFind == "" {
        return 0
    }
    for i := 0; i <= len(s)-len(toFind); i++ {
        if s[i:i+len(toFind)] == toFind {
            return i
        }
    }
    return -1
}