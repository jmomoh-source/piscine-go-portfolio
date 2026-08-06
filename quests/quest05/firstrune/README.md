# Quest05 — firstrune

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **rune handling** in Go.  
The task: write a function `FirstRune` that returns the first rune of a string.

Rules:
- Strings in Go are UTF‑8 encoded.
- To correctly handle Unicode characters, convert the string into a slice of runes.
- Return the first rune in the slice.

## Instructions
- File to submit: `firstrune.go` (inside the `piscine` package)
- Expected function signature:
```go
func FirstRune(s string) rune {
}
```

Output must match exactly:
```
bash
$ go run .
HSO
$
```

## Implementation
firstrune.go:
```
go
package piscine

func FirstRune(s string) rune {
    runes := []rune(s)
    return runes[0]
}
```

## Explanation
[]rune(s) → converts the string into a slice of runes, ensuring proper handling of Unicode characters.

runes[0] → returns the first rune in the slice.

This approach avoids issues with multi‑byte characters in UTF‑8.

## Usage
Example test program:
```
go
package main

import (
    "piscine"
    "github.com/01-edu/z01"
)

func main() {
    z01.PrintRune(piscine.FirstRune("Hello!"))
    z01.PrintRune(piscine.FirstRune("Salut!"))
    z01.PrintRune(piscine.FirstRune("Ola!"))
    z01.PrintRune('\n')
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
HSO
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same result using direct rune conversion:
```
go
[]rune("Hello!")[0]
```

This shows that your manual implementation mirrors idiomatic Go usage.

## Skills Practiced
Rune handling in Go

UTF‑8 and Unicode awareness

String indexing with runes

Basic text processing

## Notes
Using []rune is essential for correct Unicode handling.

Direct indexing into a string (s[0]) would only return the first byte, not the first rune.

## Resources
Go Runes — Tour of Go (go.dev in Bing)

Effective Go — Strings and Runes (go.dev in Bing)

Unicode and UTF‑8 Basics (unicode.org in Bing)