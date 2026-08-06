# Quest02 — isnegative

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces conditionals in Go.  
The task: write a function `IsNegative` that prints `'T'` (true) if the integer passed as parameter is negative, otherwise it prints `'F'` (false). Each result must be followed by a newline.

## Instructions
- File to submit: `isnegative.go` (inside the `piscine` package)
- Expected function signature:
```go
func IsNegative(nb int) {
}
```

## Output must match exactly:
```
bash
$ go run .
F
F
T
$
```

##Implementation
- isnegative.go:
```
go
package piscine

import "fmt"

func IsNegative(nb int) {
    if nb < 0 {
        fmt.Println("T")
    } else {
        fmt.Println("F")
    }
}
```

## Explanation
- if nb < 0 → checks if the number is negative

- fmt.Println("T") → prints T if true

- fmt.Println("F") → prints F otherwise

- Each call prints on its own line

## Usage
Example test program:
```
go
package main

import "piscine"

func main() {
    piscine.IsNegative(1)
    piscine.IsNegative(0)
    piscine.IsNegative(-1)
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
F
F
T
```

## Skills Practiced
- Go syntax basics (package, func)

- Conditional statements (if / else)

- Printing with the Go standard library (fmt)

- Handling integers and logical checks

## Notes
- Ensure each output is followed by a newline

- Do not add extra spaces or characters

- Keep the function signature exactly as specified