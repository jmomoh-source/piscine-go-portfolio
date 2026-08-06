# Quest02 — printreversealphabet

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise builds on the previous one by practicing loops in reverse order.  
The task: write a program that prints the Latin alphabet in lowercase **in reverse order** (from `z` to `a`) on a single line, followed by a newline.

## Instructions
- File to submit: `printreversealphabet/main.go`
- Allowed functions: Go standard library (`fmt`)
- Casting is not allowed — use runes directly
- Output must be:

```bash
$ go run .
zyxwvutsrqponmlkjihgfedcba
$
```

## Implementation
main.go:
```
package main

import "fmt"

func main() {
    for ch := 'z'; ch >= 'a'; ch-- {
        fmt.Printf("%c", ch)
    }
    fmt.Println()
}
```

## Explanation:
for ch := 'z'; ch >= 'a'; ch-- → loops backwards from 'z' down to 'a'

fmt.Printf("%c", ch) → prints each rune (character) without casting

fmt.Println() → prints a newline at the end


# Usage
Run the program:
```
go run .
```

## Expected output:
```
zyxwvutsrqponmlkjihgfedcba
```


## Skills Practiced
Go syntax basics (package main, func main)

Using runes ('z' down to 'a')

Reverse iteration with for

Printing with the Go standard library (fmt)


## Notes
- Do not use type casting — rely on rune literals ('a', 'z')

- Ensure the newline is included at the end

- Output must be exactly as specified, with no extra spaces or characters