package main

import "fmt"

func main() {
    DescendComb()
}

func DescendComb() {
    for a := 99; a >= 0; a-- {
        for b := a - 1; b >= 0; b-- {
            fmt.Printf("%02d %02d", a, b)
            if !(a == 1 && b == 0) {
                fmt.Print(", ")
            }
        }
    }
    fmt.Println()
}
