# Quest10 — enigma

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **pointer manipulation** in Go.  
The task: write a function `Enigma` that moves values around between multiple pointers of different depths.

Rules:
- `a` into `c`
- `c` into `d`
- `d` into `b`
- `b` into `a`

## Instructions
- File to submit: `enigma.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func Enigma(a ***int, b *int, c *******int, d ****int)
```

## Implementation
`enigma.go`:
```go
package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
    tempA := ***a
    tempB := *b
    tempC := *******c
    tempD := ****d

    ***a = tempB
    *b = tempD
    *******c = tempA
    ****d = tempC
}
```

### Explanation
- Dereference each pointer to get its value.
- Store values temporarily to avoid overwriting.
- Reassign according to the rules:
  - `a` → `c`
  - `c` → `d`
  - `d` → `b`
  - `b` → `a`

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    x := 5
    y := &x
    z := &y
    a := &z

    w := 2
    b := &w

    u := 7
    e := &u
    f := &e
    g := &f
    h := &g
    i := &h
    j := &i
    c := &j

    k := 6
    l := &k
    m := &l
    n := &m
    d := &n

    fmt.Println(***a)     // 5
    fmt.Println(*b)       // 2
    fmt.Println(*******c) // 7
    fmt.Println(****d)    // 6

    piscine.Enigma(a, b, c, d)

    fmt.Println("After using Enigma")
    fmt.Println(***a)     // 2
    fmt.Println(*b)       // 6
    fmt.Println(*******c) // 5
    fmt.Println(****d)    // 7
}
```

Output:
```text
5
2
7
6
After using Enigma
2
6
5
7
```

## Standard Library Equivalent
Go’s standard library does not provide functions for deep pointer manipulation.  
This solution demonstrates how to dereference and reassign values across multiple pointer levels.

## Skills Practiced
- Pointer dereferencing
- Temporary variable usage
- Value reassignment
- Understanding pointer depth

## Notes
- This exercise reinforces how pointers can be chained and manipulated.
- Deep pointer levels are rare in real‑world Go, but useful for learning how references work.