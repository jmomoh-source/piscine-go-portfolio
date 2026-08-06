package main

import (
    "fmt"
)

func main() {
    fmt.Println(ToUpper("Hello! How are you?"))
}


func ToUpper(s string) string {
    runes := []rune(s)
    for i, r := range runes {
        if r >= 'a' && r <= 'z' {
            runes[i] = r - 32
        }
    }
    return string(runes)
}