# Quest10 — collatzcountdown

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **mathematical sequence simulation** in Go.  
The task: write a function `CollatzCountdown` that returns the number of steps required to reach `1` using the Collatz sequence rules:
- If `n` is even → divide by 2.
- If `n` is odd → multiply by 3 and add 1.
- Continue until reaching `1`.

Rules:
- If `start` is `0` or negative, return `-1`.

## Instructions
- File to submit: `collatzcountdown.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func CollatzCountdown(start int) int
```

## Implementation
`collatzcountdown.go`:
```go
package piscine

func CollatzCountdown(start int) int {
    if start <= 0 {
        return -1
    }
    steps := 0
    for start != 1 {
        if start%2 == 0 {
            start /= 2
        } else {
            start = 3*start + 1
        }
        steps++
    }
    return steps
}
```

### Explanation
- If `start` is less than or equal to 0, return `-1`.
- Otherwise, repeatedly apply Collatz rules until `start` becomes `1`.
- Count the number of steps taken.
- Return the step count.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    steps := piscine.CollatzCountdown(12)
    fmt.Println(steps) // 9
}
```

Output:
```text
9
```

## Standard Library Equivalent
Go’s standard library does not provide Collatz sequence functions.  
This solution demonstrates how to implement iterative mathematical algorithms manually.

## Skills Practiced
- Looping and conditionals
- Mathematical sequence simulation
- Sentinel values (`-1` for invalid input)
- Algorithmic thinking

## Notes
- The Collatz sequence is also known as the “3n + 1 problem.”
- This exercise reinforces control flow and iterative computation.