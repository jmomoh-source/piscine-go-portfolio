# Quest05 — nrune

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise continues rune handling in Go.  
The task: write a function `NRune` that returns the **nth rune** of a string.  

Rules:
- Strings in Go are UTF‑8 encoded.
- Convert the string into a slice of runes to handle Unicode correctly.
- If `n` is invalid (negative, zero, or greater than the string length), return `0`.

## Instructions
- File to submit: `nrune.go` (inside the `piscine` package)
- Expected function signature:
```go
func NRune(s string, n int) rune {
}
```

Output must match exactly:
```
bash
$ go run .
la!
$
```

## Implementation
nrune.go:
```
go
package piscine

func NRune(s string, n int) rune {
    runes := []rune(s)
    if n <= 0 || n > len(runes) {
        return 0
    }
    return runes[n-1]
}
```

## Explanation
[]rune(s) → converts the string into a slice of runes, ensuring proper handling of Unicode characters.

Checks if n is valid:

n <= 0 → invalid index, return 0.

n > len(runes) → out of bounds, return 0.

Returns the rune at position n-1 (since Go slices are zero‑indexed).

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
    z01.PrintRune(piscine.NRune("Hello!", 3))
    z01.PrintRune(piscine.NRune("Salut!", 2))
    z01.PrintRune(piscine.NRune("Bye!", -1))
    z01.PrintRune(piscine.NRune("Bye!", 5))
    z01.PrintRune(piscine.NRune("Ola!", 4))
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
la!
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same result using:
```
go
runes := []rune("Hello!")
nth := runes[2] // third rune
```

This shows that your manual implementation mirrors idiomatic Go usage.

## Skills Practiced
Rune handling in Go

UTF‑8 and Unicode awareness

Index validation

String indexing with runes

## Notes
Using []rune is essential for correct Unicode handling.

Direct indexing into a string (s[n]) would only return a byte, not a rune.

Returning 0 for invalid indices ensures safe behavior.

## Resources
Go Runes — Tour of Go (go.dev in Bing) (bing.com in Bing)

Effective Go — Strings and Runes (go.dev in Bing) (bing.com in Bing)

Unicode and UTF‑8 Basics (unicode.org in Bing) (bing.com in Bing)