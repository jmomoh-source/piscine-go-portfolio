# Quest02 — printcomb

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces nested loops and conditional logic in Go.  
The task: write a function `PrintComb` that prints all unique combinations of three different digits in ascending order, separated by a comma and a space.  
Rules:
- Digits must be different.
- The first digit must be less than the second, and the second less than the third.
- The last combination (`789`) must not be followed by a comma and space.

## Instructions
- File to submit: `printcomb.go` (inside the `piscine` package)
- Expected function signature:
```go
func PrintComb() {
}
Output must match exactly:

bash
$ go run . | cat -e
012, 013, 014, ..., 789$
$
Implementation
printcomb.go:

go
package piscine

import "fmt"

func PrintComb() {
    for i := '0'; i <= '7'; i++ {
        for j := i + 1; j <= '8'; j++ {
            for k := j + 1; k <= '9'; k++ {
                fmt.Printf("%c%c%c", i, j, k)
                if !(i == '7' && j == '8' && k == '9') {
                    fmt.Print(", ")
                }
            }
        }
    }
    fmt.Println()
}
```

## Explanation
Outer loop: i runs from '0' to '7'

Middle loop: j runs from i+1 to '8'

Inner loop: k runs from j+1 to '9'

This ensures i < j < k and all digits are different

fmt.Printf("%c%c%c", i, j, k) → prints the three digits together

The condition if !(i == '7' && j == '8' && k == '9') prevents printing a trailing comma after the last combination

fmt.Println() adds the final newline

## Usage
Example test program:
```
go
package main

import "piscine"

func main() {
    piscine.PrintComb()
}
```

Run it:
```
bash
go run .
```

## Expected output (truncated):
```
text
012, 013, 014, 015, ..., 789
```

## Skills Practiced
- Nested loops in Go

- Rune arithmetic ('0' to '9')

- Conditional checks to avoid trailing separators

- Output formatting with the Go standard library (fmt)

## Notes
- Ensure no trailing comma and space after the last combination

- Each combination must be exactly three digits

- Output must be on a single line, followed by a newline