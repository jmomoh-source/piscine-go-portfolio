package main

import (
    "fmt"
)

func main() {
    fmt.Println(IsUpper("HELLO"))   // true
    fmt.Println(IsUpper("Hello"))   // false
    fmt.Println(IsUpper("123"))     // false
    fmt.Println(IsUpper(""))        // false
}

func IsUpper(s string) bool {
    if s == "" {
        return false
    }
    for _, r := range s {
        if r < 'A' || r > 'Z' {
            return false
        }
    }
    return true
}