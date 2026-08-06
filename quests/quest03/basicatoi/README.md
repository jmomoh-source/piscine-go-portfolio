# Quest03 — basicatoi

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces string-to-integer conversion in Go.  
The task: write a function `BasicAtoi` that simulates the behavior of Go’s `strconv.Atoi`, but simplified.  

Rules:
- Input strings will only contain digits (`0–9`).
- No need to handle signs (`+` or `-`).
- No need to return errors — only the integer result.
- If the string is not a valid number, return `0`.  
  *(For this exercise, only valid digit strings will be tested.)*

## Instructions
- File to submit: `basicatoi.go` (inside the `piscine` package)
- Expected function signature:
```go
func BasicAtoi(s string) int {
}
```

Output must match exactly:
```
bash
$ go run .
12345
12345
0
$
```

## Implementation
basicatoi.go:
```
go
package piscine

func BasicAtoi(s string) int {
    result := 0
    for _, r := range s {
        result = result*10 + int(r-'0')
    }
    return result
}
```

## Explanation
for _, r := range s → iterates over each rune in the string.

r - '0' → converts the rune digit into its integer value.

result = result*10 + int(r-'0') → builds the integer step by step, shifting digits left and adding the new one.

If the string is "000000", the result will be 0.

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
    fmt.Println(piscine.BasicAtoi("12345"))
    fmt.Println(piscine.BasicAtoi("0000000012345"))
    fmt.Println(piscine.BasicAtoi("000000"))
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
```

## Skills Practiced
String iteration in Go

Rune arithmetic ('0' offset)

Manual string-to-int conversion

Understanding how Atoi works internally

## Notes
Do not use strconv.Atoi — the goal is to implement the logic manually.

Only valid digit strings will be tested, so no need to handle errors or signs.

## Resources
Go Runes and Unicode — Tour of Go (go.dev in Bing) (bing.com in Bing)

Go Strings — Effective Go (go.dev in Bing) (bing.com in Bing)

Go strconv.Atoi Documentation (go.dev in Bing) (bing.com in Bing)