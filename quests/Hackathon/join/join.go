package main

import (
    "fmt"
	"strings"
)

func main() {
    toConcat := []string{"Hello!", " How", " are", " you?"}
    fmt.Println(Join(toConcat, ":"))
}

func Join(strs []string, sep string) string {
    return strings.Join(strs, sep)
}