# Quest05 — join

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **string concatenation with separators** in Go.  
The task: write a function `Join` that concatenates all strings in a slice, separated by the given `sep`.

Rules:
- Concatenate all elements of the slice in order.
- Insert the separator `sep` between each element.
- Return the resulting string.

## Instructions
- File to submit: `join.go` (inside the `piscine` package)
- Expected function signature:
```go
func Join(strs []string, sep string) string {
}
```

Output must match exactly:
```
bash
$ go run .
Hello!: How: are: you?
$
```

## Implementation
join.go:
```
go
package piscine

func Join(strs []string, sep string) string {
    if len(strs) == 0 {
        return ""
    }
    result := strs[0]
    for i := 1; i < len(strs); i++ {
        result += sep + strs[i]
    }
    return result
}
```

## Explanation
Handle the empty slice case by returning "".

Start with the first element.

For each subsequent element, append the separator and the element.

Return the final concatenated string.

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
    toConcat := []string{"Hello!", " How", " are", " you?"}
    fmt.Println(piscine.Join(toConcat, ":"))
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
Hello!: How: are: you?
```

## Standard Library Equivalent
In Go’s standard library, the same behavior is provided by strings.Join:
```
go
import "strings"

toConcat := []string{"Hello!", " How", " are", " you?"}
result := strings.Join(toConcat, ":")
```

Your manual implementation mirrors strings.Join, but demonstrates how it works internally.

## Skills Practiced
Slice iteration in Go

String concatenation

Separator handling

Understanding Go’s strings.Join

## Notes
Concatenation with + is fine for small slices.

For performance with large slices, prefer strings.Builder or strings.Join.

## Resources
Go strings.Join — Official Docs (go.dev in Bing)

Effective Go — Strings (go.dev in Bing)

Go strings.Builder — Efficient String Concatenation (go.dev in Bing)