package main

import (
	"fmt"
)

func main() {
    fmt.Print(LastRune("Hello!"))
    fmt.Print(LastRune("Salut!"))
    fmt.Print(LastRune("Ola!"))
    fmt.Print('\n')
}

func LastRune(s string) rune {
    runes := []rune(s)
    return runes[len(runes)-1]
}