package main

import (
    "fmt"
)

func main() {
    a := 0
    b := 1
    Swap(&a, &b)
    fmt.Println(a)
    fmt.Println(b)
}

func Swap(a *int, b *int) {
    temp := *a
    *a = *b
    *b = temp
}