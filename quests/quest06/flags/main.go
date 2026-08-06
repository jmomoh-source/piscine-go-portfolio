package main

import (
    "fmt"
    "os"
    "sort"
    "strings"
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
        printHelp()
        return
    }

    insertStr := ""
    order := false
    var mainArg string

    for _, arg := range args {
        if strings.HasPrefix(arg, "--insert=") {
            insertStr = strings.TrimPrefix(arg, "--insert=")
        } else if strings.HasPrefix(arg, "-i=") {
            insertStr = strings.TrimPrefix(arg, "-i=")
        } else if arg == "--order" || arg == "-o" {
            order = true
        } else {
            mainArg = arg
        }
    }

    result := mainArg + insertStr
    if order {
        runes := []rune(result)
        sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
        result = string(runes)
    }

    for _, r := range result {
        fmt.Println(r)
    }
    fmt.Println('\n')
}

func printHelp() {
    fmt.Println("--insert")
    fmt.Println("  -i")
    fmt.Println("         This flag inserts the string into the string passed as argument.")
    fmt.Println("--order")
    fmt.Println("  -o")
    fmt.Println("         This flag will behave like a boolean, if it is called it will order the argument.")
}