# Quest05 — capitalize

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **string transformation with word boundaries** in Go.  
The task: write a function `Capitalize` that capitalizes the first letter of each word and lowercases the rest.

Rules:
- A word is defined as a sequence of alphanumeric characters.
- The first letter of each word must be uppercase.
- All other letters in the word must be lowercase.
- Non‑alphanumeric characters act as separators but remain unchanged.

## Instructions
- File to submit: `capitalize.go` (inside the `piscine` package)
- Expected function signature:
```go
func Capitalize(s string) string {
}
```

Output must match exactly:
```
bash
$ go run .
Hello! How Are You? How+Are+Things+4you?
$
```

## Implementation
capitalize.go:
```
go
package piscine

import "unicode"

func Capitalize(s string) string {
    runes := []rune(s)
    newWord := true

    for i, r := range runes {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            if newWord {
                runes[i] = unicode.ToUpper(r)
                newWord = false
            } else {
                runes[i] = unicode.ToLower(r)
            }
        } else {
            newWord = true
        }
    }
    return string(runes)
}
```

## Explanation
Convert the string into a slice of runes for safe Unicode handling.

Use a flag newWord to track when a new word starts.

- If the rune is alphanumeric:

If newWord is true, capitalize it and set newWord = false.

Otherwise, lowercase it.

If the rune is not alphanumeric, reset newWord = true (next alphanumeric starts a new word).

Return the modified slice as a string.

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
    fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))
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
Hello! How Are You? How+Are+Things+4you?
```

## Standard Library Equivalent
In Go’s standard library, similar behavior can be achieved using strings.Title (deprecated) or cases.Title from golang.org/x/text/cases:
```
go
import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.Title("hello world")) // "Hello World"
}
```

- ⚠️ Note: strings.Title capitalizes words but does not strictly lowercase the rest.
Your manual implementation enforces both capitalization of the first letter and lowercasing of the rest, as required by the exercise.

## Skills Practiced
Rune iteration in Go

Word boundary detection

String transformation

Using unicode helpers (IsLetter, IsDigit, ToUpper, ToLower)

## Notes
This implementation works for Unicode letters and digits.

For locale‑aware capitalization, prefer cases.Title from golang.org/x/text.

## Resources
Go unicode Package — Official Docs (go.dev in Bing)

Go strings.Title — Official Docs (go.dev in Bing)

Go golang.org/x/text/cases — Title Case Support (pkg.go.dev in Bing)