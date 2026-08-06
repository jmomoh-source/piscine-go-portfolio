package main

import (
    "fmt"
    "strings"
)

func main() {
    summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
    for index, element := range ShoppingSummaryCounter(summary) {
        fmt.Println(index, "=>", element)
    }
}

func ShoppingSummaryCounter(str string) map[string]int {
    summary := make(map[string]int)
    items := strings.Fields(str)
    for _, item := range items {
        summary[item]++
    }
    return summary
}
