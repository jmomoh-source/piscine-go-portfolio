# Quest07 — splitwhitespaces

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **string splitting by whitespace** in Go.  
The task: write a function that separates the words of a string and returns them in a slice of strings.

Rules:
- Separators are spaces, tabs, and newlines.
- If the string is empty, return an empty slice.
- Words must be returned in the order they appear.

## Instructions
- File to submit: `splitwhitespaces.go`
- Expected function signature:
```go
func SplitWhiteSpaces(s string) []string
```

## Implementation
`splitwhitespaces.go`:
```go
package piscine

func SplitWhiteSpaces(s string) []string {
    var result []string
    word := ""
    for _, r := range s {
        if r == ' ' || r == '\t' || r == '\n' {
            if word != "" {
                result = append(result, word)
                word = ""
            }
        } else {
            word += string(r)
        }
    }
    if word != "" {
        result = append(result, word)
    }
    return result
}
```

### Explanation
- Iterate through each rune in the string.
- If the rune is a separator (`space`, `tab`, `newline`):
  - If `word` is non‑empty, append it to the result slice and reset `word`.
- Otherwise, add the rune to `word`.
- After the loop, append the last word if it exists.
- Return the slice of words.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
}
```

Output:
```text
[]string{"Hello", "how", "are", "you?"}
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with `strings.Fields`:
```go
import "strings"

func SplitWhiteSpacesStd(s string) []string {
    return strings.Fields(s)
}
```
⚠️ Note: `strings.Fields` automatically splits on any whitespace and handles multiple consecutive separators efficiently.  
Your Piscine solution demonstrates manual splitting logic for deeper understanding.

## Skills Practiced
- Rune iteration
- String building
- Slice manipulation
- Handling separators

## Notes
- This exercise demonstrates manual string splitting.
- For production code, prefer `strings.Fields` for clarity and performance.

## Resources
- Go `strings.Fields` — Official Docs (go.dev in Bing)  
- Effective Go — Strings (go.dev in Bing)  