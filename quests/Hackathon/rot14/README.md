# Quest10 — rot14

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **string transformation using ROT14** in Go.  
The task: write a function `Rot14` that shifts each alphabetical character forward by 14 positions, preserving case.  
- `'z'` becomes `'n'`  
- `'Z'` becomes `'N'`  
- Non‑alphabetic characters remain unchanged.

## Instructions
- File to submit: `rot14.go`
- Allowed functions: `github.com/01-edu/z01.PrintRune`, `--allow-builtin`
- Expected function signature:
```go
func Rot14(s string) string
```

## Implementation
`rot14.go`:
```go
package piscine

func Rot14(s string) string {
    result := []rune{}
    for _, r := range s {
        if r >= 'a' && r <= 'z' {
            result = append(result, 'a'+(r-'a'+14)%26)
        } else if r >= 'A' && r <= 'Z' {
            result = append(result, 'A'+(r-'A'+14)%26)
        } else {
            result = append(result, r)
        }
    }
    return string(result)
}
```

### Explanation
- Iterate over each rune in the string.
- If it’s lowercase (`a–z`), shift by 14 within the alphabet.
- If it’s uppercase (`A–Z`), shift by 14 within the alphabet.
- Non‑letters are appended unchanged.
- Return the transformed string.

## Usage
Example test program:
```go
package main

import (
    "piscine"
    "github.com/01-edu/z01"
)

func main() {
    result := piscine.Rot14("Hello! How are You?")
    for _, r := range result {
        z01.PrintRune(r)
    }
    z01.PrintRune('\n')
}
```

Output:
```text
Vszzc! Vck ofs Mci?
```

## Standard Library Equivalent
Go’s standard library does not provide ROT encoders.  
This solution demonstrates how to implement a Caesar‑style cipher manually using rune arithmetic.

## Skills Practiced
- Rune manipulation
- Modular arithmetic
- String transformation
- Case preservation

## Notes
- ROT14 is a variant of the Caesar cipher (ROT13 is more common).
- This exercise reinforces understanding of ASCII ranges and modular arithmetic.