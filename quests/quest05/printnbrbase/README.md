# Quest05 — printnbrbase

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **number representation in arbitrary bases** in Go.  
The task: write a function `PrintNbrBase` that prints an integer in a string base passed as a parameter.

Rules:
- A base must contain at least **2 unique characters**.
- A base must not contain `+` or `-`.
- If the base is invalid, print `NV` (Not Valid).
- Negative numbers must be handled correctly (prefix with `-`).
- Use `z01.PrintRune` for output.

## Instructions
- File to submit: `printnbrbase.go` (inside the `piscine` package)
- Expected function signature:
```go
func PrintNbrBase(nbr int, base string) {
}
```

Output must match exactly:
```
bash
$ go run .
125
-1111101
7D
-uoi
NV
$
```

## Implementation
printnbrbase.go:
```
go
package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
    // validate base
    if !isValidBase(base) {
        z01.PrintRune('N')
        z01.PrintRune('V')
        return
    }

    if nbr == 0 {
        z01.PrintRune(rune(base[0]))
        return
    }

    if nbr < 0 {
        z01.PrintRune('-')
        nbr = -nbr
    }

    b := len(base)
    digits := []rune{}
    for nbr > 0 {
        digits = append([]rune{rune(base[nbr%b])}, digits...)
        nbr /= b
    }

    for _, d := range digits {
        z01.PrintRune(d)
    }
}

func isValidBase(base string) bool {
    if len(base) < 2 {
        return false
    }
    seen := map[rune]bool{}
    for _, r := range base {
        if r == '+' || r == '-' || seen[r] {
            return false
        }
        seen[r] = true
    }
    return true
}
```

## Explanation
- Validation: Ensure base length ≥ 2, no duplicates, no + or -.

- Zero case: Print the first character of the base.

- Negative numbers: Print - and convert to positive.

- Conversion: Divide by base length, collect remainders, and map to base characters.

- Output: Print digits in correct order using z01.PrintRune.

## Usage
Example test program:
```
go
package main

import (
    "piscine"
    "github.com/01-edu/z01"
)

func main() {
    piscine.PrintNbrBase(125, "0123456789")
    z01.PrintRune('\n')
    piscine.PrintNbrBase(-125, "01")
    z01.PrintRune('\n')
    piscine.PrintNbrBase(125, "0123456789ABCDEF")
    z01.PrintRune('\n')
    piscine.PrintNbrBase(-125, "choumi")
    z01.PrintRune('\n')
    piscine.PrintNbrBase(125, "aa")
    z01.PrintRune('\n')
}
```

Run it:
```
bash
go run .
```

Expected output:
```
text
125
-1111101
7D
-uoi
NV
```

## Standard Library Equivalent
In Go’s standard library, similar behavior can be achieved using strconv.FormatInt:
```
go
import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println(strconv.FormatInt(125, 2))  // "1111101"
    fmt.Println(strconv.FormatInt(125, 16)) // "7d"
}
```

- ⚠️ Note: strconv.FormatInt supports bases 2–36 only.
Your manual implementation supports any custom base string, which is more flexible.

## Skills Practiced
Base conversion

String validation

Negative number handling

Rune printing with z01.PrintRune

## Notes
Manual base validation ensures correctness.

For production code, prefer strconv.FormatInt when working with standard bases.

## Resources
Go strconv.FormatInt — Official Docs (go.dev in Bing)

Effective Go — Numbers and Strings (go.dev in Bing)

ASCII Table Reference (ascii-code.com in Bing)