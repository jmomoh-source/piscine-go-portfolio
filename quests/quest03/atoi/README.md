# Quest03 — atoi

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise extends the previous `basicatoi` tasks by adding support for signs (`+` and `-`) and stricter validation.  
The task: write a function `Atoi` that simulates the behavior of Go’s `strconv.Atoi`, but simplified.  

Rules:
- Input strings may contain non-digit characters → return `0`.
- Handle optional leading `+` or `-` signs.
- No need to return errors — only the integer result.

## Instructions
- File to submit: `atoi.go` (inside the `piscine` package)
- Expected function signature:
```go
func Atoi(s string) int {
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
1234
-1234
0
0
$
```

## Implementation
atoi.go:
```
go
package piscine

func Atoi(s string) int {
    if len(s) == 0 {
        return 0
    }

    sign := 1
    start := 0

    if s[0] == '+' {
        start = 1
    } else if s[0] == '-' {
        sign = -1
        start = 1
    }

    result := 0
    for _, r := range s[start:] {
        if r < '0' || r > '9' {
            return 0
        }
        result = result*10 + int(r-'0')
    }

    return sign * result
}
```

## Explanation
Checks for empty string → returns 0.

Handles leading + or - signs:

+ → positive, skip sign character.

- → negative, skip sign character and set sign = -1.

Iterates through remaining runes:

If any non-digit is found → return 0.

Otherwise, build the integer step by step.

Returns the final result multiplied by the sign.

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
    fmt.Println(piscine.Atoi("12345"))
    fmt.Println(piscine.Atoi("0000000012345"))
    fmt.Println(piscine.Atoi("012 345"))
    fmt.Println(piscine.Atoi("Hello World!"))
    fmt.Println(piscine.Atoi("+1234"))
    fmt.Println(piscine.Atoi("-1234"))
    fmt.Println(piscine.Atoi("++1234"))
    fmt.Println(piscine.Atoi("--1234"))
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
1234
-1234
0
0
```

## Skills Practiced
String iteration in Go

Rune arithmetic ('0' offset)

Input validation

Handling signs in string-to-int conversion

Implementing simplified Atoi logic

## Notes
Do not use strconv.Atoi — the goal is to implement the logic manually.

Any invalid character or malformed sign sequence invalidates the string, returning 0.

## Resources
Go Runes and Unicode — Tour of Go (bing.com in Bing) (bing.com in Bing)

Go Strings — Effective Go (bing.com in Bing) (bing.com in Bing)

Go strconv.Atoi Documentation (bing.com in Bing) (bing.com in Bing)