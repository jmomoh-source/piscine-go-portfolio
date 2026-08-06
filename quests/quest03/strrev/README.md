# Quest03 — strrev

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces string manipulation in Go.  
The task: write a function `StrRev` that reverses a string and returns the reversed result.

## Instructions
- File to submit: `strrev.go` (inside the `piscine` package)
- Expected function signature:
```go
func StrRev(s string) string {
}
```

Output must match exactly:
```
bash
$ go run .
!dlroW olleH
$
```

## Implementation
strrev.go:
```
go
package piscine

func StrRev(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

## Explanation
[]rune(s) → converts the string into a slice of runes to handle Unicode characters correctly.

The loop swaps characters from the start and end moving toward the center.

Finally, string(runes) converts the rune slice back into a string.

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
    s := "Hello World!"
    s = piscine.StrRev(s)
    fmt.Println(s)
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
!dlroW olleH
```

## Skills Practiced
String manipulation in Go

Rune handling for Unicode safety

Implementing reversal algorithms

## Notes
Always use runes instead of bytes to ensure proper handling of multi‑byte characters.

Output must be exactly as specified, with no extra spaces or characters.

## Resources
Go Strings — Tour of Go (go.dev in Bing) (bing.com in Bing)

Go Runes and Unicode (go.dev in Bing) (bing.com in Bing)

Effective Go — Strings and Runes (go.dev in Bing) (bing.com in Bing)