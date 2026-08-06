# Quest05 — isprintable

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise continues character classification in Go.  
The task: write a function `IsPrintable` that checks if a string contains only **printable characters**.

Rules:
- Return `true` if all characters are printable.
- Return `false` otherwise (including control characters like `\n`, `\t`, etc.).

## Instructions
- File to submit: `isprintable.go` (inside the `piscine` package)
- Expected function signature:
```go
func IsPrintable(s string) bool {
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
isprintable.go:
```
go
package piscine

func IsPrintable(s string) bool {
    if s == "" {
        return false
    }
    for _, r := range s {
        if r < 32 || r > 126 {
            return false
        }
    }
    return true
}
```

## Explanation
ASCII printable characters range from code 32 (space ' ') to 126 (~).

Iterate over the string as runes.

If any rune falls outside this range, return false.

If all runes are within the printable range, return true.

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
    fmt.Println(piscine.IsPrintable("Hello"))   // true
    fmt.Println(piscine.IsPrintable("Hello\n")) // false
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
In Go’s standard library, you could achieve similar behavior using unicode.IsPrint:
```
go
import "unicode"

func IsPrintableStd(s string) bool {
    if s == "" {
        return false
    }
    for _, r := range s {
        if !unicode.IsPrint(r) {
            return false
        }
    }
    return true
}
```

- ⚠️ Note: unicode.IsPrint checks all Unicode printable characters, not just ASCII 32–126.
Your manual implementation restricts the check to ASCII printable characters only, as required by the exercise.

## Skills Practiced
Rune iteration in Go

ASCII range checks

Character classification

String validation

## Notes
This implementation is limited to ASCII printable characters.

For broader Unicode support, prefer unicode.IsPrint.

## Resources
Go unicode Package — Official Docs (go.dev in Bing)

Effective Go — Strings and Runes (go.dev in Bing)

ASCII Table Reference (ascii-code.com in Bing)