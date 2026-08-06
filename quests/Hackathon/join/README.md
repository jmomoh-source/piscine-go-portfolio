# Quest10 — join

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **string concatenation with separators** in Go.  
The task: write a function `Join` that concatenates all strings in a slice, separated by the given separator.

Rules:
- Input: a slice of strings and a separator string.
- Output: a single concatenated string.
- Separator must appear between elements, but not at the end.

## Instructions
- File to submit: `join.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func Join(strs []string, sep string) string
```

## Implementation
`join.go`:
```go
package piscine

import "strings"

func Join(strs []string, sep string) string {
    return strings.Join(strs, sep)
}
```

### Explanation
- Use `strings.Join` from the Go standard library.
- It automatically places the separator between elements.
- No trailing separator is added.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    toConcat := []string{"Hello!", " How", " are", " you?"}
    fmt.Println(piscine.Join(toConcat, ":"))
}
```

Output:
```text
Hello!: How: are: you?
```

## Standard Library Equivalent
This solution directly uses `strings.Join`, which is the idiomatic way in Go to concatenate slices of strings with a separator.

## Skills Practiced
- Slice handling
- String concatenation
- Using Go’s `strings` package

## Notes
- `strings.Join` is efficient and avoids manual looping.
- This exercise reinforces how to leverage Go’s standard library for common string operations.