# Quest08 — point

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **structs and pointers** in Go.  
The task: write a program that defines a `point` struct with two integer fields (`x` and `y`), then use a function to set values via a pointer.

Rules:
- Use `z01.PrintRune` for output if printing manually.
- The function `setPoint()` must work with `int`.
- Demonstrate pointer usage with structs.

## Instructions
- Directory: `point`
- File to submit: `main.go`
- Expected function signature:
```go
func setPoint(ptr *point)
```

## Implementation
`main.go`:
```go
package main

import "fmt"

type point struct {
    x int
    y int
}

func setPoint(ptr *point) {
    ptr.x = 42
    ptr.y = 21
}

func main() {
    points := &point{}
    setPoint(points)
    fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
```

### Explanation
- Define a struct `point` with fields `x` and `y`.
- `setPoint` takes a pointer to `point` and sets its fields.
- In `main`, create a pointer to a `point` struct, call `setPoint`, and print the values.

## Usage
Example test program:
```bash
$ go run .
x = 42, y = 21
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same without pointers by returning a struct:
```go
func setPoint() point {
    return point{x: 42, y: 21}
}

func main() {
    p := setPoint()
    fmt.Printf("x = %d, y = %d\n", p.x, p.y)
}
```
⚠️ Note: Returning structs is often simpler, but the Piscine solution demonstrates how to use pointers with structs for deeper understanding.

## Skills Practiced
- Struct definition
- Pointer usage
- Field assignment via pointers
- Printing struct values

## Notes
- This exercise demonstrates how pointers allow direct modification of struct fields.
- For production code, both pointer and value approaches are valid depending on context.

## Resources
- Go Structs — Official Docs (go.dev in Bing)  
- Go Pointers — Effective Go (go.dev in Bing)  