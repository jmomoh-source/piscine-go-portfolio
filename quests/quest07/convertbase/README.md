# Quest07 — convertbase

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **base conversion between arbitrary bases** in Go.  
The task: write a function that receives three arguments:

- `nbr`: A string representing a numeric value in a base.
- `baseFrom`: A string representing the base of `nbr`.
- `baseTo`: A string representing the target base.

Rules:
- Only valid bases will be tested.
- Negative numbers will not be tested.
- You must implement the conversion manually (no direct use of `strconv.ParseInt` or `strconv.FormatInt`).

## Instructions
- File to submit: `convertbase.go`
- Expected function signature:
```go
func ConvertBase(nbr, baseFrom, baseTo string) string
```

## Implementation
`convertbase.go`:
```go
package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
    // Step 1: Convert nbr from baseFrom to decimal
    baseFromLen := len(baseFrom)
    value := 0
    for _, r := range nbr {
        digit := indexOfRune(baseFrom, r)
        value = value*baseFromLen + digit
    }

    // Step 2: Convert decimal value to baseTo
    if value == 0 {
        return string(baseTo[0])
    }

    baseToLen := len(baseTo)
    result := ""
    for value > 0 {
        digit := value % baseToLen
        result = string(baseTo[digit]) + result
        value /= baseToLen
    }
    return result
}

func indexOfRune(s string, r rune) int {
    for i, v := range s {
        if v == r {
            return i
        }
    }
    return -1
}
```

### Explanation
- **Step 1:** Convert the input string `nbr` from `baseFrom` into an integer value.
- **Step 2:** Convert that integer into the target base `baseTo`.
- Use helper function `indexOfRune` to map characters to their numeric values.
- Build the result string by repeatedly dividing by the target base length.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    result := piscine.ConvertBase("101011", "01", "0123456789")
    fmt.Println(result)
}
```

Output:
```text
43
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with `strconv`:
```go
import "strconv"

func ConvertBaseStd(nbr, baseFrom, baseTo string) string {
    from := len(baseFrom)
    to := len(baseTo)

    // Parse nbr in baseFrom
    val, _ := strconv.ParseInt(nbr, from, 64)

    // Format val in baseTo
    return strconv.FormatInt(val, to)
}
```
⚠️ Note: `strconv.ParseInt` and `strconv.FormatInt` only support bases 2–36.  
Your Piscine solution demonstrates how to handle arbitrary bases manually.

## Skills Practiced
- Base conversion logic
- String/rune mapping
- Integer arithmetic
- Slice/string building

## Notes
- This exercise demonstrates manual base conversion.
- For production code, prefer `strconv` when bases are within 2–36.

## Resources
- Go `strconv.ParseInt` — Official Docs (go.dev in Bing)  
- Go `strconv.FormatInt` — Official Docs (go.dev in Bing)  