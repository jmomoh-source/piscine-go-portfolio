package main

import (
    "fmt"
)

func main() {
    fmt.Print(JumpOver("1010101010"))
    fmt.Print(JumpOver(""))
    fmt.Print(JumpOver("t w e l v e"))
    fmt.Print(JumpOver("12"))
}

func JumpOver(str string) string {
    if len(str) < 3 {
        return "\n"
    }
    result := []rune{}
    for i := 2; i < len(str); i += 3 {
        result = append(result, rune(str[i]))
    }
    return string(result) + "\n"
}