package main

import (
    "fmt"
)

const N = 6

func main() {
    a := make([]string, N)
    a[0] = "a"
    a[2] = "b"
    a[4] = "c"

    for _, v := range a {
        fmt.Println(v)
    }

    fmt.Println("Size after compacting:", Compact(&a))

    for _, v := range a {
        fmt.Println(v)
    }
}

func Compact(ptr *[]string) int {
    original := *ptr
    compacted := []string{}
    for _, v := range original {
        if v != "" {
            compacted = append(compacted, v)
        }
    }
    *ptr = compacted
    return len(compacted)
}
