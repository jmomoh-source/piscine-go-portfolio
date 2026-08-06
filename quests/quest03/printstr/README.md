# Quest03 — printstr

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces string handling in Go.  
The task: write a function `PrintStr` that prints the characters of a string one by one on the screen.

## Instructions
- File to submit: `printstr.go` (inside the `piscine` package)
- Expected function signature:
```go
func PrintStr(s string) {
}
```

Output must match exactly:
```
bash
$ go run . | cat -e
Hello World!$
```

## Implementation
printstr.go:
```
go
package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
    for _, r := range s {
        z01.PrintRune(r)
    }
}
```

## Explanation
for _, r := range s → iterates over each rune in the string.

z01.PrintRune(r) → prints each rune (character) one by one.

This ensures proper handling of Unicode characters, not just ASCII.

## Usage
Example test program:
```
go
package main

import "piscine"

func main() {
    piscine.PrintStr("Hello World!")
}
```

Run it:
```
bash
go run . | cat -e
```

Expected output:
```
text
Hello World!$
```

## Skills Practiced
Iterating over strings in Go

Understanding runes vs. bytes

Printing characters individually

Using external library function z01.PrintRune

## Notes
Strings in Go are UTF‑8 encoded, so iterating with range ensures proper rune handling.

Output must be exactly as specified, with no extra spaces or characters.

## Resources
Go Strings — Tour of Go (go.dev in Bing) (bing.com in Bing)

Go Runes and Unicode (go.dev in Bing) (bing.com in Bing)

Effective Go — Strings and Runes (go.dev in Bing) (bing.com in Bing)