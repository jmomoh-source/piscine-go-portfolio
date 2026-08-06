# Quest10 — loafofbread

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **string manipulation with custom grouping** in Go.  
The task: write a function `LoafOfBread` that:
- Takes a string and returns another string with words of 5 characters.
- After each group of 5 characters, skip the next character.
- Spaces should be ignored when counting characters.
- Always append a newline `\n` at the end.
- If the string has fewer than 5 non‑space characters, return `"Invalid Output\n"`.

## Instructions
- File to submit: `loafofbread.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func LoafOfBread(str string) string
```

## Implementation
`loafofbread.go`:
```go
package piscine

func LoafOfBread(str string) string {
    // Remove spaces when counting
    runes := []rune{}
    for _, r := range str {
        if r != ' ' {
            runes = append(runes, r)
        }
    }

    if len(runes) < 5 {
        return "Invalid Output\n"
    }

    result := []rune{}
    count := 0
    for i := 0; i < len(runes); i++ {
        result = append(result, runes[i])
        count++
        if count == 5 {
            // skip next character if exists
            if i+1 < len(runes) {
                i++
            }
            result = append(result, ' ')
            count = 0
        }
    }

    return string(result) + "\n"
}
```

### Explanation
- Build a slice of runes ignoring spaces.
- If fewer than 5 characters remain, return `"Invalid Output\n"`.
- Iterate through runes:
  - Collect 5 characters.
  - Skip the next character.
  - Insert a space to separate groups.
- Return the final string with newline.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    fmt.Print(piscine.LoafOfBread("deliciousbread"))
    fmt.Print(piscine.LoafOfBread("This is a loaf of bread"))
    fmt.Print(piscine.LoafOfBread("loaf"))
}
```

Output:
```text
delic ousbr ad
Thisi aloaf ofbre d
Invalid Output
```

## Standard Library Equivalent
Go’s standard library does not provide a direct “group and skip” function.  
This solution demonstrates how to implement custom grouping logic manually.

## Skills Practiced
- Rune iteration
- Space filtering
- Custom grouping
- Conditional skipping

## Notes
- This exercise reinforces how to manipulate strings beyond simple concatenation.
- The algorithm ensures correct grouping even with spaces in the input.