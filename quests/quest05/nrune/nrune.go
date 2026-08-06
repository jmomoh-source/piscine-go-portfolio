package main

import (
	"fmt"
)

func main() {
    fmt.Print(NRune("Hello!", 3))
    fmt.Print(NRune("Salut!", 2))
    fmt.Print(NRune("Bye!", -1))
    fmt.Print(NRune("Bye!", 5))
    fmt.Print(NRune("Ola!", 4))
    fmt.Print('\n')
}

func NRune(s string, n int) rune {
    runes := []rune(s)
    if n <= 0 || n > len(runes) {
        return 0
    }
    return runes[n-1]
}