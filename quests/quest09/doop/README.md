# Quest09 — doop

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **command‑line arithmetic operations** in Go.  
The task: write a program `doop` that takes three arguments — a value, an operator, and another value — and performs the operation.

Rules:
- Operators: `+`, `-`, `*`, `/`, `%`
- Handle division by zero → print `"No division by 0"`
- Handle modulo by zero → print `"No modulo by 0"`
- On invalid operator, invalid values, wrong number of arguments, or overflow → print nothing
- Use `os.Args` for argument parsing

## Instructions
- Directory: `doop`
- File to submit: `main.go`
- Allowed packages: `os`

## Implementation
`main.go`:
```go
package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    args := os.Args[1:]
    if len(args) != 3 {
        return
    }

    a, err1 := strconv.ParseInt(args[0], 10, 64)
    b, err2 := strconv.ParseInt(args[2], 10, 64)
    if err1 != nil || err2 != nil {
        return
    }

    op := args[1]
    var result int64
    switch op {
    case "+":
        result = a + b
    case "-":
        result = a - b
    case "*":
        result = a * b
    case "/":
        if b == 0 {
            fmt.Println("No division by 0")
            return
        }
        result = a / b
    case "%":
        if b == 0 {
            fmt.Println("No modulo by 0")
            return
        }
        result = a % b
    default:
        return
    }

    // Overflow check (within int64 range)
    if result > (1<<63-1) || result < -(1<<63) {
        return
    }

    fmt.Println(result)
}
```

### Explanation
- Parse arguments with `os.Args`.
- Convert values to `int64` using `strconv.ParseInt`.
- Switch on operator:
  - Perform arithmetic.
  - Handle division/modulo by zero with specific messages.
- If operator is invalid or parsing fails, print nothing.
- Check for overflow beyond `int64` range.
- Print result if valid.

## Usage
Example test program:
```bash
$ go run .
$ go run . 1 + 1 | cat -e
2$
$ go run . hello + 1
$ go run . 1 p 1
$ go run . 1 / 0 | cat -e
No division by 0$
$ go run . 1 % 0 | cat -e
No modulo by 0$
$ go run . 9223372036854775807 + 1
$ go run . -9223372036854775809 - 3
$ go run . 9223372036854775807 "*" 3
$ go run . 1 "*" 1
1
$ go run . 1 "*" -1
-1
```

## Standard Library Equivalent
In Go’s standard library, arithmetic is performed directly with operators:
```go
func main() {
    a := 10
    b := 5
    fmt.Println(a + b) // 15
    fmt.Println(a - b) // 5
    fmt.Println(a * b) // 50
    fmt.Println(a / b) // 2
    fmt.Println(a % b) // 0
}
```
⚠️ Note: The Piscine solution demonstrates argument parsing, error handling, and overflow checks, which are essential for building robust CLI tools.

## Skills Practiced
- Command‑line argument parsing
- String to integer conversion
- Arithmetic operations
- Error handling
- Overflow detection

## Notes
- This exercise demonstrates how to build a simplified calculator CLI.
- For production code, consider using `math/big` for arbitrary precision arithmetic.

## Resources
- Go `os.Args` — Official Docs (go.dev in Bing)  
- Go `strconv.ParseInt` — Official Docs (go.dev in Bing) 