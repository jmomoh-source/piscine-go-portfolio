# Quest05 — index

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **substring search** in Go.  
The task: write a function `Index` that behaves like Go’s standard library function `strings.Index`.

Rules:
- Return the index of the first occurrence of `toFind` in `s`.
- Return `-1` if `toFind` is not found.
- Indexing starts at `0`.

## Instructions
- File to submit: `index.go` (inside the `piscine` package)
- Expected function signature:
```go
func Index(s string, toFind string) int {
}
```

Output must match exactly:
```
bash
$ go run .
2
1
-1
$
```

## Implementation
index.go:
```
go
package piscine

func Index(s string, toFind string) int {
    if toFind == "" {
        return 0
    }
    for i := 0; i <= len(s)-len(toFind); i++ {
        if s[i:i+len(toFind)] == toFind {
            return i
        }
    }
    return -1
}
```

## Explanation
If toFind is empty, return 0 (consistent with strings.Index).

Loop through s up to the point where toFind could still fit.

Compare substrings of length len(toFind) with toFind.

Return the index of the first match.

If no match is found, return -1.

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
    fmt.Println(piscine.Index("Hello!", "l"))
    fmt.Println(piscine.Index("Salut!", "alu"))
    fmt.Println(piscine.Index("Ola!", "hOl"))
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
2
1
-1
```

## Standard Library Equivalent
In Go’s standard library, the same behavior is provided by:
```
go
import "strings"

result := strings.Index("Hello!", "l") // returns 2
```

Your manual implementation mirrors strings.Index, but demonstrates how it works internally.

## Skills Practiced
Substring search

String slicing

Iteration and conditional logic

Understanding Go’s strings.Index

## Notes
This implementation works for ASCII and UTF‑8 strings, since slicing operates on bytes.

For full Unicode substring search, Go’s strings.Index is the idiomatic choice.

## Resources
Go strings.Index — Official Docs (go.dev in Bing) (bing.com in Bing) (bing.com in Bing)

Effective Go — Strings (go.dev in Bing) (bing.com in Bing) (bing.com in Bing)

Substring Search Algorithms (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)