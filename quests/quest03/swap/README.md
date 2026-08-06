# Quest03 — swap

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces pointer manipulation in Go.  
The task: write a function `Swap` that takes two pointers to integers and swaps their contents.

## Instructions
- File to submit: `swap.go` (inside the `piscine` package)
- Expected function signature:
```go
func Swap(a *int, b *int) {
}
```

Output must match exactly:
```
bash
$ go run .
1
0
$
```

## Implementation
swap.go:
```
go
package piscine

func Swap(a *int, b *int) {
    temp := *a
    *a = *b
    *b = temp
}
```

## Explanation
*a and *b → dereference the pointers to access the actual integer values.

temp := *a → temporarily stores the value of a.

*a = *b → assigns the value of b to a.

*b = temp → assigns the stored value back to b, completing the swap.

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
    a := 0
    b := 1
    piscine.Swap(&a, &b)
    fmt.Println(a)
    fmt.Println(b)
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
1
0
```

## Skills Practiced
Using pointers to directly manipulate variable values

Understanding dereferencing (*) in Go

Implementing a classic swap operation

## Notes
Ensure the swap is done using pointers, not by returning new values.

Output must be exactly as specified, with no extra spaces or characters.

## Resources
Go Pointers — Tour of Go (go.dev in Bing) (bing.com in Bing)

Effective Go — Pointers (go.dev in Bing) (bing.com in Bing)

Go Memory Model Basics (go.dev in Bing) (bing.com in Bing)