# Quest08 — boolean

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **boolean logic and argument counting** in Go.  
The task: write a program that checks if the number of command‑line arguments is even or odd, then prints the corresponding message.

Rules:
- Use `z01.PrintRune` for output.
- Use `os.Args` to access command‑line arguments.
- Define constants/messages for even and odd cases.

## Instructions
- Directory: `boolean`
- File to submit: `main.go`
- Expected function signatures:
```go
func printStr(s string)
func isEven(nbr int) bool
```

## Implementation
`main.go`:
```go
package main

import (
    "os"
    "github.com/01-edu/z01"
)

const EvenMsg = "I have an even number of arguments"
const OddMsg = "I have an odd number of arguments"

func printStr(s string) {
    for _, r := range s {
        z01.PrintRune(r)
    }
    z01.PrintRune('\n')
}

func isEven(nbr int) bool {
    return nbr%2 == 0
}

func main() {
    lengthOfArg := len(os.Args[1:])
    if isEven(lengthOfArg) {
        printStr(EvenMsg)
    } else {
        printStr(OddMsg)
    }
}
```

### Explanation
- `printStr` prints a string rune by rune using `z01.PrintRune`.
- `isEven` checks if a number is divisible by 2.
- `main` counts the arguments (excluding the program name) and prints the appropriate message.

## Usage
Example test program:
```bash
$ go run . "not" "odd"
I have an even number of arguments

$ go run . "not even"
I have an odd number of arguments
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with `fmt.Println`:
```go
import (
    "fmt"
    "os"
)

func main() {
    lengthOfArg := len(os.Args[1:])
    if lengthOfArg%2 == 0 {
        fmt.Println("I have an even number of arguments")
    } else {
        fmt.Println("I have an odd number of arguments")
    }
}
```
⚠️ Note: `fmt.Println` is simpler and idiomatic.  
Your Piscine solution demonstrates manual rune printing for deeper understanding.

## Skills Practiced
- Boolean logic
- Argument counting
- Rune printing
- Conditional branching

## Notes
- This exercise demonstrates manual boolean checks and printing.
- For production code, prefer `fmt.Println` for clarity and simplicity.

## Resources
- Go `os.Args` — Official Docs (go.dev in Bing)  
- Effective Go — Printing (go.dev in Bing)  