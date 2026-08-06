package main

import (
    "fmt"
)

func main() {
    slice := []string{"Pineapple", "Honey", "Mushroom", "Tea", "Pepper", "Milk"}
    fmt.Println(ShoppingListSort(slice))
}

func ShoppingListSort(slice []string) []string {
    // Simple bubble sort by length
    n := len(slice)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if len(slice[j]) > len(slice[j+1]) {
                slice[j], slice[j+1] = slice[j+1], slice[j]
            }
        }
    }
    return slice
}
