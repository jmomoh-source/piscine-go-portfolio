package main

import (
    "fmt"
)

func main() {
    fmt.Println(StringToIntSlice("A quick brown fox jumps over the lazy dog"))
    fmt.Println(StringToIntSlice("Converted this string into an int"))
    fmt.Println(StringToIntSlice("hello THERE"))
}

func StringToIntSlice(str string) []int {
    result := []int{}
    for _, r := range str {
        result = append(result, int(r))
    }
    return result
}
