package main

import (
    "fmt"
)

func main() {
    fmt.Println(FindNextPrime(5)) // 5
    fmt.Println(FindNextPrime(4)) // 5
}

func IsPrime(nb int) bool {
    if nb <= 1 {
        return false
    }
    if nb == 2 {
        return true
    }
    if nb%2 == 0 {
        return false
    }
    for i := 3; i*i <= nb; i += 2 {
        if nb%i == 0 {
            return false
        }
    }
    return true
}

func FindNextPrime(nb int) int {
    if nb <= 2 {
        return 2
    }
    for {
        if IsPrime(nb) {
            return nb
        }
        nb++
    }
}