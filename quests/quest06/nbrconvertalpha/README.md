# Quest06 — nbrconvertalpha

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **numeric to alphabet conversion** in Go.  
The task: write a program that prints the corresponding letter in the `n` position of the Latin alphabet, where `n` is each argument received.

Rules:
- `1 → a`, `2 → b`, …, `26 → z`.
- If `n` is invalid (not an integer or outside 1–26), print a space `" "`.
- A flag `--upper` must be implemented. When used, the program prints the result in uppercase.
- The flag will always be the first argument.

## Instructions
- File to submit: `main.go` (inside the `nbrconvertalpha` folder)
- Expected function signature:
```go
package main

func main() {
}
```

Output examples:
```
bash
$ go run .
$ go run . 8 5 12 12 15 | cat -e
hello$
$ go run . 12 5 7 5 14 56 4 1 18 25 | cat -e
legen dary$
$ go run . 32 86 h | cat -e
   $
$ go run . --upper 8 5 25
HEY$
```

## Implementation
main.go:
```
go
package main

import (
    "os"
    "strconv"
    "github.com/01-edu/z01"
)

func main() {
    args := os.Args[1:]
    upper := false

    if len(args) > 0 && args[0] == "--upper" {
        upper = true
        args = args[1:]
    }

    for _, arg := range args {
        n, err := strconv.Atoi(arg)
        if err != nil || n < 1 || n > 26 {
            z01.PrintRune(' ')
            continue
        }
        r := rune('a' + n - 1)
        if upper {
            r = rune('A' + n - 1)
        }
        z01.PrintRune(r)
    }
    z01.PrintRune('\n')
}
```

## Explanation
Parse arguments with os.Args.

Check if the first argument is --upper; if so, set a flag and skip it.

Convert each argument to an integer with strconv.Atoi.

If invalid or out of range, print a space.

Otherwise, map the number to a letter (1 → 'a', 2 → 'b', …).

If --upper is set, use uppercase letters.

Print results with z01.PrintRune.

## Usage
Example:
```
bash
go run . --upper 8 5 25
```

Output:
```
text
HEY
```

## Standard Library Equivalent
In Go’s standard library, you could achieve similar behavior with fmt.Print:
```
go
import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    args := os.Args[1:]
    upper := false
    if len(args) > 0 && args[0] == "--upper" {
        upper = true
        args = args[1:]
    }

    for _, arg := range args {
        n, err := strconv.Atoi(arg)
        if err != nil || n < 1 || n > 26 {
            fmt.Print(" ")
            continue
        }
        if upper {
            fmt.Printf("%c", 'A'+n-1)
        } else {
            fmt.Printf("%c", 'a'+n-1)
        }
    }
    fmt.Println()
}
```

- ⚠️ Note: fmt.Print is simpler and idiomatic.
Your Piscine solution demonstrates manual rune printing with z01.PrintRune.

## Skills Practiced
Command‑line arguments (os.Args)

Flag handling

Integer parsing (strconv.Atoi)

Rune manipulation

Conditional logic

## Notes
This exercise demonstrates manual conversion and flag parsing.

For production code, prefer fmt.Print for simplicity and readability.

## Resources
Go os.Args — Official Docs (go.dev in Bing)

Go strconv.Atoi — Official Docs (go.dev in Bing)

Effective Go — Command‑line arguments (go.dev in Bing)