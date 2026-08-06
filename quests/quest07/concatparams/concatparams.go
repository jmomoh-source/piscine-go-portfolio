package main

import (
    "fmt"
)

func main() {
    test := []string{"Hello", "how", "are", "you?"}
    fmt.Println(ConcatParams(test))
}

func ConcatParams(args []string) string {
    if len(args) == 0 {
        return ""
    }
    result := ""
    for i, s := range args {
        result += s
        if i < len(args)-1 {
            result += "\n"
        }
    }
    return result
}