# Quest05 — alphacount

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **character classification** in Go.  
The task: write a function `AlphaCount` that counts the number of **Latin alphabet letters** in a string.

Rules:
- Only letters `A–Z` and `a–z` are counted.
- Digits, symbols, spaces, and other Unicode characters are ignored.

## Instructions
- File to submit: `alphacount.go` (inside the `piscine` package)
- Expected function signature:
```go
func AlphaCount(s string) int {
}
```

Output must match exactly:
```
bash
$ go run .
10
$
```

## Implementation
alphacount.go:
```
go
package piscine

func AlphaCount(s string) int {
    count := 0
    for _, r := range s {
        if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
            count++
        }
    }
    return count
}
```

## Explanation
Iterate over the string as runes (for _, r := range s).

- Check if each rune falls within the ASCII ranges:

'A'–'Z' → uppercase letters.

'a'–'z' → lowercase letters.

Increment the counter for valid letters.

Return the total count.

## Usage
Example test program:
```
go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    s := "Hello 78 World!    4455 /"
    nb := piscine.AlphaCount(s)
    fmt.Println(nb)
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
10
```

## Standard Library Equivalent
In Go’s standard library, you could achieve similar behavior using unicode.IsLetter:
```
go
import "unicode"

count := 0
for _, r := range s {
    if unicode.IsLetter(r) {
        count++
    }
}
```

- ⚠️ Note: unicode.IsLetter counts all Unicode letters, not just Latin A–Z.
Your manual implementation restricts the count to Latin alphabet letters only, as required by the exercise.

## Skills Practiced
Rune iteration in Go

ASCII range checks

Character classification

String processing

## Notes
This implementation is limited to Latin letters.

For broader Unicode support, prefer unicode.IsLetter.

## Resources
Go unicode Package — Official Docs (go.dev)

Effective Go — Strings and Runes (go.dev)

ASCII Table Reference (ascii-code.com)