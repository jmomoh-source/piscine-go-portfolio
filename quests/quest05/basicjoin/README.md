# Quest05 — basicjoin

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **string concatenation with slices** in Go.  
The task: write a function `BasicJoin` that returns a single concatenated string from the slice of strings passed as arguments.

Rules:
- Concatenate all elements of the slice in order.
- Return the resulting string.

## Instructions
- File to submit: `basicjoin.go` (inside the `piscine` package)
- Expected function signature:
```go
func BasicJoin(elems []string) string {
}
```

Output must match exactly:
```
bash
$ go run .
Hello! How are you?
$
```

## Implementation
basicjoin.go:
```
go
package piscine

func BasicJoin(elems []string) string {
    result := ""
    for _, str := range elems {
        result += str
    }
    return result
}
```

## Explanation
Initialize an empty string result.

Iterate over the slice elems.

Append each string to result.

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
    elems := []string{"Hello!", " How", " are", " you?"}
    fmt.Println(piscine.BasicJoin(elems))
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
Hello! How are you?
```

## Standard Library Equivalent
In Go’s standard library, similar behavior is provided by strings.Join:
```
go
import "strings"

elems := []string{"Hello!", " How", " are", " you?"}
result := strings.Join(elems, "")
```

- ⚠️ Note: strings.Join allows you to specify a separator.
Your manual implementation concatenates without a separator, which matches the exercise requirement.

## Skills Practiced
Slice iteration in Go

String concatenation

Understanding Go’s strings.Join

## Notes
Concatenation with + is fine for small slices.

For performance with large slices, prefer strings.Builder or strings.Join.

## Resources
Go strings.Join — Official Docs (go.dev in Bing)

Effective Go — Strings (go.dev in Bing)

Go strings.Builder — Efficient String Concatenation (go.dev in Bing)