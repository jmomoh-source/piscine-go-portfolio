# Quest03 — ultimatepointone

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise deepens your understanding of pointers in Go.  
The task: write a function `UltimatePointOne` that takes a **pointer to a pointer to a pointer to an int** as an argument and assigns the value `1` to the underlying integer.

## Instructions
- File to submit: `ultimatepointone.go` (inside the `piscine` package)
- Expected function signature:
```go
func UltimatePointOne(n ***int) {
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
ultimatepointone.go:
```
go
package piscine

func UltimatePointOne(n ***int) {
    ***n = 1
}
```

## Explanation
n ***int → declares a parameter that is a pointer to a pointer to a pointer to an integer.

***n = 1 → dereferences three times to reach the underlying integer and assigns it the value 1.

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
    a := 0
    b := &a
    n := &b
    piscine.UltimatePointOne(&n)
    fmt.Println(a)
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
Understanding multiple levels of pointers in Go

Using the * operator repeatedly to dereference values

Passing variables by reference through multiple layers

## Notes
Ensure the function modifies the value through the pointer chain, not by returning a new value.

Output must be exactly as specified, with no extra spaces or characters.