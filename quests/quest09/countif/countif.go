package main

import (
    "fmt"
)

func main() {
    tab1 := []string{"Hello", "how", "are", "you"}
    tab2 := []string{"This", "1", "is", "4", "you"}

    answer1 := CountIf(IsNumeric, tab1)
    answer2 := CountIf(IsNumeric, tab2)

    fmt.Println(answer1)
    fmt.Println(answer2)
}

func CountIf(f func(string) bool, tab []string) int {
    count := 0
    for _, v := range tab {
        if f(v) {
            count++
        }
    }
    return count
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
