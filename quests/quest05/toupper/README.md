# Quest05 — toupper

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **string transformation** in Go.  
The task: write a function `ToUpper` that capitalizes each letter in a string.

Rules:
- Convert all lowercase Latin letters (`a–z`) to uppercase (`A–Z`).
- Leave other characters unchanged (digits, symbols, spaces, punctuation).

## Instructions
- File to submit: `toupper.go` (inside the `piscine` package)
- Expected function signature:
```go
func ToUpper(s string) string {
}
```

Output must match exactly:
```
bash
$ go run .
HELLO! HOW ARE YOU?
$
```

## Implementation
toupper.go:
```
go
package piscine

func ToUpper(s string) string {
    runes := []rune(s)
    for i, r := range runes {
        if r >= 'a' && r <= 'z' {
            runes[i] = r - 32
        }
    }
    return string(runes)
}
```

## Explanation
Convert the string into a slice of runes for safe Unicode handling.

Check if each rune is in the lowercase ASCII range 'a'–'z'.

Subtract 32 from the rune value to convert it to uppercase ('a' → 'A', 'b' → 'B', etc.).

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
    fmt.Println(piscine.ToUpper("Hello! How are you?"))
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
HELLO! HOW ARE YOU?
```

## Standard Library Equivalent
In Go’s standard library, the same behavior is provided by:
```
go
import "strings"

result := strings.ToUpper("Hello! How are you?")
```

Your manual implementation mirrors strings.ToUpper, but demonstrates how it works internally.

## Skills Practiced
Rune iteration in Go

ASCII range checks

String transformation

Understanding Go’s strings.ToUpper

## Notes
This implementation is limited to Latin lowercase letters.

For broader Unicode support, prefer strings.ToUpper.

## Resources
Go strings.ToUpper — Official Docs (go.dev in Bing)

Effective Go — Strings and Runes (go.dev in Bing)

ASCII Table Reference (ascii-code.com in Bing)