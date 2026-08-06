package main

import "fmt"

func main() {
    PrintStr("Hello World!")
}

func PrintStr(s string) {
    for _, r := range s {
        fmt.Print(r)
    }
}