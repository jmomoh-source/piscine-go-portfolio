package main

import (
    "fmt"
)

func main() {
    result := ConvertBase("101011", "01", "0123456789")
    fmt.Println(result)
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
    // Step 1: Convert nbr from baseFrom to decimal
    baseFromLen := len(baseFrom)
    value := 0
    for _, r := range nbr {
        digit := indexOfRune(baseFrom, r)
        value = value*baseFromLen + digit
    }

    // Step 2: Convert decimal value to baseTo
    if value == 0 {
        return string(baseTo[0])
    }

    baseToLen := len(baseTo)
    result := ""
    for value > 0 {
        digit := value % baseToLen
        result = string(baseTo[digit]) + result
        value /= baseToLen
    }
    return result
}

func indexOfRune(s string, r rune) int {
    for i, v := range s {
        if v == r {
            return i
        }
    }
    return -1
}