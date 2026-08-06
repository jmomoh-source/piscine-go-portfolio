# Quest05 — isalpha

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise continues character classification in Go.  
The task: write a function `IsAlpha` that checks if a string contains only **alphanumeric characters** (`A–Z`, `a–z`, `0–9`) or is empty.

Rules:
- Return `true` if the string is empty or contains only alphanumeric characters.
- Return `false` otherwise (symbols, punctuation, spaces, etc. are not allowed).

## Instructions
- File to submit: `isalpha.go` (inside the `piscine` package)
- Expected function signature:
```go
func IsAlpha(s string) bool {
}
```

Output must match exactly:
```
bash
$ go run .
false
true
false
true
$
```

## Implementation
isalpha.go:
```
go
package piscine

func IsAlpha(s string) bool {
    for _, r := range s {
        if !((r >= 'A' && r <= 'Z') ||
             (r >= 'a' && r <= 'z') ||
             (r >= '0' && r <= '9')) {
            return false
        }
    }
    return true
}
```

## Explanation
Iterate over the string as runes.

Check if each rune falls within one of the valid ranges:

'A'–'Z' → uppercase letters.

'a'–'z' → lowercase letters.

'0'–'9' → digits.

If any rune is outside these ranges, return false.

If all runes are valid or the string is empty, return true.

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
    fmt.Println(piscine.IsAlpha("Hello! How are you?")) // false
    fmt.Println(piscine.IsAlpha("HelloHowareyou"))      // true
    fmt.Println(piscine.IsAlpha("What's this 4?"))      // false
    fmt.Println(piscine.IsAlpha("Whatsthis4"))          // true
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
false
true
false
true
```

## Standard Library Equivalent
In Go’s standard library, you could achieve similar behavior using unicode.IsLetter and unicode.IsDigit:
```
go
import "unicode"

func IsAlphaStd(s string) bool {
    for _, r := range s {
        if !(unicode.IsLetter(r) || unicode.IsDigit(r)) {
            return false
        }
    }
    return true
}
```

- ⚠️ Note: unicode.IsLetter and unicode.IsDigit check all Unicode letters and digits, not just Latin A–Z and 0–9.
Your manual implementation restricts the check to Latin alphanumeric characters only, as required by the exercise.

## Skills Practiced
Rune iteration in Go

ASCII range checks

Character classification

String validation

## Notes
This implementation is limited to Latin alphanumeric characters.

For broader Unicode support, prefer unicode.IsLetter and unicode.IsDigit.

## Resources
Go unicode Package — Official Docs (go.dev in Bing)

Effective Go — Strings and Runes (go.dev in Bing)

ASCII Table Reference (ascii-code.com in Bing)