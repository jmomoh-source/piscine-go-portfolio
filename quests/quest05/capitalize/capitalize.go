package main

import (
    "fmt"
	"unicode"
)

func main() {
    fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}

func Capitalize(s string) string {
    runes := []rune(s)
    newWord := true

    for i, r := range runes {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            if newWord {
                runes[i] = unicode.ToUpper(r)
                newWord = false
            } else {
                runes[i] = unicode.ToLower(r)
            }
        } else {
            newWord = true
        }
    }
    return string(runes)
}