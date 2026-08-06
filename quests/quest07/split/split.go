package main

import (
    "fmt"
)

func main() {
    s := "HelloHAhowHAareHAyou?"
    fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s, sep string) []string {
    if sep == "" {
        return []string{s}
    }

    var result []string
    start := 0
    for i := 0; i+len(sep) <= len(s); {
        if s[i:i+len(sep)] == sep {
            result = append(result, s[start:i])
            i += len(sep)
            start = i
        } else {
            i++
        }
    }
    result = append(result, s[start:])
    return result
}