package main

import "fmt"

func main() {
    for ch := 'z'; ch >= 'a'; ch-- {
        fmt.Printf("%c", ch)
    }
    fmt.Println()
}