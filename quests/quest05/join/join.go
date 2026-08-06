package main

import (
    "fmt"
)

func main() {
    toConcat := []string{"Hello!", " How", " are", " you?"}
    fmt.Println(Join(toConcat, ":"))
}

func Join(strs []string, sep string) string {
    if len(strs) == 0 {
        return ""
    }
    result := strs[0]
    for i := 1; i < len(strs); i++ {
        result += sep + strs[i]
    }
    return result
}