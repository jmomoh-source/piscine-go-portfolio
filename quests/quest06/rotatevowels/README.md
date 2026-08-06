# Quest06 — rotatevowels

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **string manipulation with vowel positions** in Go.  
The task: write a program that checks the arguments for vowels and "mirrors" their positions.

Rules:
- Vowels are `a, e, i, o, u` (case‑sensitive). `y` is not considered a vowel.
- If arguments contain vowels, the program must reverse the order of vowels across the entire input while keeping consonants in place.
- If there are no arguments, print a newline.
- If there are no vowels, print the arguments unchanged.

## Instructions
- File to submit: `main.go` (inside the `rotatevowels` folder)
- Expected function signature:
```go
package main

func main() {
}
```

Output examples:
```
bash
$ go run . "Hello World" | cat -e
Hollo Werld$

$ go run . "HEllO World" "problem solved"
Hello Werld problom sOlvEd

$ go run . "str" "shh" "psst"
str shh psst

$ go run . "happy thoughts" "good luck"
huppy thooghts guod lack

$ go run . "aEi" "Ou"
uOi Ea

$ go run .
$
```

## Implementation
main.go:
```
go
package main

import (
    "os"
    "github.com/01-edu/z01"
)

func isVowel(r rune) bool {
    return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
        r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        z01.PrintRune('\n')
        return
    }

    // Collect all vowels across arguments
    vowels := []rune{}
    for _, arg := range args {
        for _, r := range arg {
            if isVowel(r) {
                vowels = append(vowels, r)
            }
        }
    }

    // Reverse vowels
    for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
        vowels[i], vowels[j] = vowels[j], vowels[i]
    }

    // Replace vowels in order
    vi := 0
    for ai, arg := range args {
        for _, r := range arg {
            if isVowel(r) {
                z01.PrintRune(vowels[vi])
                vi++
            } else {
                z01.PrintRune(r)
            }
        }
        if ai < len(args)-1 {
            z01.PrintRune(' ')
        }
    }
    z01.PrintRune('\n')
}
```

## Explanation
Define isVowel to check if a rune is a vowel.

Collect all vowels from all arguments into a slice.

Reverse the slice of vowels.

Iterate through arguments again, replacing each vowel with the next reversed vowel.

Print consonants unchanged.

Separate arguments with spaces and end with a newline.

## Usage
Example:
```
bash
go run . "happy thoughts" "good luck"
```

Output:
```
text
huppy thooghts guod lack
```

## Standard Library Equivalent
In Go’s standard library, you could achieve similar behavior using strings.Builder for efficient string concatenation:
```
go
import (
    "fmt"
    "os"
    "strings"
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Println()
        return
    }

    // Similar vowel collection and replacement logic,
    // but use strings.Builder instead of rune printing.
    var b strings.Builder
    // ...
    fmt.Println(b.String())
}
```

- ⚠️ Note: strings.Builder is more efficient for string construction.
Your Piscine solution demonstrates manual rune handling with z01.PrintRune.

## Skills Practiced
Command‑line arguments (os.Args)

String manipulation

Rune iteration

Vowel detection and replacement

Slice reversal

## Notes
This exercise demonstrates careful handling of runes and positions.

For production code, prefer strings.Builder for efficiency.

## Resources
Go os.Args — Official Docs (go.dev in Bing)

Go strings.Builder — Official Docs (go.dev in Bing)

Effective Go — Strings and Runes (go.dev in Bing)