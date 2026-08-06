# Quest02 — printnbr

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces integer handling and output formatting in Go.  
The task: write a function `PrintNbr` that prints an `int` passed as parameter.  
Requirements:
- All possible values of type `int` must be supported.
- You cannot convert to `int64`.
- Output must be printed directly, without extra spaces or formatting.

## Instructions
- File to submit: `printnbr.go` (inside the `piscine` package)
- Expected function signature:
```go
func PrintNbr(n int) {
}
```

## Output must match exactly:
```
bash
$ go run .
-1230123
$
```

## Implementation
printnbr.go:
```
go
package piscine

import "fmt"

func PrintNbr(n int) {
    fmt.Print(n)
}
```

## Explanation
- fmt.Print(n) → prints the integer directly using the Go standard library.

- Handles negative numbers, zero, and positive numbers automatically.

- No need for manual rune arithmetic since fmt takes care of formatting.

## Usage
Example test program:
```
go
package main

import "piscine"

func main() {
    piscine.PrintNbr(-123)
    piscine.PrintNbr(0)
    piscine.PrintNbr(123)
    fmt.Println() // add newline for clarity
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
-1230123
```

## Skills Practiced
- Integer handling in Go

- Printing with the Go standard library (fmt)

- Function definition and usage

## Notes
- Do not use type casting to int64

- Ensure no extra spaces or newlines are added beyond what is required

- Always handle edge cases like 0 and negative numbers