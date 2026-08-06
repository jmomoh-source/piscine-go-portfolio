# Quest05 — tolower

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **string transformation** in Go.  
The task: write a function `ToLower` that converts each letter in a string to lowercase.

Rules:
- Convert all uppercase Latin letters (`A–Z`) to lowercase (`a–z`).
- Leave other characters unchanged (digits, symbols, spaces, punctuation).

## Instructions
- File to submit: `tolower.go` (inside the `piscine` package)
- Expected function signature:
```go
func ToLower(s string) string {
}
```

Output must match exactly:
```
bash
$ go run .
hello! how are you?
$
```

## Implementation
tolower.go:
```
go
package piscine

func ToLower(s string) string {
    runes := []rune(s)
    for i, r := range runes {
        if r >= 'A' && r <= 'Z' {
            runes[i] = r + 32
        }
    }
    return string(runes)
}
```

## Explanation
Convert the string into a slice of runes for safe Unicode handling.

Check if each rune is in the uppercase ASCII range 'A'–'Z'.

Add 32 to the rune value to convert it to lowercase ('A' → 'a', 'B' → 'b', etc.).

Return the modified slice as a string.

## Usage
Example test program:
```
go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    fmt.Println(piscine.ToLower("Hello! How are you?"))
}
```

Run it:
```
bash
go run .
```

Expected output:
```
text
hello! how are you?
```

## Standard Library Equivalent
In Go’s standard library, the same behavior is provided by:
```
go
import "strings"

result := strings.ToLower("Hello! How are you?")
```

Your manual implementation mirrors strings.ToLower, but demonstrates how it works internally.

## Skills Practiced
Rune iteration in Go

ASCII range checks

String transformation

Understanding Go’s strings.ToLower

## Notes
This implementation is limited to Latin uppercase letters.

For broader Unicode support, prefer strings.ToLower.

## Resources
Go strings.ToLower — Official Docs (go.dev in Bing)

Effective Go — Strings and Runes (go.dev in Bing)

ASCII Table Reference (ascii-code.com in Bing)