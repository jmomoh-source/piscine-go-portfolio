package main

import (
    "fmt"
)

func main() {
    fmt.Println(ToLower("Hello! How are you?"))
}

func ToLower(s string) string {
    runes := []rune(s)
    for i, r := range runes {
        if r >= 'A' && r <= 'Z' {
            runes[i] = r + 32
        }
    }
    return string(runes)
}