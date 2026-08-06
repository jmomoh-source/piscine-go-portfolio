package main

import (
	"fmt"
)

func main() {
    PrintNbrBase(125, "0123456789")
    fmt.Println('\n')
    PrintNbrBase(-125, "01")
    fmt.Println('\n')
    PrintNbrBase(125, "0123456789ABCDEF")
    fmt.Println('\n')
    PrintNbrBase(-125, "choumi")
    fmt.Println('\n')
    PrintNbrBase(125, "aa")
    fmt.Println('\n')
}

func PrintNbrBase(nbr int, base string) {
    // validate base
    if !isValidBase(base) {
        fmt.Println('N')
        fmt.Println('V')
        return
    }

    if nbr == 0 {
        fmt.Println(rune(base[0]))
        return
    }

    if nbr < 0 {
        fmt.Println('-')
        nbr = -nbr
    }

    b := len(base)
    digits := []rune{}
    for nbr > 0 {
        digits = append([]rune{rune(base[nbr%b])}, digits...)
        nbr /= b
    }

    for _, d := range digits {
        fmt.Println(d)
    }
}

func isValidBase(base string) bool {
    if len(base) < 2 {
        return false
    }
    seen := map[rune]bool{}
    for _, r := range base {
        if r == '+' || r == '-' || seen[r] {
            return false
        }
        seen[r] = true
    }
    return true
}