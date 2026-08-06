package main

import (
    "fmt"
)

func main() {
    fmt.Println(IsAlpha("Hello! How are you?")) // false
    fmt.Println(IsAlpha("HelloHowareyou"))      // true
    fmt.Println(IsAlpha("What's this 4?"))      // false
    fmt.Println(IsAlpha("Whatsthis4"))          // true
}

func IsAlpha(s string) bool {
    for _, r := range s {
        if !((r >= 'A' && r <= 'Z') ||
             (r >= 'a' && r <= 'z') ||
             (r >= '0' && r <= '9')) {
            return false
        }
    }
    return true
}