# Quest10 — rockandroll

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **conditional logic and string return values** in Go.  
The task: write a function `RockAndRoll` that:
- Returns `"rock\n"` if the number is divisible by 2.
- Returns `"roll\n"` if the number is divisible by 3.
- Returns `"rock and roll\n"` if divisible by both 2 and 3.
- Returns `"error: number is negative\n"` if the number is negative.
- Returns `"error: non divisible\n"` if none of the above conditions are met.

## Instructions
- File to submit: `rockandroll.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func RockAndRoll(n int) string
```

## Implementation
`rockandroll.go`:
```go
package piscine

func RockAndRoll(n int) string {
    if n < 0 {
        return "error: number is negative\n"
    }
    if n%2 == 0 && n%3 == 0 {
        return "rock and roll\n"
    }
    if n%2 == 0 {
        return "rock\n"
    }
    if n%3 == 0 {
        return "roll\n"
    }
    return "error: non divisible\n"
}
```

### Explanation
- Check if `n` is negative first.
- Then check divisibility by both 2 and 3.
- If not both, check divisibility by 2, then by 3.
- If none match, return the non‑divisible error.
- Each return includes a newline `\n`.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    fmt.Print(piscine.RockAndRoll(4))
    fmt.Print(piscine.RockAndRoll(9))
    fmt.Print(piscine.RockAndRoll(6))
    fmt.Print(piscine.RockAndRoll(-5))
    fmt.Print(piscine.RockAndRoll(7))
}
```

Output:
```text
rock
roll
rock and roll
error: number is negative
error: non divisible
```

## Standard Library Equivalent
Go’s standard library does not provide a direct divisibility check function.  
This solution demonstrates how to implement custom conditional logic with modulo operations.

## Skills Practiced
- Conditional branching
- Modulo arithmetic
- String return values
- Error handling via return strings

## Notes
- This exercise reinforces how to handle multiple conditions cleanly.
- The order of checks matters: negative first, then combined divisibility, then individual divisibility.