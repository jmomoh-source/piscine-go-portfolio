package main

import (
    "fmt"
)

func main() {
    fmt.Println(BasicAtoi2("12345"))
    fmt.Println(BasicAtoi2("0000000012345"))
    fmt.Println(BasicAtoi2("012 345"))
    fmt.Println(BasicAtoi2("Hello World!"))
}

func BasicAtoi2(s string) int {
    result := 0
    for _, r := range s {
        if r < '0' || r > '9' {
            return 0
        }
        result = result*10 + int(r-'0')
    }
    return result
}