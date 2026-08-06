# Quest05 — isnumeric

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise continues character classification in Go.  
The task: write a function `IsNumeric` that checks if a string contains only **numerical characters** (`0–9`).

Rules:
- Return `true` if all characters are digits.
- Return `false` otherwise (including empty string, letters, symbols, spaces).

## Instructions
- File to submit: `isnumeric.go` (inside the `piscine` package)
- Expected function signature:
```go
func IsNumeric(s string) bool {
}
```

Output must match exactly:
```
bash
$ go run .
true
false
$
```

## Implementation
isnumeric.go:
```
go
package piscine

func IsNumeric(s string) bool {
    if s == "" {
        return false
    }
    for _, r := range s {
        if r < '0' || r > '9' {
            return false
        }
    }
    return true
}
```

## Explanation
Iterate over the string as runes.

Check if each rune falls within the ASCII range '0'–'9'.

If any rune is outside this range, return false.

If all runes are digits, return true.

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
    fmt.Println(piscine.IsNumeric("010203"))   // true
    fmt.Println(piscine.IsNumeric("01,02,03")) // false
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
```

## Standard Library Equivalent
In Go’s standard library, you could achieve similar behavior using unicode.IsDigit:
```
go
import "unicode"

func IsNumericStd(s string) bool {
    if s == "" {
        return false
    }
    for _, r := range s {
        if !unicode.IsDigit(r) {
            return false
        }
    }
    return true
}
```

- ⚠️ Note: unicode.IsDigit checks all Unicode digits, not just ASCII 0–9.
Your manual implementation restricts the check to ASCII digits only, as required by the exercise.

## Skills Practiced
Rune iteration in Go

ASCII range checks

Character classification

String validation

## Notes
This implementation is limited to ASCII digits.

For broader Unicode support, prefer unicode.IsDigit.

## Resources
Go unicode Package — Official Docs (go.dev in Bing)

Effective Go — Strings and Runes (go.dev in Bing)

ASCII Table Reference (ascii-code.com in Bing)