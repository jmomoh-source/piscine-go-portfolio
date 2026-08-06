package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
    if len(args) < 2 || args[0] != "-c" {
        fmt.Println("Usage: go run . -c <number> <file> [files...]")
        os.Exit(1)
    }

    // Parse number of bytes
    var n int
    _, err := fmt.Sscanf(args[1], "%d", &n)
    if err != nil || n <= 0 {
        fmt.Println("Invalid number of bytes")
        os.Exit(1)
    }

    files := args[2:]
    exitCode := 0

    for i, filename := range files {
        data, err := os.ReadFile(filename)
        if err != nil {
            fmt.Println("open", filename+":", err)
            exitCode = 1
        } else {
            if len(files) > 1 {
                if i > 0 {
                    fmt.Println()
                }
                fmt.Printf("==> %s <==\n", filename)
            }
            if len(data) < n {
                fmt.Print(string(data))
            } else {
                fmt.Print(string(data[len(data)-n:]))
            }
        }
    }

    os.Exit(exitCode)
}