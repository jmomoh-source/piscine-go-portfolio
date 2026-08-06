package main

import "fmt"

func main() {
    PrintNbr(-123)
    PrintNbr(0)
    PrintNbr(123)
    fmt.Println() // add newline for clarity
}

func PrintNbr(n int) {
    fmt.Print(n)
}