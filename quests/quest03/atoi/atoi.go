package main

import (
    "fmt"
)

func main() {
    fmt.Println(Atoi("12345"))
    fmt.Println(Atoi("0000000012345"))
    fmt.Println(Atoi("012 345"))
    fmt.Println(Atoi("Hello World!"))
    fmt.Println(Atoi("+1234"))
    fmt.Println(Atoi("-1234"))
    fmt.Println(Atoi("++1234"))
    fmt.Println(Atoi("--1234"))
}

func Atoi(s string) int {
    if len(s) == 0 {
        return 0
    }

    sign := 1
    start := 0

    if s[0] == '+' {
        start = 1
    } else if s[0] == '-' {
        sign = -1
        start = 1
    }

    result := 0
    for _, r := range s[start:] {
        if r < '0' || r > '9' {
            return 0
        }
        result = result*10 + int(r-'0')
    }

    return sign * result
}