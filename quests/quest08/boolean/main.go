package main

import (
    "os"
    "fmt"
)

const EvenMsg = "I have an even number of arguments"
const OddMsg = "I have an odd number of arguments"

func printStr(s string) {
    for _, r := range s {
        fmt.Println(r)
    }
    fmt.Println('\n')
}

func isEven(nbr int) bool {
    return nbr%2 == 0
}

func main() {
    lengthOfArg := len(os.Args[1:])
    if isEven(lengthOfArg) {
        printStr(EvenMsg)
    } else {
        printStr(OddMsg)
    }
}