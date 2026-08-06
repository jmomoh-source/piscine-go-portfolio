package main

import (
    "fmt"
)

func main() {
    fmt.Println(IsPrintable("Hello"))   // true
    fmt.Println(IsPrintable("Hello\n")) // false
}

func IsPrintable(s string) bool {
    if s == "" {
        return false
    }
    for _, r := range s {
        if r < 32 || r > 126 {
            return false
        }
    }
    return true
}