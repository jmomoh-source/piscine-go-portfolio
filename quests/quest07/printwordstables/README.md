# Quest07 — printwordstables

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **printing slices of strings line by line** in Go.  
The task: write a function that receives a slice of strings and prints each element on a separate line.

Rules:
- Use `z01.PrintRune` for output.
- Each string in the slice must be printed exactly as it is, followed by a newline.

## Instructions
- File to submit: `printwordstables.go`
- Expected function signature:
```go
func PrintWordsTables(a []string)
```

## Implementation
`printwordstables.go`:
```go
package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
    for _, word := range a {
        for _, r := range word {
            z01.PrintRune(r)
        }
        z01.PrintRune('\n')
    }
}
```

### Explanation
- Iterate through the slice of strings.
- For each string, iterate through its runes and print them with `z01.PrintRune`.
- After each string, print a newline.

## Usage
Example test program:
```go
package main

import "piscine"

func main() {
    a := piscine.SplitWhiteSpaces("Hello how are you?")
    piscine.PrintWordsTables(a)
}
```

Output:
```text
Hello
how
are
you?
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with `fmt.Println`:
```go
import "fmt"

func PrintWordsTablesStd(a []string) {
    for _, word := range a {
        fmt.Println(word)
    }
}
```
⚠️ Note: `fmt.Println` is simpler and more idiomatic.  
Your Piscine solution demonstrates manual rune printing for deeper understanding.

## Skills Practiced
- Slice iteration
- Rune iteration
- Printing with `z01.PrintRune`
- Output formatting

## Notes
- This exercise demonstrates manual printing of slices.
- For production code, prefer `fmt.Println` for clarity and simplicity.

## Resources
- Go `fmt.Println` — Official Docs (go.dev in Bing)  
- Effective Go — Printing (go.dev in Bing)  