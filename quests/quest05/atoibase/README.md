# Quest05 — atoibase

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **string parsing in arbitrary bases** in Go.  
The task: write a function `AtoiBase` that converts a numeric string `s` into an integer, given a custom base string.

Rules:
- A base must contain at least **2 unique characters**.
- A base must not contain `+` or `-`.
- If the base is invalid, return `0`.
- The string `s` must contain only characters from the base.
- Negative numbers do not need to be handled.

## Instructions
- File to submit: `atoibase.go` (inside the `piscine` package)
- Expected function signature:
```go
func AtoiBase(s string, base string) int {
}
```

Output must match exactly:
```
bash
$ go run .
125
125
125
125
0
$
```

## Implementation
atoibase.go:
```
go
package piscine

func AtoiBase(s string, base string) int {
    if !isValidBase(base) {
        return 0
    }

    baseLen := len(base)
    indexMap := make(map[rune]int)
    for i, r := range base {
        indexMap[r] = i
    }

    result := 0
    for _, r := range s {
        value, ok := indexMap[r]
        if !ok {
            return 0
        }
        result = result*baseLen + value
    }
    return result
}

func isValidBase(base string) bool {
    if len(base) < 2 {
        return false
    }
    seen := map[rune]bool{}
    for _, r := range base {
        if r == '+' || r == '-' || seen[r] {
            return false
        }
        seen[r] = true
    }
    return true
}
```

## Explanation
- Validation: Ensure base length ≥ 2, no duplicates, no + or -.

- Mapping: Build a map of rune → index for fast lookup.

- Conversion: For each character in s, multiply the current result by base length and add the digit value.

- Invalid input: If any character is not in the base, return 0.

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
    fmt.Println(piscine.AtoiBase("125", "0123456789"))
    fmt.Println(piscine.AtoiBase("1111101", "01"))
    fmt.Println(piscine.AtoiBase("7D", "0123456789ABCDEF"))
    fmt.Println(piscine.AtoiBase("uoi", "choumi"))
    fmt.Println(piscine.AtoiBase("bbbbbab", "-ab"))
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
125
125
125
125
0
```

## Standard Library Equivalent
In Go’s standard library, similar behavior can be achieved using strconv.ParseInt for bases 2–36:
```
go
import (
    "fmt"
    "strconv"
)

func main() {
    n, _ := strconv.ParseInt("1111101", 2, 0)
    fmt.Println(n) // 125
}
```

- ⚠️ Note: strconv.ParseInt only supports bases 2–36.
Your manual implementation supports any custom base string, which is more flexible.

## Skills Practiced
Base validation

String parsing

Integer construction from digits

Mapping runes to values

## Notes
This implementation is limited to ASCII characters in the base string.

For broader Unicode support, the same logic applies since runes are used.

## Resources
Go strconv.ParseInt — Official Docs (go.dev in Bing)

Effective Go — Strings and Runes (go.dev in Bing)

ASCII Table Reference (ascii-code.com in Bing)