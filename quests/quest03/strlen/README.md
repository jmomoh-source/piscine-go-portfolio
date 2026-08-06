# Quest03 — strlen

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces string length calculation in Go.  
The task: write a function `StrLen` that counts the runes of a string and returns that count.

## Instructions
- File to submit: `strlen.go` (inside the `piscine` package)
- Expected function signature:
```go
func StrLen(s string) int {
}
```

Output must match exactly:
```
bash
$ go run .
12
$
```

## Implementation
strlen.go:
```
go
package piscine

func StrLen(s string) int {
    count := 0
    for range s {
        count++
    }
    return count
}
```

## Explanation
for range s → iterates over each rune in the string.

count++ → increments the counter for each rune encountered.

The function returns the total number of runes, which is the string length in terms of characters (not bytes).

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
    l := piscine.StrLen("Hello World!")
    fmt.Println(l)
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
12
```

## Skills Practiced
Iterating over strings in Go

Understanding runes vs. bytes

Counting characters in UTF‑8 encoded strings

## Notes
Using range ensures proper handling of Unicode characters, not just ASCII.

Do not use built‑in functions like len(s) here — the goal is to practice manual iteration.

## Resources
Go Strings — Tour of Go (go.dev in Bing) (bing.com in Bing)

Go Runes and Unicode (go.dev in Bing) (bing.com in Bing)

Effective Go — Strings and Runes (go.dev in Bing) (bing.com in Bing)