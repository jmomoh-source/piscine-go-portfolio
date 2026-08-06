# Quest05 — islower

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise continues character classification in Go.  
The task: write a function `IsLower` that checks if a string contains only **lowercase Latin letters**.

Rules:
- Return `true` if all characters are lowercase letters (`a–z`).
- Return `false` otherwise (including empty string, uppercase letters, digits, symbols, spaces).

## Instructions
- File to submit: `islower.go` (inside the `piscine` package)
- Expected function signature:
```go
func IsLower(s string) bool {
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
islower.go:
```
go
package piscine

func IsLower(s string) bool {
    if s == "" {
        return false
    }
    for _, r := range s {
        if r < 'a' || r > 'z' {
            return false
        }
    }
    return true
}
```

## Explanation
Iterate over the string as runes.

Check if each rune falls within the ASCII range 'a'–'z'.

If any rune is outside this range, return false.

If all runes are lowercase letters, return true.

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
    fmt.Println(piscine.IsLower("hello"))   // true
    fmt.Println(piscine.IsLower("hello!"))  // false
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
In Go’s standard library, you could achieve similar behavior using unicode.IsLower:
```
go
import "unicode"

func IsLowerStd(s string) bool {
    if s == "" {
        return false
    }
    for _, r := range s {
        if !unicode.IsLower(r) {
            return false
        }
    }
    return true
}
```

- ⚠️ Note: unicode.IsLower checks all Unicode lowercase letters, not just Latin a–z.
Your manual implementation restricts the check to Latin lowercase letters only, as required by the exercise.

## Skills Practiced
Rune iteration in Go

ASCII range checks

Character classification

String validation

## Notes
This implementation is limited to Latin lowercase letters.

For broader Unicode support, prefer unicode.IsLower.

## Resources
Go unicode Package — Official Docs (go.dev in Bing)

Effective Go — Strings and Runes (go.dev in Bing)

ASCII Table Reference (ascii-code.com in Bing)