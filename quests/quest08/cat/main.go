package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func main() {
    args := os.Args[1:]

    if len(args) == 0 {
        // Read from stdin
        reader := bufio.NewReader(os.Stdin)
        for {
            line, err := reader.ReadString('\n')
            if err == io.EOF {
                break
            }
            if err != nil {
                fmt.Fprintln(os.Stderr, "ERROR:", err)
                os.Exit(1)
            }
            fmt.Print(line)
        }
        return
    }

    // Read from files
    for _, filename := range args {
        file, err := os.Open(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
            os.Exit(1)
        }
        defer file.Close()

        content, err := io.ReadAll(file)
        if err != nil {
            fmt.Fprintln(os.Stderr, "ERROR:", err)
            os.Exit(1)
        }
        fmt.Print(string(content))
    }
}