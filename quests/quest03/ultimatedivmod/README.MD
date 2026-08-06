# Quest03 — ultimatedivmod

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise combines arithmetic operations with pointers in Go.  
The task: write a function `UltimateDivMod` that divides the value pointed to by `a` by the value pointed to by `b`.  
- Store the quotient in the integer that `a` points to.  
- Store the remainder in the integer that `b` points to.  

## Instructions
- File to submit: `ultimatedivmod.go` (inside the `piscine` package)
- Expected function signature:
```go
func UltimateDivMod(a *int, b *int) {
}
```

Output must match exactly:
```
bash
$ go run .
6
1
$
```

## Implementation
ultimatedivmod.go:
```
go
package piscine

func UltimateDivMod(a *int, b *int) {
    div := *a / *b
    mod := *a % *b
    *a = div
    *b = mod
}
```

## Explanation
*a / *b → performs integer division using the dereferenced values of a and b.

*a % *b → computes the remainder.

The quotient is stored back into the variable pointed to by a.

The remainder is stored back into the variable pointed to by b.

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
    a := 13
    b := 2
    piscine.UltimateDivMod(&a, &b)
    fmt.Println(a)
    fmt.Println(b)
}
```

Run it:
```
bash
go run .
```

## Expected output:
```
text
6
1
```

## Skills Practiced
Integer division and modulo operations

Using pointers to modify variables directly

Function design with multiple outputs via pointer parameters

## Notes
Ensure b is not zero to avoid division errors.

Output must be exactly as specified, with no extra spaces or characters.

## Resources
Go Pointers — Tour of Go (go.dev in Bing)

Go Arithmetic Operators — Language Spec (go.dev in Bing)

Effective Go — Pointers (go.dev in Bing)