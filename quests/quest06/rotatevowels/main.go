package main

import (
    "os"
    "fmt"
)

func isVowel(r rune) bool {
    return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
        r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Println('\n')
        return
    }

    // Collect all vowels across arguments
    vowels := []rune{}
    for _, arg := range args {
        for _, r := range arg {
            if isVowel(r) {
                vowels = append(vowels, r)
            }
        }
    }

    // Reverse vowels
    for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
        vowels[i], vowels[j] = vowels[j], vowels[i]
    }

    // Replace vowels in order
    vi := 0
    for ai, arg := range args {
        for _, r := range arg {
            if isVowel(r) {
                fmt.Println(vowels[vi])
                vi++
            } else {
                fmt.Println(r)
            }
        }
        if ai < len(args)-1 {
            fmt.Println(' ')
        }
    }
    fmt.Println('\n')
}