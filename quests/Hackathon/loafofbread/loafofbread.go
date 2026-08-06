package main

import (
    "fmt"
)

func main() {
    fmt.Print(LoafOfBread("deliciousbread"))
    fmt.Print(LoafOfBread("This is a loaf of bread"))
    fmt.Print(LoafOfBread("loaf"))
}

func LoafOfBread(str string) string {
    // Remove spaces when counting
    runes := []rune{}
    for _, r := range str {
        if r != ' ' {
            runes = append(runes, r)
        }
    }

    if len(runes) < 5 {
        return "Invalid Output\n"
    }

    result := []rune{}
    count := 0
    for i := 0; i < len(runes); i++ {
        result = append(result, runes[i])
        count++
        if count == 5 {
            // skip next character if exists
            if i+1 < len(runes) {
                i++
            }
            result = append(result, ' ')
            count = 0
        }
    }

    return string(result) + "\n"
}
