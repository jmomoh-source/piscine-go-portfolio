# Quest03 — pointone

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces pointers in Go.  
The task: write a function `PointOne` that takes a pointer to an `int` as an argument and assigns the value `1` to it.

## Instructions
- File to submit: `pointone.go` (inside the `piscine` package)
- Expected function signature:
```go
func PointOne(n *int) {
}
```

Output must match exactly:
```
bash
$ go run .
1
$
```

## Implementation
pointone.go:
```
go
package piscine

func PointOne(n *int) {
    *n = 1
}
```

## Explanation
n *int → declares a parameter that is a pointer to an integer.

*n = 1 → dereferences the pointer and assigns the value 1 to the integer it points to.

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
    n := 0
    piscine.PointOne(&n)
    fmt.Println(n)
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
1
```

## Skills Practiced
Understanding pointers in Go

Using the * operator to dereference and assign values

Passing variables by reference

## Notes
Ensure the function modifies the value through the pointer, not by returning a new value.

Output must be exactly as specified, with no extra spaces or characters.