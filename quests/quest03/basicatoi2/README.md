# Quest03 — basicatoi2

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise extends the previous `basicatoi` task by adding error handling for invalid strings.  
The task: write a function `BasicAtoi2` that simulates the behavior of Go’s `strconv.Atoi`, but simplified.  

Rules:
- Input strings may contain non-digit characters.
- If the string contains any non-digit character, return `0`.
- No need to handle signs (`+` or `-`).
- No need to return errors — only the integer result.

## Instructions
- File to submit: `basicatoi2.go` (inside the `piscine` package)
- Expected function signature:
```go
func BasicAtoi2(s string) int {
}
```

Output must match exactly:
```
bash
$ go run .

12345
12345
0
0
$
```

## Implementation
basicatoi2.go:
```
go
package piscine

func BasicAtoi2(s string) int {
    result := 0
    for _, r := range s {
        if r < '0' || r > '9' {
            return 0
        }
        result = result*10 + int(r-'0')
    }
    return result
}
```

## Explanation
for _, r := range s → iterates over each rune in the string.

if r < '0' || r > '9' → checks if the rune is not a digit; if so, return 0.

result = result*10 + int(r-'0') → builds the integer step by step.

This ensures that only valid digit strings are converted; invalid ones return 0.

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
    fmt.Println(piscine.BasicAtoi2("12345"))
    fmt.Println(piscine.BasicAtoi2("0000000012345"))
    fmt.Println(piscine.BasicAtoi2("012 345"))
    fmt.Println(piscine.BasicAtoi2("Hello World!"))
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
12345
12345
0
0
```

## Skills Practiced
String iteration in Go

Rune arithmetic ('0' offset)

Input validation

Manual string-to-int conversion

## Notes
Do not use strconv.Atoi — the goal is to implement the logic manually.

Any non-digit character invalidates the entire string, returning 0.

## Resources
Go Runes and Unicode — Tour of Go (bing.com in Bing) (bing.com in Bing)

Go Strings — Effective Go (bing.com in Bing) (bing.com in Bing)

Go strconv.Atoi Documentation (bing.com in Bing) (bing.com in Bing)