# Quest05 — isupper

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **character classification** in Go.  
The task: write a function `IsUpper` that checks if a string contains only **uppercase Latin letters**.

Rules:
- Return `true` if all characters are uppercase letters (`A–Z`).
- Return `false` otherwise (including empty string, lowercase letters, digits, symbols, spaces).

## Instructions
- File to submit: `isupper.go` (inside the `piscine` package)
- Expected function signature:
```go
func IsUpper(s string) bool {
}
```

## Implementation
isupper.go:
```
go
package piscine

func IsUpper(s string) bool {
    if s == "" {
        return false
    }
    for _, r := range s {
        if r < 'A' || r > 'Z' {
            return false
        }
    }
    return true
}
```

## Explanation
Iterate over the string as runes.

Check if each rune falls within the ASCII range 'A'–'Z'.

If any rune is outside this range, return false.

If all runes are uppercase letters, return true.

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
    fmt.Println(piscine.IsUpper("HELLO"))   // true
    fmt.Println(piscine.IsUpper("Hello"))   // false
    fmt.Println(piscine.IsUpper("123"))     // false
    fmt.Println(piscine.IsUpper(""))        // false
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
true
false
false
false
```

## Standard Library Equivalent
In Go’s standard library, you could achieve similar behavior using unicode.IsUpper:
```
go
import "unicode"

func IsUpperStd(s string) bool {
    if s == "" {
        return false
    }
    for _, r := range s {
        if !unicode.IsUpper(r) {
            return false
        }
    }
    return true
}
```

- ⚠️ Note: unicode.IsUpper checks all Unicode uppercase letters, not just Latin A–Z.
Your manual implementation restricts the check to Latin uppercase letters only, as required by the exercise.

## Skills Practiced
Rune iteration in Go

ASCII range checks

Character classification

String validation

## Notes
This implementation is limited to Latin uppercase letters.

For broader Unicode support, prefer unicode.IsUpper.

## Resources
Go unicode Package — Official Docs (go.dev in Bing)

Effective Go — Strings and Runes (go.dev in Bing)

ASCII Table Reference (ascii-code.com in Bing)