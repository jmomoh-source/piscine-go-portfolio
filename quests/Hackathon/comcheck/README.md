# Quest10 — comcheck

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **command‑line argument handling** in Go.  
The task: write a program `comcheck` that:
- Displays `Alert!!!` followed by newline if at least one of the arguments matches:
  - `"01"`
  - `"galaxy"`
  - `"galaxy 01"`
- If none of the arguments match, display nothing.

## Instructions
- File to submit: `comcheck/main.go`
- Allowed functions: `os.Args`, `fmt.Println`
- Expected function signature:
```go
func main()
```

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
    for _, arg := range args {
        if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
            fmt.Println("Alert!!!")
            return
        }
    }
}
```

### Explanation
- Use `os.Args` to access command‑line arguments.
- Skip the first element (`os.Args[0]`) which is the program name.
- Loop through arguments:
  - If any match the required strings, print `"Alert!!!"` and exit.
- If no matches are found, the program prints nothing.

## Usage
Example test runs:
```bash
$ go run . "I" "Will" "Enter" "the" "galaxy"
Alert!!!

$ go run . "galaxy 01" "do" "you" "hear" "me"
Alert!!!

$ go run . "random" "words"
# (no output)
```

## Standard Library Equivalent
Go’s standard library provides `os.Args` for command‑line arguments and `fmt.Println` for printing.  
This solution demonstrates how to combine them for simple argument checking.

## Skills Practiced
- Command‑line argument handling
- String comparison
- Conditional logic
- Program control flow

## Notes
- This exercise reinforces how to build small utilities that react to command‑line input.
- The program exits immediately after printing `Alert!!!` to avoid duplicate outputs.