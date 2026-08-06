# Quest05 — lastrune

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise continues rune handling in Go.  
The task: write a function `LastRune` that returns the **last rune** of a string.

Rules:
- Strings in Go are UTF‑8 encoded.
- Convert the string into a slice of runes to handle Unicode correctly.
- Return the last rune in the slice.

## Instructions
- File to submit: `lastrune.go` (inside the `piscine` package)
- Expected function signature:
```go
func LastRune(s string) rune {
}
Output must match exactly:

bash
$ go run .
!!!
$
Implementation
lastrune.go:

go
package piscine

func LastRune(s string) rune {
    runes := []rune(s)
    return runes[len(runes)-1]
}
Explanation
[]rune(s) → converts the string into a slice of runes, ensuring proper handling of Unicode characters.

runes[len(runes)-1] → returns the last rune in the slice.

This approach avoids issues with multi‑byte characters in UTF‑8.

Usage
Example test program:

go
package main

import (
    "piscine"
    "github.com/01-edu/z01"
)

func main() {
    z01.PrintRune(piscine.LastRune("Hello!"))
    z01.PrintRune(piscine.LastRune("Salut!"))
    z01.PrintRune(piscine.LastRune("Ola!"))
    z01.PrintRune('\n')
}
Run it:

bash
go run .
Expected output:

text
!!!
Standard Library Equivalent
In Go’s standard library, you could achieve the same result using:

go
runes := []rune("Hello!")
last := runes[len(runes)-1]
This shows that your manual implementation mirrors idiomatic Go usage.

Skills Practiced
Rune handling in Go

UTF‑8 and Unicode awareness

String indexing with runes

Basic text processing

Notes
Using []rune is essential for correct Unicode handling.

Direct indexing into a string (s[len(s)-1]) would only return the last byte, not the last rune.

Resources
Go Runes — Tour of Go (go.dev in Bing) (bing.com in Bing)

Effective Go — Strings and Runes (go.dev in Bing) (bing.com in Bing)

Unicode and UTF‑8 Basics (unicode.org in Bing) (bing.com in Bing)