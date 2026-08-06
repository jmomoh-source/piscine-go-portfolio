# Quest10 — stringtointslice

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **string to integer conversion** in Go.  
The task: write a function `StringToIntSlice` that takes a string and returns a slice of integers representing the ASCII values of each character.

Rules:
- Each rune in the string is converted to its integer value.
- Spaces and punctuation are included as their ASCII codes.
- Return the slice of integers.

## Instructions
- File to submit: `stringtointslice.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func StringToIntSlice(str string) []int
```

## Implementation
`stringtointslice.go`:
```go
package piscine

func StringToIntSlice(str string) []int {
    result := []int{}
    for _, r := range str {
        result = append(result, int(r))
    }
    return result
}
```

### Explanation
- Iterate through the string as runes.
- Convert each rune to its integer value.
- Append to the result slice.
- Return the slice.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    fmt.Println(piscine.StringToIntSlice("A quick brown fox jumps over the lazy dog"))
    fmt.Println(piscine.StringToIntSlice("Converted this string into an int"))
    fmt.Println(piscine.StringToIntSlice("hello THERE"))
}
```

Output:
```text
[65 32 113 117 105 99 107 32 98 114 111 119 110 32 102 111 120 32 106 117 109 112 115 32 111 118 101 114 32 116 104 101 32 108 97 122 121 32 100 111 103]
[67 111 110 118 101 114 116 101 100 32 116 104 105 115 32 115 116 114 105 110 103 32 105 110 116 111 32 97 110 32 105 110 116]
[104 101 108 108 111 32 84 72 69 82 69]
```

## Standard Library Equivalent
Go’s standard library does not provide a direct “string to int slice” function.  
This solution demonstrates how to iterate over runes and convert them to integers manually.

## Skills Practiced
- Rune iteration
- Type conversion (`rune` → `int`)
- Slice building
- ASCII/Unicode handling

## Notes
- This exercise reinforces how Go handles strings as UTF‑8 encoded runes.
- Works for both ASCII and Unicode characters.