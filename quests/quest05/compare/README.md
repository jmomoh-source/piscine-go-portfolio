# Quest05 — compare

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **string comparison** in Go.  
The task: write a function `Compare` that behaves like Go’s standard library function `strings.Compare`.

Rules:
- Return `0` if the strings are equal.
- Return `-1` if the first string is lexicographically smaller than the second.
- Return `1` if the first string is lexicographically greater than the second.

## Instructions
- File to submit: `compare.go` (inside the `piscine` package)
- Expected function signature:
```go
func Compare(a, b string) int {
}
```

Output must match exactly:
```
bash
$ go run .
0
-1
1
$
```

## Implementation
compare.go:
```
go
package piscine

func Compare(a, b string) int {
    minLen := len(a)
    if len(b) < minLen {
        minLen = len(b)
    }

    for i := 0; i < minLen; i++ {
        if a[i] < b[i] {
            return -1
        } else if a[i] > b[i] {
            return 1
        }
    }

    if len(a) < len(b) {
        return -1
    } else if len(a) > len(b) {
        return 1
    }
    return 0
}
```

## Explanation
Compare characters one by one until a difference is found.

If one string ends before the other, the shorter string is considered smaller.

If no differences are found and lengths are equal, return 0.

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
    fmt.Println(piscine.Compare("Hello!", "Hello!"))
    fmt.Println(piscine.Compare("Salut!", "lut!"))
    fmt.Println(piscine.Compare("Ola!", "Ol"))
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
0
-1
1
```

## Standard Library Equivalent
In Go’s standard library, the same behavior is provided by:
```
go
import "strings"

result := strings.Compare("Hello!", "Hello!") // returns 0
```

Your manual implementation mirrors strings.Compare, but demonstrates how it works internally.

## Skills Practiced
Lexicographic string comparison

Iteration and conditional logic

Understanding Go’s strings.Compare

## Notes
Lexicographic comparison is based on byte values (UTF‑8 safe for ASCII).

For full Unicode comparison, Go’s strings.Compare is the idiomatic choice.

## Resources
Go strings.Compare — Official Docs (go.dev in Bing) (bing.com in Bing)

Effective Go — Strings (go.dev in Bing) (bing.com in Bing)

Lexicographic Order — MathWorld (mathworld in Bing) (bing.com in Bing)