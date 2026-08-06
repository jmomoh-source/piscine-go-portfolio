# Quest07 — split

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **custom string splitting by a given separator** in Go.  
The task: write a function that receives a string `s` and a separator `sep`, then returns a slice of strings resulting from splitting `s` by `sep`.

Rules:
- If the string does not contain the separator, return a slice with the original string.
- If the string is empty, return an empty slice.
- You must implement the splitting logic manually (no direct use of `strings.Split`).

## Instructions
- File to submit: `split.go`
- Expected function signature:
```go
func Split(s, sep string) []string
```

## Implementation
`split.go`:
```go
package piscine

func Split(s, sep string) []string {
    if sep == "" {
        return []string{s}
    }

    var result []string
    start := 0
    for i := 0; i+len(sep) <= len(s); {
        if s[i:i+len(sep)] == sep {
            result = append(result, s[start:i])
            i += len(sep)
            start = i
        } else {
            i++
        }
    }
    result = append(result, s[start:])
    return result
}
```

### Explanation
- Handle the edge case where `sep` is empty: return the whole string.
- Iterate through the string, checking substrings of length `len(sep)`.
- When a match is found:
  - Append the substring from `start` to `i` into the result slice.
  - Move `i` forward by the length of `sep`.
  - Update `start` to the new position.
- After the loop, append the last substring.
- Return the slice of strings.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    s := "HelloHAhowHAareHAyou?"
    fmt.Printf("%#v\n", piscine.Split(s, "HA"))
}
```

Output:
```text
[]string{"Hello", "how", "are", "you?"}
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with `strings.Split`:
```go
import "strings"

func SplitStd(s, sep string) []string {
    return strings.Split(s, sep)
}
```
⚠️ Note: `strings.Split` is more efficient and idiomatic.  
Your Piscine solution demonstrates manual substring handling for deeper understanding.

## Skills Practiced
- String slicing
- Substring comparison
- Slice manipulation
- Handling edge cases

## Notes
- This exercise demonstrates manual string splitting with a custom separator.
- For production code, prefer `strings.Split` for clarity and performance.

## Resources
- Go `strings.Split` — Official Docs (go.dev in Bing)  
- Effective Go — Strings (go.dev in Bing)  