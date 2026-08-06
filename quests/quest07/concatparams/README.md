# Quest07 — concatparams

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **string concatenation with slices** in Go.  
The task: write a function that takes a slice of strings and returns them concatenated into a single string, separated by newline characters (`\n`).

Rules:
- If the slice is empty, return an empty string.
- Use `make` for efficient string building if needed.
- Each argument must be separated by a newline.

## Instructions
- File to submit: `concatparams.go`
- Expected function signature:
```go
func ConcatParams(args []string) string
```

## Implementation
`concatparams.go`:
```go
package piscine

func ConcatParams(args []string) string {
    if len(args) == 0 {
        return ""
    }
    result := ""
    for i, s := range args {
        result += s
        if i < len(args)-1 {
            result += "\n"
        }
    }
    return result
}
```

### Explanation
- Iterate through the slice of strings.
- Concatenate each string to `result`.
- Add a newline (`\n`) between elements, but not after the last one.
- Return the final string.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    test := []string{"Hello", "how", "are", "you?"}
    fmt.Println(piscine.ConcatParams(test))
}
```

Output:
```text
Hello
how
are
you?
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with `strings.Join`:
```go
import "strings"

func ConcatParamsStd(args []string) string {
    return strings.Join(args, "\n")
}
```
⚠️ Note: `strings.Join` is more efficient and idiomatic.  
Your Piscine solution demonstrates manual concatenation for deeper understanding.

## Skills Practiced
- Slice iteration
- String concatenation
- Handling separators
- Efficient string building

## Notes
- This exercise demonstrates manual concatenation logic.
- For production code, prefer `strings.Join` for clarity and performance.

## Resources
- Go `strings.Join` — Official Docs (go.dev in Bing)  
- Effective Go — Strings (go.dev in Bing)  