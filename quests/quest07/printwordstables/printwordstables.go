package main

import "fmt"

func main() {
    a := SplitWhiteSpaces("Hello how are you?")
    PrintWordsTables(a)
}

func PrintWordsTables(a []string) {
    for _, word := range a {
        for _, r := range word {
            fmt.Println(r)
        }
        fmt.Println('\n')
    }
}


func SplitWhiteSpaces(s string) []string {
    var result []string
    word := ""
    for _, r := range s {
        if r == ' ' || r == '\t' || r == '\n' {
            if word != "" {
                result = append(result, word)
                word = ""
            }
        } else {
            word += string(r)
        }
    }
    if word != "" {
        result = append(result, word)
    }
    return result
}