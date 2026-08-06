# Quest10 — jumpover

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **string indexing and conditional logic** in Go.  
The task: write a function `JumpOver` that:
- Returns a string containing every **third character** of the input.
- Always appends a newline `\n` at the end.
- If the input string is empty or has fewer than 3 characters, return just `\n`.

## Instructions
- File to submit: `jumpover.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func JumpOver(str string) string
```

## Implementation
`jumpover.go`:
```go
package piscine

func JumpOver(str string) string {
    if len(str) < 3 {
        return "\n"
    }
    result := []rune{}
    for i := 2; i < len(str); i += 3 {
        result = append(result, rune(str[i]))
    }
    return string(result) + "\n"
}
```

### Explanation
- If the string length is less than 3, return newline.
- Iterate starting at index `2` (the third character, since indexing starts at 0).
- Step by 3 each time to collect every third character.
- Build the result slice of runes.
- Return the result string with a newline appended.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    fmt.Print(piscine.JumpOver("1010101010"))
    fmt.Print(piscine.JumpOver(""))
    fmt.Print(piscine.JumpOver("t w e l v e"))
    fmt.Print(piscine.JumpOver("12"))
}
```

Output:
```text
101
w v

```

## Standard Library Equivalent
Go’s standard library does not provide a direct “every nth character” function.  
This solution demonstrates how to implement custom indexing logic with slices and runes.

## Skills Practiced
- String indexing
- Rune handling
- Conditional logic
- Custom iteration steps

## Notes
- Indexing starts at 0, so the third character is at index 2.
- This exercise reinforces control flow and string manipulation in Go.