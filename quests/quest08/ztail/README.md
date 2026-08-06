# Quest08 — ztail

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **file reading with options** in Go.  
The task: write a program that behaves like a simplified version of the Unix `tail` command.

Rules:
- The only option to handle is `-c <number>`, which specifies the number of bytes to display from the end of the file.
- At least one file must be provided.
- If multiple files are given, print a header (`==> filename <==`) before each file’s output, separated by a blank line.
- Errors must be printed, but the program should continue processing other files.
- Exit with a non‑zero status if any error occurs.

## Instructions
- Directory: `ztail`
- File to submit: `main.go`
- Allowed packages: `os`, `fmt`

## Implementation
`main.go`:
```go
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
```

### Explanation
- Check arguments: must start with `-c` and a positive integer.
- Parse the number of bytes to display.
- Loop through each file:
  - If multiple files, print a header before each file’s output.
  - Read the file with `os.ReadFile`.
  - Print the last `n` bytes (or the whole file if shorter).
  - On error, print the error message and set exit code to `1`.
- Exit with the appropriate status code.

## Usage
Example test program:
```bash
$ go run . -c 4 file1.txt
xyz

$ go run . -c 4 file1.txt file2.txt
==> file1.txt <==
xyz

==> file2.txt <==
xyz
```

Error handling:
```bash
$ go run . -c 4 file1.txt nonexisting1.txt file2.txt nonexisting2.txt
==> file1.txt <==
xyz
open nonexisting1.txt: no such file or directory

==> file2.txt <==
xyz
open nonexisting2.txt: no such file or directory
$ echo $?
1
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with `bufio.Scanner` for line‑based tailing, but here `os.ReadFile` is sufficient for byte‑based tail:
```go
import (
    "fmt"
    "os"
)

func tailFile(filename string, n int) {
    data, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    if len(data) < n {
        fmt.Print(string(data))
    } else {
        fmt.Print(string(data[len(data)-n:]))
    }
}
```

## Skills Practiced
- Command‑line argument parsing
- File reading
- Slice indexing
- Error handling
- Exit codes

## Notes
- This exercise demonstrates how to implement a simplified version of `tail`.
- For production code, prefer streaming with `bufio.Reader` for very large files.

## Resources
- Go `os.ReadFile` — Official Docs (go.dev in Bing)  
- Go `fmt.Sscanf` — Official Docs (go.dev in Bing)  