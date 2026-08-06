package main

import "fmt"


func main() {
    fmt.Println(FirstRune("Hello!"))
    fmt.Println(FirstRune("Salut!"))
    fmt.Println(FirstRune("Ola!"))
    fmt.Print('\n')
}

func FirstRune(s string) rune {
    runes := []rune(s)
    return runes[0]
}