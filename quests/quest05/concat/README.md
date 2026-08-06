# Quest05 — concat

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **string concatenation** in Go.  
The task: write a function `Concat` that returns the concatenation of two strings passed as arguments.

Rules:
- Concatenate `str1` and `str2` directly.
- Return the resulting string.

## Instructions
- File to submit: `concat.go` (inside the `piscine` package)
- Expected function signature:
```go
func Concat(str1 string, str2 string) string {
}
```

Output must match exactly:
```
bash
$ go run .
Hello! How are you?
$
```

## Implementation
concat.go:
```
go
package piscine

func Concat(str1 string, str2 string) string {
    return str1 + str2
}
```

## Explanation
In Go, the + operator concatenates strings.

Returning str1 + str2 produces the combined string.

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
    fmt.Println(piscine.Concat("Hello!", " How are you?"))
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
Hello! How are you?
```

## Standard Library Equivalent
In Go’s standard library, concatenation is done with the + operator.
For multiple strings, strings.Join is often used:
```
go
import "strings"

result := strings.Join([]string{"Hello!", " How are you?"}, "")
```

Your manual implementation mirrors the simplest idiomatic approach.

## Skills Practiced
String concatenation

Understanding Go’s + operator

Basic string manipulation

## Notes
Concatenation with + is efficient for small numbers of strings.

For joining many strings, prefer strings.Builder or strings.Join for performance.

## Resources
Go Strings — Tour of Go (go.dev in Bing) 

Go strings.Join — Official Docs (go.dev in Bing) 

Effective Go — Strings (go.dev in Bing)