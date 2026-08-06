package main

import (
    "fmt"
)

func main() {
    fmt.Println(IsNumeric("010203"))   // true
    fmt.Println(IsNumeric("01,02,03")) // false
}

func IsNumeric(s string) bool {
    if s == "" {
        return false
    }
    for _, r := range s {
        if r < '0' || r > '9' {
            return false
        }
    }
    return true
}