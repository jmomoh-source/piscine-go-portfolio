# Quest03 — divmod

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise combines arithmetic operations with pointers in Go.  
The task: write a function `DivMod` that divides two integers `a` and `b`, storing the quotient in `div` and the remainder in `mod`.

## Instructions
- File to submit: `divmod.go` (inside the `piscine` package)
- Expected function signature:
```go
func DivMod(a int, b int, div *int, mod *int) {
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
divmod.go:
```
go
package piscine

func DivMod(a int, b int, div *int, mod *int) {
    *div = a / b
    *mod = a % b
}
```

## Explanation
*div = a / b → stores the integer division result in the variable pointed to by div.

*mod = a % b → stores the remainder in the variable pointed to by mod.

This demonstrates how pointers allow functions to modify variables outside their local scope.

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
    var div int
    var mod int
    piscine.DivMod(a, b, &div, &mod)
    fmt.Println(div)
    fmt.Println(mod)
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
6
1
```

## Skills Practiced
Integer division and modulo operations

Using pointers to return multiple results

Function design with multiple outputs

## Notes
Ensure b is not zero to avoid division errors.

Output must be exactly as specified, with no extra spaces or characters.

## Resources
Go Pointers — Official Documentation (go.dev in Bing)

Go Arithmetic Operators (go.dev in Bing)

Effective Go — Pointers (go.dev in Bing)