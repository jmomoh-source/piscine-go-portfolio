package main

import (
    "fmt"
)

func main() {
    result := Rot14("Hello! How are You?")
    for _, r := range result {
        fmt.Print(r)
    }
    fmt.Print('\n')
}

func Rot14(s string) string {
    result := []rune{}
    for _, r := range s {
        if r >= 'a' && r <= 'z' {
            result = append(result, 'a'+(r-'a'+14)%26)
        } else if r >= 'A' && r <= 'Z' {
            result = append(result, 'A'+(r-'A'+14)%26)
        } else {
            result = append(result, r)
        }
    }
    return string(result)
}
