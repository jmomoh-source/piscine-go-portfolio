# Quest05 — trimatoi

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **string parsing and integer conversion** in Go.  
The task: write a function `TrimAtoi` that extracts digits from a string and returns them as an integer.

Rules:
- Ignore non‑digit characters.
- If a `-` sign appears **before any digit**, the result must be negative.
- If no digits are found, return `0`.
- Only work with `int` type (no conversion to `int64`).
- There will never be more than one sign in the test cases.

## Instructions
- File to submit: `trimatoi.go` (inside the `piscine` package)
- Expected function signature:
```go
func TrimAtoi(s string) int {
}
```

Output must match exactly:
```
bash
$ go run .
12345
12345
12345
0
1234
-1234
1234
1234
$
```

## Implementation
trimatoi.go:
```
go
package piscine

func TrimAtoi(s string) int {
    result := 0
    sign := 1
    foundDigit := false

    for _, r := range s {
        if r == '-' && !foundDigit {
            sign = -1
        }
        if r >= '0' && r <= '9' {
            foundDigit = true
            result = result*10 + int(r-'0')
        }
    }

    return result * sign
}
```

## Explanation
Iterate over the string as runes.

If a - sign is encountered before any digit, set sign = -1.

For each digit, build the integer (result = result*10 + digit).

Track whether digits have been found (foundDigit).

Return the final integer multiplied by the sign.

If no digits are found, result remains 0.

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
    fmt.Println(piscine.TrimAtoi("12345"))
    fmt.Println(piscine.TrimAtoi("str123ing45"))
    fmt.Println(piscine.TrimAtoi("012 345"))
    fmt.Println(piscine.TrimAtoi("Hello World!"))
    fmt.Println(piscine.TrimAtoi("sd+x1fa2W3s4"))
    fmt.Println(piscine.TrimAtoi("sd-x1fa2W3s4"))
    fmt.Println(piscine.TrimAtoi("sdx1-fa2W3s4"))
    fmt.Println(piscine.TrimAtoi("sdx1+fa2W3s4"))
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
12345
12345
12345
0
1234
-1234
1234
1234
```

## Standard Library Equivalent
In Go’s standard library, you could achieve similar behavior using strconv.Atoi after filtering digits:
```
go
import (
    "strconv"
    "strings"
)

func TrimAtoiStd(s string) int {
    sign := 1
    digits := ""
    for _, r := range s {
        if r == '-' && digits == "" {
            sign = -1
        }
        if r >= '0' && r <= '9' {
            digits += string(r)
        }
    }
    if digits == "" {
        return 0
    }
    n, _ := strconv.Atoi(digits)
    return sign * n
}
```

- ⚠️ Note: strconv.Atoi is the idiomatic way to convert strings to integers, but your manual implementation shows how to parse digits without relying on helpers.

## Skills Practiced
Rune iteration in Go

String parsing

Integer construction from digits

Handling signs and invalid input

## Notes
This implementation is limited to ASCII digits.

For broader Unicode digit support, prefer unicode.IsDigit.

## Resources
Go strconv.Atoi — Official Docs (go.dev in Bing)

Effective Go — Strings and Runes (go.dev in Bing)

ASCII Table Reference (ascii-code.com in Bing)