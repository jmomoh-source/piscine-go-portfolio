# Quest08 — displayfile

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **basic file I/O** in Go.  
The task: write a program that displays the content of a file given as an argument.

Rules:
- If no argument is provided, print `File name missing`.
- If more than one argument is provided, print `Too many arguments`.
- Otherwise, open the file and print its contents to standard output.

## Instructions
- Directory: `displayfile`
- File to submit: `main.go`
- Allowed packages: `fmt`, `os`, `io`

## Implementation
`main.go`:
```go
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    args := os.Args[1:]

    if len(args) == 0 {
        fmt.Println("File name missing")
        return
    }
    if len(args) > 1 {
        fmt.Println("Too many arguments")
        return
    }

    file, err := os.Open(args[0])
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    content, err := io.ReadAll(file)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    fmt.Print(string(content))
}
```

### Explanation
- Use `os.Args[1:]` to get command‑line arguments.
- Handle cases:
  - No arguments → print `"File name missing"`.
  - More than one argument → print `"Too many arguments"`.
- Open the file with `os.Open`.
- Read its contents with `io.ReadAll`.
- Print the contents to standard output.

## Usage
Example test program:
```bash
$ go run .
File name missing

$ echo 'Almost there!!' > quest8.txt
$ go run . quest8.txt main.go
Too many arguments

$ go run . quest8.txt
Almost there!!
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with `os.ReadFile`:
```go
import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Println("File name missing")
        return
    }
    if len(args) > 1 {
        fmt.Println("Too many arguments")
        return
    }

    content, err := os.ReadFile(args[0])
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Print(string(content))
}
```
⚠️ Note: `os.ReadFile` is simpler and idiomatic.  
Your Piscine solution demonstrates manual file opening and reading for deeper understanding.

## Skills Practiced
- Command‑line argument handling
- File opening and closing
- Reading file contents
- Error handling

## Notes
- This exercise demonstrates manual file I/O.
- For production code, prefer `os.ReadFile` for simplicity.

## Resources
- Go `os.Open` — Official Docs (go.dev in Bing)  
- Go `io.ReadAll` — Official Docs (go.dev in Bing)