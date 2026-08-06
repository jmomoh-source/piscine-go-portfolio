package main

import (
    "fmt"
)

func main() {
    fmt.Println(TrimAtoi("12345"))
    fmt.Println(TrimAtoi("str123ing45"))
    fmt.Println(TrimAtoi("012 345"))
    fmt.Println(TrimAtoi("Hello World!"))
    fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
    fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
    fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
    fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}

func TrimAtoi(s string) int {
    result := 0
    sign := 1
    foundDigit := false

    for _, r := range s {
        if r == '-' && !foundDigit {
            sign = -1
        }
        if r >= '0' && r <= '9' {
            foundDigit = true
            result = result*10 + int(r-'0')
        }
    }

    return result * sign
}