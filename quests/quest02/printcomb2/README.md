# Quest02 — printcomb2

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise expands on nested loops and output formatting.  
The task: write a function `PrintComb2` that prints all possible combinations of **two different two‑digit numbers** in ascending order, separated by a comma and a space.  
Rules:
- Each number must be two digits (from `00` to `99`).
- The first number must be strictly less than the second.
- The last combination (`98 99`) must not be followed by a comma and space.

## Instructions
- File to submit: `printcomb2.go` (inside the `piscine` package)
- Expected function signature:
```go
func PrintComb2() {
}
```

## Output must match exactly:
```
bash
$ go run . | cat -e
00 01, 00 02, 00 03, ..., 98 99$
$
```

## Implementation
printcomb2.go:
```
go
package piscine

import "fmt"

func PrintComb2() {
    for i := 0; i <= 98; i++ {
        for j := i + 1; j <= 99; j++ {
            fmt.Printf("%02d %02d", i, j)
            if !(i == 98 && j == 99) {
                fmt.Print(", ")
            }
        }
    }
    fmt.Println()
}
```

## Explanation
- Outer loop: i runs from 0 to 98

- Inner loop: j runs from i+1 to 99

- fmt.Printf("%02d %02d", i, j) → prints both numbers with two digits (leading zero if needed)

- The condition if !(i == 98 && j == 99) prevents printing a trailing comma after the last combination

- fmt.Println() adds the final newline

## Usage
Example test program:
```
go
package main

import "piscine"

func main() {
    piscine.PrintComb2()
}
```

## Run it:
```
bash
go run .
```

## Expected output (truncated):
```
text
00 01, 00 02, 00 03, ..., 98 99
```

## Skills Practiced
- Nested loops in Go

- Integer formatting with leading zeros (%02d)

- Conditional checks to avoid trailing separators

- Output formatting with the Go standard library (fmt)

## Notes
- Ensure no trailing comma and space after the last combination

- Each number must always be printed with two digits

- Output must be on a single line, followed by a newline