package main

import "fmt"

func main() {
    IsNegative(1)
    IsNegative(0)
    IsNegative(-1)
}

func IsNegative(nb int) {
    if nb < 0 {
        fmt.Println("T")
    } else {
        fmt.Println("F")
    }
}
