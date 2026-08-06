# Quest08 — cat

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **file concatenation and standard input handling** in Go.  
The task: write a program that behaves like a simplified version of the Unix `cat` command.

Rules:
- If no arguments are provided, read from standard input and print back to standard output.
- If one or more arguments are provided, read each file sequentially and print its contents.
- If a file cannot be opened, print an error message and exit with a non‑zero status.

## Instructions
- Directory: `cat`
- File to submit: `main.go`
- Allowed packages: `os`, `io`, `bufio`, `strings`, `z01.PrintRune`

## Implementation
`main.go`:
```go
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
```

### Explanation
- If no arguments are passed, use `bufio.NewReader(os.Stdin)` to read from standard input until EOF.
- If arguments are passed, loop through each filename:
  - Open the file with `os.Open`.
  - Read its contents with `io.ReadAll`.
  - Print the contents to standard output.
- If any error occurs, print it to `stderr` and exit with status code `1`.

## Usage
Example test program:
```bash
$ echo '"Programming is a skill best acquired by practice and example rather than from books" by Alan Turing' > quest8.txt
$ go run . quest8.txt
"Programming is a skill best acquired by practice and example rather than from books" by Alan Turing

$ go run . quest8.txt quest8T.txt
"Programming is a skill best acquired by practice and example rather than from books" by Alan Turing
"Alan Mathison Turing was an English mathematician, computer scientist, logician, cryptanalyst..."
```

Error handling:
```bash
$ go run . abc
ERROR: open abc: no such file or directory
exit status 1
```

Standard input:
```bash
$ cat quest8.txt | go run .
"Programming is a skill best acquired by practice and example rather than from books" by Alan Turing

$ go run .
Hello
Hello
^C
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
        data, _ := io.ReadAll(os.Stdin)
        fmt.Print(string(data))
        return
    }

    for _, filename := range args {
        content, err := os.ReadFile(filename)
        if err != nil {
            fmt.Println("ERROR:", err)
            os.Exit(1)
        }
        fmt.Print(string(content))
    }
}
```
⚠️ Note: `os.ReadFile` is simpler and idiomatic.  
Your Piscine solution demonstrates manual file reading and stdin handling for deeper understanding.

## Skills Practiced
- Command‑line argument handling
- File opening and reading
- Standard input reading
- Error handling
- Printing to stdout and stderr

## Notes
- This exercise demonstrates how to build a simplified version of Unix utilities.
- For production code, prefer `os.ReadFile` or `io.Copy` for simplicity and efficiency.

## Resources
- Go `os.Open` — Official Docs (go.dev in Bing)  
- Go `io.ReadAll` — Official Docs (go.dev in Bing)  
- Go `bufio.Scanner` — Official Docs (go.dev in Bing) 